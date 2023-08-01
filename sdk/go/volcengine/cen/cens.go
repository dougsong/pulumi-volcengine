// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cens
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cen"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.Cens(ctx, &cen.CensArgs{
//				Ids: []string{
//					"cen-2bzrl3srxsv0g2dx0efyoojn3",
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
func Cens(ctx *pulumi.Context, args *CensArgs, opts ...pulumi.InvokeOption) (*CensResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv CensResult
	err := ctx.Invoke("volcengine:cen/cens:Cens", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Cens.
type CensArgs struct {
	// A list of cen names.
	CenNames []string `pulumi:"cenNames"`
	// A list of cen IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of cen.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Tags.
	Tags []CensTag `pulumi:"tags"`
}

// A collection of values returned by Cens.
type CensResult struct {
	CenNames []string `pulumi:"cenNames"`
	// The collection of cen query.
	Cens []CensCen `pulumi:"cens"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// Tags.
	Tags []CensTag `pulumi:"tags"`
	// The total count of cen query.
	TotalCount int `pulumi:"totalCount"`
}

func CensOutput(ctx *pulumi.Context, args CensOutputArgs, opts ...pulumi.InvokeOption) CensResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (CensResult, error) {
			args := v.(CensArgs)
			r, err := Cens(ctx, &args, opts...)
			var s CensResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(CensResultOutput)
}

// A collection of arguments for invoking Cens.
type CensOutputArgs struct {
	// A list of cen names.
	CenNames pulumi.StringArrayInput `pulumi:"cenNames"`
	// A list of cen IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of cen.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Tags.
	Tags CensTagArrayInput `pulumi:"tags"`
}

func (CensOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CensArgs)(nil)).Elem()
}

// A collection of values returned by Cens.
type CensResultOutput struct{ *pulumi.OutputState }

func (CensResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CensResult)(nil)).Elem()
}

func (o CensResultOutput) ToCensResultOutput() CensResultOutput {
	return o
}

func (o CensResultOutput) ToCensResultOutputWithContext(ctx context.Context) CensResultOutput {
	return o
}

func (o CensResultOutput) CenNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CensResult) []string { return v.CenNames }).(pulumi.StringArrayOutput)
}

// The collection of cen query.
func (o CensResultOutput) Cens() CensCenArrayOutput {
	return o.ApplyT(func(v CensResult) []CensCen { return v.Cens }).(CensCenArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o CensResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CensResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o CensResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CensResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o CensResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CensResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o CensResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CensResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Tags.
func (o CensResultOutput) Tags() CensTagArrayOutput {
	return o.ApplyT(func(v CensResult) []CensTag { return v.Tags }).(CensTagArrayOutput)
}

// The total count of cen query.
func (o CensResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v CensResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(CensResultOutput{})
}
