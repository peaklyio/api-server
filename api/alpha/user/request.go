package user

import (
	"net/http"

	"github.com/kubicorn/kubicorn/pkg/logger"
	//mongo2 "github.com/peaklyio/api-server/mongo"

	"strings"

	"github.com/peaklyio/api-server/api"
	"github.com/peaklyio/api-server/mongo"
	"gopkg.in/mgo.v2/bson"
)

const c = "user"

func UserHandler(w http.ResponseWriter, r *http.Request) {
	logger.Debug("/user [%s]", r.Method)
	//logger.Debug("EasyStatus: JSON")
	s := api.NewEasyStatus(&api.JSONEncoder{})
	switch r.Method {

	// -----------------------------------------------------------------------------------------------------------------
	//
	// [GET]
	//
	//
	case "GET":
		userInput := User{}
		api.RequestToObject(r, &userInput)
		collection := mongo.GetCollection(c)
		result := User{}
		err := collection.Find(userInput).One(&result)
		if err != nil {
			s.Status404NotFound(w, r, api.E("Unable to find: %v", err))
			return
		}
		s.Status200Okay(w, r, result)
		return

	// -----------------------------------------------------------------------------------------------------------------
	//
	// [POST]
	//
	//
	case "POST":
		userInput := User{}
		api.RequestToObject(r, &userInput)
		collection := mongo.GetCollection(c)
		//logger.Info("%+v", userInput)
		var query *api.UniqQuery
		if userInput.Uniq != 0 {
			s.Status400BadRequest(w, r, api.E("Invalid field [Uniq] with POST request. Try PATCH request instead."))
			return
		} else if userInput.EmailAddress == "" {
			s.Status400BadRequest(w, r, api.E("Missing field [EmailAddress]"))
			return
		} else {
			query = api.UQuery("%s", userInput.EmailAddress)
		}

		// Ensure doesn't exist
		result := User{}
		err := collection.Find(query).One(&result)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				// Not found - let's save it
				_, err = collection.Upsert(query, bson.M{"$set": userInput})
				if err != nil {
					s.Status500InternalServerError(w, r, api.E("Unable to upsert: %v", err))
					return
				}
				s.Status200Okay(w, r, userInput)
				return
			}
		}
		s.Status400BadRequest(w, r, api.E("Email address [%s] already exists", userInput.EmailAddress))
		return
	// -----------------------------------------------------------------------------------------------------------------
	//
	// [PATCH]
	//
	//
	case "PATCH":
		userInput := User{}
		api.RequestToObject(r, &userInput)
		collection := mongo.GetCollection(c)
		//logger.Info("%+v", userInput)
		var query *api.UniqQuery
		if userInput.Uniq != 0 {
			query = &api.UniqQuery{
				Uniq: userInput.Uniq,
			}
		} else {
			s.Status400BadRequest(w, r, api.E("Missing field [Uniq]"))
			return
		}
		_, err := collection.Upsert(query, bson.M{"$set": userInput})
		if err != nil {
			s.Status500InternalServerError(w, r, api.E("Unable to upsert: %v", err))
			return
		}
		s.Status200Okay(w, r, userInput)
		return
	default:
		s.Status405MethodNotAllowed(w, r, api.E("Invalid method: %s", r.Method))
		return
	}
	s.Status404NotFound(w, r, api.E("Major error"))
	return
}
