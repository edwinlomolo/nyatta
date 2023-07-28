// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type UnitState string

const (
	UnitStateVacant      UnitState = "vacant"
	UnitStateUnavailable UnitState = "unavailable"
	UnitStateOccupied    UnitState = "occupied"
)

func (e *UnitState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UnitState(s)
	case string:
		*e = UnitState(s)
	default:
		return fmt.Errorf("unsupported scan type for UnitState: %T", src)
	}
	return nil
}

type NullUnitState struct {
	UnitState UnitState
	Valid     bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullUnitState) Scan(value interface{}) error {
	if value == nil {
		ns.UnitState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.UnitState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullUnitState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.UnitState, nil
}

type Amenity struct {
	ID             int64          `json:"id"`
	Name           string         `json:"name"`
	Provider       sql.NullString `json:"provider"`
	CreatedAt      time.Time      `json:"created_at"`
	Category       string         `json:"category"`
	UpdatedAt      time.Time      `json:"updated_at"`
	PropertyUnitID int64          `json:"property_unit_id"`
}

type Bedroom struct {
	ID             int64     `json:"id"`
	BedroomNumber  int32     `json:"bedroom_number"`
	EnSuite        bool      `json:"en_suite"`
	Master         bool      `json:"master"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	PropertyUnitID int64     `json:"property_unit_id"`
}

type Caretaker struct {
	ID             int64          `json:"id"`
	FirstName      string         `json:"first_name"`
	LastName       string         `json:"last_name"`
	Idverification string         `json:"idverification"`
	CountryCode    string         `json:"country_code"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	Phone          sql.NullString `json:"phone"`
	Image          string         `json:"image"`
	Verified       bool           `json:"verified"`
}

type Mailing struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

type Property struct {
	ID         int64         `json:"id"`
	Name       string        `json:"name"`
	Town       string        `json:"town"`
	PostalCode string        `json:"postal_code"`
	Type       string        `json:"type"`
	Status     string        `json:"status"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	CreatedBy  int64         `json:"created_by"`
	Caretaker  sql.NullInt64 `json:"caretaker"`
}

type PropertyUnit struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	State      UnitState `json:"state"`
	Price      int32     `json:"price"`
	Bathrooms  int32     `json:"bathrooms"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	PropertyID int64     `json:"property_id"`
}

type Shoot struct {
	ID             int64     `json:"id"`
	ShootDate      time.Time `json:"shoot_date"`
	PropertyID     int64     `json:"property_id"`
	PropertyUnitID int64     `json:"property_unit_id"`
	Status         string    `json:"status"`
	CaretakerID    int64     `json:"caretaker_id"`
}

type Tenant struct {
	ID             int64        `json:"id"`
	StartDate      time.Time    `json:"start_date"`
	EndDate        sql.NullTime `json:"end_date"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	PropertyUnitID int64        `json:"property_unit_id"`
}

type User struct {
	ID         int64          `json:"id"`
	Email      sql.NullString `json:"email"`
	FirstName  sql.NullString `json:"first_name"`
	LastName   sql.NullString `json:"last_name"`
	Phone      sql.NullString `json:"phone"`
	Onboarding sql.NullBool   `json:"onboarding"`
	Avatar     sql.NullString `json:"avatar"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}
