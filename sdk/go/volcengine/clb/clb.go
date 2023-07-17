// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package clb

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
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/clb"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := clb.NewClb(ctx, "foo", &clb.ClbArgs{
//				Description:      pulumi.String("Demo"),
//				LoadBalancerName: pulumi.String("terraform-auto-create"),
//				LoadBalancerSpec: pulumi.String("small_1"),
//				ProjectName:      pulumi.String("yyy"),
//				SubnetId:         pulumi.String("subnet-mj92ij84m5fk5smt1arvwrtw"),
//				Type:             pulumi.String("public"),
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
// CLB can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:clb/clb:Clb default clb-273y2ok6ets007fap8txvf6us
//
// ```
type Clb struct {
	pulumi.CustomResourceState

	// The description of the CLB.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The eni address of the CLB.
	EniAddress pulumi.StringOutput `pulumi:"eniAddress"`
	// The billing type of the CLB, the value can be `PostPaid`.
	LoadBalancerBillingType pulumi.StringOutput `pulumi:"loadBalancerBillingType"`
	// The name of the CLB.
	LoadBalancerName pulumi.StringOutput `pulumi:"loadBalancerName"`
	// The specification of the CLB, the value can be `small1`, `small2`, `medium1`, `medium2`, `large1`, `large2`.
	LoadBalancerSpec pulumi.StringOutput `pulumi:"loadBalancerSpec"`
	// The master zone ID of the CLB.
	MasterZoneId pulumi.StringOutput `pulumi:"masterZoneId"`
	// The reason of the console modification protection.
	ModificationProtectionReason pulumi.StringPtrOutput `pulumi:"modificationProtectionReason"`
	// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
	ModificationProtectionStatus pulumi.StringPtrOutput `pulumi:"modificationProtectionStatus"`
	// The ProjectName of the CLB.
	ProjectName pulumi.StringPtrOutput `pulumi:"projectName"`
	// The region of the request.
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// The slave zone ID of the CLB.
	SlaveZoneId pulumi.StringOutput `pulumi:"slaveZoneId"`
	// The id of the Subnet.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Tags.
	Tags ClbTagArrayOutput `pulumi:"tags"`
	// The type of the CLB. And optional choice contains `public` or `private`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The id of the VPC.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
}

