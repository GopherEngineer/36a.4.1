build:
	cd web && yarn && yarn build

start: build
	go run ./cmd/...