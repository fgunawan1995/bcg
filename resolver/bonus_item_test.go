package resolver

import (
	"reflect"
	"testing"

	"github.com/fgunawan1995/bcg/model"
)

func Test_bonusItemResolver_Item(t *testing.T) {
	type fields struct {
		item     model.Item
		bonusQty int32
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
			r := bonusItemResolver{
				item:     tt.fields.item,
				bonusQty: tt.fields.bonusQty,
			}
			got, err := r.Item()
			if (err != nil) != tt.wantErr {
				t.Errorf("bonusItemResolver.Item() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bonusItemResolver.Item() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bonusItemResolver_BonusQty(t *testing.T) {
	type fields struct {
		item     model.Item
		bonusQty int32
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
			r := bonusItemResolver{
				item:     tt.fields.item,
				bonusQty: tt.fields.bonusQty,
			}
			got, err := r.BonusQty()
			if (err != nil) != tt.wantErr {
				t.Errorf("bonusItemResolver.BonusQty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("bonusItemResolver.BonusQty() = %v, want %v", got, tt.want)
			}
		})
	}
}
