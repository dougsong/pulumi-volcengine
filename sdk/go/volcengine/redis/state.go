// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage redis instance state
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/redis"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redis.NewState(ctx, "foo", &redis.StateArgs{
//				Action:     pulumi.String("Restart"),
//				InstanceId: pulumi.String("redis-cnlficlt4974swtbz"),
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
// Redis State Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:redis/state:State default state:redis-mizl7m1kqccg5smt1bdpijuj
//
// ```
type State struct {
	pulumi.CustomResourceState

	// Instance Action, the value can be `Restart`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Id of Instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewState registers a new resource with the given unique name, arguments, and options.
func NewState(ctx *pulumi.Context,
	name string, args *StateArgs, opts ...pulumi.ResourceOption) (*State, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	var resource State
	err := ctx.RegisterResource("volcengine:redis/state:State", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetState gets an existing State resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetState(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StateState, opts ...pulumi.ResourceOption) (*State, error) {
	var resource State
	err := ctx.ReadResource("volcengine:redis/state:State", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering State resources.
type stateState struct {
	// Instance Action, the value can be `Restart`.
	Action *string `pulumi:"action"`
	// Id of Instance.
	InstanceId *string `pulumi:"instanceId"`
}

type StateState struct {
	// Instance Action, the value can be `Restart`.
	Action pulumi.StringPtrInput
	// Id of Instance.
	InstanceId pulumi.StringPtrInput
}

func (StateState) ElementType() reflect.Type {
	return reflect.TypeOf((*stateState)(nil)).Elem()
}

type stateArgs struct {
	// Instance Action, the value can be `Restart`.
	Action string `pulumi:"action"`
	// Id of Instance.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a State resource.
type StateArgs struct {
	// Instance Action, the value can be `Restart`.
	Action pulumi.StringInput
	// Id of Instance.
	InstanceId pulumi.StringInput
}

func (StateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*stateArgs)(nil)).Elem()
}

type StateInput interface {
	pulumi.Input

	ToStateOutput() StateOutput
	ToStateOutputWithContext(ctx context.Context) StateOutput
}

func (*State) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (i *State) ToStateOutput() StateOutput {
	return i.ToStateOutputWithContext(context.Background())
}

func (i *State) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateOutput)
}

// StateArrayInput is an input type that accepts StateArray and StateArrayOutput values.
// You can construct a concrete instance of `StateArrayInput` via:
//
//	StateArray{ StateArgs{...} }
type StateArrayInput interface {
	pulumi.Input

	ToStateArrayOutput() StateArrayOutput
	ToStateArrayOutputWithContext(context.Context) StateArrayOutput
}

type StateArray []StateInput

func (StateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*State)(nil)).Elem()
}

func (i StateArray) ToStateArrayOutput() StateArrayOutput {
	return i.ToStateArrayOutputWithContext(context.Background())
}

func (i StateArray) ToStateArrayOutputWithContext(ctx context.Context) StateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateArrayOutput)
}

// StateMapInput is an input type that accepts StateMap and StateMapOutput values.
// You can construct a concrete instance of `StateMapInput` via:
//
//	StateMap{ "key": StateArgs{...} }
type StateMapInput interface {
	pulumi.Input

	ToStateMapOutput() StateMapOutput
	ToStateMapOutputWithContext(context.Context) StateMapOutput
}

type StateMap map[string]StateInput

func (StateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*State)(nil)).Elem()
}

func (i StateMap) ToStateMapOutput() StateMapOutput {
	return i.ToStateMapOutputWithContext(context.Background())
}

func (i StateMap) ToStateMapOutputWithContext(ctx context.Context) StateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StateMapOutput)
}

type StateOutput struct{ *pulumi.OutputState }

func (StateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**State)(nil)).Elem()
}

func (o StateOutput) ToStateOutput() StateOutput {
	return o
}

func (o StateOutput) ToStateOutputWithContext(ctx context.Context) StateOutput {
	return o
}

// Instance Action, the value can be `Restart`.
func (o StateOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *State) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Id of Instance.
func (o StateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *State) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type StateArrayOutput struct{ *pulumi.OutputState }

func (StateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*State)(nil)).Elem()
}

func (o StateArrayOutput) ToStateArrayOutput() StateArrayOutput {
	return o
}

func (o StateArrayOutput) ToStateArrayOutputWithContext(ctx context.Context) StateArrayOutput {
	return o
}

func (o StateArrayOutput) Index(i pulumi.IntInput) StateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *State {
		return vs[0].([]*State)[vs[1].(int)]
	}).(StateOutput)
}

type StateMapOutput struct{ *pulumi.OutputState }

func (StateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*State)(nil)).Elem()
}

func (o StateMapOutput) ToStateMapOutput() StateMapOutput {
	return o
}

func (o StateMapOutput) ToStateMapOutputWithContext(ctx context.Context) StateMapOutput {
	return o
}

func (o StateMapOutput) MapIndex(k pulumi.StringInput) StateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *State {
		return vs[0].(map[string]*State)[vs[1].(string)]
	}).(StateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*StateInput)(nil)).Elem(), &State{})
	pulumi.RegisterInputType(reflect.TypeOf((*StateArrayInput)(nil)).Elem(), StateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*StateMapInput)(nil)).Elem(), StateMap{})
	pulumi.RegisterOutputType(StateOutput{})
	pulumi.RegisterOutputType(StateArrayOutput{})
	pulumi.RegisterOutputType(StateMapOutput{})
}
