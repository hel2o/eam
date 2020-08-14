package db

import (
	"fmt"

	"github.com/go-xorm/core"

	"github.com/go-xorm/xorm"
	"github.com/hel2o/eam/model"
	"github.com/lunny/log"
	_ "github.com/mattn/go-sqlite3"
)

var engine *xorm.Engine

func InitDB() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "./eam.db")
	if err != nil {
		log.Fatal(err)
	}
	engine.ShowSQL(false)
	engine.SetLogLevel(core.LOG_INFO)
	err = engine.Sync2(new(model.EAM))
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
}

func GetEAMByID(id string) (eam model.EAM, err error) {
	_, err = engine.ID(id).Get(&eam)
	return
}

func GetEAM(search string, page, limit int) (eam []model.EAM, count int64, err error) {
	var where string
	if search == "" {
		where = "1=1"
	} else {
		where = "%" + search + "%"
		where = fmt.Sprintf(`ip like '%s' 
									or place like '%s' 
									or name like '%s'
									or remarks like '%s' 
									or person like '%s'
									or company like '%s'
									or department like '%s'
									or os like '%s'
									or app_service like '%s'
									or sql_type like '%s'
									or server_model like '%s'
									or server_para like '%s'
									or developer like '%s'
									or server_code like '%s'
									or real_host like '%s'`,
			where, where, where, where, where, where, where, where, where, where, where, where, where, where, where)
	}
	count, err = engine.Where(where).Limit(limit, calcStartPage(limit, page)).FindAndCount(&eam)
	return
}

func AddEAM(eam model.EAM) (err error) {
	_, err = engine.Insert(&eam)
	return
}

func ModifyEAM(eam model.EAM) (err error) {
	_, err = engine.ID(eam.Id).Update(&eam)
	return
}

func DelEAM(eam []model.EAM) (int64, error) {
	var ids []int
	for _, v := range eam {
		ids = append(ids, v.Id)
	}
	return engine.In("id", ids).Delete(new(model.EAM))
}

func calcStartPage(limit, page int) (start int) {
	if page == 1 {
		start = 0
		return
	}
	if page > 1 {
		start = limit * (page - 1)
	}
	return
}
