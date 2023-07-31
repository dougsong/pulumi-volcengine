// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage vke default node pool batch attach
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vke"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vke"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.NewDefaultNodePoolBatchAttach(ctx, "default", &vke.DefaultNodePoolBatchAttachArgs{
//				ClusterId:         pulumi.String("ccc2umdnqtoflv91lqtq0"),
//				DefaultNodePoolId: pulumi.String("11111"),
//				Instances: vke.DefaultNodePoolBatchAttachInstanceArray{
//					&vke.DefaultNodePoolBatchAttachInstanceArgs{
//						AdditionalContainerStorageEnabled: pulumi.Bool(false),
//						InstanceId:                        pulumi.String("i-ybvza90ohwexzk8emaa3"),
//						KeepInstanceName:                  pulumi.Bool(false),
//					},
//					&vke.DefaultNodePoolBatchAttachInstanceArgs{
//						AdditionalContainerStorageEnabled: pulumi.Bool(true),
//						ContainerStoragePath:              pulumi.String("/"),
//						InstanceId:                        pulumi.String("i-ybvza90ohxexzkm4zihf"),
//						KeepInstanceName:                  pulumi.Bool(false),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DefaultNodePoolBatchAttach struct {
	pulumi.CustomResourceState

	// The ClusterId of NodePool.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The default NodePool ID.
	DefaultNodePoolId pulumi.StringOutput `pulumi:"defaultNodePoolId"`
	// The ECS InstanceIds add to NodePool.
	Instances DefaultNodePoolBatchAttachInstanceArrayOutput `pulumi:"instances"`
	// Is import of the DefaultNodePool. It only works when imported, set to true.
	IsImport pulumi.BoolOutput `pulumi:"isImport"`
	// The KubernetesConfig of NodeConfig.
	KubernetesConfigs DefaultNodePoolBatchAttachKubernetesConfigArrayOutput `pulumi:"kubernetesConfigs"`
	// The Config of NodePool.
	NodeConfigs DefaultNodePoolBatchAttachNodeConfigArrayOutput `pulumi:"nodeConfigs"`
	// Tags.
	Tags DefaultNodePoolBatchAttachTagArrayOutput `pulumi:"tags"`
}

// NewDefaultNodePoolBatchAttach registers a new resource with the given unique name, arguments, and options.
func NewDefaultNodePoolBatchAttach(ctx *pulumi.Context,
	name string, args *DefaultNodePoolBatchAttachArgs, opts ...pulumi.ResourceOption) (*DefaultNodePoolBatchAttach, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.DefaultNodePoolId == nil {
		return nil, errors.New("invalid value for required argument 'DefaultNodePoolId'")
	}
	var resource DefaultNodePoolBatchAttach
	err := ctx.RegisterResource("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultNodePoolBatchAttach gets an existing DefaultNodePoolBatchAttach resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNodePoolBatchAttach(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultNodePoolBatchAttachState, opts ...pulumi.ResourceOption) (*DefaultNodePoolBatchAttach, error) {
	var resource DefaultNodePoolBatchAttach
	err := ctx.ReadResource("volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultNodePoolBatchAttach resources.
type defaultNodePoolBatchAttachState struct {
	// The ClusterId of NodePool.
	ClusterId *string `pulumi:"clusterId"`
	// The default NodePool ID.
	DefaultNodePoolId *string `pulumi:"defaultNodePoolId"`
	// The ECS InstanceIds add to NodePool.
	Instances []DefaultNodePoolBatchAttachInstance `pulumi:"instances"`
	// Is import of the DefaultNodePool. It only works when imported, set to true.
	IsImport *bool `pulumi:"isImport"`
	// The KubernetesConfig of NodeConfig.
	KubernetesConfigs []DefaultNodePoolBatchAttachKubernetesConfig `pulumi:"kubernetesConfigs"`
	// The Config of NodePool.
	NodeConfigs []DefaultNodePoolBatchAttachNodeConfig `pulumi:"nodeConfigs"`
	// Tags.
	Tags []DefaultNodePoolBatchAttachTag `pulumi:"tags"`
}

type DefaultNodePoolBatchAttachState struct {
	// The ClusterId of NodePool.
	ClusterId pulumi.StringPtrInput
	// The default NodePool ID.
	DefaultNodePoolId pulumi.StringPtrInput
	// The ECS InstanceIds add to NodePool.
	Instances DefaultNodePoolBatchAttachInstanceArrayInput
	// Is import of the DefaultNodePool. It only works when imported, set to true.
	IsImport pulumi.BoolPtrInput
	// The KubernetesConfig of NodeConfig.
	KubernetesConfigs DefaultNodePoolBatchAttachKubernetesConfigArrayInput
	// The Config of NodePool.
	NodeConfigs DefaultNodePoolBatchAttachNodeConfigArrayInput
	// Tags.
	Tags DefaultNodePoolBatchAttachTagArrayInput
}

func (DefaultNodePoolBatchAttachState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNodePoolBatchAttachState)(nil)).Elem()
}

type defaultNodePoolBatchAttachArgs struct {
	// The ClusterId of NodePool.
	ClusterId string `pulumi:"clusterId"`
	// The default NodePool ID.
	DefaultNodePoolId string `pulumi:"defaultNodePoolId"`
	// The ECS InstanceIds add to NodePool.
	Instances []DefaultNodePoolBatchAttachInstance `pulumi:"instances"`
}

// The set of arguments for constructing a DefaultNodePoolBatchAttach resource.
type DefaultNodePoolBatchAttachArgs struct {
	// The ClusterId of NodePool.
	ClusterId pulumi.StringInput
	// The default NodePool ID.
	DefaultNodePoolId pulumi.StringInput
	// The ECS InstanceIds add to NodePool.
	Instances DefaultNodePoolBatchAttachInstanceArrayInput
}

func (DefaultNodePoolBatchAttachArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNodePoolBatchAttachArgs)(nil)).Elem()
}

