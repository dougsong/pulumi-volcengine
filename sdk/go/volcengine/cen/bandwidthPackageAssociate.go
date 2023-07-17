// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage cen bandwidth package associate
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/cen"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.NewBandwidthPackageAssociate(ctx, "foo", &cen.BandwidthPackageAssociateArgs{
//				CenBandwidthPackageId: pulumi.String("cbp-2bzeew3s8p79c2dx0eeohej4x"),
//				CenId:                 pulumi.String("cen-2bzrl3srxsv0g2dx0efyoojn3"),
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
// Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
//
// ```
type BandwidthPackageAssociate struct {
	pulumi.CustomResourceState

	// The ID of the cen bandwidth package.
	CenBandwidthPackageId pulumi.StringOutput `pulumi:"cenBandwidthPackageId"`
	// The ID of the cen.
	CenId pulumi.StringOutput `pulumi:"cenId"`
}

// NewBandwidthPackageAssociate registers a new resource with the given unique name, arguments, and options.
func NewBandwidthPackageAssociate(ctx *pulumi.Context,
	name string, args *BandwidthPackageAssociateArgs, opts ...pulumi.ResourceOption) (*BandwidthPackageAssociate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CenBandwidthPackageId == nil {
		return nil, errors.New("invalid value for required argument 'CenBandwidthPackageId'")
	}
	if args.CenId == nil {
		return nil, errors.New("invalid value for required argument 'CenId'")
	}
	var resource BandwidthPackageAssociate
	err := ctx.RegisterResource("volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthPackageAssociate gets an existing BandwidthPackageAssociate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthPackageAssociate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthPackageAssociateState, opts ...pulumi.ResourceOption) (*BandwidthPackageAssociate, error) {
	var resource BandwidthPackageAssociate
	err := ctx.ReadResource("volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthPackageAssociate resources.
type bandwidthPackageAssociateState struct {
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId *string `pulumi:"cenBandwidthPackageId"`
	// The ID of the cen.
	CenId *string `pulumi:"cenId"`
}

type BandwidthPackageAssociateState struct {
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId pulumi.StringPtrInput
	// The ID of the cen.
	CenId pulumi.StringPtrInput
}

func (BandwidthPackageAssociateState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageAssociateState)(nil)).Elem()
}

type bandwidthPackageAssociateArgs struct {
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId string `pulumi:"cenBandwidthPackageId"`
	// The ID of the cen.
	CenId string `pulumi:"cenId"`
}

// The set of arguments for constructing a BandwidthPackageAssociate resource.
type BandwidthPackageAssociateArgs struct {
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId pulumi.StringInput
	// The ID of the cen.
	CenId pulumi.StringInput
}

func (BandwidthPackageAssociateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageAssociateArgs)(nil)).Elem()
}

type BandwidthPackageAssociateInput interface {
	pulumi.Input

	ToBandwidthPackageAssociateOutput() BandwidthPackageAssociateOutput
	ToBandwidthPackageAssociateOutputWithContext(ctx context.Context) BandwidthPackageAssociateOutput
}

func (*BandwidthPackageAssociate) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthPackageAssociate)(nil)).Elem()
}

func (i *BandwidthPackageAssociate) ToBandwidthPackageAssociateOutput() BandwidthPackageAssociateOutput {
	return i.ToBandwidthPackageAssociateOutputWithContext(context.Background())
}

func (i *BandwidthPackageAssociate) ToBandwidthPackageAssociateOutputWithContext(ctx context.Context) BandwidthPackageAssociateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageAssociateOutput)
}

// BandwidthPackageAssociateArrayInput is an input type that accepts BandwidthPackageAssociateArray and BandwidthPackageAssociateArrayOutput values.
// You can construct a concrete instance of `BandwidthPackageAssociateArrayInput` via:
//
//	BandwidthPackageAssociateArray{ BandwidthPackageAssociateArgs{...} }
type BandwidthPackageAssociateArrayInput interface {
	pulumi.Input

	ToBandwidthPackageAssociateArrayOutput() BandwidthPackageAssociateArrayOutput
	ToBandwidthPackageAssociateArrayOutputWithContext(context.Context) BandwidthPackageAssociateArrayOutput
}

type BandwidthPackageAssociateArray []BandwidthPackageAssociateInput

func (BandwidthPackageAssociateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BandwidthPackageAssociate)(nil)).Elem()
}

func (i BandwidthPackageAssociateArray) ToBandwidthPackageAssociateArrayOutput() BandwidthPackageAssociateArrayOutput {
	return i.ToBandwidthPackageAssociateArrayOutputWithContext(context.Background())
}

