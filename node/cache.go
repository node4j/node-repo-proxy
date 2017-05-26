package node

import "github.com/srs/node-repo-proxy/util"

var metaData []byte

func GetMetaData() []byte {
	return metaData
}

func LoadMetaData() error {
	util.Log.Info("Loading node metadata")

	index, err := fetchNodeIndex()
	if err != nil {
		return err
	}

	data, err := indexToMaven(index)
	if err != nil {
		return err
	}

	metaData = data
	util.Log.Info("Updated node metadata")

	return nil
}

func UpdateMetaData() {
	LoadMetaData()
}
