include .env
export

service-deploy:
	@docker compose up -d app

service-undeploy:
	@docker compose down 

all-container-clear:
	@docker container prune
