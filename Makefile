server:
	go run main.go

proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        proto/*.proto
clean:
	rm pb/*.go
evans:
	evans --host localhost --port 9090 -r repl
.PHONY: server proto