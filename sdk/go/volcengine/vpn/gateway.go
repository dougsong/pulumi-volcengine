// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

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
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpn"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpn.NewGateway(ctx, "foo", &vpn.GatewayArgs{
//				Bandwidth:      pulumi.Int(20),
//				Description:    pulumi.String("tf-test"),
//				Period:         pulumi.Int(2),
//				ProjectName:    pulumi.String("default"),
//				SubnetId:       pulumi.String("subnet-12bh8g2d7fshs17q7y2nx82uk"),
//				VpcId:          pulumi.String("vpc-12b31m7z2kc8w17q7y2fih9ts"),
//				VpnGatewayName: pulumi.String("tf-test"),
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
// VpnGateway can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vpn/gateway:Gateway default vgw-273zkshb2qayo7fap8t2****
//
// ```
type Gateway struct {
	pulumi.CustomResourceState

	// The account ID of the VPN gateway.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
	Bandwidth pulumi.IntOutput `pulumi:"bandwidth"`
	// The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
	// state file, not actually remove.
	BillingType pulumi.StringPtrOutput `pulumi:"billingType"`
	// The business status of the VPN gateway.
	BusinessStatus pulumi.StringOutput `pulumi:"businessStatus"`
	// The connection count of the VPN gateway.
	ConnectionCount pulumi.IntOutput `pulumi:"connectionCount"`
	// The create time of VPN gateway.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The deleted time of the VPN gateway.
	DeletedTime pulumi.StringOutput `pulumi:"deletedTime"`
	// The description of the VPN gateway.
	Description pulumi.StringOutput `pulumi:"description"`
	// The expired time of the VPN gateway.
	ExpiredTime pulumi.StringOutput `pulumi:"expiredTime"`
	// The IP address of the VPN gateway.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The lock reason of the VPN gateway.
	LockReason pulumi.StringOutput `pulumi:"lockReason"`
	// The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
	// Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Period pulumi.IntPtrOutput `pulumi:"period"`
	// The project name of the VPN gateway.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The renew type of the VPN gateway.
	RenewType pulumi.StringOutput `pulumi:"renewType"`
	// The route count of the VPN gateway.
	RouteCount pulumi.IntOutput `pulumi:"routeCount"`
	// The status of the VPN gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// The ID of the subnet where you want to create the VPN gateway.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Tags.
	Tags GatewayTagArrayOutput `pulumi:"tags"`
	// The update time of VPN gateway.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The ID of the VPC where you want to create the VPN gateway.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringOutput `pulumi:"vpnGatewayId"`
	// The name of the VPN gateway.
	VpnGatewayName pulumi.StringOutput `pulumi:"vpnGatewayName"`
}

