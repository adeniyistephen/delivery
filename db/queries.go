package db

import (
	"context"
	"fmt"

	"github.com/adeniyistephen/delivery/database"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Delivery struct {
	log          *zap.SugaredLogger
	tr           database.Transactor
	db           sqlx.ExtContext
	isWithinTran bool
}

func NewDelivery(log *zap.SugaredLogger, db *sqlx.DB) Delivery {
	return Delivery{
		log: log,
		tr:  db,
		db:  db,
	}
}

func (d Delivery) WithinTran(ctx context.Context, fn func(sqlx.ExtContext) error) error {
	if d.isWithinTran {
		return fn(d.db)
	}
	return database.WithinTran(ctx, d.log, d.tr, fn)
}

func (d Delivery) Tran(tx sqlx.ExtContext) Delivery {
	return Delivery{
		log:          d.log,
		tr:           d.tr,
		db:           tx,
		isWithinTran: true,
	}
}

func (d Delivery) ValidateDeliveryOption(deliveryOption string) (DeliveryOption, error) {
	data := struct {
		DeliveryOption string `db:"name"`
	}{
		DeliveryOption: deliveryOption,
	}

	const q = `
	SELECT
        *
    FROM delivery_option
    WHERE 1 = 1
      AND name = :name`

	var doption DeliveryOption
	if err := database.NamedQueryStruct(d.db, q, data, &doption); err != nil {
		return DeliveryOption{}, fmt.Errorf("selecting deliveryOption deliveryOption[%q]: %w", deliveryOption, err)
	}

	return doption, nil
}

func (d Delivery) ValidateRegion(region string) (Region, error) {
	data := struct {
		Region string `db:"name"`
	}{
		Region: region,
	}

	const q = `
    SELECT
        *
    FROM region
    WHERE 1 = 1
      AND name = :name`

	var r Region
	if err := database.NamedQueryStruct(d.db, q, data, &r); err != nil {
		return Region{}, fmt.Errorf("selecting region region[%q]: %w", region, err)
	}

	return r, nil
}

func (d Delivery) RegionDeliveryOptionOverride(regionName string, deliveryOption string) DropShipperId {
	luzonDropship := "Max88official@gmail.com"
	luzonParcel := "marcoavilam88@gmail.com"
	visminDropship := "Max88cebustaff@gmail.com"
	visminParcel := "beautybeyondhealth@gmail.com"
	if regionName == "Luzon" && deliveryOption == "Parcel" {
		data := struct {
			LuzonParcel string `db:"email"`
		}{
			LuzonParcel: luzonParcel,
		}
		const q = `
		SELECT
                id
            FROM user
            WHERE 1 = 1
              AND email = :email`
		var dropshipperId DropShipperId
		if err := database.NamedQueryStruct(d.db, q, data, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}

	if regionName == "Luzon" && deliveryOption == "Dropship" {

		data := struct {
			LuzonDropship string `db:"email"`
		}{
			LuzonDropship: luzonDropship,
		}
		const q = `
		SELECT
                id
            FROM user
            WHERE 1 = 1
              AND email = :email`
		var dropid DropShipperId
		if err := database.NamedQueryStruct(d.db, q, data, &dropid); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropid
	}

	if regionName == "Vis/Min" && deliveryOption == "Parcel" {
		data := struct {
			VisminParcel string `db:"email"`
		}{
			VisminParcel: visminParcel,
		}
		const q = `
		SELECT
                id
            FROM user
            WHERE 1 = 1
              AND email = :email`
		var dropshipperId DropShipperId
		if err := database.NamedQueryStruct(d.db, q, data, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}

	if regionName == "Vis/Min" && deliveryOption == "Dropship" {
		data := struct {
			VisminDropship string `db:"email"`
		}{
			VisminDropship: visminDropship,
		}
		const q = `
		SELECT
                id
            FROM user
            WHERE 1 = 1
              AND email = :email`
		var dropshipperId DropShipperId
		if err := database.NamedQueryStruct(d.db, q, data, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}
	return DropShipperId{}
}

func (d Delivery) NewSellerQuery(sellerId int) (UserTotal, error) {
	data := struct{}{}
	const qt = `
		SELECT * FROM user_type WHERE name = 'Seller'`
	var user_type UserType
	if err := database.NamedQueryStruct(d.db, qt, data, &user_type); err != nil {
		return UserTotal{}, fmt.Errorf("selecting seller[%q]: %w", sellerId, err)
	}
	dataUser_type := struct {
		SellerID    int `db:"s_id"`
		UserType_Id int `db:"ut_id"`
	}{
		UserType_Id: user_type.Id,
		SellerID:    sellerId,
	}
	const qu = `
		SELECT id FROM user WHERE 1 = 1 AND is_active = 1 AND id = :s_id AND user_type_id = :ut_id`
	var user User
	if err := database.NamedQueryStruct(d.db, qu, dataUser_type, &user); err != nil {
		return UserTotal{}, fmt.Errorf("selecting User[%q]: %w", sellerId, err)
	}
	dataUserTotal := struct {
		UserID int `db:"u_id"`
	}{
		UserID: user.Id,
	}
	const qut = `
		SELECT
			*
		FROM user_total
		WHERE 1 = 1
		    AND userid = :u_id`

	var seller UserTotal
	if err := database.NamedQueryStruct(d.db, qut, dataUserTotal, &seller); err != nil {
		return UserTotal{}, fmt.Errorf("selecting seller[%q]: %w", sellerId, err)
	}
	return seller, nil
}

func (d Delivery) ValidateDropShipper(dropshipperId int) error {
	data := struct{}{}
	const qd = `
		SELECT id FROM user_type WHERE name = 'Dropshipper'`
	var user_type UserType
	if err := database.NamedQueryStruct(d.db, qd, data, &user_type); err != nil {
		return fmt.Errorf("selecting seller[%q]: %w", dropshipperId, err)
	}
	dataUser_type := struct {
		DropShipperID int `db:"d_id"`
		UserType_Id   int `db:"ut_id"`
	}{
		UserType_Id:   user_type.Id,
		DropShipperID: dropshipperId,
	}
	const qu = `
	        SELECT id FROM user
	        WHERE 1 = 1
	          AND is_active = 1
	          AND id = :d_id
	          AND user_type_id = :ut_id`
	var u User
	if err := database.NamedQueryStruct(d.db, qu, dataUser_type, &u); err != nil {
		return fmt.Errorf("error selecting dropshipper [%q]: %w", dropshipperId, err)
	}

	return nil
}

func (d Delivery) ValidateProduct(productName string) (Product, error) {

	data := struct {
		ProductName string `db:"name"`
	}{
		ProductName: productName,
	}

	const q = `
	SELECT
        *
    FROM product
    WHERE 1 = 1
      AND name = :name
	  AND isactive = 1`

	var product Product
	if err := database.NamedQueryStruct(d.db, q, data, &product); err != nil {
		return Product{}, fmt.Errorf("selecting product[%q]: %w", productName, err)
	}
	return product, nil
}

func (d Delivery) ValidateQuantityInventory(productId int, quantity, regionId, sellerId, dropshipperId int) (Inventory, error) {
	data := struct {
		ProductId     int `db:"pid"`
		Quantity      int `db:"qty"`
		RegionId      int `db:"rid"`
		SellerId      int `db:"sid"`
		DropShipperId int `db:"did"`
	}{
		ProductId:     productId,
		Quantity:      quantity,
		RegionId:      regionId,
		SellerId:      sellerId,
		DropShipperId: dropshipperId,
	}

	const q = `
	SELECT
		*
	FROM inventory
	WHERE 1 = 1
		AND productid = :pid
		AND quantity >= :qty
		AND regionid = :rid
		AND sellerid = :sid
		AND dropshipperid = :did`

	var inventory Inventory
	if err := database.NamedQueryStruct(d.db, q, data, &inventory); err != nil {
		return Inventory{}, fmt.Errorf("selecting & validating inventory[%q]: %w", productId, err)
	}
	return inventory, nil
}

func (d Delivery) UpdateInventorySeller(regionId, sellerId, dropshipperId, quantity int) error {
	data := struct {
		RegionId      int `db:"rid"`
		SellerId      int `db:"sid"`
		DropShipperId int `db:"did"`
		Quantity      int `db:"qty"`
	}{
		RegionId:      regionId,
		SellerId:      sellerId,
		DropShipperId: dropshipperId,
		Quantity:      quantity,
	}

	const q = `
	UPDATE 
		inventory
	SET 
		quantity = quantity - :qty
	WHERE 1 = 1
		AND regionid = :rid
		AND sellerid = :sid
		AND dropshipperid = :did`

	if err := database.NamedExecContext(d.db, q, data); err != nil {
		return fmt.Errorf("updating inventory[%q]: %w", regionId, err)
	}
	return nil
}

func (d Delivery) InsertIntoDelivery(
	createdBy int,
	createdDate string,
	isActive bool,
	name string,
	address string,
	regionId int,
	serviceFee float64,
	basePrice float64,
	declaredAmount float64,
	deliveryOptionId int,
	sellerId int,
	dropshipperId int,
	deliveryStatusId int,
	contactNumber string,
	note string,
	amountDistributor float64,
) (DeliveryID, error) {
	data := struct {
		CreatedBy         int     `db:"created_by"`
		CreatedDate       string  `db:"created_date"`
		IsActive          bool    `db:"is_active"`
		Name              string  `db:"name"`
		Address           string  `db:"address"`
		RegionId          int     `db:"region_id"`
		ServiceFee        float64 `db:"service_fee"`
		BasePrice         float64 `db:"base_price"`
		DeclaredAmount    float64 `db:"declared_amount"`
		DeliveryOptionId  int     `db:"delivery_option_id"`
		SellerId          int     `db:"seller_id"`
		DropshipperId     int     `db:"dropshipper_id"`
		DeliveryStatusId  int     `db:"delivery_status_id"`
		ContactNumber     string  `db:"contact_number"`
		Note              string  `db:"note"`
		AmountDistributor float64 `db:"amount_distributor"`
	}{
		CreatedBy:         createdBy,
		CreatedDate:       createdDate,
		IsActive:          isActive,
		Name:              name,
		Address:           address,
		RegionId:          regionId,
		ServiceFee:        serviceFee,
		BasePrice:         basePrice,
		DeclaredAmount:    declaredAmount,
		DeliveryOptionId:  deliveryOptionId,
		SellerId:          sellerId,
		DropshipperId:     dropshipperId,
		DeliveryStatusId:  deliveryStatusId,
		ContactNumber:     contactNumber,
		Note:              note,
		AmountDistributor: amountDistributor,
	}

	const q = `
	INSERT INTO delivery
		(createdby, createddate, isactive, name, address, regionid, servicefee, baseprice, declaredamount, deliveryoptionid, sellerid, dropshipperid, deliverystatusid, contactnumber, note, amountdistributor)
	VALUES
		(:created_by, :created_date, :is_active, :name, :address, :region_id, :service_fee, :base_price, :declared_amount, :delivery_option_id, :seller_id, :dropshipper_id, :delivery_status_id, :contact_number, :note, :amount_distributor)`

	if err := database.NamedExecContext(d.db, q, data); err != nil {
		return DeliveryID{}, fmt.Errorf("inserting into delivery: %w", err)
	}

	const q2 = `
	SELECT
		MAX(id) AS id
	FROM delivery`

	var deliveryId DeliveryID
	if err := database.NamedQueryStruct(d.db, q2, data, &deliveryId); err != nil {
		return DeliveryID{}, fmt.Errorf("inserting into delivery: %w", err)
	}
	return deliveryId, nil
}

func (d Delivery) GetDeliveryStatusId(deliveryStatus string) (DeliveryStatus, error) {
	data := struct {
		DeliveryStatus string `db:"delivery_status"`
	}{
		DeliveryStatus: deliveryStatus,
	}

	const q = `
	SELECT
		*
	FROM delivery_status
	WHERE 1 = 1
		AND name = :delivery_status`

	var dStatus DeliveryStatus
	if err := database.NamedQueryStruct(d.db, q, data, &dStatus); err != nil {
		return DeliveryStatus{}, fmt.Errorf("selecting delivery status[%q]: %w", deliveryStatus, err)
	}
	return dStatus, nil
}

func (d Delivery) UpdateDeliveryDetails(deliveryId int, productId int, quantity int, pricePerItemDistributor float64, totalPriceDistributor float64) error {
	data := struct {
		DeliveryId              int     `db:"delivery_id"`
		ProductId               int     `db:"product_id"`
		Quantity                int     `db:"quantity"`
		PricePerItemDistributor float64 `db:"price_per_item_distributor"`
		TotalPriceDistributor   float64 `db:"total_price_distributor"`
	}{
		DeliveryId:              deliveryId,
		ProductId:               productId,
		Quantity:                quantity,
		PricePerItemDistributor: pricePerItemDistributor,
		TotalPriceDistributor:   totalPriceDistributor,
	}

	const q = `
	INSERT INTO delivery_detail
		(deliveryid, productid, quantity, priceperitemdistributor, totalpricedistributor)
	VALUES
		(:delivery_id, :product_id, :quantity, :price_per_item_distributor, :total_price_distributor)`

	if err := database.NamedExecContext(d.db, q, data); err != nil {
		return fmt.Errorf("updating delivery details[%q]: %w", deliveryId, err)
	}
	return nil
}

func (d Delivery) DeliveryTracking(deliveryId int, deliveryStatusId int, lastUpdated string, sellerId int) error {
	data := struct {
		DeliveryId       int    `db:"delivery_id"`
		DeliveryStatusId int    `db:"delivery_status_id"`
		LastUpdated      string `db:"last_updated"`
		SellerId         int    `db:"seller_id"`
	}{
		DeliveryId:       deliveryId,
		DeliveryStatusId: deliveryStatusId,
		LastUpdated:      lastUpdated,
		SellerId:         sellerId,
	}

	const q = `
	INSERT INTO delivery_tracking
		(deliveryid, deliverystatusid, lastupdated, updatedby)
	VALUES
		(:delivery_id, :delivery_status_id, :last_updated, :seller_id)`

	if err := database.NamedExecContext(d.db, q, data); err != nil {
		return fmt.Errorf("inserting delivery tracking[%q]: %w", deliveryId, err)
	}
	return nil
}

func (d Delivery) UpdateUserTotal(sellerId int) error {
	data := struct {
		SellerId int `db:"seller_id"`
	}{
		SellerId: sellerId,
	}

	const q = `
	UPDATE 
		user_total
	SET 
		coinamount = coinamount - :seller_id
	WHERE 1 = 1
		AND userid = :seller_id`

	if err := database.NamedExecContext(d.db, q, data); err != nil {
		return fmt.Errorf("updating user total: %w", err)
	}
	return nil
}
