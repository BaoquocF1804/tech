
proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        --openapiv2_out=openapiv2\
		proto/*.proto
clean:
	rm pb/*.go
evans:
	evans --host localhost --port 9090 -r repl
server:
	go run cmd/server/main.go -port 8080
rest:
	go run cmd/server/main.go -port 8081 -type rest -endpoint 0.0.0.0:8080
client:
	go run cmd/client/main.go -address 0.0.0.0:8080
.PHONY: server proto