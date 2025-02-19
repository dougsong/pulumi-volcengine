// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "volcengine:ecs/deploymentSet:DeploymentSet":
		r = &DeploymentSet{}
	case "volcengine:ecs/deploymentSetAssociate:DeploymentSetAssociate":
		r = &DeploymentSetAssociate{}
	case "volcengine:ecs/instance:Instance":
		r = &Instance{}
	case "volcengine:ecs/keyPair:KeyPair":
		r = &KeyPair{}
	case "volcengine:ecs/keyPairAssociate:KeyPairAssociate":
		r = &KeyPairAssociate{}
	case "volcengine:ecs/launchTemplate:LaunchTemplate":
		r = &LaunchTemplate{}
	case "volcengine:ecs/state:State":
		r = &State{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := volcengine.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/deploymentSet",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/deploymentSetAssociate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/instance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/keyPair",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/keyPairAssociate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/launchTemplate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"ecs/state",
		&module{version},
	)
}
