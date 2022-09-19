package main

import (
	"net/http"


	"github.com/gin-gonic/gin"
)

const (
	carsUrl="/garage"
	carId = "/garage/:id"
)
var Garage []*Car

type Car struct {
	Id    string `json:"id"`
	Mark  string `json:"mark"`
	Model string `json:"model"`
	Gen   string `json:"gen"`
	VIN   string `json:"vin"`
}


var bmw = Car{
	Id:    "1",
	Mark:  "BMW",
	Model: "525",
	VIN:   "WBADT410MH112331",
	Gen:   "E39 2001-2003",
}
var honda = Car{
	Id:    "2",
	Mark:  "HONDA",
	Model: "FIT",
	VIN:   "JHMED63500S220141",
	Gen:   "GE 2007-2013",
}

func main() {
	Garage = append(Garage, &bmw, &honda)

	router := gin.Default()
	router.GET(carsUrl, ShowGarage)
	router.GET(carId,GetCarById)
	router.POST(carsUrl, AddCar)
	router.PATCH(carId, UpdateCar)
	router.DELETE(carId, DeleteCar)

	router.Run(":8080")

}

func ShowGarage(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, &Garage)

}

func AddCar(ctx *gin.Context) {//    СДЕЛАТЬ ЧТОБЫ ПУСТЫЕ ЗАПРОСЫ НЕ ОТПРАВЛЯЛИСЬ
	if ctx.Request.Method!=http.MethodPost{
		ctx.Writer.Header().Set("Allow","POST")

		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error":" BAD REQUEST"})
		return
	}
	
	var newCar Car
	
	ctx.BindJSON(&newCar)
	
	if newCar.Id == ""{
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error":"can't create with Id"})	
	return 
}
	Garage = append(Garage, &newCar)
	ctx.IndentedJSON(http.StatusCreated, &Garage)
}

func GetCarById(ctx *gin.Context) {
	id := ctx.Param("id")
	for _, car := range Garage {
		if id == car.Id {
			ctx.IndentedJSON(http.StatusOK, &car)
			return
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "NOT FOUND"})
}

func DeleteCar(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, car := range Garage {
		if id == car.Id {
			Garage = append(Garage[:i], Garage[i+1:]...)
			ctx.IndentedJSON(http.StatusNoContent, &car)
		}
	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "NOT FOUND"})
}

func UpdateCar(ctx *gin.Context) {
	id := ctx.Param("id")
	for i, car := range Garage {
		if id == car.Id {
			ctx.BindJSON(&car)
			Garage[i]= car
			ctx.IndentedJSON(http.StatusOK,&car)
		}

	}
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": "NOT FOUND"})

}
