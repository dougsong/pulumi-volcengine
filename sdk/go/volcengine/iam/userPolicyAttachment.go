// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage iam user policy attachment
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/Iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		user, err := Iam.NewUser(ctx, "user", &Iam.UserArgs{
// 			UserName:    pulumi.String("TfTest"),
// 			Description: pulumi.String("test"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		policy, err := Iam.NewPolicy(ctx, "policy", &Iam.PolicyArgs{
// 			PolicyName:     pulumi.String("TerraformResourceTest1"),
// 			Description:    pulumi.String("created by terraform 1"),
// 			PolicyDocument: pulumi.String("{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = Iam.NewUserPolicyAttachment(ctx, "foo", &Iam.UserPolicyAttachmentArgs{
// 			UserName:   user.UserName,
// 			PolicyName: policy.PolicyName,
// 			PolicyType: policy.PolicyType,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Iam user policy attachment can be imported using the UserName:PolicyName:PolicyType, e.g.
//
// ```sh
//  $ pulumi import volcengine:Iam/userPolicyAttachment:UserPolicyAttachment default TerraformTestUser:TerraformTestPolicy:Custom
// ```
type UserPolicyAttachment struct {
	pulumi.CustomResourceState

	// The name of the Policy.
	PolicyName pulumi.StringOutput `pulumi:"policyName"`
	// The type of the Policy.
	PolicyType pulumi.StringOutput `pulumi:"policyType"`
	// The name of the user.
	UserName pulumi.StringOutput `pulumi:"userName"`
}

// NewUserPolicyAttachment registers a new resource with the given unique name, arguments, and options.
func NewUserPolicyAttachment(ctx *pulumi.Context,
	name string, args *UserPolicyAttachmentArgs, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PolicyName == nil {
		return nil, errors.New("invalid value for required argument 'PolicyName'")
	}
	if args.PolicyType == nil {
		return nil, errors.New("invalid value for required argument 'PolicyType'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	var resource UserPolicyAttachment
	err := ctx.RegisterResource("volcengine:Iam/userPolicyAttachment:UserPolicyAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUserPolicyAttachment gets an existing UserPolicyAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUserPolicyAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserPolicyAttachmentState, opts ...pulumi.ResourceOption) (*UserPolicyAttachment, error) {
	var resource UserPolicyAttachment
	err := ctx.ReadResource("volcengine:Iam/userPolicyAttachment:UserPolicyAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering UserPolicyAttachment resources.
type userPolicyAttachmentState struct {
	// The name of the Policy.
	PolicyName *string `pulumi:"policyName"`
	// The type of the Policy.
	PolicyType *string `pulumi:"policyType"`
	// The name of the user.
	UserName *string `pulumi:"userName"`
}

type UserPolicyAttachmentState struct {
	// The name of the Policy.
	PolicyName pulumi.StringPtrInput
	// The type of the Policy.
	PolicyType pulumi.StringPtrInput
	// The name of the user.
	UserName pulumi.StringPtrInput
}

func (UserPolicyAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentState)(nil)).Elem()
}

type userPolicyAttachmentArgs struct {
	// The name of the Policy.
	PolicyName string `pulumi:"policyName"`
	// The type of the Policy.
	PolicyType string `pulumi:"policyType"`
	// The name of the user.
	UserName string `pulumi:"userName"`
}

// The set of arguments for constructing a UserPolicyAttachment resource.
type UserPolicyAttachmentArgs struct {
	// The name of the Policy.
	PolicyName pulumi.StringInput
	// The type of the Policy.
	PolicyType pulumi.StringInput
	// The name of the user.
	UserName pulumi.StringInput
}

func (UserPolicyAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userPolicyAttachmentArgs)(nil)).Elem()
}

type UserPolicyAttachmentInput interface {
	pulumi.Input

	ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput
	ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput
}

func (*UserPolicyAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (i *UserPolicyAttachment) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return i.ToUserPolicyAttachmentOutputWithContext(context.Background())
}

func (i *UserPolicyAttachment) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentOutput)
}

// UserPolicyAttachmentArrayInput is an input type that accepts UserPolicyAttachmentArray and UserPolicyAttachmentArrayOutput values.
// You can construct a concrete instance of `UserPolicyAttachmentArrayInput` via:
//
//          UserPolicyAttachmentArray{ UserPolicyAttachmentArgs{...} }
type UserPolicyAttachmentArrayInput interface {
	pulumi.Input

	ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput
	ToUserPolicyAttachmentArrayOutputWithContext(context.Context) UserPolicyAttachmentArrayOutput
}

type UserPolicyAttachmentArray []UserPolicyAttachmentInput

func (UserPolicyAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachmentArray) ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput {
	return i.ToUserPolicyAttachmentArrayOutputWithContext(context.Background())
}

func (i UserPolicyAttachmentArray) ToUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) UserPolicyAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentArrayOutput)
}

