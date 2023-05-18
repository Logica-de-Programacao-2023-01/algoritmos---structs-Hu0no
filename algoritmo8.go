package main

type Travel struct {
	Origin  string
	Destiny string
	Date    string
	Price   float64
}

func MostExpensiveTravel(travels []Travel) Travel {
	var expensive Travel
	for _, travel := range travels {
		if travel.Price > expensive.Price {
			expensive = travel
		}
	}
	return expensive
}

func main() {
	travels := []Travel{
		{
			Origin:  "Brasilia",
			Destiny: "Costa Rica",
			Date:    "July 2",
			Price:   5000,
		},
		{
			Origin:  "Brasilia",
			Destiny: "Norway",
			Date:    "October 15",
			Price:   8000,
		},
		{
			Origin:  "Brasilia",
			Destiny: "Chile",
			Date:    "December 17",
			Price:   3000,
		},
	}

	expensiveTravel := MostExpensiveTravel(travels)
	println(expensiveTravel.Origin, expensiveTravel.Destiny, expensiveTravel.Date, expensiveTravel.Price)
}
