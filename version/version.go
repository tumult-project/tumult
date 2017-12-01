package version

var (
	// Version string filled by the compiler.
	Version string
)

// Info is a type that holds tumult version information.
type Info struct {
	SemVersion string
}

// VersionNumber gets the current version number.
func (i *Info) VersionNumber() string {
	return i.SemVersion
}

// GetVersion builds version information given information like the current
// semantic versioning, git status, platform and environment data.
//
// TODO:
// * Add git support
// * Add platform data like OS type
// * Add environment data like compiler/language version
func GetVersion() *Info {
	info := &Info{
		SemVersion: Version,
	}
	return info
}