// UserPolicyAttachmentMapInput is an input type that accepts UserPolicyAttachmentMap and UserPolicyAttachmentMapOutput values.
// You can construct a concrete instance of `UserPolicyAttachmentMapInput` via:
//
//          UserPolicyAttachmentMap{ "key": UserPolicyAttachmentArgs{...} }
type UserPolicyAttachmentMapInput interface {
	pulumi.Input

	ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput
	ToUserPolicyAttachmentMapOutputWithContext(context.Context) UserPolicyAttachmentMapOutput
}

type UserPolicyAttachmentMap map[string]UserPolicyAttachmentInput

func (UserPolicyAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicyAttachment)(nil)).Elem()
}

func (i UserPolicyAttachmentMap) ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput {
	return i.ToUserPolicyAttachmentMapOutputWithContext(context.Background())
}

func (i UserPolicyAttachmentMap) ToUserPolicyAttachmentMapOutputWithContext(ctx context.Context) UserPolicyAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserPolicyAttachmentMapOutput)
}

type UserPolicyAttachmentOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutput() UserPolicyAttachmentOutput {
	return o
}

func (o UserPolicyAttachmentOutput) ToUserPolicyAttachmentOutputWithContext(ctx context.Context) UserPolicyAttachmentOutput {
	return o
}

// The name of the Policy.
func (o UserPolicyAttachmentOutput) PolicyName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyName }).(pulumi.StringOutput)
}

// The type of the Policy.
func (o UserPolicyAttachmentOutput) PolicyType() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.PolicyType }).(pulumi.StringOutput)
}

// The name of the user.
func (o UserPolicyAttachmentOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v *UserPolicyAttachment) pulumi.StringOutput { return v.UserName }).(pulumi.StringOutput)
}

type UserPolicyAttachmentArrayOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentArrayOutput) ToUserPolicyAttachmentArrayOutput() UserPolicyAttachmentArrayOutput {
	return o
}

func (o UserPolicyAttachmentArrayOutput) ToUserPolicyAttachmentArrayOutputWithContext(ctx context.Context) UserPolicyAttachmentArrayOutput {
	return o
}

func (o UserPolicyAttachmentArrayOutput) Index(i pulumi.IntInput) UserPolicyAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *UserPolicyAttachment {
		return vs[0].([]*UserPolicyAttachment)[vs[1].(int)]
	}).(UserPolicyAttachmentOutput)
}

type UserPolicyAttachmentMapOutput struct{ *pulumi.OutputState }

func (UserPolicyAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*UserPolicyAttachment)(nil)).Elem()
}

func (o UserPolicyAttachmentMapOutput) ToUserPolicyAttachmentMapOutput() UserPolicyAttachmentMapOutput {
	return o
}

func (o UserPolicyAttachmentMapOutput) ToUserPolicyAttachmentMapOutputWithContext(ctx context.Context) UserPolicyAttachmentMapOutput {
	return o
}

func (o UserPolicyAttachmentMapOutput) MapIndex(k pulumi.StringInput) UserPolicyAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *UserPolicyAttachment {
		return vs[0].(map[string]*UserPolicyAttachment)[vs[1].(string)]
	}).(UserPolicyAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentInput)(nil)).Elem(), &UserPolicyAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentArrayInput)(nil)).Elem(), UserPolicyAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserPolicyAttachmentMapInput)(nil)).Elem(), UserPolicyAttachmentMap{})
	pulumi.RegisterOutputType(UserPolicyAttachmentOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentArrayOutput{})
	pulumi.RegisterOutputType(UserPolicyAttachmentMapOutput{})
}
