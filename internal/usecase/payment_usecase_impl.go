package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	"sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
		result := make(map[string]models.PaymentMethod)
		result["bca_mobile"] = u.repo.CallBCAMobile()
		result["linkaja"] = u.repo.CallLinkAja()
		result["bni_mobile"] = u.repo.CallBNIMobile()
		result["mandiri_online"] = u.repo.CallMandiriOnline()
		result["kredivo"] = u.repo.CallKredivo()
		result["jenius"] = u.repo.CallJenius()
		return result, nil
}

func (u *paymentUsecase) GetPaymentOptions22() (map[string]models.PaymentMethod, error) {
	var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(6)

	go func() {
		defer wg.Done()
		mu.Lock()
		result["bca_mobile"] = u.repo.CallBCAMobile()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["linkaja"] = u.repo.CallLinkAja()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["bni_mobile"] = u.repo.CallBNIMobile()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["mandiri_online"] = u.repo.CallMandiriOnline()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["kredivo"] = u.repo.CallKredivo()
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["jenius"] = u.repo.CallJenius()
		mu.Unlock()
	}()

	wg.Wait()
	return result, nil
}