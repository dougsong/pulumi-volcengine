// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tls

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage tls topic
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
//			_, err := tls.NewTopic(ctx, "foo", &tls.TopicArgs{
//				AutoSplit:      pulumi.Bool(true),
//				Description:    pulumi.String("test"),
//				EnableTracking: pulumi.Bool(true),
//				MaxSplitShard:  pulumi.Int(10),
//				ProjectId:      pulumi.String("e020c978-4f05-40e1-9167-0113d3ef****"),
//				ShardCount:     pulumi.Int(2),
//				Tags: tls.TopicTagArray{
//					&tls.TopicTagArgs{
//						Key:   pulumi.String("k1"),
//						Value: pulumi.String("v1"),
//					},
//				},
//				TimeFormat: pulumi.String("%Y-%m-%dT%H:%M:%S,%f"),
//				TimeKey:    pulumi.String("request_time"),
//				TopicName:  pulumi.String("tf-test-topic"),
//				Ttl:        pulumi.Int(10),
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
// Tls Topic can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:tls/topic:Topic default edf051ed-3c46-49ba-9339-bea628fe****
//
// ```
type Topic struct {
	pulumi.CustomResourceState

	// Whether to enable automatic partition splitting function of the tls topic.
	AutoSplit pulumi.BoolOutput `pulumi:"autoSplit"`
	// The create time of the tls topic.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the tls project.
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether to enable WebTracking function of the tls topic.
	EnableTracking pulumi.BoolOutput `pulumi:"enableTracking"`
	// The max count of shards in the tls topic.
	MaxSplitShard pulumi.IntOutput `pulumi:"maxSplitShard"`
	// The modify time of the tls topic.
	ModifyTime pulumi.StringOutput `pulumi:"modifyTime"`
	// The project id of the tls topic.
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// The count of shards in the tls topic. Valid value range: 1-10.
	ShardCount pulumi.IntOutput `pulumi:"shardCount"`
	// Tags.
	Tags TopicTagArrayOutput `pulumi:"tags"`
	// The format of the time field.
	TimeFormat pulumi.StringOutput `pulumi:"timeFormat"`
	// The name of the time field.
	TimeKey pulumi.StringOutput `pulumi:"timeKey"`
	// The name of the tls topic.
	TopicName pulumi.StringOutput `pulumi:"topicName"`
	// The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
	Ttl pulumi.IntOutput `pulumi:"ttl"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.ShardCount == nil {
		return nil, errors.New("invalid value for required argument 'ShardCount'")
	}
	if args.TopicName == nil {
		return nil, errors.New("invalid value for required argument 'TopicName'")
	}
	if args.Ttl == nil {
		return nil, errors.New("invalid value for required argument 'Ttl'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Topic
	err := ctx.RegisterResource("volcengine:tls/topic:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("volcengine:tls/topic:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Whether to enable automatic partition splitting function of the tls topic.
	AutoSplit *bool `pulumi:"autoSplit"`
	// The create time of the tls topic.
	CreateTime *string `pulumi:"createTime"`
	// The description of the tls project.
	Description *string `pulumi:"description"`
	// Whether to enable WebTracking function of the tls topic.
	EnableTracking *bool `pulumi:"enableTracking"`
	// The max count of shards in the tls topic.
	MaxSplitShard *int `pulumi:"maxSplitShard"`
	// The modify time of the tls topic.
	ModifyTime *string `pulumi:"modifyTime"`
	// The project id of the tls topic.
	ProjectId *string `pulumi:"projectId"`
	// The count of shards in the tls topic. Valid value range: 1-10.
	ShardCount *int `pulumi:"shardCount"`
	// Tags.
	Tags []TopicTag `pulumi:"tags"`
	// The format of the time field.
	TimeFormat *string `pulumi:"timeFormat"`
	// The name of the time field.
	TimeKey *string `pulumi:"timeKey"`
	// The name of the tls topic.
	TopicName *string `pulumi:"topicName"`
	// The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
	Ttl *int `pulumi:"ttl"`
}

type TopicState struct {
	// Whether to enable automatic partition splitting function of the tls topic.
	AutoSplit pulumi.BoolPtrInput
	// The create time of the tls topic.
	CreateTime pulumi.StringPtrInput
	// The description of the tls project.
	Description pulumi.StringPtrInput
	// Whether to enable WebTracking function of the tls topic.
	EnableTracking pulumi.BoolPtrInput
	// The max count of shards in the tls topic.
	MaxSplitShard pulumi.IntPtrInput
	// The modify time of the tls topic.
	ModifyTime pulumi.StringPtrInput
	// The project id of the tls topic.
	ProjectId pulumi.StringPtrInput
	// The count of shards in the tls topic. Valid value range: 1-10.
	ShardCount pulumi.IntPtrInput
	// Tags.
	Tags TopicTagArrayInput
	// The format of the time field.
	TimeFormat pulumi.StringPtrInput
	// The name of the time field.
	TimeKey pulumi.StringPtrInput
	// The name of the tls topic.
	TopicName pulumi.StringPtrInput
	// The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
	Ttl pulumi.IntPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// Whether to enable automatic partition splitting function of the tls topic.
	AutoSplit *bool `pulumi:"autoSplit"`
	// The description of the tls project.
	Description *string `pulumi:"description"`
	// Whether to enable WebTracking function of the tls topic.
	EnableTracking *bool `pulumi:"enableTracking"`
	// The max count of shards in the tls topic.
	MaxSplitShard *int `pulumi:"maxSplitShard"`
	// The project id of the tls topic.
	ProjectId string `pulumi:"projectId"`
	// The count of shards in the tls topic. Valid value range: 1-10.
	ShardCount int `pulumi:"shardCount"`
	// Tags.
	Tags []TopicTag `pulumi:"tags"`
	// The format of the time field.
	TimeFormat *string `pulumi:"timeFormat"`
	// The name of the time field.
	TimeKey *string `pulumi:"timeKey"`
	// The name of the tls topic.
	TopicName string `pulumi:"topicName"`
	// The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
	Ttl int `pulumi:"ttl"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// Whether to enable automatic partition splitting function of the tls topic.
	AutoSplit pulumi.BoolPtrInput
	// The description of the tls project.
	Description pulumi.StringPtrInput
	// Whether to enable WebTracking function of the tls topic.
	EnableTracking pulumi.BoolPtrInput
	// The max count of shards in the tls topic.
	MaxSplitShard pulumi.IntPtrInput
	// The project id of the tls topic.
	ProjectId pulumi.StringInput
	// The count of shards in the tls topic. Valid value range: 1-10.
	ShardCount pulumi.IntInput
	// Tags.
	Tags TopicTagArrayInput
	// The format of the time field.
	TimeFormat pulumi.StringPtrInput
	// The name of the time field.
	TimeKey pulumi.StringPtrInput
	// The name of the tls topic.
	TopicName pulumi.StringInput
	// The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
	Ttl pulumi.IntInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}

