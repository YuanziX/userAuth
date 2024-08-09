package models

import "github.com/google/uuid"

type AuthDetails struct {
	UserEmail string
	AuthUUID  uuid.UUID
}
