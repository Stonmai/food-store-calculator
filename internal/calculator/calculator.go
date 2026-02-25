package calculator

import (
	"fmt"
	"github.com/actuallystonmai/food-store-calculator/internal/domain"
)

type Calculator struct {
	pricingRules []PricingRule
}

func New() *Calculator {
	return &Calculator{
		pricingRules: []PricingRule{
			NewMemberDiscountRule(),
			NewBuyTwoDiscountRule(),
		},
	}
}

func (c *Calculator) AddPricingRule(rule PricingRule) {
	c.pricingRules = append(c.pricingRules, rule)
}

func (c *Calculator) CalculateOrder(order *domain.Order) (*domain.CalculationResult, error) {
	// Validate order
	if err := c.ValidateOrder(order); err != nil {
		return nil, err
	}
	
	// Calculate total
	total, itemPrices, err := c.calculateTotal(order)
	if err != nil {
		return nil, err
	}
	
	// Calculate discount
	memberDiscount := 0.0
	promotionDiscount := 0.0
	
	for _, rule := range c.pricingRules {
		discount := rule.Calculate(order, total, itemPrices)

		// Track specific discounts by type
		switch rule.(type) {
		case *MemberDiscountRule:
			memberDiscount = discount
		case *BuyTwoDiscountRule:
			promotionDiscount = discount
		}
	}

	totalDiscount := memberDiscount + promotionDiscount
	grandTotal := total - totalDiscount
	
	return &domain.CalculationResult{
		Total: total,
		MemberDiscount: memberDiscount,
		PromotionDiscount: promotionDiscount,
		TotalDiscount: totalDiscount,
		GrandTotal: grandTotal,
	}, nil
}

// Calculate total price before dicount
func (c *Calculator) calculateTotal (order *domain.Order) (float64, map[domain.MenuItemType]float64, error) {
	total := 0.0
	itemPrices := make(map[domain.MenuItemType]float64)
	
	for _, orderItem := range order.Items {
		menuItem, err := domain.GetMenuItem(orderItem.ItemType)
		if err != nil {
			return 0, nil, fmt.Errorf("get menu items %s: %w", orderItem.ItemType, err)
		}
		
		itemTotal := menuItem.Price * float64(orderItem.Quantity)
		itemPrices[orderItem.ItemType] = itemTotal
		total += itemTotal
	}
	
	return total, itemPrices, nil
}

// Validate order
func (c *Calculator) ValidateOrder(order *domain.Order) error {
	if order == nil {
		return fmt.Errorf("order cannot be nil")
	}

	if len(order.Items) == 0 {
		return fmt.Errorf("order must contain at least one item")
	}

	for _, item := range order.Items {
		if item.Quantity <= 0 {
			return fmt.Errorf("item quantity must be greater than 0")
		}
	}

	return nil
}