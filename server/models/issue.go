package models

import (
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Issue struct {
    ID              primitive.ObjectID      `bson:"_id"`
    TaskedUser      *string                 `json:"taskedUser"`
    IssueLevel      *int                    `json:"issueLevel"`
    State           *string                 `json:"state"`
    StartDate       *string                 `json:"startDate"`
    FinishDate      *string                 `json:"finishDate"`
}
