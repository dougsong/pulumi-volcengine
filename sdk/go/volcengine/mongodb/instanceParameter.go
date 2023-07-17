// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage mongodb instance parameter
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/mongodb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.NewInstanceParameter(ctx, "default", &mongodb.InstanceParameterArgs{
//				InstanceId:     pulumi.String("mongo-replica-f16e9298b121"),
//				ParameterName:  pulumi.String("connPoolMaxConnsPerHost"),
//				ParameterRole:  pulumi.String("Node"),
//				ParameterValue: pulumi.String("600"),
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
// mongodb parameter can be imported using the param:instanceId:parameterName, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:mongodb/instanceParameter:InstanceParameter default param:mongo-replica-e405f8e2****:connPoolMaxConnsPerHost
//
// ```
//
//	NoteThis resource must be imported before it can be used. Please note that instance_id and parameter_name must correspond to the ID of the above import.
type InstanceParameter struct {
	pulumi.CustomResourceState

	// The instance ID. This field cannot be modified after the resource is imported.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The name of parameter. This field cannot be modified after the resource is imported.
	ParameterName pulumi.StringOutput `pulumi:"parameterName"`
	// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
	ParameterRole pulumi.StringOutput `pulumi:"parameterRole"`
	// The value of parameter.
	ParameterValue pulumi.StringOutput `pulumi:"parameterValue"`
}

// NewInstanceParameter registers a new resource with the given unique name, arguments, and options.
func NewInstanceParameter(ctx *pulumi.Context,
	name string, args *InstanceParameterArgs, opts ...pulumi.ResourceOption) (*InstanceParameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.ParameterName == nil {
		return nil, errors.New("invalid value for required argument 'ParameterName'")
	}
	if args.ParameterRole == nil {
		return nil, errors.New("invalid value for required argument 'ParameterRole'")
	}
	if args.ParameterValue == nil {
		return nil, errors.New("invalid value for required argument 'ParameterValue'")
	}
	var resource InstanceParameter
	err := ctx.RegisterResource("volcengine:mongodb/instanceParameter:InstanceParameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceParameter gets an existing InstanceParameter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceParameterState, opts ...pulumi.ResourceOption) (*InstanceParameter, error) {
	var resource InstanceParameter
	err := ctx.ReadResource("volcengine:mongodb/instanceParameter:InstanceParameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceParameter resources.
type instanceParameterState struct {
	// The instance ID. This field cannot be modified after the resource is imported.
	InstanceId *string `pulumi:"instanceId"`
	// The name of parameter. This field cannot be modified after the resource is imported.
	ParameterName *string `pulumi:"parameterName"`
	// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
	ParameterRole *string `pulumi:"parameterRole"`
	// The value of parameter.
	ParameterValue *string `pulumi:"parameterValue"`
}

type InstanceParameterState struct {
	// The instance ID. This field cannot be modified after the resource is imported.
	InstanceId pulumi.StringPtrInput
	// The name of parameter. This field cannot be modified after the resource is imported.
	ParameterName pulumi.StringPtrInput
	// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
	ParameterRole pulumi.StringPtrInput
	// The value of parameter.
	ParameterValue pulumi.StringPtrInput
}

func (InstanceParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceParameterState)(nil)).Elem()
}

type instanceParameterArgs struct {
	// The instance ID. This field cannot be modified after the resource is imported.
	InstanceId string `pulumi:"instanceId"`
	// The name of parameter. This field cannot be modified after the resource is imported.
	ParameterName string `pulumi:"parameterName"`
	// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
	ParameterRole string `pulumi:"parameterRole"`
	// The value of parameter.
	ParameterValue string `pulumi:"parameterValue"`
}

// The set of arguments for constructing a InstanceParameter resource.
type InstanceParameterArgs struct {
	// The instance ID. This field cannot be modified after the resource is imported.
	InstanceId pulumi.StringInput
	// The name of parameter. This field cannot be modified after the resource is imported.
	ParameterName pulumi.StringInput
	// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
	ParameterRole pulumi.StringInput
	// The value of parameter.
	ParameterValue pulumi.StringInput
}

func (InstanceParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceParameterArgs)(nil)).Elem()
}

type InstanceParameterInput interface {
	pulumi.Input

	ToInstanceParameterOutput() InstanceParameterOutput
	ToInstanceParameterOutputWithContext(ctx context.Context) InstanceParameterOutput
}

func (*InstanceParameter) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceParameter)(nil)).Elem()
}

func (i *InstanceParameter) ToInstanceParameterOutput() InstanceParameterOutput {
	return i.ToInstanceParameterOutputWithContext(context.Background())
}

func (i *InstanceParameter) ToInstanceParameterOutputWithContext(ctx context.Context) InstanceParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceParameterOutput)
}

