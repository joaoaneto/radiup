package repository

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joaoaneto/radiup/user/model"
)

type MySQLConfig struct {
	DbInstance *gorm.DB
}

func NewMySQLConfig() *MySQLConfig {

	db, err := gorm.Open("mysql", "radiup:radiup@/radiup?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Print("Error in DB Connection.")
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)

	db.CreateTable(&model.User{})
	db.CreateTable(&model.SpotifyToken{})
	
	db.CreateTable(&model.SimpleUser{})
	db.CreateTable(&model.AdminUser{})

	db.AutoMigrate(&model.SpotifyToken{}, &model.User{})
	db.LogMode(true)

	return &MySQLConfig{db}

}

func (mysqlcfg *MySQLConfig) InitMySQLDb() {

	//This function will be responsible for model association
	//mysqlcfg.DbInstance.Model(&model.SimpleUser{}).Related(&model.User{}, "SimpleUser")
	mysqlcfg.DbInstance.Model(&model.SpotifyToken{}).
		AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
}

func (mysqlcfg *MySQLConfig) CloseMySQLDb() {

	mysqlcfg.DbInstance.Close()

}

type SimpleUserPersistor struct {
	db *MySQLConfig
}

func NewSimpleUserPersistor(mysqlcfg *MySQLConfig) SimpleUserManager {
	return SimpleUserPersistor{mysqlcfg}
}

func (sup SimpleUserPersistor) Create(su *model.SimpleUser) error {

	err := sup.db.DbInstance.Create(su).Error
	if err != nil {
		return err
	}

	return nil
}

func (sup SimpleUserPersistor) Update(su *model.SimpleUser) error {

	err := sup.db.DbInstance.Save(su).Error
	if err != nil {
		return err
	}

	return nil
}

func (sup SimpleUserPersistor) Remove(username string) error {

	simpleUser, err := sup.Search(username)
	if err != nil {
		return err
	}

	err = sup.db.DbInstance.Delete(&simpleUser).Error
	if err != nil {
		return err
	}

	err = sup.db.DbInstance.Where(&model.SpotifyToken{FKUserID: simpleUser.UserID}).Delete(&model.SpotifyToken{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (sup SimpleUserPersistor) Search(username string) (model.SimpleUser, error) {

	simpleUser := model.SimpleUser{}

	err := sup.db.DbInstance.Where("username = ?", username).First(&simpleUser).Error
	if err != nil {
		return model.SimpleUser{}, err
	}

	return simpleUser, nil
}

func (sup SimpleUserPersistor) SearchAll() ([]model.SimpleUser, error) {
	return nil, nil
}

type AdminUserPersistor struct {
	db *MySQLConfig
}

func NewAdminUserPersistor(mysqlcfg *MySQLConfig) AdminUserManager {
	return AdminUserPersistor{mysqlcfg}
}

func (aup AdminUserPersistor) Create(au *model.AdminUser) error {

	err := aup.db.DbInstance.Create(au).Error
	if err != nil {
		return err
	}

	return nil
}

func (aup AdminUserPersistor) Update(au *model.AdminUser) error {

	err := aup.db.DbInstance.Save(au).Error
	if err != nil {
		return err
	}

	return nil
}

func (aup AdminUserPersistor) Remove(username string) error {

	adminUser, err := aup.Search(username)
	if err != nil {
		return err
	}

	err = aup.db.DbInstance.Delete(&adminUser).Error
	if err != nil {
		return err
	}

	err = aup.db.DbInstance.Where(&model.SpotifyToken{FKUserID: adminUser.UserID}).Delete(&model.SpotifyToken{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (aup AdminUserPersistor) Search(username string) (model.AdminUser, error) {

	adminUser := model.AdminUser{}

	err := aup.db.DbInstance.Where("username = ?", username).First(&adminUser).Error
	
	if err != nil {
		return model.AdminUser{}, err
	}

	return adminUser, nil
}

func (aup AdminUserPersistor) SearchAll() ([]model.AdminUser, error) {
	return nil, nil
}