// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage iam role
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/iam"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iam.NewRole(ctx, "foo", &iam.RoleArgs{
//				Description:         pulumi.String("created by terraform"),
//				DisplayName:         pulumi.String("terraform role"),
//				MaxSessionDuration:  pulumi.Int(43200),
//				RoleName:            pulumi.String("TerraformTestRole"),
//				TrustPolicyDocument: pulumi.String("{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"),
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
// Iam role can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:iam/role:Role default TerraformTestRole
//
// ```
type Role struct {
	pulumi.CustomResourceState

	// The description of the Role.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the Role.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The max session duration of the Role.
	MaxSessionDuration pulumi.IntPtrOutput `pulumi:"maxSessionDuration"`
	// The name of the Role.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The resource name of the Role.
	Trn pulumi.StringOutput `pulumi:"trn"`
	// The trust policy document of the Role.
	TrustPolicyDocument pulumi.StringOutput `pulumi:"trustPolicyDocument"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	if args.TrustPolicyDocument == nil {
		return nil, errors.New("invalid value for required argument 'TrustPolicyDocument'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Role
	err := ctx.RegisterResource("volcengine:iam/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("volcengine:iam/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The description of the Role.
	Description *string `pulumi:"description"`
	// The display name of the Role.
	DisplayName *string `pulumi:"displayName"`
	// The max session duration of the Role.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// The name of the Role.
	RoleName *string `pulumi:"roleName"`
	// The resource name of the Role.
	Trn *string `pulumi:"trn"`
	// The trust policy document of the Role.
	TrustPolicyDocument *string `pulumi:"trustPolicyDocument"`
}

type RoleState struct {
	// The description of the Role.
	Description pulumi.StringPtrInput
	// The display name of the Role.
	DisplayName pulumi.StringPtrInput
	// The max session duration of the Role.
	MaxSessionDuration pulumi.IntPtrInput
	// The name of the Role.
	RoleName pulumi.StringPtrInput
	// The resource name of the Role.
	Trn pulumi.StringPtrInput
	// The trust policy document of the Role.
	TrustPolicyDocument pulumi.StringPtrInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// The description of the Role.
	Description *string `pulumi:"description"`
	// The display name of the Role.
	DisplayName string `pulumi:"displayName"`
	// The max session duration of the Role.
	MaxSessionDuration *int `pulumi:"maxSessionDuration"`
	// The name of the Role.
	RoleName string `pulumi:"roleName"`
	// The trust policy document of the Role.
	TrustPolicyDocument string `pulumi:"trustPolicyDocument"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The description of the Role.
	Description pulumi.StringPtrInput
	// The display name of the Role.
	DisplayName pulumi.StringInput
	// The max session duration of the Role.
	MaxSessionDuration pulumi.IntPtrInput
	// The name of the Role.
	RoleName pulumi.StringInput
	// The trust policy document of the Role.
	TrustPolicyDocument pulumi.StringInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}

type RoleInput interface {
	pulumi.Input

	ToRoleOutput() RoleOutput
	ToRoleOutputWithContext(ctx context.Context) RoleOutput
}

func (*Role) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (i *Role) ToRoleOutput() RoleOutput {
	return i.ToRoleOutputWithContext(context.Background())
}

func (i *Role) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleOutput)
}

// RoleArrayInput is an input type that accepts RoleArray and RoleArrayOutput values.
// You can construct a concrete instance of `RoleArrayInput` via:
//
//	RoleArray{ RoleArgs{...} }
type RoleArrayInput interface {
	pulumi.Input

	ToRoleArrayOutput() RoleArrayOutput
	ToRoleArrayOutputWithContext(context.Context) RoleArrayOutput
}

type RoleArray []RoleInput

func (RoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (i RoleArray) ToRoleArrayOutput() RoleArrayOutput {
	return i.ToRoleArrayOutputWithContext(context.Background())
}

func (i RoleArray) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleArrayOutput)
}

// RoleMapInput is an input type that accepts RoleMap and RoleMapOutput values.
// You can construct a concrete instance of `RoleMapInput` via:
//
//	RoleMap{ "key": RoleArgs{...} }
type RoleMapInput interface {
	pulumi.Input

	ToRoleMapOutput() RoleMapOutput
	ToRoleMapOutputWithContext(context.Context) RoleMapOutput
}

type RoleMap map[string]RoleInput

func (RoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (i RoleMap) ToRoleMapOutput() RoleMapOutput {
	return i.ToRoleMapOutputWithContext(context.Background())
}

func (i RoleMap) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoleMapOutput)
}

type RoleOutput struct{ *pulumi.OutputState }

func (RoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Role)(nil)).Elem()
}

func (o RoleOutput) ToRoleOutput() RoleOutput {
	return o
}

func (o RoleOutput) ToRoleOutputWithContext(ctx context.Context) RoleOutput {
	return o
}

// The description of the Role.
func (o RoleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the Role.
func (o RoleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The max session duration of the Role.
func (o RoleOutput) MaxSessionDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Role) pulumi.IntPtrOutput { return v.MaxSessionDuration }).(pulumi.IntPtrOutput)
}

// The name of the Role.
func (o RoleOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.RoleName }).(pulumi.StringOutput)
}

// The resource name of the Role.
func (o RoleOutput) Trn() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.Trn }).(pulumi.StringOutput)
}

// The trust policy document of the Role.
func (o RoleOutput) TrustPolicyDocument() pulumi.StringOutput {
	return o.ApplyT(func(v *Role) pulumi.StringOutput { return v.TrustPolicyDocument }).(pulumi.StringOutput)
}

type RoleArrayOutput struct{ *pulumi.OutputState }

func (RoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Role)(nil)).Elem()
}

func (o RoleArrayOutput) ToRoleArrayOutput() RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) ToRoleArrayOutputWithContext(ctx context.Context) RoleArrayOutput {
	return o
}

func (o RoleArrayOutput) Index(i pulumi.IntInput) RoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Role {
		return vs[0].([]*Role)[vs[1].(int)]
	}).(RoleOutput)
}

type RoleMapOutput struct{ *pulumi.OutputState }

func (RoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Role)(nil)).Elem()
}

func (o RoleMapOutput) ToRoleMapOutput() RoleMapOutput {
	return o
}

func (o RoleMapOutput) ToRoleMapOutputWithContext(ctx context.Context) RoleMapOutput {
	return o
}

func (o RoleMapOutput) MapIndex(k pulumi.StringInput) RoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Role {
		return vs[0].(map[string]*Role)[vs[1].(string)]
	}).(RoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoleInput)(nil)).Elem(), &Role{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleArrayInput)(nil)).Elem(), RoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RoleMapInput)(nil)).Elem(), RoleMap{})
	pulumi.RegisterOutputType(RoleOutput{})
	pulumi.RegisterOutputType(RoleArrayOutput{})
	pulumi.RegisterOutputType(RoleMapOutput{})
}
