package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pb3/db"
	"pb3/entity"
	"strconv"

	"github.com/sirupsen/logrus"
)

func PostVehicle(rw http.ResponseWriter, r *http.Request) {

	reqBody := r.Body

	bodyBytes, err := ioutil.ReadAll(reqBody)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	var vehicle entity.Vehicle
	err = json.Unmarshal(bodyBytes, &vehicle)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	db.GetDB().Create(&vehicle)

	fmt.Println(vehicle)
	rw.Write(bodyBytes)
}

func GetVehicle(rw http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("id")

	var vehicle entity.Vehicle

	result := db.GetDB().Where("id=?", name).Find(&vehicle)

	if result.RecordNotFound() {
		http.Error(rw, "No record", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
	}

	vehicleBytes, err := json.Marshal(vehicle)

	if hasError(rw, err, "Internal Issue") {
		return
	}

	rw.Write(vehicleBytes)
}

func GetCanDrive(rw http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("id")

	var vehicle entity.Vehicle

	result := db.GetDB().Where("id=?", name).Find(&vehicle)

	if result.RecordNotFound() {
		http.Error(rw, "No record", http.StatusInternalServerError)
		return
	}

	if result.Error != nil {
		http.Error(rw, result.Error.Error(), http.StatusInternalServerError)
	}

	canDriveValue := strconv.FormatBool(vehicle.CanDrive)

	rw.Write([]byte(canDriveValue))
}

func DeleteVehicle(rw http.ResponseWriter, r *http.Request) {

	var vehicle entity.Vehicle

	idVal := r.URL.Query().Get("id")
	db.GetDB().Where("id=?", idVal).Delete(&vehicle)
}

func hasError(rw http.ResponseWriter, err error, message string) bool {
	logger := new(logrus.Entry)

	if err != nil {
		logger.WithError(err).Error(message)
		rw.Write([]byte(message))
		return true
	}

	return false
}
