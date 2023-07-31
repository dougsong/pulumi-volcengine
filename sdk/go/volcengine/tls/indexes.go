// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tls indexes
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/tls"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/tls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.Indexes(ctx, &tls.IndexesArgs{
//				Ids: []string{
//					"65d67d34-c5b4-4ec8-b3a9-175d3366****",
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
func Indexes(ctx *pulumi.Context, args *IndexesArgs, opts ...pulumi.InvokeOption) (*IndexesResult, error) {
	var rv IndexesResult
	err := ctx.Invoke("volcengine:tls/indexes:Indexes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Indexes.
type IndexesArgs struct {
	// The list of topic id of tls index.
	Ids []string `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Indexes.
type IndexesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
	// The collection of tls index query.
	TlsIndexes []IndexesTlsIndex `pulumi:"tlsIndexes"`
	// The total count of tls index query.
	TotalCount int `pulumi:"totalCount"`
}

func IndexesOutput(ctx *pulumi.Context, args IndexesOutputArgs, opts ...pulumi.InvokeOption) IndexesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (IndexesResult, error) {
			args := v.(IndexesArgs)
			r, err := Indexes(ctx, &args, opts...)
			var s IndexesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(IndexesResultOutput)
}

// A collection of arguments for invoking Indexes.
type IndexesOutputArgs struct {
	// The list of topic id of tls index.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (IndexesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexesArgs)(nil)).Elem()
}

// A collection of values returned by Indexes.
type IndexesResultOutput struct{ *pulumi.OutputState }

func (IndexesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexesResult)(nil)).Elem()
}

func (o IndexesResultOutput) ToIndexesResultOutput() IndexesResultOutput {
	return o
}

func (o IndexesResultOutput) ToIndexesResultOutputWithContext(ctx context.Context) IndexesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o IndexesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IndexesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o IndexesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IndexesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o IndexesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of tls index query.
func (o IndexesResultOutput) TlsIndexes() IndexesTlsIndexArrayOutput {
	return o.ApplyT(func(v IndexesResult) []IndexesTlsIndex { return v.TlsIndexes }).(IndexesTlsIndexArrayOutput)
}

// The total count of tls index query.
func (o IndexesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v IndexesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(IndexesResultOutput{})
}
