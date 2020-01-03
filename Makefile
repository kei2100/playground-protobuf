.PHONY: setup
setup:
	which protoc || brew install protobuf

PROTO_FILES ?= $(shell find defs -name '*.proto' | perl -pe 's|^defs/(.+)\.proto$$|go/$$1.pb.go|gc')
.PHONY: protoc
protoc:
	@$(MAKE) $(PROTO_FILES)

go/%.pb.go: defs/%.proto
	@mkdir -p $(@D)
	@protoc $^ --go_out=go
	@mv go/github.com/kei2100/playground-protobuf/$@ $@
	@rm -rf go/github.com

BIN_FILES ?= $(shell find cmd -name 'main.go' | perl -pe 's|^cmd/(.+)/main.go$$|bin/$$1|gc')
.PHONY: bin
bin:
	@$(MAKE) $(BIN_FILES)

bin/%: cmd/%/main.go
	@mkdir -p bin/$(*D)
	@cd cmd/$* && go build
	@mv cmd/$*/$(*F) $@
