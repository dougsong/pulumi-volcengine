// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cen

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/cen"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cen.NewBandwidthPackage(ctx, "foo", &cen.BandwidthPackageArgs{
//				Bandwidth:                  pulumi.Int(32),
//				BillingType:                pulumi.String("PrePaid"),
//				CenBandwidthPackageName:    pulumi.String("tf-test"),
//				Description:                pulumi.String("tf-test1"),
//				LocalGeographicRegionSetId: pulumi.String("China"),
//				PeerGeographicRegionSetId:  pulumi.String("China"),
//				Period:                     pulumi.Int(1),
//				PeriodUnit:                 pulumi.String("Year"),
//				ProjectName:                pulumi.String("default"),
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
// CenBandwidthPackage can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:cen/bandwidthPackage:BandwidthPackage default cbp-4c2zaavbvh5f42****
//
// ```
type BandwidthPackage struct {
	pulumi.CustomResourceState

	// The account ID of the cen bandwidth package.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The bandwidth of the cen bandwidth package. Value: 2~10000.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`. Terraform will
	// only remove the PrePaid cen bandwidth package from the state file, not actually remove.
	BillingType pulumi.StringPtrOutput `pulumi:"billingType"`
	// The business status of the cen bandwidth package.
	BusinessStatus pulumi.StringOutput `pulumi:"businessStatus"`
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId pulumi.StringOutput `pulumi:"cenBandwidthPackageId"`
	// The name of the cen bandwidth package.
	CenBandwidthPackageName pulumi.StringOutput `pulumi:"cenBandwidthPackageName"`
	// The cen IDs of the bandwidth package.
	CenIds pulumi.StringArrayOutput `pulumi:"cenIds"`
	// The create time of the cen bandwidth package.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The deleted time of the cen bandwidth package.
	DeletedTime pulumi.StringOutput `pulumi:"deletedTime"`
	// The description of the cen bandwidth package.
	Description pulumi.StringOutput `pulumi:"description"`
	// The expired time of the cen bandwidth package.
	ExpiredTime pulumi.StringOutput `pulumi:"expiredTime"`
	// The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	LocalGeographicRegionSetId pulumi.StringPtrOutput `pulumi:"localGeographicRegionSetId"`
	// The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	PeerGeographicRegionSetId pulumi.StringPtrOutput `pulumi:"peerGeographicRegionSetId"`
	// The period of the cen bandwidth package. Default value is 1.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
	PeriodUnit pulumi.StringPtrOutput `pulumi:"periodUnit"`
	// The ProjectName of the cen bandwidth package.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The remain bandwidth of the cen bandwidth package.
	RemainingBandwidth pulumi.IntOutput `pulumi:"remainingBandwidth"`
	// The status of the cen bandwidth package.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags BandwidthPackageTagArrayOutput `pulumi:"tags"`
	// The update time of the cen bandwidth package.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewBandwidthPackage registers a new resource with the given unique name, arguments, and options.
func NewBandwidthPackage(ctx *pulumi.Context,
	name string, args *BandwidthPackageArgs, opts ...pulumi.ResourceOption) (*BandwidthPackage, error) {
	if args == nil {
		args = &BandwidthPackageArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource BandwidthPackage
	err := ctx.RegisterResource("volcengine:cen/bandwidthPackage:BandwidthPackage", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBandwidthPackage gets an existing BandwidthPackage resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBandwidthPackage(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BandwidthPackageState, opts ...pulumi.ResourceOption) (*BandwidthPackage, error) {
	var resource BandwidthPackage
	err := ctx.ReadResource("volcengine:cen/bandwidthPackage:BandwidthPackage", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BandwidthPackage resources.
type bandwidthPackageState struct {
	// The account ID of the cen bandwidth package.
	AccountId *string `pulumi:"accountId"`
	// The bandwidth of the cen bandwidth package. Value: 2~10000.
	Bandwidth *int `pulumi:"bandwidth"`
	// The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`. Terraform will
	// only remove the PrePaid cen bandwidth package from the state file, not actually remove.
	BillingType *string `pulumi:"billingType"`
	// The business status of the cen bandwidth package.
	BusinessStatus *string `pulumi:"businessStatus"`
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId *string `pulumi:"cenBandwidthPackageId"`
	// The name of the cen bandwidth package.
	CenBandwidthPackageName *string `pulumi:"cenBandwidthPackageName"`
	// The cen IDs of the bandwidth package.
	CenIds []string `pulumi:"cenIds"`
	// The create time of the cen bandwidth package.
	CreationTime *string `pulumi:"creationTime"`
	// The deleted time of the cen bandwidth package.
	DeletedTime *string `pulumi:"deletedTime"`
	// The description of the cen bandwidth package.
	Description *string `pulumi:"description"`
	// The expired time of the cen bandwidth package.
	ExpiredTime *string `pulumi:"expiredTime"`
	// The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	LocalGeographicRegionSetId *string `pulumi:"localGeographicRegionSetId"`
	// The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	PeerGeographicRegionSetId *string `pulumi:"peerGeographicRegionSetId"`
	// The period of the cen bandwidth package. Default value is 1.
	Period *int `pulumi:"period"`
	// The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The ProjectName of the cen bandwidth package.
	ProjectName *string `pulumi:"projectName"`
	// The remain bandwidth of the cen bandwidth package.
	RemainingBandwidth *int `pulumi:"remainingBandwidth"`
	// The status of the cen bandwidth package.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []BandwidthPackageTag `pulumi:"tags"`
	// The update time of the cen bandwidth package.
	UpdateTime *string `pulumi:"updateTime"`
}

type BandwidthPackageState struct {
	// The account ID of the cen bandwidth package.
	AccountId pulumi.StringPtrInput
	// The bandwidth of the cen bandwidth package. Value: 2~10000.
	Bandwidth pulumi.IntPtrInput
	// The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`. Terraform will
	// only remove the PrePaid cen bandwidth package from the state file, not actually remove.
	BillingType pulumi.StringPtrInput
	// The business status of the cen bandwidth package.
	BusinessStatus pulumi.StringPtrInput
	// The ID of the cen bandwidth package.
	CenBandwidthPackageId pulumi.StringPtrInput
	// The name of the cen bandwidth package.
	CenBandwidthPackageName pulumi.StringPtrInput
	// The cen IDs of the bandwidth package.
	CenIds pulumi.StringArrayInput
	// The create time of the cen bandwidth package.
	CreationTime pulumi.StringPtrInput
	// The deleted time of the cen bandwidth package.
	DeletedTime pulumi.StringPtrInput
	// The description of the cen bandwidth package.
	Description pulumi.StringPtrInput
	// The expired time of the cen bandwidth package.
	ExpiredTime pulumi.StringPtrInput
	// The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	LocalGeographicRegionSetId pulumi.StringPtrInput
	// The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	PeerGeographicRegionSetId pulumi.StringPtrInput
	// The period of the cen bandwidth package. Default value is 1.
	Period pulumi.IntPtrInput
	// The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
	PeriodUnit pulumi.StringPtrInput
	// The ProjectName of the cen bandwidth package.
	ProjectName pulumi.StringPtrInput
	// The remain bandwidth of the cen bandwidth package.
	RemainingBandwidth pulumi.IntPtrInput
	// The status of the cen bandwidth package.
	Status pulumi.StringPtrInput
	// Tags.
	Tags BandwidthPackageTagArrayInput
	// The update time of the cen bandwidth package.
	UpdateTime pulumi.StringPtrInput
}

func (BandwidthPackageState) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageState)(nil)).Elem()
}

type bandwidthPackageArgs struct {
	// The bandwidth of the cen bandwidth package. Value: 2~10000.
	Bandwidth *int `pulumi:"bandwidth"`
	// The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`. Terraform will
	// only remove the PrePaid cen bandwidth package from the state file, not actually remove.
	BillingType *string `pulumi:"billingType"`
	// The name of the cen bandwidth package.
	CenBandwidthPackageName *string `pulumi:"cenBandwidthPackageName"`
	// The description of the cen bandwidth package.
	Description *string `pulumi:"description"`
	// The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	LocalGeographicRegionSetId *string `pulumi:"localGeographicRegionSetId"`
	// The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	PeerGeographicRegionSetId *string `pulumi:"peerGeographicRegionSetId"`
	// The period of the cen bandwidth package. Default value is 1.
	Period *int `pulumi:"period"`
	// The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
	PeriodUnit *string `pulumi:"periodUnit"`
	// The ProjectName of the cen bandwidth package.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []BandwidthPackageTag `pulumi:"tags"`
}

// The set of arguments for constructing a BandwidthPackage resource.
type BandwidthPackageArgs struct {
	// The bandwidth of the cen bandwidth package. Value: 2~10000.
	Bandwidth pulumi.IntPtrInput
	// The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`. Terraform will
	// only remove the PrePaid cen bandwidth package from the state file, not actually remove.
	BillingType pulumi.StringPtrInput
	// The name of the cen bandwidth package.
	CenBandwidthPackageName pulumi.StringPtrInput
	// The description of the cen bandwidth package.
	Description pulumi.StringPtrInput
	// The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	LocalGeographicRegionSetId pulumi.StringPtrInput
	// The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
	PeerGeographicRegionSetId pulumi.StringPtrInput
	// The period of the cen bandwidth package. Default value is 1.
	Period pulumi.IntPtrInput
	// The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
	PeriodUnit pulumi.StringPtrInput
	// The ProjectName of the cen bandwidth package.
	ProjectName pulumi.StringPtrInput
	// Tags.
	Tags BandwidthPackageTagArrayInput
}

func (BandwidthPackageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bandwidthPackageArgs)(nil)).Elem()
}

type BandwidthPackageInput interface {
	pulumi.Input

	ToBandwidthPackageOutput() BandwidthPackageOutput
	ToBandwidthPackageOutputWithContext(ctx context.Context) BandwidthPackageOutput
}

func (*BandwidthPackage) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthPackage)(nil)).Elem()
}

