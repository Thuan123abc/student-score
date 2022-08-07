package database

import "gorm.io/gorm"

type StudentRepo struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *StudentRepo {
	db.AutoMigrate(&StudentModel{})
	return &StudentRepo{
		db: db,
	}
}

func (s StudentRepo) CreateStudent(model StudentModel) error {
	err := s.db.Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}

func (s StudentRepo) DeleteStudent(mssv string) error {
	err := s.db.Where("mssv=?", mssv).Delete(&StudentModel{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s StudentRepo) UpdateStudent(model StudentModel) error {
	return nil
}

func (s StudentRepo) GetListStudent() ([]StudentModel, error) {
	return nil, nil
}

func (s StudentRepo) GetStudent(mssv string) (*StudentModel, error) {
	return nil, nil
}
