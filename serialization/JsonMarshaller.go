package serialization

import (
	"bytes"
	"encoding/json"
)

type JsonMarshaller struct {
	Indent string
	Format     bool
}

func NewJsonMarshaller() *JsonMarshaller {
	return &JsonMarshaller{
        IndentRune: "\t"
		Format: false,
	}
}

func (m JsonMarshaller) Marshal(v interface{}) ([]byte, error) {
	data, err := json.Marshal(v)

	if err == nil && m.Format {
		var buffer bytes.Buffer
		err = json.Indent(&buffer, data, "", m.Indent)
		data = buffer.Bytes()
	}

	return data, err
}

func (m JsonMarshaller) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
