package database

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

// Run is Function for init a database
func Run() {

	dbUser := beego.AppConfig.String("psqluser")
	dbPass := beego.AppConfig.String("psqlpass")
	dbHost := beego.AppConfig.String("psqlurls")
	dbPort := beego.AppConfig.String("psqlport")
	dbName := beego.AppConfig.String("psqldb")
	dns := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	if beego.BConfig.RunMode == "dev" {
		fmt.Println("==== try to connecting to ===>", dns)
	}

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default",
		"postgres",
		dns,
	)

	orm.RunSyncdb("default", false, true)

	defaultDB, err := orm.GetDB("default")
	if err != nil {
		fmt.Println("===connecting database error==")
		beego.Emergency("Get default DB error:" + err.Error())
	}
	defaultDB.SetConnMaxLifetime(3595 * time.Second)
	fmt.Println("default==", defaultDB)

	// orm.RegisterModel(new(models.News))
	// orm.RegisterModel(new(Profile))
}
