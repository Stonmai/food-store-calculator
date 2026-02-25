package domain

import "errors"

type MenuItem struct {
	Name string
	Price float64
}

type MenuItemType string
const (
	RedSet    MenuItemType = "RED"
	GreenSet  MenuItemType = "GREEN"
	BlueSet   MenuItemType = "BLUE"
	YellowSet MenuItemType = "YELLOW"
	PinkSet   MenuItemType = "PINK"
	PurpleSet MenuItemType = "PURPLE"
	OrangeSet MenuItemType = "ORANGE"
)


var (
	ErrMenuItemNotFound = errors.New("menu item not found")
)

var MenuItems = map[MenuItemType]*MenuItem{
	RedSet:    {Name: "Red Set", Price: 50.0},
	GreenSet:  {Name: "Green Set", Price: 40.0},
	BlueSet:   {Name: "Blue Set", Price: 30.0},
	YellowSet: {Name: "Yellow Set", Price: 50.0},
	PinkSet:   {Name: "Pink Set", Price: 80.0},
	PurpleSet: {Name: "Purple Set", Price: 90.0},
	OrangeSet: {Name: "Orange Set", Price: 120.0},
}

// Get menu item by type
func GetMenuItem(itemType MenuItemType) (*MenuItem, error) {
	item, exists := MenuItems[itemType]
	if !exists {
		return nil, ErrMenuItemNotFound
	}
	return item, nil
}

// Returns all menu items
func GetAllMenuItems() map[MenuItemType]*MenuItem {
	return MenuItems
}