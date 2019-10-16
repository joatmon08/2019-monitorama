package mock

import "io/ioutil"

type Wiki struct {
	BaseURL string
}

func (w *Wiki) GetDocument(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
