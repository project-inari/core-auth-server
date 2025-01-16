package dto

// SignUpReq represents the request to the signup endpoint
type SignUpReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	PhoneNo  string `json:"phoneNo" validate:"required"`
}

// SignUpReqHeader represents the header for the SignUpReq
type SignUpReqHeader struct {
	AcceptLocale string
}

// SignUpRes represents the response from the signup endpoint
type SignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Token    string `json:"token"`
}
