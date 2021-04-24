package calculate

import (
	"product-service/internal/application"
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	itemTests := []int{250, 500, 1000, 2000, 5000}

	tests := []struct {
		name     string
		ordered  int
		expected []application.ItemPacksOrder
	}{
		{
			name:    "Happy Days, 1 item ordered",
			ordered: 1,
			expected: []application.ItemPacksOrder{
				{
					ItemPack:       "250 Items",
					NumberItemPack: 1,
				},
			},
		},
		{
			name:    "Happy Days, 1 item ordered",
			ordered: 250,
			expected: []application.ItemPacksOrder{
				{
					ItemPack:       "250 Items",
					NumberItemPack: 1,
				},
			},
		},
		{
			name:    "Happy Days, 251 items ordered",
			ordered: 251,
			expected: []application.ItemPacksOrder{
				{
					ItemPack:       "500 Items",
					NumberItemPack: 1,
				},
			},
		},
		{
			name:    "Happy Days, 501 items ordered",
			ordered: 501,
			expected: []application.ItemPacksOrder{
				{ItemPack: "500 Items", NumberItemPack: 1},
				{ItemPack: "250 Items", NumberItemPack: 1},
			},
		},
		{
			name:    "Happy Days, 12001 items ordered",
			ordered: 12001,
			expected: []application.ItemPacksOrder{
				{ItemPack: "5000 Items", NumberItemPack: 2},
				{ItemPack: "2000 Items", NumberItemPack: 1},
				{ItemPack: "250 Items", NumberItemPack: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := tt.expected
			got := CalcItemsWanted(tt.ordered, itemTests)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Error got %v, want %v", got, want)
			}
		})
	}
}
