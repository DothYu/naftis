// Package version provides build version information.
package version

import (
	"fmt"
	"runtime"
)

// 在构建时，使用-ldflags -X填充以下字段。
// Note that DATE is omitted for reproducible builds
var (
	buildVersion     = "unknown"
	buildGitRevision = "unknown"
	buildUser        = "unknown"
	buildHost        = "unknown"
	buildStatus      = "unknown"
	buildTime        = "unknown"
)

// BuildInfo描述有关二进制版本的版本信息。
type BuildInfo struct {
	Version       string `json:"version"`
	GitRevision   string `json:"revision"`
	User          string `json:"user"`
	Host          string `json:"host"`
	GolangVersion string `json:"golang_version"`
	BuildStatus   string `json:"status"`
	BuildTime     string `json:"time"`
}

var (
	Info BuildInfo
)

// 生成单行版本信息
//
// This looks like:
//
// ```
// user@host-<version>-<git revision>-<build status>
// ```
func (b BuildInfo) String() string {
	return fmt.Sprintf("%v@%v-%v-%v-%v-%v",
		b.User,
		b.Host,
		b.Version,
		b.GitRevision,
		b.BuildStatus,
		b.BuildTime)
}

// 生成多行版本信息
//
// This looks like:
//
// ```
// Version: <version>
// GitRevision: <git revision>
// User: user@host
// GolangVersion: go1.10.2
// BuildStatus: <build status>
// ```
func (b BuildInfo) LongForm() string {
	return fmt.Sprintf(`Version: %v
GitRevision: %v
User: %v@%v
GolangVersion: %v
BuildStatus: %v
BuildTime: %v
`,
		b.Version,
		b.GitRevision,
		b.User,
		b.Host,
		b.GolangVersion,
		b.BuildStatus,
		b.BuildTime)
}

/**
 * description: version包初始化
 */
func init() {
	Info = BuildInfo{
		Version:       buildVersion,
		GitRevision:   buildGitRevision,
		User:          buildUser,
		Host:          buildHost,
		GolangVersion: runtime.Version(),
		BuildStatus:   buildStatus,
		BuildTime:     buildTime,
	}
}
