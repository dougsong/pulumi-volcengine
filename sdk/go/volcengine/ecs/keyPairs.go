// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of ecs key pairs
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/ecs"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			fooKeyPair, err := ecs.NewKeyPair(ctx, "fooKeyPair", &ecs.KeyPairArgs{
//				KeyPairName: pulumi.String("acc-test-key-name"),
//				Description: pulumi.String("acc-test"),
//			})
//			if err != nil {
//				return err
//			}
//			_ = ecs.KeyPairsOutput(ctx, ecs.KeyPairsOutputArgs{
//				KeyPairName: fooKeyPair.KeyPairName,
//			}, nil)
//			return nil
//		})
//	}
//
// ```
func KeyPairs(ctx *pulumi.Context, args *KeyPairsArgs, opts ...pulumi.InvokeOption) (*KeyPairsResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv KeyPairsResult
	err := ctx.Invoke("volcengine:ecs/keyPairs:KeyPairs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking KeyPairs.
type KeyPairsArgs struct {
	// The finger print info.
	FingerPrint *string `pulumi:"fingerPrint"`
	// Ids of key pair.
	KeyPairIds []string `pulumi:"keyPairIds"`
	// Name of key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// Key pair names info.
	KeyPairNames []string `pulumi:"keyPairNames"`
	// A Name Regex of ECS key pairs.
	NameRegex *string `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
}

// A collection of values returned by KeyPairs.
type KeyPairsResult struct {
	// The finger print info.
	FingerPrint *string `pulumi:"fingerPrint"`
	// The provider-assigned unique ID for this managed resource.
	Id         string   `pulumi:"id"`
	KeyPairIds []string `pulumi:"keyPairIds"`
	// The name of key pair.
	KeyPairName  *string  `pulumi:"keyPairName"`
	KeyPairNames []string `pulumi:"keyPairNames"`
	// The target query key pairs info.
	KeyPairs   []KeyPairsKeyPair `pulumi:"keyPairs"`
	NameRegex  *string           `pulumi:"nameRegex"`
	OutputFile *string           `pulumi:"outputFile"`
	// The total count of ECS key pair query.
	TotalCount int `pulumi:"totalCount"`
}

func KeyPairsOutput(ctx *pulumi.Context, args KeyPairsOutputArgs, opts ...pulumi.InvokeOption) KeyPairsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (KeyPairsResult, error) {
			args := v.(KeyPairsArgs)
			r, err := KeyPairs(ctx, &args, opts...)
			var s KeyPairsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(KeyPairsResultOutput)
}

// A collection of arguments for invoking KeyPairs.
type KeyPairsOutputArgs struct {
	// The finger print info.
	FingerPrint pulumi.StringPtrInput `pulumi:"fingerPrint"`
	// Ids of key pair.
	KeyPairIds pulumi.StringArrayInput `pulumi:"keyPairIds"`
	// Name of key pair.
	KeyPairName pulumi.StringPtrInput `pulumi:"keyPairName"`
	// Key pair names info.
	KeyPairNames pulumi.StringArrayInput `pulumi:"keyPairNames"`
	// A Name Regex of ECS key pairs.
	NameRegex pulumi.StringPtrInput `pulumi:"nameRegex"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
}

func (KeyPairsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairsArgs)(nil)).Elem()
}

// A collection of values returned by KeyPairs.
type KeyPairsResultOutput struct{ *pulumi.OutputState }

func (KeyPairsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyPairsResult)(nil)).Elem()
}

func (o KeyPairsResultOutput) ToKeyPairsResultOutput() KeyPairsResultOutput {
	return o
}

func (o KeyPairsResultOutput) ToKeyPairsResultOutputWithContext(ctx context.Context) KeyPairsResultOutput {
	return o
}

// The finger print info.
func (o KeyPairsResultOutput) FingerPrint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPairsResult) *string { return v.FingerPrint }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o KeyPairsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v KeyPairsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o KeyPairsResultOutput) KeyPairIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KeyPairsResult) []string { return v.KeyPairIds }).(pulumi.StringArrayOutput)
}

// The name of key pair.
func (o KeyPairsResultOutput) KeyPairName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPairsResult) *string { return v.KeyPairName }).(pulumi.StringPtrOutput)
}

func (o KeyPairsResultOutput) KeyPairNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v KeyPairsResult) []string { return v.KeyPairNames }).(pulumi.StringArrayOutput)
}

// The target query key pairs info.
func (o KeyPairsResultOutput) KeyPairs() KeyPairsKeyPairArrayOutput {
	return o.ApplyT(func(v KeyPairsResult) []KeyPairsKeyPair { return v.KeyPairs }).(KeyPairsKeyPairArrayOutput)
}

func (o KeyPairsResultOutput) NameRegex() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPairsResult) *string { return v.NameRegex }).(pulumi.StringPtrOutput)
}

func (o KeyPairsResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyPairsResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

// The total count of ECS key pair query.
func (o KeyPairsResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v KeyPairsResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(KeyPairsResultOutput{})
}
