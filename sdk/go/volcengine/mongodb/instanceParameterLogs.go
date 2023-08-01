// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mongodb instance parameter logs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/mongodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/mongodb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.InstanceParameterLogs(ctx, &mongodb.InstanceParameterLogsArgs{
//				EndTime:    "2023-11-14 18:15Z",
//				InstanceId: "mongo-replica-f16e9298b121",
//				StartTime:  "2022-11-14 00:00Z",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func InstanceParameterLogs(ctx *pulumi.Context, args *InstanceParameterLogsArgs, opts ...pulumi.InvokeOption) (*InstanceParameterLogsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv InstanceParameterLogsResult
	err := ctx.Invoke("volcengine:mongodb/instanceParameterLogs:InstanceParameterLogs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking InstanceParameterLogs.
type InstanceParameterLogsArgs struct {
	// The end time to query.
	EndTime string `pulumi:"endTime"`
	// The instance ID to query.
	InstanceId string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The start time to query.
	StartTime string `pulumi:"startTime"`
}

// A collection of values returned by InstanceParameterLogs.
type InstanceParameterLogsResult struct {
	EndTime string `pulumi:"endTime"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of parameter change log query.
	ParameterChangeLogs InstanceParameterLogsParameterChangeLogs `pulumi:"parameterChangeLogs"`
	StartTime           string                                   `pulumi:"startTime"`
	// The total count of mongodb instance parameter log query.
	TotalCount int `pulumi:"totalCount"`
}

func InstanceParameterLogsOutput(ctx *pulumi.Context, args InstanceParameterLogsOutputArgs, opts ...pulumi.InvokeOption) InstanceParameterLogsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (InstanceParameterLogsResult, error) {
			args := v.(InstanceParameterLogsArgs)
			r, err := InstanceParameterLogs(ctx, &args, opts...)
			var s InstanceParameterLogsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(InstanceParameterLogsResultOutput)
}

// A collection of arguments for invoking InstanceParameterLogs.
type InstanceParameterLogsOutputArgs struct {
	// The end time to query.
	EndTime pulumi.StringInput `pulumi:"endTime"`
	// The instance ID to query.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The start time to query.
	StartTime pulumi.StringInput `pulumi:"startTime"`
}

func (InstanceParameterLogsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceParameterLogsArgs)(nil)).Elem()
}

// A collection of values returned by InstanceParameterLogs.
type InstanceParameterLogsResultOutput struct{ *pulumi.OutputState }

func (InstanceParameterLogsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceParameterLogsResult)(nil)).Elem()
}

func (o InstanceParameterLogsResultOutput) ToInstanceParameterLogsResultOutput() InstanceParameterLogsResultOutput {
	return o
}

func (o InstanceParameterLogsResultOutput) ToInstanceParameterLogsResultOutputWithContext(ctx context.Context) InstanceParameterLogsResultOutput {
	return o
}

func (o InstanceParameterLogsResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o InstanceParameterLogsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o InstanceParameterLogsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o InstanceParameterLogsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of parameter change log query.
func (o InstanceParameterLogsResultOutput) ParameterChangeLogs() InstanceParameterLogsParameterChangeLogsOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) InstanceParameterLogsParameterChangeLogs {
		return v.ParameterChangeLogs
	}).(InstanceParameterLogsParameterChangeLogsOutput)
}

func (o InstanceParameterLogsResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The total count of mongodb instance parameter log query.
func (o InstanceParameterLogsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v InstanceParameterLogsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(InstanceParameterLogsResultOutput{})
}
