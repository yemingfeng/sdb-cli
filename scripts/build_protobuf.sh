rm -rf pkg/protobuf/*.go
protoc --go_out=./ --go-grpc_out=./ ./pkg/sdb-protobuf/*.proto --go-grpc_opt=require_unimplemented_servers=false