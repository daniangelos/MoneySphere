run: docker-down docker-rm docker-build docker-up

docker-down:
	docker compose down

docker-rm:
	docker image rm moneysphere 2>/dev/null ; true

docker-build:
	docker build . -t moneysphere

docker-up:
	docker compose up