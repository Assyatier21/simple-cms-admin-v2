package constant

const (
	OK      = "OK"
	SUCCESS = "success"
	FAILED  = "failed"

	SUCCESS_REGISTER_USER = "phone number succesfully registered"
	SUCCESS_LOGIN         = "successfully login"

	ERR_VALIDATION_ID      = "id must be an integer and can't be empty"
	ERR_PHONE_REGISTERED   = "phone number already registered"
	ERR_REGISTRY_NOT_FOUND = "registry not found"
	ERR_PHONE_OR_PASS      = "phone or password are incorrect"
	ERR_ROLE_FORMAT        = "role chosen aren't defined"
	ERR_GENERATE_JWT       = "failed to generate jwt token"
	ERR_EMPTY_TOKEN        = "token is empty"
	ERR_EMPTY_PAYLOAD      = "payload can't be empty"
	ERR_BINDING            = "failed to bind request"
	ERR_TOKEN_EXPIRED      = "token already expired"

	CHARSET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)
