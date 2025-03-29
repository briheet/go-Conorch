run: build
	@./cmd/bin

build: | cmd
	@go build -o cmd/bin .

cmd: 
	@mkdir -p cmd

docker-up:
	@docker compose -f docker-compose.yml up --detach

docker-down:
	@docker compose down 

hyper:
	@hyperfine --runs 100 ./cmd/bin
