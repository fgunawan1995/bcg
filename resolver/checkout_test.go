package resolver

import (
	"testing"

	"github.com/fgunawan1995/bcg/mocks"
	"github.com/stretchr/testify/mock"
)

func TestResolver_Checkout(t *testing.T) {
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
		want    string
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
				usecase.On("Checkout", mock.Anything).Return("", nil).Times(1)
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
			got, err := r.Checkout(tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolver.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Resolver.Checkout() = %v, want %v", got, tt.want)
			}
		})
	}
}
