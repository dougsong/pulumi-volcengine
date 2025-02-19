// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tls host group
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
//			_, err := tls.NewHostGroup(ctx, "foo", &tls.HostGroupArgs{
//				AutoUpdate:     pulumi.Bool(false),
//				HostGroupName:  pulumi.String("tfgroup"),
//				HostGroupType:  pulumi.String("Label"),
//				HostIdentifier: pulumi.String("tf-controller"),
//				ServiceLogging: pulumi.Bool(false),
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
// Tls Host Group can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tls/hostGroup:HostGroup default edf052s21s*******dc15
//
// ```
type HostGroup struct {
	pulumi.CustomResourceState

	// The abnormal heartbeat status count of host.
	AbnormalHeartbeatStatusCount pulumi.IntOutput `pulumi:"abnormalHeartbeatStatusCount"`
	// The latest version of log collector.
	AgentLatestVersion pulumi.StringOutput `pulumi:"agentLatestVersion"`
	// Whether enable auto update.
	AutoUpdate pulumi.BoolPtrOutput `pulumi:"autoUpdate"`
	// The create time of host group.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The count of host.
	HostCount pulumi.IntOutput `pulumi:"hostCount"`
	// The name of host group.
	HostGroupName pulumi.StringOutput `pulumi:"hostGroupName"`
	// The type of host group. The value can be IP or Label.
	HostGroupType pulumi.StringOutput `pulumi:"hostGroupType"`
	// The identifier of host.
	HostIdentifier pulumi.StringPtrOutput `pulumi:"hostIdentifier"`
	// The ip list of host group.
	HostIpLists pulumi.StringArrayOutput `pulumi:"hostIpLists"`
	// The project name of iam.
	IamProjectName pulumi.StringPtrOutput `pulumi:"iamProjectName"`
	// The modify time of host group.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// The normal heartbeat status count of host.
	NormalHeartbeatStatusCount pulumi.IntOutput `pulumi:"normalHeartbeatStatusCount"`
	// The rule count of host.
	RuleCount pulumi.IntOutput `pulumi:"ruleCount"`
	// Whether enable service logging.
	ServiceLogging pulumi.BoolPtrOutput `pulumi:"serviceLogging"`
	// The update end time of log collector.
	UpdateEndTime pulumi.StringOutput `pulumi:"updateEndTime"`
	// The update start time of log collector.
	UpdateStartTime pulumi.StringOutput `pulumi:"updateStartTime"`
}

