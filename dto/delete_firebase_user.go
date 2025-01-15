package dto

// DeleteFirebaseUserReq represents the request to delete a user in Firebase Auth
type DeleteFirebaseUserReq struct {
	UID string `json:"uid"`
}

// DeleteFirebaseUserRes represents the response from deleting a user in Firebase Auth
type DeleteFirebaseUserRes struct {
	Success bool `json:"success"`
}
