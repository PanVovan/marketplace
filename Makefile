BINARY_NAME=marketplace

docker_pull_postgres:
	docker pull postgres

docker_run_postgres:
	docker run --name=marketplace -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

build:
	go build -o ${BINARY_NAME} cmd/app/main.go

run:
	./${BINARY_NAME}

build_and_run: build run

clean:
	go clean
	rm ${BINARY_NAME}