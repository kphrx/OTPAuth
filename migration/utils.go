package migration

import (
	"encoding/base64"

	"google.golang.org/protobuf/proto"
)

func DecodeString(d string) (*Migration, error) {
	b, err := base64.StdEncoding.DecodeString(d)
	if err != nil {
		return nil, err
	}

	return Decode(b)
}

func Decode(msg []byte) (*Migration, error) {
	m := &Migration{}
	err := proto.Unmarshal(msg, m)
	return m, err
}

func (m *Migration) EecodeToString() (string, error) {
	b, err := proto.Marshal(m)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(b), nil
}

func (m *Migration) Encode() ([]byte, error) {
	return proto.Marshal(m)
}
