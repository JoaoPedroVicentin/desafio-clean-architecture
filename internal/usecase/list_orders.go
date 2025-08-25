package usecase

import (
	"github.com/JoaoPedroVicentin/desafio-clean-architecture/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func GetOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) ListOrders() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var dto []OrderOutputDTO

	for _, order := range orders {
		dto = append(dto, OrderOutputDTO{
			order.ID,
			order.Price,
			order.Tax,
			order.FinalPrice,
		})
	}

	return dto, nil
}
