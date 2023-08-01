// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage server group server
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/clb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.NewServerGroupServer(ctx, "foo", &clb.ServerGroupServerArgs{
//				Description:   pulumi.String("This is a server"),
//				InstanceId:    pulumi.String("i-ybp1scasbe72q1vq35wv"),
//				Port:          pulumi.Int(80),
//				ServerGroupId: pulumi.String("rsp-274xltv2sjoxs7fap8tlv3q3s"),
//				Type:          pulumi.String("ecs"),
//				Weight:        pulumi.Int(100),
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
// ServerGroupServer can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:clb/serverGroupServer:ServerGroupServer default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
//
// ```
type ServerGroupServer struct {
	pulumi.CustomResourceState

	// The description of the instance.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of ecs instance or the network card bound to ecs instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The private ip of the instance.
	Ip pulumi.StringOutput `pulumi:"ip"`
	// The port receiving request.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringOutput `pulumi:"serverGroupId"`
	// The server id of instance in ServerGroup.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// The type of instance. Optional choice contains `ecs`, `eni`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The weight of the instance, range in 0~100.
	Weight pulumi.IntPtrOutput `pulumi:"weight"`
}

// NewServerGroupServer registers a new resource with the given unique name, arguments, and options.
func NewServerGroupServer(ctx *pulumi.Context,
	name string, args *ServerGroupServerArgs, opts ...pulumi.ResourceOption) (*ServerGroupServer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.ServerGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ServerGroupId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource ServerGroupServer
	err := ctx.RegisterResource("volcengine:clb/serverGroupServer:ServerGroupServer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerGroupServer gets an existing ServerGroupServer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerGroupServer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerGroupServerState, opts ...pulumi.ResourceOption) (*ServerGroupServer, error) {
	var resource ServerGroupServer
	err := ctx.ReadResource("volcengine:clb/serverGroupServer:ServerGroupServer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerGroupServer resources.
type serverGroupServerState struct {
	// The description of the instance.
	Description *string `pulumi:"description"`
	// The ID of ecs instance or the network card bound to ecs instance.
	InstanceId *string `pulumi:"instanceId"`
	// The private ip of the instance.
	Ip *string `pulumi:"ip"`
	// The port receiving request.
	Port *int `pulumi:"port"`
	// The ID of the ServerGroup.
	ServerGroupId *string `pulumi:"serverGroupId"`
	// The server id of instance in ServerGroup.
	ServerId *string `pulumi:"serverId"`
	// The type of instance. Optional choice contains `ecs`, `eni`.
	Type *string `pulumi:"type"`
	// The weight of the instance, range in 0~100.
	Weight *int `pulumi:"weight"`
}

type ServerGroupServerState struct {
	// The description of the instance.
	Description pulumi.StringPtrInput
	// The ID of ecs instance or the network card bound to ecs instance.
	InstanceId pulumi.StringPtrInput
	// The private ip of the instance.
	Ip pulumi.StringPtrInput
	// The port receiving request.
	Port pulumi.IntPtrInput
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringPtrInput
	// The server id of instance in ServerGroup.
	ServerId pulumi.StringPtrInput
	// The type of instance. Optional choice contains `ecs`, `eni`.
	Type pulumi.StringPtrInput
	// The weight of the instance, range in 0~100.
	Weight pulumi.IntPtrInput
}

func (ServerGroupServerState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupServerState)(nil)).Elem()
}

type serverGroupServerArgs struct {
	// The description of the instance.
	Description *string `pulumi:"description"`
	// The ID of ecs instance or the network card bound to ecs instance.
	InstanceId string `pulumi:"instanceId"`
	// The private ip of the instance.
	Ip *string `pulumi:"ip"`
	// The port receiving request.
	Port int `pulumi:"port"`
	// The ID of the ServerGroup.
	ServerGroupId string `pulumi:"serverGroupId"`
	// The type of instance. Optional choice contains `ecs`, `eni`.
	Type string `pulumi:"type"`
	// The weight of the instance, range in 0~100.
	Weight *int `pulumi:"weight"`
}

// The set of arguments for constructing a ServerGroupServer resource.
type ServerGroupServerArgs struct {
	// The description of the instance.
	Description pulumi.StringPtrInput
	// The ID of ecs instance or the network card bound to ecs instance.
	InstanceId pulumi.StringInput
	// The private ip of the instance.
	Ip pulumi.StringPtrInput
	// The port receiving request.
	Port pulumi.IntInput
	// The ID of the ServerGroup.
	ServerGroupId pulumi.StringInput
	// The type of instance. Optional choice contains `ecs`, `eni`.
	Type pulumi.StringInput
	// The weight of the instance, range in 0~100.
	Weight pulumi.IntPtrInput
}

func (ServerGroupServerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverGroupServerArgs)(nil)).Elem()
}

type ServerGroupServerInput interface {
	pulumi.Input

	ToServerGroupServerOutput() ServerGroupServerOutput
	ToServerGroupServerOutputWithContext(ctx context.Context) ServerGroupServerOutput
}

func (*ServerGroupServer) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupServer)(nil)).Elem()
}

func (i *ServerGroupServer) ToServerGroupServerOutput() ServerGroupServerOutput {
	return i.ToServerGroupServerOutputWithContext(context.Background())
}

