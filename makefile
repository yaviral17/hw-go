run:
	@go run main.go

start-db:
	@sudo docker compose up -d

resume-db:
	@sudo docker start hw-go-db-1

stop-db:
	@sudo docker stop hw-go-db-1

remove-db:
	@sudo docker rm hw-go-db-1

logs-db:
	@sudo docker logs hw-go-db-1

