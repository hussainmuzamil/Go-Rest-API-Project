package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hussainmuzamil/students-api/internal/types"
	"github.com/hussainmuzamil/students-api/internal/utils/response"
	"io"
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating new student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			slog.Error("Request Body is Empty", err.Error())
			_ = response.WriteJson(w, http.StatusBadRequest, response.ErrorResponse(fmt.Errorf("request body is Empty")))
			return
		}

		if(err != nil){
			response.WriteJson(w,http.StatusBadRequest,response.ErrorResponse(err))
			return
		}
		
		if err := validator.New().Struct(student); err != nil{
			response.WriteJson(w,http.StatusBadRequest,response.Error(err))
			return
		}

		slog.Info("Student Request With Request Body",student)



		_ = response.WriteJson(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})
	}
}
