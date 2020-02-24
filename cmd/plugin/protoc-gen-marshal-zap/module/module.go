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
		pf.Messages = pms

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
		// .<packageName>.<Message>.<NestedMessage>
		fqName := message.FullyQualifiedName()
		packageName := message.Package().ProtoName().String()
		messageName := strings.ReplaceAll(
			strings.TrimPrefix(fqName, fmt.Sprintf(".%s.", packageName)),
			".", "_",
		)

		pfs, err := processFields(message.Fields())
		if err != nil {
			return nil, err
		}

		pms[i] = &protoMessage{
			Name:   messageName,
			Fields: pfs,
		}
	}
	return pms, nil
}

var reservedKeywords = map[string]struct{}{
	"Reset": {},
	"String": {},
	"ProtoMessage": {},
	"Descriptor": {},
}

func processFields(fields []pgs.Field) ([]*protoField, error) {
	pfs := make([]*protoField, len(fields))

	// TODO field mask

	for i, field := range fields {
		pf := protoField{}
		pf.Name = field.Name().String()

		pf.Accessor = field.Name().UpperCamelCase().String()
		if _, ok := reservedKeywords[pf.Accessor]; ok {
			pf.Accessor += "_"
		}
		if field.InOneOf() {
			pf.Accessor = fmt.Sprintf("Get%s()", pf.Accessor)
		}

		pf.Type = protoType(field.Type().ProtoType().Proto())
		pf.IsRepeated = field.Type().IsRepeated()
		pf.IsMap = field.Type().IsMap()

		if pf.IsMap {
			pf.MapType = &mapType{
				KeyType:   keyType{protoType: protoType(field.Type().Key().ProtoType().Proto())},
				ValueType: protoType(field.Type().Element().ProtoType().Proto()),
			}
		}

		pfs[i] = &pf
	}
	return pfs, nil
}
