package category

import (
	"pasarwarga/article"
	"pasarwarga/models"

	"github.com/gosimple/slug"
)


type Service interface {
	CreateCategory(input CategoryInput) (models.Category, error)
	DetailCategory(input CategoryIDInput)(models.Category, error)
	ListCategory()([]models.Category,error)
	UpdateCategory(input CategoryIDInput, categorydata CategoryInput) (models.Category, error)
	DeleteCategory(input CategoryIDInput)error
}

type service struct {
	repository Repository
	ArticleRepository article.Repository
}

func NewService(repository Repository, ArticleRepository article.Repository) *service {
	return &service{repository, ArticleRepository}
}
func (s *service)CreateCategory(input CategoryInput) (models.Category, error){

	GetCategory := models.Category{}
	GetCategory.CategoryName = input.CategoryName
	GetCategory.CategorySlug = slug.Make(input.CategoryName)
	SaveCategory, err := s.repository.CreateCategory(GetCategory)

	if err != nil{
		return GetCategory, err
	}

	return SaveCategory, nil
	

	
}

func (s *service)DetailCategory(input CategoryIDInput)(models.Category, error){

	
	FindCategoryByID, err := s.repository.FindCategoryID(input.ID)

	if err!= nil{
		return FindCategoryByID, err
	}

	return FindCategoryByID,nil
}

func(s *service)ListCategory()([]models.Category,error){

	ListAllCategory, err := s.repository.ListCategory()

	if err !=nil{

		return ListAllCategory,err
	}

	return ListAllCategory, nil

}

func (s *service)UpdateCategory(input CategoryIDInput, categorydata CategoryInput) (models.Category, error){

	GetCurrentID, err := s.repository.FindCategoryID(input.ID)

	if err !=nil{
		return GetCurrentID, err
	}

	GetCurrentID.CategoryName = categorydata.CategoryName
	GetCurrentID.CategorySlug  = slug.Make(GetCurrentID.CategoryName)
	UpdateCategory, err := s.repository.UpdateCategory(GetCurrentID)

	if err != nil{
		return UpdateCategory,err
	}

	return GetCurrentID,nil

}

func(s *service)DeleteCategory(input CategoryIDInput)error{

	GetCurrentID, err := s.repository.FindCategoryID(input.ID)

	if err !=nil{
		return err
	}
	
	DeleteArticleErr :=s.ArticleRepository.DeleteArticleByCategoryID(int(GetCurrentID.ID))

	if DeleteArticleErr != nil{
		return DeleteArticleErr
	}

	DeleteCategoryErr := s.repository.DeleteCategory(int(GetCurrentID.ID))

	if DeleteCategoryErr != nil{
		return DeleteCategoryErr
	}

	return nil

}



