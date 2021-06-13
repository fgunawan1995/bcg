package resolver

import (
	"testing"

	"github.com/fgunawan1995/bcg/model"
)

func Test_itemResolver_ID(t *testing.T) {
	type fields struct {
		item model.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := itemResolver{
				item: tt.fields.item,
			}
			if got := r.ID(); got != tt.want {
				t.Errorf("itemResolver.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemResolver_Name(t *testing.T) {
	type fields struct {
		item model.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := itemResolver{
				item: tt.fields.item,
			}
			if got := r.Name(); got != tt.want {
				t.Errorf("itemResolver.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemResolver_SKU(t *testing.T) {
	type fields struct {
		item model.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := itemResolver{
				item: tt.fields.item,
			}
			if got := r.SKU(); got != tt.want {
				t.Errorf("itemResolver.SKU() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemResolver_Price(t *testing.T) {
	type fields struct {
		item model.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := itemResolver{
				item: tt.fields.item,
			}
			if got := r.Price(); got != tt.want {
				t.Errorf("itemResolver.Price() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemResolver_CurrentStock(t *testing.T) {
	type fields struct {
		item model.Item
	}
	tests := []struct {
		name   string
		fields fields
		want   int32
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := itemResolver{
				item: tt.fields.item,
			}
			if got := r.CurrentStock(); got != tt.want {
				t.Errorf("itemResolver.CurrentStock() = %v, want %v", got, tt.want)
			}
		})
	}
}
