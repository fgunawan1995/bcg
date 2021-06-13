package usecase

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgunawan1995/bcg/mocks"
	"github.com/fgunawan1995/bcg/model"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

func Test_impl_Checkout(t *testing.T) {
	common := new(mocks.Common)
	dbDAL := new(mocks.DBDAL)
	cacheDAL := new(mocks.CacheDAL)
	db, mockSQL, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				userID: "1",
			},
			want: "Your checkout total is $9.00",
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{
					Items: []model.CartItem{
						{
							ItemID: "1",
							Qty:    1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{
					{
						ID:    "1",
						Qty:   1,
						Price: 10,
					},
				}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(map[string]*model.Promo{
					"1": {
						ID:   "1",
						Type: model.PromoTypeDiscount,
						Details: model.PromoDetail{
							MinQty:             1,
							DiscountPercentage: 0.1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				dbDAL.On("ReduceItemStock", mock.Anything, mock.Anything, mock.Anything).Return(nil).Times(1)
				mockSQL.ExpectCommit()
				cacheDAL.On("EmptyCart", mock.Anything).Times(1)
			},
		},
		{
			name: "error out of stock",
			args: args{
				userID: "1",
			},
			want:    "",
			wantErr: true,
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{
					Items: []model.CartItem{
						{
							ItemID: "1",
							Qty:    1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{
					{
						ID:    "1",
						Qty:   0,
						Price: 10,
					},
				}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(map[string]*model.Promo{
					"1": {
						ID:   "1",
						Type: model.PromoTypeDiscount,
						Details: model.PromoDetail{
							MinQty:             1,
							DiscountPercentage: 0.1,
						},
					},
				}, nil).Times(1)
			},
		},
		{
			name: "error empty cart",
			args: args{
				userID: "1",
			},
			want:    "",
			wantErr: true,
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, nil).Times(1)
			},
		},
		{
			name: "error GetCart",
			args: args{
				userID: "1",
			},
			want:    "",
			wantErr: true,
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, errors.New("aaa")).Times(1)
			},
		},
		{
			name: "error GetItemByIDs",
			args: args{
				userID: "1",
			},
			want:    "",
			wantErr: true,
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{
					Items: []model.CartItem{
						{
							ItemID: "1",
							Qty:    1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{
					{
						ID:    "1",
						Qty:   0,
						Price: 10,
					},
				}, errors.New("aaa")).Times(1)
			},
		},
		{
			name: "error GetPromoMappedByItemID",
			args: args{
				userID: "1",
			},
			want:    "",
			wantErr: true,
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{
					Items: []model.CartItem{
						{
							ItemID: "1",
							Qty:    1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{
					{
						ID:    "1",
						Qty:   1,
						Price: 10,
					},
				}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(map[string]*model.Promo{
					"1": {
						ID:   "1",
						Type: model.PromoTypeDiscount,
						Details: model.PromoDetail{
							MinQty:             1,
							DiscountPercentage: 0.1,
						},
					},
				}, errors.New("aaa")).Times(1)
			},
		},
		{
			name: "error ReduceItemStock",
			args: args{
				userID: "1",
			},
			want:    "",
			wantErr: true,
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{
					Items: []model.CartItem{
						{
							ItemID: "1",
							Qty:    1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{
					{
						ID:    "1",
						Qty:   1,
						Price: 10,
					},
				}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(map[string]*model.Promo{
					"1": {
						ID:   "1",
						Type: model.PromoTypeDiscount,
						Details: model.PromoDetail{
							MinQty:             1,
							DiscountPercentage: 0.1,
						},
					},
				}, nil).Times(1)
				dbDAL.On("GetDB").Return(sqlx.NewDb(db, "sqlmock")).Times(1)
				mockSQL.ExpectBegin()
				dbDAL.On("ReduceItemStock", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("aaa")).Times(1)
				mockSQL.ExpectRollback()
			},
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
			got, err := u.Checkout(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("impl.Checkout() = %v, want %v", got, tt.want)
			}
		})
	}
}
