package decorator

import (
	"context"
	"github.com/webook-project-go/webook-sms/service"
	"sync/atomic"
)

type FailOverSMSSlideWindowVer struct {
	svcs []service.SMSService
	idx  int32
	sd   *FailOverSlideWindow
}

func NewFailOverSMSSlideWindowVer(sd *FailOverSlideWindow, svcs []service.SMSService) *FailOverSMSSlideWindowVer {
	return &FailOverSMSSlideWindowVer{
		svcs: svcs,
		sd:   sd,
		idx:  0,
	}
}

func (f *FailOverSMSSlideWindowVer) Send(ctx context.Context, msg service.Message) error {

	idx := atomic.LoadInt32(&f.idx)

	svc := f.svcs[idx]
	err := svc.Send(ctx, msg)
	f.sd.Add(err == nil)

	switch {
	case err == nil:
		return nil
	case f.sd.ShouldFailOver():
		newIdx := (idx + 1) % int32(len(f.svcs))
		atomic.CompareAndSwapInt32(&f.idx, idx, newIdx)

	}
	return err
}
