package models

/*Users is a derived datatype that gets user data and constructs a JSON object */
type Users struct {
	ID           int64       `json:"id"`
	Name         string      `json:"name"`
	MobileNumber string      `json:"mobile_number"`
	Email        string      `json:"email"`
	Addresses    []Addresses `json:"addresses"`
}
