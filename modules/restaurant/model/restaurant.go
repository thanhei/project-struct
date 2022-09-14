package restaurantmodel

import (
	"errors"
	"strings"
)

type Restaurant struct {
	Id     int    `json:"id" gorm:"column:id;primary_key"`
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	Id   int    `json:"id" gorm:"column:id;primary_key"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return ErrNameIsEmpty
	}

	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

var (
	ErrNameIsEmpty = errors.New("name is required")
	ErrNotFound    = errors.New("restaurant not found")
)
