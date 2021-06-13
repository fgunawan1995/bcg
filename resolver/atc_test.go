package resolver

import (
	"errors"
	"testing"

	"github.com/fgunawan1995/bcg/mocks"
	"github.com/fgunawan1995/bcg/model"
	"github.com/stretchr/testify/mock"
)

func TestResolver_AddToCart(t *testing.T) {
	common := new(mocks.Common)
	usecase := new(mocks.Usecase)
	dbDAL := new(mocks.DBDAL)
	cacheDAL := new(mocks.CacheDAL)
	type args struct {
		args *struct{ Add model.AddItemToCart }
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
				&struct{ Add model.AddItemToCart }{
					Add: model.AddItemToCart{
						ItemID: "1",
						UserID: "1",
						Qty:    1,
					},
				},
			},
			mock: func() {
				usecase.On("AddItemToCart", mock.Anything).Return(nil).Times(1)
				cacheDAL.On("GetCart", mock.Anything).Return(model.Cart{}, nil).Times(1)
			},
		}, {
			name:    "error",
			wantErr: true,
			args: args{
				&struct{ Add model.AddItemToCart }{
					Add: model.AddItemToCart{
						ItemID: "1",
						UserID: "1",
						Qty:    1,
					},
				},
			},
			mock: func() {
				usecase.On("AddItemToCart", mock.Anything).Return(errors.New("aaa")).Times(1)
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
			_, err := r.AddToCart(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolver.AddToCart() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
