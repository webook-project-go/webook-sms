package service

import (
	"context"
	"github.com/webook-project-go/webook-sms/repository"
	"math/rand/v2"
	"strconv"
)

var (
	ErrSendTooFrequent      = repository.ErrSendTooFrequent
	ErrSystemError          = repository.ErrSystemError
	ErrInvalidCode          = repository.ErrInvalidCode
	ErrTooManyVerifications = repository.ErrTooManyVerifications
	ErrWrongCode            = repository.ErrWrongCode
)

type codeService struct {
	sms  SMSService
	repo repository.CodeRepository
}

func NewCodeService(sms SMSService, repo repository.CodeRepository) CodeService {
	return &codeService{
		sms:  sms,
		repo: repo,
	}
}
func (c *codeService) Send(ctx context.Context, biz, phone string) error {
	code := c.generateCode()
	err := c.sms.Send(ctx, Message{
		PhoneNumbers: phone,
		Data:         strconv.Itoa(code),
	})
	if err != nil {
		return err
	}
	return c.repo.Store(ctx, biz, phone, code)
}

func (c *codeService) VerifyCode(ctx context.Context, biz, phone string, code int) (bool, error) {
	return c.repo.VerifyCode(ctx, biz, phone, code)
}

func (c *codeService) generateCode() int {
	return rand.IntN(100000) + 100000
}
