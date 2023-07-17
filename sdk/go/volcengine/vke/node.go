// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage vke node
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vke"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.NewNode(ctx, "foo", &vke.NodeArgs{
//				AdditionalContainerStorageEnabled: pulumi.Bool(false),
//				ClusterId:                         pulumi.String("ccahbr0nqtofhiuuuajn0"),
//				ContainerStoragePath:              pulumi.String(""),
//				InstanceId:                        pulumi.String("i-ybrfa2vu2t7grbv8qa0j"),
//				KeepInstanceName:                  pulumi.Bool(true),
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
// VKE node can be imported using the node id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vke/node:Node default nc5t5epmrsf****
//
// ```
type Node struct {
	pulumi.CustomResourceState

	// The flag of additional container storage enable, the value is `true` or `false`.
	AdditionalContainerStorageEnabled pulumi.BoolPtrOutput `pulumi:"additionalContainerStorageEnabled"`
	// The client token.
	ClientToken pulumi.StringOutput `pulumi:"clientToken"`
	// The cluster id.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The container storage path.
	ContainerStoragePath pulumi.StringOutput `pulumi:"containerStoragePath"`
	// The ImageId of NodeConfig.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The initializeScript of Node.
	InitializeScript pulumi.StringPtrOutput `pulumi:"initializeScript"`
	// The instance id.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The flag of keep instance name, the value is `true` or `false`.
	KeepInstanceName pulumi.BoolPtrOutput `pulumi:"keepInstanceName"`
	// The KubernetesConfig of Node.
	KubernetesConfig NodeKubernetesConfigPtrOutput `pulumi:"kubernetesConfig"`
	// The node pool id.
	NodePoolId pulumi.StringOutput `pulumi:"nodePoolId"`
}

// NewNode registers a new resource with the given unique name, arguments, and options.
func NewNode(ctx *pulumi.Context,
	name string, args *NodeArgs, opts ...pulumi.ResourceOption) (*Node, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource Node
	err := ctx.RegisterResource("volcengine:vke/node:Node", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNode gets an existing Node resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNode(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeState, opts ...pulumi.ResourceOption) (*Node, error) {
	var resource Node
	err := ctx.ReadResource("volcengine:vke/node:Node", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Node resources.
type nodeState struct {
	// The flag of additional container storage enable, the value is `true` or `false`.
	AdditionalContainerStorageEnabled *bool `pulumi:"additionalContainerStorageEnabled"`
	// The client token.
	ClientToken *string `pulumi:"clientToken"`
	// The cluster id.
	ClusterId *string `pulumi:"clusterId"`
	// The container storage path.
	ContainerStoragePath *string `pulumi:"containerStoragePath"`
	// The ImageId of NodeConfig.
	ImageId *string `pulumi:"imageId"`
	// The initializeScript of Node.
	InitializeScript *string `pulumi:"initializeScript"`
	// The instance id.
	InstanceId *string `pulumi:"instanceId"`
	// The flag of keep instance name, the value is `true` or `false`.
	KeepInstanceName *bool `pulumi:"keepInstanceName"`
	// The KubernetesConfig of Node.
	KubernetesConfig *NodeKubernetesConfig `pulumi:"kubernetesConfig"`
	// The node pool id.
	NodePoolId *string `pulumi:"nodePoolId"`
}

type NodeState struct {
	// The flag of additional container storage enable, the value is `true` or `false`.
	AdditionalContainerStorageEnabled pulumi.BoolPtrInput
	// The client token.
	ClientToken pulumi.StringPtrInput
	// The cluster id.
	ClusterId pulumi.StringPtrInput
	// The container storage path.
	ContainerStoragePath pulumi.StringPtrInput
	// The ImageId of NodeConfig.
	ImageId pulumi.StringPtrInput
	// The initializeScript of Node.
	InitializeScript pulumi.StringPtrInput
	// The instance id.
	InstanceId pulumi.StringPtrInput
	// The flag of keep instance name, the value is `true` or `false`.
	KeepInstanceName pulumi.BoolPtrInput
	// The KubernetesConfig of Node.
	KubernetesConfig NodeKubernetesConfigPtrInput
	// The node pool id.
	NodePoolId pulumi.StringPtrInput
}

func (NodeState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeState)(nil)).Elem()
}

type nodeArgs struct {
	// The flag of additional container storage enable, the value is `true` or `false`.
	AdditionalContainerStorageEnabled *bool `pulumi:"additionalContainerStorageEnabled"`
	// The client token.
	ClientToken *string `pulumi:"clientToken"`
	// The cluster id.
	ClusterId string `pulumi:"clusterId"`
	// The container storage path.
	ContainerStoragePath *string `pulumi:"containerStoragePath"`
	// The ImageId of NodeConfig.
	ImageId *string `pulumi:"imageId"`
	// The initializeScript of Node.
	InitializeScript *string `pulumi:"initializeScript"`
	// The instance id.
	InstanceId string `pulumi:"instanceId"`
	// The flag of keep instance name, the value is `true` or `false`.
	KeepInstanceName *bool `pulumi:"keepInstanceName"`
	// The KubernetesConfig of Node.
	KubernetesConfig *NodeKubernetesConfig `pulumi:"kubernetesConfig"`
}

// The set of arguments for constructing a Node resource.
type NodeArgs struct {
	// The flag of additional container storage enable, the value is `true` or `false`.
	AdditionalContainerStorageEnabled pulumi.BoolPtrInput
	// The client token.
	ClientToken pulumi.StringPtrInput
	// The cluster id.
	ClusterId pulumi.StringInput
	// The container storage path.
	ContainerStoragePath pulumi.StringPtrInput
	// The ImageId of NodeConfig.
	ImageId pulumi.StringPtrInput
	// The initializeScript of Node.
	InitializeScript pulumi.StringPtrInput
	// The instance id.
	InstanceId pulumi.StringInput
	// The flag of keep instance name, the value is `true` or `false`.
	KeepInstanceName pulumi.BoolPtrInput
	// The KubernetesConfig of Node.
	KubernetesConfig NodeKubernetesConfigPtrInput
}

func (NodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeArgs)(nil)).Elem()
}

