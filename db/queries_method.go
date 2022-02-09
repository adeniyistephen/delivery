package db

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/adeniyistephen/delivery/database"
)

var (
	ErrNotFound              = errors.New("validation error: Not found")
	ErrInvalidID             = errors.New("ID is not in its proper form")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

func (d Delivery) AddDelivery(
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
	var basePrice float64

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
	devOption, err := d.ValidateDeliveryOption(deliveryOption)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}
	fmt.Println("Delivery Option: ", devOption)

	// Validate region
	region_exists, err := d.ValidateRegion(region)
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
	dropId := d.RegionDeliveryOptionOverride(region_exists.Name, devOption.Name)
	dropshipperId = dropId.Id
	fmt.Println("Dropshipper ID: ", dropshipperId)

	// New seller query
	S, err := d.NewSellerQuery(sellerId)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}
	fmt.Println("Seller: ", S)

	// Validate dropshipper
	if err := d.ValidateDropShipper(dropshipperId); err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}

	// Validate that the seller has ENOUGH coins that he'd be able to make that transaction
	if S.CoinAmount < serviceFee {
		fmt.Println("You don't have enough coins")
	}

	// Validate Product
	uniqueTmpTable, err := d.Product_Validation(deliveryDetails, deliveryOption, region_exists.Id, sellerId, dropshipperId)
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("product validation error: %w", err)
	}

	// Get Delivery Status Id
	deliveryStatus, err := d.GetDeliveryStatusId("Proposed")
	if err != nil {
		if errors.Is(err, database.ErrDBNotFound) {
			log.Println(ErrNotFound)
		}
		fmt.Println("query: %w", err)
	}
	//Insert into delivery table
	t := time.Now()
	for _, uTT := range uniqueTmpTable {
		deliveryId, err := d.InsertIntoDelivery(sellerId, t.String(), true, name, address, region_exists.Id, serviceFee, basePrice, declaredAmount, devOption.Id, sellerId, dropshipperId, deliveryStatus.Id, contactNumber, note, uTT.TotalPriceDistributor)
		if err != nil {
			log.Println("insert into delivery error: %w", err)
		}

		// Update DeliveryDetails into delivery_details table
		if err := d.UpdateDeliveryDetails(deliveryId.Id, uTT.ProductId, uTT.Quantity, uTT.PricePerItemDistributor, uTT.TotalPriceDistributor); err != nil {
			log.Println("update delivery details error: %w", err)
		}

		// Add to delivery tracking
		if err := d.DeliveryTracking(deliveryId.Id, deliveryStatus.Id, t.String(), sellerId); err != nil {
			log.Println("update delivery tracking error: %w", err)
		}

		// Update user total to also deduct service fee
		coinAmount, err := d.UpdateUserTotal(sellerId, serviceFee)
		if err != nil {
			log.Println("update user total error: %w", err)
		}
		fmt.Println("Coin Amount: ", coinAmount.CoinAmount)

		// Get admin account
		adminAccount, err := d.GetAdminAccount()
		if err != nil {
			log.Println("get admin account error: %w", err)
		}
		fmt.Println("Admin Account: ", adminAccount.Id)

		/**
		Profits for dropshipper and admin
		*/
		if err := d.InsertCoinTransaction(adminAccount.Id, t.String(), true, sellerId, "D", serviceFee, deliveryId.Id); err != nil {
			log.Println("insert coin transaction error: %w", err)
		}

		// Add 35 coins to dropshipper (service_fee - 35)
		dropshipperTotalExists, err := d.DropShipperTotals(dropshipperId)
		if err != nil {
			log.Println("drop shipper totals error: %w", err)
		}

		if dropshipperTotalExists.Id == 0 {
			if err := d.InsertUserTotalDropshipper(dropshipperId, adminAccount.Id, t); err != nil {
				log.Println("insert drop shipper totals error: %w", err)
			}
		} else {
			if err := d.UpdateUserTotalDropshipper(dropshipperId); err != nil {
				log.Println("update drop shipper totals error: %w", err)
			}
		}

		if err := d.InsertCoinTransactionDropshipper(adminAccount.Id, t.String(), true, dropshipperId, deliveryId.Id); err != nil {
			log.Println("insert coin transaction error: %w", err)
		}

		// Add (service_fee - 35) coins to admin
		adminTotalExists, err := d.AdminTotals(adminAccount.Id)
		if err != nil {
			log.Println("admin totals error: %w", err)
		}

		if adminTotalExists.Id == 0 {
			if err := d.InsertUserTotalAdmin(adminAccount.Id, t, serviceFee); err != nil {
				log.Println("insert admin totals error: %w", err)
			}
		} else {
			if err := d.UpdateUserTotalAdmin(adminAccount.Id, serviceFee); err != nil {
				log.Println("update admin totals error: %w", err)
			}
		}

		if err := d.InsertCoinTransactionAdmin(adminAccount.Id, t.String(), true, serviceFee, deliveryId.Id); err != nil {
			log.Println("insert coin transaction error: %w", err)
		}
	}
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

