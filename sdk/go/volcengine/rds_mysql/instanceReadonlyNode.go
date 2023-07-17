// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage rds mysql instance readonly node
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_mysql.NewInstanceReadonlyNode(ctx, "foo", &rds_mysql.InstanceReadonlyNodeArgs{
//				InstanceId: pulumi.String("mysql-b3fca7f571d6"),
//				NodeSpec:   pulumi.String("rds.mysql.1c2g"),
//				ZoneId:     pulumi.String("cn-guilin-b"),
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
// Rds Mysql Instance Readonly Node can be imported using the instance_id:node_id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode default mysql-72da4258c2c7:mysql-72da4258c2c7-r7f93
//
// ```
type InstanceReadonlyNode struct {
	pulumi.CustomResourceState

	// The RDS mysql instance id of the readonly node.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The id of the readonly node.
	NodeId pulumi.StringOutput `pulumi:"nodeId"`
	// The specification of readonly node.
	NodeSpec pulumi.StringOutput `pulumi:"nodeSpec"`
	// The available zone of readonly node.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstanceReadonlyNode registers a new resource with the given unique name, arguments, and options.
func NewInstanceReadonlyNode(ctx *pulumi.Context,
	name string, args *InstanceReadonlyNodeArgs, opts ...pulumi.ResourceOption) (*InstanceReadonlyNode, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.NodeSpec == nil {
		return nil, errors.New("invalid value for required argument 'NodeSpec'")
	}
	if args.ZoneId == nil {
		return nil, errors.New("invalid value for required argument 'ZoneId'")
	}
	var resource InstanceReadonlyNode
	err := ctx.RegisterResource("volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceReadonlyNode gets an existing InstanceReadonlyNode resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceReadonlyNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceReadonlyNodeState, opts ...pulumi.ResourceOption) (*InstanceReadonlyNode, error) {
	var resource InstanceReadonlyNode
	err := ctx.ReadResource("volcengine:rds_mysql/instanceReadonlyNode:InstanceReadonlyNode", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceReadonlyNode resources.
type instanceReadonlyNodeState struct {
	// The RDS mysql instance id of the readonly node.
	InstanceId *string `pulumi:"instanceId"`
	// The id of the readonly node.
	NodeId *string `pulumi:"nodeId"`
	// The specification of readonly node.
	NodeSpec *string `pulumi:"nodeSpec"`
	// The available zone of readonly node.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceReadonlyNodeState struct {
	// The RDS mysql instance id of the readonly node.
	InstanceId pulumi.StringPtrInput
	// The id of the readonly node.
	NodeId pulumi.StringPtrInput
	// The specification of readonly node.
	NodeSpec pulumi.StringPtrInput
	// The available zone of readonly node.
	ZoneId pulumi.StringPtrInput
}

func (InstanceReadonlyNodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceReadonlyNodeState)(nil)).Elem()
}

type instanceReadonlyNodeArgs struct {
	// The RDS mysql instance id of the readonly node.
	InstanceId string `pulumi:"instanceId"`
	// The specification of readonly node.
	NodeSpec string `pulumi:"nodeSpec"`
	// The available zone of readonly node.
	ZoneId string `pulumi:"zoneId"`
}

// The set of arguments for constructing a InstanceReadonlyNode resource.
type InstanceReadonlyNodeArgs struct {
	// The RDS mysql instance id of the readonly node.
	InstanceId pulumi.StringInput
	// The specification of readonly node.
	NodeSpec pulumi.StringInput
	// The available zone of readonly node.
	ZoneId pulumi.StringInput
}

func (InstanceReadonlyNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceReadonlyNodeArgs)(nil)).Elem()
}

type InstanceReadonlyNodeInput interface {
	pulumi.Input

	ToInstanceReadonlyNodeOutput() InstanceReadonlyNodeOutput
	ToInstanceReadonlyNodeOutputWithContext(ctx context.Context) InstanceReadonlyNodeOutput
}

func (*InstanceReadonlyNode) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceReadonlyNode)(nil)).Elem()
}

func (i *InstanceReadonlyNode) ToInstanceReadonlyNodeOutput() InstanceReadonlyNodeOutput {
	return i.ToInstanceReadonlyNodeOutputWithContext(context.Background())
}

func (i *InstanceReadonlyNode) ToInstanceReadonlyNodeOutputWithContext(ctx context.Context) InstanceReadonlyNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceReadonlyNodeOutput)
}

// InstanceReadonlyNodeArrayInput is an input type that accepts InstanceReadonlyNodeArray and InstanceReadonlyNodeArrayOutput values.
// You can construct a concrete instance of `InstanceReadonlyNodeArrayInput` via:
//
//	InstanceReadonlyNodeArray{ InstanceReadonlyNodeArgs{...} }
type InstanceReadonlyNodeArrayInput interface {
	pulumi.Input

	ToInstanceReadonlyNodeArrayOutput() InstanceReadonlyNodeArrayOutput
	ToInstanceReadonlyNodeArrayOutputWithContext(context.Context) InstanceReadonlyNodeArrayOutput
}

type InstanceReadonlyNodeArray []InstanceReadonlyNodeInput

