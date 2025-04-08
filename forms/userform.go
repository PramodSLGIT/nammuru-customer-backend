package forms

import "time"

const (
	AccountSid         = "ACbcef2267c53f7781e4259253d31adf9f"
	AuthToken          = "bb0982c09f76a5360e3dc1bd8be62d99"
	FromNumber         = "+1 567 339 2523"
	DB                 = "customer"
	CustomerCollection = "user_info"
	RideCollection     = "ride_info"
)

type UserRegistrationForm struct {
	ProfileImage string `json:"profile_image"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phone_number"`
	Gender       string `json:"gender"`
	Password     string `json:"password"`
}

type UserModel struct {
	ProfileImage []byte `json:"profile_image" bson:"profile_image" binding:"-"`
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	PhoneNumber  string `json:"phone_number" bson:"phone_number"`
	Gender       string `json:"gender" bson:"gender"`
	Password     string `json:"password"`
}

type LoginUserModel struct {
	PhoneNumber string     `json:"phone_number" bson:"phone_number"`
	Otp         string     `json:"otp" bson:"otp"`
	Time        *time.Time `json:"time" bson:"time"`
}

type LoginModel struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
