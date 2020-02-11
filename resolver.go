package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/jmoiron/sqlx"
	"github.com/scottwhite/ig-grapql/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	db *sqlx.DB
}

func NewResolver(db *sqlx.DB) *Resolver {
	return &Resolver{
		db: db,
	}
}

func (r *Resolver) Ci() CiResolver {
	return &ciResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type ciResolver struct{ *Resolver }

func (r *ciResolver) Location(ctx context.Context, obj *models.Ci) (*models.Location, error) {
	locs := make([]*models.Location, 0)
	err := r.db.Select(&locs, "select * from Location where id = $1", obj.LocationId)
	if err != nil {
		return nil, err
	}

	return locs[0], nil
}
func (r *ciResolver) OsVersion(ctx context.Context, obj *models.Ci) (*OSVersion, error) {
	panic("not implemented")
}
func (r *ciResolver) CiClass(ctx context.Context, obj *models.Ci) (CiClass, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCi(ctx context.Context, input NewCi) (*models.Ci, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Cis(ctx context.Context) ([]*models.Ci, error) {
	cis := make([]*models.Ci, 0)
	requested := graphql.CollectAllFields(ctx)
	log.Println("uh what", requested)
	sql := "select * from cis"
	err := r.db.Select(&cis, sql)
	if err != nil {
		return nil, err
	}

	return cis, nil
}