func (InstanceReadonlyNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceReadonlyNode)(nil)).Elem()
}

func (i InstanceReadonlyNodeArray) ToInstanceReadonlyNodeArrayOutput() InstanceReadonlyNodeArrayOutput {
	return i.ToInstanceReadonlyNodeArrayOutputWithContext(context.Background())
}

func (i InstanceReadonlyNodeArray) ToInstanceReadonlyNodeArrayOutputWithContext(ctx context.Context) InstanceReadonlyNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceReadonlyNodeArrayOutput)
}

// InstanceReadonlyNodeMapInput is an input type that accepts InstanceReadonlyNodeMap and InstanceReadonlyNodeMapOutput values.
// You can construct a concrete instance of `InstanceReadonlyNodeMapInput` via:
//
//	InstanceReadonlyNodeMap{ "key": InstanceReadonlyNodeArgs{...} }
type InstanceReadonlyNodeMapInput interface {
	pulumi.Input

	ToInstanceReadonlyNodeMapOutput() InstanceReadonlyNodeMapOutput
	ToInstanceReadonlyNodeMapOutputWithContext(context.Context) InstanceReadonlyNodeMapOutput
}

type InstanceReadonlyNodeMap map[string]InstanceReadonlyNodeInput

func (InstanceReadonlyNodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceReadonlyNode)(nil)).Elem()
}

func (i InstanceReadonlyNodeMap) ToInstanceReadonlyNodeMapOutput() InstanceReadonlyNodeMapOutput {
	return i.ToInstanceReadonlyNodeMapOutputWithContext(context.Background())
}

func (i InstanceReadonlyNodeMap) ToInstanceReadonlyNodeMapOutputWithContext(ctx context.Context) InstanceReadonlyNodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceReadonlyNodeMapOutput)
}

type InstanceReadonlyNodeOutput struct{ *pulumi.OutputState }

func (InstanceReadonlyNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceReadonlyNode)(nil)).Elem()
}

func (o InstanceReadonlyNodeOutput) ToInstanceReadonlyNodeOutput() InstanceReadonlyNodeOutput {
	return o
}

func (o InstanceReadonlyNodeOutput) ToInstanceReadonlyNodeOutputWithContext(ctx context.Context) InstanceReadonlyNodeOutput {
	return o
}

// The RDS mysql instance id of the readonly node.
func (o InstanceReadonlyNodeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceReadonlyNode) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The id of the readonly node.
func (o InstanceReadonlyNodeOutput) NodeId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceReadonlyNode) pulumi.StringOutput { return v.NodeId }).(pulumi.StringOutput)
}

// The specification of readonly node.
func (o InstanceReadonlyNodeOutput) NodeSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceReadonlyNode) pulumi.StringOutput { return v.NodeSpec }).(pulumi.StringOutput)
}

// The available zone of readonly node.
func (o InstanceReadonlyNodeOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceReadonlyNode) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type InstanceReadonlyNodeArrayOutput struct{ *pulumi.OutputState }

func (InstanceReadonlyNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceReadonlyNode)(nil)).Elem()
}

func (o InstanceReadonlyNodeArrayOutput) ToInstanceReadonlyNodeArrayOutput() InstanceReadonlyNodeArrayOutput {
	return o
}

func (o InstanceReadonlyNodeArrayOutput) ToInstanceReadonlyNodeArrayOutputWithContext(ctx context.Context) InstanceReadonlyNodeArrayOutput {
	return o
}

func (o InstanceReadonlyNodeArrayOutput) Index(i pulumi.IntInput) InstanceReadonlyNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceReadonlyNode {
		return vs[0].([]*InstanceReadonlyNode)[vs[1].(int)]
	}).(InstanceReadonlyNodeOutput)
}

type InstanceReadonlyNodeMapOutput struct{ *pulumi.OutputState }

func (InstanceReadonlyNodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceReadonlyNode)(nil)).Elem()
}

func (o InstanceReadonlyNodeMapOutput) ToInstanceReadonlyNodeMapOutput() InstanceReadonlyNodeMapOutput {
	return o
}

func (o InstanceReadonlyNodeMapOutput) ToInstanceReadonlyNodeMapOutputWithContext(ctx context.Context) InstanceReadonlyNodeMapOutput {
	return o
}

func (o InstanceReadonlyNodeMapOutput) MapIndex(k pulumi.StringInput) InstanceReadonlyNodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceReadonlyNode {
		return vs[0].(map[string]*InstanceReadonlyNode)[vs[1].(string)]
	}).(InstanceReadonlyNodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceReadonlyNodeInput)(nil)).Elem(), &InstanceReadonlyNode{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceReadonlyNodeArrayInput)(nil)).Elem(), InstanceReadonlyNodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceReadonlyNodeMapInput)(nil)).Elem(), InstanceReadonlyNodeMap{})
	pulumi.RegisterOutputType(InstanceReadonlyNodeOutput{})
	pulumi.RegisterOutputType(InstanceReadonlyNodeArrayOutput{})
	pulumi.RegisterOutputType(InstanceReadonlyNodeMapOutput{})
}
