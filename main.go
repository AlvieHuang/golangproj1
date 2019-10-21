package golangproj1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name       string `json: "name"`
	Age        int    `json: "age"`
	Profession string `json: "profession"`
	HairColor  string `json: "hairColor"`
}

var personmap = make(map[string]Person)

func main() {
	http.HandleFunc("/people", personFunc)
	http.HandleFunc("/people/", peopleFunc)
	http.ListenAndServe(":8080", nil)
}

func peopleFunc(w http.ResponseWriter, req *http.Request) {
	var name = req.URL.Path[8:]
	val, found := personmap[name]
	if found {
		toPrint, _ := json.Marshal(val)
		fmt.Fprintf(w, "%s", toPrint)
	} else {
		fmt.Fprint(w, "Name does not exist in map")
	}

}
func personFunc(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		// Read body
		b, err := ioutil.ReadAll(req.Body)
		defer req.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Unmarshal
		var msg Person
		err = json.Unmarshal(b, &msg)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		var name = msg.Name
		personmap[name] = msg
		toPrint, err := json.Marshal(msg)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%s", toPrint)
		// personmap[msg.name] = msg

	default:
		var arr []Person
		for _, value := range personmap {
			arr = append(arr, value)
		}
		var toPrint, err = json.Marshal(arr)
		if err != nil {
			return
		}

		fmt.Fprintf(w, "%s", toPrint)

	}
}

// Using the previous code bits shown, develop an HTTP server that will:

// Consume Person structs via a POST request to /people
// Save the data in a map
// Retrieve and return the marshaled json for the appropriate person when firing a GET request against /people/{name}
// Write the data out to a file and return the marshaled json for all people with a GET request to /people
