grpcurl -import-path ./protobuf/authPb -proto auth.proto -plaintext -d '{"token": "Bearer jwt1234"}' localhost:9000 AuthGrpcService/VerifyToken
