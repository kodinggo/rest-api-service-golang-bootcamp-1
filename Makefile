run:
	modd -f ./.modd/modd.conf

migrate-up:
	go run main.go migrate