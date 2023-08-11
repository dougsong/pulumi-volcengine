// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage network acl
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NewNetworkAcl(ctx, "foo", &vpc.NetworkAclArgs{
//				EgressAclEntries: vpc.NetworkAclEgressAclEntryArray{
//					&vpc.NetworkAclEgressAclEntryArgs{
//						DestinationCidrIp:   pulumi.String("192.168.0.0/16"),
//						NetworkAclEntryName: pulumi.String("egress2"),
//						Policy:              pulumi.String("accept"),
//						Protocol:            pulumi.String("all"),
//					},
//				},
//				IngressAclEntries: vpc.NetworkAclIngressAclEntryArray{
//					&vpc.NetworkAclIngressAclEntryArgs{
//						NetworkAclEntryName: pulumi.String("ingress1"),
//						Policy:              pulumi.String("accept"),
//						Protocol:            pulumi.String("all"),
//						SourceCidrIp:        pulumi.String("192.168.0.0/24"),
//					},
//					&vpc.NetworkAclIngressAclEntryArgs{
//						NetworkAclEntryName: pulumi.String("ingress3"),
//						Policy:              pulumi.String("accept"),
//						Port:                pulumi.String("80/80"),
//						Protocol:            pulumi.String("tcp"),
//						SourceCidrIp:        pulumi.String("192.168.0.0/24"),
//					},
//				},
//				NetworkAclName: pulumi.String("tf-test-acl"),
//				ProjectName:    pulumi.String("default"),
//				VpcId:          pulumi.String("vpc-2d6jskar243k058ozfdae13ne"),
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
// Network Acl can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vpc/networkAcl:NetworkAcl default nacl-172leak37mi9s4d1w33pswqkh
//
// ```
type NetworkAcl struct {
	pulumi.CustomResourceState

	// The description of the Network Acl.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The egress entries of Network Acl.
	EgressAclEntries NetworkAclEgressAclEntryArrayOutput `pulumi:"egressAclEntries"`
	// The ingress entries of Network Acl.
	IngressAclEntries NetworkAclIngressAclEntryArrayOutput `pulumi:"ingressAclEntries"`
	// The name of Network Acl.
	NetworkAclName pulumi.StringOutput `pulumi:"networkAclName"`
	// The project name of the network acl.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The vpc id of Network Acl.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewNetworkAcl registers a new resource with the given unique name, arguments, and options.
func NewNetworkAcl(ctx *pulumi.Context,
	name string, args *NetworkAclArgs, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource NetworkAcl
	err := ctx.RegisterResource("volcengine:vpc/networkAcl:NetworkAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkAcl gets an existing NetworkAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkAclState, opts ...pulumi.ResourceOption) (*NetworkAcl, error) {
	var resource NetworkAcl
	err := ctx.ReadResource("volcengine:vpc/networkAcl:NetworkAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkAcl resources.
type networkAclState struct {
	// The description of the Network Acl.
	Description *string `pulumi:"description"`
	// The egress entries of Network Acl.
	EgressAclEntries []NetworkAclEgressAclEntry `pulumi:"egressAclEntries"`
	// The ingress entries of Network Acl.
	IngressAclEntries []NetworkAclIngressAclEntry `pulumi:"ingressAclEntries"`
	// The name of Network Acl.
	NetworkAclName *string `pulumi:"networkAclName"`
	// The project name of the network acl.
	ProjectName *string `pulumi:"projectName"`
	// The vpc id of Network Acl.
	VpcId *string `pulumi:"vpcId"`
}

type NetworkAclState struct {
	// The description of the Network Acl.
	Description pulumi.StringPtrInput
	// The egress entries of Network Acl.
	EgressAclEntries NetworkAclEgressAclEntryArrayInput
	// The ingress entries of Network Acl.
	IngressAclEntries NetworkAclIngressAclEntryArrayInput
	// The name of Network Acl.
	NetworkAclName pulumi.StringPtrInput
	// The project name of the network acl.
	ProjectName pulumi.StringPtrInput
	// The vpc id of Network Acl.
	VpcId pulumi.StringPtrInput
}

func (NetworkAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclState)(nil)).Elem()
}

type networkAclArgs struct {
	// The description of the Network Acl.
	Description *string `pulumi:"description"`
	// The egress entries of Network Acl.
	EgressAclEntries []NetworkAclEgressAclEntry `pulumi:"egressAclEntries"`
	// The ingress entries of Network Acl.
	IngressAclEntries []NetworkAclIngressAclEntry `pulumi:"ingressAclEntries"`
	// The name of Network Acl.
	NetworkAclName *string `pulumi:"networkAclName"`
	// The project name of the network acl.
	ProjectName *string `pulumi:"projectName"`
	// The vpc id of Network Acl.
	VpcId string `pulumi:"vpcId"`
}

// The set of arguments for constructing a NetworkAcl resource.
type NetworkAclArgs struct {
	// The description of the Network Acl.
	Description pulumi.StringPtrInput
	// The egress entries of Network Acl.
	EgressAclEntries NetworkAclEgressAclEntryArrayInput
	// The ingress entries of Network Acl.
	IngressAclEntries NetworkAclIngressAclEntryArrayInput
	// The name of Network Acl.
	NetworkAclName pulumi.StringPtrInput
	// The project name of the network acl.
	ProjectName pulumi.StringPtrInput
	// The vpc id of Network Acl.
	VpcId pulumi.StringInput
}

func (NetworkAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkAclArgs)(nil)).Elem()
}

type NetworkAclInput interface {
	pulumi.Input

	ToNetworkAclOutput() NetworkAclOutput
	ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput
}

func (*NetworkAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (i *NetworkAcl) ToNetworkAclOutput() NetworkAclOutput {
	return i.ToNetworkAclOutputWithContext(context.Background())
}

func (i *NetworkAcl) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclOutput)
}

// NetworkAclArrayInput is an input type that accepts NetworkAclArray and NetworkAclArrayOutput values.
// You can construct a concrete instance of `NetworkAclArrayInput` via:
//
//	NetworkAclArray{ NetworkAclArgs{...} }
type NetworkAclArrayInput interface {
	pulumi.Input

	ToNetworkAclArrayOutput() NetworkAclArrayOutput
	ToNetworkAclArrayOutputWithContext(context.Context) NetworkAclArrayOutput
}

type NetworkAclArray []NetworkAclInput

func (NetworkAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclArray) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return i.ToNetworkAclArrayOutputWithContext(context.Background())
}

