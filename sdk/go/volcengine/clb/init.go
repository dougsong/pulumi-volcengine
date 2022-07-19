// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "volcengine:Clb/certificate:Certificate":
		r = &Certificate{}
	case "volcengine:Clb/clb:Clb":
		r = &Clb{}
	case "volcengine:Clb/clbRule:ClbRule":
		r = &ClbRule{}
	case "volcengine:Clb/listener:Listener":
		r = &Listener{}
	case "volcengine:Clb/serverGroup:ServerGroup":
		r = &ServerGroup{}
	case "volcengine:Clb/serverGroupServer:ServerGroupServer":
		r = &ServerGroupServer{}
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
		"Clb/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"Clb/clb",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"Clb/clbRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"Clb/listener",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"Clb/serverGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"volcengine",
		"Clb/serverGroupServer",
		&module{version},
	)
}
