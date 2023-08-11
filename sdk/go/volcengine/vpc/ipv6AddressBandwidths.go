// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vpc ipv6 address bandwidths
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.Ipv6AddressBandwidths(ctx, &vpc.Ipv6AddressBandwidthsArgs{
//				Ids: []string{
//					"eip-in2y2duvtlhc8gbssyfnhfre",
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
func Ipv6AddressBandwidths(ctx *pulumi.Context, args *Ipv6AddressBandwidthsArgs, opts ...pulumi.InvokeOption) (*Ipv6AddressBandwidthsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv Ipv6AddressBandwidthsResult
	err := ctx.Invoke("volcengine:vpc/ipv6AddressBandwidths:Ipv6AddressBandwidths", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Ipv6AddressBandwidths.
type Ipv6AddressBandwidthsArgs struct {
	// The ID of the associated instance.
	AssociatedInstanceId *string `pulumi:"associatedInstanceId"`
	// The type of the associated instance.
	AssociatedInstanceType *string `pulumi:"associatedInstanceType"`
	// Allocation IDs of the Ipv6 address width.
	Ids []string `pulumi:"ids"`
	// The ipv6 addresses.
	Ipv6Addresses []string `pulumi:"ipv6Addresses"`
	// ISP of the ipv6 address.
	Isp *string `pulumi:"isp"`
	// The network type of the ipv6 address.
	NetworkType *string `pulumi:"networkType"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ID of Vpc the ipv6 address in.
	VpcId *string `pulumi:"vpcId"`
}

// A collection of values returned by Ipv6AddressBandwidths.
type Ipv6AddressBandwidthsResult struct {
	AssociatedInstanceId   *string `pulumi:"associatedInstanceId"`
	AssociatedInstanceType *string `pulumi:"associatedInstanceType"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The collection of Ipv6AddressBandwidth query.
	Ipv6AddressBandwidths []Ipv6AddressBandwidthsIpv6AddressBandwidth `pulumi:"ipv6AddressBandwidths"`
	Ipv6Addresses         []string                                    `pulumi:"ipv6Addresses"`
	// The ISP of the Ipv6AddressBandwidth.
	Isp *string `pulumi:"isp"`
	// The network type of the Ipv6AddressBandwidth.
	NetworkType *string `pulumi:"networkType"`
	OutputFile  *string `pulumi:"outputFile"`
	// The total count of Ipv6AddressBandwidth query.
	TotalCount int     `pulumi:"totalCount"`
	VpcId      *string `pulumi:"vpcId"`
}

func Ipv6AddressBandwidthsOutput(ctx *pulumi.Context, args Ipv6AddressBandwidthsOutputArgs, opts ...pulumi.InvokeOption) Ipv6AddressBandwidthsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Ipv6AddressBandwidthsResult, error) {
			args := v.(Ipv6AddressBandwidthsArgs)
			r, err := Ipv6AddressBandwidths(ctx, &args, opts...)
			var s Ipv6AddressBandwidthsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(Ipv6AddressBandwidthsResultOutput)
}

// A collection of arguments for invoking Ipv6AddressBandwidths.
type Ipv6AddressBandwidthsOutputArgs struct {
	// The ID of the associated instance.
	AssociatedInstanceId pulumi.StringPtrInput `pulumi:"associatedInstanceId"`
	// The type of the associated instance.
	AssociatedInstanceType pulumi.StringPtrInput `pulumi:"associatedInstanceType"`
	// Allocation IDs of the Ipv6 address width.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The ipv6 addresses.
	Ipv6Addresses pulumi.StringArrayInput `pulumi:"ipv6Addresses"`
	// ISP of the ipv6 address.
	Isp pulumi.StringPtrInput `pulumi:"isp"`
	// The network type of the ipv6 address.
	NetworkType pulumi.StringPtrInput `pulumi:"networkType"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of Vpc the ipv6 address in.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
}

func (Ipv6AddressBandwidthsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6AddressBandwidthsArgs)(nil)).Elem()
}

// A collection of values returned by Ipv6AddressBandwidths.
type Ipv6AddressBandwidthsResultOutput struct{ *pulumi.OutputState }

func (Ipv6AddressBandwidthsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6AddressBandwidthsResult)(nil)).Elem()
}

func (o Ipv6AddressBandwidthsResultOutput) ToIpv6AddressBandwidthsResultOutput() Ipv6AddressBandwidthsResultOutput {
	return o
}

func (o Ipv6AddressBandwidthsResultOutput) ToIpv6AddressBandwidthsResultOutputWithContext(ctx context.Context) Ipv6AddressBandwidthsResultOutput {
	return o
}

func (o Ipv6AddressBandwidthsResultOutput) AssociatedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) *string { return v.AssociatedInstanceId }).(pulumi.StringPtrOutput)
}

func (o Ipv6AddressBandwidthsResultOutput) AssociatedInstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) *string { return v.AssociatedInstanceType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o Ipv6AddressBandwidthsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o Ipv6AddressBandwidthsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The collection of Ipv6AddressBandwidth query.
func (o Ipv6AddressBandwidthsResultOutput) Ipv6AddressBandwidths() Ipv6AddressBandwidthsIpv6AddressBandwidthArrayOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) []Ipv6AddressBandwidthsIpv6AddressBandwidth {
		return v.Ipv6AddressBandwidths
	}).(Ipv6AddressBandwidthsIpv6AddressBandwidthArrayOutput)
}

func (o Ipv6AddressBandwidthsResultOutput) Ipv6Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) []string { return v.Ipv6Addresses }).(pulumi.StringArrayOutput)
}

// The ISP of the Ipv6AddressBandwidth.
func (o Ipv6AddressBandwidthsResultOutput) Isp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) *string { return v.Isp }).(pulumi.StringPtrOutput)
}

// The network type of the Ipv6AddressBandwidth.
func (o Ipv6AddressBandwidthsResultOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) *string { return v.NetworkType }).(pulumi.StringPtrOutput)
}

func (o Ipv6AddressBandwidthsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of Ipv6AddressBandwidth query.
func (o Ipv6AddressBandwidthsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o Ipv6AddressBandwidthsResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressBandwidthsResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(Ipv6AddressBandwidthsResultOutput{})
}
