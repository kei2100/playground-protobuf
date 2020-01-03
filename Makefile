.PHONY: setup
setup:
	which protoc || brew install protobuf

PROTO_FILES ?= $(shell find defs -name '*.proto' | perl -pe 's|^defs/(.+)\.proto$$|go/$$1.pb.go|gc')
.PHONY: protoc
protoc:
	@$(MAKE) $(PROTO_FILES)

go/%.pb.go: defs/%.proto
	@protoc $^ --go_out=go
	@mkdir -p $(@D)
	@mv go/github.com/kei2100/playground-protobuf/$@ $@
	@rm -rf go/github.com
