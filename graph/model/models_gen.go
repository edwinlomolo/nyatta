// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Amenity struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Provider  *string    `json:"provider,omitempty"`
	Category  string     `json:"category"`
	UnitID    *uuid.UUID `json:"unitId,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

type AnyUpload struct {
	ID         uuid.UUID  `json:"id"`
	Upload     string     `json:"upload"`
	Category   string     `json:"category"`
	PropertyID *uuid.UUID `json:"propertyId,omitempty"`
	UserID     *uuid.UUID `json:"userId,omitempty"`
	UnitID     *uuid.UUID `json:"unitId,omitempty"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

type Bedroom struct {
	ID             uuid.UUID  `json:"id"`
	PropertyUnitID uuid.UUID  `json:"propertyUnitId"`
	BedroomNumber  int        `json:"bedroomNumber"`
	EnSuite        bool       `json:"enSuite"`
	Master         bool       `json:"master"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
}

type Caretaker struct {
	ID         uuid.UUID   `json:"id"`
	FirstName  string      `json:"first_name"`
	LastName   string      `json:"last_name"`
	Phone      string      `json:"phone"`
	Avatar     *AnyUpload  `json:"avatar,omitempty"`
	Verified   bool        `json:"verified"`
	Properties []*Property `json:"properties"`
	CreatedAt  *time.Time  `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time  `json:"updatedAt,omitempty"`
}

type CaretakerInput struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Image     string `json:"image"`
}

type CaretakerVerificationInput struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verifyCode"`
}

type CreatePaymentInput struct {
	Phone  string `json:"phone"`
	Amount string `json:"amount"`
}

type Gps struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type GpsInput struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type HandshakeInput struct {
	Phone string `json:"phone"`
}

type Invoice struct {
	ID            uuid.UUID     `json:"id"`
	Msid          string        `json:"msid"`
	Phone         string        `json:"phone"`
	Status        InvoiceStatus `json:"status"`
	WCoCheckoutID string        `json:"wCoCheckoutId"`
	CreatedAt     *time.Time    `json:"createdAt,omitempty"`
	UpdatedAt     *time.Time    `json:"updatedAt,omitempty"`
}

type ListingOverview struct {
	OccupiedUnits int `json:"occupiedUnits"`
	VacantUnits   int `json:"vacantUnits"`
	TotalUnits    int `json:"totalUnits"`
}

type NearByUnitsInput struct {
	Gps *GpsInput `json:"gps"`
}

type NewProperty struct {
	Name        string          `json:"name"`
	Type        string          `json:"type"`
	Location    *GpsInput       `json:"location"`
	Thumbnail   string          `json:"thumbnail"`
	IsCaretaker bool            `json:"isCaretaker"`
	Caretaker   *CaretakerInput `json:"caretaker,omitempty"`
}

type NewUser struct {
	Phone string `json:"phone"`
}

type Property struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Type        PropertyType `json:"type"`
	Location    *Gps         `json:"location,omitempty"`
	Thumbnail   *AnyUpload   `json:"thumbnail,omitempty"`
	Units       []*Unit      `json:"units"`
	UnitsCount  int          `json:"unitsCount"`
	CreatedBy   uuid.UUID    `json:"createdBy"`
	Caretaker   *Caretaker   `json:"caretaker,omitempty"`
	CaretakerID *uuid.UUID   `json:"caretakerId,omitempty"`
	Tenancy     []*Tenant    `json:"tenancy"`
	Owner       *User        `json:"owner,omitempty"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time   `json:"updatedAt,omitempty"`
}

type Shoot struct {
	ID         uuid.UUID  `json:"id"`
	PropertyID string     `json:"propertyId"`
	Date       time.Time  `json:"date"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

type ShootInput struct {
	Date time.Time `json:"date"`
}

type Status struct {
	Success string `json:"success"`
}

