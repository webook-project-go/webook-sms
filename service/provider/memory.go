package provider

import (
	"context"
	"fmt"
	"github.com/webook-project-go/webook-sms/service"
)

type MemoryService struct {
}

func NewSMSMemory() service.SMSService {
	return &MemoryService{}
}
func (M *MemoryService) Send(ctx context.Context, message service.Message) error {
	fmt.Println(message)
	return nil
}
