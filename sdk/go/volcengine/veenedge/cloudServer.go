// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package veenedge

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage veenedge cloud server
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/veenedge"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := veenedge.NewCloudServer(ctx, "foo", &veenedge.CloudServerArgs{
//				BillingConfig: &veenedge.CloudServerBillingConfigArgs{
//					BandwidthBillingMethod: pulumi.String("MonthlyP95"),
//					ComputingBillingMethod: pulumi.String("MonthlyPeak"),
//				},
//				CloudserverName: pulumi.String("tf-test"),
//				DefaultAreaName: pulumi.String("C******na"),
//				DefaultIsp:      pulumi.String("CMCC"),
//				ImageId:         pulumi.String("image*****viqm"),
//				NetworkConfig: &veenedge.CloudServerNetworkConfigArgs{
//					BandwidthPeak: pulumi.String("5"),
//				},
//				ScheduleStrategy: &veenedge.CloudServerScheduleStrategyArgs{
//					NetworkStrategy:  pulumi.String("region"),
//					PriceStrategy:    pulumi.String("high_priority"),
//					ScheduleStrategy: pulumi.String("dispersion"),
//				},
//				SecretData:      pulumi.String("sshkey-47*****wgc"),
//				SecretType:      pulumi.String("KeyPair"),
//				ServerAreaLevel: pulumi.String("region"),
//				SpecName:        pulumi.String("veEN****rge"),
//				StorageConfig: &veenedge.CloudServerStorageConfigArgs{
//					DataDiskLists: veenedge.CloudServerStorageConfigDataDiskListArray{
//						&veenedge.CloudServerStorageConfigDataDiskListArgs{
//							Capacity:    pulumi.String("20"),
//							StorageType: pulumi.String("CloudBlockSSD"),
//						},
//					},
//					SystemDisk: &veenedge.CloudServerStorageConfigSystemDiskArgs{
//						Capacity:    pulumi.String("40"),
//						StorageType: pulumi.String("CloudBlockSSD"),
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
// CloudServer can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:veenedge/cloudServer:CloudServer default cloudserver-n769ewmjjqyqh5dv
//
// ```
//
//	After the veenedge cloud server is created, a default edge instance will be created, we recommend managing this default instance as follows resource "volcengine_veenedge_instance" "foo1" {
//
//	instance_id = volcengine_veenedge_cloud_server.foo.default_instance_id }
type CloudServer struct {
	pulumi.CustomResourceState

	// The config of the billing.
	BillingConfig CloudServerBillingConfigPtrOutput `pulumi:"billingConfig"`
	// The name of cloud server.
	CloudserverName pulumi.StringOutput `pulumi:"cloudserverName"`
	// The custom data.
	CustomData CloudServerCustomDataOutput `pulumi:"customData"`
	// The name of default area.
	DefaultAreaName pulumi.StringOutput `pulumi:"defaultAreaName"`
	// The name of default cluster.
	DefaultClusterName pulumi.StringPtrOutput `pulumi:"defaultClusterName"`
	// The default instance id generate by cloud server.
	DefaultInstanceId pulumi.StringOutput `pulumi:"defaultInstanceId"`
	// The default isp info.
	DefaultIsp pulumi.StringOutput `pulumi:"defaultIsp"`
	// The image id of cloud server.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The config of the network.
	NetworkConfig CloudServerNetworkConfigOutput `pulumi:"networkConfig"`
	// The schedule strategy.
	ScheduleStrategy CloudServerScheduleStrategyOutput `pulumi:"scheduleStrategy"`
	// The data of secret. The value can be Password or KeyPair ID.
	SecretData pulumi.StringPtrOutput `pulumi:"secretData"`
	// The type of secret. The value can be `KeyPair` or `Password`.
	SecretType pulumi.StringOutput `pulumi:"secretType"`
	// The server area level. The value can be `region` or `city`.
	ServerAreaLevel pulumi.StringOutput `pulumi:"serverAreaLevel"`
	// The spec name of cloud server.
	SpecName pulumi.StringOutput `pulumi:"specName"`
	// The config of the storage.
	StorageConfig CloudServerStorageConfigOutput `pulumi:"storageConfig"`
}

