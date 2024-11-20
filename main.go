package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Define a model
type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Age   int
}

// Initialize database connection
func initDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Create a new user
func createUser(db *gorm.DB, user User) error {
	if err := db.Create(&user).Error; err != nil {
		return err
	}
	fmt.Printf("New user created: %+v\n", user)
	return nil
}

// Read user by email
func readUserByEmail(db *gorm.DB, email string) (*User, error) {
	var user User
	if err := db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	fmt.Printf("User retrieved: %+v\n", user)
	return &user, nil
}

// Update user's age
func updateUserAge(db *gorm.DB, user *User, newAge int) error {
	if err := db.Model(user).Update("Age", newAge).Error; err != nil {
		return err
	}
	fmt.Printf("User updated: %+v\n", *user)
	return nil
}

// Delete user
func deleteUser(db *gorm.DB, user *User) error {
	if err := db.Delete(user).Error; err != nil {
		return err
	}
	fmt.Println("User deleted successfully")
	return nil
}

func main() {
	// Initialize DB connection
	db, err := initDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to PostgreSQL using GORM!")

	// AutoMigrate the schema
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}

	// Perform CRUD operations
	newUser := User{Name: "John Doe", Email: "johndoe@example.com", Age: 30}

	// Create
	if err := createUser(db, newUser); err != nil {
		panic(err)
	}

	// Read
	user, err := readUserByEmail(db, "johndoe@example.com")
	if err != nil {
		panic(err)
	}

	// Update
	if err := updateUserAge(db, user, 31); err != nil {
		panic(err)
	}

	// Delete
	if err := deleteUser(db, user); err != nil {
		panic(err)
	}
}
