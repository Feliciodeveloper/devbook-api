package repositories

import (
	"api/src/models"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Users struct {
	db *gorm.DB
}

func NewRepositoryUsers(db *gorm.DB) *Users {
	return &Users{db}
}

func (u *Users) Create(user models.Users) (models.Users, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return models.Users{}, err
	}
	return user, nil
}
func (u *Users) List(param string) []models.Users {
	var user []models.Users
	param = fmt.Sprintf("%%%s%%", param)
	u.db.Debug().Where("name like ? or nick like ?", param, param).Preload(clause.Associations).Find(&user)
	return user
}
func (u *Users) Find(param uint64) models.Users {
	var user models.Users
	u.db.Debug().Where("id = ?", param).Preload("Friends").Preload("Posts.Likes").Find(&user)
	return user
}
func (u *Users) Update(param uint64, user models.Users) models.Users {
	u.db.Model(user).Where("id = ?", param).Updates(user).Find(&user)
	return user
}
func (u *Users) Delete(param uint64) error {
	if result := u.db.Where("id = ?", param).Delete(&models.Users{}).Error; result != nil {
		return result
	}
	return nil
}
func (u *Users) Login(email string) (models.Users, error) {
	var user models.Users
	result := u.db.Select("id,password").Where("email = ?", email).Find(&user)
	if result.Error != nil {
		return models.Users{}, result.Error
	}
	return user, nil
}
func (u *Users) AddFriend(ID, IDFriend uint) error {
	if err := u.db.Create(&models.Friends{UsersID: ID, FriendID: IDFriend}).Error; err != nil {
		return err
	}
	if err := u.db.Create(&models.Friends{UsersID: IDFriend, FriendID: ID}).Error; err != nil {
		return err
	}
	return nil
}
func (u *Users) ListFriends(ID uint) ([]models.Users, error) {
	var users []models.Users
	u.db.First(&users, ID)
	if err := u.db.Debug().Model(&users).Association("Friends").Find(&users); err != nil {
		return []models.Users{}, err
	}
	return users, nil
}
func (u *Users) RemoveFriend(ID, IDFriend uint) error {
	if err := u.db.Unscoped().Delete(&models.Friends{}, "users_id = ? and friend_id = ?", ID, IDFriend).Error; err != nil {
		return err
	}
	if err := u.db.Unscoped().Delete(&models.Friends{}, "users_id = ? and friend_id = ?", IDFriend, ID).Error; err != nil {
		return err
	}
	return nil
}
func (u *Users) FindPassword(param uint64) (string, error) {
	var user models.Users
	if err := u.db.Select("password").Find(&user, param).Error; err != nil {
		return "", err
	}
	return user.Password, nil
}
func (u *Users) UpdatePassword(param string, ID uint64) error {
	if err := u.db.Table("users").Where("id = ?", ID).Update("password", param).Error; err != nil {
		return err
	}
	return nil
}
