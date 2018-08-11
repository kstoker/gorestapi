package mededelingen

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	config "gorestapi/config"
	"gorestapi/database"
	logger "gorestapi/logger"
	auth "gorestapi/middleware"
	utils "gorestapi/utils"
)

// TableName : mededelingen table name to be `announcement_announcement`
func (Mededeling) TableName() string {
	return "announcement_announcement"
}

// AddToRouter : set routes and handlers
func AddToRouter(r *mux.Router) {
	r.HandleFunc("/api/mededelingen", auth.AuthMiddleware(GetMededelingen)).Methods("GET")
	r.HandleFunc("/api/mededeling/{id}", auth.AuthMiddleware(GetMededeling)).Methods("GET")
	r.HandleFunc("/api/mededeling", auth.AuthMiddleware(CreateMededeling)).Methods("POST")
	r.HandleFunc("/api/mededeling/{id}", auth.AuthMiddleware(UpdateMededeling)).Methods("PUT")
	r.HandleFunc("/api/mededeling/{id}", auth.AuthMiddleware(DeleteMededeling)).Methods("DELETE")
}

// GetMededelingen : get all records
func GetMededelingen(w http.ResponseWriter, r *http.Request) {
	var mededelingen []Mededeling
	db := database.GetDB()
	db.Find(&mededelingen)
	if err := json.NewEncoder(w).Encode(mededelingen); err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Could not write JSON.")
	}
}

// GetMededeling : get a single record based on id
func GetMededeling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var mededeling Mededeling
	db := database.GetDB()
	db.First(&mededeling, id)
	if mededeling.ID != 0 {
		// found
		if err := json.NewEncoder(w).Encode(mededeling); err != nil {
			logger.Log.Println(err)
			utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Could not write JSON.")
		}
	} else {
		// not found
		utils.WriteJSONMessage(w, http.StatusNotFound, "Record not found.")
	}
}

// CreateMededeling : create a new record
func CreateMededeling(w http.ResponseWriter, r *http.Request) {
	var mededeling Mededeling
	// CreateBodyLimit is ingesteld op 1 Mb als maximale omvang van de body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, config.CreateBodyLimit))
	if err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusBadRequest, "Error occured reading body.")
		return
	}
	if err := r.Body.Close(); err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusBadRequest, "Error occured closing body.")
		return
	}
	if err := json.Unmarshal(body, &mededeling); err != nil {
		logger.Log.Println(err.Error())
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Cannot process body.")
		return
	}
	db := database.GetDB()
	res := db.Create(&mededeling)
	if res.Error != nil {
		logger.Log.Println(res.Error)
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Could not write JSON.")
		return
	}
	//	let op: als de create slaagt dan is ID in mededeling gevuld met de waarde die door de database is toegekend
	// hier geven we alle velden terug inclusief de gevulde ID
	if err := json.NewEncoder(w).Encode(mededeling); err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Could not write JSON.")
		return
	}
}

// UpdateMededeling : update single record based on id
func UpdateMededeling(w http.ResponseWriter, r *http.Request) {
	var mededeling Mededeling
	// CreateBodyLimit is ingesteld op 1 Mb als maximale omvang van de body
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, config.CreateBodyLimit))
	if err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusBadRequest, "Error occured reading body.")
		return
	}
	if err := r.Body.Close(); err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusBadRequest, "Error occured closing body.")
		return
	}
	if err := json.Unmarshal(body, &mededeling); err != nil {
		logger.Log.Println(err.Error())
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Cannot process body.")
		return
	}
	// at this point we do have a valid body, read it and put into a map
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(body), &data); err != nil {
		logger.Log.Println(err.Error())
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Cannot process body.")
		return
	}
	// get current record based on id in request
	vars := mux.Vars(r)
	id := vars["id"]
	db := database.GetDB()
	db.First(&mededeling, id)
	if mededeling.ID != 0 {
		// found
		if err := json.NewEncoder(w).Encode(mededeling); err != nil {
			logger.Log.Println(err)
			utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Cannot write JSON.")
			return
		}
	} else {
		// not found
		utils.WriteJSONMessage(w, http.StatusNotFound, "Record to be updated with given ID not found.")
		return
	}

	// now loop through the map to update given fields
	for k, v := range data {
		switch k {
		case "text":
			mededeling.Text = v.(string)
		case "datefrom":
			t, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				logger.Log.Println(err)
			} else {
				mededeling.DateFrom = t
			}
		case "dateto":
			t, err := time.Parse(time.RFC3339, v.(string))
			if err != nil {
				logger.Log.Println(err)
			} else {
				mededeling.DateTo = t
			}
		default:
			// error: should not happen
			logger.Log.Println("Mededeling update:", "k:", k, "v:", v)
		}
	}
	// values in the struct have been updated, now save to database
	res := db.Save(&mededeling)
	if res.Error != nil {
		logger.Log.Println(res.Error)
		utils.WriteJSONMessage(w, http.StatusNotModified, "Error occured on update.")
		return
	}

	// return updated struct back to requester
	if err := json.NewEncoder(w).Encode(mededeling); err != nil {
		logger.Log.Println(err)
		utils.WriteJSONMessage(w, http.StatusUnprocessableEntity, "Error writing JSON.")
	}
}

// DeleteMededeling : delete a single record based on id
func DeleteMededeling(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var mededeling Mededeling
	db := database.GetDB()
	db.First(&mededeling, id)
	if mededeling.ID != 0 {
		// found
		res := db.Delete(&mededeling)
		if res.Error != nil {
			logger.Log.Println(res.Error)
			utils.WriteJSONMessage(w, http.StatusNotModified, "Error occured trying to delete record.")
		} else {
			utils.WriteJSONMessage(w, http.StatusOK, "Record deleted.")
		}
	} else {
		// not found
		utils.WriteJSONMessage(w, http.StatusNotFound, "Record with given id not found.")
	}
}
