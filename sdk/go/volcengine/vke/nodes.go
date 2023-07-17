// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vke

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of vke nodes
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/vke"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vke.Nodes(ctx, &vke.NodesArgs{
//				ClusterIds: []string{
//					"c123",
//					"c456",
//				},
//				Ids: []string{
//					"ncaa3e5mrsferqkomi190",
//				},
//				Statuses: []vke.NodesStatus{
//					vke.NodesStatus{
//						ConditionsType: pulumi.StringRef("Progressing"),
//						Phase:          pulumi.StringRef("Creating"),
//					},
//					vke.NodesStatus{
//						ConditionsType: pulumi.StringRef("Progressing123"),
//						Phase:          pulumi.StringRef("Creating123"),
//					},
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
func Nodes(ctx *pulumi.Context, args *NodesArgs, opts ...pulumi.InvokeOption) (*NodesResult, error) {
	var rv NodesResult
	err := ctx.Invoke("volcengine:vke/nodes:Nodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Nodes.
type NodesArgs struct {
	// A list of Cluster IDs.
	ClusterIds []string `pulumi:"clusterIds"`
	// The Create Client Token.
	CreateClientToken *string `pulumi:"createClientToken"`
	// A list of Node IDs.
	Ids []string `pulumi:"ids"`
	// The Name of Node.
	Name *string `pulumi:"name"`
	// A Name Regex of Node.
	NameRegex *string `pulumi:"nameRegex"`
	// The Node Pool IDs.
	NodePoolIds []string `pulumi:"nodePoolIds"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The Status of filter.
	Statuses []NodesStatus `pulumi:"statuses"`
	// The Zone IDs.
	ZoneIds []string `pulumi:"zoneIds"`
}

// A collection of values returned by Nodes.
type NodesResult struct {
	ClusterIds []string `pulumi:"clusterIds"`
	// The create client token of node.
	CreateClientToken *string `pulumi:"createClientToken"`
	// The provider-assigned unique ID for this managed resource.
	Id  string   `pulumi:"id"`
	Ids []string `pulumi:"ids"`
	// The name of Node.
	Name        *string  `pulumi:"name"`
	NameRegex   *string  `pulumi:"nameRegex"`
	NodePoolIds []string `pulumi:"nodePoolIds"`
	// The collection of Node query.
	Nodes      []NodesNode   `pulumi:"nodes"`
	OutputFile *string       `pulumi:"outputFile"`
	Statuses   []NodesStatus `pulumi:"statuses"`
	// The total count of Node query.
	TotalCount int      `pulumi:"totalCount"`
	ZoneIds    []string `pulumi:"zoneIds"`
}

func NodesOutput(ctx *pulumi.Context, args NodesOutputArgs, opts ...pulumi.InvokeOption) NodesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (NodesResult, error) {
			args := v.(NodesArgs)
			r, err := Nodes(ctx, &args, opts...)
			var s NodesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(NodesResultOutput)
}

// A collection of arguments for invoking Nodes.
type NodesOutputArgs struct {
	// A list of Cluster IDs.
	ClusterIds pulumi.StringArrayInput `pulumi:"clusterIds"`
	// The Create Client Token.
	CreateClientToken pulumi.StringPtrInput `pulumi:"createClientToken"`
	// A list of Node IDs.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// The Name of Node.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// A Name Regex of Node.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// The Node Pool IDs.
	NodePoolIds pulumi.StringArrayInput `pulumi:"nodePoolIds"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The Status of filter.
	Statuses NodesStatusArrayInput `pulumi:"statuses"`
	// The Zone IDs.
	ZoneIds pulumi.StringArrayInput `pulumi:"zoneIds"`
}

func (NodesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NodesArgs)(nil)).Elem()
}

// A collection of values returned by Nodes.
type NodesResultOutput struct{ *pulumi.OutputState }

func (NodesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NodesResult)(nil)).Elem()
}

func (o NodesResultOutput) ToNodesResultOutput() NodesResultOutput {
	return o
}

func (o NodesResultOutput) ToNodesResultOutputWithContext(ctx context.Context) NodesResultOutput {
	return o
}

func (o NodesResultOutput) ClusterIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodesResult) []string { return v.ClusterIds }).(pulumi.StringArrayOutput)
}

// The create client token of node.
func (o NodesResultOutput) CreateClientToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodesResult) *string { return v.CreateClientToken }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o NodesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NodesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o NodesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

// The name of Node.
func (o NodesResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodesResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NodesResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodesResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o NodesResultOutput) NodePoolIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodesResult) []string { return v.NodePoolIds }).(pulumi.StringArrayOutput)
}

// The collection of Node query.
func (o NodesResultOutput) Nodes() NodesNodeArrayOutput {
	return o.ApplyT(func(v NodesResult) []NodesNode { return v.Nodes }).(NodesNodeArrayOutput)
}

func (o NodesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NodesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o NodesResultOutput) Statuses() NodesStatusArrayOutput {
	return o.ApplyT(func(v NodesResult) []NodesStatus { return v.Statuses }).(NodesStatusArrayOutput)
}

// The total count of Node query.
func (o NodesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v NodesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func (o NodesResultOutput) ZoneIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NodesResult) []string { return v.ZoneIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(NodesResultOutput{})
}
