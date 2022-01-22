package usecase

import (
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/go-playground/validator/v10"
	"shiva/shiva-auth/helpers/s3"
	"shiva/shiva-auth/internal/categories"
	"shiva/shiva-auth/utils/baseErrors"
)

type Usecase struct {
	data     categories.Repository
	validate *validator.Validate
	uploader *s3manager.Uploader
}

func NewCategoriesUsecase(r categories.Repository, uploader *s3manager.Uploader) categories.Usecase {
	return &Usecase{
		data:     r,
		validate: validator.New(),
		uploader: uploader,
	}
}

func (uc Usecase) GetAll(search string, key string) (data []categories.Domain, err error) {
	res, err := uc.data.GetAll(search, key)
	if err != nil {
		return []categories.Domain{}, err
	}
	return res, nil
}

func (uc Usecase) GetById(id uint) (categories.Domain, error) {
	u, err := uc.data.GetById(id)
	if err != nil {
		return categories.Domain{}, err
	} else if u.ID == 0 {
		return categories.Domain{}, baseErrors.ErrNotFound
	}
	return u, nil
}

func (uc Usecase) Create(d categories.Domain) (categories.Domain, error) {
	img, err := s3.ImageUpload(uc.uploader, d.ImageHeader)
	d.ImageUrl = img.Location
	cls, err := uc.data.Create(d)
	if err != nil {
		return categories.Domain{}, err
	}
	return cls, nil
}

func (uc Usecase) Update(d categories.Domain) (categories.Domain, error) {
	if d.ImageHeader != nil {
		img, err := s3.ImageUpload(uc.uploader, d.ImageHeader)
		d.ImageUrl = img.Location
		data, err := uc.data.Update(d)
		if err != nil {
			return categories.Domain{}, err
		}
		return data, nil
	} else {
		data, err := uc.data.UpdateWithoutImage(d)
		if err != nil {
			return categories.Domain{}, err
		}
		return data, nil
	}
}

func (uc Usecase) Delete(id uint) error {
	u, err := uc.data.GetById(id)
	if err != nil {
		return err
	} else if u.ID == 0 {
		return baseErrors.ErrNotFound
	}
	err = uc.data.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