// NewCloudServer registers a new resource with the given unique name, arguments, and options.
func NewCloudServer(ctx *pulumi.Context,
	name string, args *CloudServerArgs, opts ...pulumi.ResourceOption) (*CloudServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CloudserverName == nil {
		return nil, errors.New("invalid value for required argument 'CloudserverName'")
	}
	if args.DefaultAreaName == nil {
		return nil, errors.New("invalid value for required argument 'DefaultAreaName'")
	}
	if args.DefaultIsp == nil {
		return nil, errors.New("invalid value for required argument 'DefaultIsp'")
	}
	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.NetworkConfig == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConfig'")
	}
	if args.ScheduleStrategy == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleStrategy'")
	}
	if args.SecretType == nil {
		return nil, errors.New("invalid value for required argument 'SecretType'")
	}
	if args.ServerAreaLevel == nil {
		return nil, errors.New("invalid value for required argument 'ServerAreaLevel'")
	}
	if args.SpecName == nil {
		return nil, errors.New("invalid value for required argument 'SpecName'")
	}
	if args.StorageConfig == nil {
		return nil, errors.New("invalid value for required argument 'StorageConfig'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource CloudServer
	err := ctx.RegisterResource("volcengine:veenedge/cloudServer:CloudServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCloudServer gets an existing CloudServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCloudServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CloudServerState, opts ...pulumi.ResourceOption) (*CloudServer, error) {
	var resource CloudServer
	err := ctx.ReadResource("volcengine:veenedge/cloudServer:CloudServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CloudServer resources.
type cloudServerState struct {
	// The config of the billing.
	BillingConfig *CloudServerBillingConfig `pulumi:"billingConfig"`
	// The name of cloud server.
	CloudserverName *string `pulumi:"cloudserverName"`
	// The custom data.
	CustomData *CloudServerCustomData `pulumi:"customData"`
	// The name of default area.
	DefaultAreaName *string `pulumi:"defaultAreaName"`
	// The name of default cluster.
	DefaultClusterName *string `pulumi:"defaultClusterName"`
	// The default instance id generate by cloud server.
	DefaultInstanceId *string `pulumi:"defaultInstanceId"`
	// The default isp info.
	DefaultIsp *string `pulumi:"defaultIsp"`
	// The image id of cloud server.
	ImageId *string `pulumi:"imageId"`
	// The config of the network.
	NetworkConfig *CloudServerNetworkConfig `pulumi:"networkConfig"`
	// The schedule strategy.
	ScheduleStrategy *CloudServerScheduleStrategy `pulumi:"scheduleStrategy"`
	// The data of secret. The value can be Password or KeyPair ID.
	SecretData *string `pulumi:"secretData"`
	// The type of secret. The value can be `KeyPair` or `Password`.
	SecretType *string `pulumi:"secretType"`
	// The server area level. The value can be `region` or `city`.
	ServerAreaLevel *string `pulumi:"serverAreaLevel"`
	// The spec name of cloud server.
	SpecName *string `pulumi:"specName"`
	// The config of the storage.
	StorageConfig *CloudServerStorageConfig `pulumi:"storageConfig"`
}

type CloudServerState struct {
	// The config of the billing.
	BillingConfig CloudServerBillingConfigPtrInput
	// The name of cloud server.
	CloudserverName pulumi.StringPtrInput
	// The custom data.
	CustomData CloudServerCustomDataPtrInput
	// The name of default area.
	DefaultAreaName pulumi.StringPtrInput
	// The name of default cluster.
	DefaultClusterName pulumi.StringPtrInput
	// The default instance id generate by cloud server.
	DefaultInstanceId pulumi.StringPtrInput
	// The default isp info.
	DefaultIsp pulumi.StringPtrInput
	// The image id of cloud server.
	ImageId pulumi.StringPtrInput
	// The config of the network.
	NetworkConfig CloudServerNetworkConfigPtrInput
	// The schedule strategy.
	ScheduleStrategy CloudServerScheduleStrategyPtrInput
	// The data of secret. The value can be Password or KeyPair ID.
	SecretData pulumi.StringPtrInput
	// The type of secret. The value can be `KeyPair` or `Password`.
	SecretType pulumi.StringPtrInput
	// The server area level. The value can be `region` or `city`.
	ServerAreaLevel pulumi.StringPtrInput
	// The spec name of cloud server.
	SpecName pulumi.StringPtrInput
	// The config of the storage.
	StorageConfig CloudServerStorageConfigPtrInput
}

func (CloudServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudServerState)(nil)).Elem()
}

type cloudServerArgs struct {
	// The config of the billing.
	BillingConfig *CloudServerBillingConfig `pulumi:"billingConfig"`
	// The name of cloud server.
	CloudserverName string `pulumi:"cloudserverName"`
	// The custom data.
	CustomData *CloudServerCustomData `pulumi:"customData"`
	// The name of default area.
	DefaultAreaName string `pulumi:"defaultAreaName"`
	// The name of default cluster.
	DefaultClusterName *string `pulumi:"defaultClusterName"`
	// The default isp info.
	DefaultIsp string `pulumi:"defaultIsp"`
	// The image id of cloud server.
	ImageId string `pulumi:"imageId"`
	// The config of the network.
	NetworkConfig CloudServerNetworkConfig `pulumi:"networkConfig"`
	// The schedule strategy.
	ScheduleStrategy CloudServerScheduleStrategy `pulumi:"scheduleStrategy"`
	// The data of secret. The value can be Password or KeyPair ID.
	SecretData *string `pulumi:"secretData"`
	// The type of secret. The value can be `KeyPair` or `Password`.
	SecretType string `pulumi:"secretType"`
	// The server area level. The value can be `region` or `city`.
	ServerAreaLevel string `pulumi:"serverAreaLevel"`
	// The spec name of cloud server.
	SpecName string `pulumi:"specName"`
	// The config of the storage.
	StorageConfig CloudServerStorageConfig `pulumi:"storageConfig"`
}

// The set of arguments for constructing a CloudServer resource.
type CloudServerArgs struct {
	// The config of the billing.
	BillingConfig CloudServerBillingConfigPtrInput
	// The name of cloud server.
	CloudserverName pulumi.StringInput
	// The custom data.
	CustomData CloudServerCustomDataPtrInput
	// The name of default area.
	DefaultAreaName pulumi.StringInput
	// The name of default cluster.
	DefaultClusterName pulumi.StringPtrInput
	// The default isp info.
	DefaultIsp pulumi.StringInput
	// The image id of cloud server.
	ImageId pulumi.StringInput
	// The config of the network.
	NetworkConfig CloudServerNetworkConfigInput
	// The schedule strategy.
	ScheduleStrategy CloudServerScheduleStrategyInput
	// The data of secret. The value can be Password or KeyPair ID.
	SecretData pulumi.StringPtrInput
	// The type of secret. The value can be `KeyPair` or `Password`.
	SecretType pulumi.StringInput
	// The server area level. The value can be `region` or `city`.
	ServerAreaLevel pulumi.StringInput
	// The spec name of cloud server.
	SpecName pulumi.StringInput
	// The config of the storage.
	StorageConfig CloudServerStorageConfigInput
}

func (CloudServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cloudServerArgs)(nil)).Elem()
}

