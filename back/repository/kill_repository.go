package repository

import (
	"backend-avanzada/models"
	"errors"

	"gorm.io/gorm"
)

type KillRepository struct {
	db *gorm.DB
}

func NewKillRepository(db *gorm.DB) *KillRepository {
	return &KillRepository{
		db: db,
	}
}

func (k *KillRepository) FindAll() ([]*models.Kill, error) {
	var kills []*models.Kill
	err := k.db.Find(&kills).Error
	if err != nil {
		return nil, err
	}
	return kills, nil
}

func (k *KillRepository) Save(data *models.Kill) (*models.Kill, error) {
	// var person models.Person
	// if err := k.db.First(&person, data.PersonId).Error; err != nil {
	// 	return nil, err
	// }
	// if person.State == false {
	// 	return nil, errors.New("no se puede guardar: la persona está inactiva")
	// }
	err := k.db.Save(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (k *KillRepository) FindById(id int) (*models.Kill, error) {
	var kill models.Kill
	err := k.db.Where("id = ?", id).First(&kill).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &kill, nil
}

func (k *KillRepository) Update(id int, m *models.Kill) (*models.Kill, error) {
	existing := &models.Kill{}
	if err := k.db.First(existing, id).Error; err != nil {
		return nil, err
	}

	if err := k.db.Model(existing).Updates(m).Error; err != nil {
		return nil, err
	}

	return existing, nil
}

func (k *KillRepository) Delete(data *models.Kill) error {
	return errors.New("this method is not implemented")
}
