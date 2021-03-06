package courses

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func logErrorWithoutFailing(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func errorHandler(formatter *render.Render, w http.ResponseWriter, err error) {
	switch err.(type) {
	case *BadRequestError:
		formatter.JSON(w, http.StatusBadRequest, struct {
			Success bool
			Message string
		}{
			false,
			err.Error(),
		})
		return
	case *InternalServerError:
		formatter.JSON(w, http.StatusInternalServerError, struct {
			Success bool
			Message string
		}{
			false,
			err.Error(),
		})
		return
	case *EntityNotFoundError:
		formatter.JSON(w, http.StatusNotFound, struct {
			Success bool
			Message string
		}{
			false,
			err.Error(),
		})
	default:
		log.Printf("Internal Server Error: %s", err)
		formatter.JSON(w, http.StatusInternalServerError, struct {
			Success bool
			Message string
		}{
			false,
			"Internal Server Error",
		})
	}
}

//NewRouter returns a new mux router for courses api
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	for _, route := range routes {
		var handler http.Handler
		handler = route.GetFormattedHandlerFunc(formatter)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	fmt.Println("MONGO_URL:", mongoURL)
	fmt.Println("DATABASE:", database)
	fmt.Println("COLLECTION:", collection)
	fmt.Println("KAFKA_SERVER", kafkaServer)
	fmt.Println("COURSE_CLICK_TOPIC", kafkaClickTopic)
	fmt.Println("ENROLLMENT_TOPIC", kafkaEnrollmentTopic)
	// session, err := mgo.Dial(mongoURL)
	// failOnError(err, "Mongo Dial Error")
	// defer session.Close()
	return router
}

var routes = Routes{
	Route{
		"PingHandler",
		"GET",
		"/ping",
		PingHandler,
	},
	Route{
		"CreateCourseHandler",
		"POST",
		"/courses",
		CreateCourseHandler,
	},
	Route{
		"GetCoursesHandler",
		"GET",
		"/courses",
		GetCoursesHandler,
	},
	Route{
		"GetCourseHandler",
		"GET",
		"/courses/{CourseID}",
		GetCourseHandler,
	},
	Route{
		"UpdateCourseHandler",
		"PUT",
		"/courses/{CourseID}",
		UpdateCourseHandler,
	},
	Route{
		"DeleteCourseHandler",
		"DELETE",
		"/courses/{CourseID}",
		DeleteCourseHandler,
	},
}
