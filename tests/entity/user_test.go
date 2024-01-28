package user_tests

import (
	"gochi/internal/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
	u := entity.NewUser("Ricardo", "rick@rick.com", "123456")

	assert.Equal(t, u.DisplayName, "Ricardo")
	assert.Equal(t, u.DisplayName, "rick@rick.com")
	assert.Equal(t, bcrypt.CompareHashAndPassword([]byte(u.Password), []byte("123456")), nil)
}
