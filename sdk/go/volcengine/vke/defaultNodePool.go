// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage vke default node pool
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
//			_, err := vke.NewDefaultNodePool(ctx, "default", &vke.DefaultNodePoolArgs{
//				ClusterId: pulumi.String("ccc2umdnqtoflv91lqtq0"),
//				Instances: vke.DefaultNodePoolInstanceArray{
//					&vke.DefaultNodePoolInstanceArgs{
//						AdditionalContainerStorageEnabled: pulumi.Bool(false),
//						InstanceId:                        pulumi.String("i-ybvza90ohwexzk8emaa3"),
//						KeepInstanceName:                  pulumi.Bool(false),
//					},
//					&vke.DefaultNodePoolInstanceArgs{
//						AdditionalContainerStorageEnabled: pulumi.Bool(true),
//						ContainerStoragePath:              pulumi.String("/"),
//						InstanceId:                        pulumi.String("i-ybvza90ohxexzkm4zihf"),
//						KeepInstanceName:                  pulumi.Bool(false),
//					},
//				},
//				KubernetesConfig: &vke.DefaultNodePoolKubernetesConfigArgs{
//					Cordon: pulumi.Bool(true),
//					Labels: vke.DefaultNodePoolKubernetesConfigLabelArray{
//						&vke.DefaultNodePoolKubernetesConfigLabelArgs{
//							Key:   pulumi.String("aa"),
//							Value: pulumi.String("bb"),
//						},
//						&vke.DefaultNodePoolKubernetesConfigLabelArgs{
//							Key:   pulumi.String("cccc"),
//							Value: pulumi.String("dddd"),
//						},
//					},
//					Taints: vke.DefaultNodePoolKubernetesConfigTaintArray{
//						&vke.DefaultNodePoolKubernetesConfigTaintArgs{
//							Effect: pulumi.String("NoSchedule"),
//							Key:    pulumi.String("cccc"),
//							Value:  pulumi.String("dddd"),
//						},
//						&vke.DefaultNodePoolKubernetesConfigTaintArgs{
//							Effect: pulumi.String("NoSchedule"),
//							Key:    pulumi.String("aa11"),
//							Value:  pulumi.String("111"),
//						},
//					},
//				},
//				NodeConfig: &vke.DefaultNodePoolNodeConfigArgs{
//					EcsTags: vke.DefaultNodePoolNodeConfigEcsTagArray{
//						&vke.DefaultNodePoolNodeConfigEcsTagArgs{
//							Key:   pulumi.String("ecs_k1"),
//							Value: pulumi.String("ecs_v1"),
//						},
//					},
//					InitializeScript: pulumi.String("ISMvYmluL2Jhc2gKZWNobyAx"),
//					Security: &vke.DefaultNodePoolNodeConfigSecurityArgs{
//						Login: &vke.DefaultNodePoolNodeConfigSecurityLoginArgs{
//							Password: pulumi.String("amw4WTdVcTRJVVFsUXpVTw=="),
//						},
//						SecurityGroupIds: pulumi.StringArray{
//							pulumi.String("sg-2d6t6djr2wge858ozfczv41xq"),
//							pulumi.String("sg-3re6v4lz76yv45zsk2hjvvwcj"),
//						},
//						SecurityStrategies: pulumi.StringArray{
//							pulumi.String("Hids"),
//						},
//					},
//				},
//				Tags: vke.DefaultNodePoolTagArray{
//					&vke.DefaultNodePoolTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
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
//
// ## Import
//
// VKE default node can be imported using the node id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vke/defaultNodePool:DefaultNodePool default nc5t5epmrsf****
//
// ```
type DefaultNodePool struct {
	pulumi.CustomResourceState

	// The ClusterId of NodePool.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The ECS InstanceIds add to NodePool.
	Instances DefaultNodePoolInstanceArrayOutput `pulumi:"instances"`
	// Is import of the DefaultNodePool. It only works when imported, set to true.
	IsImport pulumi.BoolOutput `pulumi:"isImport"`
	// The KubernetesConfig of NodeConfig.
	KubernetesConfig DefaultNodePoolKubernetesConfigOutput `pulumi:"kubernetesConfig"`
	// The Config of NodePool.
	NodeConfig DefaultNodePoolNodeConfigOutput `pulumi:"nodeConfig"`
	// Tags.
	Tags DefaultNodePoolTagArrayOutput `pulumi:"tags"`
}

