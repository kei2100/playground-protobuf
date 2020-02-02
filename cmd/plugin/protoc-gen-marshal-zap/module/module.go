package module

import (
	"fmt"
	"os"

	pgs "github.com/lyft/protoc-gen-star"
)

type module struct {
	*pgs.ModuleBase
}

// New creates a pgs.Module. The module generates code that implements zap.ObjectMarshaler for proto messages.
func New() pgs.Module {
	return &module{
		&pgs.ModuleBase{},
	}
}

func (m *module) Name() string {
	return "zap"
}

func (m *module) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, file := range targets {
		fmt.Fprintln(os.Stderr, file.Name())
	}
	return m.Artifacts()
}
