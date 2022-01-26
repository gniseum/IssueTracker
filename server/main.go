package main

import (
    "os"

    "server/routes"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8000"
    }

    router := gin.New()
    router.Use(gin.Logger())
    router.Use(cors.Default())

    // Router endpoints
    // Create
    router.POST("/issue/create", routes.AddIssue)

    // Remove
    router.GET("/taskedUser/:taskedUser", routes.GetIssuesByTaskedUser)
    router.GET("/issues", routes.GetIssues)
    router.GET("/issue/:id/", routes.GetIssueById)

    // Update
    router.PUT("/taskedUser/update/:id", routes.UpdateTaskedUser)
    router.PUT("/issue/update/:id", routes.UpdateIssue)

    // Delete
    router.DELETE("/issue/delete/:id", routes.DeleteIssue)

    router.Run(":" + port)
}
