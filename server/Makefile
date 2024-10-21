APP_NAME = tutorlink-api
CMD_DIR = cmd

run:
	go run $(CMD_DIR)/api/main.go

build:
	go build -o ${APP_NAME} ${CMD_DIR}/api/main.go

seed:
	go run $(CMD_DIR)/seeder/seeder.go