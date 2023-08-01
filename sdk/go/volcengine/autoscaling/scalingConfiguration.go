// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := autoscaling.NewScalingConfiguration(ctx, "foo", &autoscaling.ScalingConfigurationArgs{
//				EipBandwidth:        pulumi.Int(10),
//				EipBillingType:      pulumi.String("PostPaidByBandwidth"),
//				EipIsp:              pulumi.String("ChinaMobile"),
//				HostName:            pulumi.String(""),
//				HpcClusterId:        pulumi.String(""),
//				ImageId:             pulumi.String("image-ycgud4t4hxgso0e27bdl"),
//				InstanceDescription: pulumi.String(""),
//				InstanceName:        pulumi.String("tf-test"),
//				InstanceTypes: pulumi.StringArray{
//					pulumi.String("ecs.g2i.large"),
//				},
//				KeyPairName:                 pulumi.String("tf-keypair"),
//				Password:                    pulumi.String(""),
//				ProjectName:                 pulumi.String("default"),
//				ScalingConfigurationName:    pulumi.String("tf-test"),
//				ScalingGroupId:              pulumi.String("scg-ycinx27x25gh9y31p0fy"),
//				SecurityEnhancementStrategy: pulumi.String("InActive"),
//				SecurityGroupIds: pulumi.StringArray{
//					pulumi.String("sg-2fepz3c793g1s59gp67y21r34"),
//				},
//				SpotStrategy: pulumi.String("NoSpot"),
//				Tags: autoscaling.ScalingConfigurationTagArray{
//					&autoscaling.ScalingConfigurationTagArgs{
//						Key:   pulumi.String("tf-key1"),
//						Value: pulumi.String("tf-value1"),
//					},
//					&autoscaling.ScalingConfigurationTagArgs{
//						Key:   pulumi.String("tf-key2"),
//						Value: pulumi.String("tf-value2"),
//					},
//				},
//				UserData: pulumi.String("IyEvYmluL2Jhc2gKZWNobyAidGVzdCI="),
//				Volumes: autoscaling.ScalingConfigurationVolumeArray{
//					&autoscaling.ScalingConfigurationVolumeArgs{
//						DeleteWithInstance: pulumi.Bool(false),
//						Size:               pulumi.Int(20),
//						VolumeType:         pulumi.String("ESSD_PL0"),
//					},
//					&autoscaling.ScalingConfigurationVolumeArgs{
//						DeleteWithInstance: pulumi.Bool(true),
//						Size:               pulumi.Int(20),
//						VolumeType:         pulumi.String("ESSD_PL0"),
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
// ScalingConfiguration can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:autoscaling/scalingConfiguration:ScalingConfiguration default scc-ybkuck3mx8cm9tm5yglz
//
// ```
type ScalingConfiguration struct {
	pulumi.CustomResourceState

	// The create time of the scaling configuration.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
	EipBandwidth pulumi.IntOutput `pulumi:"eipBandwidth"`
	// The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
	EipBillingType pulumi.StringOutput `pulumi:"eipBillingType"`
	// The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
	EipIsp pulumi.StringOutput `pulumi:"eipIsp"`
	// The ECS hostname which the scaling configuration set.
	HostName pulumi.StringPtrOutput `pulumi:"hostName"`
	// The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
	HpcClusterId pulumi.StringPtrOutput `pulumi:"hpcClusterId"`
	// The ECS image id which the scaling configuration set.
	ImageId pulumi.StringOutput `pulumi:"imageId"`
	// The ECS instance description which the scaling configuration set.
	InstanceDescription pulumi.StringPtrOutput `pulumi:"instanceDescription"`
	// The ECS instance name which the scaling configuration set.
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
	InstanceTypes pulumi.StringArrayOutput `pulumi:"instanceTypes"`
	// The ECS key pair name which the scaling configuration set.
	KeyPairName pulumi.StringPtrOutput `pulumi:"keyPairName"`
	// The lifecycle state of the scaling configuration.
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The ECS password which the scaling configuration set.
	Password pulumi.StringPtrOutput `pulumi:"password"`
	// The project to which the instance created by the scaling configuration belongs.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The id of the scaling configuration.
	ScalingConfigurationId pulumi.StringOutput `pulumi:"scalingConfigurationId"`
	// The name of the scaling configuration.
	ScalingConfigurationName pulumi.StringOutput `pulumi:"scalingConfigurationName"`
	// The id of the scaling group to which the scaling configuration belongs.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
	// The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
	SecurityEnhancementStrategy pulumi.StringPtrOutput `pulumi:"securityEnhancementStrategy"`
	// The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
	SecurityGroupIds pulumi.StringArrayOutput `pulumi:"securityGroupIds"`
	// The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
	SpotStrategy pulumi.StringPtrOutput `pulumi:"spotStrategy"`
	// The label of the instance created by the scaling configuration. Up to 20 tags are supported.
	Tags ScalingConfigurationTagArrayOutput `pulumi:"tags"`
	// The create time of the scaling configuration.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// The ECS user data which the scaling configuration set.
	UserData pulumi.StringPtrOutput `pulumi:"userData"`
	// The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
	Volumes ScalingConfigurationVolumeArrayOutput `pulumi:"volumes"`
}

