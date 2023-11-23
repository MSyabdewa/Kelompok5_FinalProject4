package handler

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/infra/config"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/infra/database"
	middlewares "github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/middleware"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/repository/categoryRepository/category_pg"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/repository/productRepository/product_pg"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/repository/transactionHistoryRepository/transaction_history_pg"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/repository/userRepository/user_pg"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/service"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	config.LoadAppConfig()
	database.InitiliazeDatabase()
	db := database.GetDatabaseInstance()
	database.SeedAdmin(db)

	userRepo := user_pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	categoryRepo := category_pg.NewCategoryPG(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	productRepo := product_pg.NewProductPG(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	transactionRepo := transaction_history_pg.NewTransactionPG(db)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionHandler := NewTransactionHandler(transactionService)

	r := gin.Default()
	users := r.Group("/users")
	{
		users.POST("/register", userHandler.CreateUser)
		users.POST("/login", userHandler.Login)
		users.Use(middlewares.Authentication())
		{
			users.PATCH("/topup", userHandler.Topup)
		}
	}

	categories := r.Group("/categories")
	{
		categories.Use(middlewares.Authentication())
		{
			categories.Use(middlewares.AdminAuthorization())
			categories.GET("/", categoryHandler.GetCategories)
			categories.POST("/", categoryHandler.CreateCategory)
			categories.PATCH("/:categoryId", categoryHandler.UpdateCategory)
			categories.DELETE("/:categoryId", categoryHandler.DeleteCategory)
		}
	}

	products := r.Group("/products")
	{
		products.Use(middlewares.Authentication())
		{
			products.GET("/", productHandler.GetProduct)
			products.Use(middlewares.AdminAuthorization())
			{
				products.POST("/", productHandler.CreateProduct)
				products.PUT("/:productId", productHandler.UpdateProduct)
				products.DELETE("/:productId", productHandler.DeleteProduct)
			}
		}
	}

	transactions := r.Group("/transactions")
	{
		transactions.Use(middlewares.Authentication())
		{
			transactions.POST("/", transactionHandler.CreateTransaction)
			transactions.GET("/my-transactions", transactionHandler.GetTransactionUser)
			transactions.Use(middlewares.AdminAuthorization())
			transactions.GET("/user-transactions", transactionHandler.GetTransactionAdmin)
		}
	}

	r.Run(":" + config.GetAppConfig().Port)
}
