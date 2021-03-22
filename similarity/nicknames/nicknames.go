package nicknames

import (
	"encoding/json"
	"io/ioutil"
)

type nicknameDetails struct {
	Name      string   `json:"name"`
	Nicknames []string `json:"nicknames"`
}

var nicknameList []nicknameDetails

func contains(s []string, n string) bool {
	for _, a := range s {
		if a == n {
			return true
		}
	}
	return false
}

// JSON file with each record being of form
//{"name":xxx, "nicknames":["nick1","nick2", ....]}
func loadNamesFromFile(filename string) (map[string]([]string), error) {

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal([]byte(file), &nicknameList); err != nil {
		return nil, err
	}

	return list, nil
}

// FindNicknames . returns set of known nicknames for an given name
func FindNicknames(name, nickname string) bool {

	// name and nick name must contain something
	if len(name) == 0 || len(nickname) == 0 {
		return false
	}

	// a nickname can be the actual name
	if name == nickname {
		return true
	}

	nicknames := map[string][]string{}

	var err error
	if nicknames, err = loadNamesFromFile("test/nicknames.txt"); err != nil {
		return false // empty list
	}

	// no nicknames defined for this name
	if len(nicknames) == 0 {
		return false
	}

	//determine whether nickname is defined for the full name provided
	return contains(nicknames[name], nickname)
}
