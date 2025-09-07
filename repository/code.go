package repository

import (
	"context"
	"github.com/webook-project-go/webook-sms/repository/cache"
)

var (
	ErrSendTooFrequent      = cache.ErrSendTooFrequent
	ErrSystemError          = cache.ErrSystemError
	ErrInvalidCode          = cache.ErrInvalidCode
	ErrTooManyVerifications = cache.ErrTooManyVerifications
	ErrWrongCode            = cache.ErrWrongCode
)

type CodeRepository interface {
	Store(ctx context.Context, biz, phone string, code int) error
	VerifyCode(ctx context.Context, biz, phone string, code int) (bool, error)
}

type codeRepository struct {
	codeCache cache.CodeCache
}

func NewCodeRepository(codeCache cache.CodeCache) CodeRepository {
	return &codeRepository{codeCache: codeCache}
}
func (c *codeRepository) Store(ctx context.Context, biz, phone string, code int) error {
	return c.codeCache.Set(ctx, biz, phone, code)
}

func (c *codeRepository) VerifyCode(ctx context.Context, biz, phone string, code int) (bool, error) {
	return c.VerifyCode(ctx, biz, phone, code)
}
