package post

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindAll() ([]Post, error) {
	var posts []Post
	err := r.db.Order("id desc").Find(&posts).Error
	return posts, err
}

func (r *Repository) FindByID(id int) (Post, error) {
	var post Post
	err := r.db.First(&post, id).Error
	return post, err
}

func (r *Repository) Save(post *Post) error {
	return r.db.Create(post).Error
}

func (r *Repository) Delete(id int) error {
	return r.db.Delete(&Post{}, id).Error
}
