// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpn

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage customer gateway
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
//			_, err := vpn.NewCustomerGateway(ctx, "foo", &vpn.CustomerGatewayArgs{
//				CustomerGatewayName: pulumi.String("tf-test"),
//				Description:         pulumi.String("tf-test"),
//				IpAddress:           pulumi.String("192.0.1.3"),
//				ProjectName:         pulumi.String("default"),
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
// CustomerGateway can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:vpn/customerGateway:CustomerGateway default cgw-2byswc356dybk2dx0eed2****
//
// ```
type CustomerGateway struct {
	pulumi.CustomResourceState

	// The account ID of the customer gateway.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The connection count of the customer gateway.
	ConnectionCount pulumi.IntOutput `pulumi:"connectionCount"`
	// The create time of customer gateway.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// The ID of the customer gateway.
	CustomerGatewayId pulumi.StringOutput `pulumi:"customerGatewayId"`
	// The name of the customer gateway.
	CustomerGatewayName pulumi.StringOutput `pulumi:"customerGatewayName"`
	// The description of the customer gateway.
	Description pulumi.StringOutput `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress pulumi.StringOutput `pulumi:"ipAddress"`
	// The project name of the VPN customer gateway.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The status of the customer gateway.
	Status pulumi.StringOutput `pulumi:"status"`
	// The update time of customer gateway.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewCustomerGateway registers a new resource with the given unique name, arguments, and options.
func NewCustomerGateway(ctx *pulumi.Context,
	name string, args *CustomerGatewayArgs, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IpAddress == nil {
		return nil, errors.New("invalid value for required argument 'IpAddress'")
	}
	var resource CustomerGateway
	err := ctx.RegisterResource("volcengine:vpn/customerGateway:CustomerGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCustomerGateway gets an existing CustomerGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCustomerGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CustomerGatewayState, opts ...pulumi.ResourceOption) (*CustomerGateway, error) {
	var resource CustomerGateway
	err := ctx.ReadResource("volcengine:vpn/customerGateway:CustomerGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CustomerGateway resources.
type customerGatewayState struct {
	// The account ID of the customer gateway.
	AccountId *string `pulumi:"accountId"`
	// The connection count of the customer gateway.
	ConnectionCount *int `pulumi:"connectionCount"`
	// The create time of customer gateway.
	CreationTime *string `pulumi:"creationTime"`
	// The ID of the customer gateway.
	CustomerGatewayId *string `pulumi:"customerGatewayId"`
	// The name of the customer gateway.
	CustomerGatewayName *string `pulumi:"customerGatewayName"`
	// The description of the customer gateway.
	Description *string `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress *string `pulumi:"ipAddress"`
	// The project name of the VPN customer gateway.
	ProjectName *string `pulumi:"projectName"`
	// The status of the customer gateway.
	Status *string `pulumi:"status"`
	// The update time of customer gateway.
	UpdateTime *string `pulumi:"updateTime"`
}

type CustomerGatewayState struct {
	// The account ID of the customer gateway.
	AccountId pulumi.StringPtrInput
	// The connection count of the customer gateway.
	ConnectionCount pulumi.IntPtrInput
	// The create time of customer gateway.
	CreationTime pulumi.StringPtrInput
	// The ID of the customer gateway.
	CustomerGatewayId pulumi.StringPtrInput
	// The name of the customer gateway.
	CustomerGatewayName pulumi.StringPtrInput
	// The description of the customer gateway.
	Description pulumi.StringPtrInput
	// The IP address of the customer gateway.
	IpAddress pulumi.StringPtrInput
	// The project name of the VPN customer gateway.
	ProjectName pulumi.StringPtrInput
	// The status of the customer gateway.
	Status pulumi.StringPtrInput
	// The update time of customer gateway.
	UpdateTime pulumi.StringPtrInput
}

func (CustomerGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayState)(nil)).Elem()
}

type customerGatewayArgs struct {
	// The name of the customer gateway.
	CustomerGatewayName *string `pulumi:"customerGatewayName"`
	// The description of the customer gateway.
	Description *string `pulumi:"description"`
	// The IP address of the customer gateway.
	IpAddress string `pulumi:"ipAddress"`
	// The project name of the VPN customer gateway.
	ProjectName *string `pulumi:"projectName"`
}

// The set of arguments for constructing a CustomerGateway resource.
type CustomerGatewayArgs struct {
	// The name of the customer gateway.
	CustomerGatewayName pulumi.StringPtrInput
	// The description of the customer gateway.
	Description pulumi.StringPtrInput
	// The IP address of the customer gateway.
	IpAddress pulumi.StringInput
	// The project name of the VPN customer gateway.
	ProjectName pulumi.StringPtrInput
}

