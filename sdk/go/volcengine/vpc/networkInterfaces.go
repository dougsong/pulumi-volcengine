// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to query detailed information of network interfaces
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/vpc"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vpc.NetworkInterfaces(ctx, &vpc.NetworkInterfacesArgs{
//				Ids: []string{
//					"eni-2744htx2w0j5s7fap8t3ivwze",
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func NetworkInterfaces(ctx *pulumi.Context, args *NetworkInterfacesArgs, opts ...pulumi.InvokeOption) (*NetworkInterfacesResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv NetworkInterfacesResult
	err := ctx.Invoke("volcengine:vpc/networkInterfaces:NetworkInterfaces", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking NetworkInterfaces.
type NetworkInterfacesArgs struct {
	// A list of ENI ids.
	Ids []string `pulumi:"ids"`
	// An id of the instance to which the ENI is bound.
	InstanceId *string `pulumi:"instanceId"`
	// A list of network interface ids.
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// A name of ENI.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// File name where to save data source results.
	OutputFile *string `pulumi:"outputFile"`
	// A list of primary IP address of ENI.
	PrimaryIpAddresses []string `pulumi:"primaryIpAddresses"`
	// A list of private IP addresses.
	PrivateIpAddresses []string `pulumi:"privateIpAddresses"`
	// The ProjectName of the ENI.
	ProjectName *string `pulumi:"projectName"`
	// An id of the security group to which the secondary ENI belongs.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.
	Status *string `pulumi:"status"`
	// An id of the subnet to which the ENI is connected.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []NetworkInterfacesTag `pulumi:"tags"`
	// A type of ENI, Optional choice contains `primary`, `secondary`.
	Type *string `pulumi:"type"`
	// An id of the virtual private cloud (VPC) to which the ENI belongs.
	VpcId *string `pulumi:"vpcId"`
	// The zone ID.
	ZoneId *string `pulumi:"zoneId"`
}

// A collection of values returned by NetworkInterfaces.
type NetworkInterfacesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id                  string   `pulumi:"id"`
	Ids                 []string `pulumi:"ids"`
	InstanceId          *string  `pulumi:"instanceId"`
	NetworkInterfaceIds []string `pulumi:"networkInterfaceIds"`
	// The name of the ENI.
	NetworkInterfaceName *string `pulumi:"networkInterfaceName"`
	// The collection of ENI.
	NetworkInterfaces  []NetworkInterfacesNetworkInterface `pulumi:"networkInterfaces"`
	OutputFile         *string                             `pulumi:"outputFile"`
	PrimaryIpAddresses []string                            `pulumi:"primaryIpAddresses"`
	PrivateIpAddresses []string                            `pulumi:"privateIpAddresses"`
	// The ProjectName of the ENI.
	ProjectName     *string `pulumi:"projectName"`
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// The status of the ENI.
	Status *string `pulumi:"status"`
	// The id of the subnet to which the ENI is connected.
	SubnetId *string `pulumi:"subnetId"`
	// Tags.
	Tags []NetworkInterfacesTag `pulumi:"tags"`
	// The total count of ENI query.
	TotalCount int `pulumi:"totalCount"`
	// The type of the ENI.
	Type *string `pulumi:"type"`
	// The id of the virtual private cloud (VPC) to which the ENI belongs.
	VpcId *string `pulumi:"vpcId"`
	// The zone id of the ENI.
	ZoneId *string `pulumi:"zoneId"`
}

func NetworkInterfacesOutput(ctx *pulumi.Context, args NetworkInterfacesOutputArgs, opts ...pulumi.InvokeOption) NetworkInterfacesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (NetworkInterfacesResult, error) {
			args := v.(NetworkInterfacesArgs)
			r, err := NetworkInterfaces(ctx, &args, opts...)
			var s NetworkInterfacesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(NetworkInterfacesResultOutput)
}

// A collection of arguments for invoking NetworkInterfaces.
type NetworkInterfacesOutputArgs struct {
	// A list of ENI ids.
	Ids pulumi.StringArrayInput `pulumi:"ids"`
	// An id of the instance to which the ENI is bound.
	InstanceId pulumi.StringPtrInput `pulumi:"instanceId"`
	// A list of network interface ids.
	NetworkInterfaceIds pulumi.StringArrayInput `pulumi:"networkInterfaceIds"`
	// A name of ENI.
	NetworkInterfaceName pulumi.StringPtrInput `pulumi:"networkInterfaceName"`
	// File name where to save data source results.
	OutputFile pulumi.StringPtrInput `pulumi:"outputFile"`
	// A list of primary IP address of ENI.
	PrimaryIpAddresses pulumi.StringArrayInput `pulumi:"primaryIpAddresses"`
	// A list of private IP addresses.
	PrivateIpAddresses pulumi.StringArrayInput `pulumi:"privateIpAddresses"`
	// The ProjectName of the ENI.
	ProjectName pulumi.StringPtrInput `pulumi:"projectName"`
	// An id of the security group to which the secondary ENI belongs.
	SecurityGroupId pulumi.StringPtrInput `pulumi:"securityGroupId"`
	// A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.
	Status pulumi.StringPtrInput `pulumi:"status"`
	// An id of the subnet to which the ENI is connected.
	SubnetId pulumi.StringPtrInput `pulumi:"subnetId"`
	// Tags.
	Tags NetworkInterfacesTagArrayInput `pulumi:"tags"`
	// A type of ENI, Optional choice contains `primary`, `secondary`.
	Type pulumi.StringPtrInput `pulumi:"type"`
	// An id of the virtual private cloud (VPC) to which the ENI belongs.
	VpcId pulumi.StringPtrInput `pulumi:"vpcId"`
	// The zone ID.
	ZoneId pulumi.StringPtrInput `pulumi:"zoneId"`
}

func (NetworkInterfacesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfacesArgs)(nil)).Elem()
}

