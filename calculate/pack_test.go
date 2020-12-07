package calculate

import (
	"gymshark-interview/internal/application"
	"reflect"
	"testing"
)

var (
	itemTests = []int{250, 500, 1000, 2000, 5000}
)

func TestCalculatePacks(t *testing.T) {

	t.Run("Happy Days, 1 item ordered", func(t *testing.T) {
		want := []application.ItemPacksOrder{
			application.ItemPacksOrder{"250 Items", 1},
		}
		got := CalcItemsWanted(1, itemTests)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 250 items ordered", func(t *testing.T) {
		want := []application.ItemPacksOrder{
			application.ItemPacksOrder{"250 Items", 1},
		}
		got := CalcItemsWanted(250, itemTests)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 251 items ordered", func(t *testing.T) {
		want := []application.ItemPacksOrder{
			application.ItemPacksOrder{"500 Items", 1},
		}
		got := CalcItemsWanted(251, itemTests)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 501 items ordered", func(t *testing.T) {
		want := []application.ItemPacksOrder{
			application.ItemPacksOrder{"500 Items", 1},
			application.ItemPacksOrder{"250 Items", 1},
		}
		got := CalcItemsWanted(501, itemTests)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 12001 items ordered", func(t *testing.T) {
		want := []application.ItemPacksOrder{
			application.ItemPacksOrder{"5000 Items", 2},
			application.ItemPacksOrder{"2000 Items", 1},
			application.ItemPacksOrder{"250 Items", 1},
		}
		got := CalcItemsWanted(12001, itemTests)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})
}
