package routes

import (
    "context"
    "fmt"
    "net/http"
    "time"

    "server/models"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/bson"
    "github.com/gin-gonic/gin"
)

var validate = validator.New()
var issueCollection *mongo.Collection = OpenCollection(Client, "issues")

// Add an issue
func AddIssue(c *gin.Context) {
    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)
    var issue models.Issue

    if err := c.BindJSON(&issue); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    validationErr := validate.Struct(issue)
    if validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        return
    }

    issue.ID = primitive.NewObjectID()

    result, insertErr := issueCollection.InsertOne(ctx, issue)
    if insertErr != nil {
        msg := fmt.Sprintf("issue item was not created")
        c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
        fmt.Println(insertErr)
        return
    }

    defer cancel()

    c.JSON(http.StatusOK, result)
}

// Get all issues
func GetIssues(c *gin.Context) {
    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)
    var issues []bson.M

    cursor, err := issueCollection.Find(ctx, bson.M{})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err = cursor.All(ctx, &issues); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    defer cancel()

    fmt.Println(issues)

    c.JSON(http.StatusOK, issues)
}

// Get issues by tasked user
func GetIssuesByTaskedUser(c *gin.Context) {
    taskedUser := c.Params.ByName("taskedUser")
    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)
    var issues []bson.M

    cursor, err := issueCollection.Find(ctx, bson.M{"taskedUser": taskedUser})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    if err = cursor.All(ctx, &issues); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    defer cancel()

    fmt.Println(issues)

    c.JSON(http.StatusOK, issues)
}

// Get issue by id
func GetIssueById(c *gin.Context) {
    issueID := c.Params.ByName("id")
    docID, _ := primitive.ObjectIDFromHex(issueID)

    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)
    var issue bson.M

    if err := issueCollection.FindOne(ctx,
        bson.M{"_id": docID}).Decode(&issue); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    defer cancel()

    fmt.Println(issue)

    c.JSON(http.StatusOK, issue)
}

// Update the issue's tasked user
func UpdateTaskedUser(c *gin.Context) {
    issueID := c.Params.ByName("id")
    docID, _ := primitive.ObjectIDFromHex(issueID)

    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)

    type TaskedUser struct {
        Name *string `json:"taskedUser"`
    }

    var taskedUser TaskedUser

    if err := c.BindJSON(&taskedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    result, err := issueCollection.UpdateOne(ctx, bson.M{"_id": docID},
        bson.D{
            {"$set", bson.D{{"taskedUser", taskedUser.Name}}},
        },
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError,
               gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    defer cancel()

    c.JSON(http.StatusOK, result.ModifiedCount)
}

// Update issue
func UpdateIssue(c *gin.Context) {
    issueID := c.Params.ByName("id")
    docID, _ := primitive.ObjectIDFromHex(issueID)

    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)

    var issue models.Issue

    if err := c.BindJSON(&issue); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    validationErr := validate.Struct(issue)
    if validationErr != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
        fmt.Println(validationErr)
        return
    }

    result, err := issueCollection.ReplaceOne(
        ctx,
        bson.M{"_id": docID},
        bson.M{
            "taskedUser": issue.TaskedUser,
            "issueLevel": issue.IssueLevel,
            "state": issue.State,
            "startDate": issue.StartDate,
            "finishDate": issue.FinishDate,
        },
    )

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    defer cancel()

    c.JSON(http.StatusOK, result.ModifiedCount)
}

// Delete an issue
func DeleteIssue(c *gin.Context) {
    issueID := c.Params.ByName("id")
    docID, _ := primitive.ObjectIDFromHex(issueID)

    var ctx, cancel = context.WithTimeout(context.Background(),
                                          100 * time.Second)

    result, err := issueCollection.DeleteOne(ctx, bson.M{"_id": docID})

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        fmt.Println(err)
        return
    }

    defer cancel()

    c.JSON(http.StatusOK, result.DeletedCount)
}

