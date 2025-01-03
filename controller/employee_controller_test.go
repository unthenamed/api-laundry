package controller

import (
	"api-laundry/service"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestObjEmployeeController(t *testing.T) {
	type args struct {
		rg      *gin.RouterGroup
		service service.EmployeeService
	}
	tests := []struct {
		name string
		args args
		want *EmployeeController
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ObjEmployeeController(tt.args.rg, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObjEmployeeController() = %v, want %v", got, tt.want)
			}
		})
	}
}
