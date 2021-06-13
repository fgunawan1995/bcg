package usecase

import (
	"errors"
	"testing"

	"github.com/fgunawan1995/bcg/mocks"
	"github.com/fgunawan1995/bcg/model"
	"github.com/stretchr/testify/mock"
)

func Test_impl_AddItemToCart(t *testing.T) {
	common := new(mocks.Common)
	dbDAL := new(mocks.DBDAL)
	cacheDAL := new(mocks.CacheDAL)
	type args struct {
		input model.AddItemToCart
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				input: model.AddItemToCart{
					UserID: "1",
					ItemID: "1",
					Qty:    1,
				},
			},
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(map[string]*model.Promo{
					"1": {
						ID:   "1",
						Type: model.PromoTypeBonus,
						Details: model.PromoDetail{
							BonusItemID: "1",
							BonusQty:    1,
							PerQty:      1,
						},
					},
				}, nil).Times(1)
				cacheDAL.On("SaveCart", mock.Anything, mock.Anything).Times(1)
			},
		},
		{
			name: "error GetCart",
			args: args{
				input: model.AddItemToCart{
					UserID: "1",
					ItemID: "1",
					Qty:    1,
				},
			},
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, errors.New("test")).Times(1)
			},
			wantErr: true,
		},
		{
			name: "error GetPromoMappedByItemID",
			args: args{
				input: model.AddItemToCart{
					UserID: "1",
					ItemID: "1",
					Qty:    1,
				},
			},
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(map[string]*model.Promo{
					"1": {
						ID:   "1",
						Type: model.PromoTypeBonus,
						Details: model.PromoDetail{
							BonusItemID: "1",
							BonusQty:    1,
							PerQty:      1,
						},
					},
				}, errors.New("test")).Times(1)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &impl{
				common:   common,
				dbDAL:    dbDAL,
				cacheDAL: cacheDAL,
			}
			tt.mock()
			if err := u.AddItemToCart(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("impl.AddItemToCart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