func (i NetworkAclArray) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclArrayOutput)
}

// NetworkAclMapInput is an input type that accepts NetworkAclMap and NetworkAclMapOutput values.
// You can construct a concrete instance of `NetworkAclMapInput` via:
//
//	NetworkAclMap{ "key": NetworkAclArgs{...} }
type NetworkAclMapInput interface {
	pulumi.Input

	ToNetworkAclMapOutput() NetworkAclMapOutput
	ToNetworkAclMapOutputWithContext(context.Context) NetworkAclMapOutput
}

type NetworkAclMap map[string]NetworkAclInput

func (NetworkAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (i NetworkAclMap) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return i.ToNetworkAclMapOutputWithContext(context.Background())
}

func (i NetworkAclMap) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAclMapOutput)
}

type NetworkAclOutput struct{ *pulumi.OutputState }

func (NetworkAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkAcl)(nil)).Elem()
}

func (o NetworkAclOutput) ToNetworkAclOutput() NetworkAclOutput {
	return o
}

func (o NetworkAclOutput) ToNetworkAclOutputWithContext(ctx context.Context) NetworkAclOutput {
	return o
}

// The description of the Network Acl.
func (o NetworkAclOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The egress entries of Network Acl.
func (o NetworkAclOutput) EgressAclEntries() NetworkAclEgressAclEntryArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclEgressAclEntryArrayOutput { return v.EgressAclEntries }).(NetworkAclEgressAclEntryArrayOutput)
}

// The ingress entries of Network Acl.
func (o NetworkAclOutput) IngressAclEntries() NetworkAclIngressAclEntryArrayOutput {
	return o.ApplyT(func(v *NetworkAcl) NetworkAclIngressAclEntryArrayOutput { return v.IngressAclEntries }).(NetworkAclIngressAclEntryArrayOutput)
}

// The name of Network Acl.
func (o NetworkAclOutput) NetworkAclName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.NetworkAclName }).(pulumi.StringOutput)
}

// The project name of the network acl.
func (o NetworkAclOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The vpc id of Network Acl.
func (o NetworkAclOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkAcl) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type NetworkAclArrayOutput struct{ *pulumi.OutputState }

func (NetworkAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutput() NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) ToNetworkAclArrayOutputWithContext(ctx context.Context) NetworkAclArrayOutput {
	return o
}

func (o NetworkAclArrayOutput) Index(i pulumi.IntInput) NetworkAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].([]*NetworkAcl)[vs[1].(int)]
	}).(NetworkAclOutput)
}

type NetworkAclMapOutput struct{ *pulumi.OutputState }

func (NetworkAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkAcl)(nil)).Elem()
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutput() NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) ToNetworkAclMapOutputWithContext(ctx context.Context) NetworkAclMapOutput {
	return o
}

func (o NetworkAclMapOutput) MapIndex(k pulumi.StringInput) NetworkAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkAcl {
		return vs[0].(map[string]*NetworkAcl)[vs[1].(string)]
	}).(NetworkAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclInput)(nil)).Elem(), &NetworkAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclArrayInput)(nil)).Elem(), NetworkAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkAclMapInput)(nil)).Elem(), NetworkAclMap{})
	pulumi.RegisterOutputType(NetworkAclOutput{})
	pulumi.RegisterOutputType(NetworkAclArrayOutput{})
	pulumi.RegisterOutputType(NetworkAclMapOutput{})
}
