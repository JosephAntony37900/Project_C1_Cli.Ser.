package config

import (
    "log"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/helpers"
    productRepo "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/repository"
    userRepo "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/repository"
)

func setupRepositories() (*productRepo.ProductRepoMySQL, *userRepo.UserRepoMySQL, func()) {
    db, err := helpers.NewMySQLConnection()
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    return productRepo.NewProductRepoMySQL(db), userRepo.NewCreateUserRepoMySQL(db), func() { db.Close() }
}
