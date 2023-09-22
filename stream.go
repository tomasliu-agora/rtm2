package rtm2

type StreamQOS int
type StreamPriority int
type TopicEventType int

const (
	StreamQosUnordered StreamQOS = 0
	StreamQosOrdered   StreamQOS = 1

	StreamQosPriorityHighest StreamPriority = 0
	StreamQosPriorityHigh    StreamPriority = 1
	StreamQosPriorityNormal  StreamPriority = 2
	StreamQosPriorityLow     StreamPriority = 3

	TopicEventNone     TopicEventType = 0
	TopicEventSnapshot TopicEventType = 1
	TopicEventJoin     TopicEventType = 2
	TopicEventLeave    TopicEventType = 3
)

type TopicEvent struct {
	Type     TopicEventType
	Channel  string
	UserId   string
	Snapshot map[string][]string // Only for snapshot: topic -> [userId.]
	Topic    string
}

type StreamOptions struct {
	// Join
	Token    string
	Metadata bool
	Presence bool
	Lock     bool

	// JoinTopic
	QOS       StreamQOS
	Priority  StreamPriority
	Meta      string
	SyncMedia bool

	// Publish
	Type       MessageType
	SendTs     uint64
	CustomType string
}

type StreamOption func(*StreamOptions)

// WithStreamToken different token can be set on Join this Stream Channel
func WithStreamToken(token string) StreamOption {
	return func(c *StreamOptions) {
		c.Token = token
	}
}

// WithStreamMetadata whether to subscribe Channel Metadata in the Stream Channel.
func WithStreamMetadata(enabled bool) StreamOption {
	return func(c *StreamOptions) {
		c.Metadata = enabled
	}
}

// WithStreamPresence whether to subscribe user Presence in the Stream Channel.
func WithStreamPresence(enabled bool) StreamOption {
	return func(c *StreamOptions) {
		c.Presence = enabled
	}
}

// WithStreamLock whether to subscribe Lock in the Stream Channel.
func WithStreamLock(enabled bool) StreamOption {
	return func(c *StreamOptions) {
		c.Lock = enabled
	}
}

// WithStreamQOS sets the qos for all messages to publish on this topic
func WithStreamQOS(qos StreamQOS) StreamOption {
	return func(c *StreamOptions) {
		c.QOS = qos
	}
}

// WithStreamPriority sets the priority for all messages to publish on this topic
func WithStreamPriority(p StreamPriority) StreamOption {
	return func(c *StreamOptions) {
		c.Priority = p
	}
}

// WithStreamMeta sets the meta info on this topic
// TODO: currently not supported
func WithStreamMeta(meta string) StreamOption {
	return func(c *StreamOptions) {
		c.Meta = meta
	}
}

// WithStreamSyncMedia whether rtm message will be calibrated with rtc
func WithStreamSyncMedia(enabled bool) StreamOption {
	return func(c *StreamOptions) {
		c.SyncMedia = enabled
	}
}

// WithStreamMessageType is MessageTypeBinary by default.
func WithStreamMessageType(t MessageType) StreamOption {
	return func(c *StreamOptions) {
		c.Type = t
	}
}

// WithStreamSendTs used to calibrate data with media.
// Only valid when user join topic WithStreamSyncMedia.
func WithStreamSendTs(ts uint64) StreamOption {
	return func(c *StreamOptions) {
		c.SendTs = ts
	}
}

// WithCustomType is custom type of the message, up to 32 bytes for customize
func WithStreamCustomType(customType string) StreamOption {
	return func(c *StreamOptions) {
		c.CustomType = customType
	}
}

type StreamChannel interface {
	// Join certain Stream Channel
	// Returns the snapshot of current topic infos and a golang chan for TopicEvent
	Join(opts ...StreamOption) (map[string][]string, <-chan *TopicEvent, <-chan string, error)
	// Leave certain Stream Channel
	Leave() error
	// ChannelName returns the name for current Stream Channel
	ChannelName() string
	// JoinTopic joins certain topic. Can only be called after joined.
	JoinTopic(topic string, opt ...StreamOption) error
	// PublishTopic publishes message to certain joined topic.
	PublishTopic(topic string, message []byte, opts ...StreamOption) error
	// LeaveTopic leaves certain topic.
	LeaveTopic(topic string) error

	// SubscribeTopic subscribes certain topic on certain Users.
	// If userIds is set empty, RTM will subscribe all joined user on that topic.
	// Returns golang chan of Message.
	SubscribeTopic(topic string, userIds []string) (<-chan *Message, error)
	// UnsubscribeTopic unsubscribes certain topic on certain Users.
	// If userIds is set empty, RTM will unsubscribe all users on that topic.
	// Warn: The golang chan will not be closed even if no user is subscribed.
	UnsubscribeTopic(topic string, userIds []string) error
	// GetSubscribedUsers returns all subscribed users locally.
	GetSubscribedUsers(topic string) ([]string, error)
	// RenewToken
	RenewToken(token string) error
}
