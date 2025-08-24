package usecase

import (
	"github.com/JoaoPedroVicentin/desafio-clean-architecture/internal/entity"
	"github.com/JoaoPedroVicentin/desafio-clean-architecture/pkg/events"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	EventDispatcher events.EventDispatcherInterface
}

func GetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCase) ListOrders() ([]entity.Order, error) {
	orders, err := c.OrderRepository.ListOrders()
	if err != nil {
		return nil, err
	}

	return orders, nil
}
