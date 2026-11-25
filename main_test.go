package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "negative size", size: -1, want: 0},
		{name: "zero size", size: 0, want: 0},
		{name: "single element", size: 1, want: 1},
		{name: "multiple elements", size: 100, want: 100},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			data := generateRandomElements(test.size)
			if len(data) != test.want {
				t.Fatalf("got %d elements, want %d", len(data), test.want)
			}
			for index, value := range data {
				if value <= 0 {
					t.Errorf("element %d = %d, want a positive number", index, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "single element", data: []int{7}, want: 7},
		{name: "maximum in the middle", data: []int{3, 8, 5}, want: 8},
		{name: "duplicate maxima", data: []int{4, 1, 4}, want: 4},
		{name: "all negative", data: []int{-9, -2, -7}, want: -2},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d, want %d", test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "empty slice", data: nil, want: 0},
		{name: "fewer elements than workers", data: []int{-8, -2, -4}, want: -2},
		{name: "even chunks", data: []int{1, 5, 2, 3, 7, 4, 6, 0}, want: 7},
		{name: "maximum in uneven tail", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, want: 9},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maxChunks(test.data); got != test.want {
				t.Errorf("maxChunks(%v) = %d, want %d", test.data, got, test.want)
			}
		})
	}
}
