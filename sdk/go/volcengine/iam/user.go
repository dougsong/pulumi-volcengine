// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage iam user
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
//			_, err := iam.NewUser(ctx, "foo", &iam.UserArgs{
//				Description: pulumi.String("test"),
//				DisplayName: pulumi.String("name"),
//				UserName:    pulumi.String("tf-test"),
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
// Iam user can be imported using the UserName, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:iam/user:User default user_name
//
// ```
type User struct {
	pulumi.CustomResourceState

	// The account id of the user.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The create date of the user.
	CreateDate pulumi.StringOutput `pulumi:"createDate"`
	// The description of the user.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The display name of the user.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The email of the user.
	Email pulumi.StringPtrOutput `pulumi:"email"`
	// The mobile phone of the user.
	MobilePhone pulumi.StringPtrOutput `pulumi:"mobilePhone"`
	// The trn of the user.
	Trn pulumi.StringOutput `pulumi:"trn"`
	// The update date of the user.
	UpdateDate pulumi.StringOutput `pulumi:"updateDate"`
	// The name of the user.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource User
	err := ctx.RegisterResource("volcengine:iam/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("volcengine:iam/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The account id of the user.
	AccountId *string `pulumi:"accountId"`
	// The create date of the user.
	CreateDate *string `pulumi:"createDate"`
	// The description of the user.
	Description *string `pulumi:"description"`
	// The display name of the user.
	DisplayName *string `pulumi:"displayName"`
	// The email of the user.
	Email *string `pulumi:"email"`
	// The mobile phone of the user.
	MobilePhone *string `pulumi:"mobilePhone"`
	// The trn of the user.
	Trn *string `pulumi:"trn"`
	// The update date of the user.
	UpdateDate *string `pulumi:"updateDate"`
	// The name of the user.
	UserName *string `pulumi:"userName"`
}

type UserState struct {
	// The account id of the user.
	AccountId pulumi.StringPtrInput
	// The create date of the user.
	CreateDate pulumi.StringPtrInput
	// The description of the user.
	Description pulumi.StringPtrInput
	// The display name of the user.
	DisplayName pulumi.StringPtrInput
	// The email of the user.
	Email pulumi.StringPtrInput
	// The mobile phone of the user.
	MobilePhone pulumi.StringPtrInput
	// The trn of the user.
	Trn pulumi.StringPtrInput
	// The update date of the user.
	UpdateDate pulumi.StringPtrInput
	// The name of the user.
	UserName pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// The description of the user.
	Description *string `pulumi:"description"`
	// The display name of the user.
	DisplayName *string `pulumi:"displayName"`
	// The email of the user.
	Email *string `pulumi:"email"`
	// The mobile phone of the user.
	MobilePhone *string `pulumi:"mobilePhone"`
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// The description of the user.
	Description pulumi.StringPtrInput
	// The display name of the user.
	DisplayName pulumi.StringPtrInput
	// The email of the user.
	Email pulumi.StringPtrInput
	// The mobile phone of the user.
	MobilePhone pulumi.StringPtrInput
	// The name of the user.
	UserName pulumi.StringInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//	UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//	UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

// The account id of the user.
func (o UserOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The create date of the user.
func (o UserOutput) CreateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.CreateDate }).(pulumi.StringOutput)
}

// The description of the user.
func (o UserOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the user.
func (o UserOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The email of the user.
func (o UserOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.Email }).(pulumi.StringPtrOutput)
}

// The mobile phone of the user.
func (o UserOutput) MobilePhone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *User) pulumi.StringPtrOutput { return v.MobilePhone }).(pulumi.StringPtrOutput)
}

// The trn of the user.
func (o UserOutput) Trn() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.Trn }).(pulumi.StringOutput)
}

// The update date of the user.
func (o UserOutput) UpdateDate() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UpdateDate }).(pulumi.StringOutput)
}

// The name of the user.
func (o UserOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *User) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
