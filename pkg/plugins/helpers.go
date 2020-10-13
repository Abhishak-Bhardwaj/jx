package plugins

import (
	"fmt"
	"strings"

	jenkinsv1 "github.com/jenkins-x/jx-api/v3/pkg/apis/jenkins.io/v1"
	"github.com/jenkins-x/jx-helpers/v3/pkg/extensions"
	"github.com/jenkins-x/jx-helpers/v3/pkg/homedir"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	jenkinsxOrganisation        = "jenkins-x"
	jenkinsxPluginsOrganisation = "jenkins-x-plugins"
)

// GetJXPlugin returns the path to the locally installed jx plugin
func GetJXPlugin(name, version string) (string, error) {
	pluginBinDir, err := homedir.DefaultPluginBinDir()
	if err != nil {
		return "", err
	}
	plugin := CreateJXPlugin(jenkinsxOrganisation, name, version)
	return extensions.EnsurePluginInstalled(plugin, pluginBinDir)
}

// CreateJXPlugin creates the jx plugin
func CreateJXPlugin(org, name, version string) jenkinsv1.Plugin {
	binaries := extensions.CreateBinaries(func(p extensions.Platform) string {
		return fmt.Sprintf("https://github.com/%s/jx-%s/releases/download/v%s/jx-%s-%s-%s.%s", org, name, version, name, strings.ToLower(p.Goos), strings.ToLower(p.Goarch), p.Extension())
	})

	plugin := jenkinsv1.Plugin{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: jenkinsv1.PluginSpec{
			SubCommand:  name,
			Binaries:    binaries,
			Description: name + "  binary",
			Name:        "jx-" + name,
			Version:     version,
		},
	}
	return plugin
}
