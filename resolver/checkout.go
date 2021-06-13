package resolver

import (
	"github.com/fgunawan1995/bcg/util"
)

func (r *Resolver) Checkout(args *struct {
	UserID string
}) (string, error) {
	result, err := r.Usecase.Checkout(args.UserID)
	return result, util.FormatGQLError(err)
}
