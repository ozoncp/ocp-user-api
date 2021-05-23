package utils

import (
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	tables := []struct {
		description string
		slice       []int
		batchSize   int
		expected    [][]int
		err         bool
	}{
		{"NilSlice", nil, 3, nil, false},
		{"NegativeBatchSize", []int{1}, -3, nil, true},
		{"ZeroBatchSize", []int{1}, 0, nil, true},
		{"EmptySlice", []int{}, 4, [][]int{}, false},
		{"LessThanOneBatchSize", []int{1, 2, 3}, 4, [][]int{{1, 2, 3}}, false},
		{"MultipleBatchSize", []int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}, false},
		{"NotMultipleBatchSize", []int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}, false},
	}

	for _, table := range tables {
		actual, err := SplitSlice(table.slice, table.batchSize)

		if table.err {
			if err == nil {
				t.Errorf("%s: expected error, but got %v", table.description, actual)
			}
		} else {
			if !reflect.DeepEqual(actual, table.expected) {
				t.Errorf("%s: expected %v, but got %v", table.description, table.expected, actual)
			}
		}
	}
}

func TestInvertMap(t *testing.T) {
	tables := []struct {
		description string
		m           map[string]int
		expected    map[int]string
		isPanic     bool
	}{
		{"NilMap", nil, nil, false},
		{"EmptyMap", map[string]int{}, map[int]string{}, false},
		{"NoDuplicateValues", map[string]int{"1": 1, "2": 2}, map[int]string{1: "1", 2: "2"}, false},
		{"WithDuplicateValues", map[string]int{"1": 1, "2": 2, "3": 2}, map[int]string{1: "1", 2: "3"}, true},
	}

	for _, table := range tables {
		if table.isPanic {
			defer func() {
				r := recover()

				if r == nil {
					t.Errorf("%s must panic", table.description)
				}
			}()

			InvertMap(table.m)
		} else {
			actual := InvertMap(table.m)

			if !reflect.DeepEqual(actual, table.expected) {
				t.Errorf("%s: expected %v, but got %v", table.description, table.expected, actual)
			}
		}
	}
}

func TestFilterList(t *testing.T) {
	tables := []struct {
		description string
		list        []rune
		filter      []rune
		expected    []rune
	}{
		{"NilList", nil, []rune{'a'}, nil},
		{"EmptyList", []rune{}, []rune{'a'}, []rune{}},
		{"NilFilter", []rune{'a', 'b'}, nil, []rune{'a', 'b'}},
		{"EmptyFilter", []rune{'a', 'b'}, []rune{}, []rune{'a', 'b'}},
		{"AllInFilter", []rune{'a', 'b'}, []rune{'a', 'b'}, []rune{}},
		{"Filter", []rune{'a', 'b', 'c'}, []rune{'b'}, []rune{'a', 'c'}},
	}

	for _, table := range tables {
		actual := FilterList(table.list, table.filter)

		if !reflect.DeepEqual(actual, table.expected) {
			t.Errorf("%s: expected %v, but got %v", table.description, table.expected, actual)
		}
	}
}
