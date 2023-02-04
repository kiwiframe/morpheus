package postgres

import (
	"time"

	"github.com/kiwiframe/morpheus/backend/internal/adapters/postgres/sqlc"
	"github.com/kiwiframe/morpheus/backend/internal/core/domain"
)

func accountToDomain(a *sqlc.UsersAccount) (*domain.Account, error) {
	lastLoggedIn := time.Time{}
	if a.LastLoggedIn.Valid {
		lastLoggedIn = a.LastLoggedIn.Time
	}
	restrictionExpiry := time.Time{}
	if a.RestrictionExpiry.Valid {
		restrictionExpiry = a.RestrictionExpiry.Time
	}

	params := &domain.NewAccountParams{
		ID:                a.ID,
		Username:          a.Username,
		Email:             a.Email,
		HashedPassword:    a.HashedPassword,
		CreatedAt:         a.CreatedAt,
		UpdatedAt:         a.UpdatedAt,
		LastLoggedIn:      lastLoggedIn,
		PrivilegeLevel:    domain.AccountPrivilege(a.PrivilegeLevel),
		RestrictionLevel:  domain.AccountRestriction(a.RestrictionLevel),
		RestrictionExpiry: restrictionExpiry,
		Verified:          a.Verified,
	}
	ad, err := domain.NewAccount(params)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func verificationToDomain(v *sqlc.UsersVerification) (*domain.Verification, error) {
	params := &domain.NewVerificationParams{
		ID:        v.ID,
		UserID:    v.UserID,
		Token:     v.Token,
		ExpiresAt: v.ExpiresAt,
	}
	vd, err := domain.NewVerification(params)
	if err != nil {
		return nil, err
	}
	return vd, nil
}
