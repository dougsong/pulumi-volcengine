// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bioos

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage bioos cluster bind
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/bioos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bioos.NewClusterBind(ctx, "example", &bioos.ClusterBindArgs{
//				ClusterId:   pulumi.String("ucfhp1nteig48u8ufv8s0"),
//				Type:        pulumi.String("workflow"),
//				WorkspaceId: pulumi.String("wcfhp1vdeig48u8ufv8sg"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Cluster binder can be imported using the workspace id and cluster id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:bioos/clusterBind:ClusterBind default wc*****:uc***
//
// ```
type ClusterBind struct {
	pulumi.CustomResourceState

	// The id of the cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The type of the cluster bind.
	Type pulumi.StringOutput `pulumi:"type"`
	// The id of the workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewClusterBind registers a new resource with the given unique name, arguments, and options.
func NewClusterBind(ctx *pulumi.Context,
	name string, args *ClusterBindArgs, opts ...pulumi.ResourceOption) (*ClusterBind, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.WorkspaceId == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterBind
	err := ctx.RegisterResource("volcengine:bioos/clusterBind:ClusterBind", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClusterBind gets an existing ClusterBind resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClusterBind(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClusterBindState, opts ...pulumi.ResourceOption) (*ClusterBind, error) {
	var resource ClusterBind
	err := ctx.ReadResource("volcengine:bioos/clusterBind:ClusterBind", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ClusterBind resources.
type clusterBindState struct {
	// The id of the cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The type of the cluster bind.
	Type *string `pulumi:"type"`
	// The id of the workspace.
	WorkspaceId *string `pulumi:"workspaceId"`
}

type ClusterBindState struct {
	// The id of the cluster.
	ClusterId pulumi.StringPtrInput
	// The type of the cluster bind.
	Type pulumi.StringPtrInput
	// The id of the workspace.
	WorkspaceId pulumi.StringPtrInput
}

func (ClusterBindState) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterBindState)(nil)).Elem()
}

type clusterBindArgs struct {
	// The id of the cluster.
	ClusterId string `pulumi:"clusterId"`
	// The type of the cluster bind.
	Type string `pulumi:"type"`
	// The id of the workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}

// The set of arguments for constructing a ClusterBind resource.
type ClusterBindArgs struct {
	// The id of the cluster.
	ClusterId pulumi.StringInput
	// The type of the cluster bind.
	Type pulumi.StringInput
	// The id of the workspace.
	WorkspaceId pulumi.StringInput
}

func (ClusterBindArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterBindArgs)(nil)).Elem()
}

type ClusterBindInput interface {
	pulumi.Input

	ToClusterBindOutput() ClusterBindOutput
	ToClusterBindOutputWithContext(ctx context.Context) ClusterBindOutput
}

func (*ClusterBind) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterBind)(nil)).Elem()
}

func (i *ClusterBind) ToClusterBindOutput() ClusterBindOutput {
	return i.ToClusterBindOutputWithContext(context.Background())
}

func (i *ClusterBind) ToClusterBindOutputWithContext(ctx context.Context) ClusterBindOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBindOutput)
}

// ClusterBindArrayInput is an input type that accepts ClusterBindArray and ClusterBindArrayOutput values.
// You can construct a concrete instance of `ClusterBindArrayInput` via:
//
//	ClusterBindArray{ ClusterBindArgs{...} }
type ClusterBindArrayInput interface {
	pulumi.Input

	ToClusterBindArrayOutput() ClusterBindArrayOutput
	ToClusterBindArrayOutputWithContext(context.Context) ClusterBindArrayOutput
}

type ClusterBindArray []ClusterBindInput

func (ClusterBindArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterBind)(nil)).Elem()
}

func (i ClusterBindArray) ToClusterBindArrayOutput() ClusterBindArrayOutput {
	return i.ToClusterBindArrayOutputWithContext(context.Background())
}

func (i ClusterBindArray) ToClusterBindArrayOutputWithContext(ctx context.Context) ClusterBindArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBindArrayOutput)
}

// ClusterBindMapInput is an input type that accepts ClusterBindMap and ClusterBindMapOutput values.
// You can construct a concrete instance of `ClusterBindMapInput` via:
//
//	ClusterBindMap{ "key": ClusterBindArgs{...} }
type ClusterBindMapInput interface {
	pulumi.Input

	ToClusterBindMapOutput() ClusterBindMapOutput
	ToClusterBindMapOutputWithContext(context.Context) ClusterBindMapOutput
}

type ClusterBindMap map[string]ClusterBindInput

func (ClusterBindMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterBind)(nil)).Elem()
}

func (i ClusterBindMap) ToClusterBindMapOutput() ClusterBindMapOutput {
	return i.ToClusterBindMapOutputWithContext(context.Background())
}

func (i ClusterBindMap) ToClusterBindMapOutputWithContext(ctx context.Context) ClusterBindMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBindMapOutput)
}

type ClusterBindOutput struct{ *pulumi.OutputState }

func (ClusterBindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterBind)(nil)).Elem()
}

func (o ClusterBindOutput) ToClusterBindOutput() ClusterBindOutput {
	return o
}

func (o ClusterBindOutput) ToClusterBindOutputWithContext(ctx context.Context) ClusterBindOutput {
	return o
}

// The id of the cluster.
func (o ClusterBindOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterBind) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The type of the cluster bind.
func (o ClusterBindOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterBind) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The id of the workspace.
func (o ClusterBindOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ClusterBind) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

type ClusterBindArrayOutput struct{ *pulumi.OutputState }

func (ClusterBindArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterBind)(nil)).Elem()
}

func (o ClusterBindArrayOutput) ToClusterBindArrayOutput() ClusterBindArrayOutput {
	return o
}

func (o ClusterBindArrayOutput) ToClusterBindArrayOutputWithContext(ctx context.Context) ClusterBindArrayOutput {
	return o
}

func (o ClusterBindArrayOutput) Index(i pulumi.IntInput) ClusterBindOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterBind {
		return vs[0].([]*ClusterBind)[vs[1].(int)]
	}).(ClusterBindOutput)
}

type ClusterBindMapOutput struct{ *pulumi.OutputState }

func (ClusterBindMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterBind)(nil)).Elem()
}

func (o ClusterBindMapOutput) ToClusterBindMapOutput() ClusterBindMapOutput {
	return o
}

func (o ClusterBindMapOutput) ToClusterBindMapOutputWithContext(ctx context.Context) ClusterBindMapOutput {
	return o
}

func (o ClusterBindMapOutput) MapIndex(k pulumi.StringInput) ClusterBindOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterBind {
		return vs[0].(map[string]*ClusterBind)[vs[1].(string)]
	}).(ClusterBindOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBindInput)(nil)).Elem(), &ClusterBind{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBindArrayInput)(nil)).Elem(), ClusterBindArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBindMapInput)(nil)).Elem(), ClusterBindMap{})
	pulumi.RegisterOutputType(ClusterBindOutput{})
	pulumi.RegisterOutputType(ClusterBindArrayOutput{})
	pulumi.RegisterOutputType(ClusterBindMapOutput{})
}
