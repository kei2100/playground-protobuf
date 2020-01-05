#
# setup
#
.PHONY: setup
setup:
	which protoc || brew install protobuf

#
# protoc
#
PBGO_FILES ?= $(shell find proto -name '*.proto' | perl -pe 's|^proto/(.+)\.proto$$|go/$$1.pb.go|gc')
.PHONY: protoc
protoc:
	@$(MAKE) $(PBGO_FILES)

PATH := $(shell echo $$(pwd)/bin/plugin:$$PATH)
go/%.pb.go: proto/%.proto
	@mkdir -p $(@D)
	@PATH=$(PATH) protoc --go_out=go --example_out=go $^
	@mv go/github.com/kei2100/playground-protobuf/$@ $@
	@rm -rf go/github.com

#
# bin
#
BIN_FILES ?= $(shell find cmd -name 'main.go' | perl -pe 's|^cmd/(.+)/main.go$$|bin/$$1|gc')
.PHONY: bin
bin:
	@$(MAKE) $(BIN_FILES)

bin/%: cmd/%/main.go
	@mkdir -p bin/$(*D)
	@cd cmd/$* && go build
	@mv cmd/$*/$(*F) $@
