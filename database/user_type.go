package database

import (
	"strconv"

	"github.com/asloth/doctor-app-graphql/graph/model"
)

func (h *BaseHandler) CreateUserType(name *string) (*model.UserType, error) {

	userType := UserType{Name: *name}

	result := h.db.Create(&userType) // pass pointer of data to Create

	return &model.UserType{
		ID:   strconv.FormatUint(uint64(userType.ID), 10),
		Name: *name,
	}, result.Error

}

func (h *BaseHandler) UpdateUserType(id, name *string) (*model.UserType, error) {
	typeid, err := strconv.ParseUint(*id, 10, 32)
	if err != nil {
		panic("convertion from string to uint failed")
	}
	//declaramos el objeto que buscaremos
	userType := UserType{ID: uint(typeid)}
	//buscamos el objeto que queremos
	h.db.First(&userType)
	//cambiamos el nombre
	userType.Name = *name

	result := h.db.Save(&userType)
	return &model.UserType{
		ID:   *id,
		Name: *name,
	}, result.Error
}

func (h *BaseHandler) DeleteUserType(id *string) (*model.UserType, error) {
	result := h.db.Delete(&UserType{}, id)
	return &model.UserType{
		ID: *id,
	}, result.Error
}

func (h *BaseHandler) ListUserTypes() ([]*model.UserType, error) {
	var utypes []*UserType
	var userTypes []*model.UserType
	result := h.db.Find(&utypes)
	for i := len(utypes) - 1; i >= 0; i-- {
		ut := model.UserType{
			ID:   strconv.FormatUint(uint64(utypes[i].ID), 10),
			Name: utypes[i].Name,
		}
		userTypes = append(userTypes, &ut)

	}

	return userTypes, result.Error
}
