package services

import (
	"rol/app/interfaces"
	"rol/domain"
	"rol/dtos"

	"github.com/sirupsen/logrus"
)

//NewAppLogService preparing domain.AppLog repository for passing it in DI
//Params
//	rep - generic repository with http log instantiated
//	log - logrus logger
//Return
//	New app log service
func NewAppLogService(rep interfaces.IGenericRepository[domain.AppLog], log *logrus.Logger) (interfaces.IGenericService[
	dtos.AppLogDto,
	dtos.AppLogDto,
	dtos.AppLogDto,
	domain.AppLog], error) {
	return NewGenericService[dtos.AppLogDto, dtos.AppLogDto, dtos.AppLogDto](rep, log)
}
