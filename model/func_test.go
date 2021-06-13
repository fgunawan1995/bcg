package model

import (
	"reflect"
	"testing"
)

func TestCart_GetCartItemIDs(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.GetCartItemIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cart.GetCartItemIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_GetBonusItemIDs(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			want: []string{"1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.GetBonusItemIDs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cart.GetBonusItemIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_GetCartItemQty(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	type args struct {
		itemID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		{
			name: "success",
			fields: fields{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			args: args{
				itemID: "1",
			},
			want: 1,
		},
		{
			name: "not found",
			fields: fields{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			args: args{
				itemID: "2",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.GetCartItemQty(tt.args.itemID); got != tt.want {
				t.Errorf("Cart.GetCartItemQty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_GetBonusItemQty(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	type args struct {
		itemID string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int32
	}{
		{
			name: "success",
			fields: fields{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			args: args{
				itemID: "1",
			},
			want: 1,
		},
		{
			name: "not found",
			fields: fields{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			args: args{
				itemID: "2",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.GetBonusItemQty(tt.args.itemID); got != tt.want {
				t.Errorf("Cart.GetBonusItemQty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_AddToCart(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	type args struct {
		input AddItemToCart
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Cart
	}{
		{
			name: "success",
			args: args{
				input: AddItemToCart{
					UserID: "1",
					ItemID: "1",
					Qty:    1,
				},
			},
			want: Cart{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
		},
		{
			name: "add existing",
			args: args{
				input: AddItemToCart{
					UserID: "1",
					ItemID: "1",
					Qty:    1,
				},
			},
			fields: fields{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			want: Cart{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    2,
					},
				},
			},
		},
		{
			name: "remove",
			args: args{
				input: AddItemToCart{
					UserID: "1",
					ItemID: "1",
					Qty:    1,
				},
			},
			fields: fields{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    -1,
					},
				},
			},
			want: Cart{
				Items: []CartItem{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.AddToCart(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cart.AddToCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_ResetBonusItem(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	tests := []struct {
		name   string
		fields fields
		want   Cart
	}{
		{
			name: "success",
			fields: fields{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			want: Cart{
				Bonus: []BonusItem{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.ResetBonusItem(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cart.ResetBonusItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_AddBonusItem(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	type args struct {
		input BonusItem
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Cart
	}{
		{
			name: "success",
			args: args{
				input: BonusItem{
					ItemID: "1",
					Qty:    1,
				},
			},
			want: Cart{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
		},
		{
			name: "add existing",
			args: args{
				input: BonusItem{
					ItemID: "1",
					Qty:    1,
				},
			},
			fields: fields{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			want: Cart{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    2,
					},
				},
			},
		},
		{
			name: "remove",
			args: args{
				input: BonusItem{
					ItemID: "1",
					Qty:    1,
				},
			},
			fields: fields{
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    -1,
					},
				},
			},
			want: Cart{
				Bonus: []BonusItem{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.AddBonusItem(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cart.AddBonusItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCart_GetQtyMap(t *testing.T) {
	type fields struct {
		Items []CartItem
		Bonus []BonusItem
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]int32
	}{
		{
			name: "success",
			fields: fields{
				Items: []CartItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
				Bonus: []BonusItem{
					{
						ItemID: "1",
						Qty:    1,
					},
				},
			},
			want: map[string]int32{
				"1": 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Cart{
				Items: tt.fields.Items,
				Bonus: tt.fields.Bonus,
			}
			if got := p.GetQtyMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cart.GetQtyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
