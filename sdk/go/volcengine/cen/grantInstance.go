// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage cen grant instance
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.NewGrantInstance(ctx, "foo", &cen.GrantInstanceArgs{
//				CenId:            pulumi.String("cen-2d6zdn0c1z5s058ozfcyf4lee"),
//				CenOwnerId:       pulumi.String("210000****"),
//				InstanceId:       pulumi.String("vpc-2bysvq1xx543k2dx0eeulpeiv"),
//				InstanceRegionId: pulumi.String("cn-guilin-boe"),
//				InstanceType:     pulumi.String("VPC"),
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
// Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cen/grantInstance:GrantInstance default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
//
// ```
type GrantInstance struct {
	pulumi.CustomResourceState

	// The ID of the cen.
	CenId pulumi.StringOutput `pulumi:"cenId"`
	// The owner ID of the cen.
	CenOwnerId pulumi.StringOutput `pulumi:"cenOwnerId"`
	// The ID of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The region ID of the instance.
	InstanceRegionId pulumi.StringOutput `pulumi:"instanceRegionId"`
	// The type of the instance.
	InstanceType pulumi.StringOutput `pulumi:"instanceType"`
}

// NewGrantInstance registers a new resource with the given unique name, arguments, and options.
func NewGrantInstance(ctx *pulumi.Context,
	name string, args *GrantInstanceArgs, opts ...pulumi.ResourceOption) (*GrantInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	if args.CenOwnerId == nil {
		return nil, errors.New("invalid value for required argument 'CenOwnerId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.InstanceRegionId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceRegionId'")
	}
	if args.InstanceType == nil {
		return nil, errors.New("invalid value for required argument 'InstanceType'")
	}
	var resource GrantInstance
	err := ctx.RegisterResource("volcengine:cen/grantInstance:GrantInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrantInstance gets an existing GrantInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrantInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrantInstanceState, opts ...pulumi.ResourceOption) (*GrantInstance, error) {
	var resource GrantInstance
	err := ctx.ReadResource("volcengine:cen/grantInstance:GrantInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GrantInstance resources.
type grantInstanceState struct {
	// The ID of the cen.
	CenId *string `pulumi:"cenId"`
	// The owner ID of the cen.
	CenOwnerId *string `pulumi:"cenOwnerId"`
	// The ID of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The region ID of the instance.
	InstanceRegionId *string `pulumi:"instanceRegionId"`
	// The type of the instance.
	InstanceType *string `pulumi:"instanceType"`
}

type GrantInstanceState struct {
	// The ID of the cen.
	CenId pulumi.StringPtrInput
	// The owner ID of the cen.
	CenOwnerId pulumi.StringPtrInput
	// The ID of the instance.
	InstanceId pulumi.StringPtrInput
	// The region ID of the instance.
	InstanceRegionId pulumi.StringPtrInput
	// The type of the instance.
	InstanceType pulumi.StringPtrInput
}

func (GrantInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*grantInstanceState)(nil)).Elem()
}

type grantInstanceArgs struct {
	// The ID of the cen.
	CenId string `pulumi:"cenId"`
	// The owner ID of the cen.
	CenOwnerId string `pulumi:"cenOwnerId"`
	// The ID of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The region ID of the instance.
	InstanceRegionId string `pulumi:"instanceRegionId"`
	// The type of the instance.
	InstanceType string `pulumi:"instanceType"`
}

// The set of arguments for constructing a GrantInstance resource.
type GrantInstanceArgs struct {
	// The ID of the cen.
	CenId pulumi.StringInput
	// The owner ID of the cen.
	CenOwnerId pulumi.StringInput
	// The ID of the instance.
	InstanceId pulumi.StringInput
	// The region ID of the instance.
	InstanceRegionId pulumi.StringInput
	// The type of the instance.
	InstanceType pulumi.StringInput
}

func (GrantInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grantInstanceArgs)(nil)).Elem()
}

type GrantInstanceInput interface {
	pulumi.Input

	ToGrantInstanceOutput() GrantInstanceOutput
	ToGrantInstanceOutputWithContext(ctx context.Context) GrantInstanceOutput
}

func (*GrantInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantInstance)(nil)).Elem()
}

func (i *GrantInstance) ToGrantInstanceOutput() GrantInstanceOutput {
	return i.ToGrantInstanceOutputWithContext(context.Background())
}

func (i *GrantInstance) ToGrantInstanceOutputWithContext(ctx context.Context) GrantInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantInstanceOutput)
}

