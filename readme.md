# Silver Leaf - Full-Stack Food Ordering System

Silver Leaf is a complete, full-stack restaurant management and food ordering application. It features a robust backend built with Go, following the MVC pattern, and a modern, responsive frontend built with React.

The system is designed to serve three key roles: **Customers** who can browse the menu and place orders, **Chefs** who manage the kitchen and update order statuses, and **Administrators** who have full oversight of the platform, including user management, menu updates, and payment tracking. The entire application is containerized with Docker for easy setup and deployment.

## Features

-   **Role-Based Access Control:** Distinct interfaces and permissions for Customers, Chefs, and Administrators.
-   **Interactive Menu:** Customers can browse food categories, view items, and add them to their cart.
-   **Order Management:** A complete workflow from placing an order to preparation in the kitchen and final delivery.
-   **Chef's Dashboard:** A dedicated view for chefs to see active orders and update their status ("Yet to start", "Cooking", "Completed").
-   **Admin Panel:** A comprehensive dashboard for administrators to manage users, orders, payments, and menu items.
-   **Authentication:** Secure user login and signup using JWT (JSON Web Tokens).
-   **In-Memory Caching:** Backend caching for frequently requested data (like menus) to ensure fast response times.

## Tech Stack

-   **Backend:** Go, Gorilla/Mux (for routing), MySQL
-   **Frontend:** React, Vite, Bootstrap, Axios
-   **Database:** MySQL
-   **Containerization:** Docker

## Getting Started

Follow these instructions to get the project up and running on your local machine.

### Setup and Installation

1.  **Clone the Repository**
    Open your terminal and clone the project:
    ```bash
    git clone [https://github.com/kartikgoyal137/mvc_silverleaf.git](https://github.com/kartikgoyal137/mvc_silverleaf.git)
    cd mvc_silverleaf
    ```

2.  **Configure Environment Variables**
    The project uses a `.env` file to manage sensitive information like database credentials and JWT secrets. Create a `.env` file in the `mvc_backend` directory by copying the sample file.

    ```bash
    cp mvc_backend/.env.sample mvc_backend/.env
    ```

    Now, open `mvc_backend/.env` and fill in the required values:
    ```env
    DBUSER=your_mysql_user
    DBPASS=your_mysql_password
    TOKENKEY=your_super_secret_jwt_key
    DBNAME=silver_leaf
    DBHOST=host.docker.internal # if using docker-desktop
    ```

3.  **Build and Run with Docker**
    The included `Makefile` simplifies the Docker process. In the root directory of the project, run:
    ```bash
    make build
    make run
    ```
    This command will build the multi-stage Docker image, which compiles the Go backend, builds the React frontend, sets up the database, runs migrations, and starts the application.

    Your application should now be running!

## How to Use

-   **Access the Application:** Open your web browser and navigate to `http://localhost:8080/`.
-   **API Port:** The backend API is also available on port `8080`.

### Default Accounts

The database is pre-seeded with default accounts for the administrator and chef roles.

-   **Administrator:**
    -   **Email:** `admin@example.com`
    -   **Password:** `admin1234`

-   **Chef:**
    -   **Email:** `chef@example.com`
    -   **Password:** `chef1234`

You can also sign up as a new customer through the frontend application.