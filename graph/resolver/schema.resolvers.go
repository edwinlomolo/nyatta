package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/3dw1nM0535/nyatta/graph/generated"
	"github.com/3dw1nM0535/nyatta/graph/model"
	"github.com/3dw1nM0535/nyatta/services"
	"github.com/3dw1nM0535/nyatta/util"
	"github.com/99designs/gqlgen/graphql"
	"github.com/sirupsen/logrus"
)

// CreateUser - resolver for createUser field
func (r *mutationResolver) SignIn(ctx context.Context, input model.NewUser) (*model.SignInResponse, error) {
	res, err := ctx.Value("userService").(*services.UserServices).SignIn(&input)
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
	newProperty, err := ctx.Value("propertyService").(*services.PropertyServices).CreateProperty(&input)
	if err != nil {
		return nil, err
	}
	return newProperty, nil
}

// AddPropertyUnit is the resolver for the addPropertyUnit field.
func (r *mutationResolver) AddPropertyUnit(ctx context.Context, input model.PropertyUnitInput) (*model.PropertyUnit, error) {
	insertedPropertyUnit, err := ctx.Value("unitService").(*services.UnitServices).AddPropertyUnit(&input)
	if err != nil {
		return nil, err
	}
	return insertedPropertyUnit, err
}

// AddPropertyUnitTenant is the resolver for the addPropertyUnitTenant field.
func (r *mutationResolver) AddPropertyUnitTenant(ctx context.Context, input model.TenancyInput) (*model.Tenant, error) {
	insertedUnitTenancy, err := ctx.Value("tenancyService").(*services.TenancyServices).AddUnitTenancy(&input)
	if err != nil {
		return nil, err
	}
	return insertedUnitTenancy, err
}

// UploadImage is the resolver for the uploadImage field.
func (r *mutationResolver) UploadImage(ctx context.Context, file graphql.Upload) (string, error) {
	fileLocation, err := ctx.Value("awsService").(*services.AwsServices).UploadFile(file)
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
	status, err := ctx.Value("propertyService").(*services.PropertyServices).CaretakerPhoneVerification(&input)
	if err != nil {
		return nil, err
	}
	return status, nil
}

// Handshake is the resolver for the handshake field.
func (r *mutationResolver) Handshake(ctx context.Context, input model.HandshakeInput) (*model.User, error) {
	foundUser, err := ctx.Value("userService").(*services.UserServices).FindUserByPhone(input.Phone)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

// SaveMailing is the resolver for the saveMailing field.
func (r *mutationResolver) SaveMailing(ctx context.Context, email *string) (*model.Status, error) {
	status, err := ctx.Value("mailingService").(*services.MailingServices).SaveMailing(*email)
	if err != nil {
		return nil, err
	}
	return status, nil
}

// CreatePayment is the resolver for the createPayment field
func (r *mutationResolver) CreatePayment(ctx context.Context, input model.CreatePaymentInput) (*model.Status, error) {
	success := &model.Status{}
	logger := ctx.Value("log").(*logrus.Logger)
	phone := ctx.Value("phone").(string)

	amount, err := strconv.Atoi(input.Amount)
	if err != nil {
		logger.Errorf("%s:%v", "PaystackChargeMpesaResolverError", err)
	}

	payload := services.PaystackMpesaChargePayload{
		Email:       util.GenerateRandomEmail(),
		Amount:      amount * 100,
		Currency:    "KES",
		MobileMoney: services.MobileMoneyPayload{Phone: "+" + input.Phone},
	}

	chargeRes, err := ctx.Value("paystackService").(*services.PaystackServices).ChargeMpesaPhone(phone, payload)
	if err != nil {
		logger.Errorf("%s:%v", "PaystackChargeMpesaResolverError", err)
		return nil, err
	}

	if chargeRes.Message == "Charge attempted" {
		success.Success = "Please complete authorization process on your mobile phone"
	} else {
		success.Success = chargeRes.Message
	}

	return success, nil
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, email string) (*model.User, error) {
	return &model.User{}, nil
}

// GetProperty is the resolver for the getProperty field.
func (r *queryResolver) GetProperty(ctx context.Context, id string) (*model.Property, error) {
	foundProperty, err := ctx.Value("propertyService").(*services.PropertyServices).GetProperty(id)
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

// GetPropertyUnits is the resolver for the getPropertyUnits field.
func (r *queryResolver) GetPropertyUnits(ctx context.Context, propertyID string) ([]*model.PropertyUnit, error) {
	foundUnits, err := ctx.Value("propertyService").(*services.PropertyServices).GetPropertyUnits(propertyID)
	if err != nil {
		return nil, err
	}
	return foundUnits, nil
}

// GetPropertyTenancy is the resolver for the getPropertyTenancy field.
func (r *queryResolver) GetPropertyTenancy(ctx context.Context, propertyID string) ([]*model.Tenant, error) {
	return []*model.Tenant{}, nil
}

// GetUserProperties is the resolver for the getUserProperties field.
func (r *queryResolver) GetUserProperties(ctx context.Context) ([]*model.Property, error) {
	// Get user from authed user context
	userId := ctx.Value("userId").(*string)
	userProperties, err := ctx.Value("propertyService").(*services.PropertyServices).PropertiesCreatedBy(*userId)
	if err != nil {
		return nil, err
	}
	return userProperties, nil
}

// ListingOverview is the resolver for the listingOverview field.
func (r *queryResolver) ListingOverview(ctx context.Context, propertyID string) (*model.ListingOverview, error) {
	overview, err := ctx.Value("propertyService").(*services.PropertyServices).ListingOverview(propertyID)
	if err != nil {
		return nil, err
	}
	return overview, nil
}

// GetListings is the resolver for the getListings field.
func (r *queryResolver) GetListings(ctx context.Context) ([]*model.Property, error) {
	return make([]*model.Property, 0), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
