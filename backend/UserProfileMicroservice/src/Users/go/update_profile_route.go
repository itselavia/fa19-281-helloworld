package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//UpdateProfileHandler returns a Handler for UpdateProfile Request
func UpdateProfileHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		userID, err := strconv.Atoi(params["UserID"])
		if err != nil {
			err = NewBadRequestError("UserID must be integer")
			errorHandler(formatter, w, err)
			return
		}
		var user User
		_ = json.NewDecoder(req.Body).Decode(&user)
		user.UserID = userID
		err = updateProfile(userID, user)
		if err != nil {
			if err == mgo.ErrNotFound {
				err = NewEntityNotFoundError(strconv.Itoa(userID))
			}
			errorHandler(formatter, w, err)
			return
		}
		formatter.JSON(w, http.StatusOK, user)
	}
}

func updateProfile(userID int, user User) error {
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		logErrorWithoutFailing(err, "Mongo Dial Error")
		return err
	}
	defer session.Close()
	c := session.DB(database).C(collection)
	if err := c.Update(bson.M{
		"userid": userID,
	}, user); err != nil {
		logErrorWithoutFailing(err, "Mongo Update Error")
		return err
	}
	return nil
}