type CloudServerInput interface {
	pulumi.Input

	ToCloudServerOutput() CloudServerOutput
	ToCloudServerOutputWithContext(ctx context.Context) CloudServerOutput
}

func (*CloudServer) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServer)(nil)).Elem()
}

func (i *CloudServer) ToCloudServerOutput() CloudServerOutput {
	return i.ToCloudServerOutputWithContext(context.Background())
}

func (i *CloudServer) ToCloudServerOutputWithContext(ctx context.Context) CloudServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServerOutput)
}

// CloudServerArrayInput is an input type that accepts CloudServerArray and CloudServerArrayOutput values.
// You can construct a concrete instance of `CloudServerArrayInput` via:
//
//	CloudServerArray{ CloudServerArgs{...} }
type CloudServerArrayInput interface {
	pulumi.Input

	ToCloudServerArrayOutput() CloudServerArrayOutput
	ToCloudServerArrayOutputWithContext(context.Context) CloudServerArrayOutput
}

type CloudServerArray []CloudServerInput

func (CloudServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudServer)(nil)).Elem()
}

func (i CloudServerArray) ToCloudServerArrayOutput() CloudServerArrayOutput {
	return i.ToCloudServerArrayOutputWithContext(context.Background())
}

func (i CloudServerArray) ToCloudServerArrayOutputWithContext(ctx context.Context) CloudServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServerArrayOutput)
}

// CloudServerMapInput is an input type that accepts CloudServerMap and CloudServerMapOutput values.
// You can construct a concrete instance of `CloudServerMapInput` via:
//
//	CloudServerMap{ "key": CloudServerArgs{...} }
type CloudServerMapInput interface {
	pulumi.Input

	ToCloudServerMapOutput() CloudServerMapOutput
	ToCloudServerMapOutputWithContext(context.Context) CloudServerMapOutput
}

type CloudServerMap map[string]CloudServerInput

func (CloudServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudServer)(nil)).Elem()
}

func (i CloudServerMap) ToCloudServerMapOutput() CloudServerMapOutput {
	return i.ToCloudServerMapOutputWithContext(context.Background())
}

func (i CloudServerMap) ToCloudServerMapOutputWithContext(ctx context.Context) CloudServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloudServerMapOutput)
}

type CloudServerOutput struct{ *pulumi.OutputState }

func (CloudServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloudServer)(nil)).Elem()
}

func (o CloudServerOutput) ToCloudServerOutput() CloudServerOutput {
	return o
}

