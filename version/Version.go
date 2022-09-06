package version

import "fmt"

var (
	Version = "dev"
	CommitHash = ""
	BuildTimestamp = ""
)

func BuildVersion() string {
	return fmt.Sprintf("%s-%s (%s)", Version, CommitHash, BuildTimestamp)
}