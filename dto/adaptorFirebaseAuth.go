package dto

// AdaptorFirebaseAuthSignUpReq represents the request to the adaptor-firebase-auth signup endpoint
type AdaptorFirebaseAuthSignUpReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	PhoneNo  string `json:"phoneNo"`
}

// AdaptorFirebaseAuthSignUpReqHeader represents the header for the adaptorFirebaseAuthSignUpReq
type AdaptorFirebaseAuthSignUpReqHeader struct {
	ContentType  string
	AcceptLocale string
}

// ToMap converts the header to a map
func (h AdaptorFirebaseAuthSignUpReqHeader) ToMap() map[string]string {
	m := make(map[string]string)
	m["accept-locale"] = h.AcceptLocale
	m["Content-Type"] = h.ContentType
	return m
}

// AdaptorFirebaseAuthSignUpRes represents the response from the adaptor-firebase-auth signup endpoint
type AdaptorFirebaseAuthSignUpRes struct {
	Token string `json:"token"`
}