// NewHostGroup registers a new resource with the given unique name, arguments, and options.
func NewHostGroup(ctx *pulumi.Context,
	name string, args *HostGroupArgs, opts ...pulumi.ResourceOption) (*HostGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HostGroupName == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupName'")
	}
	if args.HostGroupType == nil {
		return nil, errors.New("invalid value for required argument 'HostGroupType'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource HostGroup
	err := ctx.RegisterResource("volcengine:tls/hostGroup:HostGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHostGroup gets an existing HostGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHostGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HostGroupState, opts ...pulumi.ResourceOption) (*HostGroup, error) {
	var resource HostGroup
	err := ctx.ReadResource("volcengine:tls/hostGroup:HostGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HostGroup resources.
type hostGroupState struct {
	// The abnormal heartbeat status count of host.
	AbnormalHeartbeatStatusCount *int `pulumi:"abnormalHeartbeatStatusCount"`
	// The latest version of log collector.
	AgentLatestVersion *string `pulumi:"agentLatestVersion"`
	// Whether enable auto update.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The create time of host group.
	CreateTime *string `pulumi:"createTime"`
	// The count of host.
	HostCount *int `pulumi:"hostCount"`
	// The name of host group.
	HostGroupName *string `pulumi:"hostGroupName"`
	// The type of host group. The value can be IP or Label.
	HostGroupType *string `pulumi:"hostGroupType"`
	// The identifier of host.
	HostIdentifier *string `pulumi:"hostIdentifier"`
	// The ip list of host group.
	HostIpLists []string `pulumi:"hostIpLists"`
	// The project name of iam.
	IamProjectName *string `pulumi:"iamProjectName"`
	// The modify time of host group.
	ModifyTime *string `pulumi:"modifyTime"`
	// The normal heartbeat status count of host.
	NormalHeartbeatStatusCount *int `pulumi:"normalHeartbeatStatusCount"`
	// The rule count of host.
	RuleCount *int `pulumi:"ruleCount"`
	// Whether enable service logging.
	ServiceLogging *bool `pulumi:"serviceLogging"`
	// The update end time of log collector.
	UpdateEndTime *string `pulumi:"updateEndTime"`
	// The update start time of log collector.
	UpdateStartTime *string `pulumi:"updateStartTime"`
}

type HostGroupState struct {
	// The abnormal heartbeat status count of host.
	AbnormalHeartbeatStatusCount pulumi.IntPtrInput
	// The latest version of log collector.
	AgentLatestVersion pulumi.StringPtrInput
	// Whether enable auto update.
	AutoUpdate pulumi.BoolPtrInput
	// The create time of host group.
	CreateTime pulumi.StringPtrInput
	// The count of host.
	HostCount pulumi.IntPtrInput
	// The name of host group.
	HostGroupName pulumi.StringPtrInput
	// The type of host group. The value can be IP or Label.
	HostGroupType pulumi.StringPtrInput
	// The identifier of host.
	HostIdentifier pulumi.StringPtrInput
	// The ip list of host group.
	HostIpLists pulumi.StringArrayInput
	// The project name of iam.
	IamProjectName pulumi.StringPtrInput
	// The modify time of host group.
	ModifyTime pulumi.StringPtrInput
	// The normal heartbeat status count of host.
	NormalHeartbeatStatusCount pulumi.IntPtrInput
	// The rule count of host.
	RuleCount pulumi.IntPtrInput
	// Whether enable service logging.
	ServiceLogging pulumi.BoolPtrInput
	// The update end time of log collector.
	UpdateEndTime pulumi.StringPtrInput
	// The update start time of log collector.
	UpdateStartTime pulumi.StringPtrInput
}

func (HostGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupState)(nil)).Elem()
}

type hostGroupArgs struct {
	// Whether enable auto update.
	AutoUpdate *bool `pulumi:"autoUpdate"`
	// The name of host group.
	HostGroupName string `pulumi:"hostGroupName"`
	// The type of host group. The value can be IP or Label.
	HostGroupType string `pulumi:"hostGroupType"`
	// The identifier of host.
	HostIdentifier *string `pulumi:"hostIdentifier"`
	// The ip list of host group.
	HostIpLists []string `pulumi:"hostIpLists"`
	// The project name of iam.
	IamProjectName *string `pulumi:"iamProjectName"`
	// Whether enable service logging.
	ServiceLogging *bool `pulumi:"serviceLogging"`
	// The update end time of log collector.
	UpdateEndTime *string `pulumi:"updateEndTime"`
	// The update start time of log collector.
	UpdateStartTime *string `pulumi:"updateStartTime"`
}

// The set of arguments for constructing a HostGroup resource.
type HostGroupArgs struct {
	// Whether enable auto update.
	AutoUpdate pulumi.BoolPtrInput
	// The name of host group.
	HostGroupName pulumi.StringInput
	// The type of host group. The value can be IP or Label.
	HostGroupType pulumi.StringInput
	// The identifier of host.
	HostIdentifier pulumi.StringPtrInput
	// The ip list of host group.
	HostIpLists pulumi.StringArrayInput
	// The project name of iam.
	IamProjectName pulumi.StringPtrInput
	// Whether enable service logging.
	ServiceLogging pulumi.BoolPtrInput
	// The update end time of log collector.
	UpdateEndTime pulumi.StringPtrInput
	// The update start time of log collector.
	UpdateStartTime pulumi.StringPtrInput
}

func (HostGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hostGroupArgs)(nil)).Elem()
}

type HostGroupInput interface {
	pulumi.Input

	ToHostGroupOutput() HostGroupOutput
	ToHostGroupOutputWithContext(ctx context.Context) HostGroupOutput
}

func (*HostGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**HostGroup)(nil)).Elem()
}

func (i *HostGroup) ToHostGroupOutput() HostGroupOutput {
	return i.ToHostGroupOutputWithContext(context.Background())
}

func (i *HostGroup) ToHostGroupOutputWithContext(ctx context.Context) HostGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupOutput)
}

// HostGroupArrayInput is an input type that accepts HostGroupArray and HostGroupArrayOutput values.
// You can construct a concrete instance of `HostGroupArrayInput` via:
//
//	HostGroupArray{ HostGroupArgs{...} }
type HostGroupArrayInput interface {
	pulumi.Input

	ToHostGroupArrayOutput() HostGroupArrayOutput
	ToHostGroupArrayOutputWithContext(context.Context) HostGroupArrayOutput
}

type HostGroupArray []HostGroupInput

func (HostGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostGroup)(nil)).Elem()
}

func (i HostGroupArray) ToHostGroupArrayOutput() HostGroupArrayOutput {
	return i.ToHostGroupArrayOutputWithContext(context.Background())
}

func (i HostGroupArray) ToHostGroupArrayOutputWithContext(ctx context.Context) HostGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupArrayOutput)
}

// HostGroupMapInput is an input type that accepts HostGroupMap and HostGroupMapOutput values.
// You can construct a concrete instance of `HostGroupMapInput` via:
//
//	HostGroupMap{ "key": HostGroupArgs{...} }
type HostGroupMapInput interface {
	pulumi.Input

	ToHostGroupMapOutput() HostGroupMapOutput
	ToHostGroupMapOutputWithContext(context.Context) HostGroupMapOutput
}

