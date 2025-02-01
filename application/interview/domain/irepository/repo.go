package irepository

import "Ai-HireSphere/application/interview/domain/irepository/idataaccess"

type IRepoBroker interface {
	idataaccess.IDataAccess
}
