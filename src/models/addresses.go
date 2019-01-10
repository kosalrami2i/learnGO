package models

/*Addresses is a derived datatype that gets address data of a user and constructs a JSON object */
type Addresses struct {
	ID       int64  `json:"id"`
	Location string `json:"location"`
}
