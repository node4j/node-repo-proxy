package npm

func GetMetaData(name string) ([]byte, error) {
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
