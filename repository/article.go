package repository

import (
	"backend-bangkit/entity"
	"backend-bangkit/pkg/errs"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *entity.Article) errs.MessageErr
	FindByID(id string) (*entity.Article, errs.MessageErr)
	FindAll(tagging string) ([]entity.Article, errs.MessageErr)
	Delete(id string) errs.MessageErr
	Update(article *entity.Article) errs.MessageErr
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

// Article Repository Implementation
func (r *articleRepository) Create(article *entity.Article) errs.MessageErr {
	err := r.db.Create(article).Error

	if err != nil {
		return errs.NewInternalServerError("cannot create article")
	}

	return nil
}

func (r *articleRepository) FindByID(id string) (*entity.Article, errs.MessageErr) {
	var article entity.Article
	err := r.db.Preload("Tag").First(&article, "id = ?", id).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NewNotFound("article not found")
		}
		return nil, errs.NewInternalServerError("cannot find article by ID")
	}

	return &article, nil
}

func (r *articleRepository) FindAll(tagging string) ([]entity.Article, errs.MessageErr) {
	var articles []entity.Article

	var err error
	if tagging != "" {
		err = r.db.Preload("Tag"). // Memuat relasi Tag
						Where("taggings.name = ?", tagging).                     // Filter berdasarkan nama tag
						Joins("JOIN taggings ON taggings.id = articles.tag_id"). // Bergabung dengan tabel Tagging
						Find(&articles).Error
	} else {
		err = r.db.Preload("Tag").Find(&articles).Error
	}

	if err != nil {
		return nil, errs.NewInternalServerError("cannot find articles")
	}

	return articles, nil
}

func (r *articleRepository) Delete(id string) errs.MessageErr {
	err := r.db.Delete(&entity.Article{}, "id = ?", id).Error

	if err != nil {
		return errs.NewInternalServerError("cannot delete article")
	}

	return nil
}

func (r *articleRepository) Update(article *entity.Article) errs.MessageErr {
	err := r.db.Save(article).Error

	if err != nil {
		return errs.NewInternalServerError("cannot update article")
	}

	return nil
}
