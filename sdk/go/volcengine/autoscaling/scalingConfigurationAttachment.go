// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage scaling configuration attachment
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/autoscaling"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := autoscaling.NewScalingConfigurationAttachment(ctx, "foo1", &autoscaling.ScalingConfigurationAttachmentArgs{
//				ScalingConfigurationId: pulumi.String("scc-ybrurj4uw6gh9zecj327"),
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
// Scaling Configuration attachment can be imported using the scaling_configuration_id e.g.
//
// ```sh
//
//	$ pulumi import volcengine:autoscaling/scalingConfigurationAttachment:ScalingConfigurationAttachment default enable:scc-ybrurj4uw6gh9zecj327
//
// ```
type ScalingConfigurationAttachment struct {
	pulumi.CustomResourceState

	// The id of the scaling configuration.
	ScalingConfigurationId pulumi.StringOutput `pulumi:"scalingConfigurationId"`
}

// NewScalingConfigurationAttachment registers a new resource with the given unique name, arguments, and options.
func NewScalingConfigurationAttachment(ctx *pulumi.Context,
	name string, args *ScalingConfigurationAttachmentArgs, opts ...pulumi.ResourceOption) (*ScalingConfigurationAttachment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ScalingConfigurationId == nil {
		return nil, errors.New("invalid value for required argument 'ScalingConfigurationId'")
	}
	var resource ScalingConfigurationAttachment
	err := ctx.RegisterResource("volcengine:autoscaling/scalingConfigurationAttachment:ScalingConfigurationAttachment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetScalingConfigurationAttachment gets an existing ScalingConfigurationAttachment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetScalingConfigurationAttachment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScalingConfigurationAttachmentState, opts ...pulumi.ResourceOption) (*ScalingConfigurationAttachment, error) {
	var resource ScalingConfigurationAttachment
	err := ctx.ReadResource("volcengine:autoscaling/scalingConfigurationAttachment:ScalingConfigurationAttachment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ScalingConfigurationAttachment resources.
type scalingConfigurationAttachmentState struct {
	// The id of the scaling configuration.
	ScalingConfigurationId *string `pulumi:"scalingConfigurationId"`
}

type ScalingConfigurationAttachmentState struct {
	// The id of the scaling configuration.
	ScalingConfigurationId pulumi.StringPtrInput
}

func (ScalingConfigurationAttachmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigurationAttachmentState)(nil)).Elem()
}

type scalingConfigurationAttachmentArgs struct {
	// The id of the scaling configuration.
	ScalingConfigurationId string `pulumi:"scalingConfigurationId"`
}

// The set of arguments for constructing a ScalingConfigurationAttachment resource.
type ScalingConfigurationAttachmentArgs struct {
	// The id of the scaling configuration.
	ScalingConfigurationId pulumi.StringInput
}

func (ScalingConfigurationAttachmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scalingConfigurationAttachmentArgs)(nil)).Elem()
}

type ScalingConfigurationAttachmentInput interface {
	pulumi.Input

	ToScalingConfigurationAttachmentOutput() ScalingConfigurationAttachmentOutput
	ToScalingConfigurationAttachmentOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentOutput
}

func (*ScalingConfigurationAttachment) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingConfigurationAttachment)(nil)).Elem()
}

func (i *ScalingConfigurationAttachment) ToScalingConfigurationAttachmentOutput() ScalingConfigurationAttachmentOutput {
	return i.ToScalingConfigurationAttachmentOutputWithContext(context.Background())
}

func (i *ScalingConfigurationAttachment) ToScalingConfigurationAttachmentOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationAttachmentOutput)
}

// ScalingConfigurationAttachmentArrayInput is an input type that accepts ScalingConfigurationAttachmentArray and ScalingConfigurationAttachmentArrayOutput values.
// You can construct a concrete instance of `ScalingConfigurationAttachmentArrayInput` via:
//
//	ScalingConfigurationAttachmentArray{ ScalingConfigurationAttachmentArgs{...} }
type ScalingConfigurationAttachmentArrayInput interface {
	pulumi.Input

	ToScalingConfigurationAttachmentArrayOutput() ScalingConfigurationAttachmentArrayOutput
	ToScalingConfigurationAttachmentArrayOutputWithContext(context.Context) ScalingConfigurationAttachmentArrayOutput
}

type ScalingConfigurationAttachmentArray []ScalingConfigurationAttachmentInput

func (ScalingConfigurationAttachmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingConfigurationAttachment)(nil)).Elem()
}

func (i ScalingConfigurationAttachmentArray) ToScalingConfigurationAttachmentArrayOutput() ScalingConfigurationAttachmentArrayOutput {
	return i.ToScalingConfigurationAttachmentArrayOutputWithContext(context.Background())
}

