// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage snat entry
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/nat"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nat.NewSnatEntry(ctx, "foo", &nat.SnatEntryArgs{
//				EipId:         pulumi.String("eip-274zlae117nr47fap8tzl24v4"),
//				NatGatewayId:  pulumi.String("ngw-2743w1f6iqby87fap8tvm9kop"),
//				SnatEntryName: pulumi.String("tf-test-up"),
//				SubnetId:      pulumi.String("subnet-2744i7u9alnnk7fap8tkq8aft"),
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
// Snat entry can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:nat/snatEntry:SnatEntry default snat-3fvhk47kf56****
//
// ```
type SnatEntry struct {
	pulumi.CustomResourceState

	// The id of the public ip address used by the SNAT entry.
	EipId pulumi.StringOutput `pulumi:"eipId"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringOutput `pulumi:"natGatewayId"`
	// The name of the SNAT entry.
	SnatEntryName pulumi.StringOutput `pulumi:"snatEntryName"`
	// The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
	SourceCidr pulumi.StringPtrOutput `pulumi:"sourceCidr"`
	// The status of the SNAT entry.
	Status pulumi.StringOutput `pulumi:"status"`
	// The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
	SubnetId pulumi.StringPtrOutput `pulumi:"subnetId"`
}

// NewSnatEntry registers a new resource with the given unique name, arguments, and options.
func NewSnatEntry(ctx *pulumi.Context,
	name string, args *SnatEntryArgs, opts ...pulumi.ResourceOption) (*SnatEntry, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EipId == nil {
		return nil, errors.New("invalid value for required argument 'EipId'")
	}
	if args.NatGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'NatGatewayId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource SnatEntry
	err := ctx.RegisterResource("volcengine:nat/snatEntry:SnatEntry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSnatEntry gets an existing SnatEntry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSnatEntry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SnatEntryState, opts ...pulumi.ResourceOption) (*SnatEntry, error) {
	var resource SnatEntry
	err := ctx.ReadResource("volcengine:nat/snatEntry:SnatEntry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SnatEntry resources.
type snatEntryState struct {
	// The id of the public ip address used by the SNAT entry.
	EipId *string `pulumi:"eipId"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// The name of the SNAT entry.
	SnatEntryName *string `pulumi:"snatEntryName"`
	// The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
	SourceCidr *string `pulumi:"sourceCidr"`
	// The status of the SNAT entry.
	Status *string `pulumi:"status"`
	// The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
	SubnetId *string `pulumi:"subnetId"`
}

type SnatEntryState struct {
	// The id of the public ip address used by the SNAT entry.
	EipId pulumi.StringPtrInput
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringPtrInput
	// The name of the SNAT entry.
	SnatEntryName pulumi.StringPtrInput
	// The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
	SourceCidr pulumi.StringPtrInput
	// The status of the SNAT entry.
	Status pulumi.StringPtrInput
	// The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
	SubnetId pulumi.StringPtrInput
}

func (SnatEntryState) ElementType() reflect.Type {
	return reflect.TypeOf((*snatEntryState)(nil)).Elem()
}

type snatEntryArgs struct {
	// The id of the public ip address used by the SNAT entry.
	EipId string `pulumi:"eipId"`
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId string `pulumi:"natGatewayId"`
	// The name of the SNAT entry.
	SnatEntryName *string `pulumi:"snatEntryName"`
	// The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
	SourceCidr *string `pulumi:"sourceCidr"`
	// The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
	SubnetId *string `pulumi:"subnetId"`
}

// The set of arguments for constructing a SnatEntry resource.
type SnatEntryArgs struct {
	// The id of the public ip address used by the SNAT entry.
	EipId pulumi.StringInput
	// The id of the nat gateway to which the entry belongs.
	NatGatewayId pulumi.StringInput
	// The name of the SNAT entry.
	SnatEntryName pulumi.StringPtrInput
	// The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
	SourceCidr pulumi.StringPtrInput
	// The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
	SubnetId pulumi.StringPtrInput
}

func (SnatEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*snatEntryArgs)(nil)).Elem()
}

type SnatEntryInput interface {
	pulumi.Input

	ToSnatEntryOutput() SnatEntryOutput
	ToSnatEntryOutputWithContext(ctx context.Context) SnatEntryOutput
}

func (*SnatEntry) ElementType() reflect.Type {
	return reflect.TypeOf((**SnatEntry)(nil)).Elem()
}

func (i *SnatEntry) ToSnatEntryOutput() SnatEntryOutput {
	return i.ToSnatEntryOutputWithContext(context.Background())
}

func (i *SnatEntry) ToSnatEntryOutputWithContext(ctx context.Context) SnatEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatEntryOutput)
}

