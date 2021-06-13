package db

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/fgunawan1995/bcg/mocks"
	"github.com/fgunawan1995/bcg/model"
	"github.com/fgunawan1995/bcg/util"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/mock"
)

func TestDAL_GetItemByIDs(t *testing.T) {
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
		want    []model.Item
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "sku", "name", "price", "qty"}).
					AddRow(1, "ABC", "abc", 50, 10)
				mock.ExpectQuery("(.*?)").WillReturnRows(rows)
			},
			want: []model.Item{
				{
					ID:    "1",
					SKU:   "ABC",
					Name:  "abc",
					Price: 50,
					Qty:   10,
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
			got, err := dal.GetItemByIDs(tt.args.itemIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("DAL.GetItemByIDs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DAL.GetCart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_impl_ReduceItemStock(t *testing.T) {
	mockTx := new(mocks.Transaction)
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		tx     util.Transaction
		itemID string
		qty    int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			mock: func() {
				mockTx.On("Exec", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dal := &impl{
				db: tt.fields.db,
			}
			tt.mock()
			if err := dal.ReduceItemStock(mockTx, tt.args.itemID, tt.args.qty); (err != nil) != tt.wantErr {
				t.Errorf("impl.ReduceItemStock() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
