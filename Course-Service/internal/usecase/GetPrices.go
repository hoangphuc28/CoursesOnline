package usecase

import (
	"github.com/hoangphuc28/CoursesOnline-ProtoFile/Course"
	"strconv"
)

func (uc *coursesUseCase) GetPrices() (*Course.GetPricesResponse, error) {
	res, err := uc.repo.GetPrices()
	if err != nil {
		return nil, err
	}
	var list Course.GetPricesResponse
	for _, item := range res {
		list.Prices = append(list.Prices, &Course.Price{
			Id:       strconv.Itoa(item.Id),
			Value:    item.Value,
			Currency: item.Currency,
		})
	}
	return &list, nil

}
