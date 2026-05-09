package auth

import (
	"context"
	"errors"

	"github.com/mmtaee/ocserv-dashboard/common/models"
	"github.com/mmtaee/ocserv-dashboard/telegram_bot/internal/repository"
)

var (
	ErrUserNotFound  = errors.New("ocserv user not found")
	ErrInvalidCreds  = errors.New("invalid credentials")
	ErrUserLocked    = errors.New("user is locked")
	ErrUserInactive  = errors.New("user is deactivated")
)

type Verifier struct {
	repo *repository.Repository
}

func NewVerifier(repo *repository.Repository) *Verifier {
	return &Verifier{repo: repo}
}

// Verify validates the given username/password against the ocserv_users table.
// On success it returns the matching user; the caller is responsible for
// linking it to a Telegram chat.
//
// A locked account (typically one whose monthly quota has been exhausted) is
// intentionally NOT rejected here: we want the user to still be able to link
// it, inspect its usage, and request a renewal through the bot. Only fully
// deactivated accounts (DeactivatedAt set) are blocked.
func (v *Verifier) Verify(ctx context.Context, username, password string) (*models.OcservUser, error) {
	user, err := v.repo.OcservUserByUsername(ctx, username)
	if err != nil {
		return nil, ErrUserNotFound
	}
	if user.Password != password {
		return nil, ErrInvalidCreds
	}
	if user.DeactivatedAt != nil {
		return nil, ErrUserInactive
	}
	return user, nil
}