func (i ScalingConfigurationAttachmentArray) ToScalingConfigurationAttachmentArrayOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationAttachmentArrayOutput)
}

// ScalingConfigurationAttachmentMapInput is an input type that accepts ScalingConfigurationAttachmentMap and ScalingConfigurationAttachmentMapOutput values.
// You can construct a concrete instance of `ScalingConfigurationAttachmentMapInput` via:
//
//	ScalingConfigurationAttachmentMap{ "key": ScalingConfigurationAttachmentArgs{...} }
type ScalingConfigurationAttachmentMapInput interface {
	pulumi.Input

	ToScalingConfigurationAttachmentMapOutput() ScalingConfigurationAttachmentMapOutput
	ToScalingConfigurationAttachmentMapOutputWithContext(context.Context) ScalingConfigurationAttachmentMapOutput
}

type ScalingConfigurationAttachmentMap map[string]ScalingConfigurationAttachmentInput

func (ScalingConfigurationAttachmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingConfigurationAttachment)(nil)).Elem()
}

func (i ScalingConfigurationAttachmentMap) ToScalingConfigurationAttachmentMapOutput() ScalingConfigurationAttachmentMapOutput {
	return i.ToScalingConfigurationAttachmentMapOutputWithContext(context.Background())
}

func (i ScalingConfigurationAttachmentMap) ToScalingConfigurationAttachmentMapOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalingConfigurationAttachmentMapOutput)
}

type ScalingConfigurationAttachmentOutput struct{ *pulumi.OutputState }

func (ScalingConfigurationAttachmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScalingConfigurationAttachment)(nil)).Elem()
}

func (o ScalingConfigurationAttachmentOutput) ToScalingConfigurationAttachmentOutput() ScalingConfigurationAttachmentOutput {
	return o
}

func (o ScalingConfigurationAttachmentOutput) ToScalingConfigurationAttachmentOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentOutput {
	return o
}

// The id of the scaling configuration.
func (o ScalingConfigurationAttachmentOutput) ScalingConfigurationId() pulumi.StringOutput {
	return o.ApplyT(func(v *ScalingConfigurationAttachment) pulumi.StringOutput { return v.ScalingConfigurationId }).(pulumi.StringOutput)
}

type ScalingConfigurationAttachmentArrayOutput struct{ *pulumi.OutputState }

func (ScalingConfigurationAttachmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ScalingConfigurationAttachment)(nil)).Elem()
}

func (o ScalingConfigurationAttachmentArrayOutput) ToScalingConfigurationAttachmentArrayOutput() ScalingConfigurationAttachmentArrayOutput {
	return o
}

func (o ScalingConfigurationAttachmentArrayOutput) ToScalingConfigurationAttachmentArrayOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentArrayOutput {
	return o
}

func (o ScalingConfigurationAttachmentArrayOutput) Index(i pulumi.IntInput) ScalingConfigurationAttachmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ScalingConfigurationAttachment {
		return vs[0].([]*ScalingConfigurationAttachment)[vs[1].(int)]
	}).(ScalingConfigurationAttachmentOutput)
}

type ScalingConfigurationAttachmentMapOutput struct{ *pulumi.OutputState }

func (ScalingConfigurationAttachmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ScalingConfigurationAttachment)(nil)).Elem()
}

func (o ScalingConfigurationAttachmentMapOutput) ToScalingConfigurationAttachmentMapOutput() ScalingConfigurationAttachmentMapOutput {
	return o
}

func (o ScalingConfigurationAttachmentMapOutput) ToScalingConfigurationAttachmentMapOutputWithContext(ctx context.Context) ScalingConfigurationAttachmentMapOutput {
	return o
}

func (o ScalingConfigurationAttachmentMapOutput) MapIndex(k pulumi.StringInput) ScalingConfigurationAttachmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ScalingConfigurationAttachment {
		return vs[0].(map[string]*ScalingConfigurationAttachment)[vs[1].(string)]
	}).(ScalingConfigurationAttachmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigurationAttachmentInput)(nil)).Elem(), &ScalingConfigurationAttachment{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigurationAttachmentArrayInput)(nil)).Elem(), ScalingConfigurationAttachmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScalingConfigurationAttachmentMapInput)(nil)).Elem(), ScalingConfigurationAttachmentMap{})
	pulumi.RegisterOutputType(ScalingConfigurationAttachmentOutput{})
	pulumi.RegisterOutputType(ScalingConfigurationAttachmentArrayOutput{})
	pulumi.RegisterOutputType(ScalingConfigurationAttachmentMapOutput{})
}
