package rtm2

const (
	// AreaCode Enum in RTMConfig

	AreaCodeCN   = 0x00000001
	AreaCodeNA   = 0x00000002
	AreaCodeEU   = 0x00000004
	AreaCodeAS   = 0x00000008
	AreaCodeJP   = 0x00000010
	AreaCodeIN   = 0x00000020
	AreaCodeGLOB = 0xFFFFFFFF
	AreaCodeOC   = 0x00000040
	AreaCodeSA   = 0x00000080
	AreaCodeAF   = 0x00000100
	AreaCodeKR   = 0x00000200
	AreaCodeHKMC = 0x00000400
	AreaCodeUS   = 0x00000800
	AreaCodeOVS  = 0xFFFFFFFE

	// Connection State Enum in ConnectionEvent

	ConnectionStateDISCONNECTED = 1
	ConnectionStateCONNECTING   = 2
	ConnectionStateCONNECTED    = 3
	ConnectionStateRECONNECTING = 4
	ConnectionStateFAILED       = 5

	// Connection State Change Reason in ConnectionEvent

	ConnectionChangedReasonConnecting                 = 0
	ConnectionChangedReasonJoinSuccess                = 1
	ConnectionChangedReasonInterrupted                = 2
	ConnectionChangedReasonBannedByServer             = 3
	ConnectionChangedReasonJoinFailed                 = 4
	ConnectionChangedReasonLeaveChannel               = 5
	ConnectionChangedReasonInvalidAppId               = 6
	ConnectionChangedReasonInvalidChannelName         = 7
	ConnectionChangedReasonInvalidToken               = 8
	ConnectionChangedReasonTokenExpired               = 9
	ConnectionChangedReasonRejectedByServer           = 10
	ConnectionChangedReasonSettingProxyServer         = 11
	ConnectionChangedReasonRenewToken                 = 12
	ConnectionChangedReasonClientIpAddrChanged        = 13
	ConnectionChangedReasonKeepaliveTimeout           = 14
	ConnectionChangedReasonRejoinSuccess              = 15
	ConnectionChangedReasonLost                       = 16
	ConnectionChangedReasonEchoLost                   = 17
	ConnectionChangedReasonClientIpAddrChangedByUser  = 18
	ConnectionChangedReasonSameUidLogin               = 19
	ConnectionChangedReasonTooManyBroadcaster         = 20
	ConnectionChangedReasonStreamChannelNotAvaliabled = 22
	ConnectionChangedReasonLoginSuccess               = 10001
)
