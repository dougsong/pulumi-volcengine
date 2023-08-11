// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage ecs key pair
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/ecs"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ecs.NewKeyPair(ctx, "foo", &ecs.KeyPairArgs{
//				Description: pulumi.String("acc-test"),
//				KeyPairName: pulumi.String("acc-test-key-name"),
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
// ECS key pair can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:ecs/keyPair:KeyPair default kp-mizl7m1kqccg5smt1bdpijuj
//
// ```
type KeyPair struct {
	pulumi.CustomResourceState

	// The description of key pair.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The finger print info.
	FingerPrint pulumi.StringOutput `pulumi:"fingerPrint"`
	// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
	KeyFile pulumi.StringPtrOutput `pulumi:"keyFile"`
	// The id of key pair.
	KeyPairId pulumi.StringOutput `pulumi:"keyPairId"`
	// The name of key pair.
	KeyPairName pulumi.StringOutput `pulumi:"keyPairName"`
	// Public key string.
	PublicKey pulumi.StringPtrOutput `pulumi:"publicKey"`
}

// NewKeyPair registers a new resource with the given unique name, arguments, and options.
func NewKeyPair(ctx *pulumi.Context,
	name string, args *KeyPairArgs, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.KeyPairName == nil {
		return nil, errors.New("invalid value for required argument 'KeyPairName'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource KeyPair
	err := ctx.RegisterResource("volcengine:ecs/keyPair:KeyPair", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyPair gets an existing KeyPair resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyPair(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyPairState, opts ...pulumi.ResourceOption) (*KeyPair, error) {
	var resource KeyPair
	err := ctx.ReadResource("volcengine:ecs/keyPair:KeyPair", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyPair resources.
type keyPairState struct {
	// The description of key pair.
	Description *string `pulumi:"description"`
	// The finger print info.
	FingerPrint *string `pulumi:"fingerPrint"`
	// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
	KeyFile *string `pulumi:"keyFile"`
	// The id of key pair.
	KeyPairId *string `pulumi:"keyPairId"`
	// The name of key pair.
	KeyPairName *string `pulumi:"keyPairName"`
	// Public key string.
	PublicKey *string `pulumi:"publicKey"`
}

type KeyPairState struct {
	// The description of key pair.
	Description pulumi.StringPtrInput
	// The finger print info.
	FingerPrint pulumi.StringPtrInput
	// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
	KeyFile pulumi.StringPtrInput
	// The id of key pair.
	KeyPairId pulumi.StringPtrInput
	// The name of key pair.
	KeyPairName pulumi.StringPtrInput
	// Public key string.
	PublicKey pulumi.StringPtrInput
}

func (KeyPairState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairState)(nil)).Elem()
}

type keyPairArgs struct {
	// The description of key pair.
	Description *string `pulumi:"description"`
	// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
	KeyFile *string `pulumi:"keyFile"`
	// The name of key pair.
	KeyPairName string `pulumi:"keyPairName"`
	// Public key string.
	PublicKey *string `pulumi:"publicKey"`
}

// The set of arguments for constructing a KeyPair resource.
type KeyPairArgs struct {
	// The description of key pair.
	Description pulumi.StringPtrInput
	// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
	KeyFile pulumi.StringPtrInput
	// The name of key pair.
	KeyPairName pulumi.StringInput
	// Public key string.
	PublicKey pulumi.StringPtrInput
}

func (KeyPairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyPairArgs)(nil)).Elem()
}

type KeyPairInput interface {
	pulumi.Input

	ToKeyPairOutput() KeyPairOutput
	ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput
}

func (*KeyPair) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil)).Elem()
}

func (i *KeyPair) ToKeyPairOutput() KeyPairOutput {
	return i.ToKeyPairOutputWithContext(context.Background())
}

func (i *KeyPair) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairOutput)
}

// KeyPairArrayInput is an input type that accepts KeyPairArray and KeyPairArrayOutput values.
// You can construct a concrete instance of `KeyPairArrayInput` via:
//
//	KeyPairArray{ KeyPairArgs{...} }
type KeyPairArrayInput interface {
	pulumi.Input

	ToKeyPairArrayOutput() KeyPairArrayOutput
	ToKeyPairArrayOutputWithContext(context.Context) KeyPairArrayOutput
}

