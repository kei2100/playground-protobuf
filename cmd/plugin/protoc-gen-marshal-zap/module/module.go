package module

import (
	"fmt"
	"html/template"
	"os"
	"strings"

	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	pgs "github.com/lyft/protoc-gen-star"
)

type module struct {
	*pgs.ModuleBase
	genGoContext pgsgo.Context
	tmpl         *template.Template
}

// New creates a pgs.Module. The module generates code that implements zap.ObjectMarshaler for proto messages.
func New() pgs.Module {
	return &module{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *module) Name() string {
	return "marshal-zap"
}

func (m *module) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.genGoContext = pgsgo.InitContext(c.Parameters())
	m.tmpl = template.Must(template.New("").Parse(tmpl))
}

func (m *module) Execute(targets map[string]pgs.File, packages map[string]pgs.Package) []pgs.Artifact {
	for _, file := range targets {
		pf := &protoFile{
			Path:      m.genGoContext.OutputPath(file).String(),
			GoPackage: m.genGoContext.PackageName(file).String(),
		}

		pms, err := processMessages(file.AllMessages())
		if err != nil {
			m.AddError(fmt.Sprintf("failed to process messages %s: %+v\n", pf.Path, err))
			os.Exit(1)
		}
		pf.ProtoMessages = pms

		m.AddGeneratorTemplateFile(
			m.genGoContext.OutputPath(file).SetExt(".marshal-zap.go").String(),
			m.tmpl,
			pf,
		)
	}
	return m.Artifacts()
}

func processMessages(messages []pgs.Message) ([]*protoMessage, error) {
	pms := make([]*protoMessage, len(messages))

	for i, message := range messages {
		var pm protoMessage

		// .<packageName>.<Message>.<NestedMessage>
		fqName := message.FullyQualifiedName()
		packageName := message.Package().ProtoName().String()
		pm.Name = strings.TrimPrefix(fqName, fmt.Sprintf(".%s.", packageName))
		pm.Name = strings.ReplaceAll(pm.Name, ".", "_")

		pms[i] = &pm
	}

	return pms, nil
}
