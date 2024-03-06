package model

import "github.com/google/uuid"

type LogDetailResponse struct {
	Xid      *uuid.UUID `json:"xid,omitempty"`
	Contents []string   `json:"contents"`
}

type LogRequest struct {
	Xid string `json:"xid" validate:"required"`
}
