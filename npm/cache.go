package npm

import "github.com/node4j/node-repo-proxy/util"

func GetMetaData(name string) ([]byte, error) {
	key := "npm." + name
	data, found := util.GetFromCache(key)
	if found {
		return data.([]byte), nil
	}

	data, err := loadMetaData(name)
	if err != nil {
		return nil, err
	}

	util.PutInCache(key, data)
	return data.([]byte), nil
}

func loadMetaData(name string) ([]byte, error) {
	index, err := fetchIndex(name)
	if err != nil {
		return nil, err
	}

	if index == nil {
		return nil, nil
	}

	data, err := indexToMaven(*index)
	if err != nil {
		return nil, err
	}

	return data, nil
}
