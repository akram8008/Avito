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

func searchBy (insertNeededPart func(result *[]house))[]house {
	result := make([]house, 0)
	insertNeededPart(&result)
	if len(result) == 0 {
		return make([]house,0)
	}
	return result
}


func searchByMaxPrice (Houses []house, limit int64)[]house{
	return searchBy(func(result *[]house){
		for _,val:=range Houses {
			if val.price < limit {
				*result = append(*result,val)
			}
		}
	})
}

func searchByIntervalPrice (Houses []house, lowerLimit,upperLimit int64)[]house{
	 return searchBy(func(result *[]house){
	 	  for _,val:=range Houses{
	 	  	if val.price < lowerLimit {
	 	  		continue
			}
			if val.price < upperLimit {
				*result=append(*result,val)
			}
		  }
	 })
}


func searchByDistrict(Houses []house, neededDistrict string) []house {
	return searchBy(func(result *[]house){
		for _,val:=range Houses {
			if val.district == neededDistrict {
				*result = append(*result,val)
			}
		}
	})
}


func searchByDistricts(Houses []house, neededDistricts []string) []house {
	return searchBy(func(result *[]house){
		for _,val:=range Houses {
			for _,district:=range neededDistricts{
				if val.district == district {
					*result = append(*result,val)
					break
				}
			}
		}
	})
}
