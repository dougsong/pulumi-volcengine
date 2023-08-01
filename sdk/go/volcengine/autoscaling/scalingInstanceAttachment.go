// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage scaling instance attachment
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := autoscaling.NewScalingInstanceAttachment(ctx, "foo", &autoscaling.ScalingInstanceAttachmentArgs{
//				DeleteType:     pulumi.String("Remove"),
//				DetachOption:   pulumi.String("none"),
//				Entrusted:      pulumi.Bool(true),
//				InstanceId:     pulumi.String("i-yc23soxj50gsnz7rxnjp"),
//				ScalingGroupId: pulumi.String("scg-yc23rtcea88hcchybf8g"),
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
// Scaling instance attachment can be imported using the scaling_group_id and instance_id, e.g. You can choose to remove or detach the instance according to the `delete_type` field.
//
// ```sh
//
//	$ pulumi import volcengine:autoscaling/scalingInstanceAttachment:ScalingInstanceAttachment default scg-mizl7m1kqccg5smt1bdpijuj:i-l8u2ai4j0fauo6mrpgk8
//
// ```
type ScalingInstanceAttachment struct {
	pulumi.CustomResourceState

	// The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
	DeleteType pulumi.StringPtrOutput `pulumi:"deleteType"`
	// Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
	DetachOption pulumi.StringPtrOutput `pulumi:"detachOption"`
	// Whether to host the instance to a scaling group. Default value is false.
	Entrusted pulumi.BoolPtrOutput `pulumi:"entrusted"`
	// The id of the instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The id of the scaling group.
	ScalingGroupId pulumi.StringOutput `pulumi:"scalingGroupId"`
}

// NewScalingInstanceAttachment registers a new resource with the given unique name, arguments, and options.
func NewScalingInstanceAttachment(ctx *pulumi.Context,
	name string, args *ScalingInstanceAttachmentArgs, opts ...pulumi.ResourceOption) (*ScalingInstanceAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ScalingGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingGroupId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ScalingInstanceAttachment
	err := ctx.RegisterResource("volcengine:autoscaling/scalingInstanceAttachment:ScalingInstanceAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingInstanceAttachment gets an existing ScalingInstanceAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingInstanceAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingInstanceAttachmentState, opts ...pulumi.ResourceOption) (*ScalingInstanceAttachment, error) {
	var resource ScalingInstanceAttachment
	err := ctx.ReadResource("volcengine:autoscaling/scalingInstanceAttachment:ScalingInstanceAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingInstanceAttachment resources.
type scalingInstanceAttachmentState struct {
	// The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
	DeleteType *string `pulumi:"deleteType"`
	// Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
	DetachOption *string `pulumi:"detachOption"`
	// Whether to host the instance to a scaling group. Default value is false.
	Entrusted *bool `pulumi:"entrusted"`
	// The id of the instance.
	InstanceId *string `pulumi:"instanceId"`
	// The id of the scaling group.
	ScalingGroupId *string `pulumi:"scalingGroupId"`
}

type ScalingInstanceAttachmentState struct {
	// The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
	DeleteType pulumi.StringPtrInput
	// Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
	DetachOption pulumi.StringPtrInput
	// Whether to host the instance to a scaling group. Default value is false.
	Entrusted pulumi.BoolPtrInput
	// The id of the instance.
	InstanceId pulumi.StringPtrInput
	// The id of the scaling group.
	ScalingGroupId pulumi.StringPtrInput
}

func (ScalingInstanceAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingInstanceAttachmentState)(nil)).Elem()
}

type scalingInstanceAttachmentArgs struct {
	// The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
	DeleteType *string `pulumi:"deleteType"`
	// Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
	DetachOption *string `pulumi:"detachOption"`
	// Whether to host the instance to a scaling group. Default value is false.
	Entrusted *bool `pulumi:"entrusted"`
	// The id of the instance.
	InstanceId string `pulumi:"instanceId"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
}

// The set of arguments for constructing a ScalingInstanceAttachment resource.
type ScalingInstanceAttachmentArgs struct {
	// The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
	DeleteType pulumi.StringPtrInput
	// Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
	DetachOption pulumi.StringPtrInput
	// Whether to host the instance to a scaling group. Default value is false.
	Entrusted pulumi.BoolPtrInput
	// The id of the instance.
	InstanceId pulumi.StringInput
	// The id of the scaling group.
	ScalingGroupId pulumi.StringInput
}

func (ScalingInstanceAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingInstanceAttachmentArgs)(nil)).Elem()
}

type ScalingInstanceAttachmentInput interface {
	pulumi.Input

	ToScalingInstanceAttachmentOutput() ScalingInstanceAttachmentOutput
	ToScalingInstanceAttachmentOutputWithContext(ctx context.Context) ScalingInstanceAttachmentOutput
}

func (*ScalingInstanceAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingInstanceAttachment)(nil)).Elem()
}

func (i *ScalingInstanceAttachment) ToScalingInstanceAttachmentOutput() ScalingInstanceAttachmentOutput {
	return i.ToScalingInstanceAttachmentOutputWithContext(context.Background())
}

func (i *ScalingInstanceAttachment) ToScalingInstanceAttachmentOutputWithContext(ctx context.Context) ScalingInstanceAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingInstanceAttachmentOutput)
}

// ScalingInstanceAttachmentArrayInput is an input type that accepts ScalingInstanceAttachmentArray and ScalingInstanceAttachmentArrayOutput values.
// You can construct a concrete instance of `ScalingInstanceAttachmentArrayInput` via:
//
//	ScalingInstanceAttachmentArray{ ScalingInstanceAttachmentArgs{...} }
type ScalingInstanceAttachmentArrayInput interface {
	pulumi.Input

	ToScalingInstanceAttachmentArrayOutput() ScalingInstanceAttachmentArrayOutput
	ToScalingInstanceAttachmentArrayOutputWithContext(context.Context) ScalingInstanceAttachmentArrayOutput
}

type ScalingInstanceAttachmentArray []ScalingInstanceAttachmentInput

func (ScalingInstanceAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingInstanceAttachment)(nil)).Elem()
}

func (i ScalingInstanceAttachmentArray) ToScalingInstanceAttachmentArrayOutput() ScalingInstanceAttachmentArrayOutput {
	return i.ToScalingInstanceAttachmentArrayOutputWithContext(context.Background())
}

func (i ScalingInstanceAttachmentArray) ToScalingInstanceAttachmentArrayOutputWithContext(ctx context.Context) ScalingInstanceAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingInstanceAttachmentArrayOutput)
}

// ScalingInstanceAttachmentMapInput is an input type that accepts ScalingInstanceAttachmentMap and ScalingInstanceAttachmentMapOutput values.
// You can construct a concrete instance of `ScalingInstanceAttachmentMapInput` via:
//
//	ScalingInstanceAttachmentMap{ "key": ScalingInstanceAttachmentArgs{...} }
type ScalingInstanceAttachmentMapInput interface {
	pulumi.Input

	ToScalingInstanceAttachmentMapOutput() ScalingInstanceAttachmentMapOutput
	ToScalingInstanceAttachmentMapOutputWithContext(context.Context) ScalingInstanceAttachmentMapOutput
}

type ScalingInstanceAttachmentMap map[string]ScalingInstanceAttachmentInput

func (ScalingInstanceAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingInstanceAttachment)(nil)).Elem()
}

func (i ScalingInstanceAttachmentMap) ToScalingInstanceAttachmentMapOutput() ScalingInstanceAttachmentMapOutput {
	return i.ToScalingInstanceAttachmentMapOutputWithContext(context.Background())
}

func (i ScalingInstanceAttachmentMap) ToScalingInstanceAttachmentMapOutputWithContext(ctx context.Context) ScalingInstanceAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingInstanceAttachmentMapOutput)
}

type ScalingInstanceAttachmentOutput struct{ *pulumi.OutputState }

func (ScalingInstanceAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingInstanceAttachment)(nil)).Elem()
}

func (o ScalingInstanceAttachmentOutput) ToScalingInstanceAttachmentOutput() ScalingInstanceAttachmentOutput {
	return o
}

func (o ScalingInstanceAttachmentOutput) ToScalingInstanceAttachmentOutputWithContext(ctx context.Context) ScalingInstanceAttachmentOutput {
	return o
}

// The type of delete activity. Valid values: Remove, Detach. Default value is Remove.
func (o ScalingInstanceAttachmentOutput) DeleteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingInstanceAttachment) pulumi.StringPtrOutput { return v.DeleteType }).(pulumi.StringPtrOutput)
}

// Whether to cancel the association of the instance with the load balancing and public network IP. Valid values: both, none. Default value is both.
func (o ScalingInstanceAttachmentOutput) DetachOption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScalingInstanceAttachment) pulumi.StringPtrOutput { return v.DetachOption }).(pulumi.StringPtrOutput)
}

// Whether to host the instance to a scaling group. Default value is false.
func (o ScalingInstanceAttachmentOutput) Entrusted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScalingInstanceAttachment) pulumi.BoolPtrOutput { return v.Entrusted }).(pulumi.BoolPtrOutput)
}

// The id of the instance.
func (o ScalingInstanceAttachmentOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingInstanceAttachment) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The id of the scaling group.
func (o ScalingInstanceAttachmentOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingInstanceAttachment) pulumi.StringOutput { return v.ScalingGroupId }).(pulumi.StringOutput)
}

type ScalingInstanceAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ScalingInstanceAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingInstanceAttachment)(nil)).Elem()
}

func (o ScalingInstanceAttachmentArrayOutput) ToScalingInstanceAttachmentArrayOutput() ScalingInstanceAttachmentArrayOutput {
	return o
}

func (o ScalingInstanceAttachmentArrayOutput) ToScalingInstanceAttachmentArrayOutputWithContext(ctx context.Context) ScalingInstanceAttachmentArrayOutput {
	return o
}

func (o ScalingInstanceAttachmentArrayOutput) Index(i pulumi.IntInput) ScalingInstanceAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingInstanceAttachment {
		return vs[0].([]*ScalingInstanceAttachment)[vs[1].(int)]
	}).(ScalingInstanceAttachmentOutput)
}

type ScalingInstanceAttachmentMapOutput struct{ *pulumi.OutputState }

func (ScalingInstanceAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingInstanceAttachment)(nil)).Elem()
}

func (o ScalingInstanceAttachmentMapOutput) ToScalingInstanceAttachmentMapOutput() ScalingInstanceAttachmentMapOutput {
	return o
}

func (o ScalingInstanceAttachmentMapOutput) ToScalingInstanceAttachmentMapOutputWithContext(ctx context.Context) ScalingInstanceAttachmentMapOutput {
	return o
}

func (o ScalingInstanceAttachmentMapOutput) MapIndex(k pulumi.StringInput) ScalingInstanceAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingInstanceAttachment {
		return vs[0].(map[string]*ScalingInstanceAttachment)[vs[1].(string)]
	}).(ScalingInstanceAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingInstanceAttachmentInput)(nil)).Elem(), &ScalingInstanceAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingInstanceAttachmentArrayInput)(nil)).Elem(), ScalingInstanceAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingInstanceAttachmentMapInput)(nil)).Elem(), ScalingInstanceAttachmentMap{})
	pulumi.RegisterOutputType(ScalingInstanceAttachmentOutput{})
	pulumi.RegisterOutputType(ScalingInstanceAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ScalingInstanceAttachmentMapOutput{})
}
