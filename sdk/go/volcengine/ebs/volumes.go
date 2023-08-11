// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of volumes
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ebs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ebs.Volumes(ctx, &ebs.VolumesArgs{
//				Ids: []string{
//					"vol-3tzg6y5imn3b9fchkedb",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Volumes(ctx *pulumi.Context, args *VolumesArgs, opts ...pulumi.InvokeOption) (*VolumesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv VolumesResult
	err := ctx.Invoke("volcengine:ebs/volumes:Volumes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Volumes.
type VolumesArgs struct {
	// A list of Volume IDs.
	Ids []string `pulumi:"ids"`
	// The Id of instance.
	InstanceId *string `pulumi:"instanceId"`
	// The Kind of Volume.
	Kind *string `pulumi:"kind"`
	// A Name Regex of Volume.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The name of Volume.
	VolumeName *string `pulumi:"volumeName"`
	// The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
	VolumeStatus *string `pulumi:"volumeStatus"`
	// The type of Volume.
	VolumeType *string `pulumi:"volumeType"`
	// The Id of Zone.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by Volumes.
type VolumesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	InstanceId *string  `pulumi:"instanceId"`
	Kind       *string  `pulumi:"kind"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The total count of Volume query.
	TotalCount   int     `pulumi:"totalCount"`
	VolumeName   *string `pulumi:"volumeName"`
	VolumeStatus *string `pulumi:"volumeStatus"`
	VolumeType   *string `pulumi:"volumeType"`
	// The collection of Volume query.
	Volumes []VolumesVolume `pulumi:"volumes"`
	ZoneId  *string         `pulumi:"zoneId"`
}

func VolumesOutput(ctx *pulumi.Context, args VolumesOutputArgs, opts ...pulumi.InvokeOption) VolumesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VolumesResult, error) {
			args := v.(VolumesArgs)
			r, err := Volumes(ctx, &args, opts...)
			var s VolumesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VolumesResultOutput)
}

// A collection of arguments for invoking Volumes.
type VolumesOutputArgs struct {
	// A list of Volume IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Id of instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// The Kind of Volume.
	Kind pulumi.StringPtrInput `pulumi:"kind"`
	// A Name Regex of Volume.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The name of Volume.
	VolumeName pulumi.StringPtrInput `pulumi:"volumeName"`
	// The Status of Volume, the value can be `available` or `attaching` or `attached` or `detaching` or `creating` or `deleting` or `error` or `extending`.
	VolumeStatus pulumi.StringPtrInput `pulumi:"volumeStatus"`
	// The type of Volume.
	VolumeType pulumi.StringPtrInput `pulumi:"volumeType"`
	// The Id of Zone.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (VolumesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumesArgs)(nil)).Elem()
}

// A collection of values returned by Volumes.
type VolumesResultOutput struct{ *pulumi.OutputState }

func (VolumesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumesResult)(nil)).Elem()
}

func (o VolumesResultOutput) ToVolumesResultOutput() VolumesResultOutput {
	return o
}

func (o VolumesResultOutput) ToVolumesResultOutputWithContext(ctx context.Context) VolumesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o VolumesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VolumesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VolumesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o VolumesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o VolumesResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VolumesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o VolumesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of Volume query.
func (o VolumesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VolumesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o VolumesResultOutput) VolumeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.VolumeName }).(pulumi.StringPtrOutput)
}

func (o VolumesResultOutput) VolumeStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.VolumeStatus }).(pulumi.StringPtrOutput)
}

func (o VolumesResultOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

// The collection of Volume query.
func (o VolumesResultOutput) Volumes() VolumesVolumeArrayOutput {
	return o.ApplyT(func(v VolumesResult) []VolumesVolume { return v.Volumes }).(VolumesVolumeArrayOutput)
}

func (o VolumesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumesResultOutput{})
}
