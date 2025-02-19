// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vke kubeconfigs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vke"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.Kubeconfigs(ctx, &vke.KubeconfigsArgs{
//				ClusterIds: []string{
//					"cce7hb97qtofmj1oi4udg",
//				},
//				Types: []string{
//					"Private",
//					"Public",
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
func Kubeconfigs(ctx *pulumi.Context, args *KubeconfigsArgs, opts ...pulumi.InvokeOption) (*KubeconfigsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv KubeconfigsResult
	err := ctx.Invoke("volcengine:vke/kubeconfigs:Kubeconfigs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Kubeconfigs.
type KubeconfigsArgs struct {
	// A list of Cluster IDs.
	ClusterIds []string `pulumi:"clusterIds"`
	// A list of Kubeconfig IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of Kubeconfig.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The page number of Kubeconfigs query.
	PageNumber *int `pulumi:"pageNumber"`
	// The page size of Kubeconfigs query.
	PageSize *int `pulumi:"pageSize"`
	// The type of Kubeconfigs query.
	Types []string `pulumi:"types"`
}

// A collection of values returned by Kubeconfigs.
type KubeconfigsResult struct {
	ClusterIds []string `pulumi:"clusterIds"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The collection of VkeKubeconfig query.
	Kubeconfigs []KubeconfigsKubeconfig `pulumi:"kubeconfigs"`
	NameRegex   *string                 `pulumi:"nameRegex"`
	OutputFile  *string                 `pulumi:"outputFile"`
	PageNumber  int                     `pulumi:"pageNumber"`
	PageSize    int                     `pulumi:"pageSize"`
	// The total count of Kubeconfig query.
	TotalCount int      `pulumi:"totalCount"`
	Types      []string `pulumi:"types"`
}

func KubeconfigsOutput(ctx *pulumi.Context, args KubeconfigsOutputArgs, opts ...pulumi.InvokeOption) KubeconfigsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (KubeconfigsResult, error) {
			args := v.(KubeconfigsArgs)
			r, err := Kubeconfigs(ctx, &args, opts...)
			var s KubeconfigsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(KubeconfigsResultOutput)
}

// A collection of arguments for invoking Kubeconfigs.
type KubeconfigsOutputArgs struct {
	// A list of Cluster IDs.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// A list of Kubeconfig IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of Kubeconfig.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The page number of Kubeconfigs query.
	PageNumber pulumi.IntPtrInput `pulumi:"pageNumber"`
	// The page size of Kubeconfigs query.
	PageSize pulumi.IntPtrInput `pulumi:"pageSize"`
	// The type of Kubeconfigs query.
	Types pulumi.StringArrayInput `pulumi:"types"`
}

func (KubeconfigsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeconfigsArgs)(nil)).Elem()
}

// A collection of values returned by Kubeconfigs.
type KubeconfigsResultOutput struct{ *pulumi.OutputState }

func (KubeconfigsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeconfigsResult)(nil)).Elem()
}

func (o KubeconfigsResultOutput) ToKubeconfigsResultOutput() KubeconfigsResultOutput {
	return o
}

func (o KubeconfigsResultOutput) ToKubeconfigsResultOutputWithContext(ctx context.Context) KubeconfigsResultOutput {
	return o
}

func (o KubeconfigsResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KubeconfigsResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o KubeconfigsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KubeconfigsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o KubeconfigsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KubeconfigsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The collection of VkeKubeconfig query.
func (o KubeconfigsResultOutput) Kubeconfigs() KubeconfigsKubeconfigArrayOutput {
	return o.ApplyT(func(v KubeconfigsResult) []KubeconfigsKubeconfig { return v.Kubeconfigs }).(KubeconfigsKubeconfigArrayOutput)
}

func (o KubeconfigsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeconfigsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o KubeconfigsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeconfigsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o KubeconfigsResultOutput) PageNumber() pulumi.IntOutput {
	return o.ApplyT(func(v KubeconfigsResult) int { return v.PageNumber }).(pulumi.IntOutput)
}

func (o KubeconfigsResultOutput) PageSize() pulumi.IntOutput {
	return o.ApplyT(func(v KubeconfigsResult) int { return v.PageSize }).(pulumi.IntOutput)
}

// The total count of Kubeconfig query.
func (o KubeconfigsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v KubeconfigsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o KubeconfigsResultOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KubeconfigsResult) []string { return v.Types }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(KubeconfigsResultOutput{})
}
