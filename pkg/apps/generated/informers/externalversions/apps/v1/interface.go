// This file was automatically generated by informer-gen

package v1

import (
	internalinterfaces "github.com/openshift/origin/pkg/apps/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// DeploymentConfigs returns a DeploymentConfigInformer.
	DeploymentConfigs() DeploymentConfigInformer
}

type version struct {
	internalinterfaces.SharedInformerFactory
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory) Interface {
	return &version{f}
}

// DeploymentConfigs returns a DeploymentConfigInformer.
func (v *version) DeploymentConfigs() DeploymentConfigInformer {
	return &deploymentConfigInformer{factory: v.SharedInformerFactory}
}
