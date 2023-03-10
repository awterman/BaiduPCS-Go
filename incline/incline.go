package incline

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

const (
	prefix = "\nINC-PROTO-BEGIN\n"
	suffix = "\nINC-PROTO-END\n"
)

func Encode[T string | []byte](data T) T {
	var variant any = data

	var encodedData string
	switch v := variant.(type) {
	case string:
		encodedData = prefix + v + suffix
	case []byte:
		encodedData = prefix + string(v) + suffix
	}
	return T(encodedData)
}

func Print[T string | []byte](data T) {
	fmt.Print(Encode(data))
}

type Session struct {
	Id string
}

func NewSession(id string) *Session {
	return &Session{
		Id: id,
	}
}

// NewSessionWithUUID creates a new session with a random UUID.
func NewSessionWithUUID() *Session {
	return NewSession(uuid.NewString())
}

func (s *Session) Encode(data any) ([]byte, error) {
	x := struct {
		Id   string `json:"id"`
		Data any    `json:"data"`
	}{
		Id:   s.Id,
		Data: data,
	}

	variant, err := json.Marshal(x)
	if err != nil {
		return nil, err
	}
	return Encode(variant), nil
}

func (s *Session) Print(data any) error {
	b, err := s.Encode(data)
	if err != nil {
		return err
	}
	Print(b)
	return nil
}