// NewScalingConfiguration registers a new resource with the given unique name, arguments, and options.
func NewScalingConfiguration(ctx *pulumi.Context,
	name string, args *ScalingConfigurationArgs, opts ...pulumi.ResourceOption) (*ScalingConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ImageId == nil {
		return nil, errors.New("invalid value for required argument 'ImageId'")
	}
	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.InstanceTypes == nil {
		return nil, errors.New("invalid value for required argument 'InstanceTypes'")
	}
	if args.ScalingConfigurationName == nil {
		return nil, errors.New("invalid value for required argument 'ScalingConfigurationName'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	if args.SecurityGroupIds == nil {
		return nil, errors.New("invalid value for required argument 'SecurityGroupIds'")
	}
	if args.Volumes == nil {
		return nil, errors.New("invalid value for required argument 'Volumes'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ScalingConfiguration
	err := ctx.RegisterResource("volcengine:autoscaling/scalingConfiguration:ScalingConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingConfiguration gets an existing ScalingConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingConfigurationState, opts ...pulumi.ResourceOption) (*ScalingConfiguration, error) {
	var resource ScalingConfiguration
	err := ctx.ReadResource("volcengine:autoscaling/scalingConfiguration:ScalingConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingConfiguration resources.
type scalingConfigurationState struct {
	// The create time of the scaling configuration.
	CreatedAt *string `pulumi:"createdAt"`
	// The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
	EipBandwidth *int `pulumi:"eipBandwidth"`
	// The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
	EipBillingType *string `pulumi:"eipBillingType"`
	// The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
	EipIsp *string `pulumi:"eipIsp"`
	// The ECS hostname which the scaling configuration set.
	HostName *string `pulumi:"hostName"`
	// The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
	HpcClusterId *string `pulumi:"hpcClusterId"`
	// The ECS image id which the scaling configuration set.
	ImageId *string `pulumi:"imageId"`
	// The ECS instance description which the scaling configuration set.
	InstanceDescription *string `pulumi:"instanceDescription"`
	// The ECS instance name which the scaling configuration set.
	InstanceName *string `pulumi:"instanceName"`
	// The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// The ECS key pair name which the scaling configuration set.
	KeyPairName *string `pulumi:"keyPairName"`
	// The lifecycle state of the scaling configuration.
	LifecycleState *string `pulumi:"lifecycleState"`
	// The ECS password which the scaling configuration set.
	Password *string `pulumi:"password"`
	// The project to which the instance created by the scaling configuration belongs.
	ProjectName *string `pulumi:"projectName"`
	// The id of the scaling configuration.
	ScalingConfigurationId *string `pulumi:"scalingConfigurationId"`
	// The name of the scaling configuration.
	ScalingConfigurationName *string `pulumi:"scalingConfigurationName"`
	// The id of the scaling group to which the scaling configuration belongs.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
	// The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
	SecurityEnhancementStrategy *string `pulumi:"securityEnhancementStrategy"`
	// The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
	SpotStrategy *string `pulumi:"spotStrategy"`
	// The label of the instance created by the scaling configuration. Up to 20 tags are supported.
	Tags []ScalingConfigurationTag `pulumi:"tags"`
	// The create time of the scaling configuration.
	UpdatedAt *string `pulumi:"updatedAt"`
	// The ECS user data which the scaling configuration set.
	UserData *string `pulumi:"userData"`
	// The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
	Volumes []ScalingConfigurationVolume `pulumi:"volumes"`
}

type ScalingConfigurationState struct {
	// The create time of the scaling configuration.
	CreatedAt pulumi.StringPtrInput
	// The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
	EipBandwidth pulumi.IntPtrInput
	// The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
	EipBillingType pulumi.StringPtrInput
	// The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
	EipIsp pulumi.StringPtrInput
	// The ECS hostname which the scaling configuration set.
	HostName pulumi.StringPtrInput
	// The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
	HpcClusterId pulumi.StringPtrInput
	// The ECS image id which the scaling configuration set.
	ImageId pulumi.StringPtrInput
	// The ECS instance description which the scaling configuration set.
	InstanceDescription pulumi.StringPtrInput
	// The ECS instance name which the scaling configuration set.
	InstanceName pulumi.StringPtrInput
	// The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
	InstanceTypes pulumi.StringArrayInput
	// The ECS key pair name which the scaling configuration set.
	KeyPairName pulumi.StringPtrInput
	// The lifecycle state of the scaling configuration.
	LifecycleState pulumi.StringPtrInput
	// The ECS password which the scaling configuration set.
	Password pulumi.StringPtrInput
	// The project to which the instance created by the scaling configuration belongs.
	ProjectName pulumi.StringPtrInput
	// The id of the scaling configuration.
	ScalingConfigurationId pulumi.StringPtrInput
	// The name of the scaling configuration.
	ScalingConfigurationName pulumi.StringPtrInput
	// The id of the scaling group to which the scaling configuration belongs.
	ScalingGroupId pulumi.StringPtrInput
	// The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
	SecurityEnhancementStrategy pulumi.StringPtrInput
	// The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
	SecurityGroupIds pulumi.StringArrayInput
	// The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
	SpotStrategy pulumi.StringPtrInput
	// The label of the instance created by the scaling configuration. Up to 20 tags are supported.
	Tags ScalingConfigurationTagArrayInput
	// The create time of the scaling configuration.
	UpdatedAt pulumi.StringPtrInput
	// The ECS user data which the scaling configuration set.
	UserData pulumi.StringPtrInput
	// The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
	Volumes ScalingConfigurationVolumeArrayInput
}

func (ScalingConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigurationState)(nil)).Elem()
}

type scalingConfigurationArgs struct {
	// The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
	EipBandwidth *int `pulumi:"eipBandwidth"`
	// The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
	EipBillingType *string `pulumi:"eipBillingType"`
	// The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
	EipIsp *string `pulumi:"eipIsp"`
	// The ECS hostname which the scaling configuration set.
	HostName *string `pulumi:"hostName"`
	// The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
	HpcClusterId *string `pulumi:"hpcClusterId"`
	// The ECS image id which the scaling configuration set.
	ImageId string `pulumi:"imageId"`
	// The ECS instance description which the scaling configuration set.
	InstanceDescription *string `pulumi:"instanceDescription"`
	// The ECS instance name which the scaling configuration set.
	InstanceName string `pulumi:"instanceName"`
	// The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
	InstanceTypes []string `pulumi:"instanceTypes"`
	// The ECS key pair name which the scaling configuration set.
	KeyPairName *string `pulumi:"keyPairName"`
	// The ECS password which the scaling configuration set.
	Password *string `pulumi:"password"`
	// The project to which the instance created by the scaling configuration belongs.
	ProjectName *string `pulumi:"projectName"`
	// The name of the scaling configuration.
	ScalingConfigurationName string `pulumi:"scalingConfigurationName"`
	// The id of the scaling group to which the scaling configuration belongs.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
	SecurityEnhancementStrategy *string `pulumi:"securityEnhancementStrategy"`
	// The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
	SpotStrategy *string `pulumi:"spotStrategy"`
	// The label of the instance created by the scaling configuration. Up to 20 tags are supported.
	Tags []ScalingConfigurationTag `pulumi:"tags"`
	// The ECS user data which the scaling configuration set.
	UserData *string `pulumi:"userData"`
	// The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
	Volumes []ScalingConfigurationVolume `pulumi:"volumes"`
}

// The set of arguments for constructing a ScalingConfiguration resource.
type ScalingConfigurationArgs struct {
	// The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
	EipBandwidth pulumi.IntPtrInput
	// The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
	EipBillingType pulumi.StringPtrInput
	// The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
	EipIsp pulumi.StringPtrInput
	// The ECS hostname which the scaling configuration set.
	HostName pulumi.StringPtrInput
	// The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
	HpcClusterId pulumi.StringPtrInput
	// The ECS image id which the scaling configuration set.
	ImageId pulumi.StringInput
	// The ECS instance description which the scaling configuration set.
	InstanceDescription pulumi.StringPtrInput
	// The ECS instance name which the scaling configuration set.
	InstanceName pulumi.StringInput
	// The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
	InstanceTypes pulumi.StringArrayInput
	// The ECS key pair name which the scaling configuration set.
	KeyPairName pulumi.StringPtrInput
	// The ECS password which the scaling configuration set.
	Password pulumi.StringPtrInput
	// The project to which the instance created by the scaling configuration belongs.
	ProjectName pulumi.StringPtrInput
	// The name of the scaling configuration.
	ScalingConfigurationName pulumi.StringInput
	// The id of the scaling group to which the scaling configuration belongs.
	ScalingGroupId pulumi.StringInput
	// The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
	SecurityEnhancementStrategy pulumi.StringPtrInput
	// The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
	SecurityGroupIds pulumi.StringArrayInput
	// The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
	SpotStrategy pulumi.StringPtrInput
	// The label of the instance created by the scaling configuration. Up to 20 tags are supported.
	Tags ScalingConfigurationTagArrayInput
	// The ECS user data which the scaling configuration set.
	UserData pulumi.StringPtrInput
	// The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
	Volumes ScalingConfigurationVolumeArrayInput
}

func (ScalingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigurationArgs)(nil)).Elem()
}

type ScalingConfigurationInput interface {
	pulumi.Input

	ToScalingConfigurationOutput() ScalingConfigurationOutput
	ToScalingConfigurationOutputWithContext(ctx context.Context) ScalingConfigurationOutput
}

func (*ScalingConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingConfiguration)(nil)).Elem()
}

