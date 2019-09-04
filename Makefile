build:
	go build -o package/api/api src/cmd/api/main.go
	go build -o package/cmdmanager/cmdmanager src/cmd/cmdmanager/main.go
	go build -o package/nats/nats src/cmd/nats/main.go
	go build -o package/migration/migration src/cmd/nats/main.go
up:
	docker-compose up
upb:
	docker-compose up --build
stop:
	docker-compose stop
upd:
	docker-compose up -d
