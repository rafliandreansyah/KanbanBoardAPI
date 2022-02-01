package models

import (
	"KanbanBoardAPI/helpers"
	"database/sql/driver"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Role string

const(
	Admin = "admin"
	Member = "member"
)

func (r *Role) Scan(value interface{}) error  {
	*r = Role(value.([]byte))
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

type User struct {
	ID        uuid.UUID `gorm:"type=uuid;primaryKey" json:"id" form:"id"`
	FullName  string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~full name is empty"`
	Email     string    `gorm:"email;unique;not null" json:"email" form:"email" valid:"required~email is empty,email~invalid format email"`
	Password  string    `gorm:"not null" json:"password" form:"password" valid:"required~password is empty,minstringlength(6)~minimum length password is 6 character"`
	Role      string    `gorm:"not null" sql:"type:user_role" json:"role" form:"role" valid:"required~role is empty,in(admin|member)~role not found"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.NewV4()

	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	hashPass, err := helpers.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashPass
	return nil
}
