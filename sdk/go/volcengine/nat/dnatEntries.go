// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package nat

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of dnat entries
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/nat"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := nat.DnatEntries(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func DnatEntries(ctx *pulumi.Context, args *DnatEntriesArgs, opts ...pulumi.InvokeOption) (*DnatEntriesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv DnatEntriesResult
	err := ctx.Invoke("volcengine:nat/dnatEntries:DnatEntries", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking DnatEntries.
type DnatEntriesArgs struct {
	// The name of the DNAT entry.
	DnatEntryName *string `pulumi:"dnatEntryName"`
	// Provides the public IP address for public network access.
	ExternalIp *string `pulumi:"externalIp"`
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort *string `pulumi:"externalPort"`
	// A list of DNAT entry ids.
	Ids []string `pulumi:"ids"`
	// Provides the internal IP address.
	InternalIp *string `pulumi:"internalIp"`
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort *string `pulumi:"internalPort"`
	// The id of the NAT gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The network protocol.
	Protocol *string `pulumi:"protocol"`
}

// A collection of values returned by DnatEntries.
type DnatEntriesResult struct {
	// List of DNAT entries.
	DnatEntries []DnatEntriesDnatEntry `pulumi:"dnatEntries"`
	// The name of the DNAT entry.
	DnatEntryName *string `pulumi:"dnatEntryName"`
	// Provides the public IP address for public network access.
	ExternalIp *string `pulumi:"externalIp"`
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort *string `pulumi:"externalPort"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// Provides the internal IP address.
	InternalIp *string `pulumi:"internalIp"`
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort *string `pulumi:"internalPort"`
	// The ID of the NAT gateway.
	NatGatewayId *string `pulumi:"natGatewayId"`
	OutputFile   *string `pulumi:"outputFile"`
	// The network protocol.
	Protocol *string `pulumi:"protocol"`
	// The total count of snat entries query.
	TotalCount int `pulumi:"totalCount"`
}

func DnatEntriesOutput(ctx *pulumi.Context, args DnatEntriesOutputArgs, opts ...pulumi.InvokeOption) DnatEntriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (DnatEntriesResult, error) {
			args := v.(DnatEntriesArgs)
			r, err := DnatEntries(ctx, &args, opts...)
			var s DnatEntriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(DnatEntriesResultOutput)
}

// A collection of arguments for invoking DnatEntries.
type DnatEntriesOutputArgs struct {
	// The name of the DNAT entry.
	DnatEntryName pulumi.StringPtrInput `pulumi:"dnatEntryName"`
	// Provides the public IP address for public network access.
	ExternalIp pulumi.StringPtrInput `pulumi:"externalIp"`
	// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
	ExternalPort pulumi.StringPtrInput `pulumi:"externalPort"`
	// A list of DNAT entry ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Provides the internal IP address.
	InternalIp pulumi.StringPtrInput `pulumi:"internalIp"`
	// The port or port segment on which the cloud server instance provides services to the public network.
	InternalPort pulumi.StringPtrInput `pulumi:"internalPort"`
	// The id of the NAT gateway.
	NatGatewayId pulumi.StringPtrInput `pulumi:"natGatewayId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The network protocol.
	Protocol pulumi.StringPtrInput `pulumi:"protocol"`
}

func (DnatEntriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DnatEntriesArgs)(nil)).Elem()
}

// A collection of values returned by DnatEntries.
type DnatEntriesResultOutput struct{ *pulumi.OutputState }

func (DnatEntriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DnatEntriesResult)(nil)).Elem()
}

func (o DnatEntriesResultOutput) ToDnatEntriesResultOutput() DnatEntriesResultOutput {
	return o
}

func (o DnatEntriesResultOutput) ToDnatEntriesResultOutputWithContext(ctx context.Context) DnatEntriesResultOutput {
	return o
}

// List of DNAT entries.
func (o DnatEntriesResultOutput) DnatEntries() DnatEntriesDnatEntryArrayOutput {
	return o.ApplyT(func(v DnatEntriesResult) []DnatEntriesDnatEntry { return v.DnatEntries }).(DnatEntriesDnatEntryArrayOutput)
}

// The name of the DNAT entry.
func (o DnatEntriesResultOutput) DnatEntryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.DnatEntryName }).(pulumi.StringPtrOutput)
}

// Provides the public IP address for public network access.
func (o DnatEntriesResultOutput) ExternalIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.ExternalIp }).(pulumi.StringPtrOutput)
}

// The port or port segment that receives requests from the public network. If InternalPort is passed into the port segment, ExternalPort must also be passed into the port segment.
func (o DnatEntriesResultOutput) ExternalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.ExternalPort }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o DnatEntriesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DnatEntriesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o DnatEntriesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DnatEntriesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// Provides the internal IP address.
func (o DnatEntriesResultOutput) InternalIp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.InternalIp }).(pulumi.StringPtrOutput)
}

// The port or port segment on which the cloud server instance provides services to the public network.
func (o DnatEntriesResultOutput) InternalPort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.InternalPort }).(pulumi.StringPtrOutput)
}

// The ID of the NAT gateway.
func (o DnatEntriesResultOutput) NatGatewayId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.NatGatewayId }).(pulumi.StringPtrOutput)
}

func (o DnatEntriesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The network protocol.
func (o DnatEntriesResultOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DnatEntriesResult) *string { return v.Protocol }).(pulumi.StringPtrOutput)
}

// The total count of snat entries query.
func (o DnatEntriesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v DnatEntriesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(DnatEntriesResultOutput{})
}
