package main

// GitCommit ... The git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

const (
	// Version ... The main version number that is being run at the moment.
	Version = "0.0.1"

	// VersionPrerelease ... A pre-release marker for the version. If this is "" (empty string)
	// then it means that it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "dev"
)
