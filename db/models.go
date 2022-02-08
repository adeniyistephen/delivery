package db

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type DeliveryOption struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Region struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	Id            int                 `json:"id"`
	FistName      string              `json:"firstname"`
	LastName      string              `json:"lastname"`
	Email         string              `json:"email"`
	Mobile_Number string              `json:"mobile_number"`
	Password      string              `json:"password"`
	User_TypeID   int                 `json:"user_type_id"`
	CreatedBy     int                 `json:"created_by"`
	CreatedDate   timestamp.Timestamp `json:"created_date"`
	LastUpdated   timestamp.Timestamp `json:"last_updated"`
	UpdatedBy     int                 `json:"updated_by"`
	IsActive      bool                `json:"is_active"`
	BankTypeID    int                 `json:"bank_type_id"`
	BankNo        string              `json:"bank_no"`
	Address       string              `json:"address"`
	Birthday      string              `json:"birthday"`
	Gender        string              `json:"gender"`
	M88Account    string              `json:"m88_account"`
	RegionID      int                 `json:"region_id"`
}

type SysParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type UserTotal struct {
	Id          int     `json:"id"`
	UserId      int     `json:"userid"`
	Amount      float64 `json:"amount"`
	CoinAmount  float64 `json:"coinamount"`
	CreatedBy   int     `json:"createdby"`
	UpdatedBy   int     `json:"updatedby"`
	LastUpdated string  `json:"lastupdated"`
}

type UserType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DropShipperId struct {
	Id int `json:"id"`
}

type Product struct {
	Id                      int     `json:"id"`
	Name                    string  `json:"name"`
	ProductTypeID           int     `json:"producttypeid"`
	CreatedBy               int     `json:"createdby"`
	CreatedDate             string  `json:"createddate"`
	LastUpdated             string  `json:"lastupdated"`
	UpdatedBy               string  `json:"updatedby"`
	IsActive                bool    `json:"isactive"`
	URL                     string  `json:"url"`
	PricePerItem            float64 `json:"priceperitem"`
	PricePerItemDropshipper float64 `json:"priceperitemdropshipper"`
}

type Inventory struct {
	Id            int    `json:"id"`
	ProductId     int    `json:"productid"`
	CreatedBy     int    `json:"createdby"`
	CreatedDate   string `json:"createddate"`
	Quantity      int    `json:"quantity"`
	UpdatedBy     int    `json:"updatedby"`
	IsActive      bool   `json:"isactive"`
	RegionId      int    `json:"regionid"`
	SellerId      int    `json:"sellerid"`
	DropShipperId int    `json:"dropshipperid"`
}

type ProductType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DeliveryModel struct {
	Id                   int     `json:"id"`
	CreatedBy            int     `json:"createdby"`
	UpdatedBy            int     `json:"updatedby"`
	CreatedDate          string  `json:"createddate"`
	IsActive             bool    `json:"isactive"`
	Name                 string  `json:"name"`
	Address              string  `json:"address"`
	RegionID             int     `json:"regionid"`
	ServiceFee           float64 `json:"servicefee"`
	DeclaredAmount       float64 `json:"declaredamount"`
	DeliveryOptionId     int     `json:"deliveryoptionid"`
	DeliveryStatusId     int     `json:"deliverystatusid"`
	SellerId             int     `json:"sellerid"`
	DropShipperId        int     `json:"dropshipperid"`
	RiderId              int     `json:"riderid"`
	TrackingNumber       string  `json:"trackingnumber"`
	ContactNumber        string  `json:"contactnumber"`
	Note                 string  `json:"note"`
	BasePrice            float64 `json:"baseprice"`
	AmountDistributor    float64 `json:"amountdistributor"`
	VoidOrRejectedReason string  `json:"voidorrejectedreason"`
}

type DeliveryStatus struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}