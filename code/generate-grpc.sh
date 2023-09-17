# pentru server in go
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./app/protobuf/*.proto

# pentru client in dart
protoc --dart_out=grpc:../../app/dart_protobuf ./alfie_api.proto && \