func (i *BandwidthPackage) ToBandwidthPackageOutput() BandwidthPackageOutput {
	return i.ToBandwidthPackageOutputWithContext(context.Background())
}

func (i *BandwidthPackage) ToBandwidthPackageOutputWithContext(ctx context.Context) BandwidthPackageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageOutput)
}

// BandwidthPackageArrayInput is an input type that accepts BandwidthPackageArray and BandwidthPackageArrayOutput values.
// You can construct a concrete instance of `BandwidthPackageArrayInput` via:
//
//	BandwidthPackageArray{ BandwidthPackageArgs{...} }
type BandwidthPackageArrayInput interface {
	pulumi.Input

	ToBandwidthPackageArrayOutput() BandwidthPackageArrayOutput
	ToBandwidthPackageArrayOutputWithContext(context.Context) BandwidthPackageArrayOutput
}

type BandwidthPackageArray []BandwidthPackageInput

func (BandwidthPackageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BandwidthPackage)(nil)).Elem()
}

func (i BandwidthPackageArray) ToBandwidthPackageArrayOutput() BandwidthPackageArrayOutput {
	return i.ToBandwidthPackageArrayOutputWithContext(context.Background())
}

func (i BandwidthPackageArray) ToBandwidthPackageArrayOutputWithContext(ctx context.Context) BandwidthPackageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageArrayOutput)
}

