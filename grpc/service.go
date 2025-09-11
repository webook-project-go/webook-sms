package grpc

import (
	"context"
	"github.com/webook-project-go/webook-apis/code"
	"github.com/webook-project-go/webook-apis/gen/go/apis/sms/v1"
	"github.com/webook-project-go/webook-sms/repository"
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
		var codeVal int32
		switch err {
		case repository.ErrSendTooFrequent:
			codeVal = code.ErrCodeSendTooFrequent
		case repository.ErrTooManyVerifications:
			codeVal = code.ErrCodeTooManyVerifications
		}
		return &v1.SendCodeResponse{
			Code: codeVal,
		}, err
	}

	return &v1.SendCodeResponse{
		Code: 0, // 成功
	}, nil
}

func (s *Service) VerifyCode(ctx context.Context, request *v1.VerifyCodeRequest) (*v1.VerifyCodeResponse, error) {
	ok, err := s.code.VerifyCode(ctx, request.GetBiz(), request.GetPhone(), int(request.GetCode()))
	if err != nil {
		var codeVal int32
		switch err {
		case repository.ErrInvalidCode:
			codeVal = code.ErrCodeInvalidCode
		case repository.ErrWrongCode:
			codeVal = code.ErrCodeWrongCode
		case repository.ErrTooManyVerifications:
			codeVal = code.ErrCodeTooManyVerifications
		}
		return &v1.VerifyCodeResponse{
			OK:   false,
			Code: codeVal,
		}, err
	}

	return &v1.VerifyCodeResponse{
		OK:   ok,
		Code: 0, // 成功
	}, nil
}
