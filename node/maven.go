package node

import (
	"encoding/xml"
	"time"

	"github.com/srs/node-repo-proxy/maven"
)

func indexToMaven(index Index) ([]byte, error) {
	versions := make([]string, len(index))
	for i := range index {
		versions[i] = index[i].Version[1:]
	}

	m := maven.MetaData{}
	m.GroupId = "org.node"
	m.ArtifactId = "node"
	m.Versioning = maven.Versioning{}
	m.Versioning.Versions = maven.Versions{versions}
	m.Versioning.LastUpdated = uint64(time.Now().UnixNano() / int64(time.Millisecond))
	m.Versioning.Latest = versions[0]
	m.Versioning.Release = versions[0]

	data, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}
