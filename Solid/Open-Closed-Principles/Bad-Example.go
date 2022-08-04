package solid

import (
	"io/ioutil"
	"net/http"
)

func GetCities(sourceType string, source string) ([]City, error) {
	var data []byte
	var err error

	if sourceType == "file" {
		data, err = ioutil.ReadFile(source)
		if err != nil {
			return nil, err
		}
	} else if sourceType == "link" {
		resp, err := http.Get(source)
		if err != nil {
			return nil, err
		}

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
	}

	var cities []City
	err = yaml.Unmarshal(data, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
