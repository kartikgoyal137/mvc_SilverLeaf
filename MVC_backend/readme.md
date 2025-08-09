# Velvet Plate - Food Ordering System

Velvet Plate is a backend for a food ordering system built in Go, using the MVC pattern. It features a robust API for managing users, menus, orders, and payments, with role-based access control for customers, chefs, and administrators.

### Installation

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/kartikgoyal137/mvc-food_ordering_system.git](https://github.com/kartikgoyal137/mvc-food_ordering_system.git)
    cd mvc-food_ordering_system/MVC_backend
    ```

2.  **Install dependencies:**
    ```bash
    go mod download
    ```

3.  **Set up the database:**
    * Create a MySQL database named `velvet_plate`.
    * Configure your database credentials in a `.env` file.
    * Run the SQL migrations from the `database/migrations` directory.

4.  **Run the server:**
    ```bash
    go run cmd/main.go
    ```

## Project Structure

The project follows a standard Go project structure, separating concerns into `cmd` for the application entry point, and `pkg` for decoupled packages handling the API, controllers, middleware, models, and types.