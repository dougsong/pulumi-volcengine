// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of rds mysql accounts
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds_mysql.Accounts(ctx, &rds_mysql.AccountsArgs{
//				AccountName: pulumi.StringRef(""),
//				InstanceId:  "mysql-47d6bc58762b",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Accounts(ctx *pulumi.Context, args *AccountsArgs, opts ...pulumi.InvokeOption) (*AccountsResult, error) {
	var rv AccountsResult
	err := ctx.Invoke("volcengine:rds_mysql/accounts:Accounts", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Accounts.
type AccountsArgs struct {
	// The name of the database account. This field supports fuzzy query.
	AccountName *string `pulumi:"accountName"`
	// The id of the RDS instance.
	InstanceId string `pulumi:"instanceId"`
	// A Name Regex of database account.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by Accounts.
type AccountsResult struct {
	// The name of the database account.
	AccountName *string `pulumi:"accountName"`
	// The collection of RDS instance account query.
	Accounts []AccountsAccount `pulumi:"accounts"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	InstanceId string  `pulumi:"instanceId"`
	NameRegex  *string `pulumi:"nameRegex"`
	OutputFile *string `pulumi:"outputFile"`
	// The total count of database account query.
	TotalCount int `pulumi:"totalCount"`
}

func AccountsOutput(ctx *pulumi.Context, args AccountsOutputArgs, opts ...pulumi.InvokeOption) AccountsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AccountsResult, error) {
			args := v.(AccountsArgs)
			r, err := Accounts(ctx, &args, opts...)
			var s AccountsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AccountsResultOutput)
}

// A collection of arguments for invoking Accounts.
type AccountsOutputArgs struct {
	// The name of the database account. This field supports fuzzy query.
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
	// The id of the RDS instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// A Name Regex of database account.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (AccountsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountsArgs)(nil)).Elem()
}

// A collection of values returned by Accounts.
type AccountsResultOutput struct{ *pulumi.OutputState }

func (AccountsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccountsResult)(nil)).Elem()
}

func (o AccountsResultOutput) ToAccountsResultOutput() AccountsResultOutput {
	return o
}

func (o AccountsResultOutput) ToAccountsResultOutputWithContext(ctx context.Context) AccountsResultOutput {
	return o
}

// The name of the database account.
func (o AccountsResultOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountsResult) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

// The collection of RDS instance account query.
func (o AccountsResultOutput) Accounts() AccountsAccountArrayOutput {
	return o.ApplyT(func(v AccountsResult) []AccountsAccount { return v.Accounts }).(AccountsAccountArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AccountsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AccountsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AccountsResultOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v AccountsResult) string { return v.InstanceId }).(pulumi.StringOutput)
}

func (o AccountsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o AccountsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AccountsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of database account query.
func (o AccountsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AccountsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(AccountsResultOutput{})
}
