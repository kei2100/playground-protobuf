#
# setup
#
.PHONY: setup
setup:
	which protoc || brew install protobuf

#
# protoc
#
PBGO_FILES ?= $(shell find proto -type f -name '*.proto' | perl -pe 's|^proto/(.+)\.proto$$|go/$$1.pb.go|gc')
BIN_FILES ?= $(shell  find cmd/plugin -type d -maxdepth 1 -mindepth 1 | perl -pe 's|^cmd/|bin/|gc')
BIN_PREREQ_FILES ?= $(shell find cmd -name '*.go')
PLUGIN_OPT ?=

.PHONY: protoc
protoc:
	@PLUGIN_OPT=$(PLUGIN_OPT) $(MAKE) $(PBGO_FILES)

.PHONY: protoc.example
protoc-gen-example:
	@$(MAKE) bin/plugin/$@
	@PLUGIN_OPT=--example_out=go $(MAKE) protoc

.PHONY: protoc.marshal-zap
protoc-gen-marshal-zap:
	@$(MAKE) bin/plugin/$@
	@PLUGIN_OPT=--marshal-zap_out=go $(MAKE) protoc

go/%.pb.go: proto/%.proto $(BIN_PREREQ_FILES)
	@mkdir -p $(@D)
	@PATH=$(shell echo $$(pwd)/bin/plugin:$$PATH) protoc $(strip --go_out=go $(PLUGIN_OPT)) proto/$*.proto
	@[ -f go/github.com/kei2100/playground-protobuf/$@ ] && mv go/github.com/kei2100/playground-protobuf/$@ $@
	@rm -rf go/github.com

#
# bin
#
.PHONY: bin
bin:
	@$(MAKE) $(BIN_FILES)

bin/%: $(BIN_PREREQ_FILES)
	@mkdir -p bin/$(*D)
	@cd cmd/$* && go build
	@mv cmd/$*/$(*F) $@
