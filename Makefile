build:
	@docker build -t troydai/echoserver:latest .

run: build
	@docker run -p 8081:8081 -d troydai/echoserver:latest

deploy:
	kubectl apply --namespace pilot01 -f ./kubernetes/deployment.yaml

undeploy:
	kubectl delete --namespace pilot01 -f ./kubernetes/deployment.yaml

