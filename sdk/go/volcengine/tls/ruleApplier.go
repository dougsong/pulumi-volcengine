// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tls rule applier
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/tls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.NewRuleApplier(ctx, "foo", &tls.RuleApplierArgs{
//				HostGroupId: pulumi.String("a2a9e8c5-9835-434f-b866-2c1cfa82887d"),
//				RuleId:      pulumi.String("25104b0f-28b7-4a5c-8339-7f9c431d77c8"),
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
// tls rule applier can be imported using the rule id and host group id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tls/ruleApplier:RuleApplier default fa************:bcb*******
//
// ```
type RuleApplier struct {
	pulumi.CustomResourceState

	// The id of the host group.
	HostGroupId pulumi.StringOutput `pulumi:"hostGroupId"`
	// The id of the rule.
	RuleId pulumi.StringOutput `pulumi:"ruleId"`
}

// NewRuleApplier registers a new resource with the given unique name, arguments, and options.
func NewRuleApplier(ctx *pulumi.Context,
	name string, args *RuleApplierArgs, opts ...pulumi.ResourceOption) (*RuleApplier, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostGroupId == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupId'")
	}
	if args.RuleId == nil {
		return nil, errors.New("invalid value for required argument 'RuleId'")
	}
	var resource RuleApplier
	err := ctx.RegisterResource("volcengine:tls/ruleApplier:RuleApplier", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRuleApplier gets an existing RuleApplier resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRuleApplier(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RuleApplierState, opts ...pulumi.ResourceOption) (*RuleApplier, error) {
	var resource RuleApplier
	err := ctx.ReadResource("volcengine:tls/ruleApplier:RuleApplier", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RuleApplier resources.
type ruleApplierState struct {
	// The id of the host group.
	HostGroupId *string `pulumi:"hostGroupId"`
	// The id of the rule.
	RuleId *string `pulumi:"ruleId"`
}

type RuleApplierState struct {
	// The id of the host group.
	HostGroupId pulumi.StringPtrInput
	// The id of the rule.
	RuleId pulumi.StringPtrInput
}

func (RuleApplierState) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleApplierState)(nil)).Elem()
}

type ruleApplierArgs struct {
	// The id of the host group.
	HostGroupId string `pulumi:"hostGroupId"`
	// The id of the rule.
	RuleId string `pulumi:"ruleId"`
}

// The set of arguments for constructing a RuleApplier resource.
type RuleApplierArgs struct {
	// The id of the host group.
	HostGroupId pulumi.StringInput
	// The id of the rule.
	RuleId pulumi.StringInput
}

func (RuleApplierArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ruleApplierArgs)(nil)).Elem()
}

type RuleApplierInput interface {
	pulumi.Input

	ToRuleApplierOutput() RuleApplierOutput
	ToRuleApplierOutputWithContext(ctx context.Context) RuleApplierOutput
}

func (*RuleApplier) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleApplier)(nil)).Elem()
}

func (i *RuleApplier) ToRuleApplierOutput() RuleApplierOutput {
	return i.ToRuleApplierOutputWithContext(context.Background())
}

func (i *RuleApplier) ToRuleApplierOutputWithContext(ctx context.Context) RuleApplierOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleApplierOutput)
}

// RuleApplierArrayInput is an input type that accepts RuleApplierArray and RuleApplierArrayOutput values.
// You can construct a concrete instance of `RuleApplierArrayInput` via:
//
//	RuleApplierArray{ RuleApplierArgs{...} }
type RuleApplierArrayInput interface {
	pulumi.Input

	ToRuleApplierArrayOutput() RuleApplierArrayOutput
	ToRuleApplierArrayOutputWithContext(context.Context) RuleApplierArrayOutput
}

type RuleApplierArray []RuleApplierInput

func (RuleApplierArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleApplier)(nil)).Elem()
}

func (i RuleApplierArray) ToRuleApplierArrayOutput() RuleApplierArrayOutput {
	return i.ToRuleApplierArrayOutputWithContext(context.Background())
}

func (i RuleApplierArray) ToRuleApplierArrayOutputWithContext(ctx context.Context) RuleApplierArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleApplierArrayOutput)
}

// RuleApplierMapInput is an input type that accepts RuleApplierMap and RuleApplierMapOutput values.
// You can construct a concrete instance of `RuleApplierMapInput` via:
//
//	RuleApplierMap{ "key": RuleApplierArgs{...} }
type RuleApplierMapInput interface {
	pulumi.Input

	ToRuleApplierMapOutput() RuleApplierMapOutput
	ToRuleApplierMapOutputWithContext(context.Context) RuleApplierMapOutput
}

type RuleApplierMap map[string]RuleApplierInput

func (RuleApplierMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleApplier)(nil)).Elem()
}

func (i RuleApplierMap) ToRuleApplierMapOutput() RuleApplierMapOutput {
	return i.ToRuleApplierMapOutputWithContext(context.Background())
}

func (i RuleApplierMap) ToRuleApplierMapOutputWithContext(ctx context.Context) RuleApplierMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RuleApplierMapOutput)
}

type RuleApplierOutput struct{ *pulumi.OutputState }

func (RuleApplierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RuleApplier)(nil)).Elem()
}

func (o RuleApplierOutput) ToRuleApplierOutput() RuleApplierOutput {
	return o
}

func (o RuleApplierOutput) ToRuleApplierOutputWithContext(ctx context.Context) RuleApplierOutput {
	return o
}

// The id of the host group.
func (o RuleApplierOutput) HostGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleApplier) pulumi.StringOutput { return v.HostGroupId }).(pulumi.StringOutput)
}

// The id of the rule.
func (o RuleApplierOutput) RuleId() pulumi.StringOutput {
	return o.ApplyT(func(v *RuleApplier) pulumi.StringOutput { return v.RuleId }).(pulumi.StringOutput)
}

type RuleApplierArrayOutput struct{ *pulumi.OutputState }

func (RuleApplierArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RuleApplier)(nil)).Elem()
}

func (o RuleApplierArrayOutput) ToRuleApplierArrayOutput() RuleApplierArrayOutput {
	return o
}

func (o RuleApplierArrayOutput) ToRuleApplierArrayOutputWithContext(ctx context.Context) RuleApplierArrayOutput {
	return o
}

func (o RuleApplierArrayOutput) Index(i pulumi.IntInput) RuleApplierOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RuleApplier {
		return vs[0].([]*RuleApplier)[vs[1].(int)]
	}).(RuleApplierOutput)
}

type RuleApplierMapOutput struct{ *pulumi.OutputState }

func (RuleApplierMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RuleApplier)(nil)).Elem()
}

func (o RuleApplierMapOutput) ToRuleApplierMapOutput() RuleApplierMapOutput {
	return o
}

func (o RuleApplierMapOutput) ToRuleApplierMapOutputWithContext(ctx context.Context) RuleApplierMapOutput {
	return o
}

func (o RuleApplierMapOutput) MapIndex(k pulumi.StringInput) RuleApplierOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RuleApplier {
		return vs[0].(map[string]*RuleApplier)[vs[1].(string)]
	}).(RuleApplierOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RuleApplierInput)(nil)).Elem(), &RuleApplier{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleApplierArrayInput)(nil)).Elem(), RuleApplierArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RuleApplierMapInput)(nil)).Elem(), RuleApplierMap{})
	pulumi.RegisterOutputType(RuleApplierOutput{})
	pulumi.RegisterOutputType(RuleApplierArrayOutput{})
	pulumi.RegisterOutputType(RuleApplierMapOutput{})
}
