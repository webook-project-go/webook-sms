package service

import "context"

type Message struct {
	PhoneNumbers string
	Data         string
}

type SMSService interface {
	Send(context.Context, Message) error
}

type CodeService interface {
	Send(ctx context.Context, biz, phone string) error
	VerifyCode(ctx context.Context, biz, phone string, code int) (bool, error)
}
