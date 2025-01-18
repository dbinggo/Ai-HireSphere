package repo_gorm

import (
	"Ai-HireSphere/application/interview-center/repository"
	"Ai-HireSphere/common/coze"
	"Ai-HireSphere/common/utils"
	"errors"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

var _ repository.Repo = (*Repo)(nil)

func NewRepo(db *gorm.DB, api *coze.CozeApi) *Repo {
	return &Repo{
		DB:   db,
		Coze: api,
	}
}

func (b *Repo) Get(where repository.Resume) (repository.Resume, error) {
	var repoWhere, repoData ResumeModel
	var data repository.Resume

	err := copier.Copy(&repoWhere, where)
	if err != nil {
		return repository.Resume{}, err
	}

	err = b.DB.Model(b.Model).Where(repoWhere).First(&repoData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.Resume{}, repository.NotFound
		}
		return repository.Resume{}, err
	}

	err = copier.Copy(&data, repoData)
	if err != nil {
		return repository.Resume{}, err
	}

	return data, nil
}

func (b *Repo) Update(where repository.Resume, data repository.Resume) error {
	var repoWhere, repoData ResumeModel

	err := copier.Copy(&repoWhere, where)
	if err != nil {
		return err
	}

	err = copier.Copy(&repoData, data)
	if err != nil {
		return err
	}

	err = b.DB.Model(b.Model).Where(repoWhere).Updates(repoData).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) Create(data repository.Resume) error {
	var repoData ResumeModel
	err := copier.Copy(&repoData, data)
	if err != nil {
		return err
	}

	err = b.Coze.Doc.CreateDoc([]utils.FileBase{data.File})
	if err != nil {
		return err
	}

	err = b.DB.Model(b.Model).Create(repoData).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) Delete(where repository.Resume) error {
	var repoWhere ResumeModel
	err := copier.Copy(&repoWhere, where)
	if err != nil {
		return err
	}

	err = b.DB.Model(b.Model).Delete(repoWhere).Error
	if err != nil {
		return err
	}

	return nil
}

func (b *Repo) Begin() repository.Repo {
	tx := new(Repo)
	tx.DB = b.DB.Begin()
	return tx
}

func (b *Repo) Rollback() repository.Repo {
	b.DB = b.DB.Rollback()
	return b
}

func (b *Repo) Commit() repository.Repo {
	b.DB = b.DB.Commit()
	return b
}

func (b *Repo) Tx(f func(repo repository.Repo) error) error {
	var err error
	tx := b.Begin()
	defer func() {
		if wrong := recover(); wrong != nil {
			tx.Rollback()
			return
		}
	}()

	err = f(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
