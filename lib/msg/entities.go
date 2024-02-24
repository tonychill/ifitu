package msg

import "context"

type Messenger interface {
	SendMessage(context.Context, Message) SendMessageResponse
}

type Message struct {
	Email *EmailMessage
	Sms   *SmsMessage
}

type EmailMessage struct{}

type SmsMessage struct{}

type SendMessageResponse struct {
}

type messengerImpl struct {
	tc interface{} // twillio client
}
