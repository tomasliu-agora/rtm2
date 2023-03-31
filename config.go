package rtm2

import "go.uber.org/zap"

// RTMConfig stores configurations for RTM Client.
type RTMConfig struct {
	Appid  string // The App ID of your project. Must have.
	UserId string // The ID of the user. Must have.
	Vid    uint32 // The Vendor ID of your project.

	// The region for connection. This advanced feature applies to scenarios that have regional restrictions.
	// For the regions that Agora supports, see consts.go for detail.
	// After specifying the region, the SDK connects to the Agora servers within that region.
	AreaCode uint32
	// PresenceTimeout in seconds, specify the time to preserve presence after disconnecting with rtm service.
	PresenceTimeout uint32
	// The log file path, default is empty, which stands for disable.
	FilePath string
	// External logger for golang side log file.
	Logger *zap.Logger
}