// NewClb registers a new resource with the given unique name, arguments, and options.
func NewClb(ctx *pulumi.Context,
	name string, args *ClbArgs, opts ...pulumi.ResourceOption) (*Clb, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LoadBalancerSpec == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancerSpec'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Clb
	err := ctx.RegisterResource("volcengine:clb/clb:Clb", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetClb gets an existing Clb resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetClb(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ClbState, opts ...pulumi.ResourceOption) (*Clb, error) {
	var resource Clb
	err := ctx.ReadResource("volcengine:clb/clb:Clb", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Clb resources.
type clbState struct {
	// The description of the CLB.
	Description *string `pulumi:"description"`
	// The eni address of the CLB.
	EniAddress *string `pulumi:"eniAddress"`
	// The billing type of the CLB, the value can be `PostPaid`.
	LoadBalancerBillingType *string `pulumi:"loadBalancerBillingType"`
	// The name of the CLB.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The specification of the CLB, the value can be `small1`, `small2`, `medium1`, `medium2`, `large1`, `large2`.
	LoadBalancerSpec *string `pulumi:"loadBalancerSpec"`
	// The master zone ID of the CLB.
	MasterZoneId *string `pulumi:"masterZoneId"`
	// The reason of the console modification protection.
	ModificationProtectionReason *string `pulumi:"modificationProtectionReason"`
	// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
	ModificationProtectionStatus *string `pulumi:"modificationProtectionStatus"`
	// The ProjectName of the CLB.
	ProjectName *string `pulumi:"projectName"`
	// The region of the request.
	RegionId *string `pulumi:"regionId"`
	// The slave zone ID of the CLB.
	SlaveZoneId *string `pulumi:"slaveZoneId"`
	// The id of the Subnet.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []ClbTag `pulumi:"tags"`
	// The type of the CLB. And optional choice contains `public` or `private`.
	Type *string `pulumi:"type"`
	// The id of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

type ClbState struct {
	// The description of the CLB.
	Description pulumi.StringPtrInput
	// The eni address of the CLB.
	EniAddress pulumi.StringPtrInput
	// The billing type of the CLB, the value can be `PostPaid`.
	LoadBalancerBillingType pulumi.StringPtrInput
	// The name of the CLB.
	LoadBalancerName pulumi.StringPtrInput
	// The specification of the CLB, the value can be `small1`, `small2`, `medium1`, `medium2`, `large1`, `large2`.
	LoadBalancerSpec pulumi.StringPtrInput
	// The master zone ID of the CLB.
	MasterZoneId pulumi.StringPtrInput
	// The reason of the console modification protection.
	ModificationProtectionReason pulumi.StringPtrInput
	// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
	ModificationProtectionStatus pulumi.StringPtrInput
	// The ProjectName of the CLB.
	ProjectName pulumi.StringPtrInput
	// The region of the request.
	RegionId pulumi.StringPtrInput
	// The slave zone ID of the CLB.
	SlaveZoneId pulumi.StringPtrInput
	// The id of the Subnet.
	SubnetId pulumi.StringPtrInput
	// Tags.
	Tags ClbTagArrayInput
	// The type of the CLB. And optional choice contains `public` or `private`.
	Type pulumi.StringPtrInput
	// The id of the VPC.
	VpcId pulumi.StringPtrInput
}

func (ClbState) ElementType() reflect.Type {
	return reflect.TypeOf((*clbState)(nil)).Elem()
}

type clbArgs struct {
	// The description of the CLB.
	Description *string `pulumi:"description"`
	// The eni address of the CLB.
	EniAddress *string `pulumi:"eniAddress"`
	// The billing type of the CLB, the value can be `PostPaid`.
	LoadBalancerBillingType *string `pulumi:"loadBalancerBillingType"`
	// The name of the CLB.
	LoadBalancerName *string `pulumi:"loadBalancerName"`
	// The specification of the CLB, the value can be `small1`, `small2`, `medium1`, `medium2`, `large1`, `large2`.
	LoadBalancerSpec string `pulumi:"loadBalancerSpec"`
	// The master zone ID of the CLB.
	MasterZoneId *string `pulumi:"masterZoneId"`
	// The reason of the console modification protection.
	ModificationProtectionReason *string `pulumi:"modificationProtectionReason"`
	// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
	ModificationProtectionStatus *string `pulumi:"modificationProtectionStatus"`
	// The ProjectName of the CLB.
	ProjectName *string `pulumi:"projectName"`
	// The region of the request.
	RegionId *string `pulumi:"regionId"`
	// The slave zone ID of the CLB.
	SlaveZoneId *string `pulumi:"slaveZoneId"`
	// The id of the Subnet.
	SubnetId string `pulumi:"subnetId"`
	// Tags.
	Tags []ClbTag `pulumi:"tags"`
	// The type of the CLB. And optional choice contains `public` or `private`.
	Type string `pulumi:"type"`
	// The id of the VPC.
	VpcId *string `pulumi:"vpcId"`
}

// The set of arguments for constructing a Clb resource.
type ClbArgs struct {
	// The description of the CLB.
	Description pulumi.StringPtrInput
	// The eni address of the CLB.
	EniAddress pulumi.StringPtrInput
	// The billing type of the CLB, the value can be `PostPaid`.
	LoadBalancerBillingType pulumi.StringPtrInput
	// The name of the CLB.
	LoadBalancerName pulumi.StringPtrInput
	// The specification of the CLB, the value can be `small1`, `small2`, `medium1`, `medium2`, `large1`, `large2`.
	LoadBalancerSpec pulumi.StringInput
	// The master zone ID of the CLB.
	MasterZoneId pulumi.StringPtrInput
	// The reason of the console modification protection.
	ModificationProtectionReason pulumi.StringPtrInput
	// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
	ModificationProtectionStatus pulumi.StringPtrInput
	// The ProjectName of the CLB.
	ProjectName pulumi.StringPtrInput
	// The region of the request.
	RegionId pulumi.StringPtrInput
	// The slave zone ID of the CLB.
	SlaveZoneId pulumi.StringPtrInput
	// The id of the Subnet.
	SubnetId pulumi.StringInput
	// Tags.
	Tags ClbTagArrayInput
	// The type of the CLB. And optional choice contains `public` or `private`.
	Type pulumi.StringInput
	// The id of the VPC.
	VpcId pulumi.StringPtrInput
}

func (ClbArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clbArgs)(nil)).Elem()
}

type ClbInput interface {
	pulumi.Input

	ToClbOutput() ClbOutput
	ToClbOutputWithContext(ctx context.Context) ClbOutput
}

func (*Clb) ElementType() reflect.Type {
	return reflect.TypeOf((**Clb)(nil)).Elem()
}

func (i *Clb) ToClbOutput() ClbOutput {
	return i.ToClbOutputWithContext(context.Background())
}

func (i *Clb) ToClbOutputWithContext(ctx context.Context) ClbOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClbOutput)
}

// ClbArrayInput is an input type that accepts ClbArray and ClbArrayOutput values.
// You can construct a concrete instance of `ClbArrayInput` via:
//
//	ClbArray{ ClbArgs{...} }
type ClbArrayInput interface {
	pulumi.Input

	ToClbArrayOutput() ClbArrayOutput
	ToClbArrayOutputWithContext(context.Context) ClbArrayOutput
}

type ClbArray []ClbInput

func (ClbArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Clb)(nil)).Elem()
}

