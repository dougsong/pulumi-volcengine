// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package eip

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/eip"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := eip.NewAddress(ctx, "foo", &eip.AddressArgs{
//				Bandwidth:   pulumi.Int(1),
//				BillingType: pulumi.String("PostPaidByBandwidth"),
//				Description: pulumi.String("acc-test"),
//				Isp:         pulumi.String("ChinaUnicom"),
//				ProjectName: pulumi.String("default"),
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
// Eip address can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:eip/address:Address default eip-274oj9a8rs9a87fap8sf9515b
//
// ```
type Address struct {
	pulumi.CustomResourceState

	// The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
	BillingType pulumi.StringOutput `pulumi:"billingType"`
	// The deleted time of the EIP.
	DeletedTime pulumi.StringOutput `pulumi:"deletedTime"`
	// The description of the EIP.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ip address of the EIP.
	EipAddress pulumi.StringOutput `pulumi:"eipAddress"`
	// The expired time of the EIP.
	ExpiredTime pulumi.StringOutput `pulumi:"expiredTime"`
	// The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
	Isp pulumi.StringOutput `pulumi:"isp"`
	// The name of the EIP Address.
	Name pulumi.StringOutput `pulumi:"name"`
	// The overdue time of the EIP.
	OverdueTime pulumi.StringOutput `pulumi:"overdueTime"`
	// The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billingType from PostPaid to PrePaid.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The ProjectName of the EIP.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The status of the EIP.
	Status pulumi.StringOutput `pulumi:"status"`
	// Tags.
	Tags AddressTagArrayOutput `pulumi:"tags"`
}

// NewAddress registers a new resource with the given unique name, arguments, and options.
func NewAddress(ctx *pulumi.Context,
	name string, args *AddressArgs, opts ...pulumi.ResourceOption) (*Address, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingType == nil {
		return nil, errors.New("invalid value for required argument 'BillingType'")
	}
	var resource Address
	err := ctx.RegisterResource("volcengine:eip/address:Address", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAddress gets an existing Address resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAddress(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AddressState, opts ...pulumi.ResourceOption) (*Address, error) {
	var resource Address
	err := ctx.ReadResource("volcengine:eip/address:Address", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Address resources.
type addressState struct {
	// The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
	Bandwidth *int `pulumi:"bandwidth"`
	// The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
	BillingType *string `pulumi:"billingType"`
	// The deleted time of the EIP.
	DeletedTime *string `pulumi:"deletedTime"`
	// The description of the EIP.
	Description *string `pulumi:"description"`
	// The ip address of the EIP.
	EipAddress *string `pulumi:"eipAddress"`
	// The expired time of the EIP.
	ExpiredTime *string `pulumi:"expiredTime"`
	// The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
	Isp *string `pulumi:"isp"`
	// The name of the EIP Address.
	Name *string `pulumi:"name"`
	// The overdue time of the EIP.
	OverdueTime *string `pulumi:"overdueTime"`
	// The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billingType from PostPaid to PrePaid.
	Period *int `pulumi:"period"`
	// The ProjectName of the EIP.
	ProjectName *string `pulumi:"projectName"`
	// The status of the EIP.
	Status *string `pulumi:"status"`
	// Tags.
	Tags []AddressTag `pulumi:"tags"`
}

type AddressState struct {
	// The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
	Bandwidth pulumi.IntPtrInput
	// The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
	BillingType pulumi.StringPtrInput
	// The deleted time of the EIP.
	DeletedTime pulumi.StringPtrInput
	// The description of the EIP.
	Description pulumi.StringPtrInput
	// The ip address of the EIP.
	EipAddress pulumi.StringPtrInput
	// The expired time of the EIP.
	ExpiredTime pulumi.StringPtrInput
	// The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
	Isp pulumi.StringPtrInput
	// The name of the EIP Address.
	Name pulumi.StringPtrInput
	// The overdue time of the EIP.
	OverdueTime pulumi.StringPtrInput
	// The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billingType from PostPaid to PrePaid.
	Period pulumi.IntPtrInput
	// The ProjectName of the EIP.
	ProjectName pulumi.StringPtrInput
	// The status of the EIP.
	Status pulumi.StringPtrInput
	// Tags.
	Tags AddressTagArrayInput
}

func (AddressState) ElementType() reflect.Type {
	return reflect.TypeOf((*addressState)(nil)).Elem()
}

type addressArgs struct {
	// The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
	Bandwidth *int `pulumi:"bandwidth"`
	// The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
	BillingType string `pulumi:"billingType"`
	// The description of the EIP.
	Description *string `pulumi:"description"`
	// The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
	Isp *string `pulumi:"isp"`
	// The name of the EIP Address.
	Name *string `pulumi:"name"`
	// The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billingType from PostPaid to PrePaid.
	Period *int `pulumi:"period"`
	// The ProjectName of the EIP.
	ProjectName *string `pulumi:"projectName"`
	// Tags.
	Tags []AddressTag `pulumi:"tags"`
}

// The set of arguments for constructing a Address resource.
type AddressArgs struct {
	// The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
	Bandwidth pulumi.IntPtrInput
	// The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
	BillingType pulumi.StringInput
	// The description of the EIP.
	Description pulumi.StringPtrInput
	// The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
	Isp pulumi.StringPtrInput
	// The name of the EIP Address.
	Name pulumi.StringPtrInput
	// The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billingType from PostPaid to PrePaid.
	Period pulumi.IntPtrInput
	// The ProjectName of the EIP.
	ProjectName pulumi.StringPtrInput
	// Tags.
	Tags AddressTagArrayInput
}

func (AddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*addressArgs)(nil)).Elem()
}

type AddressInput interface {
	pulumi.Input

	ToAddressOutput() AddressOutput
	ToAddressOutputWithContext(ctx context.Context) AddressOutput
}

func (*Address) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (i *Address) ToAddressOutput() AddressOutput {
	return i.ToAddressOutputWithContext(context.Background())
}

func (i *Address) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressOutput)
}

