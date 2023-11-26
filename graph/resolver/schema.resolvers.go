package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"strconv"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

// CreateUser - resolver for createUser field
func (r *mutationResolver) SignIn(ctx context.Context, input model.NewUser) (*model.SignInResponse, error) {
	res, err := ctx.Value("userService").(*services.UserServices).SignIn(ctx, &input)
	if err != nil {
		return nil, err
	}
	return &model.SignInResponse{
		User:  res.User,
		Token: res.Token,
	}, nil
}

// CreateProperty is the resolver for the createProperty field.
func (r *mutationResolver) CreateProperty(ctx context.Context, input model.NewProperty) (*model.Property, error) {
	userId := ctx.Value("userId").(string)

	newProperty, err := ctx.Value("propertyService").(*services.PropertyServices).CreateProperty(ctx, &input, uuid.MustParse(userId))
	if err != nil {
		return nil, err
	}
	return newProperty, nil
}

// AddUnit is the resolver for the addUnit field.
func (r *mutationResolver) AddUnit(ctx context.Context, input model.UnitInput) (*model.Unit, error) {
	insertedUnit, err := ctx.Value("unitService").(*services.UnitServices).AddUnit(ctx, &input)
	if err != nil {
		return nil, err
	}
	return insertedUnit, err
}

// AddUnitTenant is the resolver for the addUnitTenant field.
func (r *mutationResolver) AddUnitTenant(ctx context.Context, input model.TenancyInput) (*model.Tenant, error) {
	insertedUnitTenancy, err := ctx.Value("tenancyService").(*services.TenancyServices).AddUnitTenancy(ctx, &input)
	if err != nil {
		return nil, err
	}
	return insertedUnitTenancy, err
}

// UploadImage is the resolver for the uploadImage field.
func (r *mutationResolver) UploadImage(ctx context.Context, file graphql.Upload) (string, error) {
	fileLocation, err := ctx.Value("awsService").(*services.AwsServices).UploadGqlFile(file)
	if err != nil {
		return "", err
	}
	return fileLocation, nil
}

// SendVerificationCode is the resolver for the sendVerificationCode field.
func (r *mutationResolver) SendVerificationCode(ctx context.Context, input model.VerificationInput) (*model.Status, error) {
	status, err := ctx.Value("twilioService").(*services.TwilioServices).SendVerification(input.Phone)
	if err != nil {
		return nil, err
	}
	return &model.Status{Success: status}, nil
}

// VerifyUserVerificationCode is the resolver for the verifyUserVerificationCode field.
func (r *mutationResolver) VerifyUserVerificationCode(ctx context.Context, input model.UserVerificationInput) (*model.Status, error) {
	return &model.Status{}, nil
}

// VerifyCaretakerVerificationCode is the resolver for the verifyCaretakerVerificationCode field.
func (r *mutationResolver) VerifyCaretakerVerificationCode(ctx context.Context, input model.CaretakerVerificationInput) (*model.Status, error) {
	status, err := ctx.Value("propertyService").(*services.PropertyServices).CaretakerPhoneVerification(ctx, &input)
	if err != nil {
		return nil, err
	}
	return status, nil
}

