package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := RandomString(6)
	hp1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hp1)

	err = CheckPassword(password, hp1)
	require.NoError(t, err)

	wrongPassword := RandomString(6)

	err = CheckPassword(wrongPassword, hp1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hp2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hp2)

	require.NotEqual(t, hp1, hp2)

}
