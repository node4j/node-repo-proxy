package npm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DistTags struct {
	Latest string `json:"latest"`
}

type Version struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Versions map[string]Version

type Index struct {
	Name     string   `json:"name"`
	DistTags DistTags `json:"dist-tags"`
	Versions Versions `json:"versions"`
}

func fetchIndex(name string) (*Index, error) {
	resp, err := http.Get("https://registry.npmjs.org/" + name)
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

	return &v, nil
}
