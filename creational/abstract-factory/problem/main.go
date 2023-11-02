package main

import "fmt"

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

func main() {
	fmt.Println([]Voucher{
		{
			Drink: Coffee{},
			Food:  Cake{},
		},
		{
			Drink: Beer{},
			Food:  GrilledFish{},
		},
		{
			Drink: Coffee{},
			Food:  GrilledFish{},
		},
	})
}
