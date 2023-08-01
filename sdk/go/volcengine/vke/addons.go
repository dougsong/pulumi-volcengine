// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vke addons
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vke"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vke"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.Addons(ctx, &vke.AddonsArgs{
//				ClusterIds: []string{
//					"cccctv1vqtofp49d96ujg",
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
func Addons(ctx *pulumi.Context, args *AddonsArgs, opts ...pulumi.InvokeOption) (*AddonsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv AddonsResult
	err := ctx.Invoke("volcengine:vke/addons:Addons", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Addons.
type AddonsArgs struct {
	// The IDs of Cluster.
	ClusterIds []string `pulumi:"clusterIds"`
	// ClientToken when the addon is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The deploy model, the value is `Managed` or `Unmanaged`.
	DeployModes []string `pulumi:"deployModes"`
	// The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deployMode is `Unmanaged`.
	DeployNodeTypes []string `pulumi:"deployNodeTypes"`
	// A Name Regex of addon.
	NameRegex *string `pulumi:"nameRegex"`
	// The Names of addons.
	Names []string `pulumi:"names"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// Array of addon states to filter.
	Statuses []AddonsStatus `pulumi:"statuses"`
	// The ClientToken when the last addon update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

// A collection of values returned by Addons.
type AddonsResult struct {
	// The collection of addon query.
	Addons            []AddonsAddon `pulumi:"addons"`
	ClusterIds        []string      `pulumi:"clusterIds"`
	CreateClientToken *string       `pulumi:"createClientToken"`
	DeployModes       []string      `pulumi:"deployModes"`
	DeployNodeTypes   []string      `pulumi:"deployNodeTypes"`
	// The provider-assigned unique ID for this managed resource.
	Id         string         `pulumi:"id"`
	NameRegex  *string        `pulumi:"nameRegex"`
	Names      []string       `pulumi:"names"`
	OutputFile *string        `pulumi:"outputFile"`
	Statuses   []AddonsStatus `pulumi:"statuses"`
	// The total count of addon query.
	TotalCount        int     `pulumi:"totalCount"`
	UpdateClientToken *string `pulumi:"updateClientToken"`
}

func AddonsOutput(ctx *pulumi.Context, args AddonsOutputArgs, opts ...pulumi.InvokeOption) AddonsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (AddonsResult, error) {
			args := v.(AddonsArgs)
			r, err := Addons(ctx, &args, opts...)
			var s AddonsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(AddonsResultOutput)
}

// A collection of arguments for invoking Addons.
type AddonsOutputArgs struct {
	// The IDs of Cluster.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// ClientToken when the addon is created successfully. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// The deploy model, the value is `Managed` or `Unmanaged`.
	DeployModes pulumi.StringArrayInput `pulumi:"deployModes"`
	// The deploy node types, the value is `Node` or `VirtualNode`. Only effected when deployMode is `Unmanaged`.
	DeployNodeTypes pulumi.StringArrayInput `pulumi:"deployNodeTypes"`
	// A Name Regex of addon.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The Names of addons.
	Names pulumi.StringArrayInput `pulumi:"names"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// Array of addon states to filter.
	Statuses AddonsStatusArrayInput `pulumi:"statuses"`
	// The ClientToken when the last addon update succeeded. ClientToken is a string that guarantees the idempotency of the request. This string is passed in by the caller.
	UpdateClientToken pulumi.StringPtrInput `pulumi:"updateClientToken"`
}

func (AddonsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonsArgs)(nil)).Elem()
}

// A collection of values returned by Addons.
type AddonsResultOutput struct{ *pulumi.OutputState }

func (AddonsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddonsResult)(nil)).Elem()
}

func (o AddonsResultOutput) ToAddonsResultOutput() AddonsResultOutput {
	return o
}

func (o AddonsResultOutput) ToAddonsResultOutputWithContext(ctx context.Context) AddonsResultOutput {
	return o
}

// The collection of addon query.
func (o AddonsResultOutput) Addons() AddonsAddonArrayOutput {
	return o.ApplyT(func(v AddonsResult) []AddonsAddon { return v.Addons }).(AddonsAddonArrayOutput)
}

func (o AddonsResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

func (o AddonsResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonsResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

func (o AddonsResultOutput) DeployModes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.DeployModes }).(pulumi.StringArrayOutput)
}

func (o AddonsResultOutput) DeployNodeTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.DeployNodeTypes }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o AddonsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AddonsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o AddonsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o AddonsResultOutput) Names() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddonsResult) []string { return v.Names }).(pulumi.StringArrayOutput)
}

func (o AddonsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o AddonsResultOutput) Statuses() AddonsStatusArrayOutput {
	return o.ApplyT(func(v AddonsResult) []AddonsStatus { return v.Statuses }).(AddonsStatusArrayOutput)
}

// The total count of addon query.
func (o AddonsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v AddonsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o AddonsResultOutput) UpdateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AddonsResult) *string { return v.UpdateClientToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AddonsResultOutput{})
}
