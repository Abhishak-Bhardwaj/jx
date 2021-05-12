package upgrade

import (
	"github.com/jenkins-x/jx-cli/pkg/plugins"
	"github.com/jenkins-x/jx-helpers/pkg/cobras/helper"
	"github.com/jenkins-x/jx-helpers/pkg/cobras/templates"
	"github.com/jenkins-x/jx-helpers/pkg/extensions"
	"github.com/jenkins-x/jx-helpers/pkg/termcolor"
	"github.com/jenkins-x/jx-logging/pkg/log"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	cmdLong = templates.LongDesc(`
		Upgrades all of the plugins in your local Jenkins X CLI
`)

	cmdExample = templates.Examples(`
		# upgrades your plugin binaries
		jx upgrade
	`)
)

// UpgradeOptions the options for upgrading a cluster
type UpgradeOptions struct {
	GitCredentials bool
}

// NewCmdUpgrade creates a command object for the command
func NewCmdUpgrade() (*cobra.Command, *UpgradeOptions) {
	o := &UpgradeOptions{}

	cmd := &cobra.Command{
		Use:     "upgrade",
		Short:   "Upgrades all of the plugins in your local Jenkins X CLI",
		Long:    cmdLong,
		Example: cmdExample,
		Run: func(cmd *cobra.Command, args []string) {
			err := o.Run()
			helper.CheckErr(err)
		},
	}
	return cmd, o
}

// Run implements the command
func (o *UpgradeOptions) Run() error {
	pluginBinDir, err := plugins.PluginBinDir()
	if err != nil {
		return errors.Wrap(err, "failed to find plugin bin directory")
	}

	for _, p := range plugins.Plugins {
		log.Logger().Infof("checking binary jx plugin %s version %s is installed", termcolor.ColorInfo(p.Name), termcolor.ColorInfo(p.Spec.Version))
		_, err := extensions.EnsurePluginInstalled(p, pluginBinDir)
		if err != nil {
			return errors.Wrapf(err, "failed to ensure plugin is installed %s", p.Name)
		}
	}
	return nil
}

/*
// upgradeBinaryPlugins eagerly installs/upgrades any binary plugins which have Plugin CRDs defined
// in the current development namespace
func (o *UpgradeOptions) upgradeBinaryPlugins() error {
	jxClient, ns, err := o.JXClientAndDevNamespace()
	if err != nil {
		return errors.Wrapf(err, "failed to create jx client")
	}
	pluginList, err := jxClient.JenkinsV1().Plugins(ns).List(metav1.ListOptions{})
	if err != nil && apierrors.IsNotFound(err) {
		err = nil
	}
	if err != nil {
		return errors.Wrapf(err, "failed to query Jenkins X plugins in namespace %s", ns)
	}
	if pluginList != nil {
		for _, plugin := range pluginList.Items {
			if plugin.Labels != nil && plugin.Labels[extensions.PluginCommandLabel] != "" {
				log.Logger().Infof("checking binary jx plugin %s version %s is installed", util.ColorInfo(plugin.Name), util.ColorInfo(plugin.Spec.Version))
				_, err = extensions.EnsurePluginInstalled(plugin)
				if err != nil {
					return errors.Wrapf(err, "failed to ensure plugin is installed %s", plugin.Name)
				}
			}
		}
	}
	return nil
}
*/
