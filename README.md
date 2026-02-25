# Food Store Calculator
a Calculator program for food store. The application features in interactive CLI. 

## Setup Instructions

### Prerequisites
- Go 1.24+

### Installation guide
```bash
git clone https://github.com/actuallystonmai/food-store-calculator.git && cd food-store-calculator
go mod download
```

#### Run cli program
```bash
go run cmd/cli/main.go
```

#### Run tests
```bash
# Run all tests
go test ./internal/calculator

# Run with verbose output
go test -v ./internal/calculator

# Run with coverage
go test -v -cover ./internal/calculator
```
---

## Interactive CLI

After run, there are 5 options:
```
=========================
Food store calculator
=========================

1. Add Item
2. View Cart
3. Toggle Member Card
4. Checkout
5. Exit
```

**1. Add Item** 
  Select items and quntity.

**2. View Cart**
  See all items in the current order and member card status.

**3. Toggle Member Card**
  Activate or deactivate the 10% member discount.
  
**4. Checkout**
  Calculate final price and reset order.
  
**5. Exit**
  Close the application.

---

## Architecture

```
          Presentation
┌─────────────────────────────────┐
│          Interactive            │
│   Command-line interface (CLI)  │
└──────────────┬──────────────────┘
               │
               ▼
         Businses Logic
┌─────────────────────────────────┐
│           Calculator            │
│  - pricingRules: []PricingRule  │
└──────────────┬──────────────────┘
               │ uses
               ▼
┌──────────────────────────────────┐
│      PricingRule Interface       │
│  + Calculate(...)  float64       │
│  + Name()  string                │
└──────────────┬───────────────────┘
               │
       ┌───────┴────────┐
       │                │
       ▼                ▼
┌─────────────┐  ┌──────────────┐
│    Member   │  │    BuyTwo    │
│  Discount   │  │   Discount   │
│     Rule    │  │     Rule     │
└──────┬──────┘  └──────┬───────┘
       │                │ 
       └────────┬───────┘
                │
                ▼
              Domain
┌──────────────┐    ┌─────────────────────────┐          
│  Menu Data   │    │       Order Types       │          
│  - MenuItems │    │   - Order               │          
│  - MenuItem  │    │   - OrderItem           │          
│              │    │   - CalculationResult   │          
└──────────────┘    └─────────────────────────┘             
```      
                
#### Layer Responsilities

**Presentaion Layer** (`cmd/cli`)
  Handle user input/output and display formatted results

**Business Logic Layer** (`internal/calculator`)
  Handle calucation logic
   - Order validation
   - Discount calculation
   - Pricing rule management (Member and Buy Two discounts)

**Domain Layer** (`internal/domain`)
   Define buisiness entities
    - Menu for menu items and prices
    - Order for order structure and calculation results
   