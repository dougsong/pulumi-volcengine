// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package mongodb

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage mongodb allow list associate
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
//			_, err := mongodb.NewMongoAllowListAssociate(ctx, "foo", &mongodb.MongoAllowListAssociateArgs{
//				AllowListId: pulumi.String("acl-9e307ce4efe843fb9ffd8cb6a6cb225f"),
//				InstanceId:  pulumi.String("mongo-replica-f16e9298b121"),
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
// mongodb allow list associate can be imported using the instanceId:allowListId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
//
// ```
type MongoAllowListAssociate struct {
	pulumi.CustomResourceState

	// Id of allow list to associate.
	AllowListId pulumi.StringOutput `pulumi:"allowListId"`
	// Id of instance to associate.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
}

// NewMongoAllowListAssociate registers a new resource with the given unique name, arguments, and options.
func NewMongoAllowListAssociate(ctx *pulumi.Context,
	name string, args *MongoAllowListAssociateArgs, opts ...pulumi.ResourceOption) (*MongoAllowListAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowListId == nil {
		return nil, errors.New("invalid value for required argument 'AllowListId'")
	}
	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource MongoAllowListAssociate
	err := ctx.RegisterResource("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoAllowListAssociate gets an existing MongoAllowListAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoAllowListAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoAllowListAssociateState, opts ...pulumi.ResourceOption) (*MongoAllowListAssociate, error) {
	var resource MongoAllowListAssociate
	err := ctx.ReadResource("volcengine:mongodb/mongoAllowListAssociate:MongoAllowListAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoAllowListAssociate resources.
type mongoAllowListAssociateState struct {
	// Id of allow list to associate.
	AllowListId *string `pulumi:"allowListId"`
	// Id of instance to associate.
	InstanceId *string `pulumi:"instanceId"`
}

type MongoAllowListAssociateState struct {
	// Id of allow list to associate.
	AllowListId pulumi.StringPtrInput
	// Id of instance to associate.
	InstanceId pulumi.StringPtrInput
}

func (MongoAllowListAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoAllowListAssociateState)(nil)).Elem()
}

type mongoAllowListAssociateArgs struct {
	// Id of allow list to associate.
	AllowListId string `pulumi:"allowListId"`
	// Id of instance to associate.
	InstanceId string `pulumi:"instanceId"`
}

// The set of arguments for constructing a MongoAllowListAssociate resource.
type MongoAllowListAssociateArgs struct {
	// Id of allow list to associate.
	AllowListId pulumi.StringInput
	// Id of instance to associate.
	InstanceId pulumi.StringInput
}

func (MongoAllowListAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoAllowListAssociateArgs)(nil)).Elem()
}

type MongoAllowListAssociateInput interface {
	pulumi.Input

	ToMongoAllowListAssociateOutput() MongoAllowListAssociateOutput
	ToMongoAllowListAssociateOutputWithContext(ctx context.Context) MongoAllowListAssociateOutput
}

func (*MongoAllowListAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoAllowListAssociate)(nil)).Elem()
}

func (i *MongoAllowListAssociate) ToMongoAllowListAssociateOutput() MongoAllowListAssociateOutput {
	return i.ToMongoAllowListAssociateOutputWithContext(context.Background())
}

func (i *MongoAllowListAssociate) ToMongoAllowListAssociateOutputWithContext(ctx context.Context) MongoAllowListAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoAllowListAssociateOutput)
}

// MongoAllowListAssociateArrayInput is an input type that accepts MongoAllowListAssociateArray and MongoAllowListAssociateArrayOutput values.
// You can construct a concrete instance of `MongoAllowListAssociateArrayInput` via:
//
//	MongoAllowListAssociateArray{ MongoAllowListAssociateArgs{...} }
type MongoAllowListAssociateArrayInput interface {
	pulumi.Input

	ToMongoAllowListAssociateArrayOutput() MongoAllowListAssociateArrayOutput
	ToMongoAllowListAssociateArrayOutputWithContext(context.Context) MongoAllowListAssociateArrayOutput
}

type MongoAllowListAssociateArray []MongoAllowListAssociateInput

func (MongoAllowListAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoAllowListAssociate)(nil)).Elem()
}

func (i MongoAllowListAssociateArray) ToMongoAllowListAssociateArrayOutput() MongoAllowListAssociateArrayOutput {
	return i.ToMongoAllowListAssociateArrayOutputWithContext(context.Background())
}

