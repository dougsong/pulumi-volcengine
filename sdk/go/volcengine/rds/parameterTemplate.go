// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds parameter template
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := rds.NewParameterTemplate(ctx, "foo", &rds.ParameterTemplateArgs{
//				TemplateDesc: pulumi.String("created by terraform"),
//				TemplateName: pulumi.String("tf-template"),
//				TemplateParams: rds.ParameterTemplateTemplateParamArray{
//					&rds.ParameterTemplateTemplateParamArgs{
//						Name:         pulumi.String("auto_increment_increment"),
//						RunningValue: pulumi.String("2"),
//					},
//					&rds.ParameterTemplateTemplateParamArgs{
//						Name:         pulumi.String("slow_query_log"),
//						RunningValue: pulumi.String("ON"),
//					},
//					&rds.ParameterTemplateTemplateParamArgs{
//						Name:         pulumi.String("net_retry_count"),
//						RunningValue: pulumi.String("33"),
//					},
//				},
//				TemplateType:        pulumi.String("MySQL"),
//				TemplateTypeVersion: pulumi.String("MySQL_Community_5_7"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// RDS Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:rds/parameterTemplate:ParameterTemplate default mysql-sys-80bb93aa14be22d0
//
// ```
type ParameterTemplate struct {
	pulumi.CustomResourceState

	// Parameter template description.
	TemplateDesc pulumi.StringPtrOutput `pulumi:"templateDesc"`
	// Parameter template name.
	TemplateName pulumi.StringOutput `pulumi:"templateName"`
	// Template parameters. InstanceParam only needs to pass Name and RunningValue.
	TemplateParams ParameterTemplateTemplateParamArrayOutput `pulumi:"templateParams"`
	// Parameter template database type, range of values:
	// MySQL - MySQL database. (Defaults).
	TemplateType pulumi.StringPtrOutput `pulumi:"templateType"`
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7 (default)
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion pulumi.StringPtrOutput `pulumi:"templateTypeVersion"`
}

// NewParameterTemplate registers a new resource with the given unique name, arguments, and options.
func NewParameterTemplate(ctx *pulumi.Context,
	name string, args *ParameterTemplateArgs, opts ...pulumi.ResourceOption) (*ParameterTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TemplateName == nil {
		return nil, errors.New("invalid value for required argument 'TemplateName'")
	}
	if args.TemplateParams == nil {
		return nil, errors.New("invalid value for required argument 'TemplateParams'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ParameterTemplate
	err := ctx.RegisterResource("volcengine:rds/parameterTemplate:ParameterTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameterTemplate gets an existing ParameterTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameterTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterTemplateState, opts ...pulumi.ResourceOption) (*ParameterTemplate, error) {
	var resource ParameterTemplate
	err := ctx.ReadResource("volcengine:rds/parameterTemplate:ParameterTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ParameterTemplate resources.
type parameterTemplateState struct {
	// Parameter template description.
	TemplateDesc *string `pulumi:"templateDesc"`
	// Parameter template name.
	TemplateName *string `pulumi:"templateName"`
	// Template parameters. InstanceParam only needs to pass Name and RunningValue.
	TemplateParams []ParameterTemplateTemplateParam `pulumi:"templateParams"`
	// Parameter template database type, range of values:
	// MySQL - MySQL database. (Defaults).
	TemplateType *string `pulumi:"templateType"`
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7 (default)
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion *string `pulumi:"templateTypeVersion"`
}

type ParameterTemplateState struct {
	// Parameter template description.
	TemplateDesc pulumi.StringPtrInput
	// Parameter template name.
	TemplateName pulumi.StringPtrInput
	// Template parameters. InstanceParam only needs to pass Name and RunningValue.
	TemplateParams ParameterTemplateTemplateParamArrayInput
	// Parameter template database type, range of values:
	// MySQL - MySQL database. (Defaults).
	TemplateType pulumi.StringPtrInput
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7 (default)
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion pulumi.StringPtrInput
}

func (ParameterTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterTemplateState)(nil)).Elem()
}

type parameterTemplateArgs struct {
	// Parameter template description.
	TemplateDesc *string `pulumi:"templateDesc"`
	// Parameter template name.
	TemplateName string `pulumi:"templateName"`
	// Template parameters. InstanceParam only needs to pass Name and RunningValue.
	TemplateParams []ParameterTemplateTemplateParam `pulumi:"templateParams"`
	// Parameter template database type, range of values:
	// MySQL - MySQL database. (Defaults).
	TemplateType *string `pulumi:"templateType"`
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7 (default)
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion *string `pulumi:"templateTypeVersion"`
}

// The set of arguments for constructing a ParameterTemplate resource.
type ParameterTemplateArgs struct {
	// Parameter template description.
	TemplateDesc pulumi.StringPtrInput
	// Parameter template name.
	TemplateName pulumi.StringInput
	// Template parameters. InstanceParam only needs to pass Name and RunningValue.
	TemplateParams ParameterTemplateTemplateParamArrayInput
	// Parameter template database type, range of values:
	// MySQL - MySQL database. (Defaults).
	TemplateType pulumi.StringPtrInput
	// Parameter template database version, value range:
	// MySQL_Community_5_7 - MySQL 5.7 (default)
	// MySQL_8_0 - MySQL 8.0.
	TemplateTypeVersion pulumi.StringPtrInput
}

func (ParameterTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterTemplateArgs)(nil)).Elem()
}

type ParameterTemplateInput interface {
	pulumi.Input

	ToParameterTemplateOutput() ParameterTemplateOutput
	ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput
}

func (*ParameterTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterTemplate)(nil)).Elem()
}

func (i *ParameterTemplate) ToParameterTemplateOutput() ParameterTemplateOutput {
	return i.ToParameterTemplateOutputWithContext(context.Background())
}

func (i *ParameterTemplate) ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateOutput)
}

