package node

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type IndexEntry struct {
	Version string `json:"version"`
}

type Index []IndexEntry

func fetchNodeIndex() (Index, error) {
	resp, err := http.Get("https://nodejs.org/dist/index.json")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, nil
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	v := Index{}
	err = json.Unmarshal([]byte(body), &v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