func (CustomerGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*customerGatewayArgs)(nil)).Elem()
}

type CustomerGatewayInput interface {
	pulumi.Input

	ToCustomerGatewayOutput() CustomerGatewayOutput
	ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput
}

func (*CustomerGateway) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (i *CustomerGateway) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return i.ToCustomerGatewayOutputWithContext(context.Background())
}

func (i *CustomerGateway) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayOutput)
}

// CustomerGatewayArrayInput is an input type that accepts CustomerGatewayArray and CustomerGatewayArrayOutput values.
// You can construct a concrete instance of `CustomerGatewayArrayInput` via:
//
//	CustomerGatewayArray{ CustomerGatewayArgs{...} }
type CustomerGatewayArrayInput interface {
	pulumi.Input

	ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput
	ToCustomerGatewayArrayOutputWithContext(context.Context) CustomerGatewayArrayOutput
}

type CustomerGatewayArray []CustomerGatewayInput

func (CustomerGatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return i.ToCustomerGatewayArrayOutputWithContext(context.Background())
}

func (i CustomerGatewayArray) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayArrayOutput)
}

// CustomerGatewayMapInput is an input type that accepts CustomerGatewayMap and CustomerGatewayMapOutput values.
// You can construct a concrete instance of `CustomerGatewayMapInput` via:
//
//	CustomerGatewayMap{ "key": CustomerGatewayArgs{...} }
type CustomerGatewayMapInput interface {
	pulumi.Input

	ToCustomerGatewayMapOutput() CustomerGatewayMapOutput
	ToCustomerGatewayMapOutputWithContext(context.Context) CustomerGatewayMapOutput
}

type CustomerGatewayMap map[string]CustomerGatewayInput

func (CustomerGatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return i.ToCustomerGatewayMapOutputWithContext(context.Background())
}

func (i CustomerGatewayMap) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerGatewayMapOutput)
}

type CustomerGatewayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutput() CustomerGatewayOutput {
	return o
}

func (o CustomerGatewayOutput) ToCustomerGatewayOutputWithContext(ctx context.Context) CustomerGatewayOutput {
	return o
}

// The account ID of the customer gateway.
func (o CustomerGatewayOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.AccountId }).(pulumi.StringOutput)
}

// The connection count of the customer gateway.
func (o CustomerGatewayOutput) ConnectionCount() pulumi.IntOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.IntOutput { return v.ConnectionCount }).(pulumi.IntOutput)
}

// The create time of customer gateway.
func (o CustomerGatewayOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// The ID of the customer gateway.
func (o CustomerGatewayOutput) CustomerGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.CustomerGatewayId }).(pulumi.StringOutput)
}

// The name of the customer gateway.
func (o CustomerGatewayOutput) CustomerGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.CustomerGatewayName }).(pulumi.StringOutput)
}

// The description of the customer gateway.
func (o CustomerGatewayOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The IP address of the customer gateway.
func (o CustomerGatewayOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.IpAddress }).(pulumi.StringOutput)
}

// The project name of the VPN customer gateway.
func (o CustomerGatewayOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The status of the customer gateway.
func (o CustomerGatewayOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// The update time of customer gateway.
func (o CustomerGatewayOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CustomerGateway) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type CustomerGatewayArrayOutput struct{ *pulumi.OutputState }

func (CustomerGatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutput() CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) ToCustomerGatewayArrayOutputWithContext(ctx context.Context) CustomerGatewayArrayOutput {
	return o
}

func (o CustomerGatewayArrayOutput) Index(i pulumi.IntInput) CustomerGatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].([]*CustomerGateway)[vs[1].(int)]
	}).(CustomerGatewayOutput)
}

type CustomerGatewayMapOutput struct{ *pulumi.OutputState }

func (CustomerGatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CustomerGateway)(nil)).Elem()
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutput() CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) ToCustomerGatewayMapOutputWithContext(ctx context.Context) CustomerGatewayMapOutput {
	return o
}

func (o CustomerGatewayMapOutput) MapIndex(k pulumi.StringInput) CustomerGatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CustomerGateway {
		return vs[0].(map[string]*CustomerGateway)[vs[1].(string)]
	}).(CustomerGatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayInput)(nil)).Elem(), &CustomerGateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayArrayInput)(nil)).Elem(), CustomerGatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CustomerGatewayMapInput)(nil)).Elem(), CustomerGatewayMap{})
	pulumi.RegisterOutputType(CustomerGatewayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayArrayOutput{})
	pulumi.RegisterOutputType(CustomerGatewayMapOutput{})
}
