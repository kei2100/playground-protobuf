package module

type protoFile struct {
	Path          string
	GoPackage     string
	ProtoMessages []*protoMessage
}

type protoMessage struct {
	Name string
}

const tmpl = `// Code generated by protoc-gen-marshal-zap. DO NOT EDIT.
// source: {{ .Path }}

package {{ .GoPackage }}

import  (
	"go.uber.org/zap/zapcore"
)

{{ range .ProtoMessages }}
func (m *{{ .Name }}) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if u == nil {
		return nil
	}

	return nil
}
{{ end }}
`
