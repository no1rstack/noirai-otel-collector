package metering

import "regexp"

var (
	ExcludeNoirAIWorkspaceResourceAttrs = regexp.MustCompile("^noirai.workspace.*")
)
