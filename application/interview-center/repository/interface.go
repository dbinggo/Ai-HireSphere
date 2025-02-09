package repository

// 抽象化存储层
type Repo interface {
	Get(where Interview) (Interview, error)
	Update(where Interview, data Interview) error
	Create(data Interview) error
	Delete(where Interview) error

	Begin() Repo
	Rollback() Repo
	Commit() Repo
	Tx(func(repo Repo) error) error
}
