package types

import "time"

type Category struct {
	CategoryID   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Description  string `json:"description"`
	ImageURL     string `json:"image_url"`
}

type MenuItem struct {
	ProductID      int     `json:"product_id"`
	ProductName    string  `json:"product_name"`
	CategoryID     *int    `json:"category_id"`
	Price          float64 `json:"price"`
	ImageURL       string  `json:"image_url"`
	IngredientList string  `json:"ingredient_list"`
}

type Order struct {
	OrderID      int       `json:"order_id"`
	UserID       *int      `json:"user_id"`
	Status       string    `json:"status"` // "Yet to start", "Cooking", "Completed"
	CreatedAt    time.Time `json:"created_at"`
	Instructions string    `json:"instructions"`
	TableNo      *int      `json:"table_no"`
}

type CreateOrder struct {
	OrderID      int       `json:"order_id"`
	TableNo      *int       `json:"table_no"`
	Tip          *int       `json:"tip"`
	Instructions string     `json:"instructions"`
}

type CartItem struct {
	OrderID   int `json:"order_id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type Payment struct {
	TransactionID int       `json:"transaction_id"`
	OrderID       *int      `json:"order_id"`
	UserID        *int      `json:"user_id"`
	FoodTotal     float64   `json:"food_total"`
	CreatedAt     time.Time `json:"created_at"`
	Tip           *int      `json:"tip"`
	Status        string    `json:"status"` // "Pending" or "Completed"
}

type MakePayment struct {
	OrderID       *int      `json:"order_id"`
	UserID        *int      `json:"user_id"`
	FoodTotal     float64   `json:"food_total"`
	Tip           *int      `json:"tip"`
}

type User struct {
	UserID       int    `json:"user_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Contact      string `json:"contact"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"` // "administrator", "customer", or "chef"
}

type RegisterUser struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Contact      string `json:"contact"`
	Email        string `json:"email"`
	Password string     `json:"password"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	CreateNewUser(user User) error
	GetUserById(id int) (*User, error)
	GetAllUsers() ([]User, error)
}

type PaymentStore interface {
	PaymentsByUserId(id int) ([]Payment, error)
	GetAllPayments() ([]Payment, error)
}

type OrderStore interface {
	GetAllOrders() ([]Order, error)
	OrdersByStatus(status string) ([]Order, error)
	OrdersByUserId(id int) ([]Order, error)
	CreateOrder(order CreateOrder) error
	GetOneOrder(id int) ([]CartItem, error)
	CreateEmptyOrder(userId int) (int,error)
}

type MenuStore interface {
	ListOfCategory() ([]Category ,error)
	GetMenuByCategoryId(id int) ([]MenuItem ,error)
}

type CartStore interface {
    AddToCart(place CartItem) error
    GetCartItems(orderID int) ([]CartItem, error)
    UpdateCartItemQuantity(place CartItem) error
    RemoveFromCart(orderID, productID int) error
}

