package util

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// usiamo bcrypt per generare l'hash della password prima di memorizzarla nel db
func HashPassword(password string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hp), nil
}

// CheckPassword serve per verificare il matching tra password fornita dall'utente e quella memorizzata nel db
func CheckPassword(password string, hashedPassword string) error {
	// usiamo bcrypt per comparare, c'e' gia' il metodo.
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
