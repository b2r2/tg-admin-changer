BIN := "./bin/tg_admin_changer_bot"
IMAGE := "tg_admin_changer_bot:dev"


#build:
#	docker build -t $(IMAGE) --secret id=TOKEN,env=$TOKEN -f Dockerfile .
build:
	DOCKER_BUILDKIT=1 docker build --secret id=TOKEN,env=TOKEN -t $(IMAGE) -f Dockerfile .

run:
	go run cmd/tg-admin-changer/main.go

PHONY: build run