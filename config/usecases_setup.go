package config

import (
    productApp "github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
    userApp "github.com/JosephAntony37900/ArquitecturaHexagonal/users/application"
    productRepo "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/repository"
    userRepo "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/repository"
)

func setupUseCases(productRepo *productRepo.ProductRepoMySQL, userRepo *userRepo.UserRepoMySQL) (
    *productApp.CreateProduct,
    *productApp.GetProducts,
    *productApp.UpdateProduct,
    *productApp.DeleteProduct,
    *userApp.CreateUsers,
    *userApp.GetUsers,
    *userApp.DeleteUser,
    *userApp.UpdateUser,
) {
    return productApp.NewCreateProduct(productRepo),
        productApp.NewGetProducts(productRepo),
        productApp.NewUpdateProduct(productRepo),
        productApp.NewDeleteProduct(productRepo),
        userApp.NewCreateUser(userRepo),
        userApp.NewGetUsers(userRepo),
        userApp.NewDeleteUser(userRepo),
        userApp.NewUpdateUser(userRepo)
}
