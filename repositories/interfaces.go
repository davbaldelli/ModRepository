package repositories

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarRepository interface {
	InsertCar(car *entities.Car) error
	SelectAllCars(premium bool, admin bool) ([]entities.Car, error)
	SelectAllCarCategories() ([]entities.CarCategory, error)
	UpdateCar(car entities.Car) (bool, error)
}

type LogRepository interface {
	SelectAllTrackLogs() ([]entities.TrackLog, error)
	SelectAllCarLogs() ([]entities.CarLog, error)
}

type TrackRepository interface {
	SelectAllTracks(premium bool, admin bool) ([]entities.Track, error)
	InsertTrack(track *entities.Track) error
	UpdateTrack(track entities.Track) (bool, error)
}

type NationRepository interface {
	SelectAllBrandsNations() ([]entities.Nation, error)
	SelectAllTrackNations() ([]entities.Nation, error)
}

type BrandRepository interface {
	SelectAllBrands() ([]entities.CarBrand, error)
}

type UserRepository interface {
	Login(user entities.User) (entities.User, error)
	SignIn(user entities.User) (entities.User, error)
	UpdatePassword(string, string) error
}

type AuthorRepository interface {
	SelectAllAuthors() ([]entities.Author, error)
	SelectAllCarAuthors() ([]entities.Author, error)
	SelectAllTrackAuthors() ([]entities.Author, error)
}

type ServersRepository interface {
	GetAllServers() ([]entities.Server, error)
	UpdateServer(server entities.Server) error
	AddServer(server entities.Server) error
	DeleteServer(server entities.Server) error
}

type SkinRepository interface {
	SelectCarSkins(carId uint) ([]entities.Skin, error)
	AddSkin(skin entities.Skin) error
	UpdateSkin(skin entities.Skin) error
}
