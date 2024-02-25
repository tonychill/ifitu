package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tonychill/ifitu/apis/pb/go/finance"
	finSvc "github.com/tonychill/ifitu/apis/pb/go/finance_service"
)

func (r *routerImpl) handleAddPayment(c *fiber.Ctx) error {
	resp, err := r.finImpl.AddPaymentMethod(c.Context(), &finSvc.AddPaymentMethodRequest{
		PaymentMethod: &finance.PaymentMethod{
			CustomerId: "cus_123",
		},
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "this was bad",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"rules":         "test",
		"client_secret": resp.PaymentMethod.CustomerId,
	})
}
