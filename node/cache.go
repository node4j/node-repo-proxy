package node

import "github.com/srs/node-repo-proxy/util"

var MetaData []byte

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

	MetaData = data
	util.Log.Info("Updated node metadata")

	return nil
}

func UpdateMetaData() {
	LoadMetaData()
}
