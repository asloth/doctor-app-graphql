package database

import (
	"time"

	"gorm.io/gorm"
)

var Handler *BaseHandler

type BaseHandler struct {
	db *gorm.DB
}

func NewBaseHandler(db *gorm.DB) *BaseHandler {
	return &BaseHandler{
		db: db,
	}
}

//Funcion que migra los schemas
func (h *BaseHandler) Migrate() error {
	// Migrate the schema
	err := h.db.AutoMigrate(&UserType{},
		&User{},
		&Patient{},
		&MedicalCenter{},
		&IdentificationDocument{},
		&Doctor{},
		&DoctorSchedule{},
		&ClinicHistory{},
		&Attention{},
	)
	if err != nil {
		panic("migration failed ")
	}
	return err
}

//Schemas que representan las tablas de la base de datos
type UserType struct {
	ID    uint `gorm:"primaryKey;autoIncrement"`
	Name  string
	Users []User
}

type User struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	Name       string `gorm:"not null"`
	Password   string `gorm:"not null"`
	UserTypeID uint
	UserType   UserType `gorm:"foreignKey:UserTypeID;references:ID"`
}

type Patient struct {
	ID                       uint `gorm:"primaryKey;autoIncrement"`
	UserID                   uint
	User                     User
	ClinicHistory            ClinicHistory
	DocumentId               string
	Name                     string
	Lastname                 string
	Email                    string
	PhoneNumber              string
	SisId                    string
	Genre                    string
	Birthdate                time.Time
	Address                  string
	District                 string
	Region                   string
	IdentificationDocumentID uint
	IdentificationDocument   IdentificationDocument `gorm:"foreignKey:IdentificationDocumentID;references:ID"`
	DocumentIdUrl            string
	Attentions               []Attention
}

type Doctor struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	UserID     uint
	User       User `gorm:"foreignKey:UserID;references:ID"`
	Name       string
	Lastname   string
	DocumentId string
	Email      string
	Schedule   []DoctorSchedule
	Attentions []Attention
}

type DoctorSchedule struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	InitDate  time.Time
	EndDate   time.Time
	StartHour uint8 `gorm:"check:start_hour<24;check:start_hour>0"`
	EndHour   uint8 `gorm:"check:end_hour<24;check:end_hour>0"`
	DoctorID  uint
}

type IdentificationDocument struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Name     string
	Patients []Patient
}

type ClinicHistory struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	DocumentUrl string
	PatientID   uint
}

type MedicalCenter struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Name       string
	Address    string
	OpenHour   uint8 `gorm:"check:open_hour<24;check:open_hour>0"`
	CloseHour  uint8 `gorm:"check:close_hour<24;check:close_hour>0"`
	Attentions []Attention
}

type Attention struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	DoctorID        uint
	Doctor          Doctor `gorm:"foreignKey:DoctorID;references:ID"`
	PatientID       uint
	Patient         Patient `gorm:"foreignKey:PatientID;references:ID"`
	MedicalCenterID uint
	MedicalCenter   MedicalCenter `gorm:"foreignKey:MedicalCenterID;references:ID"`
	Date            time.Time
}