// GrantInstanceArrayInput is an input type that accepts GrantInstanceArray and GrantInstanceArrayOutput values.
// You can construct a concrete instance of `GrantInstanceArrayInput` via:
//
//	GrantInstanceArray{ GrantInstanceArgs{...} }
type GrantInstanceArrayInput interface {
	pulumi.Input

	ToGrantInstanceArrayOutput() GrantInstanceArrayOutput
	ToGrantInstanceArrayOutputWithContext(context.Context) GrantInstanceArrayOutput
}

type GrantInstanceArray []GrantInstanceInput

func (GrantInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrantInstance)(nil)).Elem()
}

func (i GrantInstanceArray) ToGrantInstanceArrayOutput() GrantInstanceArrayOutput {
	return i.ToGrantInstanceArrayOutputWithContext(context.Background())
}

func (i GrantInstanceArray) ToGrantInstanceArrayOutputWithContext(ctx context.Context) GrantInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantInstanceArrayOutput)
}

// GrantInstanceMapInput is an input type that accepts GrantInstanceMap and GrantInstanceMapOutput values.
// You can construct a concrete instance of `GrantInstanceMapInput` via:
//
//	GrantInstanceMap{ "key": GrantInstanceArgs{...} }
type GrantInstanceMapInput interface {
	pulumi.Input

	ToGrantInstanceMapOutput() GrantInstanceMapOutput
	ToGrantInstanceMapOutputWithContext(context.Context) GrantInstanceMapOutput
}

type GrantInstanceMap map[string]GrantInstanceInput

func (GrantInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrantInstance)(nil)).Elem()
}

func (i GrantInstanceMap) ToGrantInstanceMapOutput() GrantInstanceMapOutput {
	return i.ToGrantInstanceMapOutputWithContext(context.Background())
}

func (i GrantInstanceMap) ToGrantInstanceMapOutputWithContext(ctx context.Context) GrantInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrantInstanceMapOutput)
}

type GrantInstanceOutput struct{ *pulumi.OutputState }

func (GrantInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrantInstance)(nil)).Elem()
}

func (o GrantInstanceOutput) ToGrantInstanceOutput() GrantInstanceOutput {
	return o
}

func (o GrantInstanceOutput) ToGrantInstanceOutputWithContext(ctx context.Context) GrantInstanceOutput {
	return o
}

// The ID of the cen.
func (o GrantInstanceOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantInstance) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

// The owner ID of the cen.
func (o GrantInstanceOutput) CenOwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantInstance) pulumi.StringOutput { return v.CenOwnerId }).(pulumi.StringOutput)
}

// The ID of the instance.
func (o GrantInstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantInstance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The region ID of the instance.
func (o GrantInstanceOutput) InstanceRegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantInstance) pulumi.StringOutput { return v.InstanceRegionId }).(pulumi.StringOutput)
}

// The type of the instance.
func (o GrantInstanceOutput) InstanceType() pulumi.StringOutput {
	return o.ApplyT(func(v *GrantInstance) pulumi.StringOutput { return v.InstanceType }).(pulumi.StringOutput)
}

type GrantInstanceArrayOutput struct{ *pulumi.OutputState }

func (GrantInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrantInstance)(nil)).Elem()
}

func (o GrantInstanceArrayOutput) ToGrantInstanceArrayOutput() GrantInstanceArrayOutput {
	return o
}

func (o GrantInstanceArrayOutput) ToGrantInstanceArrayOutputWithContext(ctx context.Context) GrantInstanceArrayOutput {
	return o
}

func (o GrantInstanceArrayOutput) Index(i pulumi.IntInput) GrantInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GrantInstance {
		return vs[0].([]*GrantInstance)[vs[1].(int)]
	}).(GrantInstanceOutput)
}

type GrantInstanceMapOutput struct{ *pulumi.OutputState }

func (GrantInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrantInstance)(nil)).Elem()
}

func (o GrantInstanceMapOutput) ToGrantInstanceMapOutput() GrantInstanceMapOutput {
	return o
}

func (o GrantInstanceMapOutput) ToGrantInstanceMapOutputWithContext(ctx context.Context) GrantInstanceMapOutput {
	return o
}

func (o GrantInstanceMapOutput) MapIndex(k pulumi.StringInput) GrantInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GrantInstance {
		return vs[0].(map[string]*GrantInstance)[vs[1].(string)]
	}).(GrantInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrantInstanceInput)(nil)).Elem(), &GrantInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantInstanceArrayInput)(nil)).Elem(), GrantInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrantInstanceMapInput)(nil)).Elem(), GrantInstanceMap{})
	pulumi.RegisterOutputType(GrantInstanceOutput{})
	pulumi.RegisterOutputType(GrantInstanceArrayOutput{})
	pulumi.RegisterOutputType(GrantInstanceMapOutput{})
}
