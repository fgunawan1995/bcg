package resolver

import (
	"errors"
	"testing"

	"github.com/fgunawan1995/bcg/mocks"
	"github.com/fgunawan1995/bcg/model"
	"github.com/stretchr/testify/mock"
)

func Test_cartResolver_Items(t *testing.T) {
	common := new(mocks.Common)
	dbDAL := new(mocks.DBDAL)
	tests := []struct {
		name    string
		want    []cartItemResolver
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{model.Item{}}, nil).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(make(map[string]*model.Promo), nil).Times(1)
			},
		},
		{
			name:    "error GetItemByIDs",
			wantErr: true,
			mock: func() {
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{model.Item{}}, errors.New("aaa")).Times(1)
				common.On("GetPromoMappedByItemID", mock.Anything).Return(make(map[string]*model.Promo), nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartResolver{
				dbDAL:  dbDAL,
				common: common,
			}
			tt.mock()
			_, err := r.Items()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartResolver.Items() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_cartResolver_BonusItems(t *testing.T) {
	common := new(mocks.Common)
	dbDAL := new(mocks.DBDAL)
	tests := []struct {
		name    string
		want    []bonusItemResolver
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{model.Item{}}, nil).Times(1)
			},
		},
		{
			name:    "error GetItemByIDs",
			wantErr: true,
			mock: func() {
				dbDAL.On("GetItemByIDs", mock.Anything).Return([]model.Item{model.Item{}}, errors.New("aaa")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := cartResolver{
				dbDAL:  dbDAL,
				common: common,
			}
			tt.mock()
			_, err := r.BonusItems()
			if (err != nil) != tt.wantErr {
				t.Errorf("cartResolver.BonusItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestResolver_Cart(t *testing.T) {
	common := new(mocks.Common)
	usecase := new(mocks.Usecase)
	dbDAL := new(mocks.DBDAL)
	cacheDAL := new(mocks.CacheDAL)
	type args struct {
		args *struct{ UserID string }
	}
	tests := []struct {
		name    string
		args    args
		want    cartResolver
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				&struct{ UserID string }{
					UserID: "1",
				},
			},
			mock: func() {
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, nil).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resolver{
				DBDAL:    dbDAL,
				CacheDAL: cacheDAL,
				Usecase:  usecase,
				Common:   common,
			}
			tt.mock()
			_, err := r.Cart(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolver.Cart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
