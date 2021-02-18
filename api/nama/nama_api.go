package nama

import (
	"config"
	"encoding/json"
	"models"
	"entities"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)


func FindAll(response http.ResponseWriter, request *http.Request){
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		namaModel := models.NamaModel{
			Db : db,
		}
		namas, err2 := namaModel.FindAll()
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, namas)
		}
	}
}

func Search(response http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	keyword := vars["keyword"]
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		namaModel := models.NamaModel{
			Db : db,
		}
		namas, err2 := namaModel.Search(keyword)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, namas)
		}
	}
}

func Create(response http.ResponseWriter, request *http.Request){
	var nama entities.Nama
	err := json.NewDecoder(request.Body).Decode(&nama)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		namaModel := models.NamaModel{
			Db : db,
		}
		err2 := namaModel.Create(&nama)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nama)
		}
	}
}

func Update(response http.ResponseWriter, request *http.Request){
	var nama entities.Nama
	err := json.NewDecoder(request.Body).Decode(&nama)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		namaModel := models.NamaModel{
			Db : db,
		}
		_, err2 := namaModel.Update(&nama)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nama)
		}
	}
}

func Delete(response http.ResponseWriter, request *http.Request){
	vars := mux.Vars(request)
	sid := vars["id"]
	id , _ := strconv.ParseInt(sid,10,64)
	db, err := config.GetDB()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
	} else {
		namaModel := models.NamaModel{
			Db : db,
		}
		_, err2 := namaModel.Delete(id)
		if err2 != nil {
			respondWithError(response, http.StatusBadRequest, err2.Error())
		} else {
			respondWithJson(response, http.StatusOK, nil)
		}
	}
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJson(w, code, map[string]string{"error":msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}