type DefaultNodePoolBatchAttachInput interface {
	pulumi.Input

	ToDefaultNodePoolBatchAttachOutput() DefaultNodePoolBatchAttachOutput
	ToDefaultNodePoolBatchAttachOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachOutput
}

func (*DefaultNodePoolBatchAttach) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNodePoolBatchAttach)(nil)).Elem()
}

func (i *DefaultNodePoolBatchAttach) ToDefaultNodePoolBatchAttachOutput() DefaultNodePoolBatchAttachOutput {
	return i.ToDefaultNodePoolBatchAttachOutputWithContext(context.Background())
}

func (i *DefaultNodePoolBatchAttach) ToDefaultNodePoolBatchAttachOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNodePoolBatchAttachOutput)
}

// DefaultNodePoolBatchAttachArrayInput is an input type that accepts DefaultNodePoolBatchAttachArray and DefaultNodePoolBatchAttachArrayOutput values.
// You can construct a concrete instance of `DefaultNodePoolBatchAttachArrayInput` via:
//
//	DefaultNodePoolBatchAttachArray{ DefaultNodePoolBatchAttachArgs{...} }
type DefaultNodePoolBatchAttachArrayInput interface {
	pulumi.Input

	ToDefaultNodePoolBatchAttachArrayOutput() DefaultNodePoolBatchAttachArrayOutput
	ToDefaultNodePoolBatchAttachArrayOutputWithContext(context.Context) DefaultNodePoolBatchAttachArrayOutput
}

type DefaultNodePoolBatchAttachArray []DefaultNodePoolBatchAttachInput

func (DefaultNodePoolBatchAttachArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNodePoolBatchAttach)(nil)).Elem()
}

func (i DefaultNodePoolBatchAttachArray) ToDefaultNodePoolBatchAttachArrayOutput() DefaultNodePoolBatchAttachArrayOutput {
	return i.ToDefaultNodePoolBatchAttachArrayOutputWithContext(context.Background())
}

func (i DefaultNodePoolBatchAttachArray) ToDefaultNodePoolBatchAttachArrayOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNodePoolBatchAttachArrayOutput)
}

// DefaultNodePoolBatchAttachMapInput is an input type that accepts DefaultNodePoolBatchAttachMap and DefaultNodePoolBatchAttachMapOutput values.
// You can construct a concrete instance of `DefaultNodePoolBatchAttachMapInput` via:
//
//	DefaultNodePoolBatchAttachMap{ "key": DefaultNodePoolBatchAttachArgs{...} }
type DefaultNodePoolBatchAttachMapInput interface {
	pulumi.Input

	ToDefaultNodePoolBatchAttachMapOutput() DefaultNodePoolBatchAttachMapOutput
	ToDefaultNodePoolBatchAttachMapOutputWithContext(context.Context) DefaultNodePoolBatchAttachMapOutput
}

type DefaultNodePoolBatchAttachMap map[string]DefaultNodePoolBatchAttachInput

func (DefaultNodePoolBatchAttachMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNodePoolBatchAttach)(nil)).Elem()
}

func (i DefaultNodePoolBatchAttachMap) ToDefaultNodePoolBatchAttachMapOutput() DefaultNodePoolBatchAttachMapOutput {
	return i.ToDefaultNodePoolBatchAttachMapOutputWithContext(context.Background())
}

func (i DefaultNodePoolBatchAttachMap) ToDefaultNodePoolBatchAttachMapOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNodePoolBatchAttachMapOutput)
}

type DefaultNodePoolBatchAttachOutput struct{ *pulumi.OutputState }

func (DefaultNodePoolBatchAttachOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNodePoolBatchAttach)(nil)).Elem()
}

