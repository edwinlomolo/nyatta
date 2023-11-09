// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type AnyUpload struct {
	ID         string     `json:"id"`
	Upload     string     `json:"upload"`
	Category   string     `json:"category"`
	PropertyID string     `json:"propertyId"`
	UserID     string     `json:"userId"`
	UnitID     string     `json:"unitId"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type Bedroom struct {
	ID             string     `json:"id"`
	PropertyUnitID string     `json:"propertyUnitId"`
	BedroomNumber  int        `json:"bedroomNumber"`
	EnSuite        bool       `json:"enSuite"`
	Master         bool       `json:"master"`
	CreatedAt      *time.Time `json:"createdAt"`
	UpdatedAt      *time.Time `json:"updatedAt"`
}

type Caretaker struct {
	ID         string       `json:"id"`
	FirstName  string       `json:"first_name"`
	LastName   string       `json:"last_name"`
	Phone      string       `json:"phone"`
	Uploads    []*AnyUpload `json:"uploads"`
	Verified   bool         `json:"verified"`
	Properties []*Property  `json:"properties"`
	CreatedAt  *time.Time   `json:"createdAt"`
	UpdatedAt  *time.Time   `json:"updatedAt"`
}

type CaretakerInput struct {
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     string         `json:"phone"`
	Image     graphql.Upload `json:"image"`
}

type CaretakerVerificationInput struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verifyCode"`
}

type CreatePaymentInput struct {
	Phone       string `json:"phone"`
	Description string `json:"description"`
	Amount      string `json:"amount"`
	Email       string `json:"email"`
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
	ID            string        `json:"id"`
	Msid          string        `json:"msid"`
	Phone         string        `json:"phone"`
	Status        InvoiceStatus `json:"status"`
	WCoCheckoutID string        `json:"wCoCheckoutId"`
	CreatedAt     *time.Time    `json:"createdAt"`
	UpdatedAt     *time.Time    `json:"updatedAt"`
}

type ListingOverview struct {
	OccupiedUnits int `json:"occupiedUnits"`
	VacantUnits   int `json:"vacantUnits"`
	TotalUnits    int `json:"totalUnits"`
}

type NewUser struct {
	Phone string `json:"phone"`
}

type PropertyUnit struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Bedrooms     []*Bedroom   `json:"bedrooms"`
	PropertyID   string       `json:"propertyId"`
	Property     *Property    `json:"property"`
	Price        string       `json:"price"`
	AmenityCount int          `json:"amenityCount"`
	Bathrooms    int          `json:"bathrooms"`
	Amenities    []*Amenity   `json:"amenities"`
	State        UnitState    `json:"state"`
	Type         string       `json:"type"`
	Uploads      []*AnyUpload `json:"uploads"`
	Tenancy      []*Tenant    `json:"tenancy"`
	CreatedAt    *time.Time   `json:"createdAt"`
	UpdatedAt    *time.Time   `json:"updatedAt"`
}

type PropertyUnitInput struct {
	PropertyID string              `json:"propertyId"`
	Baths      int                 `json:"baths"`
	Name       string              `json:"name"`
	Location   *GpsInput           `json:"location"`
	Type       string              `json:"type"`
	Amenities  []*UnitAmenityInput `json:"amenities"`
	Bedrooms   []*UnitBedroomInput `json:"bedrooms"`
	Price      string              `json:"price"`
	Uploads    []*UploadImages     `json:"uploads"`
}

type Shoot struct {
	ID         string     `json:"id"`
	PropertyID string     `json:"propertyId"`
	Date       time.Time  `json:"date"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
}

type ShootInput struct {
	Date time.Time `json:"date"`
}

type SignInResponse struct {
	User  *User  `json:"user"`
	Token string `json:"Token"`
}

type Status struct {
	Success string `json:"success"`
}

type TenancyInput struct {
	StartDate      time.Time `json:"startDate"`
	PropertyUnitID string    `json:"propertyUnitId"`
}

type Tenant struct {
	ID             string       `json:"id"`
	StartDate      time.Time    `json:"startDate"`
	EndDate        *time.Time   `json:"endDate"`
	Upload         []*AnyUpload `json:"upload"`
	PropertyUnitID string       `json:"propertyUnitId"`
	CreatedAt      *time.Time   `json:"createdAt"`
	UpdatedAt      *time.Time   `json:"updatedAt"`
}

type Token struct {
	Token string `json:"token"`
}

type Town struct {
	ID         string `json:"id"`
	Town       string `json:"town"`
	PostalCode string `json:"postalCode"`
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

type UploadImages struct {
	Image    string `json:"image"`
	Category string `json:"category"`
}

type UserVerificationInput struct {
	Phone      string `json:"phone"`
	VerifyCode string `json:"verifyCode"`
}

type VerificationInput struct {
	Phone      string  `json:"phone"`
	VerifyCode *string `json:"verifyCode"`
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
	UploadCategoryProfileImg   UploadCategory = "PROFILE_IMG"
	UploadCategoryUnitImage    UploadCategory = "UNIT_IMAGE"
	UploadCategoryCaretakerImg UploadCategory = "CARETAKER_IMG"
)

var AllUploadCategory = []UploadCategory{
	UploadCategoryProfileImg,
	UploadCategoryUnitImage,
	UploadCategoryCaretakerImg,
}

func (e UploadCategory) IsValid() bool {
	switch e {
	case UploadCategoryProfileImg, UploadCategoryUnitImage, UploadCategoryCaretakerImg:
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
