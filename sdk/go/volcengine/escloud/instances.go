// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package escloud

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of escloud instances
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/escloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/escloud"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := escloud.Instances(ctx, &escloud.InstancesArgs{
//				Ids: []string{
//					"d3gftqjvnah74eie",
//				},
//				Statuses: []string{
//					"Running",
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
func Instances(ctx *pulumi.Context, args *InstancesArgs, opts ...pulumi.InvokeOption) (*InstancesResult, error) {
	var rv InstancesResult
	err := ctx.Invoke("volcengine:escloud/instances:Instances", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Instances.
type InstancesArgs struct {
	// The charge types of instance.
	ChargeTypes []string `pulumi:"chargeTypes"`
	// A list of instance IDs.
	Ids []string `pulumi:"ids"`
	// The names of instance.
	Names []string `pulumi:"names"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The list status of instance.
	Statuses []string `pulumi:"statuses"`
	// The versions of instance.
	Versions []string `pulumi:"versions"`
	// The available zone IDs of instance.
	ZoneIds []string `pulumi:"zoneIds"`
}

// A collection of values returned by Instances.
type InstancesResult struct {
	ChargeTypes []string `pulumi:"chargeTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The collection of instance query.
	Instances  []InstancesInstance `pulumi:"instances"`
	Names      []string            `pulumi:"names"`
	OutputFile *string             `pulumi:"outputFile"`
	Statuses   []string            `pulumi:"statuses"`
	// The total count of instance query.
	TotalCount int      `pulumi:"totalCount"`
	Versions   []string `pulumi:"versions"`
	ZoneIds    []string `pulumi:"zoneIds"`
}

func InstancesOutput(ctx *pulumi.Context, args InstancesOutputArgs, opts ...pulumi.InvokeOption) InstancesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (InstancesResult, error) {
			args := v.(InstancesArgs)
			r, err := Instances(ctx, &args, opts...)
			var s InstancesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(InstancesResultOutput)
}

// A collection of arguments for invoking Instances.
type InstancesOutputArgs struct {
	// The charge types of instance.
	ChargeTypes pulumi.StringArrayInput `pulumi:"chargeTypes"`
	// A list of instance IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The names of instance.
	Names pulumi.StringArrayInput `pulumi:"names"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The list status of instance.
	Statuses pulumi.StringArrayInput `pulumi:"statuses"`
	// The versions of instance.
	Versions pulumi.StringArrayInput `pulumi:"versions"`
	// The available zone IDs of instance.
	ZoneIds pulumi.StringArrayInput `pulumi:"zoneIds"`
}

func (InstancesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesArgs)(nil)).Elem()
}

// A collection of values returned by Instances.
type InstancesResultOutput struct{ *pulumi.OutputState }

func (InstancesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstancesResult)(nil)).Elem()
}

func (o InstancesResultOutput) ToInstancesResultOutput() InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) ToInstancesResultOutputWithContext(ctx context.Context) InstancesResultOutput {
	return o
}

func (o InstancesResultOutput) ChargeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.ChargeTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o InstancesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstancesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o InstancesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The collection of instance query.
func (o InstancesResultOutput) Instances() InstancesInstanceArrayOutput {
	return o.ApplyT(func(v InstancesResult) []InstancesInstance { return v.Instances }).(InstancesInstanceArrayOutput)
}

func (o InstancesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o InstancesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstancesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o InstancesResultOutput) Statuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.Statuses }).(pulumi.StringArrayOutput)
}

// The total count of instance query.
func (o InstancesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v InstancesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o InstancesResultOutput) Versions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.Versions }).(pulumi.StringArrayOutput)
}

func (o InstancesResultOutput) ZoneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstancesResult) []string { return v.ZoneIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(InstancesResultOutput{})
}
