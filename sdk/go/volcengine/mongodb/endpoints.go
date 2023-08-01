// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mongodb endpoints
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
//			_, err := mongodb.Endpoints(ctx, &mongodb.EndpointsArgs{
//				InstanceId: pulumi.StringRef("mongo-shard-xxx"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Endpoints(ctx *pulumi.Context, args *EndpointsArgs, opts ...pulumi.InvokeOption) (*EndpointsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv EndpointsResult
	err := ctx.Invoke("volcengine:mongodb/endpoints:Endpoints", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Endpoints.
type EndpointsArgs struct {
	// The instance ID to query.
	InstanceId *string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Endpoints.
type EndpointsResult struct {
	// The collection of mongodb endpoints query.
	Endpoints []EndpointsEndpoint `pulumi:"endpoints"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of mongodb endpoint query.
	TotalCount int `pulumi:"totalCount"`
}

func EndpointsOutput(ctx *pulumi.Context, args EndpointsOutputArgs, opts ...pulumi.InvokeOption) EndpointsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (EndpointsResult, error) {
			args := v.(EndpointsArgs)
			r, err := Endpoints(ctx, &args, opts...)
			var s EndpointsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(EndpointsResultOutput)
}

// A collection of arguments for invoking Endpoints.
type EndpointsOutputArgs struct {
	// The instance ID to query.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (EndpointsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsArgs)(nil)).Elem()
}

// A collection of values returned by Endpoints.
type EndpointsResultOutput struct{ *pulumi.OutputState }

func (EndpointsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResult)(nil)).Elem()
}

func (o EndpointsResultOutput) ToEndpointsResultOutput() EndpointsResultOutput {
	return o
}

func (o EndpointsResultOutput) ToEndpointsResultOutputWithContext(ctx context.Context) EndpointsResultOutput {
	return o
}

// The collection of mongodb endpoints query.
func (o EndpointsResultOutput) Endpoints() EndpointsEndpointArrayOutput {
	return o.ApplyT(func(v EndpointsResult) []EndpointsEndpoint { return v.Endpoints }).(EndpointsEndpointArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o EndpointsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o EndpointsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o EndpointsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EndpointsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of mongodb endpoint query.
func (o EndpointsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v EndpointsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(EndpointsResultOutput{})
}
