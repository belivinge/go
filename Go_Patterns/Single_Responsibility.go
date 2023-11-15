package main

type Food_Order struct {
	OrderID int
	meal    []Menu
	cost    float64
}

type Menu struct {
	MealID int
	food   string
	cost   float64
}

func (param *Food_Order) AddMeal(menu Menu) {
	param.meal = append(param.meal, menu)
	param.cost += menu.cost
}

func (param *Food_Order) Total() float64 {
	return param.cost
}
