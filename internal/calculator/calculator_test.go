package calculator

import (
	"github.com/actuallystonmai/food-store-calculator/internal/domain"
	"testing"
)

func TestCalculateOrder_BasicOrder(t *testing.T) {
	calc := New()

	order := &domain.Order{
		Items: []domain.OrderItem{
			{ItemType: domain.RedSet, Quantity: 1},
			{ItemType: domain.GreenSet, Quantity: 1},
		},
		IsMember: false,
	}

	result, err := calc.CalculateOrder(order)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Total != 90.0 {
		t.Errorf("Expected total 90.00, got %.2f", result.Total)
	}

	if result.GrandTotal != 90.0 {
		t.Errorf("Expected grand total 90.00, got %.2f", result.GrandTotal)
	}
}

func TestCalculateOrder_WithMemberDiscount(t *testing.T) {
	calc := New()

	order := &domain.Order{
		Items: []domain.OrderItem{
			{ItemType: domain.RedSet, Quantity: 1},
			{ItemType: domain.GreenSet, Quantity: 1},
		},
		IsMember: true,
	}

	result, err := calc.CalculateOrder(order)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Total != 90.0 {
		t.Errorf("Expected total 90.00, got %.2f", result.Total)
	}

	if result.MemberDiscount != 9.0 {
		t.Errorf("Expected member discount 9.00, got %.2f", result.MemberDiscount)
	}

	if result.GrandTotal != 81.0 {
		t.Errorf("Expected grand total 81.00, got %.2f", result.GrandTotal)
	}
}

func TestCalculateOrder_WithBuyTwoDiscount(t *testing.T) {
	calc := New()

	order := &domain.Order{
		Items: []domain.OrderItem{
			{ItemType: domain.OrangeSet, Quantity: 2},
		},
		IsMember: false,
	}

	result, err := calc.CalculateOrder(order)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Total != 240.0 {
		t.Errorf("Expected total 240.00, got %.2f", result.Total)
	}

	if result.PromotionDiscount != 12.0 {
		t.Errorf("Expected buy two discount 12.00, got %.2f", result.PromotionDiscount)
	}

	if result.GrandTotal != 228.0 {
		t.Errorf("Expected grand total 228.00, got %.2f", result.GrandTotal)
	}
}

func TestCalculateOrder_BothDiscounts(t *testing.T) {
	calc := New()

	order := &domain.Order{
		Items: []domain.OrderItem{
			{ItemType: domain.OrangeSet, Quantity: 2},
			{ItemType: domain.RedSet, Quantity: 1},
		},
		IsMember: true,
	}

	result, err := calc.CalculateOrder(order)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Total != 290.0 {
		t.Errorf("Expected total 290.00, got %.2f", result.Total)
	}

	if result.MemberDiscount != 29.0 {
		t.Errorf("Expected member discount 29.00, got %.2f", result.MemberDiscount)
	}

	if result.PromotionDiscount != 12.0 {
		t.Errorf("Expected buy two discount 12.00, got %.2f", result.PromotionDiscount)
	}

	if result.TotalDiscount != 41.0 {
		t.Errorf("Expected total discount 41.00, got %.2f", result.TotalDiscount)
	}

	if result.GrandTotal != 249.0 {
		t.Errorf("Expected grand total 249.00, got %.2f", result.GrandTotal)
	}
}

func TestMemberDiscountRule(t *testing.T) {
	rule := NewMemberDiscountRule()

	order := &domain.Order{IsMember: true}
	discount := rule.Calculate(order, 100.0, nil)

	if discount != 10.0 {
		t.Errorf("Expected 10.0, got %.2f", discount)
	}

	if rule.Name() != "Member Discount" {
		t.Errorf("Expected 'Member Discount', got %s", rule.Name())
	}
}

func TestButTwoDiscountRule(t *testing.T) {
	rule := NewBuyTwoDiscountRule()

	itemPrices := map[domain.MenuItemType]float64{
		domain.OrangeSet: 240.0,
	}

	order := &domain.Order{
		Items: []domain.OrderItem{
			{ItemType: domain.OrangeSet, Quantity: 2},
		},
	}

	discount := rule.Calculate(order, 240.0, itemPrices)

	if discount != 12.0 {
		t.Errorf("Expected 12.0, got %.2f", discount)
	}

	if rule.Name() != "Buy Two Discount" {
		t.Errorf("Expected Buy Two Discount', got %s", rule.Name())
	}
}

func TestAddPricingRule(t *testing.T) {
	calc := New()

	// Custom rule that gives 20% discount
	customRule := &MemberDiscountRule{Rate: 0.20}
	calc.AddPricingRule(customRule)

	// Should have 3 rules now (2 default + 1 custom)
	if len(calc.pricingRules) != 3 {
		t.Errorf("Expected 3 rules, got %d", len(calc.pricingRules))
	}
}

func TestValidateOrder_NilOrder(t *testing.T) {
	calc := New()

	err := calc.ValidateOrder(nil)

	if err == nil {
		t.Error("Expected error for nil order, got nil")
	}
}

func TestValidateOrder_EmptyOrder(t *testing.T) {
	calc := New()

	order := &domain.Order{
		Items:    []domain.OrderItem{},
		IsMember: false,
	}

	err := calc.ValidateOrder(order)

	if err == nil {
		t.Error("Expected error for empty order, got nil")
	}
}

func TestValidateOrder_InvalidQuantity(t *testing.T) {
	calc := New()

	order := &domain.Order{
		Items: []domain.OrderItem{
			{ItemType: domain.RedSet, Quantity: 0},
		},
		IsMember: false,
	}

	err := calc.ValidateOrder(order)

	if err == nil {
		t.Error("Expected error for invalid quantity, got nil")
	}
}