type TopicInput interface {
	pulumi.Input

	ToTopicOutput() TopicOutput
	ToTopicOutputWithContext(ctx context.Context) TopicOutput
}

func (*Topic) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (i *Topic) ToTopicOutput() TopicOutput {
	return i.ToTopicOutputWithContext(context.Background())
}

func (i *Topic) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicOutput)
}

// TopicArrayInput is an input type that accepts TopicArray and TopicArrayOutput values.
// You can construct a concrete instance of `TopicArrayInput` via:
//
//	TopicArray{ TopicArgs{...} }
type TopicArrayInput interface {
	pulumi.Input

	ToTopicArrayOutput() TopicArrayOutput
	ToTopicArrayOutputWithContext(context.Context) TopicArrayOutput
}

type TopicArray []TopicInput

func (TopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (i TopicArray) ToTopicArrayOutput() TopicArrayOutput {
	return i.ToTopicArrayOutputWithContext(context.Background())
}

func (i TopicArray) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicArrayOutput)
}

// TopicMapInput is an input type that accepts TopicMap and TopicMapOutput values.
// You can construct a concrete instance of `TopicMapInput` via:
//
//	TopicMap{ "key": TopicArgs{...} }
type TopicMapInput interface {
	pulumi.Input

	ToTopicMapOutput() TopicMapOutput
	ToTopicMapOutputWithContext(context.Context) TopicMapOutput
}

