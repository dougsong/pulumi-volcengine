// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package redis

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage redis backup restore
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/redis"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := redis.NewBackupRestore(ctx, "default", &redis.BackupRestoreArgs{
//				InstanceId: pulumi.String("redis-cnlfvrv4qye6u4lpa"),
//				TimePoint:  pulumi.String("2023-04-14T02:51:51Z"),
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
// Redis Backup Restore can be imported using the restore:instanceId, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:redis/backupRestore:BackupRestore default restore:redis-asdljioeixxxx
//
// ```
type BackupRestore struct {
	pulumi.CustomResourceState

	// The type of backup. The value can be Full or Inc.
	BackupType pulumi.StringPtrOutput `pulumi:"backupType"`
	// Id of instance.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignoreChanges in import.
	TimePoint pulumi.StringOutput `pulumi:"timePoint"`
}

// NewBackupRestore registers a new resource with the given unique name, arguments, and options.
func NewBackupRestore(ctx *pulumi.Context,
	name string, args *BackupRestoreArgs, opts ...pulumi.ResourceOption) (*BackupRestore, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.TimePoint == nil {
		return nil, errors.New("invalid value for required argument 'TimePoint'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource BackupRestore
	err := ctx.RegisterResource("volcengine:redis/backupRestore:BackupRestore", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBackupRestore gets an existing BackupRestore resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBackupRestore(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BackupRestoreState, opts ...pulumi.ResourceOption) (*BackupRestore, error) {
	var resource BackupRestore
	err := ctx.ReadResource("volcengine:redis/backupRestore:BackupRestore", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BackupRestore resources.
type backupRestoreState struct {
	// The type of backup. The value can be Full or Inc.
	BackupType *string `pulumi:"backupType"`
	// Id of instance.
	InstanceId *string `pulumi:"instanceId"`
	// Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignoreChanges in import.
	TimePoint *string `pulumi:"timePoint"`
}

type BackupRestoreState struct {
	// The type of backup. The value can be Full or Inc.
	BackupType pulumi.StringPtrInput
	// Id of instance.
	InstanceId pulumi.StringPtrInput
	// Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignoreChanges in import.
	TimePoint pulumi.StringPtrInput
}

func (BackupRestoreState) ElementType() reflect.Type {
	return reflect.TypeOf((*backupRestoreState)(nil)).Elem()
}

type backupRestoreArgs struct {
	// The type of backup. The value can be Full or Inc.
	BackupType *string `pulumi:"backupType"`
	// Id of instance.
	InstanceId string `pulumi:"instanceId"`
	// Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignoreChanges in import.
	TimePoint string `pulumi:"timePoint"`
}

// The set of arguments for constructing a BackupRestore resource.
type BackupRestoreArgs struct {
	// The type of backup. The value can be Full or Inc.
	BackupType pulumi.StringPtrInput
	// Id of instance.
	InstanceId pulumi.StringInput
	// Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignoreChanges in import.
	TimePoint pulumi.StringInput
}

func (BackupRestoreArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*backupRestoreArgs)(nil)).Elem()
}

type BackupRestoreInput interface {
	pulumi.Input

	ToBackupRestoreOutput() BackupRestoreOutput
	ToBackupRestoreOutputWithContext(ctx context.Context) BackupRestoreOutput
}

func (*BackupRestore) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupRestore)(nil)).Elem()
}

func (i *BackupRestore) ToBackupRestoreOutput() BackupRestoreOutput {
	return i.ToBackupRestoreOutputWithContext(context.Background())
}

func (i *BackupRestore) ToBackupRestoreOutputWithContext(ctx context.Context) BackupRestoreOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupRestoreOutput)
}

// BackupRestoreArrayInput is an input type that accepts BackupRestoreArray and BackupRestoreArrayOutput values.
// You can construct a concrete instance of `BackupRestoreArrayInput` via:
//
//	BackupRestoreArray{ BackupRestoreArgs{...} }
type BackupRestoreArrayInput interface {
	pulumi.Input

	ToBackupRestoreArrayOutput() BackupRestoreArrayOutput
	ToBackupRestoreArrayOutputWithContext(context.Context) BackupRestoreArrayOutput
}

