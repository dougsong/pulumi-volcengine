// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of iam users
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/Iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Iam.Users(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func Users(ctx *pulumi.Context, args *UsersArgs, opts ...pulumi.InvokeOption) (*UsersResult, error) {
	var rv UsersResult
	err := ctx.Invoke("volcengine:Iam/users:Users", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Users.
type UsersArgs struct {
	// A Name Regex of IAM.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// A list of user names.
	UserNames []string `pulumi:"userNames"`
}

// A collection of values returned by Users.
type UsersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of user query.
	TotalCount int      `pulumi:"totalCount"`
	UserNames  []string `pulumi:"userNames"`
	// The collection of user.
	Users []UsersUser `pulumi:"users"`
}

func UsersOutput(ctx *pulumi.Context, args UsersOutputArgs, opts ...pulumi.InvokeOption) UsersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (UsersResult, error) {
			args := v.(UsersArgs)
			r, err := Users(ctx, &args, opts...)
			var s UsersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(UsersResultOutput)
}

// A collection of arguments for invoking Users.
type UsersOutputArgs struct {
	// A Name Regex of IAM.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A list of user names.
	UserNames pulumi.StringArrayInput `pulumi:"userNames"`
}

func (UsersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UsersArgs)(nil)).Elem()
}

// A collection of values returned by Users.
type UsersResultOutput struct{ *pulumi.OutputState }

func (UsersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UsersResult)(nil)).Elem()
}

func (o UsersResultOutput) ToUsersResultOutput() UsersResultOutput {
	return o
}

func (o UsersResultOutput) ToUsersResultOutputWithContext(ctx context.Context) UsersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o UsersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v UsersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o UsersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o UsersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UsersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of user query.
func (o UsersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v UsersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o UsersResultOutput) UserNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UsersResult) []string { return v.UserNames }).(pulumi.StringArrayOutput)
}

// The collection of user.
func (o UsersResultOutput) Users() UsersUserArrayOutput {
	return o.ApplyT(func(v UsersResult) []UsersUser { return v.Users }).(UsersUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(UsersResultOutput{})
}
