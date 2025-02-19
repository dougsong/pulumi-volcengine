// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of redis allow lists
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/redis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redis.AllowLists(ctx, &redis.AllowListsArgs{
//				RegionId: "cn-beijing",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func AllowLists(ctx *pulumi.Context, args *AllowListsArgs, opts ...pulumi.InvokeOption) (*AllowListsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv AllowListsResult
	err := ctx.Invoke("volcengine:redis/allowLists:AllowLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking AllowLists.
type AllowListsArgs struct {
	// The Id of instance.
	InstanceId *string `pulumi:"instanceId"`
	// A Name Regex of Allow List.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The Id of region.
	RegionId string `pulumi:"regionId"`
}

// A collection of values returned by AllowLists.
type AllowListsResult struct {
	// Information of list of allow list.
	AllowLists []AllowListsAllowList `pulumi:"allowLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Id of instance.
	InstanceId *string `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	RegionId   string  `pulumi:"regionId"`
	// The total count of allow list query.
	TotalCount int `pulumi:"totalCount"`
}

func AllowListsOutput(ctx *pulumi.Context, args AllowListsOutputArgs, opts ...pulumi.InvokeOption) AllowListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AllowListsResult, error) {
			args := v.(AllowListsArgs)
			r, err := AllowLists(ctx, &args, opts...)
			var s AllowListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AllowListsResultOutput)
}

// A collection of arguments for invoking AllowLists.
type AllowListsOutputArgs struct {
	// The Id of instance.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// A Name Regex of Allow List.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Id of region.
	RegionId pulumi.StringInput `pulumi:"regionId"`
}

func (AllowListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowListsArgs)(nil)).Elem()
}

// A collection of values returned by AllowLists.
type AllowListsResultOutput struct{ *pulumi.OutputState }

func (AllowListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AllowListsResult)(nil)).Elem()
}

func (o AllowListsResultOutput) ToAllowListsResultOutput() AllowListsResultOutput {
	return o
}

func (o AllowListsResultOutput) ToAllowListsResultOutputWithContext(ctx context.Context) AllowListsResultOutput {
	return o
}

// Information of list of allow list.
func (o AllowListsResultOutput) AllowLists() AllowListsAllowListArrayOutput {
	return o.ApplyT(func(v AllowListsResult) []AllowListsAllowList { return v.AllowLists }).(AllowListsAllowListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AllowListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AllowListsResult) string { return v.Id }).(pulumi.StringOutput)
}

// Id of instance.
func (o AllowListsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AllowListsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o AllowListsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AllowListsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o AllowListsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AllowListsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o AllowListsResultOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v AllowListsResult) string { return v.RegionId }).(pulumi.StringOutput)
}

// The total count of allow list query.
func (o AllowListsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AllowListsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AllowListsResultOutput{})
}
