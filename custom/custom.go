package custom

import (
	"fmt"
	"log"
	"net/http"
)

func Show(msg string, n int) {
	dots := ""

	for i := 0; i < n; i++ {
		dots += ":"
	}

	m := fmt.Sprintf("%v %v %v", dots, msg, dots)
	fmt.Println(m)
}

func AppLog(r *http.Request){
	log.Println("::::",r.URL, r.Method,)
}

func GenericError() map[string]interface{} {
	return map[string]interface{}{
		"Message": "Something went wrong",
	}
}