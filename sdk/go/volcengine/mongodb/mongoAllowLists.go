// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of mongodb allow lists
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/mongodb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.MongoAllowLists(ctx, &mongodb.MongoAllowListsArgs{
//				InstanceId: pulumi.StringRef("mongo-replica-xxx"),
//				RegionId:   "cn-xxx",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func MongoAllowLists(ctx *pulumi.Context, args *MongoAllowListsArgs, opts ...pulumi.InvokeOption) (*MongoAllowListsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv MongoAllowListsResult
	err := ctx.Invoke("volcengine:mongodb/mongoAllowLists:MongoAllowLists", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking MongoAllowLists.
type MongoAllowListsArgs struct {
	// The allow list IDs to query.
	AllowListIds []string `pulumi:"allowListIds"`
	// The instance ID to query.
	InstanceId *string `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The region ID.
	RegionId string `pulumi:"regionId"`
}

// A collection of values returned by MongoAllowLists.
type MongoAllowListsResult struct {
	AllowListIds []string `pulumi:"allowListIds"`
	// The collection of mongodb allow list query.
	AllowLists []MongoAllowListsAllowList `pulumi:"allowLists"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The instance id that bound to the allow list.
	InstanceId *string `pulumi:"instanceId"`
	OutputFile *string `pulumi:"outputFile"`
	RegionId   string  `pulumi:"regionId"`
	// The total count of mongodb allow lists query.
	TotalCount int `pulumi:"totalCount"`
}

func MongoAllowListsOutput(ctx *pulumi.Context, args MongoAllowListsOutputArgs, opts ...pulumi.InvokeOption) MongoAllowListsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (MongoAllowListsResult, error) {
			args := v.(MongoAllowListsArgs)
			r, err := MongoAllowLists(ctx, &args, opts...)
			var s MongoAllowListsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(MongoAllowListsResultOutput)
}

// A collection of arguments for invoking MongoAllowLists.
type MongoAllowListsOutputArgs struct {
	// The allow list IDs to query.
	AllowListIds pulumi.StringArrayInput `pulumi:"allowListIds"`
	// The instance ID to query.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The region ID.
	RegionId pulumi.StringInput `pulumi:"regionId"`
}

func (MongoAllowListsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoAllowListsArgs)(nil)).Elem()
}

// A collection of values returned by MongoAllowLists.
type MongoAllowListsResultOutput struct{ *pulumi.OutputState }

func (MongoAllowListsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoAllowListsResult)(nil)).Elem()
}

func (o MongoAllowListsResultOutput) ToMongoAllowListsResultOutput() MongoAllowListsResultOutput {
	return o
}

func (o MongoAllowListsResultOutput) ToMongoAllowListsResultOutputWithContext(ctx context.Context) MongoAllowListsResultOutput {
	return o
}

func (o MongoAllowListsResultOutput) AllowListIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoAllowListsResult) []string { return v.AllowListIds }).(pulumi.StringArrayOutput)
}

// The collection of mongodb allow list query.
func (o MongoAllowListsResultOutput) AllowLists() MongoAllowListsAllowListArrayOutput {
	return o.ApplyT(func(v MongoAllowListsResult) []MongoAllowListsAllowList { return v.AllowLists }).(MongoAllowListsAllowListArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o MongoAllowListsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoAllowListsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The instance id that bound to the allow list.
func (o MongoAllowListsResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MongoAllowListsResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o MongoAllowListsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MongoAllowListsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o MongoAllowListsResultOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v MongoAllowListsResult) string { return v.RegionId }).(pulumi.StringOutput)
}

// The total count of mongodb allow lists query.
func (o MongoAllowListsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v MongoAllowListsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(MongoAllowListsResultOutput{})
}