// NewDefaultNodePool registers a new resource with the given unique name, arguments, and options.
func NewDefaultNodePool(ctx *pulumi.Context,
	name string, args *DefaultNodePoolArgs, opts ...pulumi.ResourceOption) (*DefaultNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.KubernetesConfig == nil {
		return nil, errors.New("invalid value for required argument 'KubernetesConfig'")
	}
	if args.NodeConfig == nil {
		return nil, errors.New("invalid value for required argument 'NodeConfig'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource DefaultNodePool
	err := ctx.RegisterResource("volcengine:vke/defaultNodePool:DefaultNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDefaultNodePool gets an existing DefaultNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDefaultNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DefaultNodePoolState, opts ...pulumi.ResourceOption) (*DefaultNodePool, error) {
	var resource DefaultNodePool
	err := ctx.ReadResource("volcengine:vke/defaultNodePool:DefaultNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DefaultNodePool resources.
type defaultNodePoolState struct {
	// The ClusterId of NodePool.
	ClusterId *string `pulumi:"clusterId"`
	// The ECS InstanceIds add to NodePool.
	Instances []DefaultNodePoolInstance `pulumi:"instances"`
	// Is import of the DefaultNodePool. It only works when imported, set to true.
	IsImport *bool `pulumi:"isImport"`
	// The KubernetesConfig of NodeConfig.
	KubernetesConfig *DefaultNodePoolKubernetesConfig `pulumi:"kubernetesConfig"`
	// The Config of NodePool.
	NodeConfig *DefaultNodePoolNodeConfig `pulumi:"nodeConfig"`
	// Tags.
	Tags []DefaultNodePoolTag `pulumi:"tags"`
}

type DefaultNodePoolState struct {
	// The ClusterId of NodePool.
	ClusterId pulumi.StringPtrInput
	// The ECS InstanceIds add to NodePool.
	Instances DefaultNodePoolInstanceArrayInput
	// Is import of the DefaultNodePool. It only works when imported, set to true.
	IsImport pulumi.BoolPtrInput
	// The KubernetesConfig of NodeConfig.
	KubernetesConfig DefaultNodePoolKubernetesConfigPtrInput
	// The Config of NodePool.
	NodeConfig DefaultNodePoolNodeConfigPtrInput
	// Tags.
	Tags DefaultNodePoolTagArrayInput
}

func (DefaultNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNodePoolState)(nil)).Elem()
}

type defaultNodePoolArgs struct {
	// The ClusterId of NodePool.
	ClusterId string `pulumi:"clusterId"`
	// The ECS InstanceIds add to NodePool.
	Instances []DefaultNodePoolInstance `pulumi:"instances"`
	// The KubernetesConfig of NodeConfig.
	KubernetesConfig DefaultNodePoolKubernetesConfig `pulumi:"kubernetesConfig"`
	// The Config of NodePool.
	NodeConfig DefaultNodePoolNodeConfig `pulumi:"nodeConfig"`
	// Tags.
	Tags []DefaultNodePoolTag `pulumi:"tags"`
}

// The set of arguments for constructing a DefaultNodePool resource.
type DefaultNodePoolArgs struct {
	// The ClusterId of NodePool.
	ClusterId pulumi.StringInput
	// The ECS InstanceIds add to NodePool.
	Instances DefaultNodePoolInstanceArrayInput
	// The KubernetesConfig of NodeConfig.
	KubernetesConfig DefaultNodePoolKubernetesConfigInput
	// The Config of NodePool.
	NodeConfig DefaultNodePoolNodeConfigInput
	// Tags.
	Tags DefaultNodePoolTagArrayInput
}

func (DefaultNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*defaultNodePoolArgs)(nil)).Elem()
}

type DefaultNodePoolInput interface {
	pulumi.Input

	ToDefaultNodePoolOutput() DefaultNodePoolOutput
	ToDefaultNodePoolOutputWithContext(ctx context.Context) DefaultNodePoolOutput
}

func (*DefaultNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNodePool)(nil)).Elem()
}

func (i *DefaultNodePool) ToDefaultNodePoolOutput() DefaultNodePoolOutput {
	return i.ToDefaultNodePoolOutputWithContext(context.Background())
}

func (i *DefaultNodePool) ToDefaultNodePoolOutputWithContext(ctx context.Context) DefaultNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNodePoolOutput)
}

// DefaultNodePoolArrayInput is an input type that accepts DefaultNodePoolArray and DefaultNodePoolArrayOutput values.
// You can construct a concrete instance of `DefaultNodePoolArrayInput` via:
//
//	DefaultNodePoolArray{ DefaultNodePoolArgs{...} }
type DefaultNodePoolArrayInput interface {
	pulumi.Input

	ToDefaultNodePoolArrayOutput() DefaultNodePoolArrayOutput
	ToDefaultNodePoolArrayOutputWithContext(context.Context) DefaultNodePoolArrayOutput
}

type DefaultNodePoolArray []DefaultNodePoolInput

func (DefaultNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNodePool)(nil)).Elem()
}

func (i DefaultNodePoolArray) ToDefaultNodePoolArrayOutput() DefaultNodePoolArrayOutput {
	return i.ToDefaultNodePoolArrayOutputWithContext(context.Background())
}

func (i DefaultNodePoolArray) ToDefaultNodePoolArrayOutputWithContext(ctx context.Context) DefaultNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNodePoolArrayOutput)
}

