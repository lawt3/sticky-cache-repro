temporal: ## Starts temporalite (does this result in different behaviour?)
	temporalite start --namespace default

prometheus: ## Start Prometheus
	docker run --rm -p 9090:9090 --name prometheus prom/prometheus