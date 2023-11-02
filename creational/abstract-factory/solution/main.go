package main

import (
	"errors"
	"fmt"
	"log"
)

type Drink interface {
	Drink()
}
type Food interface {
	Eat()
}
type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (Coffee) Drink() {
	fmt.Println("It's coffee , drinkable")
}

type Beer struct{}

func (Beer) Drink() {
	fmt.Println("It's beer , drinkable")
}

type Cake struct{}

func (Cake) Eat() {
	fmt.Println("It's Cake , eatable")
}

type GrilledFish struct{}

func (GrilledFish) Eat() {
	fmt.Println("It's GrilledFish , eatable")
}

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}
type DrinkMorningVoucher struct{}

func (DrinkMorningVoucher) GetDrink() Drink {
	return Coffee{}
}
func (DrinkMorningVoucher) GetFood() Food {
	return Cake{}
}

type DrinkEveningVoucher struct{}

func (DrinkEveningVoucher) GetDrink() Drink {
	return Beer{}
}
func (DrinkEveningVoucher) GetFood() Food {
	return GrilledFish{}
}

func GetVoucherFactory(campaignName string) (VoucherAbstractFactory, error) {
	switch campaignName {
	case "creative-morning":
		return DrinkMorningVoucher{}, nil
	case "creative-evening":
		return DrinkEveningVoucher{}, nil
	default:
		return nil, errors.New("campaign not found")
	}
}

func GetVoucher(factory VoucherAbstractFactory) Voucher {
	return Voucher{
		Drink: factory.GetDrink(),
		Food:  factory.GetFood(),
	}
}

func main() {
	voucherFactory, err := GetVoucherFactory("creative-morning")
	if err != nil {
		log.Fatal(err)
	}
	myVoucher := GetVoucher(voucherFactory)
	fmt.Println("I'm happy with this Voucher and come back to use it next time:", myVoucher)
}