type KeyPairArray []KeyPairInput

func (KeyPairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyPair)(nil)).Elem()
}

func (i KeyPairArray) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return i.ToKeyPairArrayOutputWithContext(context.Background())
}

func (i KeyPairArray) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairArrayOutput)
}

// KeyPairMapInput is an input type that accepts KeyPairMap and KeyPairMapOutput values.
// You can construct a concrete instance of `KeyPairMapInput` via:
//
//	KeyPairMap{ "key": KeyPairArgs{...} }
type KeyPairMapInput interface {
	pulumi.Input

	ToKeyPairMapOutput() KeyPairMapOutput
	ToKeyPairMapOutputWithContext(context.Context) KeyPairMapOutput
}

type KeyPairMap map[string]KeyPairInput

func (KeyPairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyPair)(nil)).Elem()
}

func (i KeyPairMap) ToKeyPairMapOutput() KeyPairMapOutput {
	return i.ToKeyPairMapOutputWithContext(context.Background())
}

func (i KeyPairMap) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyPairMapOutput)
}

type KeyPairOutput struct{ *pulumi.OutputState }

func (KeyPairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyPair)(nil)).Elem()
}

func (o KeyPairOutput) ToKeyPairOutput() KeyPairOutput {
	return o
}

func (o KeyPairOutput) ToKeyPairOutputWithContext(ctx context.Context) KeyPairOutput {
	return o
}

// The description of key pair.
func (o KeyPairOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyPair) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The finger print info.
func (o KeyPairOutput) FingerPrint() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyPair) pulumi.StringOutput { return v.FingerPrint }).(pulumi.StringOutput)
}

// Target file to save private key. It is recommended that the value not be empty. You only have one chance to download the private key, the volcengine will not save your private key, please keep it safe. In the TF import scenario, this field will not write the private key locally.
func (o KeyPairOutput) KeyFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyPair) pulumi.StringPtrOutput { return v.KeyFile }).(pulumi.StringPtrOutput)
}

// The id of key pair.
func (o KeyPairOutput) KeyPairId() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyPair) pulumi.StringOutput { return v.KeyPairId }).(pulumi.StringOutput)
}

// The name of key pair.
func (o KeyPairOutput) KeyPairName() pulumi.StringOutput {
	return o.ApplyT(func(v *KeyPair) pulumi.StringOutput { return v.KeyPairName }).(pulumi.StringOutput)
}

// Public key string.
func (o KeyPairOutput) PublicKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyPair) pulumi.StringPtrOutput { return v.PublicKey }).(pulumi.StringPtrOutput)
}

type KeyPairArrayOutput struct{ *pulumi.OutputState }

func (KeyPairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*KeyPair)(nil)).Elem()
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutput() KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) ToKeyPairArrayOutputWithContext(ctx context.Context) KeyPairArrayOutput {
	return o
}

func (o KeyPairArrayOutput) Index(i pulumi.IntInput) KeyPairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *KeyPair {
		return vs[0].([]*KeyPair)[vs[1].(int)]
	}).(KeyPairOutput)
}

type KeyPairMapOutput struct{ *pulumi.OutputState }

func (KeyPairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*KeyPair)(nil)).Elem()
}

func (o KeyPairMapOutput) ToKeyPairMapOutput() KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) ToKeyPairMapOutputWithContext(ctx context.Context) KeyPairMapOutput {
	return o
}

func (o KeyPairMapOutput) MapIndex(k pulumi.StringInput) KeyPairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *KeyPair {
		return vs[0].(map[string]*KeyPair)[vs[1].(string)]
	}).(KeyPairOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairInput)(nil)).Elem(), &KeyPair{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairArrayInput)(nil)).Elem(), KeyPairArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyPairMapInput)(nil)).Elem(), KeyPairMap{})
	pulumi.RegisterOutputType(KeyPairOutput{})
	pulumi.RegisterOutputType(KeyPairArrayOutput{})
	pulumi.RegisterOutputType(KeyPairMapOutput{})
}
