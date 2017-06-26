package node

import "github.com/node4j/node-repo-proxy/util"

func GetMetaData() ([]byte, error) {
	data, found := util.GetFromCache("node")
	if found {
		return data.([]byte), nil
	}

	data, err := loadMetaData()
	if err != nil {
		return nil, err
	}

	util.PutInCache("node", data)
	return data.([]byte), nil
}

func loadMetaData() ([]byte, error) {
	index, err := fetchNodeIndex()
	if err != nil {
		return nil, err
	}

	data, err := indexToMaven(index)
	if err != nil {
		return nil, err
	}

	return data, nil
}
