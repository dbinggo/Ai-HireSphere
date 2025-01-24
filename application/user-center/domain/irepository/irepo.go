package irepository

import (
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	ireidsaccess "Ai-HireSphere/application/user-center/domain/irepository/ireids_access"
)

type IRepoBroker interface {
	idataaccess.IDataAccess
	ireidsaccess.IRedisAccess
}
