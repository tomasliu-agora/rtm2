package rtm2

// ChannelType enum
type ChannelType int

const (
	ChannelTypeMessage ChannelType = 0 // Message Channel
	ChannelTypeStream  ChannelType = 1 // Stream Channel
)

// ConnectionEvent will be notified when the connection state changes between rtm sdk and agora service.
type ConnectionEvent struct {
	// Connection states between rtm sdk and agora server.
	// See consts.go for detail.
	State int32
	// Reasons for connection state change.
	// See consts.go for detail.
	Reason int32
}

type RTMClient interface {
	// Login the Agora RTM service. The operation result will be returned.
	// Connection Events will be notified by event channel.
	Login(token string) (<-chan *ConnectionEvent, error)
	// Logout the Agora RTM service.
	// Be noticed that this method will break the rtm service including storage/lock/presence.
	Logout() error
	// SetParameters to rtm sdk and service.
	SetParameters(params map[string]interface{}) error
	// GetParameters returns current parameters set to sdk and service.
	GetParameters() map[string]interface{}
	// RenewToken will renew the token to rtm service.
	// Once a token is enabled and used, it expires after a certain period of time.
	// You should generate a new token from your token server and call RenewToken to renew it.
	RenewToken(token string) error

	// Storage gets the storage interface.
	Storage() Storage
	// Lock gets the lock interface.
	Lock() Lock
	// Presence gets the presence interface.
	Presence() Presence

	// Publish a message into certain Message Channel.
	Publish(channel string, message []byte, opts ...MessageOption) error
	// Subscribe certain Message Channel.
	// Options can be set if Storage / Lock / Presence subscription is needed.
	Subscribe(channel string, opts ...MessageOption) (chan *Message, error)
	// Unsubscribe certain Message Channel.
	// Storage / Lock / Presence subscriptions will be canceled as well.
	Unsubscribe(channel string) error

	// StreamChannel returns a Stream Channel interface.
	// Call StreamChannel multiple times on same channel will return the same interface.
	StreamChannel(channel string) StreamChannel
}
