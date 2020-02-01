#
# setup
#
.PHONY: setup
setup:
	which protoc || brew install protobuf

#
# protoc
#
TARGETS ?= $(shell find proto -name '*.proto' | perl -pe 's|^proto/(.+)\.proto$$|go/$$1.pb.go|gc')
.PHONY: protoc
protoc:
	@$(MAKE) $(TARGETS)

PATH := $(shell echo $$(pwd)/bin/plugin:$$PATH)
PLUGIN_OPT ?=
go/%.pb.go: proto/%.proto
	@mkdir -p $(@D)
	@PATH=$(PATH) protoc $(strip --go_out=go $(PLUGIN_OPT)) $^
	@[ -f go/github.com/kei2100/playground-protobuf/$@ ] && mv go/github.com/kei2100/playground-protobuf/$@ $@
	@rm -rf go/github.com

#.PHONY: protoc.example
#protoc.example:
#	$(eval TARGETS = $(shell find proto -name '*.proto' | perl -pe 's|^proto/(.+)\.proto$$|go/$$1.proto.dump|gc'))
#	$(eval TARGET_PATTERN = go/%.proto.dump)
#	$(eval PLUGIN_OPT = --example_out=go)
#	@PLUGIN_OPT=$(PLUGIN_OPT) TARGETS='$(TARGETS)' TARGET_PATTERN=$(TARGET_PATTERN) $(MAKE) protoc

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
