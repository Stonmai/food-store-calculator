package domain

type OrderItem struct {
	ItemType MenuItemType
	Quantity int
}

type Order struct {
	Items    []OrderItem
	IsMember bool
}

type CalculationResult struct {
	Total             float64
	MemberDiscount    float64
	PromotionDiscount float64
	TotalDiscount     float64
	GrandTotal        float64
}
