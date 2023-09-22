package rtm2

type LockEventType int

const (
	LockTypeNone     LockEventType = 0
	LockTypeSnapshot LockEventType = 1 // Triggered after reconnecting with rtm service.
	LockTypeSet      LockEventType = 2 // Triggered after certain lock is set
	LockTypeRemove   LockEventType = 3 // Triggered after certain lock is removed
	LockTypeAcquired LockEventType = 4 // Triggered after certain lock is acquired by certain User (including self)
	LockTypeReleased LockEventType = 5 // Triggered after certain lock is released by certain User (including self)
	LockTypeExpired  LockEventType = 6 // Triggered after certain lock is expired
)

type LockEvent struct {
	Type    LockEventType // The change type of Lock
	Details []*LockDetail // The detail information of Lock changes
}

// LockDetail stores the information of a Lock.
type LockDetail struct {
	Name  string // The name of the lock.
	Owner string // The owner of the lock. Empty if the lock is not acquired by anyone.
	TTL   uint32 // Time to live in seconds. The lock will be expired afterwards.
}

// Lock provides the rtm lock interfaces that can be invoked by your app.
type Lock interface {
	// GetLockChan must be called after Lock is subscribed on certain channel.
	// Otherwise, the golang chan will be blocked which will cause fatal error.
	// Return error if no subscription found.
	GetLockChan(channel string, channelType ChannelType) (map[string]*LockDetail, <-chan *LockEvent, error)

	// Set a lock in certain channel. No matter Stream Channel or Message Channel.
	Set(channel string, channelType ChannelType, name string, ttl uint32) error
	// Get all locks' detail in certain channel. No matter Stream Channel or Message Channel.
	Get(channel string, channelType ChannelType) (map[string]*LockDetail, error)
	// Remove a lock in certain channel. No matter Stream Channel or Message Channel.
	Remove(channel string, channelType ChannelType, name string) error

	// Acquire a lock in certain channel.
	// No matter Stream Channel or Message Channel. No matter joined or not.
	// Returns a golang chan to wait for result.
	// If error is fetched, means acquire failed.
	Acquire(channel string, channelType ChannelType, name string, retry bool) <-chan error
	// Release a lock in certain channel.
	// No matter Stream Channel or Message Channel. No matter joined or not.
	// TODO: Will close the golang chan returned by Acquire if retry is set.
	Release(channel string, channelType ChannelType, name string) error
	// Revoke a lock from certain user in certain channel.
	// No matter Stream Channel or Message Channel. No matter joined or not.
	// Returns error if the lock's owner does not match.
	Revoke(channel string, channelType ChannelType, name string, owner string) error
}
