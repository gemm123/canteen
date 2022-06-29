package service

import "github.com/gemm123/canteen/models"

func (s *service) GetCash() (models.Money, error) {
	money, err := s.repository.GetCash()
	if err != nil {
		return money, err
	}

	return money, nil
}

func (s *service) UpdateCash(id int, moneyRequest models.Money) (models.Money, error) {
	money, err := s.repository.GetCashByID(id)
	if err != nil {
		return money, err
	}

	money.Cash = moneyRequest.Cash

	newMoney, err := s.repository.UpdateCash(money)
	if err != nil {
		return newMoney, err
	}

	return newMoney, nil
}
