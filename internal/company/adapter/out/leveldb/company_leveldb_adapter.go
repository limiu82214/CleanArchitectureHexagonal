package leveldb

import (
	"strconv"

	"github.com/limiu82214/CleanArchitectureHexagonal/internal/company/application/port/out"
	"github.com/limiu82214/CleanArchitectureHexagonal/internal/company/domain"
	"github.com/limiu82214/CleanArchitectureHexagonal/pkg/gob"
	"github.com/syndtr/goleveldb/leveldb"
)

type companyLevelDBAdapter struct {
	db *leveldb.DB
}

func NewCompanyGinAdapter(db *leveldb.DB) out.ILoadCompanyPort {
	return &companyLevelDBAdapter{
		db: db,
	}
}

func (cla *companyLevelDBAdapter) GetSite(siteID int) domain.Site {
	// 假資料---
	s1 := domain.Site{Id: 1, SiteName: "site_name"}
	s2 := domain.Site{Id: 2, SiteName: "site_name2"}
	s3 := domain.Site{Id: 3, SiteName: "site_name3"}
	gs1, _ := gob.StoreStructToByte(s1)
	gs2, _ := gob.StoreStructToByte(s2)
	gs3, _ := gob.StoreStructToByte(s3)
	cla.db.Put([]byte([]byte(strconv.Itoa(s1.Id))), gs1, nil)
	cla.db.Put([]byte([]byte(strconv.Itoa(s2.Id))), gs2, nil)
	cla.db.Put([]byte([]byte(strconv.Itoa(s3.Id))), gs3, nil)
	// ---

	s := domain.Site{}
	id := []byte(strconv.Itoa(siteID))
	data, _ := cla.db.Get(id, nil)
	if data != nil {
		gob.GetStrutFromByte(data, &s)
	}
	return s
}
