package controller

import (
	"assigment3/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"text/template"
)

type Wind struct {
	Val int
}

type Water struct {
	Val int
}

func (w *Water) StatusWater() string {
	var Status string

	switch {
	case w.Val < 5:
		Status = "aman"
	case w.Val >= 6 && w.Val <= 8:
		Status = "siaga"
	case w.Val > 8:
		Status = "bahaya"
	}

	return Status
}

func (w *Wind) StatusWind() string {
	var Status string

	switch {
	case w.Val < 6:
		Status = "aman"
	case w.Val >= 7 && w.Val <= 15:
		Status = "siaga"
	case w.Val > 15:
		Status = "bahaya"
	}

	return Status
}

func UpdateWater(c chan int) {
	wa := rand.Intn(99) + 1
	fmt.Println("Water", wa)
	c <- wa
}

func UpdateWind(c chan int) {
	wi := rand.Intn(99) + 1
	fmt.Println("Wind", wi)
	c <- wi
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	Result := model.Value{}
	var waterVal int
	var windVal int
	c := make(chan int)

	go UpdateWater(c)
	waterVal = <-c

	go UpdateWind(c)
	windVal = <-c

	dataWater := Water{Val: waterVal}
	dataWind := Wind{Val: windVal}

	statWater := dataWater.StatusWater()
	statWind := dataWind.StatusWind()

	Result = model.Value{
		WaterValue:  waterVal,
		WindValue:   windVal,
		WaterStatus: statWater,
		WindStatus:  statWind,
	}

	_, err := json.Marshal(Result)
	if err != nil {
		panic(err)
	}

	if r.Method == "GET" {
		tpl, err := template.ParseFiles("./html/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tpl.Execute(w, Result)
		return
	}
	http.Error(w, "invalid method", http.StatusBadRequest)
}
