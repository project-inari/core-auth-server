package dto

const (
	adaptorFirebaseAuthContentType = "application/json"
)

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
	m["Accept-Locale"] = h.AcceptLocale
	if h.ContentType == "" {
		m["Content-Type"] = adaptorFirebaseAuthContentType
	} else {
		m["Content-Type"] = h.ContentType
	}
	return m
}

// AdaptorFirebaseAuthSignUpRes represents the response from the adaptor-firebase-auth signup endpoint
type AdaptorFirebaseAuthSignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Token    string `json:"token"`
}

// AdaptorFirebaseAuthVerifyTokenReq represents the request to the adaptor-firebase-auth verify token endpoint
type AdaptorFirebaseAuthVerifyTokenReq struct {
	Token string `json:"token"`
}

// AdaptorFirebaseAuthVerifyTokenReqHeader represents the header for the adaptorFirebaseAuthVerifyTokenReq
type AdaptorFirebaseAuthVerifyTokenReqHeader struct {
	ContentType string
}

// ToMap converts the header to a map
func (h AdaptorFirebaseAuthVerifyTokenReqHeader) ToMap() map[string]string {
	m := make(map[string]string)
	if h.ContentType == "" {
		m["Content-Type"] = adaptorFirebaseAuthContentType
	} else {
		m["Content-Type"] = h.ContentType
	}
	return m
}

// AdaptorFirebaseAuthVerifyTokenRes represents the response from the adaptor-firebase-auth verify token endpoint
type AdaptorFirebaseAuthVerifyTokenRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Success  bool   `json:"success"`
}

// AdaptorFirebaseAuthDeleteUserReq represents the request to the adaptor-firebase-auth delete user endpoint
type AdaptorFirebaseAuthDeleteUserReq struct {
	UID string `json:"uid"`
}

// AdaptorFirebaseAuthDeleteUserReqHeader represents the header for the adaptorFirebaseAuthDeleteUserReq
type AdaptorFirebaseAuthDeleteUserReqHeader struct {
	ContentType string
}

// ToMap converts the header to a map
func (h AdaptorFirebaseAuthDeleteUserReqHeader) ToMap() map[string]string {
	m := make(map[string]string)
	if h.ContentType == "" {
		m["Content-Type"] = adaptorFirebaseAuthContentType
	} else {
		m["Content-Type"] = h.ContentType
	}
	return m
}

// AdaptorFirebaseAuthDeleteUserRes represents the response from the adaptor-firebase-auth delete user endpoint
type AdaptorFirebaseAuthDeleteUserRes struct {
	Success bool `json:"success"`
}
