// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of cr tags
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.Tags(ctx, &cr.TagsArgs{
//				Namespace:  "test",
//				Registry:   "enterprise-1",
//				Repository: "repo",
//				Types: []string{
//					"Image",
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
func Tags(ctx *pulumi.Context, args *TagsArgs, opts ...pulumi.InvokeOption) (*TagsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv TagsResult
	err := ctx.Invoke("volcengine:cr/tags:Tags", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Tags.
type TagsArgs struct {
	// The list of instance names.
	Names []string `pulumi:"names"`
	// The CR namespace.
	Namespace string `pulumi:"namespace"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The CR instance name.
	Registry string `pulumi:"registry"`
	// The repository name.
	Repository string `pulumi:"repository"`
	// The list of OCI product tag type.
	Types []string `pulumi:"types"`
}

// A collection of values returned by Tags.
type TagsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Names      []string `pulumi:"names"`
	Namespace  string   `pulumi:"namespace"`
	OutputFile *string  `pulumi:"outputFile"`
	Registry   string   `pulumi:"registry"`
	Repository string   `pulumi:"repository"`
	// The collection of repository query.
	Tags []TagsTag `pulumi:"tags"`
	// The total count of tag query.
	TotalCount int      `pulumi:"totalCount"`
	Types      []string `pulumi:"types"`
}

func TagsOutput(ctx *pulumi.Context, args TagsOutputArgs, opts ...pulumi.InvokeOption) TagsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (TagsResult, error) {
			args := v.(TagsArgs)
			r, err := Tags(ctx, &args, opts...)
			var s TagsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(TagsResultOutput)
}

// A collection of arguments for invoking Tags.
type TagsOutputArgs struct {
	// The list of instance names.
	Names pulumi.StringArrayInput `pulumi:"names"`
	// The CR namespace.
	Namespace pulumi.StringInput `pulumi:"namespace"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The CR instance name.
	Registry pulumi.StringInput `pulumi:"registry"`
	// The repository name.
	Repository pulumi.StringInput `pulumi:"repository"`
	// The list of OCI product tag type.
	Types pulumi.StringArrayInput `pulumi:"types"`
}

func (TagsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsArgs)(nil)).Elem()
}

// A collection of values returned by Tags.
type TagsResultOutput struct{ *pulumi.OutputState }

func (TagsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagsResult)(nil)).Elem()
}

func (o TagsResultOutput) ToTagsResultOutput() TagsResultOutput {
	return o
}

func (o TagsResultOutput) ToTagsResultOutputWithContext(ctx context.Context) TagsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o TagsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TagsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o TagsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TagsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o TagsResultOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v TagsResult) string { return v.Namespace }).(pulumi.StringOutput)
}

func (o TagsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o TagsResultOutput) Registry() pulumi.StringOutput {
	return o.ApplyT(func(v TagsResult) string { return v.Registry }).(pulumi.StringOutput)
}

func (o TagsResultOutput) Repository() pulumi.StringOutput {
	return o.ApplyT(func(v TagsResult) string { return v.Repository }).(pulumi.StringOutput)
}

// The collection of repository query.
func (o TagsResultOutput) Tags() TagsTagArrayOutput {
	return o.ApplyT(func(v TagsResult) []TagsTag { return v.Tags }).(TagsTagArrayOutput)
}

// The total count of tag query.
func (o TagsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v TagsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o TagsResultOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v TagsResult) []string { return v.Types }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(TagsResultOutput{})
}
