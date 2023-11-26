// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package store

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Amenity struct {
	ID        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Provider  sql.NullString `json:"provider"`
	Category  string         `json:"category"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	UnitID    uuid.UUID      `json:"unit_id"`
}

type Bedroom struct {
	ID            uuid.UUID `json:"id"`
	BedroomNumber int32     `json:"bedroom_number"`
	EnSuite       bool      `json:"en_suite"`
	Master        bool      `json:"master"`
	UnitID        uuid.UUID `json:"unit_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Caretaker struct {
	ID        uuid.UUID     `json:"id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Phone     string        `json:"phone"`
	Verified  bool          `json:"verified"`
	CreatedBy uuid.NullUUID `json:"created_by"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type Invoice struct {
	ID          uuid.UUID      `json:"id"`
	Msid        sql.NullString `json:"msid"`
	Channel     sql.NullString `json:"channel"`
	Currency    sql.NullString `json:"currency"`
	Bank        sql.NullString `json:"bank"`
	AuthCode    sql.NullString `json:"auth_code"`
	CountryCode sql.NullString `json:"country_code"`
	Fees        sql.NullString `json:"fees"`
	Amount      sql.NullString `json:"amount"`
	Phone       sql.NullString `json:"phone"`
	Status      string         `json:"status"`
	Reference   sql.NullString `json:"reference"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

type Mailing struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Property struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Location    interface{}   `json:"location"`
	Type        string        `json:"type"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	CreatedBy   uuid.NullUUID `json:"created_by"`
	CaretakerID uuid.NullUUID `json:"caretaker_id"`
}

type Shoot struct {
	ID          uuid.UUID `json:"id"`
	ShootDate   time.Time `json:"shoot_date"`
	PropertyID  uuid.UUID `json:"property_id"`
	UnitID      uuid.UUID `json:"unit_id"`
	Status      string    `json:"status"`
	CaretakerID uuid.UUID `json:"caretaker_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Tenant struct {
	ID        uuid.UUID    `json:"id"`
	StartDate time.Time    `json:"start_date"`
	EndDate   sql.NullTime `json:"end_date"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	UnitID    uuid.UUID    `json:"unit_id"`
	UserID    uuid.UUID    `json:"user_id"`
}

type Unit struct {
	ID          uuid.UUID     `json:"id"`
	Name        string        `json:"name"`
	Location    interface{}   `json:"location"`
	Type        string        `json:"type"`
	State       string        `json:"state"`
	Price       int32         `json:"price"`
	Bathrooms   int32         `json:"bathrooms"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	CreatedBy   uuid.NullUUID `json:"created_by"`
	CaretakerID uuid.NullUUID `json:"caretaker_id"`
	PropertyID  uuid.NullUUID `json:"property_id"`
}

type Upload struct {
	ID          uuid.UUID      `json:"id"`
	Upload      string         `json:"upload"`
	Category    string         `json:"category"`
	Label       sql.NullString `json:"label"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	UnitID      uuid.NullUUID  `json:"unit_id"`
	PropertyID  uuid.NullUUID  `json:"property_id"`
	UserID      uuid.NullUUID  `json:"user_id"`
	CaretakerID uuid.NullUUID  `json:"caretaker_id"`
	TenantID    uuid.NullUUID  `json:"tenant_id"`
}

type User struct {
	ID               uuid.UUID      `json:"id"`
	FirstName        sql.NullString `json:"first_name"`
	LastName         sql.NullString `json:"last_name"`
	SubscribeRetries int32          `json:"subscribe_retries"`
	NextRenewal      int64          `json:"next_renewal"`
	Phone            string         `json:"phone"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
}
