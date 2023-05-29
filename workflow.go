package repro

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"

	// TODO(cretz): Remove when tagged
	_ "go.temporal.io/sdk/contrib/tools/workflowcheck/determinism"
)

func Workflow(ctx workflow.Context, name string) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Workflow started", "name", name)

	for i := 0; i < 100; i++ {
		var result string
		err := workflow.ExecuteActivity(ctx, Activity, name).Get(ctx, &result)
		if err != nil {
			logger.Error("Activity failed.", "Error", err)
			return err
		}
		logger.Info("Activity succeeded.", "Result", result)
	}

	logger.Info("Workflow continuing as new")
	return workflow.NewContinueAsNewError(ctx, Workflow, "world")
	// logger.Info("Workflow finished.")
	// return nil
}

func Activity(ctx context.Context, name string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "name", name)
	return "Hello " + name + "! " + strconv.Itoa(rand.Intn(10)), nil
}
