// used to unmarshall the JSON data
// user send a request.. the data of the req will be in json.. so we unmarshall
// the data so that it is understood

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// in create function later, when we send the req to server
// the data will be in json format, we need to unmarshal it.. done here
func ParseBody(r *http.Request, x interface{}) {
	// reading the body
	if body, err := ioutil.ReadAll(r.Body); err == nil {

		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}

}

// explaination of above code:

/*
1. Read the body in the memory of the http req as a byte slice
2. if no err in reading the body, it unmarshals the body
3. unmarshal takes the byte slice and a pointer x interface
4. it tries to decode data of the json byte slice and store the result in x
5. value stored by x should match the structure value of the json data
6. if any err is found in unmarshalling, it returns without modifying x
*/
