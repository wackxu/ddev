package version

import "fmt"

// VERSION is supplied with the git committish this is built from
var VERSION = ""

// IMPORTANT: These versions are overridden by version ldflags specifications VERSION_VARIABLES in the Makefile

// DdevVersion is the current version of ddev, by default the git committish (should be current git tag)
var DdevVersion = "v0.0.0-overridden-by-make" // Note that this is overridden by make

// DockerVersionConstraint is the current minimum version of docker required for ddev.
// See https://godoc.org/github.com/Masterminds/semver#hdr-Checking_Version_Constraints
// for examples defining version constraints.
// REMEMBER TO CHANGE docs/index.md if you touch this!
var DockerVersionConstraint = ">= 18.06.0-ce"

// DockerComposeVersionConstraint is the current minimum version of docker-compose required for ddev.
// REMEMBER TO CHANGE docs/index.md if you touch this!
var DockerComposeVersionConstraint = ">= 1.20.0"

// DockerComposeFileFormatVersion is the compose version to be used
var DockerComposeFileFormatVersion = "3.6"

// WebImg defines the default web image used for applications.
var WebImg = "drud/ddev-webserver"

// WebTag defines the default web image tag for drud dev
var WebTag = "20181022_add_ssh_auth" // Note that this can be overridden by make

// DBImg defines the default db image used for applications.
var DBImg = "drud/ddev-dbserver"

// DBTag defines the default db image tag for drud dev
var DBTag = "20181106_export_db" // Note that this may be overridden by make

// DBAImg defines the default phpmyadmin image tag used for applications.
var DBAImg = "drud/phpmyadmin"

// DBATag defines the default phpmyadmin image tag used for applications.
var DBATag = "v1.3.0" // Note that this can be overridden by make

// RouterImage defines the image used for the router.
var RouterImage = "drud/ddev-router"

// RouterTag defines the tag used for the router.
var RouterTag = "v1.3.0" // Note that this can be overridden by make

var SSHAuthImage = "drud/ddev-ssh-agent"

var SSHAuthTag = "20181022_add_ssh_auth"

// COMMIT is the actual committish, supplied by make
var COMMIT = "COMMIT should be overridden"

// BUILDINFO is information with date and context, supplied by make
var BUILDINFO = "BUILDINFO should have new info"

// DDevTLD defines the tld to use for DDev site URLs.
const DDevTLD = "ddev.local"

// GetVersionInfo returns a map containing the version info defined above.
func GetVersionInfo() map[string]string {
	versionInfo := make(map[string]string)

	versionInfo["cli"] = DdevVersion
	versionInfo["web"] = WebImg + ":" + WebTag
	versionInfo["db"] = DBImg + ":" + DBTag
	versionInfo["dba"] = DBAImg + ":" + DBATag
	versionInfo["router"] = RouterImage + ":" + RouterTag
	versionInfo["ddev-ssh-agent"] = SSHAuthImage + ":" + SSHAuthTag
	versionInfo["commit"] = COMMIT
	versionInfo["domain"] = DDevTLD
	versionInfo["build info"] = BUILDINFO

	return versionInfo
}

// GetWebImage returns the correctly formatted web image:tag reference
func GetWebImage() string {
	return fmt.Sprintf("%s:%s", WebImg, WebTag)
}

// GetDBImage returns the correctly formatted db image:tag reference
func GetDBImage() string {
	return fmt.Sprintf("%s:%s", DBImg, DBTag)
}

// GetDBAImage returns the correctly formatted dba image:tag reference
func GetDBAImage() string {
	return fmt.Sprintf("%s:%s", DBAImg, DBATag)
}
