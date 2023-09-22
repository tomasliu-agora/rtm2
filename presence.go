package rtm2

type PresenceEventType int

const (
	PresenceTypeNone         PresenceEventType = 0
	PresenceTypeSnapshot     PresenceEventType = 1
	PresenceTypeInterval     PresenceEventType = 2
	PresenceTypeJoinChannel  PresenceEventType = 3
	PresenceTypeLeaveChannel PresenceEventType = 4
	PresenceTypeTimeout      PresenceEventType = 5
	PresenceTypeStateChange  PresenceEventType = 6
	PresenceTypeOutOfService PresenceEventType = 7
)

type PresenceEvent struct {
	Type   PresenceEventType
	UserId string
	Items  map[string]string
	// For interval
	Joined  []string
	Left    []string
	Timeout []string
	// For snapshot and interval event
	States map[string]map[string]string
}

type UserState struct {
	UserId string
	State  map[string]string
}

type ChannelInfo struct {
	Channel string
	Type    ChannelType
}

type PresenceOptions struct {
	UserId bool
	State  bool
	Page   string
}

type PresenceOption func(*PresenceOptions)

// WithPresenceUserId whether to display user id in query result.
func WithPresenceUserId(enabled bool) PresenceOption {
	return func(c *PresenceOptions) {
		c.UserId = enabled
	}
}

// WithPresenceState whether to display user state in query result.
func WithPresenceState(enabled bool) PresenceOption {
	return func(c *PresenceOptions) {
		c.State = enabled
	}
}

// WithPage stores next page index.
func WithPage(page string) PresenceOption {
	return func(c *PresenceOptions) {
		c.Page = page
	}
}

type Presence interface {
	// GetPresenceChan must be called after Presence is subscribed on certain channel.
	// Otherwise, the golang chan will be blocked which will cause fatal error.
	// Return error if no subscription found.
	GetPresenceChan(channel string, channelType ChannelType) (map[string]*UserState, <-chan *PresenceEvent, error)

	// WhoNow returns all users joined certain channel.
	// Paging is supported: if there are more Users, return "next page index" on second return value.
	// Use WithPage to start from "next page index".
	WhoNow(channel string, channelType ChannelType, opts ...PresenceOption) (map[string]*UserState, string, error)
	// WhereNow all channels certain user has joined. No matter Message Channel or Stream Channel.
	WhereNow(userId string) ([]*ChannelInfo, error)
	// GetOnlineUsers returns all users joined certain channel.
	// Paging is supported: if there are more Users, return "next page index" on second return value.
	// Use WithPage to start from "next page index".
	GetOnlineUsers(channel string, channelType ChannelType, opts ...PresenceOption) (map[string]*UserState, string, error)
	// GetUserChannels all channels certain user has joined. No matter Message Channel or Stream Channel.
	GetUserChannels(userId string) ([]*ChannelInfo, error)
	// SetState can be called before Join Stream Channel or Subscribe Message Channel.
	// Cache the state and auto set after Join or Subscribe.
	SetState(channel string, channelType ChannelType, data map[string]string) error
	// RemoveState can be called before Join Stream Channel or Subscribe Message Channel.
	// Cache state can be removed as well.
	RemoveState(channel string, channelType ChannelType, keys []string) error
	// GetState can query certain user's state on certain channel.
	GetState(channel string, channelType ChannelType, userId string) (map[string]string, error)
}
