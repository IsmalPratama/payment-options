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
	result["go_pay"] = u.repo.CallGoPay()         
	result["ovo"] = u.repo.CallOVO()            
	result["dana"] = u.repo.CallDANA()           
	result["linkaja"] = u.repo.CallLinkAja()     
	result["kredivo"] = u.repo.CallKredivo()     
	result["shopeepay"] = u.repo.CallShopeePay() 
	result["bca_virtual_account"] = u.repo.CallVirtualAccountBCA() 
	return result, nil
}

func (u *paymentUsecase) GetPaymentOptions22() (map[string]models.PaymentMethod, error) {
	var wg sync.WaitGroup
	result := make(map[string]models.PaymentMethod)
	mu := sync.Mutex{}

	wg.Add(7)

	go func() {
		defer wg.Done()
		mu.Lock()
		result["go_pay"] = u.repo.CallGoPay()         
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["ovo"] = u.repo.CallOVO()             
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["dana"] = u.repo.CallDANA()           
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
		result["kredivo"] = u.repo.CallKredivo()     
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["shopeepay"] = u.repo.CallShopeePay() 
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		result["bca_virtual_account"] = u.repo.CallVirtualAccountBCA()
		mu.Unlock()
	}()

	wg.Wait()
	return result, nil
}
