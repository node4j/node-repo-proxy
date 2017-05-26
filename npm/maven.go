package npm

import (
	"encoding/xml"
	"sort"
	"time"

	"github.com/srs/node-repo-proxy/maven"
)

func indexToMaven(index Index) ([]byte, error) {
	versions := make([]string, len(index.Versions))
	i := 0

	for _, v := range index.Versions {
		versions[i] = v.Version
		i++
	}

	sort.Sort(sort.Reverse(sort.StringSlice(versions)))

	m := maven.MetaData{}
	m.GroupId = "org.npmjs"
	m.ArtifactId = index.Name
	m.Versioning = maven.Versioning{}
	m.Versioning.Versions = maven.Versions{versions}
	m.Versioning.LastUpdated = uint64(time.Now().UnixNano() / int64(time.Millisecond))
	m.Versioning.Latest = index.DistTags.Latest
	m.Versioning.Release = index.DistTags.Latest

	data, err := xml.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, err
	}

	return data, nil
}
