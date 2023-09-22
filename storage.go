package rtm2

type StorageEventType int32

const (
	StorageTypeNone     StorageEventType = 0
	StorageTypeSnapshot StorageEventType = 1
	StorageTypeSet      StorageEventType = 2
	StorageTypeUpdate   StorageEventType = 3
	StorageTypeRemove   StorageEventType = 4
)

type MetadataItem struct {
	Key      string
	Value    string
	Author   string
	Revision int64
	UpdateTs int64
}

type StorageEvent struct {
	EventType     StorageEventType
	MajorRevision int64
	Items         map[string]*MetadataItem
}

type StorageOptions struct {
	MajorRev     int64
	Lock         string
	RecordTs     bool
	RecordAuthor bool
}

type StorageOption func(*StorageOptions)

// WithStorageLock set the lock name for the Set, Update or Remove.
// If set, the operation will only success if lock is acquired by self.
func WithStorageLock(lock string) StorageOption {
	return func(c *StorageOptions) {
		c.Lock = lock
	}
}

// WithStorageMajorRev set the major revision for the Set, Update or Remove.
// If set, the operation will only success if major revision matches.
func WithStorageMajorRev(rev int64) StorageOption {
	return func(c *StorageOptions) {
		c.MajorRev = rev
	}
}

// WithStorageRecordTs whether to record timestamp along with Metadata Set or Update
func WithStorageRecordTs(enabled bool) StorageOption {
	return func(c *StorageOptions) {
		c.RecordTs = enabled
	}
}

// WithStorageRecordAuthor whether to record the modifier's id along with Metadata Set or Update
func WithStorageRecordAuthor(enabled bool) StorageOption {
	return func(c *StorageOptions) {
		c.RecordAuthor = enabled
	}
}

type Storage interface {
	// GetChannelMetadataChan must be called after Metadata is subscribed on certain channel.
	// Otherwise, the golang chan will be blocked which will cause fatal error.
	// Return error if no subscription found.
	GetChannelMetadataChan(channel string, channelType ChannelType) (map[string]*MetadataItem, <-chan *StorageEvent, error)

	// SetChannelMetadata set multiple metadata items to certain channel.
	// Will create new Metadata Item if not exists.
	// No matter Message Channel or Stream Channel. Can be called without Join Stream Channel or Subscribe Message Channel.
	SetChannelMetadata(channel string, channelType ChannelType, data map[string]*MetadataItem, opts ...StorageOption) error
	// UpdateChannelMetadata set multiple metadata items to certain channel.
	// Returns error if the Metadata Item not exists.
	// No matter Message Channel or Stream Channel. Can be called without Join Stream Channel or Subscribe Message Channel.
	UpdateChannelMetadata(channel string, channelType ChannelType, data map[string]*MetadataItem, opts ...StorageOption) error
	// RemoveChannelMetadata remove multiple metadata items to certain channel.
	// No matter Message Channel or Stream Channel. Can be called without Join Stream Channel or Subscribe Message Channel.
	RemoveChannelMetadata(channel string, channelType ChannelType, data map[string]*MetadataItem, opts ...StorageOption) error
	// GetChannelMetadata returns all metadata items on certain channel.
	// No matter Message Channel or Stream Channel. Can be called without Join Stream Channel or Subscribe Message Channel.
	GetChannelMetadata(channel string, channelType ChannelType) (int64, map[string]*MetadataItem, error)

	// SetUserMetadata set multiple metadata items on certain User. Will create new Metadata Item if not exists.
	SetUserMetadata(userId string, data map[string]*MetadataItem, opts ...StorageOption) error
	// UpdateUserMetadata set multiple metadata items on certain User. Returns error if the Metadata Item not exists.
	UpdateUserMetadata(userId string, data map[string]*MetadataItem, opts ...StorageOption) error
	// RemoveUserMetadata remove multiple metadata items on certain User.
	RemoveUserMetadata(userId string, data map[string]*MetadataItem, opts ...StorageOption) error
	// GetUserMetadata returns all metadata items on certain User.
	GetUserMetadata(userId string) (int64, map[string]*MetadataItem, error)
	// SubscribeUserMetadata will subscribe metadata items on certain User.
	// Returns a snapshot and golang chan of StorageEvent.
	SubscribeUserMetadata(userId string) (map[string]*MetadataItem, <-chan *StorageEvent, error)
	// UnsubscribeUserMetadata will unsubscribe metadata items on certain User.
	UnsubscribeUserMetadata(userId string) error
}
