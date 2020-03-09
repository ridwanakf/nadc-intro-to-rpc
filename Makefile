# go build command for rpc server
build-server:
	@echo " >> building binaries for RPC server"
	@go build -v -o server/server server/app.go

# go run command for rpc server
run-server: build-server
	./server/server

# go build command for rpc client
build-client:
	@echo " >> building binaries for RPC client"
	@go build -v -o client/client client/app.go

# go run command for rpc server
run-client: build-client
	./client/client

