// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package autoscaling

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of scaling instances
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
//			_, err := autoscaling.ScalingInstances(ctx, &autoscaling.ScalingInstancesArgs{
//				Ids: []string{
//					"i-ybzl4u5uogl8j07hgcbg",
//					"i-ybyncrcpzpgh9zmlct0w",
//					"i-ybyncrcpzogh9z4ax9tv",
//				},
//				ScalingConfigurationId: pulumi.StringRef("scc-ybtawzucw95pkgon0wst"),
//				ScalingGroupId:         "scg-ybtawtznszgh9yv8agcp",
//				Status:                 pulumi.StringRef("InService"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func ScalingInstances(ctx *pulumi.Context, args *ScalingInstancesArgs, opts ...pulumi.InvokeOption) (*ScalingInstancesResult, error) {
	var rv ScalingInstancesResult
	err := ctx.Invoke("volcengine:autoscaling/scalingInstances:ScalingInstances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ScalingInstances.
type ScalingInstancesArgs struct {
	// The creation type of the instances. Valid values: AutoCreated, Attached.
	CreationType *string `pulumi:"creationType"`
	// A list of instance ids.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The id of the scaling configuration id.
	ScalingConfigurationId *string `pulumi:"scalingConfigurationId"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.
	Status *string `pulumi:"status"`
}

// A collection of values returned by ScalingInstances.
type ScalingInstancesResult struct {
	// The creation type of the instance. Valid values: AutoCreated, Attached.
	CreationType *string `pulumi:"creationType"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The id of the scaling configuration.
	ScalingConfigurationId *string `pulumi:"scalingConfigurationId"`
	// The id of the scaling group.
	ScalingGroupId string `pulumi:"scalingGroupId"`
	// The collection of scaling instances.
	ScalingInstances []ScalingInstancesScalingInstance `pulumi:"scalingInstances"`
	// The status of instances.
	Status *string `pulumi:"status"`
	// The total count of scaling instances query.
	TotalCount int `pulumi:"totalCount"`
}

func ScalingInstancesOutput(ctx *pulumi.Context, args ScalingInstancesOutputArgs, opts ...pulumi.InvokeOption) ScalingInstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ScalingInstancesResult, error) {
			args := v.(ScalingInstancesArgs)
			r, err := ScalingInstances(ctx, &args, opts...)
			var s ScalingInstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ScalingInstancesResultOutput)
}

// A collection of arguments for invoking ScalingInstances.
type ScalingInstancesOutputArgs struct {
	// The creation type of the instances. Valid values: AutoCreated, Attached.
	CreationType pulumi.StringPtrInput `pulumi:"creationType"`
	// A list of instance ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The id of the scaling configuration id.
	ScalingConfigurationId pulumi.StringPtrInput `pulumi:"scalingConfigurationId"`
	// The id of the scaling group.
	ScalingGroupId pulumi.StringInput `pulumi:"scalingGroupId"`
	// The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (ScalingInstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingInstancesArgs)(nil)).Elem()
}

// A collection of values returned by ScalingInstances.
type ScalingInstancesResultOutput struct{ *pulumi.OutputState }

func (ScalingInstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScalingInstancesResult)(nil)).Elem()
}

func (o ScalingInstancesResultOutput) ToScalingInstancesResultOutput() ScalingInstancesResultOutput {
	return o
}

func (o ScalingInstancesResultOutput) ToScalingInstancesResultOutputWithContext(ctx context.Context) ScalingInstancesResultOutput {
	return o
}

// The creation type of the instance. Valid values: AutoCreated, Attached.
func (o ScalingInstancesResultOutput) CreationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.CreationType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o ScalingInstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ScalingInstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ScalingInstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScalingInstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o ScalingInstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The id of the scaling configuration.
func (o ScalingInstancesResultOutput) ScalingConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.ScalingConfigurationId }).(pulumi.StringPtrOutput)
}

// The id of the scaling group.
func (o ScalingInstancesResultOutput) ScalingGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ScalingInstancesResult) string { return v.ScalingGroupId }).(pulumi.StringOutput)
}

// The collection of scaling instances.
func (o ScalingInstancesResultOutput) ScalingInstances() ScalingInstancesScalingInstanceArrayOutput {
	return o.ApplyT(func(v ScalingInstancesResult) []ScalingInstancesScalingInstance { return v.ScalingInstances }).(ScalingInstancesScalingInstanceArrayOutput)
}

// The status of instances.
func (o ScalingInstancesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScalingInstancesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The total count of scaling instances query.
func (o ScalingInstancesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ScalingInstancesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ScalingInstancesResultOutput{})
}
