start_back:
	go run ./cmd/api/

start_front:
	go run ./cmd/web/

start: start_back start_front