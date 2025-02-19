// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage iam login profile
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
//			_, err := iam.NewLoginProfile(ctx, "foo", &iam.LoginProfileArgs{
//				LoginAllowed:          pulumi.Bool(true),
//				Password:              pulumi.String("******"),
//				PasswordResetRequired: pulumi.Bool(false),
//				UserName:              pulumi.String("tf-test"),
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
// Login profile can be imported using the UserName, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:iam/loginProfile:LoginProfile default user_name
//
// ```
type LoginProfile struct {
	pulumi.CustomResourceState

	// The flag of login allowed.
	LoginAllowed pulumi.BoolPtrOutput `pulumi:"loginAllowed"`
	// The password.
	Password pulumi.StringOutput `pulumi:"password"`
	// Is required reset password when next time login in.
	PasswordResetRequired pulumi.BoolPtrOutput `pulumi:"passwordResetRequired"`
	// The user name.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewLoginProfile registers a new resource with the given unique name, arguments, and options.
func NewLoginProfile(ctx *pulumi.Context,
	name string, args *LoginProfileArgs, opts ...pulumi.ResourceOption) (*LoginProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = pkgResourceDefaultOpts(opts)
	var resource LoginProfile
	err := ctx.RegisterResource("volcengine:iam/loginProfile:LoginProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLoginProfile gets an existing LoginProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLoginProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LoginProfileState, opts ...pulumi.ResourceOption) (*LoginProfile, error) {
	var resource LoginProfile
	err := ctx.ReadResource("volcengine:iam/loginProfile:LoginProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LoginProfile resources.
type loginProfileState struct {
	// The flag of login allowed.
	LoginAllowed *bool `pulumi:"loginAllowed"`
	// The password.
	Password *string `pulumi:"password"`
	// Is required reset password when next time login in.
	PasswordResetRequired *bool `pulumi:"passwordResetRequired"`
	// The user name.
	UserName *string `pulumi:"userName"`
}

type LoginProfileState struct {
	// The flag of login allowed.
	LoginAllowed pulumi.BoolPtrInput
	// The password.
	Password pulumi.StringPtrInput
	// Is required reset password when next time login in.
	PasswordResetRequired pulumi.BoolPtrInput
	// The user name.
	UserName pulumi.StringPtrInput
}

func (LoginProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*loginProfileState)(nil)).Elem()
}

type loginProfileArgs struct {
	// The flag of login allowed.
	LoginAllowed *bool `pulumi:"loginAllowed"`
	// The password.
	Password string `pulumi:"password"`
	// Is required reset password when next time login in.
	PasswordResetRequired *bool `pulumi:"passwordResetRequired"`
	// The user name.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a LoginProfile resource.
type LoginProfileArgs struct {
	// The flag of login allowed.
	LoginAllowed pulumi.BoolPtrInput
	// The password.
	Password pulumi.StringInput
	// Is required reset password when next time login in.
	PasswordResetRequired pulumi.BoolPtrInput
	// The user name.
	UserName pulumi.StringInput
}

func (LoginProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*loginProfileArgs)(nil)).Elem()
}

type LoginProfileInput interface {
	pulumi.Input

	ToLoginProfileOutput() LoginProfileOutput
	ToLoginProfileOutputWithContext(ctx context.Context) LoginProfileOutput
}

func (*LoginProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginProfile)(nil)).Elem()
}

func (i *LoginProfile) ToLoginProfileOutput() LoginProfileOutput {
	return i.ToLoginProfileOutputWithContext(context.Background())
}

func (i *LoginProfile) ToLoginProfileOutputWithContext(ctx context.Context) LoginProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginProfileOutput)
}

// LoginProfileArrayInput is an input type that accepts LoginProfileArray and LoginProfileArrayOutput values.
// You can construct a concrete instance of `LoginProfileArrayInput` via:
//
//	LoginProfileArray{ LoginProfileArgs{...} }
type LoginProfileArrayInput interface {
	pulumi.Input

	ToLoginProfileArrayOutput() LoginProfileArrayOutput
	ToLoginProfileArrayOutputWithContext(context.Context) LoginProfileArrayOutput
}

