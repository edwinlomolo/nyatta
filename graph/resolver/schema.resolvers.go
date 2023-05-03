package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
)

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.NewUser) (*model.Token, error) {
	token, err := ctx.Value("userService").(*services.UserServices).SignIn(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return &model.Token{Token: *token}, nil
}

// CreateProperty is the resolver for the createProperty field.
func (r *mutationResolver) CreateProperty(ctx context.Context, input model.NewProperty) (*model.Property, error) {
	newProperty, err := ctx.Value("propertyService").(*services.PropertyServices).CreateProperty(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return newProperty, nil
}

// AddAmenity is the resolver for the addAmenity field.
func (r *mutationResolver) AddAmenity(ctx context.Context, input model.AmenityInput) (*model.Amenity, error) {
	insertedAmenity, err := ctx.Value("amenityService").(*services.AmenityServices).AddAmenity(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return insertedAmenity, err
}

// AddPropertyUnit is the resolver for the addPropertyUnit field.
func (r *mutationResolver) AddPropertyUnit(ctx context.Context, input model.PropertyUnitInput) (*model.PropertyUnit, error) {
	insertedPropertyUnit, err := ctx.Value("unitService").(*services.UnitServices).AddPropertyUnit(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return insertedPropertyUnit, err
}

// AddUnitBedrooms is the resolver for the addUnitBedrooms field.
func (r *mutationResolver) AddUnitBedrooms(ctx context.Context, input []*model.UnitBedroomInput) ([]*model.Bedroom, error) {
	insertedUnitBedrooms, err := ctx.Value("unitService").(*services.UnitServices).AddUnitBedrooms(input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return insertedUnitBedrooms, err
}

// AddPropertyUnitTenant is the resolver for the addPropertyUnitTenant field.
func (r *mutationResolver) AddPropertyUnitTenant(ctx context.Context, input model.TenancyInput) (*model.Tenant, error) {
	insertedUnitTenancy, err := ctx.Value("tenancyService").(*services.TenancyServices).AddUnitTenancy(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return insertedUnitTenancy, err
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	foundUser, err := ctx.Value("userService").(*services.UserServices).FindById(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundUser, nil
}

// GetProperty is the resolver for the getProperty field.
func (r *queryResolver) GetProperty(ctx context.Context, id string) (*model.Property, error) {
	foundProperty, err := ctx.Value("propertyService").(*services.PropertyServices).GetProperty(id)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundProperty, nil
}

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, World", nil
}

// GetListings is the resolver for the getListings field.
func (r *queryResolver) GetListings(ctx context.Context, input model.ListingsInput) ([]*model.Property, error) {
	foundListings, err := ctx.Value("listingService").(*services.ListingServices).GetListings(input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return foundListings, nil
}

// SearchTown is the resolver for the searchTown field.
func (r *queryResolver) SearchTown(ctx context.Context, town string) ([]*model.Town, error) {
	store := ctx.Value("store").(*sql.DB)

	query := `SELECT id, town, postal_code FROM postal_towns WHERE town ~* $1`
	rows, err := store.Query(query, town)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	defer rows.Close()

	var towns []*model.Town

	for rows.Next() {
		// Scan row into variables
		var id string
		var town string
		var postalCode string

		err = rows.Scan(&id, &town, &postalCode)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
		}
		towns = append(towns, &model.Town{ID: id, Town: town, PostalCode: postalCode})
	}

	// Rows.Err will report last error encountered by Rows.Scan
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}

	return towns, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
