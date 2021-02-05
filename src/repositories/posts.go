package repositories

import (
	"api/src/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Posts struct {
	db *gorm.DB
}

func NewRepositoryPosts(db *gorm.DB) *Posts {
	return &Posts{db}
}

func (p *Posts) Create(post models.Posts) (models.Posts, error) {
	if err := p.db.Create(&post).Error; err != nil {
		return models.Posts{}, err
	}
	return post, nil
}
func (p *Posts) Find(ID uint) (models.Posts, error) {
	var posts models.Posts
	if err := p.db.Preload(clause.Associations).Find(&posts, ID).Error; err != nil {
		return models.Posts{}, err
	}
	return posts, nil
}
func (p *Posts) List(param string) []models.Posts {
	var posts []models.Posts
	p.db.Where("title like ?", "%"+param+"%").Preload(clause.Associations).Find(&posts)
	return posts
}
func (p *Posts) Update(posts models.Posts, IDPost uint, ID uint) error {
	if err := p.db.Model(posts).Where("id = ? and author_id = ?", IDPost, ID).
		Updates(posts).Error; err != nil {
		return err
	}
	return nil
}
func (p *Posts) Delete(ID uint) error {
	if err := p.db.Delete(&models.Posts{}, "id = ?", ID).Error; err != nil {
		return err
	}
	return nil
}
func (p *Posts) Like(IDPost, ID uint) error {
	if err := p.db.Create(models.Likes{UsersID: ID, PostsID: IDPost}).Error; err != nil {
		return err
	}
	return nil
}
func (p *Posts) UnLike(IDPost, ID uint) error {
	if err := p.db.Unscoped().Delete(models.Likes{}, "posts_id = ? and users_id = ?", IDPost, ID).Error; err != nil {
		return err
	}
	return nil
}
