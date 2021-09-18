package constants

// VersionTypes contains the different types of versions.
type VersionTypes struct {
	AppVersion string
}

// Versions contains the different kinds of versions and their values.
var Versions = VersionTypes{
	AppVersion: "0.0.1",
}
