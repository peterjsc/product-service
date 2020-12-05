package calculate

import (
	"reflect"
	"testing"
)

func TestCalculatePacks(t *testing.T) {

	t.Run("Happy Days, 1 item ordered", func(t *testing.T) {
		want := []string{"250 Items x 1"}
		got := itemsWanted(1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 250 items ordered", func(t *testing.T) {
		want := []string{"250 Items x 1"}
		got := itemsWanted(250)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 251 items ordered", func(t *testing.T) {
		want := []string{"500 Items x 1"}
		got := itemsWanted(251)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 501 items ordered", func(t *testing.T) {
		want := []string{"500 Items x 1", "250 Items x 1"}
		got := itemsWanted(501)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 12001 items ordered", func(t *testing.T) {
		want := []string{"5000 Items x 2", "2000 Items x 1", "250 Items x 1"}
		got := itemsWanted(12001)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})
}

func TestMinItem(t *testing.T) {

	t.Run("Happy Days, 1 item ordered", func(t *testing.T) {
		want := 250
		intArr := []int{250, 500, 1000, 2000, 5000}
		got := minItem(intArr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 250 item ordered", func(t *testing.T) {
		want := 250
		intArr := []int{250, 500, 1000, 2000, 5000}
		got := minItem(intArr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 500 item ordered", func(t *testing.T) {
		want := 250
		intArr := []int{250, 500, 1000, 2000, 5000}
		got := minItem(intArr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

	t.Run("Happy Days, 251 item ordered", func(t *testing.T) {
		want := 250
		intArr := []int{250, 500, 1000, 2000, 5000}
		got := minItem(intArr)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Error got %v, want %v", got, want)
		}
	})

}
