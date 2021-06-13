package usecase

import (
	"testing"

	"github.com/fgunawan1995/bcg/common"
	cachedal "github.com/fgunawan1995/bcg/dal/cache"
	dbdal "github.com/fgunawan1995/bcg/dal/db"
)

func TestNew(t *testing.T) {
	type args struct {
		common common.Common
		cache  cachedal.CacheDAL
		db     dbdal.DBDAL
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "success",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			New(tt.args.common, tt.args.cache, tt.args.db)
		})
	}
}