// A collection of values returned by NetworkInterfaces.
type NetworkInterfacesResultOutput struct{ *pulumi.OutputState }

func (NetworkInterfacesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkInterfacesResult)(nil)).Elem()
}

func (o NetworkInterfacesResultOutput) ToNetworkInterfacesResultOutput() NetworkInterfacesResultOutput {
	return o
}

func (o NetworkInterfacesResultOutput) ToNetworkInterfacesResultOutputWithContext(ctx context.Context) NetworkInterfacesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o NetworkInterfacesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o NetworkInterfacesResultOutput) Ids() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) []string { return v.Ids }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfacesResultOutput) InstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.InstanceId }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResultOutput) NetworkInterfaceIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) []string { return v.NetworkInterfaceIds }).(pulumi.StringArrayOutput)
}

// The name of the ENI.
func (o NetworkInterfacesResultOutput) NetworkInterfaceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.NetworkInterfaceName }).(pulumi.StringPtrOutput)
}

// The collection of ENI.
func (o NetworkInterfacesResultOutput) NetworkInterfaces() NetworkInterfacesNetworkInterfaceArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) []NetworkInterfacesNetworkInterface { return v.NetworkInterfaces }).(NetworkInterfacesNetworkInterfaceArrayOutput)
}

func (o NetworkInterfacesResultOutput) OutputFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.OutputFile }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResultOutput) PrimaryIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) []string { return v.PrimaryIpAddresses }).(pulumi.StringArrayOutput)
}

func (o NetworkInterfacesResultOutput) PrivateIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) []string { return v.PrivateIpAddresses }).(pulumi.StringArrayOutput)
}

// The ProjectName of the ENI.
func (o NetworkInterfacesResultOutput) ProjectName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.ProjectName }).(pulumi.StringPtrOutput)
}

func (o NetworkInterfacesResultOutput) SecurityGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.SecurityGroupId }).(pulumi.StringPtrOutput)
}

// The status of the ENI.
func (o NetworkInterfacesResultOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.Status }).(pulumi.StringPtrOutput)
}

// The id of the subnet to which the ENI is connected.
func (o NetworkInterfacesResultOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.SubnetId }).(pulumi.StringPtrOutput)
}

// Tags.
func (o NetworkInterfacesResultOutput) Tags() NetworkInterfacesTagArrayOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) []NetworkInterfacesTag { return v.Tags }).(NetworkInterfacesTagArrayOutput)
}

// The total count of ENI query.
func (o NetworkInterfacesResultOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) int { return v.TotalCount }).(pulumi.IntOutput)
}

// The type of the ENI.
func (o NetworkInterfacesResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// The id of the virtual private cloud (VPC) to which the ENI belongs.
func (o NetworkInterfacesResultOutput) VpcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.VpcId }).(pulumi.StringPtrOutput)
}

// The zone id of the ENI.
func (o NetworkInterfacesResultOutput) ZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkInterfacesResult) *string { return v.ZoneId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(NetworkInterfacesResultOutput{})
}
