package db

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgunawan1995/bcg/model"
	"github.com/jmoiron/sqlx"
)

func Test_impl_GetItemPromoByItemIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type args struct {
		itemIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.ItemPromo
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "item_id", "promo_id"}).
					AddRow(1, 1, 1)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: []model.ItemPromo{
				{
					ID:      "1",
					ItemID:  "1",
					PromoID: "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: sqlx.NewDb(db, "sqlmock"),
			}
			tt.mock()
			got, err := dal.GetItemPromoByItemIDs(tt.args.itemIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetItemPromoByItemIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetItemPromoByItemIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_GetPromoByIDs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	type args struct {
		itemIDs []string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Promo
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id"}).
					AddRow(1)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: []model.Promo{
				{
					ID: "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: sqlx.NewDb(db, "sqlmock"),
			}
			tt.mock()
			got, err := dal.GetPromoByIDs(tt.args.itemIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("impl.GetPromoByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("impl.GetPromoByIDs() = %v, want %v", got, tt.want)
			}
		})
	}
}