// BandwidthPackageMapInput is an input type that accepts BandwidthPackageMap and BandwidthPackageMapOutput values.
// You can construct a concrete instance of `BandwidthPackageMapInput` via:
//
//	BandwidthPackageMap{ "key": BandwidthPackageArgs{...} }
type BandwidthPackageMapInput interface {
	pulumi.Input

	ToBandwidthPackageMapOutput() BandwidthPackageMapOutput
	ToBandwidthPackageMapOutputWithContext(context.Context) BandwidthPackageMapOutput
}

type BandwidthPackageMap map[string]BandwidthPackageInput

func (BandwidthPackageMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BandwidthPackage)(nil)).Elem()
}

func (i BandwidthPackageMap) ToBandwidthPackageMapOutput() BandwidthPackageMapOutput {
	return i.ToBandwidthPackageMapOutputWithContext(context.Background())
}

func (i BandwidthPackageMap) ToBandwidthPackageMapOutputWithContext(ctx context.Context) BandwidthPackageMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BandwidthPackageMapOutput)
}

type BandwidthPackageOutput struct{ *pulumi.OutputState }

func (BandwidthPackageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BandwidthPackage)(nil)).Elem()
}

func (o BandwidthPackageOutput) ToBandwidthPackageOutput() BandwidthPackageOutput {
	return o
}

func (o BandwidthPackageOutput) ToBandwidthPackageOutputWithContext(ctx context.Context) BandwidthPackageOutput {
	return o
}

// The account ID of the cen bandwidth package.
func (o BandwidthPackageOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The bandwidth of the cen bandwidth package. Value: 2~10000.
func (o BandwidthPackageOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The billing type of the cen bandwidth package. Only support `PrePaid` and default value is `PrePaid`. Terraform will
// only remove the PrePaid cen bandwidth package from the state file, not actually remove.
func (o BandwidthPackageOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.BillingType }).(pulumi.StringPtrOutput)
}