// SnatEntryArrayInput is an input type that accepts SnatEntryArray and SnatEntryArrayOutput values.
// You can construct a concrete instance of `SnatEntryArrayInput` via:
//
//	SnatEntryArray{ SnatEntryArgs{...} }
type SnatEntryArrayInput interface {
	pulumi.Input

	ToSnatEntryArrayOutput() SnatEntryArrayOutput
	ToSnatEntryArrayOutputWithContext(context.Context) SnatEntryArrayOutput
}

type SnatEntryArray []SnatEntryInput

func (SnatEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnatEntry)(nil)).Elem()
}

func (i SnatEntryArray) ToSnatEntryArrayOutput() SnatEntryArrayOutput {
	return i.ToSnatEntryArrayOutputWithContext(context.Background())
}

func (i SnatEntryArray) ToSnatEntryArrayOutputWithContext(ctx context.Context) SnatEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatEntryArrayOutput)
}

// SnatEntryMapInput is an input type that accepts SnatEntryMap and SnatEntryMapOutput values.
// You can construct a concrete instance of `SnatEntryMapInput` via:
//
//	SnatEntryMap{ "key": SnatEntryArgs{...} }
type SnatEntryMapInput interface {
	pulumi.Input

	ToSnatEntryMapOutput() SnatEntryMapOutput
	ToSnatEntryMapOutputWithContext(context.Context) SnatEntryMapOutput
}

type SnatEntryMap map[string]SnatEntryInput

func (SnatEntryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnatEntry)(nil)).Elem()
}

func (i SnatEntryMap) ToSnatEntryMapOutput() SnatEntryMapOutput {
	return i.ToSnatEntryMapOutputWithContext(context.Background())
}

func (i SnatEntryMap) ToSnatEntryMapOutputWithContext(ctx context.Context) SnatEntryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SnatEntryMapOutput)
}

type SnatEntryOutput struct{ *pulumi.OutputState }

func (SnatEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SnatEntry)(nil)).Elem()
}

func (o SnatEntryOutput) ToSnatEntryOutput() SnatEntryOutput {
	return o
}

func (o SnatEntryOutput) ToSnatEntryOutputWithContext(ctx context.Context) SnatEntryOutput {
	return o
}

// The id of the public ip address used by the SNAT entry.
func (o SnatEntryOutput) EipId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.EipId }).(pulumi.StringOutput)
}

// The id of the nat gateway to which the entry belongs.
func (o SnatEntryOutput) NatGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.NatGatewayId }).(pulumi.StringOutput)
}

// The name of the SNAT entry.
func (o SnatEntryOutput) SnatEntryName() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.SnatEntryName }).(pulumi.StringOutput)
}

// The SourceCidr of the SNAT entry. Only one of `subnet_id,source_cidr` can be specified.
func (o SnatEntryOutput) SourceCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringPtrOutput { return v.SourceCidr }).(pulumi.StringPtrOutput)
}

// The status of the SNAT entry.
func (o SnatEntryOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The id of the subnet that is required to access the internet. Only one of `subnet_id,source_cidr` can be specified.
func (o SnatEntryOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SnatEntry) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

type SnatEntryArrayOutput struct{ *pulumi.OutputState }

func (SnatEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SnatEntry)(nil)).Elem()
}

func (o SnatEntryArrayOutput) ToSnatEntryArrayOutput() SnatEntryArrayOutput {
	return o
}

func (o SnatEntryArrayOutput) ToSnatEntryArrayOutputWithContext(ctx context.Context) SnatEntryArrayOutput {
	return o
}

func (o SnatEntryArrayOutput) Index(i pulumi.IntInput) SnatEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SnatEntry {
		return vs[0].([]*SnatEntry)[vs[1].(int)]
	}).(SnatEntryOutput)
}

type SnatEntryMapOutput struct{ *pulumi.OutputState }

func (SnatEntryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SnatEntry)(nil)).Elem()
}

func (o SnatEntryMapOutput) ToSnatEntryMapOutput() SnatEntryMapOutput {
	return o
}

func (o SnatEntryMapOutput) ToSnatEntryMapOutputWithContext(ctx context.Context) SnatEntryMapOutput {
	return o
}

func (o SnatEntryMapOutput) MapIndex(k pulumi.StringInput) SnatEntryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SnatEntry {
		return vs[0].(map[string]*SnatEntry)[vs[1].(string)]
	}).(SnatEntryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SnatEntryInput)(nil)).Elem(), &SnatEntry{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatEntryArrayInput)(nil)).Elem(), SnatEntryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SnatEntryMapInput)(nil)).Elem(), SnatEntryMap{})
	pulumi.RegisterOutputType(SnatEntryOutput{})
	pulumi.RegisterOutputType(SnatEntryArrayOutput{})
	pulumi.RegisterOutputType(SnatEntryMapOutput{})
}
