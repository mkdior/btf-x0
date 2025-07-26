package user

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	ID      int
	Balance int
}

func Serialize(u User) string {
	return fmt.Sprintf("(%d,%d)", u.ID, u.Balance)
}

func Deserialize(s string) (User, error) {
	s = strings.Trim(s, "()")
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return User{}, errors.New("malformed user data")
	}

	id, err := strconv.Atoi(parts[0])
	if err != nil {
		return User{}, fmt.Errorf("invalid ID: %w", err)
	}

	balance, err := strconv.Atoi(parts[1])
	if err != nil {
		return User{}, fmt.Errorf("invalid Balance: %w", err)
	}

	return User{ID: id, Balance: balance}, nil
}
