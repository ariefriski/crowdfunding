package user

import "gorm.io/gorm"

// Buat interface nge Savenya
type Repository interface {
	Save(user User) (User, error)
}

// Implementasi dari Interface Repository, tipe data konek DB
type repository struct {
	db *gorm.DB
}

// Instance dari struct repository
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Simpan ke databasenya pake methode interface Repository (Save)
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
