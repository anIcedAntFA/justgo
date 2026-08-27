package exercises

import (
	"reflect"
	"testing"
)

func TestListAppendAndValues(t *testing.T) {
	t.Skip("Chapter 08 exercise: implement Append and Values, then delete this Skip")

	cases := []struct {
		name string
		in   []int
		want []int
	}{
		{"empty", nil, nil},
		{"one", []int{1}, []int{1}},
		{"three", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var head *Node
			for _, v := range tc.in {
				head = Append(head, v)
			}

			got := Values(head)
			if len(got) == 0 && len(tc.want) == 0 {
				return // both empty — nil vs []int{} should both pass
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Values(built from %v) = %v, want %v", tc.in, got, tc.want)
			}
		})
	}
}
