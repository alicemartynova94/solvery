package lesson_three_test

import (
	"github.com/stretchr/testify/assert"
	intnl "solvery/03_lesson/internal"
	"testing"
)

func TestSet_Union(t *testing.T) {
	tests := []struct {
		name   string
		setOne *intnl.Set[int]
		setTwo *intnl.Set[int]
		want   []int
	}{
		{
			name: "equal size maps",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{1: {}, 2: {}, 3: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}},
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "zero size maps",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{},
			},
			want: []int{},
		},
		{
			name: "diff size maps",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{10: {}, 15: {}},
			},
			want: []int{4, 5, 6, 7, 10, 15},
		},
		{
			name: "maps with duplicates",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			want: []int{4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setOne.Union(tt.setTwo)
			actual := make([]int, 0)
			for k, _ := range tt.setOne.values {
				actual = append(actual, k)
			}
			assert.ElementsMatch(t, tt.want, actual)
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	tests := []struct {
		name   string
		setOne *intnl.Set[int]
		setTwo *intnl.Set[int]
		want   []int
	}{
		{
			name: "no intersection",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{1: {}, 2: {}, 3: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}},
			},
			want: []int{},
		},
		{
			name: "zero size maps",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{},
			},
			want: []int{},
		},
		{
			name: "has intersection",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 7: {}},
			},
			want: []int{4, 7},
		},
		{
			name: "full intersection",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			want: []int{4, 5, 6, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intersection := tt.setOne.Intersection(tt.setTwo)
			actual := make([]int, 0)
			for k, _ := range intersection.values {
				actual = append(actual, k)
			}
			assert.ElementsMatch(t, tt.want, actual)
		})
	}
}

func TestSet_Difference(t *testing.T) {
	tests := []struct {
		name   string
		setOne *intnl.Set[int]
		setTwo *intnl.Set[int]
		want   []int
	}{
		{
			name: "same elements",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{1: {}, 2: {}, 3: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{1: {}, 2: {}, 3: {}},
			},
			want: []int{},
		},
		{
			name: "no overlap",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{1: {}, 2: {}, 3: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "partial overlap",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{4: {}, 5: {}, 6: {}, 7: {}},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{5: {}, 7: {}},
			},
			want: []int{4, 6},
		},
		{
			name: "one set is empty",
			setOne: &intnl.Set[int]{
				values: map[int]struct{}{},
			},
			setTwo: &intnl.Set[int]{
				values: map[int]struct{}{5: {}, 7: {}},
			},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intersection := tt.setOne.Difference(tt.setTwo)
			actual := make([]int, 0)
			for k, _ := range intersection.values {
				actual = append(actual, k)
			}
			assert.ElementsMatch(t, tt.want, actual)
		})
	}
}
