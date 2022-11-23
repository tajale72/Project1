package controller

import "log"

type VerifyToken struct {
	Token string `json:"token"`
}

type JwtVerifyToken struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	ExpiresAt string `json:"expiresAt"`
	IssuesAt  string `json:"issuesAt"`
}

func (s *Service) VerifyToken(token string) (string, error) {
	log.Println("VerifyToken")
	return "VerifyToken", nil
}

func (s *Service) GenerateToken(body []byte) (string, error) {
	log.Println("GenerateToken")
	return "Generatetoken", nil
}
