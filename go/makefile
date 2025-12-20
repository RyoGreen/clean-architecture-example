build:
	cd assets && npm run build

watch:
	cd assets && npm run dev

local-server-start:
	go run cmd/app/app.go

setup-server: build
	docker-compose up -d
