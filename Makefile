
# Image URL to use all building/pushing image targets
IMG ?= istioext:latest


docker-build: ## Build docker image with the manager.
	docker build -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}