type TenancyInput struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Phone     string    `json:"phone"`
	StartDate time.Time `json:"start_date"`
	UnitID    uuid.UUID `json:"unit_id"`
}

type Tenant struct {
	ID         uuid.UUID  `json:"id"`
	StartDate  time.Time  `json:"startDate"`
	EndDate    *time.Time `json:"endDate,omitempty"`
	UnitID     uuid.UUID  `json:"unitId"`
	UserID     uuid.UUID  `json:"userId"`
	PropertyID uuid.UUID  `json:"propertyId"`
	User       *User      `json:"user"`
	Unit       *Unit      `json:"unit"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

type Token struct {
	Token string `json:"token"`
}

type Town struct {
	ID         string `json:"id"`
	Town       string `json:"town"`
	PostalCode string `json:"postalCode"`
}

type Unit struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Bedrooms    []*Bedroom   `json:"bedrooms"`
	PropertyID  uuid.UUID    `json:"propertyId"`
	Location    *Gps         `json:"location,omitempty"`
	Thumbnail   *AnyUpload   `json:"thumbnail,omitempty"`
	CaretakerID *uuid.UUID   `json:"caretakerId,omitempty"`
	Caretaker   *Caretaker   `json:"caretaker,omitempty"`
	Property    *Property    `json:"property,omitempty"`
	Tenant      *Tenant      `json:"tenant,omitempty"`
	Distance    *string      `json:"distance,omitempty"`
	Price       string       `json:"price"`
	CreatedBy   *uuid.UUID   `json:"createdBy,omitempty"`
	Owner       *User        `json:"owner,omitempty"`
	Bathrooms   int          `json:"bathrooms"`
	Amenities   []*Amenity   `json:"amenities"`
	State       UnitState    `json:"state"`
	Type        string       `json:"type"`
	Images      []*AnyUpload `json:"images"`
	Tenancy     []*Tenant    `json:"tenancy"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time   `json:"updatedAt,omitempty"`
}

type UnitAmenityInput struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type UnitBedroomInput struct {
	BedroomNumber int  `json:"bedroomNumber"`
	EnSuite       bool `json:"enSuite"`
	Master        bool `json:"master"`
}

type UnitInput struct {
	PropertyID  *uuid.UUID          `json:"propertyId,omitempty"`
	Baths       int                 `json:"baths"`
	Name        string              `json:"name"`
	Type        string              `json:"type"`
	State       UnitState           `json:"state"`
	IsCaretaker *bool               `json:"isCaretaker,omitempty"`
	Location    *GpsInput           `json:"location,omitempty"`
	Caretaker   *CaretakerInput     `json:"caretaker,omitempty"`
	Amenities   []*UnitAmenityInput `json:"amenities"`
	Bedrooms    []*UnitBedroomInput `json:"bedrooms"`
	Price       string              `json:"price"`
	Uploads     []*UploadImages     `json:"uploads,omitempty"`
}

type UploadImages struct {
	Image    string `json:"image"`
	Category string `json:"category"`
}

type User struct {
	ID               uuid.UUID   `json:"id"`
	FirstName        *string     `json:"first_name,omitempty"`
	LastName         *string     `json:"last_name,omitempty"`
	Phone            string      `json:"phone"`
	IsLandlord       bool        `json:"is_landlord"`
	Avatar           *AnyUpload  `json:"avatar,omitempty"`
	SubscribeRetries int         `json:"subscribe_retries"`
	Properties       []*Property `json:"properties"`
	Units            []*Unit     `json:"units"`
	Tenancy          []*Tenant   `json:"tenancy"`
	CreatedAt        *time.Time  `json:"createdAt,omitempty"`
	UpdatedAt        *time.Time  `json:"updatedAt,omitempty"`
}

type UserVerificationInput struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verifyCode"`
}

type VerificationInput struct {
	Phone      string  `json:"phone"`
	VerifyCode *string `json:"verifyCode,omitempty"`
}

type CountryCode string

