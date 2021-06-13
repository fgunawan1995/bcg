package common

import (
	"errors"
	"reflect"
	"testing"

	"github.com/fgunawan1995/bcg/mocks"
	"github.com/fgunawan1995/bcg/model"
	mock "github.com/stretchr/testify/mock"
)

func Test_impl_GetPromoMappedByItemID(t *testing.T) {
	dbDAL := new(mocks.DBDAL)
	cacheDAL := new(mocks.CacheDAL)
	type args struct {
		itemIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]*model.Promo
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				itemIDs: []string{"1"},
			},
			want: map[string]*model.Promo{
				"1": {
					ID: "1",
				},
			},
			mock: func() {
				dbDAL.On("GetItemPromoByItemIDs", mock.Anything).Return([]model.ItemPromo{
					{
						ID:      "1",
						ItemID:  "1",
						PromoID: "1",
					},
				}, nil).Times(1)
				dbDAL.On("GetPromoByIDs", mock.Anything).Return([]model.Promo{
					{
						ID: "1",
					},
				}, nil).Times(1)
			},
		},
		{
			name: "error GetItemPromoByItemIDs",
			args: args{
				itemIDs: []string{"1"},
			},
			want:    make(map[string]*model.Promo),
			wantErr: true,
			mock: func() {
				dbDAL.On("GetItemPromoByItemIDs", mock.Anything).Return([]model.ItemPromo{
					{
						ID:      "1",
						ItemID:  "1",
						PromoID: "1",
					},
				}, errors.New("aaa")).Times(1)
			},
		},
		{
			name: "error GetPromoByIDs",
			args: args{
				itemIDs: []string{"1"},
			},
			want:    make(map[string]*model.Promo),
			wantErr: true,
			mock: func() {
				dbDAL.On("GetItemPromoByItemIDs", mock.Anything).Return([]model.ItemPromo{
					{
						ID:      "1",
						ItemID:  "1",
						PromoID: "1",
					},
				}, nil).Times(1)
				dbDAL.On("GetPromoByIDs", mock.Anything).Return([]model.Promo{
					{
						ID: "1",
					},
				}, errors.New("aaa")).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &impl{
				dbDAL:    dbDAL,
				cacheDAL: cacheDAL,
			}
			tt.mock()
			got, err := c.GetPromoMappedByItemID(tt.args.itemIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetPromoMappedByItemID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetPromoMappedByItemID() = %v, want %v", got, tt.want)
			}
		})
	}
}
