package organization

import (
	"NickyShell/nicky/db"
	. "NickyShell/nicky/exception"
	"time"
)

type Organization struct {
	Id         int64           `json:"-"`
	Oid        int64           `json:"id,string" xorm:"unique notnull comment('组织ID')"`
	Pid        int64           `json:"pid,string" xorm:"notnull comment('父ID')"`
	Name       string          `json:"label" xorm:"varchar(256) notnull comment('组织名称')"`
	Children   []*Organization `json:"children,omitempty" xorm:"-"`
	CreateTime time.Time       `json:"-" xorm:"created"`
	UpdateTime time.Time       `json:"-" xorm:"updated"`
}

func Prepare(dbEngine *db.DbEngine) error {
	if err := dbEngine.Sync2(new(Organization)); err != nil {
		return err
	}

	// 添加默认组织
	var org Organization
	if ok, err := dbEngine.Where("oid=0").Get(&org); err != nil {
		return err
	} else if ok != true {
		_, err := dbEngine.Insert(Organization{
			Name: "默认组织",
		})
		return err
	}
	return nil
}

func find(orgs []Organization, id int64) *Organization {
	for k, v := range orgs {
		if v.Oid == id {
			return &orgs[k]
		}
	}
	return nil
}

func ToTree(orgs []Organization) *Organization {
	if len(orgs) == 0 {
		return nil
	}
	for k, v := range orgs {
		if v.Oid > 0 {
			if f := find(orgs, v.Pid); f != nil {
				f.Children = append(f.Children, &orgs[k])
			}
		}
	}
	return &orgs[0]
}

func ListOrganizations(dbEngine *db.DbEngine) *Organization {
	orgs := make([]Organization, 0)
	err := dbEngine.Find(&orgs)
	ThrowIfError(err)
	return ToTree(orgs)
}

func UpdateOrganization(dbEngine *db.DbEngine, org *Organization) {
	_, err := dbEngine.Exec("insert into organization(oid,pid,name) values (?,?,?) "+
		"on duplicate key update name = values(name)", org.Oid, org.Pid, org.Name)
	ThrowIfError(err)
}

func DeleteOrganization(dbEngine *db.DbEngine, org *Organization) {
	_, err := dbEngine.Delete(org)
	ThrowIfError(err)
}
