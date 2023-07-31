// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage cr registry
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.NewRegistry(ctx, "foo", &cr.RegistryArgs{
//				DeleteImmediately: pulumi.Bool(false),
//				Password:          pulumi.String("1qaz!QAZ"),
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
// CR Instance can be imported using the name, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cr/registry:Registry default enterprise-x
//
// ```
type Registry struct {
	pulumi.CustomResourceState

	// The charge type of registry.
	ChargeType pulumi.StringOutput `pulumi:"chargeType"`
	// The creation time of registry.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether delete registry immediately. Only effected in delete action.
	DeleteImmediately pulumi.BoolPtrOutput `pulumi:"deleteImmediately"`
	// The domain of registry.
	Domains RegistryDomainArrayOutput `pulumi:"domains"`
	// The name of registry.
	Name pulumi.StringOutput `pulumi:"name"`
	// The password of registry user.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The status of registry.
	Statuses RegistryStatusArrayOutput `pulumi:"statuses"`
	// The type of registry.
	Type pulumi.StringOutput `pulumi:"type"`
	// The status of user.
	UserStatus pulumi.StringOutput `pulumi:"userStatus"`
	// The username of cr instance.
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil {
		args = &RegistryArgs{}
	}

	var resource Registry
	err := ctx.RegisterResource("volcengine:cr/registry:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("volcengine:cr/registry:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// The charge type of registry.
	ChargeType *string `pulumi:"chargeType"`
	// The creation time of registry.
	CreateTime *string `pulumi:"createTime"`
	// Whether delete registry immediately. Only effected in delete action.
	DeleteImmediately *bool `pulumi:"deleteImmediately"`
	// The domain of registry.
	Domains []RegistryDomain `pulumi:"domains"`
	// The name of registry.
	Name *string `pulumi:"name"`
	// The password of registry user.
	Password *string `pulumi:"password"`
	// The status of registry.
	Statuses []RegistryStatus `pulumi:"statuses"`
	// The type of registry.
	Type *string `pulumi:"type"`
	// The status of user.
	UserStatus *string `pulumi:"userStatus"`
	// The username of cr instance.
	Username *string `pulumi:"username"`
}

type RegistryState struct {
	// The charge type of registry.
	ChargeType pulumi.StringPtrInput
	// The creation time of registry.
	CreateTime pulumi.StringPtrInput
	// Whether delete registry immediately. Only effected in delete action.
	DeleteImmediately pulumi.BoolPtrInput
	// The domain of registry.
	Domains RegistryDomainArrayInput
	// The name of registry.
	Name pulumi.StringPtrInput
	// The password of registry user.
	Password pulumi.StringPtrInput
	// The status of registry.
	Statuses RegistryStatusArrayInput
	// The type of registry.
	Type pulumi.StringPtrInput
	// The status of user.
	UserStatus pulumi.StringPtrInput
	// The username of cr instance.
	Username pulumi.StringPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// Whether delete registry immediately. Only effected in delete action.
	DeleteImmediately *bool `pulumi:"deleteImmediately"`
	// The name of registry.
	Name *string `pulumi:"name"`
	// The password of registry user.
	Password *string `pulumi:"password"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// Whether delete registry immediately. Only effected in delete action.
	DeleteImmediately pulumi.BoolPtrInput
	// The name of registry.
	Name pulumi.StringPtrInput
	// The password of registry user.
	Password pulumi.StringPtrInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}

type RegistryInput interface {
	pulumi.Input

	ToRegistryOutput() RegistryOutput
	ToRegistryOutputWithContext(ctx context.Context) RegistryOutput
}

func (*Registry) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (i *Registry) ToRegistryOutput() RegistryOutput {
	return i.ToRegistryOutputWithContext(context.Background())
}

func (i *Registry) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryOutput)
}

// RegistryArrayInput is an input type that accepts RegistryArray and RegistryArrayOutput values.
// You can construct a concrete instance of `RegistryArrayInput` via:
//
//	RegistryArray{ RegistryArgs{...} }
type RegistryArrayInput interface {
	pulumi.Input

	ToRegistryArrayOutput() RegistryArrayOutput
	ToRegistryArrayOutputWithContext(context.Context) RegistryArrayOutput
}

type RegistryArray []RegistryInput

func (RegistryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registry)(nil)).Elem()
}

func (i RegistryArray) ToRegistryArrayOutput() RegistryArrayOutput {
	return i.ToRegistryArrayOutputWithContext(context.Background())
}

func (i RegistryArray) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryArrayOutput)
}

// RegistryMapInput is an input type that accepts RegistryMap and RegistryMapOutput values.
// You can construct a concrete instance of `RegistryMapInput` via:
//
//	RegistryMap{ "key": RegistryArgs{...} }
type RegistryMapInput interface {
	pulumi.Input

	ToRegistryMapOutput() RegistryMapOutput
	ToRegistryMapOutputWithContext(context.Context) RegistryMapOutput
}

type RegistryMap map[string]RegistryInput

func (RegistryMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registry)(nil)).Elem()
}

func (i RegistryMap) ToRegistryMapOutput() RegistryMapOutput {
	return i.ToRegistryMapOutputWithContext(context.Background())
}

func (i RegistryMap) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryMapOutput)
}

type RegistryOutput struct{ *pulumi.OutputState }

func (RegistryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Registry)(nil)).Elem()
}

func (o RegistryOutput) ToRegistryOutput() RegistryOutput {
	return o
}

func (o RegistryOutput) ToRegistryOutputWithContext(ctx context.Context) RegistryOutput {
	return o
}

// The charge type of registry.
func (o RegistryOutput) ChargeType() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.ChargeType }).(pulumi.StringOutput)
}

// The creation time of registry.
func (o RegistryOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether delete registry immediately. Only effected in delete action.
func (o RegistryOutput) DeleteImmediately() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.BoolPtrOutput { return v.DeleteImmediately }).(pulumi.BoolPtrOutput)
}

// The domain of registry.
func (o RegistryOutput) Domains() RegistryDomainArrayOutput {
	return o.ApplyT(func(v *Registry) RegistryDomainArrayOutput { return v.Domains }).(RegistryDomainArrayOutput)
}

// The name of registry.
func (o RegistryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password of registry user.
func (o RegistryOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The status of registry.
func (o RegistryOutput) Statuses() RegistryStatusArrayOutput {
	return o.ApplyT(func(v *Registry) RegistryStatusArrayOutput { return v.Statuses }).(RegistryStatusArrayOutput)
}

// The type of registry.
func (o RegistryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The status of user.
func (o RegistryOutput) UserStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.UserStatus }).(pulumi.StringOutput)
}

// The username of cr instance.
func (o RegistryOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *Registry) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type RegistryArrayOutput struct{ *pulumi.OutputState }

func (RegistryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Registry)(nil)).Elem()
}

func (o RegistryArrayOutput) ToRegistryArrayOutput() RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) ToRegistryArrayOutputWithContext(ctx context.Context) RegistryArrayOutput {
	return o
}

func (o RegistryArrayOutput) Index(i pulumi.IntInput) RegistryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Registry {
		return vs[0].([]*Registry)[vs[1].(int)]
	}).(RegistryOutput)
}

type RegistryMapOutput struct{ *pulumi.OutputState }

func (RegistryMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Registry)(nil)).Elem()
}

func (o RegistryMapOutput) ToRegistryMapOutput() RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) ToRegistryMapOutputWithContext(ctx context.Context) RegistryMapOutput {
	return o
}

func (o RegistryMapOutput) MapIndex(k pulumi.StringInput) RegistryOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Registry {
		return vs[0].(map[string]*Registry)[vs[1].(string)]
	}).(RegistryOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryInput)(nil)).Elem(), &Registry{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryArrayInput)(nil)).Elem(), RegistryArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegistryMapInput)(nil)).Elem(), RegistryMap{})
	pulumi.RegisterOutputType(RegistryOutput{})
	pulumi.RegisterOutputType(RegistryArrayOutput{})
	pulumi.RegisterOutputType(RegistryMapOutput{})
}
