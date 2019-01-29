package models

/*Addresses is a derived datatype that gets address data of a user and constructs a JSON object */
type Addresses struct {
	Id       int64
	Location string
	User     *Users
}
