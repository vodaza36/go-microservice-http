package domain

// User entity
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Customer entity
type Customer struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

// Address entity
type Address struct {
	Street      string `json:"street"`
	HouseNo     string `json:"houseno"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	Additional1 string `json:"add1"`
	Additional2 string `json:"add2"`
}

// Registration entity
type Registration struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

// UserService for user management
type UserService interface {
	CreateUser(u *User) (*User, error)
	GetUserByEmail(email string) (*User, error)
}

// CustomerService for customer management
type CustomerService interface {
	CreateCustomer(u *Customer) (*Customer, error)
	GetCustomerByID(ID string) (*Customer, error)
}

// RegistrationService for signing up a user
type RegistrationService interface {
	Register(r *Registration) error
}