func (i *ScalingConfiguration) ToScalingConfigurationOutput() ScalingConfigurationOutput {
	return i.ToScalingConfigurationOutputWithContext(context.Background())
}

func (i *ScalingConfiguration) ToScalingConfigurationOutputWithContext(ctx context.Context) ScalingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationOutput)
}

// ScalingConfigurationArrayInput is an input type that accepts ScalingConfigurationArray and ScalingConfigurationArrayOutput values.
// You can construct a concrete instance of `ScalingConfigurationArrayInput` via:
//
//	ScalingConfigurationArray{ ScalingConfigurationArgs{...} }
type ScalingConfigurationArrayInput interface {
	pulumi.Input

	ToScalingConfigurationArrayOutput() ScalingConfigurationArrayOutput
	ToScalingConfigurationArrayOutputWithContext(context.Context) ScalingConfigurationArrayOutput
}

type ScalingConfigurationArray []ScalingConfigurationInput

func (ScalingConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingConfiguration)(nil)).Elem()
}

func (i ScalingConfigurationArray) ToScalingConfigurationArrayOutput() ScalingConfigurationArrayOutput {
	return i.ToScalingConfigurationArrayOutputWithContext(context.Background())
}

func (i ScalingConfigurationArray) ToScalingConfigurationArrayOutputWithContext(ctx context.Context) ScalingConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationArrayOutput)
}

