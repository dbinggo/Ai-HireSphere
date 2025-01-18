package repository

// 抽象化存储层
type Repo interface {
	Get(where Resume) (Resume, error)
	Update(where Resume, data Resume) error
	Create(data Resume) error
	Delete(where Resume) error

	Begin() Repo
	Rollback() Repo
	Commit() Repo
	Tx(func(repo Repo) error) error
}
