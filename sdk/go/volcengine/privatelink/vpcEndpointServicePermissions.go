// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package privatelink

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of privatelink vpc endpoint service permissions
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/privatelink"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/privatelink"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := privatelink.VpcEndpointServicePermissions(ctx, &privatelink.VpcEndpointServicePermissionsArgs{
//				ServiceId: "epsvc-3rel73uf2ewao5zsk2j2l58ro",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func VpcEndpointServicePermissions(ctx *pulumi.Context, args *VpcEndpointServicePermissionsArgs, opts ...pulumi.InvokeOption) (*VpcEndpointServicePermissionsResult, error) {
	var rv VpcEndpointServicePermissionsResult
	err := ctx.Invoke("volcengine:privatelink/vpcEndpointServicePermissions:VpcEndpointServicePermissions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking VpcEndpointServicePermissions.
type VpcEndpointServicePermissionsArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The Id of permit account.
	PermitAccountId *string `pulumi:"permitAccountId"`
	// The Id of service.
	ServiceId string `pulumi:"serviceId"`
}

// A collection of values returned by VpcEndpointServicePermissions.
type VpcEndpointServicePermissionsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	// The collection of query.
	Permissions []VpcEndpointServicePermissionsPermission `pulumi:"permissions"`
	// The permit account id.
	PermitAccountId *string `pulumi:"permitAccountId"`
	ServiceId       string  `pulumi:"serviceId"`
	// Returns the total amount of the data list.
	TotalCount int `pulumi:"totalCount"`
}

func VpcEndpointServicePermissionsOutput(ctx *pulumi.Context, args VpcEndpointServicePermissionsOutputArgs, opts ...pulumi.InvokeOption) VpcEndpointServicePermissionsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VpcEndpointServicePermissionsResult, error) {
			args := v.(VpcEndpointServicePermissionsArgs)
			r, err := VpcEndpointServicePermissions(ctx, &args, opts...)
			var s VpcEndpointServicePermissionsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VpcEndpointServicePermissionsResultOutput)
}

// A collection of arguments for invoking VpcEndpointServicePermissions.
type VpcEndpointServicePermissionsOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Id of permit account.
	PermitAccountId pulumi.StringPtrInput `pulumi:"permitAccountId"`
	// The Id of service.
	ServiceId pulumi.StringInput `pulumi:"serviceId"`
}

func (VpcEndpointServicePermissionsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointServicePermissionsArgs)(nil)).Elem()
}

// A collection of values returned by VpcEndpointServicePermissions.
type VpcEndpointServicePermissionsResultOutput struct{ *pulumi.OutputState }

func (VpcEndpointServicePermissionsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcEndpointServicePermissionsResult)(nil)).Elem()
}

func (o VpcEndpointServicePermissionsResultOutput) ToVpcEndpointServicePermissionsResultOutput() VpcEndpointServicePermissionsResultOutput {
	return o
}

func (o VpcEndpointServicePermissionsResultOutput) ToVpcEndpointServicePermissionsResultOutputWithContext(ctx context.Context) VpcEndpointServicePermissionsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o VpcEndpointServicePermissionsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VpcEndpointServicePermissionsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VpcEndpointServicePermissionsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointServicePermissionsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The collection of query.
func (o VpcEndpointServicePermissionsResultOutput) Permissions() VpcEndpointServicePermissionsPermissionArrayOutput {
	return o.ApplyT(func(v VpcEndpointServicePermissionsResult) []VpcEndpointServicePermissionsPermission {
		return v.Permissions
	}).(VpcEndpointServicePermissionsPermissionArrayOutput)
}

// The permit account id.
func (o VpcEndpointServicePermissionsResultOutput) PermitAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcEndpointServicePermissionsResult) *string { return v.PermitAccountId }).(pulumi.StringPtrOutput)
}

func (o VpcEndpointServicePermissionsResultOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v VpcEndpointServicePermissionsResult) string { return v.ServiceId }).(pulumi.StringOutput)
}

// Returns the total amount of the data list.
func (o VpcEndpointServicePermissionsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VpcEndpointServicePermissionsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcEndpointServicePermissionsResultOutput{})
}
