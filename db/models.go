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
	FistName      string              `json:"fistname"`
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
	Id          int                 `json:"id"`
	UserId      int                 `json:"user_id"`
	Amount      float64             `json:"amount"`
	CoinAmount  float64             `json:"coin_amount"`
	CreatedBy   int                 `json:"created_by"`
	UpdatedBy   int                 `json:"updated_by"`
	LastUpdated timestamp.Timestamp `json:"last_updated"`
}

type UserType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type DropShipperId struct {
	Id int `json:"id"`
}