type TopicMap map[string]TopicInput

func (TopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (i TopicMap) ToTopicMapOutput() TopicMapOutput {
	return i.ToTopicMapOutputWithContext(context.Background())
}

func (i TopicMap) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TopicMapOutput)
}

type TopicOutput struct{ *pulumi.OutputState }

func (TopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Topic)(nil)).Elem()
}

func (o TopicOutput) ToTopicOutput() TopicOutput {
	return o
}

func (o TopicOutput) ToTopicOutputWithContext(ctx context.Context) TopicOutput {
	return o
}

// Whether to enable automatic partition splitting function of the tls topic.
func (o TopicOutput) AutoSplit() pulumi.BoolOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolOutput { return v.AutoSplit }).(pulumi.BoolOutput)
}

// The create time of the tls topic.
func (o TopicOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the tls project.
func (o TopicOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Whether to enable WebTracking function of the tls topic.
func (o TopicOutput) EnableTracking() pulumi.BoolOutput {
	return o.ApplyT(func(v *Topic) pulumi.BoolOutput { return v.EnableTracking }).(pulumi.BoolOutput)
}

// The max count of shards in the tls topic.
func (o TopicOutput) MaxSplitShard() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.MaxSplitShard }).(pulumi.IntOutput)
}

// The modify time of the tls topic.
func (o TopicOutput) ModifyTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.ModifyTime }).(pulumi.StringOutput)
}

// The project id of the tls topic.
func (o TopicOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// The count of shards in the tls topic. Valid value range: 1-10.
func (o TopicOutput) ShardCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.ShardCount }).(pulumi.IntOutput)
}

// Tags.
func (o TopicOutput) Tags() TopicTagArrayOutput {
	return o.ApplyT(func(v *Topic) TopicTagArrayOutput { return v.Tags }).(TopicTagArrayOutput)
}

// The format of the time field.
func (o TopicOutput) TimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TimeFormat }).(pulumi.StringOutput)
}

// The name of the time field.
func (o TopicOutput) TimeKey() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TimeKey }).(pulumi.StringOutput)
}

// The name of the tls topic.
func (o TopicOutput) TopicName() pulumi.StringOutput {
	return o.ApplyT(func(v *Topic) pulumi.StringOutput { return v.TopicName }).(pulumi.StringOutput)
}

// The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
func (o TopicOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v *Topic) pulumi.IntOutput { return v.Ttl }).(pulumi.IntOutput)
}

type TopicArrayOutput struct{ *pulumi.OutputState }

func (TopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Topic)(nil)).Elem()
}

func (o TopicArrayOutput) ToTopicArrayOutput() TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) ToTopicArrayOutputWithContext(ctx context.Context) TopicArrayOutput {
	return o
}

func (o TopicArrayOutput) Index(i pulumi.IntInput) TopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].([]*Topic)[vs[1].(int)]
	}).(TopicOutput)
}

type TopicMapOutput struct{ *pulumi.OutputState }

func (TopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Topic)(nil)).Elem()
}

func (o TopicMapOutput) ToTopicMapOutput() TopicMapOutput {
	return o
}

func (o TopicMapOutput) ToTopicMapOutputWithContext(ctx context.Context) TopicMapOutput {
	return o
}

func (o TopicMapOutput) MapIndex(k pulumi.StringInput) TopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Topic {
		return vs[0].(map[string]*Topic)[vs[1].(string)]
	}).(TopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TopicInput)(nil)).Elem(), &Topic{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicArrayInput)(nil)).Elem(), TopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TopicMapInput)(nil)).Elem(), TopicMap{})
	pulumi.RegisterOutputType(TopicOutput{})
	pulumi.RegisterOutputType(TopicArrayOutput{})
	pulumi.RegisterOutputType(TopicMapOutput{})
}
