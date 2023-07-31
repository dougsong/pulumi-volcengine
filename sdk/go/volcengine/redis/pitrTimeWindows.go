// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func PitrTimeWindows(ctx *pulumi.Context, args *PitrTimeWindowsArgs, opts ...pulumi.InvokeOption) (*PitrTimeWindowsResult, error) {
	var rv PitrTimeWindowsResult
	err := ctx.Invoke("volcengine:redis/pitrTimeWindows:PitrTimeWindows", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking PitrTimeWindows.
type PitrTimeWindowsArgs struct {
	Ids        []string `pulumi:"ids"`
	OutputFile *string  `pulumi:"outputFile"`
}

// A collection of values returned by PitrTimeWindows.
type PitrTimeWindowsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string                  `pulumi:"id"`
	Ids        []string                `pulumi:"ids"`
	OutputFile *string                 `pulumi:"outputFile"`
	Periods    []PitrTimeWindowsPeriod `pulumi:"periods"`
	TotalCount int                     `pulumi:"totalCount"`
}

func PitrTimeWindowsOutput(ctx *pulumi.Context, args PitrTimeWindowsOutputArgs, opts ...pulumi.InvokeOption) PitrTimeWindowsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (PitrTimeWindowsResult, error) {
			args := v.(PitrTimeWindowsArgs)
			r, err := PitrTimeWindows(ctx, &args, opts...)
			var s PitrTimeWindowsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(PitrTimeWindowsResultOutput)
}

// A collection of arguments for invoking PitrTimeWindows.
type PitrTimeWindowsOutputArgs struct {
	Ids        pulumi.StringArrayInput `pulumi:"ids"`
	OutputFile pulumi.StringPtrInput   `pulumi:"outputFile"`
}

func (PitrTimeWindowsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PitrTimeWindowsArgs)(nil)).Elem()
}

// A collection of values returned by PitrTimeWindows.
type PitrTimeWindowsResultOutput struct{ *pulumi.OutputState }

func (PitrTimeWindowsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PitrTimeWindowsResult)(nil)).Elem()
}

func (o PitrTimeWindowsResultOutput) ToPitrTimeWindowsResultOutput() PitrTimeWindowsResultOutput {
	return o
}

func (o PitrTimeWindowsResultOutput) ToPitrTimeWindowsResultOutputWithContext(ctx context.Context) PitrTimeWindowsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o PitrTimeWindowsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PitrTimeWindowsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o PitrTimeWindowsResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PitrTimeWindowsResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o PitrTimeWindowsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PitrTimeWindowsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o PitrTimeWindowsResultOutput) Periods() PitrTimeWindowsPeriodArrayOutput {
	return o.ApplyT(func(v PitrTimeWindowsResult) []PitrTimeWindowsPeriod { return v.Periods }).(PitrTimeWindowsPeriodArrayOutput)
}

func (o PitrTimeWindowsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v PitrTimeWindowsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(PitrTimeWindowsResultOutput{})
}
