// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ebs

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VolumesVolume struct {
	BillingType        int    `pulumi:"billingType"`
	CreatedAt          string `pulumi:"createdAt"`
	DeleteWithInstance bool   `pulumi:"deleteWithInstance"`
	Description        string `pulumi:"description"`
	DeviceName         string `pulumi:"deviceName"`
	ExpiredTime        string `pulumi:"expiredTime"`
	Id                 string `pulumi:"id"`
	ImageId            string `pulumi:"imageId"`
	// The Id of instance.
	InstanceId string `pulumi:"instanceId"`
	// The Kind of Volume.
	Kind        string `pulumi:"kind"`
	PayType     string `pulumi:"payType"`
	RenewType   int    `pulumi:"renewType"`
	Size        int    `pulumi:"size"`
	Status      string `pulumi:"status"`
	TradeStatus int    `pulumi:"tradeStatus"`
	UpdatedAt   string `pulumi:"updatedAt"`
	VolumeId    string `pulumi:"volumeId"`
	// The name of Volume.
	VolumeName string `pulumi:"volumeName"`
	// The type of Volume.
	VolumeType string `pulumi:"volumeType"`
	// The Id of Zone.
	ZoneId string `pulumi:"zoneId"`
}

// VolumesVolumeInput is an input type that accepts VolumesVolumeArgs and VolumesVolumeOutput values.
// You can construct a concrete instance of `VolumesVolumeInput` via:
//
//	VolumesVolumeArgs{...}
type VolumesVolumeInput interface {
	pulumi.Input

	ToVolumesVolumeOutput() VolumesVolumeOutput
	ToVolumesVolumeOutputWithContext(context.Context) VolumesVolumeOutput
}

type VolumesVolumeArgs struct {
	BillingType        pulumi.IntInput    `pulumi:"billingType"`
	CreatedAt          pulumi.StringInput `pulumi:"createdAt"`
	DeleteWithInstance pulumi.BoolInput   `pulumi:"deleteWithInstance"`
	Description        pulumi.StringInput `pulumi:"description"`
	DeviceName         pulumi.StringInput `pulumi:"deviceName"`
	ExpiredTime        pulumi.StringInput `pulumi:"expiredTime"`
	Id                 pulumi.StringInput `pulumi:"id"`
	ImageId            pulumi.StringInput `pulumi:"imageId"`
	// The Id of instance.
	InstanceId pulumi.StringInput `pulumi:"instanceId"`
	// The Kind of Volume.
	Kind        pulumi.StringInput `pulumi:"kind"`
	PayType     pulumi.StringInput `pulumi:"payType"`
	RenewType   pulumi.IntInput    `pulumi:"renewType"`
	Size        pulumi.IntInput    `pulumi:"size"`
	Status      pulumi.StringInput `pulumi:"status"`
	TradeStatus pulumi.IntInput    `pulumi:"tradeStatus"`
	UpdatedAt   pulumi.StringInput `pulumi:"updatedAt"`
	VolumeId    pulumi.StringInput `pulumi:"volumeId"`
	// The name of Volume.
	VolumeName pulumi.StringInput `pulumi:"volumeName"`
	// The type of Volume.
	VolumeType pulumi.StringInput `pulumi:"volumeType"`
	// The Id of Zone.
	ZoneId pulumi.StringInput `pulumi:"zoneId"`
}

func (VolumesVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumesVolume)(nil)).Elem()
}

func (i VolumesVolumeArgs) ToVolumesVolumeOutput() VolumesVolumeOutput {
	return i.ToVolumesVolumeOutputWithContext(context.Background())
}

func (i VolumesVolumeArgs) ToVolumesVolumeOutputWithContext(ctx context.Context) VolumesVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumesVolumeOutput)
}

// VolumesVolumeArrayInput is an input type that accepts VolumesVolumeArray and VolumesVolumeArrayOutput values.
// You can construct a concrete instance of `VolumesVolumeArrayInput` via:
//
//	VolumesVolumeArray{ VolumesVolumeArgs{...} }
type VolumesVolumeArrayInput interface {
	pulumi.Input

	ToVolumesVolumeArrayOutput() VolumesVolumeArrayOutput
	ToVolumesVolumeArrayOutputWithContext(context.Context) VolumesVolumeArrayOutput
}

