package controllers

import (
    "log"
    "github.com/JosephAntony37900/ArquitecturaHexagonal/users/application"
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

type GetUsersController struct {
    getUsers *application.GetUsers
}

func NewUsersController(getUsers *application.GetUsers) *GetUsersController {
    return &GetUsersController{getUsers: getUsers}
}

func (gu *GetUsersController) Handle(ctx *gin.Context) {
    log.Println("Petición de listar todos los usuarios, recibido")

    user, err := gu.getUsers.Run()
    if err != nil {
        log.Printf("Error buscando usuarios")
        ctx.JSON(500, gin.H{"error": err.Error()})
        return
    }

    log.Printf("Retornando %d usuarios", len(user))
    ctx.JSON(200, user)
}

// Controlador para Short Polling
func (gu *GetUsersController) ShortPoll(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
}

// Controlador para Long Polling
func (gu *GetUsersController) LongPoll(ctx *gin.Context) {
    timeout := time.After(30 * time.Second)
    select {
    case <-timeout:
        ctx.JSON(http.StatusOK, gin.H{"message": "No hay datos nuevos"})
    case newData := <-waitForNewData():
        ctx.JSON(http.StatusOK, gin.H{"data": newData})
    }
}
