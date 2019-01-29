package models

/*Users is a derived datatype that gets user data and constructs a JSON object */
type Users struct {
	Id            string `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name          string `json:"name" gorm:"type:string"`
	Mobile_Number string `json:"mobile_number" gorm:"type:string"`
	Email         string `json:"email" gorm:"type:string"`
}