type HostGroupMap map[string]HostGroupInput

func (HostGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostGroup)(nil)).Elem()
}

func (i HostGroupMap) ToHostGroupMapOutput() HostGroupMapOutput {
	return i.ToHostGroupMapOutputWithContext(context.Background())
}

func (i HostGroupMap) ToHostGroupMapOutputWithContext(ctx context.Context) HostGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostGroupMapOutput)
}

type HostGroupOutput struct{ *pulumi.OutputState }

func (HostGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostGroup)(nil)).Elem()
}

func (o HostGroupOutput) ToHostGroupOutput() HostGroupOutput {
	return o
}

func (o HostGroupOutput) ToHostGroupOutputWithContext(ctx context.Context) HostGroupOutput {
	return o
}

// The abnormal heartbeat status count of host.
func (o HostGroupOutput) AbnormalHeartbeatStatusCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.IntOutput { return v.AbnormalHeartbeatStatusCount }).(pulumi.IntOutput)
}

// The latest version of log collector.
func (o HostGroupOutput) AgentLatestVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.AgentLatestVersion }).(pulumi.StringOutput)
}

// Whether enable auto update.
func (o HostGroupOutput) AutoUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.BoolPtrOutput { return v.AutoUpdate }).(pulumi.BoolPtrOutput)
}

// The create time of host group.
func (o HostGroupOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The count of host.
func (o HostGroupOutput) HostCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.IntOutput { return v.HostCount }).(pulumi.IntOutput)
}

// The name of host group.
func (o HostGroupOutput) HostGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.HostGroupName }).(pulumi.StringOutput)
}

// The type of host group. The value can be IP or Label.
func (o HostGroupOutput) HostGroupType() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.HostGroupType }).(pulumi.StringOutput)
}

// The identifier of host.
func (o HostGroupOutput) HostIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringPtrOutput { return v.HostIdentifier }).(pulumi.StringPtrOutput)
}

// The ip list of host group.
func (o HostGroupOutput) HostIpLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringArrayOutput { return v.HostIpLists }).(pulumi.StringArrayOutput)
}

// The project name of iam.
func (o HostGroupOutput) IamProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringPtrOutput { return v.IamProjectName }).(pulumi.StringPtrOutput)
}

// The modify time of host group.
func (o HostGroupOutput) ModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.ModifyTime }).(pulumi.StringOutput)
}

// The normal heartbeat status count of host.
func (o HostGroupOutput) NormalHeartbeatStatusCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.IntOutput { return v.NormalHeartbeatStatusCount }).(pulumi.IntOutput)
}

// The rule count of host.
func (o HostGroupOutput) RuleCount() pulumi.IntOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.IntOutput { return v.RuleCount }).(pulumi.IntOutput)
}

// Whether enable service logging.
func (o HostGroupOutput) ServiceLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.BoolPtrOutput { return v.ServiceLogging }).(pulumi.BoolPtrOutput)
}

// The update end time of log collector.
func (o HostGroupOutput) UpdateEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.UpdateEndTime }).(pulumi.StringOutput)
}

// The update start time of log collector.
func (o HostGroupOutput) UpdateStartTime() pulumi.StringOutput {
	return o.ApplyT(func(v *HostGroup) pulumi.StringOutput { return v.UpdateStartTime }).(pulumi.StringOutput)
}

type HostGroupArrayOutput struct{ *pulumi.OutputState }

func (HostGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HostGroup)(nil)).Elem()
}

func (o HostGroupArrayOutput) ToHostGroupArrayOutput() HostGroupArrayOutput {
	return o
}

func (o HostGroupArrayOutput) ToHostGroupArrayOutputWithContext(ctx context.Context) HostGroupArrayOutput {
	return o
}

func (o HostGroupArrayOutput) Index(i pulumi.IntInput) HostGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HostGroup {
		return vs[0].([]*HostGroup)[vs[1].(int)]
	}).(HostGroupOutput)
}

type HostGroupMapOutput struct{ *pulumi.OutputState }

func (HostGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HostGroup)(nil)).Elem()
}

func (o HostGroupMapOutput) ToHostGroupMapOutput() HostGroupMapOutput {
	return o
}

func (o HostGroupMapOutput) ToHostGroupMapOutputWithContext(ctx context.Context) HostGroupMapOutput {
	return o
}

func (o HostGroupMapOutput) MapIndex(k pulumi.StringInput) HostGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HostGroup {
		return vs[0].(map[string]*HostGroup)[vs[1].(string)]
	}).(HostGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupInput)(nil)).Elem(), &HostGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupArrayInput)(nil)).Elem(), HostGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HostGroupMapInput)(nil)).Elem(), HostGroupMap{})
	pulumi.RegisterOutputType(HostGroupOutput{})
	pulumi.RegisterOutputType(HostGroupArrayOutput{})
	pulumi.RegisterOutputType(HostGroupMapOutput{})
}
