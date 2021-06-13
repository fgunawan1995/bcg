package resolver

import (
	"testing"

	"github.com/fgunawan1995/bcg/model"
)

func Test_promoResolver_ID(t *testing.T) {
	type fields struct {
		promo model.Promo
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
			r := promoResolver{
				promo: tt.fields.promo,
			}
			if got := r.ID(); got != tt.want {
				t.Errorf("promoResolver.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_promoResolver_Name(t *testing.T) {
	type fields struct {
		promo model.Promo
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
			r := promoResolver{
				promo: tt.fields.promo,
			}
			if got := r.Name(); got != tt.want {
				t.Errorf("promoResolver.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}
