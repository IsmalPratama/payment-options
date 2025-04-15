package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallBCAMobile() models.PaymentMethod
	CallLinkAja() models.PaymentMethod
	CallBNIMobile() models.PaymentMethod
	CallMandiriOnline() models.PaymentMethod
	CallKredivo() models.PaymentMethod
	CallJenius() models.PaymentMethod
}