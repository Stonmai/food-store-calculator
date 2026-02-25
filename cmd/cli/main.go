package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/actuallystonmai/food-store-calculator/internal/calculator"
	"github.com/actuallystonmai/food-store-calculator/internal/domain"
)

type CLI struct {
	calc    *calculator.Calculator
	order   *domain.Order
	scanner *bufio.Scanner
}

func main() {
	cli := &CLI{
		calc: calculator.New(),
		order: &domain.Order{
			Items:    []domain.OrderItem{},
			IsMember: false,
		},
		scanner: bufio.NewScanner(os.Stdin),
	}

	cli.Run()
}

func (c *CLI) Run() {
	fmt.Println("=========================")
	fmt.Println("Food store calculator")
	fmt.Println("=========================")

	for {
		c.showMenu()
		choice := c.getInput("Your choice: ")

		switch choice {
		case "1":
			c.addItem()
		case "2":
			c.viewCart()
		case "3":
			c.toggleMember()
		case "4":
			c.checkout()
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice!, Choose choices between 1-5")
		}
	}
}

func (c *CLI) showMenu() {
	fmt.Println("\n1. Add Item")
	fmt.Println("2. View Cart")
	fmt.Println("3. Toggle Member Card")
	fmt.Println("4. Checkout")
	fmt.Println("5. Exit")
}

func (c *CLI) getInput(prompt string) string {
	fmt.Print(prompt)
	c.scanner.Scan()
	return strings.TrimSpace(c.scanner.Text())
}

func (c *CLI) viewMenu() {
	fmt.Println("\nMENU:")
	// items := domain.GetAllMenuItems()
	fmt.Println("1. Red Set     - 50 THB")
	fmt.Println("2. Green Set   - 40 THB")
	fmt.Println("3. Blue Set    - 30 THB")
	fmt.Println("4. Yellow Set  - 50 THB")
	fmt.Println("5. Pink Set    - 80 THB")
	fmt.Println("6. Purple Set  - 90 THB")
	fmt.Println("7. Orange Set  - 120 THB")
}

func (c *CLI) addItem() {
	c.viewMenu()
	choice := c.getInput("\nSelect item (1-7): ")
	qty := c.getInput("Quantity: ")

	// Convert and validate...
	itemNum, _ := strconv.Atoi(choice)
	quantity, _ := strconv.Atoi(qty)

	// Map to item type
	itemTypes := []domain.MenuItemType{
		domain.RedSet, domain.GreenSet, domain.BlueSet,
		domain.YellowSet, domain.PinkSet, domain.PurpleSet,
		domain.OrangeSet,
	}

	if itemNum >= 1 && itemNum <= 7 && quantity > 0 {
		c.order.Items = append(c.order.Items, domain.OrderItem{
			ItemType: itemTypes[itemNum-1],
			Quantity: quantity,
		})
		fmt.Println("Item added!")
	} else {
		fmt.Println("Invalid input")
	}
}

func (c *CLI) viewCart() {
	fmt.Println("\nYOUR CART:")
	if len(c.order.Items) == 0 {
		fmt.Println("Empty")
		return
	}

	for _, item := range c.order.Items {
		menuItem, _ := domain.GetMenuItem(item.ItemType)
		fmt.Printf("%s × %d - %.2f THB\n",
			menuItem.Name, item.Quantity,
			menuItem.Price*float64(item.Quantity))
	}

	if c.order.IsMember {
		fmt.Println("Member card used")
	} else {
		fmt.Println("No member card used")
	}
}

func (c *CLI) toggleMember() {
	c.order.IsMember = !c.order.IsMember
	if c.order.IsMember {
		fmt.Println("Use member card!")
	} else {
		fmt.Println("Dont use member card!")
	}
}

func (c *CLI) checkout() {
	if len(c.order.Items) == 0 {
		fmt.Println("❌ Cart is empty!")
		return
	}

	result, _ := c.calc.CalculateOrder(c.order)

	fmt.Println("==================================")
	fmt.Printf("Total price is: %.2f THB\n", result.GrandTotal)
	fmt.Println("==================================")

	// Reset order
	c.order = &domain.Order{
		Items:    []domain.OrderItem{},
		IsMember: c.order.IsMember,
	}
}