// The business status of the cen bandwidth package.
func (o BandwidthPackageOutput) BusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.BusinessStatus }).(pulumi.StringOutput)
}

// The ID of the cen bandwidth package.
func (o BandwidthPackageOutput) CenBandwidthPackageId() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.CenBandwidthPackageId }).(pulumi.StringOutput)
}

// The name of the cen bandwidth package.
func (o BandwidthPackageOutput) CenBandwidthPackageName() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.CenBandwidthPackageName }).(pulumi.StringOutput)
}

// The cen IDs of the bandwidth package.
func (o BandwidthPackageOutput) CenIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringArrayOutput { return v.CenIds }).(pulumi.StringArrayOutput)
}

// The create time of the cen bandwidth package.
func (o BandwidthPackageOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The deleted time of the cen bandwidth package.
func (o BandwidthPackageOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.DeletedTime }).(pulumi.StringOutput)
}

// The description of the cen bandwidth package.
func (o BandwidthPackageOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The expired time of the cen bandwidth package.
func (o BandwidthPackageOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.ExpiredTime }).(pulumi.StringOutput)
}

// The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
func (o BandwidthPackageOutput) LocalGeographicRegionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.LocalGeographicRegionSetId }).(pulumi.StringPtrOutput)
}

// The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.
func (o BandwidthPackageOutput) PeerGeographicRegionSetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.PeerGeographicRegionSetId }).(pulumi.StringPtrOutput)
}

// The period of the cen bandwidth package. Default value is 1.
func (o BandwidthPackageOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.
func (o BandwidthPackageOutput) PeriodUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.PeriodUnit }).(pulumi.StringPtrOutput)
}

// The ProjectName of the cen bandwidth package.
func (o BandwidthPackageOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The remain bandwidth of the cen bandwidth package.
func (o BandwidthPackageOutput) RemainingBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.IntOutput { return v.RemainingBandwidth }).(pulumi.IntOutput)
}

// The status of the cen bandwidth package.
func (o BandwidthPackageOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o BandwidthPackageOutput) Tags() BandwidthPackageTagArrayOutput {
	return o.ApplyT(func(v *BandwidthPackage) BandwidthPackageTagArrayOutput { return v.Tags }).(BandwidthPackageTagArrayOutput)
}

// The update time of the cen bandwidth package.
func (o BandwidthPackageOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BandwidthPackage) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type BandwidthPackageArrayOutput struct{ *pulumi.OutputState }

func (BandwidthPackageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BandwidthPackage)(nil)).Elem()
}

func (o BandwidthPackageArrayOutput) ToBandwidthPackageArrayOutput() BandwidthPackageArrayOutput {
	return o
}

func (o BandwidthPackageArrayOutput) ToBandwidthPackageArrayOutputWithContext(ctx context.Context) BandwidthPackageArrayOutput {
	return o
}

func (o BandwidthPackageArrayOutput) Index(i pulumi.IntInput) BandwidthPackageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BandwidthPackage {
		return vs[0].([]*BandwidthPackage)[vs[1].(int)]
	}).(BandwidthPackageOutput)
}

type BandwidthPackageMapOutput struct{ *pulumi.OutputState }

func (BandwidthPackageMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BandwidthPackage)(nil)).Elem()
}

func (o BandwidthPackageMapOutput) ToBandwidthPackageMapOutput() BandwidthPackageMapOutput {
	return o
}

func (o BandwidthPackageMapOutput) ToBandwidthPackageMapOutputWithContext(ctx context.Context) BandwidthPackageMapOutput {
	return o
}

func (o BandwidthPackageMapOutput) MapIndex(k pulumi.StringInput) BandwidthPackageOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BandwidthPackage {
		return vs[0].(map[string]*BandwidthPackage)[vs[1].(string)]
	}).(BandwidthPackageOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageInput)(nil)).Elem(), &BandwidthPackage{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageArrayInput)(nil)).Elem(), BandwidthPackageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BandwidthPackageMapInput)(nil)).Elem(), BandwidthPackageMap{})
	pulumi.RegisterOutputType(BandwidthPackageOutput{})
	pulumi.RegisterOutputType(BandwidthPackageArrayOutput{})
	pulumi.RegisterOutputType(BandwidthPackageMapOutput{})
}
