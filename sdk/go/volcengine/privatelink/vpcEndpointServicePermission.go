// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage privatelink vpc endpoint service permission
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/privatelink"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := privatelink.NewVpcEndpointServicePermission(ctx, "foo", &privatelink.VpcEndpointServicePermissionArgs{
//				PermitAccountId: pulumi.String("210000000"),
//				ServiceId:       pulumi.String("epsvc-3rel73uf2ewao5zsk2j2l58ro"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = privatelink.NewVpcEndpointServicePermission(ctx, "foo1", &privatelink.VpcEndpointServicePermissionArgs{
//				PermitAccountId: pulumi.String("210000001"),
//				ServiceId:       pulumi.String("epsvc-3rel73uf2ewao5zsk2j2l58ro"),
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
// VpcEndpointServicePermission can be imported using the serviceId:permitAccountId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:privatelink/vpcEndpointServicePermission:VpcEndpointServicePermission default epsvc-2fe630gurkl37k5gfuy33****:2100000000
//
// ```
type VpcEndpointServicePermission struct {
	pulumi.CustomResourceState

	// The id of account.
	PermitAccountId pulumi.StringOutput `pulumi:"permitAccountId"`
	// The id of service.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
}

// NewVpcEndpointServicePermission registers a new resource with the given unique name, arguments, and options.
func NewVpcEndpointServicePermission(ctx *pulumi.Context,
	name string, args *VpcEndpointServicePermissionArgs, opts ...pulumi.ResourceOption) (*VpcEndpointServicePermission, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PermitAccountId == nil {
		return nil, errors.New("invalid value for required argument 'PermitAccountId'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource VpcEndpointServicePermission
	err := ctx.RegisterResource("volcengine:privatelink/vpcEndpointServicePermission:VpcEndpointServicePermission", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpcEndpointServicePermission gets an existing VpcEndpointServicePermission resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpcEndpointServicePermission(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpcEndpointServicePermissionState, opts ...pulumi.ResourceOption) (*VpcEndpointServicePermission, error) {
	var resource VpcEndpointServicePermission
	err := ctx.ReadResource("volcengine:privatelink/vpcEndpointServicePermission:VpcEndpointServicePermission", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpcEndpointServicePermission resources.
type vpcEndpointServicePermissionState struct {
	// The id of account.
	PermitAccountId *string `pulumi:"permitAccountId"`
	// The id of service.
	ServiceId *string `pulumi:"serviceId"`
}

type VpcEndpointServicePermissionState struct {
	// The id of account.
	PermitAccountId pulumi.StringPtrInput
	// The id of service.
	ServiceId pulumi.StringPtrInput
}

func (VpcEndpointServicePermissionState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServicePermissionState)(nil)).Elem()
}

type vpcEndpointServicePermissionArgs struct {
	// The id of account.
	PermitAccountId string `pulumi:"permitAccountId"`
	// The id of service.
	ServiceId string `pulumi:"serviceId"`
}

// The set of arguments for constructing a VpcEndpointServicePermission resource.
type VpcEndpointServicePermissionArgs struct {
	// The id of account.
	PermitAccountId pulumi.StringInput
	// The id of service.
	ServiceId pulumi.StringInput
}

func (VpcEndpointServicePermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpcEndpointServicePermissionArgs)(nil)).Elem()
}

type VpcEndpointServicePermissionInput interface {
	pulumi.Input

	ToVpcEndpointServicePermissionOutput() VpcEndpointServicePermissionOutput
	ToVpcEndpointServicePermissionOutputWithContext(ctx context.Context) VpcEndpointServicePermissionOutput
}

func (*VpcEndpointServicePermission) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServicePermission)(nil)).Elem()
}

func (i *VpcEndpointServicePermission) ToVpcEndpointServicePermissionOutput() VpcEndpointServicePermissionOutput {
	return i.ToVpcEndpointServicePermissionOutputWithContext(context.Background())
}

func (i *VpcEndpointServicePermission) ToVpcEndpointServicePermissionOutputWithContext(ctx context.Context) VpcEndpointServicePermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServicePermissionOutput)
}

// VpcEndpointServicePermissionArrayInput is an input type that accepts VpcEndpointServicePermissionArray and VpcEndpointServicePermissionArrayOutput values.
// You can construct a concrete instance of `VpcEndpointServicePermissionArrayInput` via:
//
//	VpcEndpointServicePermissionArray{ VpcEndpointServicePermissionArgs{...} }
type VpcEndpointServicePermissionArrayInput interface {
	pulumi.Input

	ToVpcEndpointServicePermissionArrayOutput() VpcEndpointServicePermissionArrayOutput
	ToVpcEndpointServicePermissionArrayOutputWithContext(context.Context) VpcEndpointServicePermissionArrayOutput
}

type VpcEndpointServicePermissionArray []VpcEndpointServicePermissionInput

func (VpcEndpointServicePermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServicePermission)(nil)).Elem()
}

