package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type AccountPrivilege string
type AccountRestriction string

const (
	AccountPrivilegeUser            AccountPrivilege   = "user"
	AccountPrivilegeModerator       AccountPrivilege   = "moderator"
	AccountPrivilegeAdmin           AccountPrivilege   = "admin"
	AccountRestrictionNone          AccountRestriction = "none"
	AccountRestrictionCommunication AccountRestriction = "communication"
	AccountRestrictionFull          AccountRestriction = "full"
)

var (
	ErrInvalidPrivilege   = errors.New("account: invalid privilege level")
	ErrInvalidRestriction = errors.New("account: invalid restriction level")
)

func (e AccountPrivilege) Valid() bool {
	switch e {
	case AccountPrivilegeUser,
		AccountPrivilegeModerator,
		AccountPrivilegeAdmin:
		return true
	}
	return false
}

func (e AccountRestriction) Valid() bool {
	switch e {
	case AccountRestrictionNone,
		AccountRestrictionCommunication,
		AccountRestrictionFull:
		return true
	}
	return false
}

type Account struct {
	ID                uuid.UUID
	Username          string
	Email             string
	HashedPassword    string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastLoggedIn      time.Time
	PrivilegeLevel    AccountPrivilege
	RestrictionLevel  AccountRestriction
	RestrictionExpiry time.Time
	Verified          bool
}

type NewAccountParams struct {
	ID                uuid.UUID
	Username          string
	Email             string
	HashedPassword    string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	LastLoggedIn      time.Time
	PrivilegeLevel    AccountPrivilege
	RestrictionLevel  AccountRestriction
	RestrictionExpiry time.Time
	Verified          bool
}

func NewAccount(params *NewAccountParams) (*Account, error) {
	a := &Account{
		ID:                params.ID,
		Username:          params.Username,
		Email:             params.Email,
		HashedPassword:    params.HashedPassword,
		CreatedAt:         params.CreatedAt,
		UpdatedAt:         params.UpdatedAt,
		LastLoggedIn:      params.LastLoggedIn,
		PrivilegeLevel:    params.PrivilegeLevel,
		RestrictionLevel:  params.RestrictionLevel,
		RestrictionExpiry: params.RestrictionExpiry,
		Verified:          params.Verified,
	}
	_, err := a.Valid()
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Account) Valid() (bool, error) {
	if !a.PrivilegeLevel.Valid() {
		return false, ErrInvalidPrivilege
	}
	if !a.RestrictionLevel.Valid() {
		return false, ErrInvalidRestriction
	}
	return true, nil
}
