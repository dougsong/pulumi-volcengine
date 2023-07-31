// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tls host groups
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
//			_, err := tls.HostGroups(ctx, &tls.HostGroupsArgs{
//				HostGroupId:   pulumi.StringRef("fbea6619-7b0c-40f3-ac7e-45c63e3f676e"),
//				HostGroupName: pulumi.StringRef("cn"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func HostGroups(ctx *pulumi.Context, args *HostGroupsArgs, opts ...pulumi.InvokeOption) (*HostGroupsResult, error) {
	var rv HostGroupsResult
	err := ctx.Invoke("volcengine:tls/hostGroups:HostGroups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking HostGroups.
type HostGroupsArgs struct {
	// Whether enable auto update.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The id of host group.
	HostGroupId *string `pulumi:"hostGroupId"`
	// The name of host group.
	HostGroupName *string `pulumi:"hostGroupName"`
	// The identifier of host.
	HostIdentifier *string `pulumi:"hostIdentifier"`
	// The project name of iam.
	IamProjectName *string `pulumi:"iamProjectName"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Whether enable service logging.
	ServiceLogging *bool `pulumi:"serviceLogging"`
}

// A collection of values returned by HostGroups.
type HostGroupsResult struct {
	// Whether enable auto update.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The id of host group.
	HostGroupId *string `pulumi:"hostGroupId"`
	// The name of host group.
	HostGroupName *string `pulumi:"hostGroupName"`
	// The identifier of host.
	HostIdentifier *string `pulumi:"hostIdentifier"`
	// The project name of iam.
	IamProjectName *string `pulumi:"iamProjectName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The collection of query.
	Infos      []HostGroupsInfo `pulumi:"infos"`
	OutputFile *string          `pulumi:"outputFile"`
	// Whether enable service logging.
	ServiceLogging *bool `pulumi:"serviceLogging"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func HostGroupsOutput(ctx *pulumi.Context, args HostGroupsOutputArgs, opts ...pulumi.InvokeOption) HostGroupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (HostGroupsResult, error) {
			args := v.(HostGroupsArgs)
			r, err := HostGroups(ctx, &args, opts...)
			var s HostGroupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(HostGroupsResultOutput)
}

// A collection of arguments for invoking HostGroups.
type HostGroupsOutputArgs struct {
	// Whether enable auto update.
	AutoUpdate pulumi.BoolPtrInput `pulumi:"autoUpdate"`
	// The id of host group.
	HostGroupId pulumi.StringPtrInput `pulumi:"hostGroupId"`
	// The name of host group.
	HostGroupName pulumi.StringPtrInput `pulumi:"hostGroupName"`
	// The identifier of host.
	HostIdentifier pulumi.StringPtrInput `pulumi:"hostIdentifier"`
	// The project name of iam.
	IamProjectName pulumi.StringPtrInput `pulumi:"iamProjectName"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Whether enable service logging.
	ServiceLogging pulumi.BoolPtrInput `pulumi:"serviceLogging"`
}

func (HostGroupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostGroupsArgs)(nil)).Elem()
}

// A collection of values returned by HostGroups.
type HostGroupsResultOutput struct{ *pulumi.OutputState }

func (HostGroupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostGroupsResult)(nil)).Elem()
}

func (o HostGroupsResultOutput) ToHostGroupsResultOutput() HostGroupsResultOutput {
	return o
}

func (o HostGroupsResultOutput) ToHostGroupsResultOutputWithContext(ctx context.Context) HostGroupsResultOutput {
	return o
}

// Whether enable auto update.
func (o HostGroupsResultOutput) AutoUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *bool { return v.AutoUpdate }).(pulumi.BoolPtrOutput)
}

// The id of host group.
func (o HostGroupsResultOutput) HostGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *string { return v.HostGroupId }).(pulumi.StringPtrOutput)
}

// The name of host group.
func (o HostGroupsResultOutput) HostGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *string { return v.HostGroupName }).(pulumi.StringPtrOutput)
}

// The identifier of host.
func (o HostGroupsResultOutput) HostIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *string { return v.HostIdentifier }).(pulumi.StringPtrOutput)
}

// The project name of iam.
func (o HostGroupsResultOutput) IamProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *string { return v.IamProjectName }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o HostGroupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v HostGroupsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The collection of query.
func (o HostGroupsResultOutput) Infos() HostGroupsInfoArrayOutput {
	return o.ApplyT(func(v HostGroupsResult) []HostGroupsInfo { return v.Infos }).(HostGroupsInfoArrayOutput)
}

func (o HostGroupsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// Whether enable service logging.
func (o HostGroupsResultOutput) ServiceLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostGroupsResult) *bool { return v.ServiceLogging }).(pulumi.BoolPtrOutput)
}

// The total count of query.
func (o HostGroupsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v HostGroupsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(HostGroupsResultOutput{})
}
