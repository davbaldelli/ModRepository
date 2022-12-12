package db

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarMods struct {
	ModModel
	ModelName    string `gorm:"column:model"`
	Year         uint
	Brand        string
	Categories   []CarCategory `gorm:"foreignKey:IdCar"`
	Transmission string
	Drivetrain   string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	Author       string
	AuthorLink   string
	Nation       string
	NationCode   string
	BrandLogo    string
}

type Car struct {
	ModModel
	ModelName    string `gorm:"column:model"`
	Year         int
	IdBrand      uint
	Categories   []CarCategory `gorm:"foreignKey:IdCar"`
	Transmission string
	Drivetrain   string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
}

type CarCategory struct {
	Id       uint `gorm:"primarykey"`
	Category string
	IdCar    uint
}

func (CarCategory) TableName() string {
	return "car_categories"
}

func (cat CarCategory) toEntity() entities.CarCategory {
	return entities.CarCategory{Name: entities.CarType(cat.Category)}
}

func (c CarMods) ToEntity(premium bool, admin bool) entities.Car {
	download := c.DownloadLink
	if (c.Premium && !premium) || (c.Personal && !admin) {
		download = c.Source
	}
	return entities.Car{
		Mod: entities.Mod{
			Id:           c.Id,
			DownloadLink: download,
			Source:       c.Source,
			Premium:      c.Premium,
			Personal:     c.Personal,
			Image:        c.Image,
			Author: entities.Author{
				Name: c.Author,
				Link: c.AuthorLink,
			},
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Rating:    c.Rating,
			Version:   c.Version,
			Official: c.Official,
		},
		Brand: entities.CarBrand{
			Name:   c.Brand,
			Nation: entities.Nation{Name: c.Nation, Code: c.NationCode},
		},
		ModelName: c.ModelName,
		Categories: mapCategories(c.Categories, func(category CarCategory) entities.CarCategory {
			return category.toEntity()
		}),
		Drivetrain:   entities.Drivetrain(c.Drivetrain),
		Transmission: entities.Transmission(c.Transmission),
		Year:         c.Year,
		Torque:       c.Torque,
		TopSpeed:     c.TopSpeed,
		Weight:       c.Weight,
		BHP:          c.BHP,
	}
}

func mapCategories(vs []CarCategory, f func(category CarCategory) entities.CarCategory) []entities.CarCategory {
	vsm := make([]entities.CarCategory, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func CarFromEntity(car entities.Car, idBrand uint, idAuthor uint) Car {
	return Car{
		ModModel:     ModModel{
			Id:           car.Id,
			Rating:       car.Rating,
			Version:      car.Version,
			DownloadLink: car.DownloadLink,
			Source:       car.Source,
			Premium:      car.Premium,
			Personal:     car.Personal,
			IdAuthor:     idAuthor,
			Image:        car.Image,
			Official: car.Official,
		},
		ModelName:    car.ModelName,
		IdBrand:      idBrand,
		Categories:   allCarCategoryFromEntity(car.Categories, car.Id),
		Year:         int(car.Year),
		Drivetrain:   string(car.Drivetrain),
		Transmission: string(car.Transmission),
		BHP:          car.BHP,
		Torque:       car.Torque,
		Weight:       car.Weight,
		TopSpeed:     car.TopSpeed,
	}
}

func carCategoryFromEntity(category entities.CarCategory, id uint) CarCategory {
	return CarCategory{
		IdCar:    id,
		Category: string(category.Name),
	}
}

func allCarCategoryFromEntity(categories []entities.CarCategory, id uint) []CarCategory {
	var dbCats []CarCategory
	for _, cat := range categories {
		dbCats = append(dbCats, carCategoryFromEntity(cat, id))
	}
	return dbCats
}
