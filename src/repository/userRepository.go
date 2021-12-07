package repository 
import (
	"todo/graph/model"
	"gorm.io/gorm"
)
type UserRepository interface {
	Save(user *model.User) (*model.User, error)
	GetAll() ([]*model.User, error)
	Update(input model.UpdateUser) (bool,error)
	FindOne(id int) (*model.User, error)
	Destroy(id int) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	} 
}

func (repo *userRepository)Save(user *model.User) (*model.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepository)GetAll() ([]*model.User, error) {
	var users []*model.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (repo *userRepository)FindOne(id int) (*model.User, error) {
	var user *model.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}
func (repo *userRepository)Update(input model.UpdateUser) (bool, error) {
	if err := repo.db.Model(model.User{}).Where("id = ?", input.ID).Updates(input).Error; err != nil {
		return false, err
	}
	return true, nil
}
func (repo *userRepository)Destroy(id int) (*model.User, error) {
	var user *model.User
	if err := repo.db.Delete(&user, id).Error; err != nil {
		return nil, err
	}	
	return user, nil
}