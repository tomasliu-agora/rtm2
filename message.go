package rtm2

// MessageType for Message Channel and Stream Channel
type MessageType int

const (
	MessageTypeBinary MessageType = 0
	MessageTypeString MessageType = 1
)

type Message struct {
	UserId     string
	Type       MessageType
	Message    []byte
	CustomType string
}

type MessageOptions struct {
	Type       MessageType
	CustomType string

	Message  bool
	Metadata bool
	Presence bool
	Lock     bool
}

func DefaultMessageOptions() *MessageOptions {
	return &MessageOptions{Type: MessageTypeBinary, Message: true, Metadata: false, Presence: true, Lock: false}
}

type MessageOption func(c *MessageOptions)

// WithMessageType is MessageTypeBinary by default.
func WithMessageType(t MessageType) MessageOption {
	return func(c *MessageOptions) {
		c.Type = t
	}
}

// WithMessage whether to subscribe message in the Message Channel.
func WithMessage(enabled bool) MessageOption {
	return func(c *MessageOptions) {
		c.Message = enabled
	}
}

// WithMessageMetadata whether to subscribe Channel Metadata in the Message Channel.
func WithMessageMetadata(enabled bool) MessageOption {
	return func(c *MessageOptions) {
		c.Metadata = enabled
	}
}

// WithMessagePresence whether to subscribe user Presence in the Message Channel.
func WithMessagePresence(enabled bool) MessageOption {
	return func(c *MessageOptions) {
		c.Presence = enabled
	}
}

// WithMessageLock whether to subscribe Lock in the Message Channel.
func WithMessageLock(enabled bool) MessageOption {
	return func(c *MessageOptions) {
		c.Lock = enabled
	}
}

// WithMessageCustomType is custom type of the message, up to 32 bytes for customize
func WithMessageCustomType(customType string) MessageOption {
	return func(c *MessageOptions) {
		c.CustomType = customType
	}
}
