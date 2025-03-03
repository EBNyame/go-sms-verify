package data

type OTPData struct {
	PhoneNumber string `json: "PhoneNumber, omitempty" validate: "required"`
}

type VerifyData struct {
	User *OTPData `json: "User, omitempty" validate: "required"`
	Code string   `json: "Code, omitempty" validate: "required"`
}