func (i VpcEndpointServicePermissionArray) ToVpcEndpointServicePermissionArrayOutput() VpcEndpointServicePermissionArrayOutput {
	return i.ToVpcEndpointServicePermissionArrayOutputWithContext(context.Background())
}

func (i VpcEndpointServicePermissionArray) ToVpcEndpointServicePermissionArrayOutputWithContext(ctx context.Context) VpcEndpointServicePermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServicePermissionArrayOutput)
}

// VpcEndpointServicePermissionMapInput is an input type that accepts VpcEndpointServicePermissionMap and VpcEndpointServicePermissionMapOutput values.
// You can construct a concrete instance of `VpcEndpointServicePermissionMapInput` via:
//
//	VpcEndpointServicePermissionMap{ "key": VpcEndpointServicePermissionArgs{...} }
type VpcEndpointServicePermissionMapInput interface {
	pulumi.Input

	ToVpcEndpointServicePermissionMapOutput() VpcEndpointServicePermissionMapOutput
	ToVpcEndpointServicePermissionMapOutputWithContext(context.Context) VpcEndpointServicePermissionMapOutput
}

type VpcEndpointServicePermissionMap map[string]VpcEndpointServicePermissionInput

func (VpcEndpointServicePermissionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServicePermission)(nil)).Elem()
}

func (i VpcEndpointServicePermissionMap) ToVpcEndpointServicePermissionMapOutput() VpcEndpointServicePermissionMapOutput {
	return i.ToVpcEndpointServicePermissionMapOutputWithContext(context.Background())
}

func (i VpcEndpointServicePermissionMap) ToVpcEndpointServicePermissionMapOutputWithContext(ctx context.Context) VpcEndpointServicePermissionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpcEndpointServicePermissionMapOutput)
}

type VpcEndpointServicePermissionOutput struct{ *pulumi.OutputState }

func (VpcEndpointServicePermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpcEndpointServicePermission)(nil)).Elem()
}

func (o VpcEndpointServicePermissionOutput) ToVpcEndpointServicePermissionOutput() VpcEndpointServicePermissionOutput {
	return o
}

func (o VpcEndpointServicePermissionOutput) ToVpcEndpointServicePermissionOutputWithContext(ctx context.Context) VpcEndpointServicePermissionOutput {
	return o
}

// The id of account.
func (o VpcEndpointServicePermissionOutput) PermitAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServicePermission) pulumi.StringOutput { return v.PermitAccountId }).(pulumi.StringOutput)
}

// The id of service.
func (o VpcEndpointServicePermissionOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpcEndpointServicePermission) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

type VpcEndpointServicePermissionArrayOutput struct{ *pulumi.OutputState }

func (VpcEndpointServicePermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpcEndpointServicePermission)(nil)).Elem()
}

func (o VpcEndpointServicePermissionArrayOutput) ToVpcEndpointServicePermissionArrayOutput() VpcEndpointServicePermissionArrayOutput {
	return o
}

func (o VpcEndpointServicePermissionArrayOutput) ToVpcEndpointServicePermissionArrayOutputWithContext(ctx context.Context) VpcEndpointServicePermissionArrayOutput {
	return o
}

func (o VpcEndpointServicePermissionArrayOutput) Index(i pulumi.IntInput) VpcEndpointServicePermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpcEndpointServicePermission {
		return vs[0].([]*VpcEndpointServicePermission)[vs[1].(int)]
	}).(VpcEndpointServicePermissionOutput)
}

type VpcEndpointServicePermissionMapOutput struct{ *pulumi.OutputState }

func (VpcEndpointServicePermissionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpcEndpointServicePermission)(nil)).Elem()
}

func (o VpcEndpointServicePermissionMapOutput) ToVpcEndpointServicePermissionMapOutput() VpcEndpointServicePermissionMapOutput {
	return o
}

func (o VpcEndpointServicePermissionMapOutput) ToVpcEndpointServicePermissionMapOutputWithContext(ctx context.Context) VpcEndpointServicePermissionMapOutput {
	return o
}

func (o VpcEndpointServicePermissionMapOutput) MapIndex(k pulumi.StringInput) VpcEndpointServicePermissionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpcEndpointServicePermission {
		return vs[0].(map[string]*VpcEndpointServicePermission)[vs[1].(string)]
	}).(VpcEndpointServicePermissionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServicePermissionInput)(nil)).Elem(), &VpcEndpointServicePermission{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServicePermissionArrayInput)(nil)).Elem(), VpcEndpointServicePermissionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpcEndpointServicePermissionMapInput)(nil)).Elem(), VpcEndpointServicePermissionMap{})
	pulumi.RegisterOutputType(VpcEndpointServicePermissionOutput{})
	pulumi.RegisterOutputType(VpcEndpointServicePermissionArrayOutput{})
	pulumi.RegisterOutputType(VpcEndpointServicePermissionMapOutput{})
}
