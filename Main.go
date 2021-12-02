package main

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	repo "github.com/davide/ModRepository/repositories/mysql"
	"github.com/davide/ModRepository/routes"
	"github.com/davide/ModRepository/routes/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
)

type Credentials struct {
	Username string
	Password string
	Host     string
}

func main() {

	jsonFile, err := os.Open("credentials.json")

	if err != nil {
		log.Fatal("no credentials file")
	}

	byteValue, err1 := ioutil.ReadAll(jsonFile)

	if err1 != nil {
		log.Fatal("err pasrsing json")
	}

	var cred Credentials

	if err2 := json.Unmarshal(byteValue, &cred); err2 != nil {
		log.Fatal("err pasrsing json")
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", cred.Username, cred.Password, cred.Host, "mod_repo")
	dbase, err3 := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err3 != nil {
		log.Fatal("error connecting to database")
	}

	carRepo := repo.CarRepositoryImpl{Db: dbase}
	trackRepo := repo.TrackRepositoryImpl{Db: dbase}
	nationRepo := repo.NationsRepositoryImpl{Db: dbase}
	brandRepo := repo.BrandRepositoryImpl{Db: dbase}
	userRepo := repo.UserRepositoryImpl{Db: dbase}
	authorRepo := repo.AuthorsRepositoryImpl{Db: dbase}
	logsRepo := repo.LogRepositoryImpl{Db : dbase}

	web := routes.Web{
		CarHandler:     handlers.CarsHandlerImpl{CarCtrl: controllers.CarControllerImpl{Repo: carRepo}},
		TracksHandler:  handlers.TrackHandlerImpl{TrackCtrl: controllers.TrackControllerImpl{Repo: trackRepo}},
		NationHandler:  handlers.NationsHandlerImpl{CtrlNations: controllers.NationControllerImpl{Repo: nationRepo}},
		BrandsHandler:  handlers.BrandsHandlerImpl{BrandCtrl: controllers.BrandControllerImpl{Repo: brandRepo}},
		UsersHandler:   handlers.UserHandlerImpl{UserCtrl: controllers.UserControllerImpl{Repo: userRepo}},
		AuthorsHandler: handlers.AuthorHandlerImpl{AuthorsCtrl: controllers.AuthorsControllerImpl{Repo: authorRepo}},
		LogsHandler: handlers.LogsHandlerImpl{Ctrl: controllers.LogControllerImpl{Repo :logsRepo}},
	}
	web.Listen()
}
