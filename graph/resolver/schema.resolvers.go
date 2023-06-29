package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/nyatta/config"
	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/99designs/gqlgen/graphql"
)

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input model.NewUser) (*model.Token, error) {
	token, err := ctx.Value("userService").(*services.UserServices).SignIn(&input)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return &model.Token{Token: *token}, nil
}

// CreateUser - resolver for createUser field
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	_, err := ctx.Value("userService").(*services.UserServices).CreateUser(&input)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.ResolverError, err)
	}
	return &model.User{}, nil
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

// UploadImage is the resolver for the uploadImage field.
func (r *mutationResolver) UploadImage(ctx context.Context, file graphql.Upload) (string, error) {
	fileLocation, err := ctx.Value("awsService").(*services.AwsServices).UploadFile(file)
	if err != nil {
		return "", fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return fileLocation, nil
}

// SendVerificationCode is the resolver for the sendVerificationCode field.
func (r *mutationResolver) SendVerificationCode(ctx context.Context, input model.VerificationInput) (*model.Status, error) {
	status, err := ctx.Value("twilioService").(*services.TwilioServices).SendVerification(input.Phone, input.CountryCode)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return &model.Status{Success: status}, nil
}

// VerifyVerificationCode is the resolver for the verifyVerificationCode field.
func (r *mutationResolver) VerifyVerificationCode(ctx context.Context, input model.VerificationInput) (*model.Status, error) {
	status, err := ctx.Value("twilioService").(*services.TwilioServices).VerifyCode(input.Phone, *input.VerifyCode, input.CountryCode)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", config.ResolverError, err)
	}
	return &model.Status{Success: status}, nil
}

// Handshake is the resolver for the handshake field.
func (r *mutationResolver) Handshake(ctx context.Context, input model.HandshakeInput) (*model.User, error) {
	foundUser, err := ctx.Value("userService").(*services.UserServices).FindUserByPhone(input.Phone)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.ResolverError, err)
	}
	return foundUser, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	user, err := ctx.Value("userService").(*services.UserServices).UpdateUser(&input)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.ResolverError, err)
	}
	return user, nil
}

// SetupProperty is the resolver for the setupProperty field.
func (r *mutationResolver) SetupProperty(ctx context.Context, input model.SetupPropertyInput) (*model.Status, error) {
	fmt.Println(input)
	return &model.Status{Success: "pending"}, nil
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

// SearchTown is the resolver for the searchTown field.
func (r *queryResolver) SearchTown(ctx context.Context, town string) ([]*model.Town, error) {
	var towns []*model.Town
	towns, err := ctx.Value("postaService").(*services.PostaServices).SearchTown(town)
	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.ResolverError, err)
	}

	return towns, nil
}

// GetTowns is the resolver for the getTowns field.
func (r *queryResolver) GetTowns(ctx context.Context) ([]*model.Town, error) {
	var towns []*model.Town
	towns, err := ctx.Value("postaService").(*services.PostaServices).GetTowns()

	if err != nil {
		return nil, fmt.Errorf("%s:%v", config.ResolverError, err)
	}

	return towns, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
