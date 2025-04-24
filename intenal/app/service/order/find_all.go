package orderservice

import (
	"context"

	svc "github.com/zhora-ip/libraries-management-system/intenal/models/service"
)

func (s *OrderService) FindAll(ctx context.Context, req *svc.FindAllOrdersRequest) (*svc.FindAllOrdersResponse, error) {

	orders, err := s.oRepo.FindAll(ctx, req)

	if err != nil {
		return nil, err
	}

	return &svc.FindAllOrdersResponse{Data: orders}, nil
}
