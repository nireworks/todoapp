package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortByTitle(t *testing.T) {
	tests := []struct {
		name     string
		unsorted []*Todo
		want     []*Todo
	}{
		{
			name: "already sorted",
			unsorted: []*Todo{
				{Id: 1, Title: "aaaa"},
				{Id: 2, Title: "bbbb"},
				{Id: 3, Title: "cccc"},
			},
			want: []*Todo{
				{Id: 1, Title: "aaaa"},
				{Id: 2, Title: "bbbb"},
				{Id: 3, Title: "cccc"},
			},
		},
		{
			name: "reverse sorted",
			unsorted: []*Todo{
				{Id: 3, Title: "cccc"},
				{Id: 2, Title: "bbbb"},
				{Id: 1, Title: "aaaa"},
			},
			want: []*Todo{
				{Id: 1, Title: "aaaa"},
				{Id: 2, Title: "bbbb"},
				{Id: 3, Title: "cccc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortByTitle(tt.unsorted)

			for idx, todo := range tt.unsorted {
				assert.Equal(t, todo, tt.want[idx])
			}
		})
	}
}

func TestSortById(t *testing.T) {
	tests := []struct {
		name     string
		unsorted []*Todo
		want     []*Todo
	}{
		{
			name: "already sorted",
			unsorted: []*Todo{
				{Id: 1, Title: "aaaa"},
				{Id: 2, Title: "bbbb"},
				{Id: 3, Title: "cccc"},
			},
			want: []*Todo{
				{Id: 1, Title: "aaaa"},
				{Id: 2, Title: "bbbb"},
				{Id: 3, Title: "cccc"},
			},
		},
		{
			name: "reverse sorted",
			unsorted: []*Todo{
				{Id: 3, Title: "cccc"},
				{Id: 2, Title: "bbbb"},
				{Id: 1, Title: "aaaa"},
			},
			want: []*Todo{
				{Id: 1, Title: "aaaa"},
				{Id: 2, Title: "bbbb"},
				{Id: 3, Title: "cccc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortById(tt.unsorted)

			for idx, todo := range tt.unsorted {
				assert.Equal(t, todo, tt.want[idx])
			}
		})
	}
}