const (
	CountryCodeKe CountryCode = "KE"
)

var AllCountryCode = []CountryCode{
	CountryCodeKe,
}

func (e CountryCode) IsValid() bool {
	switch e {
	case CountryCodeKe:
		return true
	}
	return false
}

func (e CountryCode) String() string {
	return string(e)
}

func (e *CountryCode) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CountryCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CountryCode", str)
	}
	return nil
}

func (e CountryCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InvoiceStatus string

const (
	InvoiceStatusProcessed  InvoiceStatus = "PROCESSED"
	InvoiceStatusProcessing InvoiceStatus = "PROCESSING"
)

var AllInvoiceStatus = []InvoiceStatus{
	InvoiceStatusProcessed,
	InvoiceStatusProcessing,
}

func (e InvoiceStatus) IsValid() bool {
	switch e {
	case InvoiceStatusProcessed, InvoiceStatusProcessing:
		return true
	}
	return false
}

func (e InvoiceStatus) String() string {
	return string(e)
}

func (e *InvoiceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InvoiceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InvoiceStatus", str)
	}
	return nil
}

func (e InvoiceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PropertyType string

const (
	PropertyTypeApartmentsBuilding PropertyType = "APARTMENTS_BUILDING"
	PropertyTypeApartment          PropertyType = "APARTMENT"
)

var AllPropertyType = []PropertyType{
	PropertyTypeApartmentsBuilding,
	PropertyTypeApartment,
}

func (e PropertyType) IsValid() bool {
	switch e {
	case PropertyTypeApartmentsBuilding, PropertyTypeApartment:
		return true
	}
	return false
}

func (e PropertyType) String() string {
	return string(e)
}

func (e *PropertyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PropertyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PropertyType", str)
	}
	return nil
}

func (e PropertyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ShootStatus string

const (
	ShootStatusPending ShootStatus = "PENDING"
	ShootStatusDone    ShootStatus = "DONE"
)

var AllShootStatus = []ShootStatus{
	ShootStatusPending,
	ShootStatusDone,
}

func (e ShootStatus) IsValid() bool {
	switch e {
	case ShootStatusPending, ShootStatusDone:
		return true
	}
	return false
}

func (e ShootStatus) String() string {
	return string(e)
}

func (e *ShootStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ShootStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ShootStatus", str)
	}
	return nil
}

func (e ShootStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UnitState string

const (
	UnitStateVacant      UnitState = "VACANT"
	UnitStateOccupied    UnitState = "OCCUPIED"
	UnitStateUnavailable UnitState = "UNAVAILABLE"
)

var AllUnitState = []UnitState{
	UnitStateVacant,
	UnitStateOccupied,
	UnitStateUnavailable,
}

func (e UnitState) IsValid() bool {
	switch e {
	case UnitStateVacant, UnitStateOccupied, UnitStateUnavailable:
		return true
	}
	return false
}

func (e UnitState) String() string {
	return string(e)
}

func (e *UnitState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UnitState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UnitState", str)
	}
	return nil
}

func (e UnitState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UploadCategory string

const (
	UploadCategoryProfileImg        UploadCategory = "PROFILE_IMG"
	UploadCategoryUnitImages        UploadCategory = "UNIT_IMAGES"
	UploadCategoryPropertyThumbnail UploadCategory = "PROPERTY_THUMBNAIL"
)

var AllUploadCategory = []UploadCategory{
	UploadCategoryProfileImg,
	UploadCategoryUnitImages,
	UploadCategoryPropertyThumbnail,
}

func (e UploadCategory) IsValid() bool {
	switch e {
	case UploadCategoryProfileImg, UploadCategoryUnitImages, UploadCategoryPropertyThumbnail:
		return true
	}
	return false
}

func (e UploadCategory) String() string {
	return string(e)
}

func (e *UploadCategory) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UploadCategory(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UploadCategory", str)
	}
	return nil
}

func (e UploadCategory) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
