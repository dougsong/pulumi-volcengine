// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cr namespaces
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cr"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.Namespaces(ctx, &cr.NamespacesArgs{
//				Names: []string{
//					"namespace-*",
//				},
//				Registry: "tf-1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Namespaces(ctx *pulumi.Context, args *NamespacesArgs, opts ...pulumi.InvokeOption) (*NamespacesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv NamespacesResult
	err := ctx.Invoke("volcengine:cr/namespaces:Namespaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Namespaces.
type NamespacesArgs struct {
	// The list of instance IDs.
	Names []string `pulumi:"names"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The target cr instance name.
	Registry string `pulumi:"registry"`
}

// A collection of values returned by Namespaces.
type NamespacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id    string   `pulumi:"id"`
	Names []string `pulumi:"names"`
	// The collection of namespaces query.
	Namespaces []NamespacesNamespace `pulumi:"namespaces"`
	OutputFile *string               `pulumi:"outputFile"`
	Registry   string                `pulumi:"registry"`
	// The total count of instance query.
	TotalCount int `pulumi:"totalCount"`
}

func NamespacesOutput(ctx *pulumi.Context, args NamespacesOutputArgs, opts ...pulumi.InvokeOption) NamespacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (NamespacesResult, error) {
			args := v.(NamespacesArgs)
			r, err := Namespaces(ctx, &args, opts...)
			var s NamespacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(NamespacesResultOutput)
}

// A collection of arguments for invoking Namespaces.
type NamespacesOutputArgs struct {
	// The list of instance IDs.
	Names pulumi.StringArrayInput `pulumi:"names"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The target cr instance name.
	Registry pulumi.StringInput `pulumi:"registry"`
}

func (NamespacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespacesArgs)(nil)).Elem()
}

// A collection of values returned by Namespaces.
type NamespacesResultOutput struct{ *pulumi.OutputState }

func (NamespacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespacesResult)(nil)).Elem()
}

func (o NamespacesResultOutput) ToNamespacesResultOutput() NamespacesResultOutput {
	return o
}

func (o NamespacesResultOutput) ToNamespacesResultOutputWithContext(ctx context.Context) NamespacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o NamespacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NamespacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o NamespacesResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NamespacesResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

// The collection of namespaces query.
func (o NamespacesResultOutput) Namespaces() NamespacesNamespaceArrayOutput {
	return o.ApplyT(func(v NamespacesResult) []NamespacesNamespace { return v.Namespaces }).(NamespacesNamespaceArrayOutput)
}

func (o NamespacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NamespacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o NamespacesResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v NamespacesResult) string { return v.Registry }).(pulumi.StringOutput)
}

// The total count of instance query.
func (o NamespacesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v NamespacesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(NamespacesResultOutput{})
}
