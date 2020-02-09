package main

import (
	"sort"
)

type house struct {
	id      uint64
	district string
	roomNumber uint64
	price int64
	distanceFromCenter uint64
}

func commonSortPart (Houses []house, compare func(a,b house)bool) []house{
	ready  := make([]house, len(Houses))
	copy(ready,Houses)
	sort.Slice(ready,func(i,j int)bool{
		return compare(ready[i],ready[j])
	})
    return ready
}

func sortByPriceAsc (Houses []house)[]house{
	return commonSortPart(Houses,func(a,b house)bool{
		return a.price < b.price
	})
}

func sortByPriceDesc (Houses []house)[]house{
	return commonSortPart(Houses,func(a,b house)bool{
		return a.price > b.price
	})
}

func sortByDistrictAsc (Houses []house)[]house{
	return commonSortPart(Houses,func(a,b house)bool{
		return a.distanceFromCenter < b.distanceFromCenter
	})
}

func sortByDistrictDesc (Houses []house)[]house{
	return commonSortPart(Houses,func(a,b house)bool{
		return a.distanceFromCenter > b.distanceFromCenter
	})
}

func searchBy (Houses []house,insertNeededPart func(home house) bool)[]house {
	result := make([]house, 0)
	for _,home:=range Houses {
		if insertNeededPart(home) {
			result = append(result, home)
		}
	}

	if len(result) == 0 {
		return make([]house,0)
	}
	return result
}


func searchByMaxPrice (Houses []house, limit int64)[]house{
	return searchBy(Houses,func(home house) bool{
		           return home.price <= limit
	})
}



func searchByIntervalPrice (Houses []house, lowerLimit,upperLimit int64)[]house{
	return searchBy(Houses,func(home house) bool{
		return  lowerLimit<= home.price && home.price <= upperLimit
	})
}


func searchByDistrict(Houses []house, neededDistrict string) []house{
	return searchBy(Houses,func(home house) bool{
		return  home.district == neededDistrict
	})
}

func searchByDistricts(Houses []house, neededDistricts []string) []house{
	return searchBy(Houses,func(home house) bool{
			for _,val :=range neededDistricts {
				if home.district == val{
					return true
				}
			}
			return false
	})
}


func main() {
}