type NodeInput interface {
	pulumi.Input

	ToNodeOutput() NodeOutput
	ToNodeOutputWithContext(ctx context.Context) NodeOutput
}

func (*Node) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil)).Elem()
}

func (i *Node) ToNodeOutput() NodeOutput {
	return i.ToNodeOutputWithContext(context.Background())
}

func (i *Node) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeOutput)
}

// NodeArrayInput is an input type that accepts NodeArray and NodeArrayOutput values.
// You can construct a concrete instance of `NodeArrayInput` via:
//
//	NodeArray{ NodeArgs{...} }
type NodeArrayInput interface {
	pulumi.Input

	ToNodeArrayOutput() NodeArrayOutput
	ToNodeArrayOutputWithContext(context.Context) NodeArrayOutput
}

type NodeArray []NodeInput

func (NodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Node)(nil)).Elem()
}

func (i NodeArray) ToNodeArrayOutput() NodeArrayOutput {
	return i.ToNodeArrayOutputWithContext(context.Background())
}

func (i NodeArray) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeArrayOutput)
}

// NodeMapInput is an input type that accepts NodeMap and NodeMapOutput values.
// You can construct a concrete instance of `NodeMapInput` via:
//
//	NodeMap{ "key": NodeArgs{...} }
type NodeMapInput interface {
	pulumi.Input

	ToNodeMapOutput() NodeMapOutput
	ToNodeMapOutputWithContext(context.Context) NodeMapOutput
}

type NodeMap map[string]NodeInput

func (NodeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Node)(nil)).Elem()
}

func (i NodeMap) ToNodeMapOutput() NodeMapOutput {
	return i.ToNodeMapOutputWithContext(context.Background())
}

func (i NodeMap) ToNodeMapOutputWithContext(ctx context.Context) NodeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NodeMapOutput)
}

type NodeOutput struct{ *pulumi.OutputState }

func (NodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Node)(nil)).Elem()
}

func (o NodeOutput) ToNodeOutput() NodeOutput {
	return o
}

func (o NodeOutput) ToNodeOutputWithContext(ctx context.Context) NodeOutput {
	return o
}

// The flag of additional container storage enable, the value is `true` or `false`.
func (o NodeOutput) AdditionalContainerStorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Node) pulumi.BoolPtrOutput { return v.AdditionalContainerStorageEnabled }).(pulumi.BoolPtrOutput)
}

// The client token.
func (o NodeOutput) ClientToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.ClientToken }).(pulumi.StringOutput)
}

// The cluster id.
func (o NodeOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The container storage path.
func (o NodeOutput) ContainerStoragePath() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.ContainerStoragePath }).(pulumi.StringOutput)
}

// The ImageId of NodeConfig.
func (o NodeOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The initializeScript of Node.
func (o NodeOutput) InitializeScript() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Node) pulumi.StringPtrOutput { return v.InitializeScript }).(pulumi.StringPtrOutput)
}

// The instance id.
func (o NodeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The flag of keep instance name, the value is `true` or `false`.
func (o NodeOutput) KeepInstanceName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Node) pulumi.BoolPtrOutput { return v.KeepInstanceName }).(pulumi.BoolPtrOutput)
}

// The KubernetesConfig of Node.
func (o NodeOutput) KubernetesConfig() NodeKubernetesConfigPtrOutput {
	return o.ApplyT(func(v *Node) NodeKubernetesConfigPtrOutput { return v.KubernetesConfig }).(NodeKubernetesConfigPtrOutput)
}

// The node pool id.
func (o NodeOutput) NodePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *Node) pulumi.StringOutput { return v.NodePoolId }).(pulumi.StringOutput)
}

type NodeArrayOutput struct{ *pulumi.OutputState }

func (NodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Node)(nil)).Elem()
}

func (o NodeArrayOutput) ToNodeArrayOutput() NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) ToNodeArrayOutputWithContext(ctx context.Context) NodeArrayOutput {
	return o
}

func (o NodeArrayOutput) Index(i pulumi.IntInput) NodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Node {
		return vs[0].([]*Node)[vs[1].(int)]
	}).(NodeOutput)
}

type NodeMapOutput struct{ *pulumi.OutputState }

func (NodeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Node)(nil)).Elem()
}

func (o NodeMapOutput) ToNodeMapOutput() NodeMapOutput {
	return o
}

func (o NodeMapOutput) ToNodeMapOutputWithContext(ctx context.Context) NodeMapOutput {
	return o
}

func (o NodeMapOutput) MapIndex(k pulumi.StringInput) NodeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Node {
		return vs[0].(map[string]*Node)[vs[1].(string)]
	}).(NodeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NodeInput)(nil)).Elem(), &Node{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeArrayInput)(nil)).Elem(), NodeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NodeMapInput)(nil)).Elem(), NodeMap{})
	pulumi.RegisterOutputType(NodeOutput{})
	pulumi.RegisterOutputType(NodeArrayOutput{})
	pulumi.RegisterOutputType(NodeMapOutput{})
}
