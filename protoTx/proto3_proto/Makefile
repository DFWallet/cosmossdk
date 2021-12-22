regenerate:
	go install cosmossdk/protoTx/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=cosmossdk/protoTx/proto/test_proto,\
	Mgoogle/protobuf/any.proto=cosmossdk/protoTx/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
