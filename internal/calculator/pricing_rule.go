package calculator

import "github.com/actuallystonmai/food-store-calculator/internal/domain"


type PricingRule interface {
	Calculate(order *domain.Order, total float64, itemPrices map[domain.MenuItemType]float64) float64
	Name() string
}

// Member discount rule -> 10% discount for members
type MemberDiscountRule struct {
	Rate float64
}

func NewMemberDiscountRule() *MemberDiscountRule {
	return &MemberDiscountRule{
		Rate: 0.10,
	}
}

func (r *MemberDiscountRule) Calculate(order *domain.Order, total float64, itemPrices map[domain.MenuItemType]float64) float64 {
	if order.IsMember {
		return total * r.Rate
	}
	return 0.0
}

func (r *MemberDiscountRule) Name() string {
	return "Member Discount"
}

// Buy 2 dicount rule -> 5% discount of buying multiple items
type BuyTwoDiscountRule struct {
	PromotionItems map[domain.MenuItemType]bool
	MinQuantity   int
	Rate          float64
}

func NewBuyTwoDiscountRule() *BuyTwoDiscountRule {
	return &BuyTwoDiscountRule{
		PromotionItems: map[domain.MenuItemType]bool{
			domain.OrangeSet: true,
			domain.PinkSet:   true,
			domain.GreenSet:  true,
		},
		MinQuantity: 2,
		Rate:        0.05,
	}
}

func (r *BuyTwoDiscountRule) Calculate(order *domain.Order, total float64, itemPrices map[domain.MenuItemType]float64) float64 {
	discount := 0.0

	for _, orderItem := range order.Items {
		if r.PromotionItems[orderItem.ItemType] && orderItem.Quantity >= r.MinQuantity {
			itemTotal := itemPrices[orderItem.ItemType]
			discount += itemTotal * r.Rate
		}
	}

	return discount
}

func (r *BuyTwoDiscountRule) Name() string {
	return "Buy Two Discount"
}
