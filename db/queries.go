package db

import (
	"fmt"

	"github.com/adeniyistephen/delivery/database"
	"github.com/jmoiron/sqlx"
)

type Delivery struct {
	db sqlx.ExtContext
}

func NewDelivery(db *sqlx.DB) Delivery {
	return Delivery{db: db}
}

func (d Delivery) Tran(tx sqlx.ExtContext) Delivery {
	return Delivery{db: tx}
}

func (d Delivery) ValidateDeliveryOption(deliveryOption string) (DeliveryOption, error) {
	data := struct {
		DeliveryOption string `db:"name"`
	}{
		DeliveryOption: deliveryOption,
	}

	const q = `
	SELECT
        COUNT(id),
        id,
        name
    INTO @delivery_option_exists, @delivery_option_id, @delivery_option
    FROM delivery_option
    WHERE 1 = 1
      AND name = :name;
    IF @delivery_option_exists = 0 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Delivery type does not exist!';
    END IF;`

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
	SET @region_exists = (
        SELECT
            COUNT(id)
        FROM
            region
        WHERE 1 = 1
          AND name = :name
    );
    IF @region_exists = 0 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Region does not exist';
    END IF;
    SELECT
        id,
        name
    INTO @region_id, @region_name
    FROM region
    WHERE 1 = 1
      AND name = :name;`

	var r Region
	if err := database.NamedQueryStruct(d.db, q, data, &r); err != nil {
		return Region{}, fmt.Errorf("selecting region region[%q]: %w", region, err)
	}

	return r, nil
}

func (d Delivery) RegionDeliveryOptionOverride(regionName string, deliveryOption string) int {
	var dropshipperId int
	if regionName == "Luzon" && deliveryOption == "Parcel" {
		const q = `
		SELECT
                id
            FROM user u
            WHERE 1 = 1
              AND email = (SELECT value FROM sysparam WHERE key = 'HANDLER_PARCEL_LUZON')`
		if err := database.NamedQueryStruct(d.db, q, nil, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}

	if regionName == "Luzon" && deliveryOption == "Dropship" {
		const q = `
		SELECT
                id
            FROM user u
            WHERE 1 = 1
              AND email = (SELECT value FROM sysparam WHERE key = 'HANDLER_DROPSHIP_LUZON')`
		if err := database.NamedQueryStruct(d.db, q, nil, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}


	if regionName == "Vis/Min" && deliveryOption == "Parcel" {
		const q = `
		SELECT
                id
            FROM user u
            WHERE 1 = 1
              AND email = (SELECT value FROM sysparam WHERE key = 'HANDLER_PARCEL_VISMIN')`
		if err := database.NamedQueryStruct(d.db, q, nil, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}

	if regionName == "Vis/Min" && deliveryOption == "Dropship" {
		const q = `
		SELECT
                id
            FROM user u
            WHERE 1 = 1
              AND email = (SELECT value FROM sysparam WHERE key = 'HANDLER_DROPSHIP_VISMIN')`
		if err := database.NamedQueryStruct(d.db, q, nil, &dropshipperId); err != nil {
			fmt.Println("selecting dropshipperId: %w", err)
		}
		return dropshipperId
	}
	return dropshipperId
}

func (d Delivery) NewSellerQuery(sellerId int) (UserTotal,error) {
	data := struct {
		SellerID int `db:"id"`
	}{
		SellerID: sellerId,
	}

	const q = `
	SELECT
        COUNT(u.id),
        ut.coin_amount
    INTO @user_exists, @user_coin_balance
    FROM user u
             INNER JOIN user_total ut
                        ON  1 = 1
                            AND u.id = ut.user_id
    WHERE 1 = 1
      AND u.is_active = 1
      AND u.id = :id
      AND u.user_type_id = (SELECT id FROM user_type WHERE name = 'Seller');
    IF @user_exists = 0 THEN
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Seller does not exist';
    END IF;`

	var seller UserTotal
	if err := database.NamedQueryStruct(d.db, q, data, &seller); err != nil {
		return UserTotal{}, fmt.Errorf("selecting deliveryOption deliveryOption[%q]: %w", sellerId, err)
	}

	return seller, nil
}

func (d Delivery) ValidateDropShipper(dropshipperId int) error {
	data := struct {
		DropshipperId int `db:"id"`
	}{
		DropshipperId: dropshipperId,
	}

	const q = `
	SET @dropshipper_exists = (
        SELECT COUNT(id) FROM user
        WHERE 1 = 1
          AND is_active = 1
          AND id = :id
          AND user_type_id = (SELECT id FROM user_type WHERE name = 'Dropshipper')
    );
    IF @dropshipper_exists = 0 THEN
        SET @error_message = CONCAT('Dropshipper does not exist for user id: ', @dropshipper_exists, ' charles ', :id);
        -- SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = 'Dropshipper does not exist';
        SIGNAL SQLSTATE '45000' SET MESSAGE_TEXT = @error_message;
    END IF;`

	var u User
	if err := database.NamedQueryStruct(d.db, q, data, &u); err != nil {
		return fmt.Errorf("selecting region region[%q]: %w", dropshipperId, err)
	}

	return nil
}