func GetBasePrice(declaredAmount float64) float64 {
	var basePrice float64
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

type UniqueTmpTable struct {
	ProductId               int
	Quantity                int
	PricePerItemDistributor float64
	TotalPriceDistributor   float64
}

func (d Delivery) Product_Validation(delivery_details string, devOption string, regionId int, sellerId int, dropshipperId int) ([]UniqueTmpTable, error) {
	var utt []UniqueTmpTable
	for product, quantity := range gettingProductAndQuantity(delivery_details) {
		validateProduct, err := d.ValidateProduct(product)
		if err != nil {
			if errors.Is(err, database.ErrDBNotFound) {
				log.Println(ErrNotFound)
			}
			fmt.Println("query: %w", err)
		}
		fmt.Println("Product: ", validateProduct)

		if devOption == "Dropship" && validateProduct.Name == "Max-Cee Blister" {
			fmt.Errorf("Max-Cee Blister is not available for Dropship")
		}
		if devOption == "Parcel" && validateProduct.Name == "Max-Cee" {
			fmt.Errorf("Max-Cee is not available for Parcel")
		}
		if devOption == "Parcel" {
			inventory, err := d.ValidateQuantityInventory(validateProduct.Id, quantity, regionId, sellerId, dropshipperId)
			if err != nil {
				if errors.Is(err, database.ErrDBNotFound) {
					log.Println(ErrNotFound)
				}
				fmt.Println("query, no enough product: %w", err)
			}
			fmt.Println("Validate Quantity: ", inventory)
		}
		utt = append(utt,
			UniqueTmpTable{
				ProductId:               validateProduct.Id,
				Quantity:                quantity,
				PricePerItemDistributor: validateProduct.PricePerItemDropshipper,
				TotalPriceDistributor:   float64(quantity) * validateProduct.PricePerItemDropshipper,
			})
	}

	for _, ut := range utt {
		if devOption == "Parcel" {
			if err := d.UpdateInventorySeller(regionId, sellerId, dropshipperId, ut.Quantity); err != nil {
				if errors.Is(err, database.ErrDBNotFound) {
					log.Println(ErrNotFound)
				}
				fmt.Println("query: %w", err)
			}
		}
	}
	return utt, nil
}

func gettingProductAndQuantity(delivery_details string) map[string]int {
	product := make(map[string]bool)
	product["Max-Cee"] = true
	product["PPAR"] = true
	product["Maxijuice"] = true
	product["Tamaraw"] = true
	product["Vert"] = true
	product["Rouge"] = true
	product["Kogen"] = true
	product["Glutagen"] = true
	product["Vert Lotion"] = true
	product["Rouge Lotion"] = true
	product["Max-Cee Blister"] = true
	product["Maxigold"] = true
	product["Shakura Glutathone"] = true

	str := strings.Split(delivery_details, " ")
	productQuantity := make(map[string]int)
	for i := 0; i < len(str); i++ {
		_, ok := product[str[i]]
		if ok {
			productQuantity[str[i]], _ = strconv.Atoi(str[i-1])
		}
	}

	return productQuantity
}