// DefaultNodePoolMapInput is an input type that accepts DefaultNodePoolMap and DefaultNodePoolMapOutput values.
// You can construct a concrete instance of `DefaultNodePoolMapInput` via:
//
//	DefaultNodePoolMap{ "key": DefaultNodePoolArgs{...} }
type DefaultNodePoolMapInput interface {
	pulumi.Input

	ToDefaultNodePoolMapOutput() DefaultNodePoolMapOutput
	ToDefaultNodePoolMapOutputWithContext(context.Context) DefaultNodePoolMapOutput
}

type DefaultNodePoolMap map[string]DefaultNodePoolInput

func (DefaultNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNodePool)(nil)).Elem()
}

func (i DefaultNodePoolMap) ToDefaultNodePoolMapOutput() DefaultNodePoolMapOutput {
	return i.ToDefaultNodePoolMapOutputWithContext(context.Background())
}

func (i DefaultNodePoolMap) ToDefaultNodePoolMapOutputWithContext(ctx context.Context) DefaultNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DefaultNodePoolMapOutput)
}

type DefaultNodePoolOutput struct{ *pulumi.OutputState }

func (DefaultNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultNodePool)(nil)).Elem()
}

func (o DefaultNodePoolOutput) ToDefaultNodePoolOutput() DefaultNodePoolOutput {
	return o
}

func (o DefaultNodePoolOutput) ToDefaultNodePoolOutputWithContext(ctx context.Context) DefaultNodePoolOutput {
	return o
}

// The ClusterId of NodePool.
func (o DefaultNodePoolOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DefaultNodePool) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The ECS InstanceIds add to NodePool.
func (o DefaultNodePoolOutput) Instances() DefaultNodePoolInstanceArrayOutput {
	return o.ApplyT(func(v *DefaultNodePool) DefaultNodePoolInstanceArrayOutput { return v.Instances }).(DefaultNodePoolInstanceArrayOutput)
}

// Is import of the DefaultNodePool. It only works when imported, set to true.
func (o DefaultNodePoolOutput) IsImport() pulumi.BoolOutput {
	return o.ApplyT(func(v *DefaultNodePool) pulumi.BoolOutput { return v.IsImport }).(pulumi.BoolOutput)
}

// The KubernetesConfig of NodeConfig.
func (o DefaultNodePoolOutput) KubernetesConfig() DefaultNodePoolKubernetesConfigOutput {
	return o.ApplyT(func(v *DefaultNodePool) DefaultNodePoolKubernetesConfigOutput { return v.KubernetesConfig }).(DefaultNodePoolKubernetesConfigOutput)
}

// The Config of NodePool.
func (o DefaultNodePoolOutput) NodeConfig() DefaultNodePoolNodeConfigOutput {
	return o.ApplyT(func(v *DefaultNodePool) DefaultNodePoolNodeConfigOutput { return v.NodeConfig }).(DefaultNodePoolNodeConfigOutput)
}

// Tags.
func (o DefaultNodePoolOutput) Tags() DefaultNodePoolTagArrayOutput {
	return o.ApplyT(func(v *DefaultNodePool) DefaultNodePoolTagArrayOutput { return v.Tags }).(DefaultNodePoolTagArrayOutput)
}

type DefaultNodePoolArrayOutput struct{ *pulumi.OutputState }

func (DefaultNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DefaultNodePool)(nil)).Elem()
}

func (o DefaultNodePoolArrayOutput) ToDefaultNodePoolArrayOutput() DefaultNodePoolArrayOutput {
	return o
}

func (o DefaultNodePoolArrayOutput) ToDefaultNodePoolArrayOutputWithContext(ctx context.Context) DefaultNodePoolArrayOutput {
	return o
}

func (o DefaultNodePoolArrayOutput) Index(i pulumi.IntInput) DefaultNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DefaultNodePool {
		return vs[0].([]*DefaultNodePool)[vs[1].(int)]
	}).(DefaultNodePoolOutput)
}

type DefaultNodePoolMapOutput struct{ *pulumi.OutputState }

func (DefaultNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DefaultNodePool)(nil)).Elem()
}

func (o DefaultNodePoolMapOutput) ToDefaultNodePoolMapOutput() DefaultNodePoolMapOutput {
	return o
}

func (o DefaultNodePoolMapOutput) ToDefaultNodePoolMapOutputWithContext(ctx context.Context) DefaultNodePoolMapOutput {
	return o
}

func (o DefaultNodePoolMapOutput) MapIndex(k pulumi.StringInput) DefaultNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DefaultNodePool {
		return vs[0].(map[string]*DefaultNodePool)[vs[1].(string)]
	}).(DefaultNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNodePoolInput)(nil)).Elem(), &DefaultNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNodePoolArrayInput)(nil)).Elem(), DefaultNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultNodePoolMapInput)(nil)).Elem(), DefaultNodePoolMap{})
	pulumi.RegisterOutputType(DefaultNodePoolOutput{})
	pulumi.RegisterOutputType(DefaultNodePoolArrayOutput{})
	pulumi.RegisterOutputType(DefaultNodePoolMapOutput{})
}
