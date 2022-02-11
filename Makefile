docker-run:
	docker run --publish 5000:5000 auth-microservice
docker-build:
	docker build --tag auth-microservice .