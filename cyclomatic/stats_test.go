package cyclomatic

import (
	"reflect"
	"testing"
)

func TestAverageComplexity(t *testing.T) {
	tests := []struct {
		stats Stats
		want  float64
	}{
		{Stats{
			{Complexity: 2},
		}, 2},
		{Stats{
			{Complexity: 2},
			{Complexity: 3},
		}, 2.5},
		{Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 4},
		}, 3},
		{Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 3},
			{Complexity: 3},
		}, 2.75},
	}
	for _, tt := range tests {
		got := tt.stats.AverageComplexity()
		if got != tt.want {
			t.Errorf("Average complexity for %q, got: %g, want: %g", tt.stats, got, tt.want)
		}
	}
}

func TestTotalComplexity(t *testing.T) {
	tests := []struct {
		stats Stats
		want  uint64
	}{
		{Stats{
			{Complexity: 2},
		}, 2},
		{Stats{
			{Complexity: 2},
			{Complexity: 3},
		}, 5},
		{Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 4},
		}, 9},
		{Stats{
			{Complexity: 2},
			{Complexity: 3},
			{Complexity: 3},
			{Complexity: 3},
		}, 11},
	}
	for _, tt := range tests {
		got := tt.stats.TotalComplexity()
		if got != tt.want {
			t.Errorf("Total complexity for %q, got: %d, want: %d", tt.stats, got, tt.want)
		}
	}
}

func TestSortAndFilter(t *testing.T) {
	tests := []struct {
		stats Stats
		top   int
		over  int
		want  Stats
	}{
		{
			stats: Stats{
				{Complexity: 1},
				{Complexity: 4},
				{Complexity: 2},
				{Complexity: 3},
			},
			top: -1, over: 0,
			want: Stats{
				{Complexity: 4},
				{Complexity: 3},
				{Complexity: 2},
				{Complexity: 1},
			},
		},
		{
			stats: Stats{
				{Complexity: 1},
				{Complexity: 2},
				{Complexity: 3},
				{Complexity: 4},
			},
			top: 2, over: 0,
			want: Stats{
				{Complexity: 4},
				{Complexity: 3},
			},
		},
		{
			stats: Stats{
				{Complexity: 1},
				{Complexity: 2},
				{Complexity: 4},
				{Complexity: 4},
				{Complexity: 5},
			},
			top: -1, over: 3,
			want: Stats{
				{Complexity: 5},
				{Complexity: 4},
				{Complexity: 4},
			},
		},
		{
			stats: Stats{
				{Complexity: 1},
				{Complexity: 2},
				{Complexity: 3},
				{Complexity: 4},
				{Complexity: 5},
			},
			top: 2, over: 2,
			want: Stats{
				{Complexity: 5},
				{Complexity: 4},
			},
		},
	}
	for _, tt := range tests {
		got := tt.stats.SortAndFilter(tt.top, tt.over)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Sort and filter (top %d over %d) for %q, got: %q, want: %q",
				tt.top, tt.over, tt.stats, got, tt.want)
		}
	}
}
