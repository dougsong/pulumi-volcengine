// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bioos

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of bioos workspaces
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/bioos"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/bioos"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bioos.Workspaces(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Workspaces(ctx *pulumi.Context, args *WorkspacesArgs, opts ...pulumi.InvokeOption) (*WorkspacesResult, error) {
	var rv WorkspacesResult
	err := ctx.Invoke("volcengine:bioos/workspaces:Workspaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Workspaces.
type WorkspacesArgs struct {
	// A list of workspace ids.
	Ids []string `pulumi:"ids"`
	// Keyword to filter by workspace name or description.
	Keyword *string `pulumi:"keyword"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Sort Field (Name CreateTime).
	SortBy *string `pulumi:"sortBy"`
	// The sort order.
	SortOrder *string `pulumi:"sortOrder"`
}

// A collection of values returned by Workspaces.
type WorkspacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// A list of workspaces.
	Items      []WorkspacesItem `pulumi:"items"`
	Keyword    *string          `pulumi:"keyword"`
	OutputFile *string          `pulumi:"outputFile"`
	SortBy     *string          `pulumi:"sortBy"`
	SortOrder  *string          `pulumi:"sortOrder"`
	// The total count of Workspace query.
	TotalCount int `pulumi:"totalCount"`
}

func WorkspacesOutput(ctx *pulumi.Context, args WorkspacesOutputArgs, opts ...pulumi.InvokeOption) WorkspacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (WorkspacesResult, error) {
			args := v.(WorkspacesArgs)
			r, err := Workspaces(ctx, &args, opts...)
			var s WorkspacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(WorkspacesResultOutput)
}

// A collection of arguments for invoking Workspaces.
type WorkspacesOutputArgs struct {
	// A list of workspace ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// Keyword to filter by workspace name or description.
	Keyword pulumi.StringPtrInput `pulumi:"keyword"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Sort Field (Name CreateTime).
	SortBy pulumi.StringPtrInput `pulumi:"sortBy"`
	// The sort order.
	SortOrder pulumi.StringPtrInput `pulumi:"sortOrder"`
}

func (WorkspacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacesArgs)(nil)).Elem()
}

// A collection of values returned by Workspaces.
type WorkspacesResultOutput struct{ *pulumi.OutputState }

func (WorkspacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspacesResult)(nil)).Elem()
}

func (o WorkspacesResultOutput) ToWorkspacesResultOutput() WorkspacesResultOutput {
	return o
}

func (o WorkspacesResultOutput) ToWorkspacesResultOutputWithContext(ctx context.Context) WorkspacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o WorkspacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v WorkspacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o WorkspacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkspacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// A list of workspaces.
func (o WorkspacesResultOutput) Items() WorkspacesItemArrayOutput {
	return o.ApplyT(func(v WorkspacesResult) []WorkspacesItem { return v.Items }).(WorkspacesItemArrayOutput)
}

func (o WorkspacesResultOutput) Keyword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspacesResult) *string { return v.Keyword }).(pulumi.StringPtrOutput)
}

func (o WorkspacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o WorkspacesResultOutput) SortBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspacesResult) *string { return v.SortBy }).(pulumi.StringPtrOutput)
}

func (o WorkspacesResultOutput) SortOrder() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkspacesResult) *string { return v.SortOrder }).(pulumi.StringPtrOutput)
}

// The total count of Workspace query.
func (o WorkspacesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v WorkspacesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspacesResultOutput{})
}