func (i ClbArray) ToClbArrayOutput() ClbArrayOutput {
	return i.ToClbArrayOutputWithContext(context.Background())
}

func (i ClbArray) ToClbArrayOutputWithContext(ctx context.Context) ClbArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClbArrayOutput)
}

// ClbMapInput is an input type that accepts ClbMap and ClbMapOutput values.
// You can construct a concrete instance of `ClbMapInput` via:
//
//	ClbMap{ "key": ClbArgs{...} }
type ClbMapInput interface {
	pulumi.Input

	ToClbMapOutput() ClbMapOutput
	ToClbMapOutputWithContext(context.Context) ClbMapOutput
}

type ClbMap map[string]ClbInput

func (ClbMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Clb)(nil)).Elem()
}

func (i ClbMap) ToClbMapOutput() ClbMapOutput {
	return i.ToClbMapOutputWithContext(context.Background())
}

func (i ClbMap) ToClbMapOutputWithContext(ctx context.Context) ClbMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClbMapOutput)
}

type ClbOutput struct{ *pulumi.OutputState }

func (ClbOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Clb)(nil)).Elem()
}

func (o ClbOutput) ToClbOutput() ClbOutput {
	return o
}

func (o ClbOutput) ToClbOutputWithContext(ctx context.Context) ClbOutput {
	return o
}

// The description of the CLB.
func (o ClbOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The eni address of the CLB.
func (o ClbOutput) EniAddress() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.EniAddress }).(pulumi.StringOutput)
}

// The billing type of the CLB, the value can be `PostPaid`.
func (o ClbOutput) LoadBalancerBillingType() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.LoadBalancerBillingType }).(pulumi.StringOutput)
}

// The name of the CLB.
func (o ClbOutput) LoadBalancerName() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.LoadBalancerName }).(pulumi.StringOutput)
}

// The specification of the CLB, the value can be `small1`, `small2`, `medium1`, `medium2`, `large1`, `large2`.
func (o ClbOutput) LoadBalancerSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.LoadBalancerSpec }).(pulumi.StringOutput)
}

// The master zone ID of the CLB.
func (o ClbOutput) MasterZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.MasterZoneId }).(pulumi.StringOutput)
}

// The reason of the console modification protection.
func (o ClbOutput) ModificationProtectionReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringPtrOutput { return v.ModificationProtectionReason }).(pulumi.StringPtrOutput)
}

// The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.
func (o ClbOutput) ModificationProtectionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringPtrOutput { return v.ModificationProtectionStatus }).(pulumi.StringPtrOutput)
}

// The ProjectName of the CLB.
func (o ClbOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringPtrOutput { return v.ProjectName }).(pulumi.StringPtrOutput)
}

// The region of the request.
func (o ClbOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// The slave zone ID of the CLB.
func (o ClbOutput) SlaveZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.SlaveZoneId }).(pulumi.StringOutput)
}

// The id of the Subnet.
func (o ClbOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Tags.
func (o ClbOutput) Tags() ClbTagArrayOutput {
	return o.ApplyT(func(v *Clb) ClbTagArrayOutput { return v.Tags }).(ClbTagArrayOutput)
}

// The type of the CLB. And optional choice contains `public` or `private`.
func (o ClbOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The id of the VPC.
func (o ClbOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Clb) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

type ClbArrayOutput struct{ *pulumi.OutputState }

func (ClbArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Clb)(nil)).Elem()
}

func (o ClbArrayOutput) ToClbArrayOutput() ClbArrayOutput {
	return o
}

func (o ClbArrayOutput) ToClbArrayOutputWithContext(ctx context.Context) ClbArrayOutput {
	return o
}

func (o ClbArrayOutput) Index(i pulumi.IntInput) ClbOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Clb {
		return vs[0].([]*Clb)[vs[1].(int)]
	}).(ClbOutput)
}

type ClbMapOutput struct{ *pulumi.OutputState }

func (ClbMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Clb)(nil)).Elem()
}

func (o ClbMapOutput) ToClbMapOutput() ClbMapOutput {
	return o
}

func (o ClbMapOutput) ToClbMapOutputWithContext(ctx context.Context) ClbMapOutput {
	return o
}

func (o ClbMapOutput) MapIndex(k pulumi.StringInput) ClbOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Clb {
		return vs[0].(map[string]*Clb)[vs[1].(string)]
	}).(ClbOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClbInput)(nil)).Elem(), &Clb{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClbArrayInput)(nil)).Elem(), ClbArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClbMapInput)(nil)).Elem(), ClbMap{})
	pulumi.RegisterOutputType(ClbOutput{})
	pulumi.RegisterOutputType(ClbArrayOutput{})
	pulumi.RegisterOutputType(ClbMapOutput{})
}
