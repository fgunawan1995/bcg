package resolver

import (
	"testing"

	"github.com/fgunawan1995/bcg/common"
	cachedal "github.com/fgunawan1995/bcg/dal/cache"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
	"github.com/fgunawan1995/bcg/usecase"
)

func TestNew(t *testing.T) {
	type args struct {
		usecase usecase.Usecase
		common  common.Common
		cache   cachedal.CacheDAL
		db      dbdal.DBDAL
	}
	tests := []struct {
		name string
		args args
		want *Resolver
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.usecase, tt.args.common, tt.args.cache, tt.args.db)
		})
	}
}

func TestResolver_Hello(t *testing.T) {
	type fields struct {
		DBDAL    dbdal.DBDAL
		CacheDAL cachedal.CacheDAL
		Usecase  usecase.Usecase
		Common   common.Common
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "success",
			want: "Hello, world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resolver{
				DBDAL:    tt.fields.DBDAL,
				CacheDAL: tt.fields.CacheDAL,
				Usecase:  tt.fields.Usecase,
				Common:   tt.fields.Common,
			}
			if got := r.Hello(); got != tt.want {
				t.Errorf("Resolver.Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
