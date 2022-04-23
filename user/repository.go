package user

import (
	"gorm.io/gorm"
)

// Buat interface nge Savenya
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *gorm.DB
}

// Fungsi dari struct reposiroty
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Menerapkan method dari interface nya Repository ke struct repository
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
