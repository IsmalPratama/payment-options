package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallGoPay() models.PaymentMethod
	CallOVO() models.PaymentMethod
	CallDANA() models.PaymentMethod
	CallLinkAja() models.PaymentMethod
	CallKredivo() models.PaymentMethod
	CallShopeePay() models.PaymentMethod
	CallVirtualAccountBCA() models.PaymentMethod // Add this method here
}
