package repository

import (
	"log"

	"github.com/joaoaneto/radiup/user/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySQLConfig struct {
	DbInstance *gorm.DB
}

func NewMySQLConfig() *MySQLConfig {
	
	db, err := gorm.Open("mysql", "radiup:radiupProject2017@/radiup?charset=utf8&parseTime=True&loc=Local")
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

func (sup SimpleUserPersistor) Create(u *model.User) error {
	
	err := sup.db.DbInstance.Create(u).Error
	if err != nil {
		return err
	}

	return nil
}

func (sup SimpleUserPersistor) Update(u *model.User) error {

	err := sup.db.DbInstance.Save(u).Error
	if err != nil {
		return err
	}

	return nil
}

func (sup SimpleUserPersistor) Remove(username string) error {
	
	user, err := sup.Search(username)
	if err != nil {
		return err
	}

	err = sup.db.DbInstance.Delete(&user).Error
	if err != nil {
		return err
	}

	err = sup.db.DbInstance.Where(&model.SpotifyToken{UserID: user.UserID}).Delete(&model.SpotifyToken{}).Error
	if err != nil {
		return err
	}
	
	return nil
}

func (sup SimpleUserPersistor) Search(username string) (model.User, error) {
	
	user := model.User{}

	err := sup.db.DbInstance.Where(&model.User{Username: username}).First(&user).Error
	if err != nil {
		return model.User{}, err
	}
	
	return user, nil
}

func (sup SimpleUserPersistor) SearchAll() ([]model.User, error) {
	return nil, nil
}