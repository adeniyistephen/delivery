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
