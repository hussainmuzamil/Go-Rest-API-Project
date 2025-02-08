package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/hussainmuzamil/students-api/internal/types"
	"github.com/hussainmuzamil/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating new student")

		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			slog.Error("Request Body is Empty")
			_ = response.WriteJson(w, http.StatusBadRequest, response.ErrorResponse(fmt.Errorf("request body is Empty")))
			return
		}

		if(err != nil){
			response.WriteJson(w,http.StatusBadRequest,response.ErrorResponse(err))
			return
		}
		
		if err := validator.New().Struct(student); err != nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationErrors(validateErrs))
			return
		}

		_ = response.WriteJson(w, http.StatusCreated, map[string]string{"message": "Student created successfully"})
	}
}
