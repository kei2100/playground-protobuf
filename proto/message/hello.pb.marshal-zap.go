// Code generated by protoc-gen-marshal-zap. DO NOT EDIT.
// source: github.com/kei2100/playground-protobuf/proto/message/hello.pb.go

package message

import (
	"go.uber.org/zap/zapcore"
	"strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = zapcore.NewNopCore
var _ = strconv.FormatInt

func (m *Hello) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if m == nil {
		return nil
	}

	enc.AddString("message", m.Message)

	return nil
}
