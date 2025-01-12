package dto

// SignUpReq represents the request to the signup endpoint
type SignUpReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phoneNo"`
}

// SignUpReqHeader represents the header for the SignUpReq
type SignUpReqHeader struct {
	AcceptLocale string `json:"accept-locale"`
}

// SignUpRes represents the response from the signup endpoint
type SignUpRes struct {
	Token string `json:"token"`
}