// InstanceParameterArrayInput is an input type that accepts InstanceParameterArray and InstanceParameterArrayOutput values.
// You can construct a concrete instance of `InstanceParameterArrayInput` via:
//
//	InstanceParameterArray{ InstanceParameterArgs{...} }
type InstanceParameterArrayInput interface {
	pulumi.Input

	ToInstanceParameterArrayOutput() InstanceParameterArrayOutput
	ToInstanceParameterArrayOutputWithContext(context.Context) InstanceParameterArrayOutput
}

type InstanceParameterArray []InstanceParameterInput

func (InstanceParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceParameter)(nil)).Elem()
}

func (i InstanceParameterArray) ToInstanceParameterArrayOutput() InstanceParameterArrayOutput {
	return i.ToInstanceParameterArrayOutputWithContext(context.Background())
}

func (i InstanceParameterArray) ToInstanceParameterArrayOutputWithContext(ctx context.Context) InstanceParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceParameterArrayOutput)
}

// InstanceParameterMapInput is an input type that accepts InstanceParameterMap and InstanceParameterMapOutput values.
// You can construct a concrete instance of `InstanceParameterMapInput` via:
//
//	InstanceParameterMap{ "key": InstanceParameterArgs{...} }
type InstanceParameterMapInput interface {
	pulumi.Input

	ToInstanceParameterMapOutput() InstanceParameterMapOutput
	ToInstanceParameterMapOutputWithContext(context.Context) InstanceParameterMapOutput
}

type InstanceParameterMap map[string]InstanceParameterInput

func (InstanceParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceParameter)(nil)).Elem()
}

func (i InstanceParameterMap) ToInstanceParameterMapOutput() InstanceParameterMapOutput {
	return i.ToInstanceParameterMapOutputWithContext(context.Background())
}

func (i InstanceParameterMap) ToInstanceParameterMapOutputWithContext(ctx context.Context) InstanceParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceParameterMapOutput)
}

type InstanceParameterOutput struct{ *pulumi.OutputState }

func (InstanceParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceParameter)(nil)).Elem()
}

func (o InstanceParameterOutput) ToInstanceParameterOutput() InstanceParameterOutput {
	return o
}

func (o InstanceParameterOutput) ToInstanceParameterOutputWithContext(ctx context.Context) InstanceParameterOutput {
	return o
}

// The instance ID. This field cannot be modified after the resource is imported.
func (o InstanceParameterOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceParameter) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The name of parameter. This field cannot be modified after the resource is imported.
func (o InstanceParameterOutput) ParameterName() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceParameter) pulumi.StringOutput { return v.ParameterName }).(pulumi.StringOutput)
}

// The node type to which the parameter belongs. The value range is as follows: Node, Shard, ConfigServer, Mongos.
func (o InstanceParameterOutput) ParameterRole() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceParameter) pulumi.StringOutput { return v.ParameterRole }).(pulumi.StringOutput)
}

// The value of parameter.
func (o InstanceParameterOutput) ParameterValue() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceParameter) pulumi.StringOutput { return v.ParameterValue }).(pulumi.StringOutput)
}

type InstanceParameterArrayOutput struct{ *pulumi.OutputState }

func (InstanceParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceParameter)(nil)).Elem()
}

func (o InstanceParameterArrayOutput) ToInstanceParameterArrayOutput() InstanceParameterArrayOutput {
	return o
}

func (o InstanceParameterArrayOutput) ToInstanceParameterArrayOutputWithContext(ctx context.Context) InstanceParameterArrayOutput {
	return o
}

func (o InstanceParameterArrayOutput) Index(i pulumi.IntInput) InstanceParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceParameter {
		return vs[0].([]*InstanceParameter)[vs[1].(int)]
	}).(InstanceParameterOutput)
}

type InstanceParameterMapOutput struct{ *pulumi.OutputState }

func (InstanceParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceParameter)(nil)).Elem()
}

func (o InstanceParameterMapOutput) ToInstanceParameterMapOutput() InstanceParameterMapOutput {
	return o
}

func (o InstanceParameterMapOutput) ToInstanceParameterMapOutputWithContext(ctx context.Context) InstanceParameterMapOutput {
	return o
}

func (o InstanceParameterMapOutput) MapIndex(k pulumi.StringInput) InstanceParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceParameter {
		return vs[0].(map[string]*InstanceParameter)[vs[1].(string)]
	}).(InstanceParameterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceParameterInput)(nil)).Elem(), &InstanceParameter{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceParameterArrayInput)(nil)).Elem(), InstanceParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceParameterMapInput)(nil)).Elem(), InstanceParameterMap{})
	pulumi.RegisterOutputType(InstanceParameterOutput{})
	pulumi.RegisterOutputType(InstanceParameterArrayOutput{})
	pulumi.RegisterOutputType(InstanceParameterMapOutput{})
}
