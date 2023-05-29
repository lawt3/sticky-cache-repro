### Steps to run this sample:
1) Start Temporal: `temporalite start --namespace default`.
2) Start Prometheus: `docker compose up`
3) Run the following command to start the worker
```
go run worker/main.go
```
3) Run the following command to start the workflow
```
go run starter/main.go
```
