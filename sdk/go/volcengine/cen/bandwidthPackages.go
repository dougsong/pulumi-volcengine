// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cen bandwidth packages
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cen"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.BandwidthPackages(ctx, &cen.BandwidthPackagesArgs{
//				CenId: pulumi.StringRef("cen-2bzrl3srxsv0g2dx0efyoojn3"),
//				Ids: []string{
//					"cbp-2bzeew3s8p79c2dx0eeohej4x",
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
func BandwidthPackages(ctx *pulumi.Context, args *BandwidthPackagesArgs, opts ...pulumi.InvokeOption) (*BandwidthPackagesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv BandwidthPackagesResult
	err := ctx.Invoke("volcengine:cen/bandwidthPackages:BandwidthPackages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking BandwidthPackages.
type BandwidthPackagesArgs struct {
	// A list of cen bandwidth package names.
	CenBandwidthPackageNames []string `pulumi:"cenBandwidthPackageNames"`
	// A cen id.
	CenId *string `pulumi:"cenId"`
	// A list of cen bandwidth package IDs.
	Ids []string `pulumi:"ids"`
	// A local geographic region set id.
	LocalGeographicRegionSetId *string `pulumi:"localGeographicRegionSetId"`
	// A Name Regex of cen bandwidth package.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// A peer geographic region set id.
	PeerGeographicRegionSetId *string `pulumi:"peerGeographicRegionSetId"`
	// Tags.
	Tags []BandwidthPackagesTag `pulumi:"tags"`
}

// A collection of values returned by BandwidthPackages.
type BandwidthPackagesResult struct {
	// The collection of cen bandwidth package query.
	BandwidthPackages        []BandwidthPackagesBandwidthPackage `pulumi:"bandwidthPackages"`
	CenBandwidthPackageNames []string                            `pulumi:"cenBandwidthPackageNames"`
	CenId                    *string                             `pulumi:"cenId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The local geographic region set id of the cen bandwidth package.
	LocalGeographicRegionSetId *string `pulumi:"localGeographicRegionSetId"`
	NameRegex                  *string `pulumi:"nameRegex"`
	OutputFile                 *string `pulumi:"outputFile"`
	// The peer geographic region set id of the cen bandwidth package.
	PeerGeographicRegionSetId *string `pulumi:"peerGeographicRegionSetId"`
	// Tags.
	Tags []BandwidthPackagesTag `pulumi:"tags"`
	// The total count of cen bandwidth package query.
	TotalCount int `pulumi:"totalCount"`
}

func BandwidthPackagesOutput(ctx *pulumi.Context, args BandwidthPackagesOutputArgs, opts ...pulumi.InvokeOption) BandwidthPackagesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BandwidthPackagesResult, error) {
			args := v.(BandwidthPackagesArgs)
			r, err := BandwidthPackages(ctx, &args, opts...)
			var s BandwidthPackagesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BandwidthPackagesResultOutput)
}

// A collection of arguments for invoking BandwidthPackages.
type BandwidthPackagesOutputArgs struct {
	// A list of cen bandwidth package names.
	CenBandwidthPackageNames pulumi.StringArrayInput `pulumi:"cenBandwidthPackageNames"`
	// A cen id.
	CenId pulumi.StringPtrInput `pulumi:"cenId"`
	// A list of cen bandwidth package IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A local geographic region set id.
	LocalGeographicRegionSetId pulumi.StringPtrInput `pulumi:"localGeographicRegionSetId"`
	// A Name Regex of cen bandwidth package.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A peer geographic region set id.
	PeerGeographicRegionSetId pulumi.StringPtrInput `pulumi:"peerGeographicRegionSetId"`
	// Tags.
	Tags BandwidthPackagesTagArrayInput `pulumi:"tags"`
}

func (BandwidthPackagesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthPackagesArgs)(nil)).Elem()
}

// A collection of values returned by BandwidthPackages.
type BandwidthPackagesResultOutput struct{ *pulumi.OutputState }

func (BandwidthPackagesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BandwidthPackagesResult)(nil)).Elem()
}

func (o BandwidthPackagesResultOutput) ToBandwidthPackagesResultOutput() BandwidthPackagesResultOutput {
	return o
}

func (o BandwidthPackagesResultOutput) ToBandwidthPackagesResultOutputWithContext(ctx context.Context) BandwidthPackagesResultOutput {
	return o
}

// The collection of cen bandwidth package query.
func (o BandwidthPackagesResultOutput) BandwidthPackages() BandwidthPackagesBandwidthPackageArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []BandwidthPackagesBandwidthPackage { return v.BandwidthPackages }).(BandwidthPackagesBandwidthPackageArrayOutput)
}

func (o BandwidthPackagesResultOutput) CenBandwidthPackageNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []string { return v.CenBandwidthPackageNames }).(pulumi.StringArrayOutput)
}

func (o BandwidthPackagesResultOutput) CenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.CenId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o BandwidthPackagesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o BandwidthPackagesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The local geographic region set id of the cen bandwidth package.
func (o BandwidthPackagesResultOutput) LocalGeographicRegionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.LocalGeographicRegionSetId }).(pulumi.StringPtrOutput)
}

func (o BandwidthPackagesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o BandwidthPackagesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The peer geographic region set id of the cen bandwidth package.
func (o BandwidthPackagesResultOutput) PeerGeographicRegionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) *string { return v.PeerGeographicRegionSetId }).(pulumi.StringPtrOutput)
}

// Tags.
func (o BandwidthPackagesResultOutput) Tags() BandwidthPackagesTagArrayOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) []BandwidthPackagesTag { return v.Tags }).(BandwidthPackagesTagArrayOutput)
}

// The total count of cen bandwidth package query.
func (o BandwidthPackagesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v BandwidthPackagesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(BandwidthPackagesResultOutput{})
}
