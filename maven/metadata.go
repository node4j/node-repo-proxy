package maven

type Versions struct {
	Version []string `xml:"version"`
}

type Versioning struct {
	Latest      string   `xml:"latest"`
	Release     string   `xml:"release"`
	Versions    Versions `xml:"versions"`
	LastUpdated uint64   `xml:"lastUpdated"`
}

type MetaData struct {
	XMLName    struct{}   `xml:"metadata"`
	GroupId    string     `xml:"groupId"`
	ArtifactId string     `xml:"artifactId"`
	Versioning Versioning `xml:"versioning"`
}
