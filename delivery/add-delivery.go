package delivery

import (
	"errors"
	"fmt"
	"log"

	"github.com/adeniyistephen/delivery/database"
	"github.com/adeniyistephen/delivery/db"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Core struct {
	delivery db.Delivery
}

func NewCore(log *zap.SugaredLogger, sqlxDB *sqlx.DB) Core {
	return Core{
		delivery: db.NewDelivery(log, sqlxDB),
	}
}

var (
	ErrNotFound              = errors.New("Validation error: Not found")
	ErrInvalidID             = errors.New("ID is not in its proper form")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

func (c Core) AddDelivery(
	deliveryOption string,
	sellerId int,
	dropshipperId int,
	name string,
	contactNumber string,
	address string,
	note string,
	region string,
	serviceFee float64,
	declaredAmount float64,
	deliveryDetails string) {

	fmt.Println("Add Delivery Hit")
	var basePrice int

	// Validate declared amount
	if declaredAmount <= 0 {
		log.Fatal("Declared amount must be greater than 0")
	}

	// Modify declared amount to then be adjusted to service fee
	serviceFee = GetServiceFee(declaredAmount)
	basePrice = GetBasePrice(declaredAmount)

	// Validate additional details
	if name == "" || contactNumber == "" || note == "" {
		log.Fatal("Name, contact number, and note are required")
	}

	// Validate delivery option
	devOption, err := c.delivery.ValidateDeliveryOption(deliveryOption)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}
	fmt.Println("Delivery Option: ", devOption)

	// Validate region
	region_exists, err := c.delivery.ValidateRegion(region)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}
	fmt.Println("Region: ", region_exists)

	/*
	   Use system based config and override the dropshipper ID to use the system based rules
	*/
	dropId := c.delivery.RegionDeliveryOptionOverride(region_exists.Name, devOption.Name)
	dropshipperId = dropId.Id
	fmt.Println("Dropshipper ID: ", dropshipperId)

	// New seller query
	S, err := c.delivery.NewSellerQuery(sellerId)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}
	fmt.Println("Seller: ", S)

	// Validate dropshipper
	if err := c.delivery.ValidateDropShipper(dropshipperId); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}

	// Validate that the seller has ENOUGH coins that he'd be able to make that transaction
	if S.CoinAmount < serviceFee {
		fmt.Println("You don't have enough coins")
	}

	fmt.Println(basePrice)
}

func GetServiceFee(declaredAmount float64) float64 {
	var serviceFee float64
	if declaredAmount >= 0 && declaredAmount <= 1499 {
		serviceFee = 195
	} else if declaredAmount >= 1500 && declaredAmount <= 1999 {
		serviceFee = 200
	} else if declaredAmount >= 2000 && declaredAmount <= 2499 {
		serviceFee = 205
	} else if declaredAmount >= 2500 && declaredAmount <= 2999 {
		serviceFee = 240
	} else if declaredAmount >= 3000 && declaredAmount <= 3499 {
		serviceFee = 245
	} else if declaredAmount >= 3500 && declaredAmount <= 3999 {
		serviceFee = 250
	} else if declaredAmount >= 4000 && declaredAmount <= 4499 {
		serviceFee = 255
	} else if (declaredAmount >= 4500 && declaredAmount <= 4999) || declaredAmount > 4999 {
		serviceFee = 260
	}

	return serviceFee
}

func GetBasePrice(declaredAmount float64) int {
	var basePrice int
	if declaredAmount >= 0 && declaredAmount <= 1499 {
		basePrice = 130
	} else if declaredAmount >= 1500 && declaredAmount <= 1999 {
		declaredAmount = 135
	} else if declaredAmount >= 2000 && declaredAmount <= 2499 {
		declaredAmount = 135
	} else if declaredAmount >= 2500 && declaredAmount <= 2999 {
		declaredAmount = 175
	} else if declaredAmount >= 3000 && declaredAmount <= 3499 {
		declaredAmount = 180
	} else if declaredAmount >= 3500 && declaredAmount <= 3999 {
		declaredAmount = 185
	} else if declaredAmount >= 4000 && declaredAmount <= 4499 {
		declaredAmount = 190
	} else if (declaredAmount >= 4500 && declaredAmount <= 4999) || declaredAmount > 4999 {
		declaredAmount = 195
	}

	return basePrice
}
