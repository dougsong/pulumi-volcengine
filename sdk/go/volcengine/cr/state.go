// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cr

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage cr registry state
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cr"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cr.NewState(ctx, "foo", &cr.StateArgs{
//				Action: pulumi.String("Start"),
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
// CR registry state can be imported using the state:registry_name, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cr/state:State default state:cr-basic
//
// ```
type State struct {
	pulumi.CustomResourceState

	// Start cr instance action,the value must be `Start`.
	Action pulumi.StringOutput `pulumi:"action"`
	// The cr instance id.
	Name pulumi.StringOutput `pulumi:"name"`
	// The status of cr instance.
	Status StateStatusOutput `pulumi:"status"`
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
	var resource State
	err := ctx.RegisterResource("volcengine:cr/state:State", name, args, &resource, opts...)
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
	err := ctx.ReadResource("volcengine:cr/state:State", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering State resources.
type stateState struct {
	// Start cr instance action,the value must be `Start`.
	Action *string `pulumi:"action"`
	// The cr instance id.
	Name *string `pulumi:"name"`
	// The status of cr instance.
	Status *StateStatus `pulumi:"status"`
}

type StateState struct {
	// Start cr instance action,the value must be `Start`.
	Action pulumi.StringPtrInput
	// The cr instance id.
	Name pulumi.StringPtrInput
	// The status of cr instance.
	Status StateStatusPtrInput
}

func (StateState) ElementType() reflect.Type {
	return reflect.TypeOf((*stateState)(nil)).Elem()
}

type stateArgs struct {
	// Start cr instance action,the value must be `Start`.
	Action string `pulumi:"action"`
	// The cr instance id.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a State resource.
type StateArgs struct {
	// Start cr instance action,the value must be `Start`.
	Action pulumi.StringInput
	// The cr instance id.
	Name pulumi.StringPtrInput
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

// Start cr instance action,the value must be `Start`.
func (o StateOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *State) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// The cr instance id.
func (o StateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *State) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The status of cr instance.
func (o StateOutput) Status() StateStatusOutput {
	return o.ApplyT(func(v *State) StateStatusOutput { return v.Status }).(StateStatusOutput)
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
