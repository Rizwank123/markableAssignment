package service

import (
	"context"
	"errors"

	"github.com/gofrs/uuid/v5"

	"github.com/markable/internal/domain"
	"github.com/markable/internal/pkg/config"
	"github.com/markable/internal/pkg/security"
	"github.com/markable/internal/pkg/util"
)

type userServiceImpl struct {
	apu util.AppUtil
	cfg config.MarkAbleConfig
	scm security.Manager
	tr  domain.Transactioner
	ur  domain.UserRepository
}

func NewUserService(apu util.AppUtil, cfg config.MarkAbleConfig, scm security.Manager, tr domain.Transactioner, ur domain.UserRepository) domain.UserService {
	return &userServiceImpl{
		apu: apu,
		cfg: cfg,
		scm: scm,
		tr:  tr,
		ur:  ur,
	}
}

// FindByID implements domain.UserService.
func (s *userServiceImpl) FindByID(id uuid.UUID) (result domain.User, err error) {
	return s.ur.FindByID(context.Background(), id)
}

// FindByUserName implements domain.UserService.
func (s *userServiceImpl) FindByUserName(username string) (result domain.User, err error) {
	return s.ur.FindByUserName(context.Background(), username)
}

// Login implements domain.UserService.
func (s *userServiceImpl) Login(in domain.LoginInput) (result domain.LoginOutput, err error) {
	usr, err := s.ur.FindByUserName(context.Background(), in.UserName)
	if err != nil {
		return result, err
	}
	if usr.UserName != "" && usr.ID.IsNil() {
		return result, errors.New("user with this username:" + in.UserName + " does not exist")
	}
	match, err := s.apu.PasswordCheck(*usr.Password, in.Password)
	if err != nil || !match {
		return result, errors.New("wrong password")
	}
	// generate token
	ti := security.TokenMetadata{
		UserID: usr.ID.String(),
		Role:   usr.Role,
	}
	token, err := s.scm.GenerateAuthToken(ti)
	if err != nil {
		return result, err
	}
	result = domain.LoginOutput{
		Token:     token,
		ExpiresIn: int64(s.cfg.AuthExpiryPeriod),
	}
	return result, err
}

// Register implements domain.UserService.
func (s *userServiceImpl) Register(in domain.RegisterUserInput) (result domain.User, err error) {
	pass, err := s.apu.EncryptPassword(in.Password)
	if err != nil {
		return result, err
	}
	result = domain.User{
		UserName: in.UserName,
		FullName: in.FullName,
		Role:     string(in.Role),
		Password: &pass,
	}
	err = s.ur.CreateUser(context.Background(), &result)
	if err != nil {
		return result, err
	}
	return result, err
}

// UpdateUser implements domain.UserService.
func (s *userServiceImpl) UpdateUser(id uuid.UUID, in domain.UpdateUserInput) (result domain.User, err error) {
	result, err = s.ur.FindByID(context.Background(), id)
	if err != nil {
		return result, err
	}
	if in.UserName != "" {
		result.UserName = in.UserName
	}
	if in.FullName != "" {
		result.FullName = in.FullName
	}
	if in.Password != "" {
		pass, err := s.apu.EncryptPassword(in.Password)
		if err != nil {
			return result, err
		}
		result.Password = &pass
	}
	err = s.ur.UpdateUser(context.Background(), &result)
	return result, err

}

// DeleteUser implements domain.UserService.
func (s *userServiceImpl) DeleteUser(id uuid.UUID) (err error) {
	result, err := s.ur.FindByID(context.Background(), id)
	if err != nil {
		return err
	}
	if result.ID.IsNil() {
		return errors.New("user with this id:" + id.String() + " does not exist")
	}
	err = s.ur.DeleteUser(context.Background(), result.ID)
	if err != nil {
		return err
	}
	return nil

}