func (i MongoAllowListAssociateArray) ToMongoAllowListAssociateArrayOutputWithContext(ctx context.Context) MongoAllowListAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoAllowListAssociateArrayOutput)
}

// MongoAllowListAssociateMapInput is an input type that accepts MongoAllowListAssociateMap and MongoAllowListAssociateMapOutput values.
// You can construct a concrete instance of `MongoAllowListAssociateMapInput` via:
//
//	MongoAllowListAssociateMap{ "key": MongoAllowListAssociateArgs{...} }
type MongoAllowListAssociateMapInput interface {
	pulumi.Input

	ToMongoAllowListAssociateMapOutput() MongoAllowListAssociateMapOutput
	ToMongoAllowListAssociateMapOutputWithContext(context.Context) MongoAllowListAssociateMapOutput
}

type MongoAllowListAssociateMap map[string]MongoAllowListAssociateInput

func (MongoAllowListAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoAllowListAssociate)(nil)).Elem()
}

func (i MongoAllowListAssociateMap) ToMongoAllowListAssociateMapOutput() MongoAllowListAssociateMapOutput {
	return i.ToMongoAllowListAssociateMapOutputWithContext(context.Background())
}

func (i MongoAllowListAssociateMap) ToMongoAllowListAssociateMapOutputWithContext(ctx context.Context) MongoAllowListAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoAllowListAssociateMapOutput)
}

type MongoAllowListAssociateOutput struct{ *pulumi.OutputState }

func (MongoAllowListAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoAllowListAssociate)(nil)).Elem()
}

func (o MongoAllowListAssociateOutput) ToMongoAllowListAssociateOutput() MongoAllowListAssociateOutput {
	return o
}

func (o MongoAllowListAssociateOutput) ToMongoAllowListAssociateOutputWithContext(ctx context.Context) MongoAllowListAssociateOutput {
	return o
}

// Id of allow list to associate.
func (o MongoAllowListAssociateOutput) AllowListId() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoAllowListAssociate) pulumi.StringOutput { return v.AllowListId }).(pulumi.StringOutput)
}

// Id of instance to associate.
func (o MongoAllowListAssociateOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoAllowListAssociate) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

type MongoAllowListAssociateArrayOutput struct{ *pulumi.OutputState }

func (MongoAllowListAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoAllowListAssociate)(nil)).Elem()
}

func (o MongoAllowListAssociateArrayOutput) ToMongoAllowListAssociateArrayOutput() MongoAllowListAssociateArrayOutput {
	return o
}

func (o MongoAllowListAssociateArrayOutput) ToMongoAllowListAssociateArrayOutputWithContext(ctx context.Context) MongoAllowListAssociateArrayOutput {
	return o
}

func (o MongoAllowListAssociateArrayOutput) Index(i pulumi.IntInput) MongoAllowListAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MongoAllowListAssociate {
		return vs[0].([]*MongoAllowListAssociate)[vs[1].(int)]
	}).(MongoAllowListAssociateOutput)
}

type MongoAllowListAssociateMapOutput struct{ *pulumi.OutputState }

func (MongoAllowListAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoAllowListAssociate)(nil)).Elem()
}

func (o MongoAllowListAssociateMapOutput) ToMongoAllowListAssociateMapOutput() MongoAllowListAssociateMapOutput {
	return o
}

func (o MongoAllowListAssociateMapOutput) ToMongoAllowListAssociateMapOutputWithContext(ctx context.Context) MongoAllowListAssociateMapOutput {
	return o
}

func (o MongoAllowListAssociateMapOutput) MapIndex(k pulumi.StringInput) MongoAllowListAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MongoAllowListAssociate {
		return vs[0].(map[string]*MongoAllowListAssociate)[vs[1].(string)]
	}).(MongoAllowListAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MongoAllowListAssociateInput)(nil)).Elem(), &MongoAllowListAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoAllowListAssociateArrayInput)(nil)).Elem(), MongoAllowListAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoAllowListAssociateMapInput)(nil)).Elem(), MongoAllowListAssociateMap{})
	pulumi.RegisterOutputType(MongoAllowListAssociateOutput{})
	pulumi.RegisterOutputType(MongoAllowListAssociateArrayOutput{})
	pulumi.RegisterOutputType(MongoAllowListAssociateMapOutput{})
}