// Handshake is the resolver for the handshake field.
func (r *mutationResolver) Handshake(ctx context.Context, input model.HandshakeInput) (*model.User, error) {
	foundUser, err := ctx.Value("userService").(*services.UserServices).FindUserByPhone(ctx, input.Phone)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

// SaveMailing is the resolver for the saveMailing field.
func (r *mutationResolver) SaveMailing(ctx context.Context, email *string) (*model.Status, error) {
	status, err := ctx.Value("mailingService").(*services.MailingServices).SaveMailing(ctx, *email)
	if err != nil {
		return nil, err
	}
	return status, nil
}

// CreatePayment is the resolver for the createPayment field
func (r *mutationResolver) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (*model.Status, error) {
	success := &model.Status{}

	amount, err := strconv.Atoi(input.Amount)
	if err != nil {
		return nil, err
	}

	payload := model.PaystackMpesaChargePayload{
		Email:       util.GenerateRandomEmail(),
		Amount:      amount * 100,
		Currency:    "KES",
		MobileMoney: model.MobileMoneyPayload{Phone: "+" + input.Phone},
	}

	chargeRes, err := ctx.Value("paystackService").(*services.PaystackServices).ChargeMpesaPhone(ctx, payload)
	if err != nil {
		return nil, err
	}

	if chargeRes.Message == "Charge attempted" {
		success.Success = "Please complete authorization process on your mobile phone"
	} else {
		success.Success = chargeRes.Message
	}

	return success, nil
}

// UpdateUserInfo is the resolver for the updateUserInfo field.
func (r *mutationResolver) UpdateUserInfo(ctx context.Context, firstName string, lastName string, avatar string) (*model.User, error) {
	userId := ctx.Value("userId").(string)

	updatedUser, err := ctx.Value("userService").(*services.UserServices).UpdateUserInfo(ctx, uuid.MustParse(userId), firstName, lastName, avatar)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	userId := ctx.Value("userId").(string)
	foundUser, err := ctx.Value("userService").(*services.UserServices).GetUser(ctx, uuid.MustParse(userId))
	if err != nil {
		return nil, err
	}

	return foundUser, nil
}

// GetProperty is the resolver for the getProperty field.
func (r *queryResolver) GetProperty(ctx context.Context, id uuid.UUID) (*model.Property, error) {
	foundProperty, err := ctx.Value("propertyService").(*services.PropertyServices).GetProperty(ctx, id)
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return towns, nil
}

// GetTowns is the resolver for the getTowns field.
func (r *queryResolver) GetTowns(ctx context.Context) ([]*model.Town, error) {
	var towns []*model.Town
	towns, err := ctx.Value("postaService").(*services.PostaServices).GetTowns()

	if err != nil {
		return nil, err
	}

	return towns, nil
}

// GetUnits is the resolver for the getUnits field.
func (r *queryResolver) GetUnits(ctx context.Context, propertyID uuid.UUID) ([]*model.Unit, error) {
	foundUnits, err := ctx.Value("propertyService").(*services.PropertyServices).GetUnits(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	return foundUnits, nil
}

// GetPropertyTenancy is the resolver for the getPropertyTenancy field.
func (r *queryResolver) GetPropertyTenancy(ctx context.Context, propertyID uuid.UUID) ([]*model.Tenant, error) {
	return []*model.Tenant{}, nil
}

// GetUserProperties is the resolver for the getUserProperties field.
func (r *queryResolver) GetUserProperties(ctx context.Context) ([]*model.Property, error) {
	// Get user from authed user context
	userId := ctx.Value("userId").(string)
	userProperties, err := ctx.Value("propertyService").(*services.PropertyServices).PropertiesCreatedBy(ctx, uuid.MustParse(userId))
	if err != nil {
		return nil, err
	}
	return userProperties, nil
}

// ListingOverview is the resolver for the listingOverview field.
func (r *queryResolver) ListingOverview(ctx context.Context, propertyID uuid.UUID) (*model.ListingOverview, error) {
	overview, err := ctx.Value("propertyService").(*services.PropertyServices).ListingOverview(ctx, propertyID)
	if err != nil {
		return nil, err
	}
	return overview, nil
}

// RefreshToken is the resolver for the refreshToken field.
func (r *queryResolver) RefreshToken(ctx context.Context) (*model.SignInResponse, error) {
	phone := ctx.Value("phone").(string)
	res, err := ctx.Value("userService").(*services.UserServices).SignIn(ctx, &model.NewUser{Phone: phone})
	if err != nil {
		return nil, err
	}

	return &model.SignInResponse{
		User:  res.User,
		Token: res.Token,
	}, nil
}

// GetNearByUnits is the resolver for the getNearByUnits field.
func (r *queryResolver) GetNearByUnits(ctx context.Context, input model.NearByUnitsInput) ([]*model.Unit, error) {
	units, err := ctx.Value("listingService").(*services.ListingServices).GetNearByUnits(ctx, &input)
	if err != nil {
		return nil, err
	}

	return units, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
