package main

import "fmt"

//type house struct {
//	id      uint64
//	district string
//	roomNumber uint64
//	price uint64
//	distanceFromCenter uint64
//}

var houses = []house{
	{
		id: 				1,
		district:			"Sino",
		roomNumber:			45,
		price:              70_000,
		distanceFromCenter: 10,
	},
	{
		id: 				2,
		district:			"Somoni",
		roomNumber:			54,
		price:              110_000,
		distanceFromCenter: 3,
	},
	{
		id: 				3,
		district:			"Rudaki",
		roomNumber:			64,
		price:              95_000,
		distanceFromCenter: 1,
	},
	{
		id: 				4,
		district:			"Shohmansur",
		roomNumber:			74,
		price:              10_000,
		distanceFromCenter: 60,
	},
	{
		id: 				4,
		district:			"Shohmansur",
		roomNumber:			74,
		price:              100_000,
		distanceFromCenter: 6,
	},
}

func ExampleSortByPriceAsc() {
	fmt.Println(sortByPriceAsc(houses))
	// Output:[{4 Shohmansur 74 10000 60} {1 Sino 45 70000 10} {3 Rudaki 64 95000 1} {4 Shohmansur 74 100000 6} {2 Somoni 54 110000 3}]
}

func ExampleSortByPriceDesc() {
	fmt.Println(sortByPriceDesc(houses))
	// Output:[{2 Somoni 54 110000 3} {4 Shohmansur 74 100000 6} {3 Rudaki 64 95000 1} {1 Sino 45 70000 10} {4 Shohmansur 74 10000 60}]
}

func ExampleSortByDistanceFromCenterAsc() {
	fmt.Println(sortByDistrictAsc(houses))
	// Output:[{3 Rudaki 64 95000 1} {2 Somoni 54 110000 3} {4 Shohmansur 74 100000 6} {1 Sino 45 70000 10} {4 Shohmansur 74 10000 60}]
}

func ExampleSortByDistanceFromCenterDesc() {
	fmt.Println(sortByDistrictDesc(houses))
	// Output:[{4 Shohmansur 74 10000 60} {1 Sino 45 70000 10} {4 Shohmansur 74 100000 6} {2 Somoni 54 110000 3} {3 Rudaki 64 95000 1}]
}

func ExampleSearchByMaxPrice_ManyResults() {
	fmt.Println(searchByMaxPrice(houses, 100_000))
	// Output:[{1 Sino 45 70000 10} {3 Rudaki 64 95000 1} {4 Shohmansur 74 10000 60}]
}

func ExampleSearchByMaxPrice_NoResults() {
	fmt.Println(searchByMaxPrice(houses, 0))
	// Output:[]
}

func ExampleSearchByIntervalPrice_ManyResults() {
	fmt.Println(searchByIntervalPrice(houses, 50_000, 100_001))
	// Output: [{1 Sino 45 70000 10} {3 Rudaki 64 95000 1} {4 Shohmansur 74 100000 6}]
}

func ExampleSearchByIntervalPrice_NoResults() {
	fmt.Println(searchByIntervalPrice(houses, 0, 0))
	// Output: []
}


func ExampleSearchByDistrict_ManyResults() {
	fmt.Println(searchByDistrict(houses, "Shohmansur"))
	// Output:[{4 Shohmansur 74 10000 60} {4 Shohmansur 74 100000 6}]
}

func ExampleSearchByDistrict_NoResults() {
	fmt.Println(searchByDistrict(houses, ""))
	// Output:[]
}

func ExampleSearchByDistricts_ManyResults() {
	fmt.Println(searchByDistricts(houses, []string{"Rudaki","Shohmansur",}))
	// Output:[{3 Rudaki 64 95000 1} {4 Shohmansur 74 10000 60} {4 Shohmansur 74 100000 6}]
}

func ExampleSearchByDistricts_NoResult() {
	fmt.Println(searchByDistricts(houses, []string{"","",}))
	// Output:[]
}
