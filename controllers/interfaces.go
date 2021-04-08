package controllers

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarController interface {
	GetAllCars() []entities.Car
	GetCarsByNation(string) []entities.Car
	GetCarByModel(string) []entities.Car
	GetCarsByBrand(string) []entities.Car
	GetCarsByType(string) []entities.Car

	AddCar(modelName string, downloadUrl string, brand entities.CarBrand, categories []entities.CarCategory, year uint) error
}

type TrackController interface {
	GetAllTracks() []entities.Track
	GetTracksByNation(string) []entities.Track
	GetTracksByLayoutType(string) []entities.Track
	GetTracksByName(string) []entities.Track

	AddTrack(name string, downloadUrl string, layouts []entities.Layout, location string, nation entities.Nation, year uint) error
}

type BrandController interface {
	GetAllBrands() []entities.CarBrand
	GetBrandByNation(string) []entities.CarBrand
	GetBrandByName(string) []entities.CarBrand
}

type NationController interface {
	GetAllTracksNations() []entities.Nation
	GetAllBrandsNations() []entities.Nation
}
