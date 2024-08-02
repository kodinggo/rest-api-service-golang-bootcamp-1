run:
	modd -f ./.modd/modd.conf

migrate-up:
	go run main.go migrate
	
proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
--go-grpc_opt=paths=source_relative pb/story/*.proto