// ParameterTemplateArrayInput is an input type that accepts ParameterTemplateArray and ParameterTemplateArrayOutput values.
// You can construct a concrete instance of `ParameterTemplateArrayInput` via:
//
//	ParameterTemplateArray{ ParameterTemplateArgs{...} }
type ParameterTemplateArrayInput interface {
	pulumi.Input

	ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput
	ToParameterTemplateArrayOutputWithContext(context.Context) ParameterTemplateArrayOutput
}

type ParameterTemplateArray []ParameterTemplateInput

func (ParameterTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterTemplate)(nil)).Elem()
}

func (i ParameterTemplateArray) ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput {
	return i.ToParameterTemplateArrayOutputWithContext(context.Background())
}

func (i ParameterTemplateArray) ToParameterTemplateArrayOutputWithContext(ctx context.Context) ParameterTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateArrayOutput)
}

// ParameterTemplateMapInput is an input type that accepts ParameterTemplateMap and ParameterTemplateMapOutput values.
// You can construct a concrete instance of `ParameterTemplateMapInput` via:
//
//	ParameterTemplateMap{ "key": ParameterTemplateArgs{...} }
type ParameterTemplateMapInput interface {
	pulumi.Input

	ToParameterTemplateMapOutput() ParameterTemplateMapOutput
	ToParameterTemplateMapOutputWithContext(context.Context) ParameterTemplateMapOutput
}

type ParameterTemplateMap map[string]ParameterTemplateInput

func (ParameterTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterTemplate)(nil)).Elem()
}

func (i ParameterTemplateMap) ToParameterTemplateMapOutput() ParameterTemplateMapOutput {
	return i.ToParameterTemplateMapOutputWithContext(context.Background())
}

func (i ParameterTemplateMap) ToParameterTemplateMapOutputWithContext(ctx context.Context) ParameterTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterTemplateMapOutput)
}

type ParameterTemplateOutput struct{ *pulumi.OutputState }

func (ParameterTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateOutput) ToParameterTemplateOutput() ParameterTemplateOutput {
	return o
}

func (o ParameterTemplateOutput) ToParameterTemplateOutputWithContext(ctx context.Context) ParameterTemplateOutput {
	return o
}

// Parameter template description.
func (o ParameterTemplateOutput) TemplateDesc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringPtrOutput { return v.TemplateDesc }).(pulumi.StringPtrOutput)
}

// Parameter template name.
func (o ParameterTemplateOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringOutput { return v.TemplateName }).(pulumi.StringOutput)
}

// Template parameters. InstanceParam only needs to pass Name and RunningValue.
func (o ParameterTemplateOutput) TemplateParams() ParameterTemplateTemplateParamArrayOutput {
	return o.ApplyT(func(v *ParameterTemplate) ParameterTemplateTemplateParamArrayOutput { return v.TemplateParams }).(ParameterTemplateTemplateParamArrayOutput)
}

// Parameter template database type, range of values:
// MySQL - MySQL database. (Defaults).
func (o ParameterTemplateOutput) TemplateType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringPtrOutput { return v.TemplateType }).(pulumi.StringPtrOutput)
}

// Parameter template database version, value range:
// MySQL_Community_5_7 - MySQL 5.7 (default)
// MySQL_8_0 - MySQL 8.0.
func (o ParameterTemplateOutput) TemplateTypeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParameterTemplate) pulumi.StringPtrOutput { return v.TemplateTypeVersion }).(pulumi.StringPtrOutput)
}

type ParameterTemplateArrayOutput struct{ *pulumi.OutputState }

func (ParameterTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateArrayOutput) ToParameterTemplateArrayOutput() ParameterTemplateArrayOutput {
	return o
}

func (o ParameterTemplateArrayOutput) ToParameterTemplateArrayOutputWithContext(ctx context.Context) ParameterTemplateArrayOutput {
	return o
}

func (o ParameterTemplateArrayOutput) Index(i pulumi.IntInput) ParameterTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ParameterTemplate {
		return vs[0].([]*ParameterTemplate)[vs[1].(int)]
	}).(ParameterTemplateOutput)
}

type ParameterTemplateMapOutput struct{ *pulumi.OutputState }

func (ParameterTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ParameterTemplate)(nil)).Elem()
}

func (o ParameterTemplateMapOutput) ToParameterTemplateMapOutput() ParameterTemplateMapOutput {
	return o
}

func (o ParameterTemplateMapOutput) ToParameterTemplateMapOutputWithContext(ctx context.Context) ParameterTemplateMapOutput {
	return o
}

func (o ParameterTemplateMapOutput) MapIndex(k pulumi.StringInput) ParameterTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ParameterTemplate {
		return vs[0].(map[string]*ParameterTemplate)[vs[1].(string)]
	}).(ParameterTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateInput)(nil)).Elem(), &ParameterTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateArrayInput)(nil)).Elem(), ParameterTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTemplateMapInput)(nil)).Elem(), ParameterTemplateMap{})
	pulumi.RegisterOutputType(ParameterTemplateOutput{})
	pulumi.RegisterOutputType(ParameterTemplateArrayOutput{})
	pulumi.RegisterOutputType(ParameterTemplateMapOutput{})
}
