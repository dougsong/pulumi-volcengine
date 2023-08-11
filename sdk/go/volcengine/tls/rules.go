// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of tls rules
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/tls"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := tls.Rules(ctx, &tls.RulesArgs{
//				ProjectId: "cc44f8b6-0328-4622-b043-023fca735cd4",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func Rules(ctx *pulumi.Context, args *RulesArgs, opts ...pulumi.InvokeOption) (*RulesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv RulesResult
	err := ctx.Invoke("volcengine:tls/rules:Rules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking Rules.
type RulesArgs struct {
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// The project id.
	ProjectId string `pulumi:"projectId"`
	// The rule id.
	RuleId *string `pulumi:"ruleId"`
	// The rule name.
	RuleName *string `pulumi:"ruleName"`
	// The topic id.
	TopicId *string `pulumi:"topicId"`
	// The topic name.
	TopicName *string `pulumi:"topicName"`
}

// A collection of values returned by Rules.
type RulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	OutputFile *string `pulumi:"outputFile"`
	ProjectId  string  `pulumi:"projectId"`
	// The rule id.
	RuleId *string `pulumi:"ruleId"`
	// The rule name.
	RuleName *string `pulumi:"ruleName"`
	// The rules list.
	Rules []RulesRule `pulumi:"rules"`
	// The topic id.
	TopicId *string `pulumi:"topicId"`
	// The topic name.
	TopicName *string `pulumi:"topicName"`
	// The total count of query.
	TotalCount int `pulumi:"totalCount"`
}

func RulesOutput(ctx *pulumi.Context, args RulesOutputArgs, opts ...pulumi.InvokeOption) RulesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (RulesResult, error) {
			args := v.(RulesArgs)
			r, err := Rules(ctx, &args, opts...)
			var s RulesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(RulesResultOutput)
}

// A collection of arguments for invoking Rules.
type RulesOutputArgs struct {
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// The project id.
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// The rule id.
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
	// The rule name.
	RuleName pulumi.StringPtrInput `pulumi:"ruleName"`
	// The topic id.
	TopicId pulumi.StringPtrInput `pulumi:"topicId"`
	// The topic name.
	TopicName pulumi.StringPtrInput `pulumi:"topicName"`
}

func (RulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesArgs)(nil)).Elem()
}

// A collection of values returned by Rules.
type RulesResultOutput struct{ *pulumi.OutputState }

func (RulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RulesResult)(nil)).Elem()
}

func (o RulesResultOutput) ToRulesResultOutput() RulesResultOutput {
	return o
}

func (o RulesResultOutput) ToRulesResultOutputWithContext(ctx context.Context) RulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o RulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o RulesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o RulesResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v RulesResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// The rule id.
func (o RulesResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

// The rule name.
func (o RulesResultOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.RuleName }).(pulumi.StringPtrOutput)
}

// The rules list.
func (o RulesResultOutput) Rules() RulesRuleArrayOutput {
	return o.ApplyT(func(v RulesResult) []RulesRule { return v.Rules }).(RulesRuleArrayOutput)
}

// The topic id.
func (o RulesResultOutput) TopicId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.TopicId }).(pulumi.StringPtrOutput)
}

// The topic name.
func (o RulesResultOutput) TopicName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RulesResult) *string { return v.TopicName }).(pulumi.StringPtrOutput)
}

// The total count of query.
func (o RulesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v RulesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(RulesResultOutput{})
}
