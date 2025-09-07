package grpc

import (
	"context"
	"github.com/webook-project-go/webook-apis/gen/go/apis/sms/v1"
	"github.com/webook-project-go/webook-sms/service"
)

type Service struct {
	sms  service.SMSService
	code service.CodeService
	v1.UnimplementedSMSServiceServer
}

func New(sms service.SMSService, code service.CodeService) *Service {
	return &Service{
		sms:  sms,
		code: code,
	}
}
func (s *Service) Send(ctx context.Context, request *v1.SendRequest) (*v1.SendResponse, error) {
	err := s.sms.Send(ctx, service.Message{
		PhoneNumbers: request.GetPhone(),
		Data:         request.GetContent(),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) SendCode(ctx context.Context, request *v1.SendCodeRequest) (*v1.SendCodeResponse, error) {
	err := s.code.Send(ctx, request.GetBiz(), request.GetPhone())
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *Service) VerifyCode(ctx context.Context, request *v1.VerifyCodeRequest) (*v1.VerifyCodeResponse, error) {
	ok, err := s.code.VerifyCode(ctx, request.GetBiz(), request.GetPhone(), int(request.GetCode()))
	if err != nil {
		return nil, err
	}
	return &v1.VerifyCodeResponse{OK: ok}, nil
}