func (o DefaultNodePoolBatchAttachOutput) ToDefaultNodePoolBatchAttachOutput() DefaultNodePoolBatchAttachOutput {
	return o
}

func (o DefaultNodePoolBatchAttachOutput) ToDefaultNodePoolBatchAttachOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachOutput {
	return o
}

// The ClusterId of NodePool.
func (o DefaultNodePoolBatchAttachOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The default NodePool ID.
func (o DefaultNodePoolBatchAttachOutput) DefaultNodePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) pulumi.StringOutput { return v.DefaultNodePoolId }).(pulumi.StringOutput)
}

// The ECS InstanceIds add to NodePool.
func (o DefaultNodePoolBatchAttachOutput) Instances() DefaultNodePoolBatchAttachInstanceArrayOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) DefaultNodePoolBatchAttachInstanceArrayOutput { return v.Instances }).(DefaultNodePoolBatchAttachInstanceArrayOutput)
}

// Is import of the DefaultNodePool. It only works when imported, set to true.
func (o DefaultNodePoolBatchAttachOutput) IsImport() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) pulumi.BoolOutput { return v.IsImport }).(pulumi.BoolOutput)
}

// The KubernetesConfig of NodeConfig.
func (o DefaultNodePoolBatchAttachOutput) KubernetesConfigs() DefaultNodePoolBatchAttachKubernetesConfigArrayOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) DefaultNodePoolBatchAttachKubernetesConfigArrayOutput {
		return v.KubernetesConfigs
	}).(DefaultNodePoolBatchAttachKubernetesConfigArrayOutput)
}

// The Config of NodePool.
func (o DefaultNodePoolBatchAttachOutput) NodeConfigs() DefaultNodePoolBatchAttachNodeConfigArrayOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) DefaultNodePoolBatchAttachNodeConfigArrayOutput {
		return v.NodeConfigs
	}).(DefaultNodePoolBatchAttachNodeConfigArrayOutput)
}

// Tags.
func (o DefaultNodePoolBatchAttachOutput) Tags() DefaultNodePoolBatchAttachTagArrayOutput {
	return o.ApplyT(func(v *DefaultNodePoolBatchAttach) DefaultNodePoolBatchAttachTagArrayOutput { return v.Tags }).(DefaultNodePoolBatchAttachTagArrayOutput)
}

type DefaultNodePoolBatchAttachArrayOutput struct{ *pulumi.OutputState }

func (DefaultNodePoolBatchAttachArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNodePoolBatchAttach)(nil)).Elem()
}

func (o DefaultNodePoolBatchAttachArrayOutput) ToDefaultNodePoolBatchAttachArrayOutput() DefaultNodePoolBatchAttachArrayOutput {
	return o
}

func (o DefaultNodePoolBatchAttachArrayOutput) ToDefaultNodePoolBatchAttachArrayOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachArrayOutput {
	return o
}

func (o DefaultNodePoolBatchAttachArrayOutput) Index(i pulumi.IntInput) DefaultNodePoolBatchAttachOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultNodePoolBatchAttach {
		return vs[0].([]*DefaultNodePoolBatchAttach)[vs[1].(int)]
	}).(DefaultNodePoolBatchAttachOutput)
}

type DefaultNodePoolBatchAttachMapOutput struct{ *pulumi.OutputState }

func (DefaultNodePoolBatchAttachMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNodePoolBatchAttach)(nil)).Elem()
}

func (o DefaultNodePoolBatchAttachMapOutput) ToDefaultNodePoolBatchAttachMapOutput() DefaultNodePoolBatchAttachMapOutput {
	return o
}

func (o DefaultNodePoolBatchAttachMapOutput) ToDefaultNodePoolBatchAttachMapOutputWithContext(ctx context.Context) DefaultNodePoolBatchAttachMapOutput {
	return o
}

func (o DefaultNodePoolBatchAttachMapOutput) MapIndex(k pulumi.StringInput) DefaultNodePoolBatchAttachOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultNodePoolBatchAttach {
		return vs[0].(map[string]*DefaultNodePoolBatchAttach)[vs[1].(string)]
	}).(DefaultNodePoolBatchAttachOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNodePoolBatchAttachInput)(nil)).Elem(), &DefaultNodePoolBatchAttach{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNodePoolBatchAttachArrayInput)(nil)).Elem(), DefaultNodePoolBatchAttachArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNodePoolBatchAttachMapInput)(nil)).Elem(), DefaultNodePoolBatchAttachMap{})
	pulumi.RegisterOutputType(DefaultNodePoolBatchAttachOutput{})
	pulumi.RegisterOutputType(DefaultNodePoolBatchAttachArrayOutput{})
	pulumi.RegisterOutputType(DefaultNodePoolBatchAttachMapOutput{})
}
