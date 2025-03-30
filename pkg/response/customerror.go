package response

const (
	ERROR_FIELD_ENTITY string = "please insert all field"
	// Unique record that already exists
	ERROR_ALREADY_EXISTS string = "already_exists"
	// Invalid credentials
	ERROR_INVALID_CREDENTIAL string = "invalid_credential"
	// Invalid request
	ERROR_INVALID_REQUEST string = "invalid_request"
	// Record not found
	ERROR_NOT_FOUND string = "not_found"
	//wrong  password
	ERROR_WRONG_PASSWORD string = "wrong_password"

	ERROR_NOT_VERIFIED string = "not_verified"

	ERROR_INVALID_OTP string = "invalid_otp"

	ERROR_EXPIRED_OTP string = "expired_otp"

	ERROR_OTP_ALREADY_USED string = "otp_already_used"

	ERROR_ALREADY_VERIFIED string = "already_verified"
)

type Err struct {
	Code   string `json:"code"`
	Errors any    `json:"error"`
}

func (e *Err) Error() string {
	return e.Code
}
