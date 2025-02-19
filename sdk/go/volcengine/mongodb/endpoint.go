// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage mongodb endpoint
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/mongodb"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := mongodb.NewEndpoint(ctx, "foo", &mongodb.EndpointArgs{
//				EipIds: pulumi.StringArray{
//					pulumi.String("eip-3rfe12dvmz8qo5zsk2h91q05p"),
//				},
//				InstanceId:  pulumi.String("mongo-replica-38cf5badeb9e"),
//				NetworkType: pulumi.String("Public"),
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
// mongodb endpoint can be imported using the instanceId:endpointId `instanceId`represents the instance that endpoint related to. `endpointId`the id of endpoint. e.g.
//
// ```sh
//
//	$ pulumi import volcengine:mongodb/endpoint:Endpoint default mongo-replica-e405f8e2****:BRhFA0pDAk0XXkxCZQ
//
// ```
type Endpoint struct {
	pulumi.CustomResourceState

	// A list of EIP IDs that need to be bound when applying for endpoint.
	EipIds pulumi.StringArrayOutput `pulumi:"eipIds"`
	// The id of endpoint.
	EndpointId pulumi.StringOutput `pulumi:"endpointId"`
	// The instance where the endpoint resides.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// A list of the Mongos node that needs to apply for the endpoint.
	MongosNodeIds pulumi.StringArrayOutput `pulumi:"mongosNodeIds"`
	// The network type of endpoint.
	NetworkType pulumi.StringPtrOutput `pulumi:"networkType"`
	// The object ID corresponding to the endpoint.
	ObjectId pulumi.StringOutput `pulumi:"objectId"`
}

// NewEndpoint registers a new resource with the given unique name, arguments, and options.
func NewEndpoint(ctx *pulumi.Context,
	name string, args *EndpointArgs, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Endpoint
	err := ctx.RegisterResource("volcengine:mongodb/endpoint:Endpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEndpoint gets an existing Endpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EndpointState, opts ...pulumi.ResourceOption) (*Endpoint, error) {
	var resource Endpoint
	err := ctx.ReadResource("volcengine:mongodb/endpoint:Endpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Endpoint resources.
type endpointState struct {
	// A list of EIP IDs that need to be bound when applying for endpoint.
	EipIds []string `pulumi:"eipIds"`
	// The id of endpoint.
	EndpointId *string `pulumi:"endpointId"`
	// The instance where the endpoint resides.
	InstanceId *string `pulumi:"instanceId"`
	// A list of the Mongos node that needs to apply for the endpoint.
	MongosNodeIds []string `pulumi:"mongosNodeIds"`
	// The network type of endpoint.
	NetworkType *string `pulumi:"networkType"`
	// The object ID corresponding to the endpoint.
	ObjectId *string `pulumi:"objectId"`
}

type EndpointState struct {
	// A list of EIP IDs that need to be bound when applying for endpoint.
	EipIds pulumi.StringArrayInput
	// The id of endpoint.
	EndpointId pulumi.StringPtrInput
	// The instance where the endpoint resides.
	InstanceId pulumi.StringPtrInput
	// A list of the Mongos node that needs to apply for the endpoint.
	MongosNodeIds pulumi.StringArrayInput
	// The network type of endpoint.
	NetworkType pulumi.StringPtrInput
	// The object ID corresponding to the endpoint.
	ObjectId pulumi.StringPtrInput
}

func (EndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointState)(nil)).Elem()
}

type endpointArgs struct {
	// A list of EIP IDs that need to be bound when applying for endpoint.
	EipIds []string `pulumi:"eipIds"`
	// The instance where the endpoint resides.
	InstanceId string `pulumi:"instanceId"`
	// A list of the Mongos node that needs to apply for the endpoint.
	MongosNodeIds []string `pulumi:"mongosNodeIds"`
	// The network type of endpoint.
	NetworkType *string `pulumi:"networkType"`
	// The object ID corresponding to the endpoint.
	ObjectId *string `pulumi:"objectId"`
}

// The set of arguments for constructing a Endpoint resource.
type EndpointArgs struct {
	// A list of EIP IDs that need to be bound when applying for endpoint.
	EipIds pulumi.StringArrayInput
	// The instance where the endpoint resides.
	InstanceId pulumi.StringInput
	// A list of the Mongos node that needs to apply for the endpoint.
	MongosNodeIds pulumi.StringArrayInput
	// The network type of endpoint.
	NetworkType pulumi.StringPtrInput
	// The object ID corresponding to the endpoint.
	ObjectId pulumi.StringPtrInput
}

func (EndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*endpointArgs)(nil)).Elem()
}

