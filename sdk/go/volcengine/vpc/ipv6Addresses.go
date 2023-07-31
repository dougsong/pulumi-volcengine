// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vpc ipv6 addresses
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.Ipv6Addresses(ctx, &vpc.Ipv6AddressesArgs{
//				AssociatedInstanceId: pulumi.StringRef("i-yca53yuhj6gh9zl53kav"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Ipv6Addresses(ctx *pulumi.Context, args *Ipv6AddressesArgs, opts ...pulumi.InvokeOption) (*Ipv6AddressesResult, error) {
	var rv Ipv6AddressesResult
	err := ctx.Invoke("volcengine:vpc/ipv6Addresses:Ipv6Addresses", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Ipv6Addresses.
type Ipv6AddressesArgs struct {
	// The ID of the ECS instance that is assigned the IPv6 address.
	AssociatedInstanceId *string `pulumi:"associatedInstanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Ipv6Addresses.
type Ipv6AddressesResult struct {
	AssociatedInstanceId *string `pulumi:"associatedInstanceId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The collection of Ipv6Address query.
	Ipv6Addresses []Ipv6AddressesIpv6Address `pulumi:"ipv6Addresses"`
	OutputFile    *string                    `pulumi:"outputFile"`
	// The total count of Ipv6Address query.
	TotalCount int `pulumi:"totalCount"`
}

func Ipv6AddressesOutput(ctx *pulumi.Context, args Ipv6AddressesOutputArgs, opts ...pulumi.InvokeOption) Ipv6AddressesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (Ipv6AddressesResult, error) {
			args := v.(Ipv6AddressesArgs)
			r, err := Ipv6Addresses(ctx, &args, opts...)
			var s Ipv6AddressesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(Ipv6AddressesResultOutput)
}

// A collection of arguments for invoking Ipv6Addresses.
type Ipv6AddressesOutputArgs struct {
	// The ID of the ECS instance that is assigned the IPv6 address.
	AssociatedInstanceId pulumi.StringPtrInput `pulumi:"associatedInstanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (Ipv6AddressesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6AddressesArgs)(nil)).Elem()
}

// A collection of values returned by Ipv6Addresses.
type Ipv6AddressesResultOutput struct{ *pulumi.OutputState }

func (Ipv6AddressesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ipv6AddressesResult)(nil)).Elem()
}

func (o Ipv6AddressesResultOutput) ToIpv6AddressesResultOutput() Ipv6AddressesResultOutput {
	return o
}

func (o Ipv6AddressesResultOutput) ToIpv6AddressesResultOutputWithContext(ctx context.Context) Ipv6AddressesResultOutput {
	return o
}

func (o Ipv6AddressesResultOutput) AssociatedInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressesResult) *string { return v.AssociatedInstanceId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o Ipv6AddressesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Ipv6AddressesResult) string { return v.Id }).(pulumi.StringOutput)
}

// The collection of Ipv6Address query.
func (o Ipv6AddressesResultOutput) Ipv6Addresses() Ipv6AddressesIpv6AddressArrayOutput {
	return o.ApplyT(func(v Ipv6AddressesResult) []Ipv6AddressesIpv6Address { return v.Ipv6Addresses }).(Ipv6AddressesIpv6AddressArrayOutput)
}

func (o Ipv6AddressesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ipv6AddressesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of Ipv6Address query.
func (o Ipv6AddressesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v Ipv6AddressesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(Ipv6AddressesResultOutput{})
}
