#
# setup
#
.PHONY: setup
setup:
	which protoc || brew install protobuf

#
# protoc
#
PROTO_FILES = $(shell find proto -name '*.proto')
TARGET_FILES ?= $(shell find proto -name '*.proto' | perl -pe 's|^proto/(.+)\.proto$$|go/$$1.pb.go|gc')
PLUGIN_OPT ?=

.PHONY: protoc
protoc: bin
	@PLUGIN_OPT=$(PLUGIN_OPT) $(MAKE) $(TARGET_FILES)

.PHONY: protoc.example
protoc.example:
	@PLUGIN_OPT=--example_out=go $(MAKE) protoc

PATH := $(shell echo $$(pwd)/bin/plugin:$$PATH)
.PHONY: $(PROTO_FILES)
go/%.pb.go: proto/%.proto
	@mkdir -p $(@D)
	@PATH=$(PATH) protoc $(strip --go_out=go $(PLUGIN_OPT)) $^
	@[ -f go/github.com/kei2100/playground-protobuf/$@ ] && mv go/github.com/kei2100/playground-protobuf/$@ $@
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
