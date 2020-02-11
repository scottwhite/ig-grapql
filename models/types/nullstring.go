package types

import (
	"database/sql"

	"github.com/99designs/gqlgen/graphql"
	"gopkg.in/guregu/null.v3"
)

func MarshalNullString(ns null.String) graphql.Marshaler {
	if !ns.Valid {
		// this is also important, so we can detect if this scalar is used in a not null context and return an appropriate error
		return graphql.Null
	}
	return graphql.MarshalString(ns.String)
}

func UnmarshalNullString(v interface{}) (null.String, error) {
	if v == nil {
		return null.String{sql.NullString{Valid: false}}, nil
	}
	// again you can delegate to the default implementation to save yourself some work.
	s, err := graphql.UnmarshalString(v)
	return null.String{sql.NullString{String: s}}, err
}
