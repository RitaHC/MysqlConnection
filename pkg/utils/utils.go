// Maeshal / Unmarshal Json data that is recieved from the requests

package utils

import(
	// Json conversion Packages
	"encoding/json"
	"io/ioutil"
	// Server Connection Package
	"net/http"
)

// 2 Parameters: request from web & interface
func ParseBody(r *http.Request, x interface{}){
	// Read the whole body or error using the ioutil package 
	// Then assgn the whole result in variable 'Body' or error in variable 'err'
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		// If NO ERROR: then unmarshal it in the variable 'x'
		// Any error encountered is assigned to 'err'
		// The 'if' statement checks that the result is not 'nil' then retur the result
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return 
		}
	}
}