type BackupRestoreArray []BackupRestoreInput

func (BackupRestoreArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupRestore)(nil)).Elem()
}

func (i BackupRestoreArray) ToBackupRestoreArrayOutput() BackupRestoreArrayOutput {
	return i.ToBackupRestoreArrayOutputWithContext(context.Background())
}

func (i BackupRestoreArray) ToBackupRestoreArrayOutputWithContext(ctx context.Context) BackupRestoreArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupRestoreArrayOutput)
}

// BackupRestoreMapInput is an input type that accepts BackupRestoreMap and BackupRestoreMapOutput values.
// You can construct a concrete instance of `BackupRestoreMapInput` via:
//
//	BackupRestoreMap{ "key": BackupRestoreArgs{...} }
type BackupRestoreMapInput interface {
	pulumi.Input

	ToBackupRestoreMapOutput() BackupRestoreMapOutput
	ToBackupRestoreMapOutputWithContext(context.Context) BackupRestoreMapOutput
}

type BackupRestoreMap map[string]BackupRestoreInput

func (BackupRestoreMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupRestore)(nil)).Elem()
}

func (i BackupRestoreMap) ToBackupRestoreMapOutput() BackupRestoreMapOutput {
	return i.ToBackupRestoreMapOutputWithContext(context.Background())
}

func (i BackupRestoreMap) ToBackupRestoreMapOutputWithContext(ctx context.Context) BackupRestoreMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupRestoreMapOutput)
}

type BackupRestoreOutput struct{ *pulumi.OutputState }

func (BackupRestoreOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupRestore)(nil)).Elem()
}

func (o BackupRestoreOutput) ToBackupRestoreOutput() BackupRestoreOutput {
	return o
}

func (o BackupRestoreOutput) ToBackupRestoreOutputWithContext(ctx context.Context) BackupRestoreOutput {
	return o
}

// The type of backup. The value can be Full or Inc.
func (o BackupRestoreOutput) BackupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupRestore) pulumi.StringPtrOutput { return v.BackupType }).(pulumi.StringPtrOutput)
}

// Id of instance.
func (o BackupRestoreOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRestore) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignoreChanges in import.
func (o BackupRestoreOutput) TimePoint() pulumi.StringOutput {
	return o.ApplyT(func(v *BackupRestore) pulumi.StringOutput { return v.TimePoint }).(pulumi.StringOutput)
}

type BackupRestoreArrayOutput struct{ *pulumi.OutputState }

func (BackupRestoreArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BackupRestore)(nil)).Elem()
}

func (o BackupRestoreArrayOutput) ToBackupRestoreArrayOutput() BackupRestoreArrayOutput {
	return o
}

func (o BackupRestoreArrayOutput) ToBackupRestoreArrayOutputWithContext(ctx context.Context) BackupRestoreArrayOutput {
	return o
}

func (o BackupRestoreArrayOutput) Index(i pulumi.IntInput) BackupRestoreOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BackupRestore {
		return vs[0].([]*BackupRestore)[vs[1].(int)]
	}).(BackupRestoreOutput)
}

type BackupRestoreMapOutput struct{ *pulumi.OutputState }

func (BackupRestoreMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BackupRestore)(nil)).Elem()
}

func (o BackupRestoreMapOutput) ToBackupRestoreMapOutput() BackupRestoreMapOutput {
	return o
}

func (o BackupRestoreMapOutput) ToBackupRestoreMapOutputWithContext(ctx context.Context) BackupRestoreMapOutput {
	return o
}

func (o BackupRestoreMapOutput) MapIndex(k pulumi.StringInput) BackupRestoreOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BackupRestore {
		return vs[0].(map[string]*BackupRestore)[vs[1].(string)]
	}).(BackupRestoreOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BackupRestoreInput)(nil)).Elem(), &BackupRestore{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupRestoreArrayInput)(nil)).Elem(), BackupRestoreArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BackupRestoreMapInput)(nil)).Elem(), BackupRestoreMap{})
	pulumi.RegisterOutputType(BackupRestoreOutput{})
	pulumi.RegisterOutputType(BackupRestoreArrayOutput{})
	pulumi.RegisterOutputType(BackupRestoreMapOutput{})
}
