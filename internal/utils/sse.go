package utils

import (
	"fmt"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type SSEMarshaler struct {
	runtime.JSONPb
}

func NewSSEMarshaler() *SSEMarshaler {
	return &SSEMarshaler{
		JSONPb: runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		},
	}
}

func (m *SSEMarshaler) ContentType(_ interface{}) string {
	return "text/event-stream"
}

func (m *SSEMarshaler) Marshal(v interface{}) ([]byte, error) {
	if v == nil {
		return nil, nil
	}

	var msg proto.Message
	switch tv := v.(type) {
	case proto.Message:
		msg = tv
	default:
		return nil, fmt.Errorf("unexpected type %T", v)
	}

	jsData, err := m.JSONPb.Marshal(msg)
	if err != nil {
		return nil, err
	}

	return []byte(fmt.Sprintf("data: %s\n\n", jsData)), nil
}
