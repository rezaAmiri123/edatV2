package jetstream

import (
	"time"

	"github.com/rezaAmiri123/edatV2/am"
	"github.com/rezaAmiri123/edatV2/ddd"
)

type rawMessage struct {
	id         string
	name       string
	subjet     string
	data       []byte
	metadata   ddd.Metadata
	sentAt     time.Time
	receivedAt time.Time
	acked      bool
	ackFu      func() error
	nackFn     func() error
	extendFn   func() error
	killFn     func() error
}

var _ am.IncomingMessage = (*rawMessage)(nil)

func (m *rawMessage) ID() string             { return m.id }
func (m *rawMessage) Subject() string        { return m.subjet }
func (m *rawMessage) MessageName() string    { return m.name }
func (m *rawMessage) Metadata() ddd.Metadata { return m.metadata }
func (m *rawMessage) SentAt() time.Time      { return m.sentAt }
func (m *rawMessage) Data() []byte           { return m.data }
func (m *rawMessage) ReceivedAt() time.Time  { return m.receivedAt }

func (m *rawMessage) Ack() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.ackFu()
}

func (m *rawMessage) NAck() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.nackFn()
}

func (m *rawMessage) Extend() error {
	return m.extendFn()
}

func (m *rawMessage) Kill() error {
	if m.acked {
		return nil
	}
	m.acked = true
	return m.killFn()
}