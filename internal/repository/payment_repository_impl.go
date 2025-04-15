package repository

import (
	"payment-options/internal/models"
	// "time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallGoPay() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx",
		Status:  "Active",
		Balance: "500000", 
		Icon:    "https://sampleurl.com/gopay.jpg",
	}
}

func (r *paymentRepo) CallOVO() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx",
		Status:  "Active", 
		Balance: "300000",
		Icon:    "https://sampleurl.com/ovo.jpg",
	}
}

func (r *paymentRepo) CallDANA() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx",
		Status:  "Active",
		Balance: "200000", 
		Icon:    "https://sampleurl.com/dana.jpg",
	}
}

func (r *paymentRepo) CallLinkAja() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx", 
		Status:  "Active", 
		Balance: "150000", 
		Icon:    "https://sampleurl.com/linkaja.jpg",
	}
}

func (r *paymentRepo) CallKredivo() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx",
		Status:  "Active", 
		Balance: "100000",   
		Icon:    "https://sampleurl.com/kredivo.jpg",
	}
}

func (r *paymentRepo) CallShopeePay() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "0812xx",  
		Status:  "Active",  
		Balance: "250000",  
		Icon:    "https://sampleurl.com/shopeepay.jpg", 
	}
}

func (r *paymentRepo) CallVirtualAccountBCA() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "BCA12345", 
		Status:  "Active", 
		Balance: "1000000", 
		Icon:    "https://sampleurl.com/bca_va.jpg",
	}
}