// ScalingConfigurationMapInput is an input type that accepts ScalingConfigurationMap and ScalingConfigurationMapOutput values.
// You can construct a concrete instance of `ScalingConfigurationMapInput` via:
//
//	ScalingConfigurationMap{ "key": ScalingConfigurationArgs{...} }
type ScalingConfigurationMapInput interface {
	pulumi.Input

	ToScalingConfigurationMapOutput() ScalingConfigurationMapOutput
	ToScalingConfigurationMapOutputWithContext(context.Context) ScalingConfigurationMapOutput
}

type ScalingConfigurationMap map[string]ScalingConfigurationInput

func (ScalingConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingConfiguration)(nil)).Elem()
}

func (i ScalingConfigurationMap) ToScalingConfigurationMapOutput() ScalingConfigurationMapOutput {
	return i.ToScalingConfigurationMapOutputWithContext(context.Background())
}

func (i ScalingConfigurationMap) ToScalingConfigurationMapOutputWithContext(ctx context.Context) ScalingConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationMapOutput)
}

type ScalingConfigurationOutput struct{ *pulumi.OutputState }

func (ScalingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingConfiguration)(nil)).Elem()
}

func (o ScalingConfigurationOutput) ToScalingConfigurationOutput() ScalingConfigurationOutput {
	return o
}

