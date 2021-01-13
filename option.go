package jpushclient

type Option struct {
	SendNo          int   `json:"sendno,omitempty"`
	TimeLive        int   `json:"time_to_live,omitempty"`
	ApnsProduction  bool  `json:"apns_production"`
	OverrideMsgId   int64 `json:"override_msg_id,omitempty"`
	BigPushDuration int   `json:"big_push_duration,omitempty"`

	// 推送请求下发通道，仅对厂商VIP用户有效
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel"`
}

func (this *Option) SetSendno(no int) {
	this.SendNo = no
}

func (this *Option) SetTimelive(timelive int) {
	this.TimeLive = timelive
}

func (this *Option) SetOverrideMsgId(id int64) {
	this.OverrideMsgId = id
}

func (this *Option) SetApns(apns bool) {
	this.ApnsProduction = apns
}

func (this *Option) SetBigPushDuration(bigPushDuration int) {
	this.BigPushDuration = bigPushDuration
}

func (this *Option) SetThirdPartyChannel(thirdPartyChannel map[string]interface{}) {
	this.ThirdPartyChannel = thirdPartyChannel
}