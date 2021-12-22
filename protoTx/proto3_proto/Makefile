regenerate:
	go install github.com/DFWallet/cosmossdk/protoTx/protoc-min-version
	protoc-min-version --version="3.0.0" --gogo_out=\
	Mtest_proto/test.proto=github.com/DFWallet/cosmossdk/protoTx/proto/test_proto,\
	Mgoogle/protobuf/any.proto=github.com/DFWallet/cosmossdk/protoTx/types:. \
	--proto_path=../../protobuf:../:. proto3.proto
	
