#
# setup
#
.PHONY: setup
setup:
	which protoc || brew install protobuf
	go get golang.org/x/tools/cmd/goimports

fmt:
	@goimports -w .

#
# protoc
#
PBGO_FILES ?= $(shell find proto -type f -name '*.proto' | perl -pe 's/proto$$/pb.go/gc')
BIN_FILES ?= $(shell  find cmd/plugin -type d -maxdepth 1 -mindepth 1 | perl -pe 's|^cmd/|bin/|gc')
BIN_PREREQ_FILES ?= $(shell find cmd -name '*.go')
PLUGIN_OPT ?=

.PHONY: protoc
protoc:
	@PLUGIN_OPT=$(PLUGIN_OPT) $(MAKE) $(PBGO_FILES)

.PHONY: protoc-gen-example
protoc-gen-example:
	@$(MAKE) bin/plugin/$@
	@PLUGIN_OPT=--example_out=proto $(MAKE) protoc

.PHONY: protoc-gen-marshal-zap
protoc-gen-marshal-zap:
	@$(MAKE) bin/plugin/$@
	@PLUGIN_OPT=--marshal-zap_out=:./proto $(MAKE) protoc

# TODO fix 依存関係 - types.pb.goを削除しないとtypes.pb.marshal-zap.goを作成できないなど
proto/%.pb.go: proto/%.proto $(BIN_PREREQ_FILES) Makefile
	@mkdir -p $(@D)
	@PATH=$(shell echo $$(pwd)/bin/plugin:$$PATH) protoc $(strip --go_out=proto $(PLUGIN_OPT)) proto/$*.proto
	@find proto/github.com/kei2100/playground-protobuf -type f -name '*.go' | xargs -n 1 -I {} mv {} $(@D)
	@rm -rf proto/github.com

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
