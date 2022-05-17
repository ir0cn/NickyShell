package service

import (
	"NickyShell/nicky/db"
	"NickyShell/nicky/users/organization"
)

type Service struct {
	dbEngine *db.DbEngine
}

func (svc *Service) StartDbService() error {
	var err error
	if svc.dbEngine, err = db.NewDbEngine("127.0.0.1:3306", "root", "123456", "nicky"); err != nil {
		return err
	}

	if err := organization.Prepare(svc.dbEngine); err != nil {
		return err
	}
	return nil
}

func Start() error {
	svc := &Service{}
	if err := svc.StartDbService(); err != nil {
		return err
	}
	return svc.StartWebService()
}
