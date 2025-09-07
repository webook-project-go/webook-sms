package decorator

import (
	"context"
	"errors"
	"github.com/webook-project-go/webook-sms/service"
	"sync/atomic"
)

type FailOverSMSService struct {
	svcs        []service.SMSService
	threshold   int32
	idx         int32
	failedTimes int32
}

func NewFailOverSMSService(svcs []service.SMSService) service.SMSService {
	return &FailOverSMSService{
		svcs:        svcs,
		threshold:   3,
		idx:         0,
		failedTimes: 0,
	}
}

func (f *FailOverSMSService) Send(ctx context.Context, msg service.Message) error {
	idx := atomic.LoadInt32(&f.idx)
	if atomic.LoadInt32(&f.failedTimes) >= f.threshold {
		newIdx := (idx + 1) % int32(len(f.svcs))
		if atomic.CompareAndSwapInt32(&f.idx, idx, newIdx) {
			atomic.StoreInt32(&f.failedTimes, 0)
			idx = newIdx
		} else {
			idx = atomic.LoadInt32(&f.idx)
		}
	}

	svc := f.svcs[idx]
	err := svc.Send(ctx, msg)
	switch {
	case err == nil:
		atomic.StoreInt32(&f.failedTimes, 0)
		return nil
	case errors.Is(err, context.DeadlineExceeded):
		atomic.AddInt32(&f.failedTimes, 1)

	default:
		atomic.StoreInt32(&f.idx, (idx+1)%int32(len(f.svcs)))
		atomic.StoreInt32(&f.failedTimes, 0)

	}
	return err
}