func (i BandwidthPackageAssociateArray) ToBandwidthPackageAssociateArrayOutputWithContext(ctx context.Context) BandwidthPackageAssociateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageAssociateArrayOutput)
}

// BandwidthPackageAssociateMapInput is an input type that accepts BandwidthPackageAssociateMap and BandwidthPackageAssociateMapOutput values.
// You can construct a concrete instance of `BandwidthPackageAssociateMapInput` via:
//
//	BandwidthPackageAssociateMap{ "key": BandwidthPackageAssociateArgs{...} }
type BandwidthPackageAssociateMapInput interface {
	pulumi.Input

	ToBandwidthPackageAssociateMapOutput() BandwidthPackageAssociateMapOutput
	ToBandwidthPackageAssociateMapOutputWithContext(context.Context) BandwidthPackageAssociateMapOutput
}

type BandwidthPackageAssociateMap map[string]BandwidthPackageAssociateInput

func (BandwidthPackageAssociateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BandwidthPackageAssociate)(nil)).Elem()
}

func (i BandwidthPackageAssociateMap) ToBandwidthPackageAssociateMapOutput() BandwidthPackageAssociateMapOutput {
	return i.ToBandwidthPackageAssociateMapOutputWithContext(context.Background())
}

func (i BandwidthPackageAssociateMap) ToBandwidthPackageAssociateMapOutputWithContext(ctx context.Context) BandwidthPackageAssociateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageAssociateMapOutput)
}

type BandwidthPackageAssociateOutput struct{ *pulumi.OutputState }

func (BandwidthPackageAssociateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthPackageAssociate)(nil)).Elem()
}

func (o BandwidthPackageAssociateOutput) ToBandwidthPackageAssociateOutput() BandwidthPackageAssociateOutput {
	return o
}

func (o BandwidthPackageAssociateOutput) ToBandwidthPackageAssociateOutputWithContext(ctx context.Context) BandwidthPackageAssociateOutput {
	return o
}

// The ID of the cen bandwidth package.
func (o BandwidthPackageAssociateOutput) CenBandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackageAssociate) pulumi.StringOutput { return v.CenBandwidthPackageId }).(pulumi.StringOutput)
}

// The ID of the cen.
func (o BandwidthPackageAssociateOutput) CenId() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackageAssociate) pulumi.StringOutput { return v.CenId }).(pulumi.StringOutput)
}

type BandwidthPackageAssociateArrayOutput struct{ *pulumi.OutputState }

func (BandwidthPackageAssociateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BandwidthPackageAssociate)(nil)).Elem()
}

func (o BandwidthPackageAssociateArrayOutput) ToBandwidthPackageAssociateArrayOutput() BandwidthPackageAssociateArrayOutput {
	return o
}

func (o BandwidthPackageAssociateArrayOutput) ToBandwidthPackageAssociateArrayOutputWithContext(ctx context.Context) BandwidthPackageAssociateArrayOutput {
	return o
}

func (o BandwidthPackageAssociateArrayOutput) Index(i pulumi.IntInput) BandwidthPackageAssociateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BandwidthPackageAssociate {
		return vs[0].([]*BandwidthPackageAssociate)[vs[1].(int)]
	}).(BandwidthPackageAssociateOutput)
}

type BandwidthPackageAssociateMapOutput struct{ *pulumi.OutputState }

func (BandwidthPackageAssociateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BandwidthPackageAssociate)(nil)).Elem()
}

func (o BandwidthPackageAssociateMapOutput) ToBandwidthPackageAssociateMapOutput() BandwidthPackageAssociateMapOutput {
	return o
}

func (o BandwidthPackageAssociateMapOutput) ToBandwidthPackageAssociateMapOutputWithContext(ctx context.Context) BandwidthPackageAssociateMapOutput {
	return o
}

func (o BandwidthPackageAssociateMapOutput) MapIndex(k pulumi.StringInput) BandwidthPackageAssociateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BandwidthPackageAssociate {
		return vs[0].(map[string]*BandwidthPackageAssociate)[vs[1].(string)]
	}).(BandwidthPackageAssociateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageAssociateInput)(nil)).Elem(), &BandwidthPackageAssociate{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageAssociateArrayInput)(nil)).Elem(), BandwidthPackageAssociateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageAssociateMapInput)(nil)).Elem(), BandwidthPackageAssociateMap{})
	pulumi.RegisterOutputType(BandwidthPackageAssociateOutput{})
	pulumi.RegisterOutputType(BandwidthPackageAssociateArrayOutput{})
	pulumi.RegisterOutputType(BandwidthPackageAssociateMapOutput{})
}