type LoginProfileArray []LoginProfileInput

func (LoginProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoginProfile)(nil)).Elem()
}

func (i LoginProfileArray) ToLoginProfileArrayOutput() LoginProfileArrayOutput {
	return i.ToLoginProfileArrayOutputWithContext(context.Background())
}

func (i LoginProfileArray) ToLoginProfileArrayOutputWithContext(ctx context.Context) LoginProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginProfileArrayOutput)
}

// LoginProfileMapInput is an input type that accepts LoginProfileMap and LoginProfileMapOutput values.
// You can construct a concrete instance of `LoginProfileMapInput` via:
//
//	LoginProfileMap{ "key": LoginProfileArgs{...} }
type LoginProfileMapInput interface {
	pulumi.Input

	ToLoginProfileMapOutput() LoginProfileMapOutput
	ToLoginProfileMapOutputWithContext(context.Context) LoginProfileMapOutput
}

type LoginProfileMap map[string]LoginProfileInput

func (LoginProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoginProfile)(nil)).Elem()
}

func (i LoginProfileMap) ToLoginProfileMapOutput() LoginProfileMapOutput {
	return i.ToLoginProfileMapOutputWithContext(context.Background())
}

func (i LoginProfileMap) ToLoginProfileMapOutputWithContext(ctx context.Context) LoginProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginProfileMapOutput)
}

type LoginProfileOutput struct{ *pulumi.OutputState }

func (LoginProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginProfile)(nil)).Elem()
}

func (o LoginProfileOutput) ToLoginProfileOutput() LoginProfileOutput {
	return o
}

func (o LoginProfileOutput) ToLoginProfileOutputWithContext(ctx context.Context) LoginProfileOutput {
	return o
}

// The flag of login allowed.
func (o LoginProfileOutput) LoginAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoginProfile) pulumi.BoolPtrOutput { return v.LoginAllowed }).(pulumi.BoolPtrOutput)
}

// The password.
func (o LoginProfileOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *LoginProfile) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Is required reset password when next time login in.
func (o LoginProfileOutput) PasswordResetRequired() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LoginProfile) pulumi.BoolPtrOutput { return v.PasswordResetRequired }).(pulumi.BoolPtrOutput)
}

// The user name.
func (o LoginProfileOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *LoginProfile) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type LoginProfileArrayOutput struct{ *pulumi.OutputState }

func (LoginProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LoginProfile)(nil)).Elem()
}

func (o LoginProfileArrayOutput) ToLoginProfileArrayOutput() LoginProfileArrayOutput {
	return o
}

func (o LoginProfileArrayOutput) ToLoginProfileArrayOutputWithContext(ctx context.Context) LoginProfileArrayOutput {
	return o
}

func (o LoginProfileArrayOutput) Index(i pulumi.IntInput) LoginProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LoginProfile {
		return vs[0].([]*LoginProfile)[vs[1].(int)]
	}).(LoginProfileOutput)
}

type LoginProfileMapOutput struct{ *pulumi.OutputState }

func (LoginProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LoginProfile)(nil)).Elem()
}

func (o LoginProfileMapOutput) ToLoginProfileMapOutput() LoginProfileMapOutput {
	return o
}

func (o LoginProfileMapOutput) ToLoginProfileMapOutputWithContext(ctx context.Context) LoginProfileMapOutput {
	return o
}

func (o LoginProfileMapOutput) MapIndex(k pulumi.StringInput) LoginProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LoginProfile {
		return vs[0].(map[string]*LoginProfile)[vs[1].(string)]
	}).(LoginProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LoginProfileInput)(nil)).Elem(), &LoginProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoginProfileArrayInput)(nil)).Elem(), LoginProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LoginProfileMapInput)(nil)).Elem(), LoginProfileMap{})
	pulumi.RegisterOutputType(LoginProfileOutput{})
	pulumi.RegisterOutputType(LoginProfileArrayOutput{})
	pulumi.RegisterOutputType(LoginProfileMapOutput{})
}
