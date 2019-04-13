package users

// User struct for registering users
type User struct {
	ID        int    `json:"id" form:"id" query:"id"`
	FirstName string `json:"firstname" form:"firstname" query:"firstname"`
	LastName  string `json:"lastname" form:"lastname" query:"lastname"`
	Email     string `json:"email" form:"email" query:"email"`
	Pass      string `json:"password" form:"password" query:"password"` //encrypt
}