// AddressArrayInput is an input type that accepts AddressArray and AddressArrayOutput values.
// You can construct a concrete instance of `AddressArrayInput` via:
//
//	AddressArray{ AddressArgs{...} }
type AddressArrayInput interface {
	pulumi.Input

	ToAddressArrayOutput() AddressArrayOutput
	ToAddressArrayOutputWithContext(context.Context) AddressArrayOutput
}

type AddressArray []AddressInput

func (AddressArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address)(nil)).Elem()
}

func (i AddressArray) ToAddressArrayOutput() AddressArrayOutput {
	return i.ToAddressArrayOutputWithContext(context.Background())
}

func (i AddressArray) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressArrayOutput)
}

// AddressMapInput is an input type that accepts AddressMap and AddressMapOutput values.
// You can construct a concrete instance of `AddressMapInput` via:
//
//	AddressMap{ "key": AddressArgs{...} }
type AddressMapInput interface {
	pulumi.Input

	ToAddressMapOutput() AddressMapOutput
	ToAddressMapOutputWithContext(context.Context) AddressMapOutput
}

type AddressMap map[string]AddressInput

func (AddressMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address)(nil)).Elem()
}

func (i AddressMap) ToAddressMapOutput() AddressMapOutput {
	return i.ToAddressMapOutputWithContext(context.Background())
}

func (i AddressMap) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddressMapOutput)
}

type AddressOutput struct{ *pulumi.OutputState }

func (AddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Address)(nil)).Elem()
}

func (o AddressOutput) ToAddressOutput() AddressOutput {
	return o
}

func (o AddressOutput) ToAddressOutputWithContext(ctx context.Context) AddressOutput {
	return o
}

// The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.
func (o AddressOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Address) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
func (o AddressOutput) BillingType() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.BillingType }).(pulumi.StringOutput)
}

// The deleted time of the EIP.
func (o AddressOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.DeletedTime }).(pulumi.StringOutput)
}

// The description of the EIP.
func (o AddressOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ip address of the EIP.
func (o AddressOutput) EipAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.EipAddress }).(pulumi.StringOutput)
}

// The expired time of the EIP.
func (o AddressOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.ExpiredTime }).(pulumi.StringOutput)
}

// The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
func (o AddressOutput) Isp() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Isp }).(pulumi.StringOutput)
}

// The name of the EIP Address.
func (o AddressOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The overdue time of the EIP.
func (o AddressOutput) OverdueTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.OverdueTime }).(pulumi.StringOutput)
}

// The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billingType from PostPaid to PrePaid.
func (o AddressOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The ProjectName of the EIP.
func (o AddressOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Address) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The status of the EIP.
func (o AddressOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Address) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Tags.
func (o AddressOutput) Tags() AddressTagArrayOutput {
	return o.ApplyT(func(v *Address) AddressTagArrayOutput { return v.Tags }).(AddressTagArrayOutput)
}

type AddressArrayOutput struct{ *pulumi.OutputState }

func (AddressArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Address)(nil)).Elem()
}

func (o AddressArrayOutput) ToAddressArrayOutput() AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) ToAddressArrayOutputWithContext(ctx context.Context) AddressArrayOutput {
	return o
}

func (o AddressArrayOutput) Index(i pulumi.IntInput) AddressOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Address {
		return vs[0].([]*Address)[vs[1].(int)]
	}).(AddressOutput)
}

type AddressMapOutput struct{ *pulumi.OutputState }

func (AddressMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Address)(nil)).Elem()
}

func (o AddressMapOutput) ToAddressMapOutput() AddressMapOutput {
	return o
}

func (o AddressMapOutput) ToAddressMapOutputWithContext(ctx context.Context) AddressMapOutput {
	return o
}

func (o AddressMapOutput) MapIndex(k pulumi.StringInput) AddressOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Address {
		return vs[0].(map[string]*Address)[vs[1].(string)]
	}).(AddressOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AddressInput)(nil)).Elem(), &Address{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressArrayInput)(nil)).Elem(), AddressArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AddressMapInput)(nil)).Elem(), AddressMap{})
	pulumi.RegisterOutputType(AddressOutput{})
	pulumi.RegisterOutputType(AddressArrayOutput{})
	pulumi.RegisterOutputType(AddressMapOutput{})
}
