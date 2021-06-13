package cache

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fgunawan1995/bcg/model"
	cache "github.com/patrickmn/go-cache"
)

func TestDAL_GetCart(t *testing.T) {
	mockCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	type args struct {
		userID string
	}
	tests := []struct {
		name    string
		args    args
		want    model.Cart
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				userID: "1",
			},
			mock: func() {
				mockCache.SetDefault(fmt.Sprintf(model.KeyCart, "1"), model.Cart{})
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mockCache,
			}
			tt.mock()
			got, err := dal.GetCart(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DAL.GetCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DAL.GetCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_SaveCart(t *testing.T) {
	mockCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	type args struct {
		userID string
		cart   model.Cart
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mockCache,
			}
			dal.SaveCart(tt.args.userID, tt.args.cart)
		})
	}
}

func Test_impl_EmptyCart(t *testing.T) {
	mockCache := cache.New(cache.DefaultExpiration, cache.NoExpiration)
	type args struct {
		userID string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				cacheObj: mockCache,
			}
			dal.EmptyCart(tt.args.userID)
		})
	}
}
