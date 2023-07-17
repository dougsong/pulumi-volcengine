// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of clb zones
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.Zones(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Zones(ctx *pulumi.Context, args *ZonesArgs, opts ...pulumi.InvokeOption) (*ZonesResult, error) {
	var rv ZonesResult
	err := ctx.Invoke("volcengine:clb/zones:Zones", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Zones.
type ZonesArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Zones.
type ZonesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The master zones list.
	MasterZones []ZonesMasterZone `pulumi:"masterZones"`
	OutputFile  *string           `pulumi:"outputFile"`
	// The total count of zone query.
	TotalCount int `pulumi:"totalCount"`
}

func ZonesOutput(ctx *pulumi.Context, args ZonesOutputArgs, opts ...pulumi.InvokeOption) ZonesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ZonesResult, error) {
			args := v.(ZonesArgs)
			r, err := Zones(ctx, &args, opts...)
			var s ZonesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ZonesResultOutput)
}

// A collection of arguments for invoking Zones.
type ZonesOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (ZonesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonesArgs)(nil)).Elem()
}

// A collection of values returned by Zones.
type ZonesResultOutput struct{ *pulumi.OutputState }

func (ZonesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ZonesResult)(nil)).Elem()
}

func (o ZonesResultOutput) ToZonesResultOutput() ZonesResultOutput {
	return o
}

func (o ZonesResultOutput) ToZonesResultOutputWithContext(ctx context.Context) ZonesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o ZonesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ZonesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The master zones list.
func (o ZonesResultOutput) MasterZones() ZonesMasterZoneArrayOutput {
	return o.ApplyT(func(v ZonesResult) []ZonesMasterZone { return v.MasterZones }).(ZonesMasterZoneArrayOutput)
}

func (o ZonesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ZonesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of zone query.
func (o ZonesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ZonesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ZonesResultOutput{})
}