func (o CloudServerOutput) ToCloudServerOutputWithContext(ctx context.Context) CloudServerOutput {
	return o
}

// The config of the billing.
func (o CloudServerOutput) BillingConfig() CloudServerBillingConfigPtrOutput {
	return o.ApplyT(func(v *CloudServer) CloudServerBillingConfigPtrOutput { return v.BillingConfig }).(CloudServerBillingConfigPtrOutput)
}

// The name of cloud server.
func (o CloudServerOutput) CloudserverName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.CloudserverName }).(pulumi.StringOutput)
}

// The custom data.
func (o CloudServerOutput) CustomData() CloudServerCustomDataOutput {
	return o.ApplyT(func(v *CloudServer) CloudServerCustomDataOutput { return v.CustomData }).(CloudServerCustomDataOutput)
}

// The name of default area.
func (o CloudServerOutput) DefaultAreaName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.DefaultAreaName }).(pulumi.StringOutput)
}

// The name of default cluster.
func (o CloudServerOutput) DefaultClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringPtrOutput { return v.DefaultClusterName }).(pulumi.StringPtrOutput)
}

// The default instance id generate by cloud server.
func (o CloudServerOutput) DefaultInstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.DefaultInstanceId }).(pulumi.StringOutput)
}

// The default isp info.
func (o CloudServerOutput) DefaultIsp() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.DefaultIsp }).(pulumi.StringOutput)
}

// The image id of cloud server.
func (o CloudServerOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The config of the network.
func (o CloudServerOutput) NetworkConfig() CloudServerNetworkConfigOutput {
	return o.ApplyT(func(v *CloudServer) CloudServerNetworkConfigOutput { return v.NetworkConfig }).(CloudServerNetworkConfigOutput)
}

// The schedule strategy.
func (o CloudServerOutput) ScheduleStrategy() CloudServerScheduleStrategyOutput {
	return o.ApplyT(func(v *CloudServer) CloudServerScheduleStrategyOutput { return v.ScheduleStrategy }).(CloudServerScheduleStrategyOutput)
}

// The data of secret. The value can be Password or KeyPair ID.
func (o CloudServerOutput) SecretData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringPtrOutput { return v.SecretData }).(pulumi.StringPtrOutput)
}

// The type of secret. The value can be `KeyPair` or `Password`.
func (o CloudServerOutput) SecretType() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.SecretType }).(pulumi.StringOutput)
}

// The server area level. The value can be `region` or `city`.
func (o CloudServerOutput) ServerAreaLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.ServerAreaLevel }).(pulumi.StringOutput)
}

// The spec name of cloud server.
func (o CloudServerOutput) SpecName() pulumi.StringOutput {
	return o.ApplyT(func(v *CloudServer) pulumi.StringOutput { return v.SpecName }).(pulumi.StringOutput)
}

// The config of the storage.
func (o CloudServerOutput) StorageConfig() CloudServerStorageConfigOutput {
	return o.ApplyT(func(v *CloudServer) CloudServerStorageConfigOutput { return v.StorageConfig }).(CloudServerStorageConfigOutput)
}

type CloudServerArrayOutput struct{ *pulumi.OutputState }

func (CloudServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CloudServer)(nil)).Elem()
}

func (o CloudServerArrayOutput) ToCloudServerArrayOutput() CloudServerArrayOutput {
	return o
}

func (o CloudServerArrayOutput) ToCloudServerArrayOutputWithContext(ctx context.Context) CloudServerArrayOutput {
	return o
}

func (o CloudServerArrayOutput) Index(i pulumi.IntInput) CloudServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CloudServer {
		return vs[0].([]*CloudServer)[vs[1].(int)]
	}).(CloudServerOutput)
}

type CloudServerMapOutput struct{ *pulumi.OutputState }

func (CloudServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CloudServer)(nil)).Elem()
}

func (o CloudServerMapOutput) ToCloudServerMapOutput() CloudServerMapOutput {
	return o
}

func (o CloudServerMapOutput) ToCloudServerMapOutputWithContext(ctx context.Context) CloudServerMapOutput {
	return o
}

func (o CloudServerMapOutput) MapIndex(k pulumi.StringInput) CloudServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CloudServer {
		return vs[0].(map[string]*CloudServer)[vs[1].(string)]
	}).(CloudServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CloudServerInput)(nil)).Elem(), &CloudServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudServerArrayInput)(nil)).Elem(), CloudServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CloudServerMapInput)(nil)).Elem(), CloudServerMap{})
	pulumi.RegisterOutputType(CloudServerOutput{})
	pulumi.RegisterOutputType(CloudServerArrayOutput{})
	pulumi.RegisterOutputType(CloudServerMapOutput{})
}
