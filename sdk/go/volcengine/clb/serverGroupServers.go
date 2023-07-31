// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of server group servers
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.ServerGroupServers(ctx, &clb.ServerGroupServersArgs{
//				Ids: []string{
//					"rs-273z9pv8mtfcw7fap8sp6ie8k",
//				},
//				ServerGroupId: "rsp-273z9pt9lpdds7fap8sqdvfrf",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func ServerGroupServers(ctx *pulumi.Context, args *ServerGroupServersArgs, opts ...pulumi.InvokeOption) (*ServerGroupServersResult, error) {
	var rv ServerGroupServersResult
	err := ctx.Invoke("volcengine:clb/serverGroupServers:ServerGroupServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking ServerGroupServers.
type ServerGroupServersArgs struct {
	// The list of ServerGroupServer IDs.
	Ids []string `pulumi:"ids"`
	// A Name Regex of ServerGroupServer.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The ID of the ServerGroup.
	ServerGroupId string `pulumi:"serverGroupId"`
}

// A collection of values returned by ServerGroupServers.
type ServerGroupServersResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id            string   `pulumi:"id"`
	Ids           []string `pulumi:"ids"`
	NameRegex     *string  `pulumi:"nameRegex"`
	OutputFile    *string  `pulumi:"outputFile"`
	ServerGroupId string   `pulumi:"serverGroupId"`
	// The server list of ServerGroup.
	Servers []ServerGroupServersServer `pulumi:"servers"`
	// The total count of ServerGroupServer query.
	TotalCount int `pulumi:"totalCount"`
}

func ServerGroupServersOutput(ctx *pulumi.Context, args ServerGroupServersOutputArgs, opts ...pulumi.InvokeOption) ServerGroupServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ServerGroupServersResult, error) {
			args := v.(ServerGroupServersArgs)
			r, err := ServerGroupServers(ctx, &args, opts...)
			var s ServerGroupServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ServerGroupServersResultOutput)
}

// A collection of arguments for invoking ServerGroupServers.
type ServerGroupServersOutputArgs struct {
	// The list of ServerGroupServer IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// A Name Regex of ServerGroupServer.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringInput `pulumi:"serverGroupId"`
}

func (ServerGroupServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupServersArgs)(nil)).Elem()
}

// A collection of values returned by ServerGroupServers.
type ServerGroupServersResultOutput struct{ *pulumi.OutputState }

func (ServerGroupServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerGroupServersResult)(nil)).Elem()
}

func (o ServerGroupServersResultOutput) ToServerGroupServersResultOutput() ServerGroupServersResultOutput {
	return o
}

func (o ServerGroupServersResultOutput) ToServerGroupServersResultOutputWithContext(ctx context.Context) ServerGroupServersResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o ServerGroupServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServerGroupServersResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServerGroupServersResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ServerGroupServersResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o ServerGroupServersResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupServersResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o ServerGroupServersResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServerGroupServersResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o ServerGroupServersResultOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ServerGroupServersResult) string { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The server list of ServerGroup.
func (o ServerGroupServersResultOutput) Servers() ServerGroupServersServerArrayOutput {
	return o.ApplyT(func(v ServerGroupServersResult) []ServerGroupServersServer { return v.Servers }).(ServerGroupServersServerArrayOutput)
}

// The total count of ServerGroupServer query.
func (o ServerGroupServersResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v ServerGroupServersResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(ServerGroupServersResultOutput{})
}
