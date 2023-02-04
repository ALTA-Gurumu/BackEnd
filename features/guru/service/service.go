package service

import (
	"Gurumu/features/guru"
	"Gurumu/helper"
	"errors"
	"fmt"
	"log"
	"mime/multipart"
	"strings"

	"github.com/go-playground/validator"
)

type guruUseCase struct {
	qry guru.GuruData
	vld *validator.Validate
}

func New(gd guru.GuruData) guru.GuruService {
	return &guruUseCase{
		qry: gd,
		vld: validator.New(),
	}
}

// Register implements guru.GuruService
func (guc *guruUseCase) Register(newGuru guru.Core) (guru.Core, error) {
	hashed, err := helper.GeneratePassword(newGuru.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return guru.Core{}, errors.New("password process error")
	}

	err = guc.vld.Struct(&newGuru)
	if err != nil {
		msg := helper.ValidationErrorHandle(err)
		fmt.Println("msg", msg)
		return guru.Core{}, errors.New(msg)
	}

	newGuru.Password = string(hashed)
	res, err := guc.qry.Register(newGuru)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "data sudah terdaftar"
		} else {
			msg = "terdapat masalah pada server"
		}
		return guru.Core{}, errors.New(msg)
	}

	return res, nil
}

// Delete implements guru.GuruService
func (guc *guruUseCase) Delete(token interface{}) error {
	id := helper.ExtractToken(token)
	if id <= 0 {
		return errors.New("data not found")
	}

	err := guc.qry.Delete(uint(id))
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data tidak ditemukan"
		} else {
			msg = "terdapat masalah pada server"
		}
		return errors.New(msg)
	}

	return nil
}

// Profile implements guru.GuruService
func (*guruUseCase) Profile(token interface{}) (guru.Core, error) {
	panic("unimplemented")
}

// Update implements guru.GuruService
func (guc *guruUseCase) Update(token interface{}, updateData guru.Core, avatar *multipart.FileHeader, ijazah *multipart.FileHeader) error {
	userID := helper.ExtractToken(token)
	if userID <= 0 {
		return fmt.Errorf("token tidak valid")
	}

	if err := guc.vld.Struct(&updateData); err != nil {
		log.Println(err)
		return fmt.Errorf("validation error: %s", helper.ValidationErrorHandle(err))
	}

	var avatarURL, ijazahURL string
	if avatar != nil {
		res, err := guc.qry.GetByID(uint(userID))
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "not found") {
				return fmt.Errorf("data guru tidak ditemukan")
			}
			return fmt.Errorf("gagal mendapatkan data guru: %s", err)
		}

		avatarURL, err = helper.UploadTeacherProfilePhotoS3(*avatar, res.Email)
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "kesalahan input") {
				return fmt.Errorf("gagal upload avatar: %s", err)
			}
			return fmt.Errorf("gagal upload avatar: system server error")
		}
	}

	if ijazah != nil {
		res, err := guc.qry.GetByID(uint(userID))
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "not found") {
				return fmt.Errorf("data guru tidak ditemukan")
			}
			return fmt.Errorf("gagal mendapatkan data guru: %s", err)
		}

		ijazahURL, err = helper.UploadTeacherCertificateS3(*ijazah, res.Email)
		if err != nil {
			log.Println(err)
			if strings.Contains(err.Error(), "kesalahan input") {
				return fmt.Errorf("gagal upload ijazah: %s", err)
			}
			return fmt.Errorf("gagal upload ijazah: system server error")
		}
	}

	if avatarURL != "" {
		updateData.Avatar = avatarURL
	}
	if ijazahURL != "" {
		updateData.Ijazah = ijazahURL
	}

	if err := guc.qry.Update(uint(userID), updateData); err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "tidak ditemukan") {
			return fmt.Errorf("data guru tidak ditemukan")
		}
		return fmt.Errorf("gagal update data guru: %s", err)
	}

	return nil
}
