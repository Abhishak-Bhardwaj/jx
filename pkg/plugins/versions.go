package plugins

import (
	jenkinsv1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
)

const (
	// AdminVersion the version of the jx admin plugin
	AdminVersion = "0.0.31"

	// ApplicationVersion the version of the jx application plugin
	ApplicationVersion = "0.0.1"

	// GitOpsVersion the version of the jx gitops plugin
	GitOpsVersion = "0.0.54"

	// PipelineVersion the version of the jx pipeline plugin
	PipelineVersion = "0.0.2"

	// ProjectVersion the version of the jx project plugin
	ProjectVersion = "0.0.17"

	// PromoteVersion the version of the jx promote plugin
	PromoteVersion = "0.0.54"

	// SecretVersion the version of the jx secret plugin
	SecretVersion = "0.0.34"

	// VerifyVersion the version of the jx verify plugin
	VerifyVersion = "0.0.4"
)

var (
	// Plugins default plugins
	Plugins = []jenkinsv1.Plugin{
		CreateJXPlugin("admin", AdminVersion),
		CreateJXPlugin("application", ApplicationVersion),
		CreateJXPlugin("gitops", GitOpsVersion),
		CreateJXPlugin("pipeline", PipelineVersion),
		CreateJXPlugin("project", ProjectVersion),
		CreateJXPlugin("promote", PromoteVersion),
		CreateJXPlugin("secret", SecretVersion),
		CreateJXPlugin("verify", VerifyVersion),
	}
)
