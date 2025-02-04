package config

import (
    productController "github.com/JosephAntony37900/ArquitecturaHexagonal/products/infrastructure/controllers"
    userController "github.com/JosephAntony37900/ArquitecturaHexagonal/users/infrastructure/controllers"
    productApp "github.com/JosephAntony37900/ArquitecturaHexagonal/products/application"
    userApp "github.com/JosephAntony37900/ArquitecturaHexagonal/users/application"
)

func setupControllers(
    createProduct *productApp.CreateProduct,
    getProducts *productApp.GetProducts,
    updateProduct *productApp.UpdateProduct,
    deleteProduct *productApp.DeleteProduct,
    createUser *userApp.CreateUsers,
    getUsers *userApp.GetUsers,
    deleteUser *userApp.DeleteUser,
    updateUser *userApp.UpdateUser,
) (
    *productController.CreateProductController,
    *productController.GetProductsController,
    *productController.UpdateProductController,
    *productController.DeleteProductController,
    *userController.CreateUserController,
    *userController.GetUsersController,
    *userController.DeleteUserController,
    *userController.UpdateUserController,
) {
    return productController.NewCreateProductController(createProduct),
        productController.NewGetProductsController(getProducts),
        productController.NewUpdateProductController(updateProduct),
        productController.NewDeleteProductController(deleteProduct),
        userController.NewCreateUserController(createUser),
        userController.NewUsersController(getUsers),
        userController.NewDeleteUserController(deleteUser),
        userController.NewUpdateUserController(updateUser)
}