// NewGateway registers a new resource with the given unique name, arguments, and options.
func NewGateway(ctx *pulumi.Context,
	name string, args *GatewayArgs, opts ...pulumi.ResourceOption) (*Gateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bandwidth == nil {
		return nil, errors.New("invalid value for required argument 'Bandwidth'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.VpcId == nil {
		return nil, errors.New("invalid value for required argument 'VpcId'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Gateway
	err := ctx.RegisterResource("volcengine:vpn/gateway:Gateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGateway gets an existing Gateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GatewayState, opts ...pulumi.ResourceOption) (*Gateway, error) {
	var resource Gateway
	err := ctx.ReadResource("volcengine:vpn/gateway:Gateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gateway resources.
type gatewayState struct {
	// The account ID of the VPN gateway.
	AccountId *string `pulumi:"accountId"`
	// The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
	Bandwidth *int `pulumi:"bandwidth"`
	// The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
	// state file, not actually remove.
	BillingType *string `pulumi:"billingType"`
	// The business status of the VPN gateway.
	BusinessStatus *string `pulumi:"businessStatus"`
	// The connection count of the VPN gateway.
	ConnectionCount *int `pulumi:"connectionCount"`
	// The create time of VPN gateway.
	CreationTime *string `pulumi:"creationTime"`
	// The deleted time of the VPN gateway.
	DeletedTime *string `pulumi:"deletedTime"`
	// The description of the VPN gateway.
	Description *string `pulumi:"description"`
	// The expired time of the VPN gateway.
	ExpiredTime *string `pulumi:"expiredTime"`
	// The IP address of the VPN gateway.
	IpAddress *string `pulumi:"ipAddress"`
	// The lock reason of the VPN gateway.
	LockReason *string `pulumi:"lockReason"`
	// The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
	// Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Period *int `pulumi:"period"`
	// The project name of the VPN gateway.
	ProjectName *string `pulumi:"projectName"`
	// The renew type of the VPN gateway.
	RenewType *string `pulumi:"renewType"`
	// The route count of the VPN gateway.
	RouteCount *int `pulumi:"routeCount"`
	// The status of the VPN gateway.
	Status *string `pulumi:"status"`
	// The ID of the subnet where you want to create the VPN gateway.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []GatewayTag `pulumi:"tags"`
	// The update time of VPN gateway.
	UpdateTime *string `pulumi:"updateTime"`
	// The ID of the VPC where you want to create the VPN gateway.
	VpcId *string `pulumi:"vpcId"`
	// The ID of the VPN gateway.
	VpnGatewayId *string `pulumi:"vpnGatewayId"`
	// The name of the VPN gateway.
	VpnGatewayName *string `pulumi:"vpnGatewayName"`
}

type GatewayState struct {
	// The account ID of the VPN gateway.
	AccountId pulumi.StringPtrInput
	// The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
	Bandwidth pulumi.IntPtrInput
	// The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
	// state file, not actually remove.
	BillingType pulumi.StringPtrInput
	// The business status of the VPN gateway.
	BusinessStatus pulumi.StringPtrInput
	// The connection count of the VPN gateway.
	ConnectionCount pulumi.IntPtrInput
	// The create time of VPN gateway.
	CreationTime pulumi.StringPtrInput
	// The deleted time of the VPN gateway.
	DeletedTime pulumi.StringPtrInput
	// The description of the VPN gateway.
	Description pulumi.StringPtrInput
	// The expired time of the VPN gateway.
	ExpiredTime pulumi.StringPtrInput
	// The IP address of the VPN gateway.
	IpAddress pulumi.StringPtrInput
	// The lock reason of the VPN gateway.
	LockReason pulumi.StringPtrInput
	// The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
	// Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Period pulumi.IntPtrInput
	// The project name of the VPN gateway.
	ProjectName pulumi.StringPtrInput
	// The renew type of the VPN gateway.
	RenewType pulumi.StringPtrInput
	// The route count of the VPN gateway.
	RouteCount pulumi.IntPtrInput
	// The status of the VPN gateway.
	Status pulumi.StringPtrInput
	// The ID of the subnet where you want to create the VPN gateway.
	SubnetId pulumi.StringPtrInput
	// Tags.
	Tags GatewayTagArrayInput
	// The update time of VPN gateway.
	UpdateTime pulumi.StringPtrInput
	// The ID of the VPC where you want to create the VPN gateway.
	VpcId pulumi.StringPtrInput
	// The ID of the VPN gateway.
	VpnGatewayId pulumi.StringPtrInput
	// The name of the VPN gateway.
	VpnGatewayName pulumi.StringPtrInput
}

func (GatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayState)(nil)).Elem()
}

type gatewayArgs struct {
	// The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
	Bandwidth int `pulumi:"bandwidth"`
	// The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
	// state file, not actually remove.
	BillingType *string `pulumi:"billingType"`
	// The description of the VPN gateway.
	Description *string `pulumi:"description"`
	// The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
	// Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Period *int `pulumi:"period"`
	// The project name of the VPN gateway.
	ProjectName *string `pulumi:"projectName"`
	// The ID of the subnet where you want to create the VPN gateway.
	SubnetId string `pulumi:"subnetId"`
	// Tags.
	Tags []GatewayTag `pulumi:"tags"`
	// The ID of the VPC where you want to create the VPN gateway.
	VpcId string `pulumi:"vpcId"`
	// The name of the VPN gateway.
	VpnGatewayName *string `pulumi:"vpnGatewayName"`
}

// The set of arguments for constructing a Gateway resource.
type GatewayArgs struct {
	// The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
	Bandwidth pulumi.IntInput
	// The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
	// state file, not actually remove.
	BillingType pulumi.StringPtrInput
	// The description of the VPN gateway.
	Description pulumi.StringPtrInput
	// The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
	// Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	Period pulumi.IntPtrInput
	// The project name of the VPN gateway.
	ProjectName pulumi.StringPtrInput
	// The ID of the subnet where you want to create the VPN gateway.
	SubnetId pulumi.StringInput
	// Tags.
	Tags GatewayTagArrayInput
	// The ID of the VPC where you want to create the VPN gateway.
	VpcId pulumi.StringInput
	// The name of the VPN gateway.
	VpnGatewayName pulumi.StringPtrInput
}

func (GatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gatewayArgs)(nil)).Elem()
}

type GatewayInput interface {
	pulumi.Input

	ToGatewayOutput() GatewayOutput
	ToGatewayOutputWithContext(ctx context.Context) GatewayOutput
}

func (*Gateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (i *Gateway) ToGatewayOutput() GatewayOutput {
	return i.ToGatewayOutputWithContext(context.Background())
}

func (i *Gateway) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayOutput)
}

// GatewayArrayInput is an input type that accepts GatewayArray and GatewayArrayOutput values.
// You can construct a concrete instance of `GatewayArrayInput` via:
//
//	GatewayArray{ GatewayArgs{...} }
type GatewayArrayInput interface {
	pulumi.Input

	ToGatewayArrayOutput() GatewayArrayOutput
	ToGatewayArrayOutputWithContext(context.Context) GatewayArrayOutput
}

type GatewayArray []GatewayInput

func (GatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (i GatewayArray) ToGatewayArrayOutput() GatewayArrayOutput {
	return i.ToGatewayArrayOutputWithContext(context.Background())
}

func (i GatewayArray) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayArrayOutput)
}

