package solid

import (
	"io/ioutil"
	"net/http"
)

type DataReader func(source string) ([]byte, error)

func ReadFromFile(fileName string) ([]byte, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ReadFromLink(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return data, nil
}

func GetCities(reader DataReader, source string) ([]City, error) {
	data, err := reader(source)
	if err != nil {
		return nil, err
	}

	var cities []City
	err = yaml.Unmarshal(data, &cities)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
