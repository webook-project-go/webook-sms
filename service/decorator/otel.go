package decorator

import (
	"context"
	"github.com/webook-project-go/webook-sms/service"
	"go.opentelemetry.io/otel/trace"
)

type TraceService struct {
	service.SMSService
	tracer trace.Tracer
}

func NewTraceService(svc service.SMSService, trace trace.Tracer) service.SMSService {
	return &TraceService{
		SMSService: svc,
		tracer:     trace,
	}
}
func (t *TraceService) Send(ctx context.Context, message service.Message) error {
	ctx, span := t.tracer.Start(ctx, "sms_service")
	defer span.End()
	err := t.SMSService.Send(ctx, message)
	span.RecordError(err)
	return err
}
