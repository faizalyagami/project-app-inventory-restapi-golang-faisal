package model

type Session struct {
	Token  string `json:"token"`
	UserID int64  `json:"user_id"`
}
