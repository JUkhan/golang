package examples

import (
	"sort"
)

type KData struct {
	Product     string
	Price       uint16
	Weight      uint16
	greedyValue float32
}

//this is fractional knapsack method
func Knapsack(data []KData, capacity uint16) (res float32) {

	for i, d := range data {
		data[i].greedyValue = float32(d.Price) / float32(d.Weight)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i].greedyValue > data[j].greedyValue
	})
	var weightCount uint16 = 0
	for _, v := range data {
		if weightCount+v.Weight > capacity {
			res += v.greedyValue * float32(capacity-weightCount)
			return
		} else {
			weightCount += v.Weight
			res += float32(v.Price)
		}
	}

	return
}
