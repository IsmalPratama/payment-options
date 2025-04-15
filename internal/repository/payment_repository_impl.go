package repository

import (
	"payment-options/internal/models"
	// "time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallBCAMobile() models.PaymentMethod {
	// time.Sleep(2 * time.Second) // Simulate network delay
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "50000",
		Icon:    "https://sampleurl.com/bca.jpg",
	}
}

func (r *paymentRepo) CallLinkAja() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "20000",
		Icon:    "https://sampleurl.com/linkaja.jpg",
	}
}

func (r *paymentRepo) CallBNIMobile() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "30000",
		Icon:    "https://sampleurl.com/bni.jpg",
	}
}

func (r *paymentRepo) CallMandiriOnline() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "45000",
		Icon:    "https://sampleurl.com/mandiri.jpg",
	}
}

func (r *paymentRepo) CallKredivo() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "8000",
		Icon:    "https://sampleurl.com/kredivo.jpg",
	}
}

func (r *paymentRepo) CallJenius() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "6288xx",
		Status:  "Active",
		Balance: "12000",
		Icon:    "https://sampleurl.com/jenius.jpg",
	}
}