func (i *ServerGroupServer) ToServerGroupServerOutputWithContext(ctx context.Context) ServerGroupServerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupServerOutput)
}

// ServerGroupServerArrayInput is an input type that accepts ServerGroupServerArray and ServerGroupServerArrayOutput values.
// You can construct a concrete instance of `ServerGroupServerArrayInput` via:
//
//	ServerGroupServerArray{ ServerGroupServerArgs{...} }
type ServerGroupServerArrayInput interface {
	pulumi.Input

	ToServerGroupServerArrayOutput() ServerGroupServerArrayOutput
	ToServerGroupServerArrayOutputWithContext(context.Context) ServerGroupServerArrayOutput
}

type ServerGroupServerArray []ServerGroupServerInput

func (ServerGroupServerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroupServer)(nil)).Elem()
}

func (i ServerGroupServerArray) ToServerGroupServerArrayOutput() ServerGroupServerArrayOutput {
	return i.ToServerGroupServerArrayOutputWithContext(context.Background())
}

func (i ServerGroupServerArray) ToServerGroupServerArrayOutputWithContext(ctx context.Context) ServerGroupServerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupServerArrayOutput)
}

// ServerGroupServerMapInput is an input type that accepts ServerGroupServerMap and ServerGroupServerMapOutput values.
// You can construct a concrete instance of `ServerGroupServerMapInput` via:
//
//	ServerGroupServerMap{ "key": ServerGroupServerArgs{...} }
type ServerGroupServerMapInput interface {
	pulumi.Input

	ToServerGroupServerMapOutput() ServerGroupServerMapOutput
	ToServerGroupServerMapOutputWithContext(context.Context) ServerGroupServerMapOutput
}

type ServerGroupServerMap map[string]ServerGroupServerInput

func (ServerGroupServerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroupServer)(nil)).Elem()
}

func (i ServerGroupServerMap) ToServerGroupServerMapOutput() ServerGroupServerMapOutput {
	return i.ToServerGroupServerMapOutputWithContext(context.Background())
}

func (i ServerGroupServerMap) ToServerGroupServerMapOutputWithContext(ctx context.Context) ServerGroupServerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerGroupServerMapOutput)
}

type ServerGroupServerOutput struct{ *pulumi.OutputState }

func (ServerGroupServerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerGroupServer)(nil)).Elem()
}

func (o ServerGroupServerOutput) ToServerGroupServerOutput() ServerGroupServerOutput {
	return o
}

func (o ServerGroupServerOutput) ToServerGroupServerOutputWithContext(ctx context.Context) ServerGroupServerOutput {
	return o
}

// The description of the instance.
func (o ServerGroupServerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of ecs instance or the network card bound to ecs instance.
func (o ServerGroupServerOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The private ip of the instance.
func (o ServerGroupServerOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

// The port receiving request.
func (o ServerGroupServerOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The ID of the ServerGroup.
func (o ServerGroupServerOutput) ServerGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.StringOutput { return v.ServerGroupId }).(pulumi.StringOutput)
}

// The server id of instance in ServerGroup.
func (o ServerGroupServerOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// The type of instance. Optional choice contains `ecs`, `eni`.
func (o ServerGroupServerOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The weight of the instance, range in 0~100.
func (o ServerGroupServerOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerGroupServer) pulumi.IntPtrOutput { return v.Weight }).(pulumi.IntPtrOutput)
}

type ServerGroupServerArrayOutput struct{ *pulumi.OutputState }

func (ServerGroupServerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerGroupServer)(nil)).Elem()
}

func (o ServerGroupServerArrayOutput) ToServerGroupServerArrayOutput() ServerGroupServerArrayOutput {
	return o
}

func (o ServerGroupServerArrayOutput) ToServerGroupServerArrayOutputWithContext(ctx context.Context) ServerGroupServerArrayOutput {
	return o
}

func (o ServerGroupServerArrayOutput) Index(i pulumi.IntInput) ServerGroupServerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerGroupServer {
		return vs[0].([]*ServerGroupServer)[vs[1].(int)]
	}).(ServerGroupServerOutput)
}

type ServerGroupServerMapOutput struct{ *pulumi.OutputState }

func (ServerGroupServerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerGroupServer)(nil)).Elem()
}

func (o ServerGroupServerMapOutput) ToServerGroupServerMapOutput() ServerGroupServerMapOutput {
	return o
}

func (o ServerGroupServerMapOutput) ToServerGroupServerMapOutputWithContext(ctx context.Context) ServerGroupServerMapOutput {
	return o
}

func (o ServerGroupServerMapOutput) MapIndex(k pulumi.StringInput) ServerGroupServerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerGroupServer {
		return vs[0].(map[string]*ServerGroupServer)[vs[1].(string)]
	}).(ServerGroupServerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupServerInput)(nil)).Elem(), &ServerGroupServer{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupServerArrayInput)(nil)).Elem(), ServerGroupServerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerGroupServerMapInput)(nil)).Elem(), ServerGroupServerMap{})
	pulumi.RegisterOutputType(ServerGroupServerOutput{})
	pulumi.RegisterOutputType(ServerGroupServerArrayOutput{})
	pulumi.RegisterOutputType(ServerGroupServerMapOutput{})
}