type VolumesVolumeArray []VolumesVolumeInput

func (VolumesVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumesVolume)(nil)).Elem()
}

func (i VolumesVolumeArray) ToVolumesVolumeArrayOutput() VolumesVolumeArrayOutput {
	return i.ToVolumesVolumeArrayOutputWithContext(context.Background())
}

func (i VolumesVolumeArray) ToVolumesVolumeArrayOutputWithContext(ctx context.Context) VolumesVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumesVolumeArrayOutput)
}

type VolumesVolumeOutput struct{ *pulumi.OutputState }

func (VolumesVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumesVolume)(nil)).Elem()
}

func (o VolumesVolumeOutput) ToVolumesVolumeOutput() VolumesVolumeOutput {
	return o
}

func (o VolumesVolumeOutput) ToVolumesVolumeOutputWithContext(ctx context.Context) VolumesVolumeOutput {
	return o
}

func (o VolumesVolumeOutput) BillingType() pulumi.IntOutput {
	return o.ApplyT(func(v VolumesVolume) int { return v.BillingType }).(pulumi.IntOutput)
}

func (o VolumesVolumeOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) DeleteWithInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v VolumesVolume) bool { return v.DeleteWithInstance }).(pulumi.BoolOutput)
}

func (o VolumesVolumeOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.Description }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) DeviceName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.DeviceName }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) ExpiredTime() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.ExpiredTime }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.Id }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) ImageId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.ImageId }).(pulumi.StringOutput)
}

// The Id of instance.
func (o VolumesVolumeOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.InstanceId }).(pulumi.StringOutput)
}

// The Kind of Volume.
func (o VolumesVolumeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.Kind }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) PayType() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.PayType }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) RenewType() pulumi.IntOutput {
	return o.ApplyT(func(v VolumesVolume) int { return v.RenewType }).(pulumi.IntOutput)
}

func (o VolumesVolumeOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v VolumesVolume) int { return v.Size }).(pulumi.IntOutput)
}

func (o VolumesVolumeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.Status }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) TradeStatus() pulumi.IntOutput {
	return o.ApplyT(func(v VolumesVolume) int { return v.TradeStatus }).(pulumi.IntOutput)
}

func (o VolumesVolumeOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.UpdatedAt }).(pulumi.StringOutput)
}

func (o VolumesVolumeOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.VolumeId }).(pulumi.StringOutput)
}

// The name of Volume.
func (o VolumesVolumeOutput) VolumeName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.VolumeName }).(pulumi.StringOutput)
}

// The type of Volume.
func (o VolumesVolumeOutput) VolumeType() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.VolumeType }).(pulumi.StringOutput)
}

// The Id of Zone.
func (o VolumesVolumeOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v VolumesVolume) string { return v.ZoneId }).(pulumi.StringOutput)
}

type VolumesVolumeArrayOutput struct{ *pulumi.OutputState }

func (VolumesVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VolumesVolume)(nil)).Elem()
}

func (o VolumesVolumeArrayOutput) ToVolumesVolumeArrayOutput() VolumesVolumeArrayOutput {
	return o
}

func (o VolumesVolumeArrayOutput) ToVolumesVolumeArrayOutputWithContext(ctx context.Context) VolumesVolumeArrayOutput {
	return o
}

func (o VolumesVolumeArrayOutput) Index(i pulumi.IntInput) VolumesVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VolumesVolume {
		return vs[0].([]VolumesVolume)[vs[1].(int)]
	}).(VolumesVolumeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VolumesVolumeInput)(nil)).Elem(), VolumesVolumeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VolumesVolumeArrayInput)(nil)).Elem(), VolumesVolumeArray{})
	pulumi.RegisterOutputType(VolumesVolumeOutput{})
	pulumi.RegisterOutputType(VolumesVolumeArrayOutput{})
}
