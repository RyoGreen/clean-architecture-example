build:
	cd assets && npm run build

watch:
	cd assets && npm run dev

server-start:
	go run cmd/app/app.go
