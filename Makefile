docker-run:
	docker run --network=host auth-microservice
docker-build:
	docker build --tag auth-microservice .