func (o ScalingConfigurationOutput) ToScalingConfigurationOutputWithContext(ctx context.Context) ScalingConfigurationOutput {
	return o
}

// The create time of the scaling configuration.
func (o ScalingConfigurationOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// The EIP bandwidth which the scaling configuration set. When the value of Eip.BillingType is PostPaidByBandwidth, the value is 1 to 500. When the value of Eip.BillingType is PostPaidByTraffic, the value is 1 to 200.
func (o ScalingConfigurationOutput) EipBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.IntOutput { return v.EipBandwidth }).(pulumi.IntOutput)
}

// The EIP billing type which the scaling configuration set. Valid values: PostPaidByBandwidth, PostPaidByTraffic.
func (o ScalingConfigurationOutput) EipBillingType() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.EipBillingType }).(pulumi.StringOutput)
}

// The EIP ISP which the scaling configuration set. Valid values: BGP, ChinaMobile, ChinaUnicom, ChinaTelecom.
func (o ScalingConfigurationOutput) EipIsp() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.EipIsp }).(pulumi.StringOutput)
}

// The ECS hostname which the scaling configuration set.
func (o ScalingConfigurationOutput) HostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.HostName }).(pulumi.StringPtrOutput)
}

// The ID of the HPC cluster to which the instance belongs. Valid only when InstanceTypes.N specifies High Performance Computing GPU Type.
func (o ScalingConfigurationOutput) HpcClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.HpcClusterId }).(pulumi.StringPtrOutput)
}

// The ECS image id which the scaling configuration set.
func (o ScalingConfigurationOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.ImageId }).(pulumi.StringOutput)
}

// The ECS instance description which the scaling configuration set.
func (o ScalingConfigurationOutput) InstanceDescription() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.InstanceDescription }).(pulumi.StringPtrOutput)
}

// The ECS instance name which the scaling configuration set.
func (o ScalingConfigurationOutput) InstanceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.InstanceName }).(pulumi.StringOutput)
}

// The list of the ECS instance type which the scaling configuration set. The maximum number of instance types is 10.
func (o ScalingConfigurationOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringArrayOutput { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

// The ECS key pair name which the scaling configuration set.
func (o ScalingConfigurationOutput) KeyPairName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.KeyPairName }).(pulumi.StringPtrOutput)
}

// The lifecycle state of the scaling configuration.
func (o ScalingConfigurationOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// The ECS password which the scaling configuration set.
func (o ScalingConfigurationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.Password }).(pulumi.StringPtrOutput)
}

