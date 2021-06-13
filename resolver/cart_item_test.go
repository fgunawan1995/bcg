package resolver

import (
	"reflect"
	"testing"

	"github.com/fgunawan1995/bcg/model"
)

func Test_cartItemResolver_Item(t *testing.T) {
	type fields struct {
		item    model.Item
		promo   *model.Promo
		cartQty int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    itemResolver
		wantErr bool
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartItemResolver{
				item:    tt.fields.item,
				promo:   tt.fields.promo,
				cartQty: tt.fields.cartQty,
			}
			got, err := r.Item()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemResolver.Item() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cartItemResolver.Item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cartItemResolver_Promo(t *testing.T) {
	type fields struct {
		item    model.Item
		promo   *model.Promo
		cartQty int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    *promoResolver
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				promo: &model.Promo{},
			},
			want: &promoResolver{
				promo: model.Promo{},
			},
		},
		{
			name: "no promo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartItemResolver{
				item:    tt.fields.item,
				promo:   tt.fields.promo,
				cartQty: tt.fields.cartQty,
			}
			got, err := r.Promo()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemResolver.Promo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cartItemResolver.Promo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cartItemResolver_CartQty(t *testing.T) {
	type fields struct {
		item    model.Item
		promo   *model.Promo
		cartQty int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    int32
		wantErr bool
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartItemResolver{
				item:    tt.fields.item,
				promo:   tt.fields.promo,
				cartQty: tt.fields.cartQty,
			}
			got, err := r.CartQty()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemResolver.CartQty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cartItemResolver.CartQty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cartItemResolver_SubTotal(t *testing.T) {
	type fields struct {
		item    model.Item
		promo   *model.Promo
		cartQty int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartItemResolver{
				item:    tt.fields.item,
				promo:   tt.fields.promo,
				cartQty: tt.fields.cartQty,
			}
			got, err := r.SubTotal()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemResolver.SubTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cartItemResolver.SubTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cartItemResolver_DiscTotal(t *testing.T) {
	type fields struct {
		item    model.Item
		promo   *model.Promo
		cartQty int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				promo: &model.Promo{
					Type: model.PromoTypeDiscount,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartItemResolver{
				item:    tt.fields.item,
				promo:   tt.fields.promo,
				cartQty: tt.fields.cartQty,
			}
			got, err := r.DiscTotal()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemResolver.DiscTotal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cartItemResolver.DiscTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cartItemResolver_Total(t *testing.T) {
	type fields struct {
		item    model.Item
		promo   *model.Promo
		cartQty int32
	}
	tests := []struct {
		name    string
		fields  fields
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				promo: &model.Promo{
					Type: model.PromoTypeDiscount,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartItemResolver{
				item:    tt.fields.item,
				promo:   tt.fields.promo,
				cartQty: tt.fields.cartQty,
			}
			got, err := r.Total()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartItemResolver.Total() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("cartItemResolver.Total() = %v, want %v", got, tt.want)
			}
		})
	}
}