type EndpointInput interface {
	pulumi.Input

	ToEndpointOutput() EndpointOutput
	ToEndpointOutputWithContext(ctx context.Context) EndpointOutput
}

func (*Endpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (i *Endpoint) ToEndpointOutput() EndpointOutput {
	return i.ToEndpointOutputWithContext(context.Background())
}

func (i *Endpoint) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointOutput)
}

// EndpointArrayInput is an input type that accepts EndpointArray and EndpointArrayOutput values.
// You can construct a concrete instance of `EndpointArrayInput` via:
//
//	EndpointArray{ EndpointArgs{...} }
type EndpointArrayInput interface {
	pulumi.Input

	ToEndpointArrayOutput() EndpointArrayOutput
	ToEndpointArrayOutputWithContext(context.Context) EndpointArrayOutput
}

type EndpointArray []EndpointInput

func (EndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (i EndpointArray) ToEndpointArrayOutput() EndpointArrayOutput {
	return i.ToEndpointArrayOutputWithContext(context.Background())
}

func (i EndpointArray) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointArrayOutput)
}

// EndpointMapInput is an input type that accepts EndpointMap and EndpointMapOutput values.
// You can construct a concrete instance of `EndpointMapInput` via:
//
//	EndpointMap{ "key": EndpointArgs{...} }
type EndpointMapInput interface {
	pulumi.Input

	ToEndpointMapOutput() EndpointMapOutput
	ToEndpointMapOutputWithContext(context.Context) EndpointMapOutput
}

type EndpointMap map[string]EndpointInput

func (EndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (i EndpointMap) ToEndpointMapOutput() EndpointMapOutput {
	return i.ToEndpointMapOutputWithContext(context.Background())
}

func (i EndpointMap) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointMapOutput)
}

type EndpointOutput struct{ *pulumi.OutputState }

func (EndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Endpoint)(nil)).Elem()
}

func (o EndpointOutput) ToEndpointOutput() EndpointOutput {
	return o
}

func (o EndpointOutput) ToEndpointOutputWithContext(ctx context.Context) EndpointOutput {
	return o
}

// A list of EIP IDs that need to be bound when applying for endpoint.
func (o EndpointOutput) EipIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.EipIds }).(pulumi.StringArrayOutput)
}

// The id of endpoint.
func (o EndpointOutput) EndpointId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.EndpointId }).(pulumi.StringOutput)
}

// The instance where the endpoint resides.
func (o EndpointOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// A list of the Mongos node that needs to apply for the endpoint.
func (o EndpointOutput) MongosNodeIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringArrayOutput { return v.MongosNodeIds }).(pulumi.StringArrayOutput)
}

// The network type of endpoint.
func (o EndpointOutput) NetworkType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringPtrOutput { return v.NetworkType }).(pulumi.StringPtrOutput)
}

// The object ID corresponding to the endpoint.
func (o EndpointOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Endpoint) pulumi.StringOutput { return v.ObjectId }).(pulumi.StringOutput)
}

type EndpointArrayOutput struct{ *pulumi.OutputState }

func (EndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Endpoint)(nil)).Elem()
}

func (o EndpointArrayOutput) ToEndpointArrayOutput() EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) ToEndpointArrayOutputWithContext(ctx context.Context) EndpointArrayOutput {
	return o
}

func (o EndpointArrayOutput) Index(i pulumi.IntInput) EndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].([]*Endpoint)[vs[1].(int)]
	}).(EndpointOutput)
}

type EndpointMapOutput struct{ *pulumi.OutputState }

func (EndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Endpoint)(nil)).Elem()
}

func (o EndpointMapOutput) ToEndpointMapOutput() EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) ToEndpointMapOutputWithContext(ctx context.Context) EndpointMapOutput {
	return o
}

func (o EndpointMapOutput) MapIndex(k pulumi.StringInput) EndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Endpoint {
		return vs[0].(map[string]*Endpoint)[vs[1].(string)]
	}).(EndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointInput)(nil)).Elem(), &Endpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointArrayInput)(nil)).Elem(), EndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointMapInput)(nil)).Elem(), EndpointMap{})
	pulumi.RegisterOutputType(EndpointOutput{})
	pulumi.RegisterOutputType(EndpointArrayOutput{})
	pulumi.RegisterOutputType(EndpointMapOutput{})
}
