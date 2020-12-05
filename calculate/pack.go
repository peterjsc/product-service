package calculate

import (
	"fmt"
	"sort"
)

var (
	item_stock = map[int]string{
		5000: "5000 Items",
		2000: "2000 Items",
		1000: "1000 Items",
		500:  "500 Items",
		250:  "250 Items",
	}
)

func itemsWanted(itemsOrdered int) []string {

	packsArr := intArr(item_stock)

	sort.Sort(sort.Reverse(sort.IntSlice(packsArr)))

	packs := calculatePacks(itemsOrdered, packsArr)

	total := sum(packs, packsArr)

	packs = calculatePacks(total, packsArr)

	chosenPacks := formatPacks(packs, packsArr)

	return chosenPacks
}

func formatPacks(packs []int, packsArr []int) []string {
	packsChosen := []string{}
	fmt.Println(packs)
	for i, count := range packs {
		if count != 0 {
			packs := fmt.Sprintf("%v Items x %v", packsArr[i], count)
			packsChosen = append(packsChosen, packs)
		}
	}
	return packsChosen
}

func calculatePacks(itemsOrdered int, packsArr []int) []int {

	orderArr := []int{}
	for _, packSize := range packsArr {
		orderCount := 0

		if packSize <= itemsOrdered {
			orderCount = itemsOrdered / packSize
			itemsOrdered -= packSize * orderCount
		}

		orderArr = append(orderArr, orderCount)
	}

	if itemsOrdered > 0 {
		orderArr[len(orderArr)-1]++
	}

	return orderArr
}

func minItem(arr []int) int {
	smallest := arr[0]
	for _, val := range arr {
		if val <= smallest {
			smallest = val
		}
	}
	return smallest
}

func intArr(item_stock map[int]string) []int {
	packsArr := []int{}
	for key, _ := range item_stock {
		packsArr = append(packsArr, key)
	}
	return packsArr
}

func sum(packs []int, packsArr []int) int {
	total := 0
	for i, count := range packs {
		total += count * packsArr[i]
	}
	return total
}
