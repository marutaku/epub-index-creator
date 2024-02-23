package expand

type RootFile struct {
	FullPath  string `xml:"full-path,attr"`
	MediaType string `xml:"media-type,attr"`
}

type Container struct {
	XmlNS     string     `xml:"xmlns,attr"`
	Version   string     `xml:"version,attr"`
	RootFiles []RootFile `xml:"rootfiles>rootfile"`
}