// GatewayMapInput is an input type that accepts GatewayMap and GatewayMapOutput values.
// You can construct a concrete instance of `GatewayMapInput` via:
//
//	GatewayMap{ "key": GatewayArgs{...} }
type GatewayMapInput interface {
	pulumi.Input

	ToGatewayMapOutput() GatewayMapOutput
	ToGatewayMapOutputWithContext(context.Context) GatewayMapOutput
}

type GatewayMap map[string]GatewayInput

func (GatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (i GatewayMap) ToGatewayMapOutput() GatewayMapOutput {
	return i.ToGatewayMapOutputWithContext(context.Background())
}

func (i GatewayMap) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GatewayMapOutput)
}

type GatewayOutput struct{ *pulumi.OutputState }

func (GatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Gateway)(nil)).Elem()
}

func (o GatewayOutput) ToGatewayOutput() GatewayOutput {
	return o
}

func (o GatewayOutput) ToGatewayOutputWithContext(ctx context.Context) GatewayOutput {
	return o
}

// The account ID of the VPN gateway.
func (o GatewayOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The bandwidth of the VPN gateway. Unit: Mbps. Values: 5, 10, 20, 50, 100, 200, 500.
func (o GatewayOutput) Bandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.Bandwidth }).(pulumi.IntOutput)
}

// The BillingType of the VPN gateway. Only support `PrePaid`. Terraform will only remove the PrePaid VPN gateway from the
// state file, not actually remove.
func (o GatewayOutput) BillingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.BillingType }).(pulumi.StringPtrOutput)
}

// The business status of the VPN gateway.
func (o GatewayOutput) BusinessStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.BusinessStatus }).(pulumi.StringOutput)
}

// The connection count of the VPN gateway.
func (o GatewayOutput) ConnectionCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.ConnectionCount }).(pulumi.IntOutput)
}

// The create time of VPN gateway.
func (o GatewayOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The deleted time of the VPN gateway.
func (o GatewayOutput) DeletedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.DeletedTime }).(pulumi.StringOutput)
}

// The description of the VPN gateway.
func (o GatewayOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The expired time of the VPN gateway.
func (o GatewayOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.ExpiredTime }).(pulumi.StringOutput)
}

// The IP address of the VPN gateway.
func (o GatewayOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The lock reason of the VPN gateway.
func (o GatewayOutput) LockReason() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.LockReason }).(pulumi.StringOutput)
}

// The Period of the VPN gateway. Default value is 12. This parameter is only useful when creating vpn gateway. Default period unit is Month.
// Value range: 1~9, 12, 24, 36. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
func (o GatewayOutput) Period() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntPtrOutput { return v.Period }).(pulumi.IntPtrOutput)
}

// The project name of the VPN gateway.
func (o GatewayOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The renew type of the VPN gateway.
func (o GatewayOutput) RenewType() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.RenewType }).(pulumi.StringOutput)
}

// The route count of the VPN gateway.
func (o GatewayOutput) RouteCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Gateway) pulumi.IntOutput { return v.RouteCount }).(pulumi.IntOutput)
}

// The status of the VPN gateway.
func (o GatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The ID of the subnet where you want to create the VPN gateway.
func (o GatewayOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Tags.
func (o GatewayOutput) Tags() GatewayTagArrayOutput {
	return o.ApplyT(func(v *Gateway) GatewayTagArrayOutput { return v.Tags }).(GatewayTagArrayOutput)
}

// The update time of VPN gateway.
func (o GatewayOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The ID of the VPC where you want to create the VPN gateway.
func (o GatewayOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The ID of the VPN gateway.
func (o GatewayOutput) VpnGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpnGatewayId }).(pulumi.StringOutput)
}

// The name of the VPN gateway.
func (o GatewayOutput) VpnGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Gateway) pulumi.StringOutput { return v.VpnGatewayName }).(pulumi.StringOutput)
}

type GatewayArrayOutput struct{ *pulumi.OutputState }

func (GatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Gateway)(nil)).Elem()
}

func (o GatewayArrayOutput) ToGatewayArrayOutput() GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) ToGatewayArrayOutputWithContext(ctx context.Context) GatewayArrayOutput {
	return o
}

func (o GatewayArrayOutput) Index(i pulumi.IntInput) GatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].([]*Gateway)[vs[1].(int)]
	}).(GatewayOutput)
}

type GatewayMapOutput struct{ *pulumi.OutputState }

func (GatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Gateway)(nil)).Elem()
}

func (o GatewayMapOutput) ToGatewayMapOutput() GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) ToGatewayMapOutputWithContext(ctx context.Context) GatewayMapOutput {
	return o
}

func (o GatewayMapOutput) MapIndex(k pulumi.StringInput) GatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Gateway {
		return vs[0].(map[string]*Gateway)[vs[1].(string)]
	}).(GatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayInput)(nil)).Elem(), &Gateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayArrayInput)(nil)).Elem(), GatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GatewayMapInput)(nil)).Elem(), GatewayMap{})
	pulumi.RegisterOutputType(GatewayOutput{})
	pulumi.RegisterOutputType(GatewayArrayOutput{})
	pulumi.RegisterOutputType(GatewayMapOutput{})
}
