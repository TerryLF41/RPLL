package model

import (
	"github.com/dgrijalva/jwt-go"
)

// Claim represents a claim model
type Claim struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	UserType int    `json:"usertype"`
	jwt.StandardClaims
}

// ClaimModelFactory interface defines methods for claim model
type ClaimModelFactory interface {
	CreateClaim(id int, name string, userType int) *Claim
}

// ConcreteClaimModelFactory struct implements ClaimModelFactory interface
type ConcreteClaimModelFactory struct{}

// CreateClaim creates a new claim instance
func (factory *ConcreteClaimModelFactory) CreateClaim(id int, name string, userType int) *Claim {
	return &Claim{
		ID:       id,
		Name:     name,
		UserType: userType,
	}
}

// NewClaimModelFactory creates a new claim model factory
func NewClaimModelFactory() ClaimModelFactory {
	return &ConcreteClaimModelFactory{}
}
