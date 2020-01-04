package plume

const (
	MsgHello = iota
	MsgText

	MsgAuth
	MsgAuthAck
	MsgAuthRst

	MsgMayNotEnter

	// In the future, we can support things like presence
	// by using additioal codes liek MsgTypingStart/Stop
)

type Message struct {
	Type int    `json:"type"`
	User User   `json:"user"`
	Text string `json:"text"`
}
