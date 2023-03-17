package incproto

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

type Type string

const (
	TypeEvent Type = "event"
)

type Message struct {
	SessionId string `json:"session_id,omitempty"`

	Type Type   `json:"type"`
	Name string `json:"name"`

	Data json.RawMessage `json:"data"`
}

func (m *Message) Encode() string {
	var s string

	b, err := json.Marshal(m)
	if err != nil {
		s = fmt.Sprintf("error: %v", err)
	} else {
		s = string(b)
	}

	return Encode(s)
}

func (m *Message) String() string {
	return m.Encode()
}

func (m *Message) Print() {
	fmt.Print(m.Encode())
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

func (s *Session) Message(eventType Type, name string, data json.RawMessage) *Message {
	return &Message{
		SessionId: s.Id,

		Type: eventType,
		Name: name,

		Data: data,
	}
}

func (s *Session) Event(name string, data json.RawMessage) *Message {
	return s.Message(TypeEvent, name, data)
}

func MustMarshalJSON(v interface{}) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
