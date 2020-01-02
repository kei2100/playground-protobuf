.PHONY: setup
setup:
	which protoc || brew install protobuf

.PHONY: protoc

go/message/%.pb.go: defs/message/%.proto
	protoc $^ --go_out=go
	mkdir -p $(@D)
	mv go/github.com/kei2100/playground-protobuf/$@ $@
	rm -rf go/github.com
