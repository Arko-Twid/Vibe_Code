package main

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func setupRoutes(app *fiber.App) {
	app.Post("/signup", signupHandler)
	app.Post("/login", loginHandler)

	app.Get("/products", getProductsHandler)
	app.Get("/products/:id", getProductHandler)
	app.Post("/products", createProductHandler)
	app.Put("/products/:id", updateProductHandler)
	app.Delete("/products/:id", deleteProductHandler)
}

func signupHandler(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	user := User{ID: uuid.New().String(), Email: body.Email, Role: "customer"}
	if err := user.SetPassword(body.Password); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	if err := DB.Create(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Email already exists"})
	}
	return c.JSON(fiber.Map{"message": "Signup successful"})
}

func loginHandler(c *fiber.Ctx) error {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var body req
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	var user User
	if err := DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	if !user.CheckPassword(body.Password) {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}
	return c.JSON(fiber.Map{"token": tokenString})
}

func getProductsHandler(c *fiber.Ctx) error {
	var products []Product
	if err := DB.Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to fetch products"})
	}
	return c.JSON(products)
}

func getProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := DB.Where("id = ?", id).First(&product).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

func createProductHandler(c *fiber.Ctx) error {
	var product Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	product.ID = uuid.New().String()
	product.CreatedAt = time.Now()
	if err := DB.Create(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to create product"})
	}
	return c.JSON(product)
}

func updateProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var product Product
	if err := DB.Where("id = ?", id).First(&product).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Product not found"})
	}
	var update Product
	if err := c.BodyParser(&update); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	product.Name = update.Name
	product.Description = update.Description
	product.Price = update.Price
	product.ImageURL = update.ImageURL
	product.Category = update.Category
	product.Tags = update.Tags
	product.Stock = update.Stock
	if err := DB.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update product"})
	}
	return c.JSON(product)
}

func deleteProductHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := DB.Delete(&Product{}, "id = ?", id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete product"})
	}
	return c.JSON(fiber.Map{"message": "Product deleted"})
}
