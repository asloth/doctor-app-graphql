package database

import (
	"strconv"

	"github.com/asloth/doctor-app-graphql/graph/model"
)

func (h *BaseHandler) CreateUser(name *string, email *string, password *string, usertypeid *int) (*model.User, error) {
	ut := UserType{ID: uint(*usertypeid)}
	user := User{Name: *name, Email: *email, Password: *password, UserType: ut}

	h.db.First(&ut)
	result := h.db.Create(&user) // pass pointer of data to Create

	return &model.User{
		ID:       strconv.FormatUint(uint64(user.ID), 10),
		Name:     *name,
		Email:    *email,
		Password: *password,
		UserType: &model.UserType{
			ID:   strconv.FormatUint(uint64(ut.ID), 10),
			Name: ut.Name,
		},
	}, result.Error
}
