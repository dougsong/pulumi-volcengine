// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vpcs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vpc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.Vpcs(ctx, &vpc.VpcsArgs{
//				Ids: []string{
//					"vpc-mizl7m1kqccg5smt1bdpijuj",
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
func Vpcs(ctx *pulumi.Context, args *VpcsArgs, opts ...pulumi.InvokeOption) (*VpcsResult, error) {
	var rv VpcsResult
	err := ctx.Invoke("volcengine:vpc/vpcs:Vpcs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Vpcs.
type VpcsArgs struct {
	// A list of VPC IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of Vpc.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ProjectName of the VPC.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []VpcsTag `pulumi:"tags"`
	// The vpc name to query.
	VpcName *string `pulumi:"vpcName"`
}

// A collection of values returned by Vpcs.
type VpcsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	Ids        []string `pulumi:"ids"`
	NameRegex  *string  `pulumi:"nameRegex"`
	OutputFile *string  `pulumi:"outputFile"`
	// The ProjectName of the VPC.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []VpcsTag `pulumi:"tags"`
	// The total count of Vpc query.
	TotalCount int `pulumi:"totalCount"`
	// The name of VPC.
	VpcName *string `pulumi:"vpcName"`
	// The collection of Vpc query.
	Vpcs []VpcsVpc `pulumi:"vpcs"`
}

func VpcsOutput(ctx *pulumi.Context, args VpcsOutputArgs, opts ...pulumi.InvokeOption) VpcsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (VpcsResult, error) {
			args := v.(VpcsArgs)
			r, err := Vpcs(ctx, &args, opts...)
			var s VpcsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(VpcsResultOutput)
}

// A collection of arguments for invoking Vpcs.
type VpcsOutputArgs struct {
	// A list of VPC IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of Vpc.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ProjectName of the VPC.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// Tags.
	Tags VpcsTagArrayInput `pulumi:"tags"`
	// The vpc name to query.
	VpcName pulumi.StringPtrInput `pulumi:"vpcName"`
}

func (VpcsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcsArgs)(nil)).Elem()
}

// A collection of values returned by Vpcs.
type VpcsResultOutput struct{ *pulumi.OutputState }

func (VpcsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VpcsResult)(nil)).Elem()
}

func (o VpcsResultOutput) ToVpcsResultOutput() VpcsResultOutput {
	return o
}

func (o VpcsResultOutput) ToVpcsResultOutputWithContext(ctx context.Context) VpcsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o VpcsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VpcsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o VpcsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VpcsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o VpcsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o VpcsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The ProjectName of the VPC.
func (o VpcsResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcsResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// Tags.
func (o VpcsResultOutput) Tags() VpcsTagArrayOutput {
	return o.ApplyT(func(v VpcsResult) []VpcsTag { return v.Tags }).(VpcsTagArrayOutput)
}

// The total count of Vpc query.
func (o VpcsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v VpcsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The name of VPC.
func (o VpcsResultOutput) VpcName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VpcsResult) *string { return v.VpcName }).(pulumi.StringPtrOutput)
}

// The collection of Vpc query.
func (o VpcsResultOutput) Vpcs() VpcsVpcArrayOutput {
	return o.ApplyT(func(v VpcsResult) []VpcsVpc { return v.Vpcs }).(VpcsVpcArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VpcsResultOutput{})
}
