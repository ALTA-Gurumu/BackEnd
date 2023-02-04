package guru

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Core struct {
	ID        uint
	Nama      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	Telepon   string
	Deskripsi string
	Pelajaran string
	Alamat    string
	Avatar    string
	Ijazah    string
	Role      string
	Latitude  string
	Longitude string
}

type GuruHandler interface {
	Register() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
}

type GuruService interface {
	Register(newGuru Core) (Core, error)
	Profile(token interface{}) (Core, error)
	Update(token interface{}, updateData Core, avatar *multipart.FileHeader, ijazah *multipart.FileHeader) error
	Delete(token interface{}) error
}

type GuruData interface {
	Register(newGuru Core) (Core, error)
	GetByID(id uint) (Core, error)
	Update(id uint, updateData Core) error
	Delete(id uint) error
}