// The project to which the instance created by the scaling configuration belongs.
func (o ScalingConfigurationOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The id of the scaling configuration.
func (o ScalingConfigurationOutput) ScalingConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.ScalingConfigurationId }).(pulumi.StringOutput)
}

// The name of the scaling configuration.
func (o ScalingConfigurationOutput) ScalingConfigurationName() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.ScalingConfigurationName }).(pulumi.StringOutput)
}

// The id of the scaling group to which the scaling configuration belongs.
func (o ScalingConfigurationOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The Ecs security enhancement strategy which the scaling configuration set. Valid values: Active, InActive.
func (o ScalingConfigurationOutput) SecurityEnhancementStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.SecurityEnhancementStrategy }).(pulumi.StringPtrOutput)
}

// The list of the security group id of the networkInterface which the scaling configuration set. A maximum of 5 security groups can be bound at the same time, and the value ranges from 1 to 5.
func (o ScalingConfigurationOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringArrayOutput { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The preemption policy of the instance. Valid Value: NoSpot (default), SpotAsPriceGo.
func (o ScalingConfigurationOutput) SpotStrategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.SpotStrategy }).(pulumi.StringPtrOutput)
}

// The label of the instance created by the scaling configuration. Up to 20 tags are supported.
func (o ScalingConfigurationOutput) Tags() ScalingConfigurationTagArrayOutput {
	return o.ApplyT(func(v *ScalingConfiguration) ScalingConfigurationTagArrayOutput { return v.Tags }).(ScalingConfigurationTagArrayOutput)
}

// The create time of the scaling configuration.
func (o ScalingConfigurationOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// The ECS user data which the scaling configuration set.
func (o ScalingConfigurationOutput) UserData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingConfiguration) pulumi.StringPtrOutput { return v.UserData }).(pulumi.StringPtrOutput)
}

// The list of volume of the scaling configuration. The number of supported volumes ranges from 1 to 15.
func (o ScalingConfigurationOutput) Volumes() ScalingConfigurationVolumeArrayOutput {
	return o.ApplyT(func(v *ScalingConfiguration) ScalingConfigurationVolumeArrayOutput { return v.Volumes }).(ScalingConfigurationVolumeArrayOutput)
}

type ScalingConfigurationArrayOutput struct{ *pulumi.OutputState }

func (ScalingConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingConfiguration)(nil)).Elem()
}

func (o ScalingConfigurationArrayOutput) ToScalingConfigurationArrayOutput() ScalingConfigurationArrayOutput {
	return o
}

func (o ScalingConfigurationArrayOutput) ToScalingConfigurationArrayOutputWithContext(ctx context.Context) ScalingConfigurationArrayOutput {
	return o
}

func (o ScalingConfigurationArrayOutput) Index(i pulumi.IntInput) ScalingConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingConfiguration {
		return vs[0].([]*ScalingConfiguration)[vs[1].(int)]
	}).(ScalingConfigurationOutput)
}

type ScalingConfigurationMapOutput struct{ *pulumi.OutputState }

func (ScalingConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingConfiguration)(nil)).Elem()
}

func (o ScalingConfigurationMapOutput) ToScalingConfigurationMapOutput() ScalingConfigurationMapOutput {
	return o
}

func (o ScalingConfigurationMapOutput) ToScalingConfigurationMapOutputWithContext(ctx context.Context) ScalingConfigurationMapOutput {
	return o
}

func (o ScalingConfigurationMapOutput) MapIndex(k pulumi.StringInput) ScalingConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingConfiguration {
		return vs[0].(map[string]*ScalingConfiguration)[vs[1].(string)]
	}).(ScalingConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigurationInput)(nil)).Elem(), &ScalingConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigurationArrayInput)(nil)).Elem(), ScalingConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigurationMapInput)(nil)).Elem(), ScalingConfigurationMap{})
	pulumi.RegisterOutputType(ScalingConfigurationOutput{})
	pulumi.RegisterOutputType(ScalingConfigurationArrayOutput{})
	pulumi.RegisterOutputType(ScalingConfigurationMapOutput{})
}
