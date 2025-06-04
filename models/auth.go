package models

type (
	AuthLogin struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	AuthToken struct {
		Token string `json:"token"`
	}
)
