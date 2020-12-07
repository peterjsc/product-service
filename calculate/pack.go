package calculate

import (
	"fmt"
	"gymshark-interview/internal/application"
	"sort"
)

func CalcItemsWanted(itemsOrdered int, itemPacks []int) []application.ItemPacksOrder {

	sort.Sort(sort.Reverse(sort.IntSlice(itemPacks)))

	packs := calculatePacks(itemsOrdered, itemPacks)

	total := sum(packs, itemPacks)

	packs = calculatePacks(total, itemPacks)

	chosenPacks := makeItemsOrdered(packs, itemPacks)

	return chosenPacks
}

func makeItemsOrdered(packs []int, packsArr []int) []application.ItemPacksOrder {

	var packsChosen []application.ItemPacksOrder
	fmt.Println(packs)
	for i, count := range packs {
		if count != 0 {

			packs := fmt.Sprintf("%v Items", packsArr[i])
			packOrder := application.ItemPacksOrder{ItemPack: packs,
				NumberItemPack: count,
			}

			packsChosen = append(packsChosen, packOrder)
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

func sum(packs []int, packsArr []int) int {
	total := 0
	for i, count := range packs {
		total += count * packsArr[i]
	}
	return total
}
