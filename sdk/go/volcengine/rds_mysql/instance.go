// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package rds_mysql

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage rds mysql instance
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/volcengine/pulumi-volcengine/sdk/go/volcengine/rds_mysql"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			foo, err := rds_mysql.NewInstance(ctx, "foo", &rds_mysql.InstanceArgs{
//				DbEngineVersion:     pulumi.String("MySQL_5_7"),
//				NodeSpec:            pulumi.String("rds.mysql.1c2g"),
//				PrimaryZoneId:       pulumi.String("cn-guilin-a"),
//				SecondaryZoneId:     pulumi.String("cn-guilin-b"),
//				StorageSpace:        pulumi.Int(80),
//				SubnetId:            pulumi.String("subnet-2d72yi377stts58ozfdrlk9f6"),
//				InstanceName:        pulumi.String("tf-test"),
//				LowerCaseTableNames: pulumi.String("1"),
//				ChargeInfo: &rds_mysql.InstanceChargeInfoArgs{
//					ChargeType: pulumi.String("PostPaid"),
//				},
//				AllowListIds: pulumi.StringArray{
//					pulumi.String("acl-2dd8f8317e4d4159b21630d13ae2e6ec"),
//					pulumi.String("acl-2eaa2a053b2a4a58b988e38ae975e81c"),
//				},
//				Parameters: rds_mysql.InstanceParameterArray{
//					&rds_mysql.InstanceParameterArgs{
//						ParameterName:  pulumi.String("auto_increment_increment"),
//						ParameterValue: pulumi.String("2"),
//					},
//					&rds_mysql.InstanceParameterArgs{
//						ParameterName:  pulumi.String("auto_increment_offset"),
//						ParameterValue: pulumi.String("4"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = rds_mysql.NewInstanceReadonlyNode(ctx, "readonly", &rds_mysql.InstanceReadonlyNodeArgs{
//				InstanceId: foo.ID(),
//				NodeSpec:   pulumi.String("rds.mysql.2c4g"),
//				ZoneId:     pulumi.String("cn-guilin-a"),
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
// Rds Mysql Instance can be imported using the id, e.g.
//
// ```sh
//
//	$ pulumi import volcengine:rds_mysql/instance:Instance default mysql-72da4258c2c7
//
// ```
type Instance struct {
	pulumi.CustomResourceState

	// Allow list Ids of the RDS instance.
	AllowListIds pulumi.StringArrayOutput `pulumi:"allowListIds"`
	// The version of allow list.
	AllowListVersion pulumi.StringOutput `pulumi:"allowListVersion"`
	// The instance has used backup space. Unit: GB.
	BackupUse pulumi.IntOutput `pulumi:"backupUse"`
	// Payment methods.
	ChargeDetails InstanceChargeDetailArrayOutput `pulumi:"chargeDetails"`
	// Payment methods.
	ChargeInfo InstanceChargeInfoOutput `pulumi:"chargeInfo"`
	// Node creation local time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Data synchronization mode.
	DataSyncMode pulumi.StringOutput `pulumi:"dataSyncMode"`
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion pulumi.StringOutput `pulumi:"dbEngineVersion"`
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone pulumi.StringOutput `pulumi:"dbTimeZone"`
	// The endpoint info of the RDS instance.
	Endpoints InstanceEndpointArrayOutput `pulumi:"endpoints"`
	// Instance ID.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// The status of the RDS instance.
	InstanceStatus pulumi.StringOutput `pulumi:"instanceStatus"`
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames pulumi.StringPtrOutput `pulumi:"lowerCaseTableNames"`
	// Maintenance Window.
	MaintenanceWindows InstanceMaintenanceWindowArrayOutput `pulumi:"maintenanceWindows"`
	// Memory size in GB.
	Memory pulumi.IntOutput `pulumi:"memory"`
	// The number of nodes.
	NodeNumber pulumi.IntOutput `pulumi:"nodeNumber"`
	// The specification of primary node and secondary node.
	NodeSpec pulumi.StringOutput `pulumi:"nodeSpec"`
	// Instance node information.
	Nodes InstanceNodeArrayOutput `pulumi:"nodes"`
	// Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
	Parameters InstanceParameterArrayOutput `pulumi:"parameters"`
	// The available zone of primary node.
	PrimaryZoneId pulumi.StringOutput `pulumi:"primaryZoneId"`
	// The region of the RDS instance.
	RegionId pulumi.StringOutput `pulumi:"regionId"`
	// The available zone of secondary node.
	SecondaryZoneId pulumi.StringOutput `pulumi:"secondaryZoneId"`
	// Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
	StorageSpace pulumi.IntPtrOutput `pulumi:"storageSpace"`
	// Instance storage type.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// The instance has used storage space. Unit: GB.
	StorageUse pulumi.IntOutput `pulumi:"storageUse"`
	// Subnet ID of the RDS instance.
	SubnetId pulumi.StringOutput `pulumi:"subnetId"`
	// Time zone.
	TimeZone pulumi.StringOutput `pulumi:"timeZone"`
	// The update time of the RDS instance.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// CPU size.
	VCpu pulumi.IntOutput `pulumi:"vCpu"`
	// The vpc ID of the RDS instance.
	VpcId pulumi.StringOutput `pulumi:"vpcId"`
	// The available zone of the RDS instance.
	ZoneId pulumi.StringOutput `pulumi:"zoneId"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChargeInfo == nil {
		return nil, errors.New("invalid value for required argument 'ChargeInfo'")
	}
	if args.DbEngineVersion == nil {
		return nil, errors.New("invalid value for required argument 'DbEngineVersion'")
	}
	if args.NodeSpec == nil {
		return nil, errors.New("invalid value for required argument 'NodeSpec'")
	}
	if args.PrimaryZoneId == nil {
		return nil, errors.New("invalid value for required argument 'PrimaryZoneId'")
	}
	if args.SecondaryZoneId == nil {
		return nil, errors.New("invalid value for required argument 'SecondaryZoneId'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	var resource Instance
	err := ctx.RegisterResource("volcengine:rds_mysql/instance:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("volcengine:rds_mysql/instance:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
	// Allow list Ids of the RDS instance.
	AllowListIds []string `pulumi:"allowListIds"`
	// The version of allow list.
	AllowListVersion *string `pulumi:"allowListVersion"`
	// The instance has used backup space. Unit: GB.
	BackupUse *int `pulumi:"backupUse"`
	// Payment methods.
	ChargeDetails []InstanceChargeDetail `pulumi:"chargeDetails"`
	// Payment methods.
	ChargeInfo *InstanceChargeInfo `pulumi:"chargeInfo"`
	// Node creation local time.
	CreateTime *string `pulumi:"createTime"`
	// Data synchronization mode.
	DataSyncMode *string `pulumi:"dataSyncMode"`
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion *string `pulumi:"dbEngineVersion"`
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone *string `pulumi:"dbTimeZone"`
	// The endpoint info of the RDS instance.
	Endpoints []InstanceEndpoint `pulumi:"endpoints"`
	// Instance ID.
	InstanceId *string `pulumi:"instanceId"`
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName *string `pulumi:"instanceName"`
	// The status of the RDS instance.
	InstanceStatus *string `pulumi:"instanceStatus"`
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames *string `pulumi:"lowerCaseTableNames"`
	// Maintenance Window.
	MaintenanceWindows []InstanceMaintenanceWindow `pulumi:"maintenanceWindows"`
	// Memory size in GB.
	Memory *int `pulumi:"memory"`
	// The number of nodes.
	NodeNumber *int `pulumi:"nodeNumber"`
	// The specification of primary node and secondary node.
	NodeSpec *string `pulumi:"nodeSpec"`
	// Instance node information.
	Nodes []InstanceNode `pulumi:"nodes"`
	// Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
	Parameters []InstanceParameter `pulumi:"parameters"`
	// The available zone of primary node.
	PrimaryZoneId *string `pulumi:"primaryZoneId"`
	// The region of the RDS instance.
	RegionId *string `pulumi:"regionId"`
	// The available zone of secondary node.
	SecondaryZoneId *string `pulumi:"secondaryZoneId"`
	// Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
	StorageSpace *int `pulumi:"storageSpace"`
	// Instance storage type.
	StorageType *string `pulumi:"storageType"`
	// The instance has used storage space. Unit: GB.
	StorageUse *int `pulumi:"storageUse"`
	// Subnet ID of the RDS instance.
	SubnetId *string `pulumi:"subnetId"`
	// Time zone.
	TimeZone *string `pulumi:"timeZone"`
	// The update time of the RDS instance.
	UpdateTime *string `pulumi:"updateTime"`
	// CPU size.
	VCpu *int `pulumi:"vCpu"`
	// The vpc ID of the RDS instance.
	VpcId *string `pulumi:"vpcId"`
	// The available zone of the RDS instance.
	ZoneId *string `pulumi:"zoneId"`
}

type InstanceState struct {
	// Allow list Ids of the RDS instance.
	AllowListIds pulumi.StringArrayInput
	// The version of allow list.
	AllowListVersion pulumi.StringPtrInput
	// The instance has used backup space. Unit: GB.
	BackupUse pulumi.IntPtrInput
	// Payment methods.
	ChargeDetails InstanceChargeDetailArrayInput
	// Payment methods.
	ChargeInfo InstanceChargeInfoPtrInput
	// Node creation local time.
	CreateTime pulumi.StringPtrInput
	// Data synchronization mode.
	DataSyncMode pulumi.StringPtrInput
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion pulumi.StringPtrInput
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone pulumi.StringPtrInput
	// The endpoint info of the RDS instance.
	Endpoints InstanceEndpointArrayInput
	// Instance ID.
	InstanceId pulumi.StringPtrInput
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName pulumi.StringPtrInput
	// The status of the RDS instance.
	InstanceStatus pulumi.StringPtrInput
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames pulumi.StringPtrInput
	// Maintenance Window.
	MaintenanceWindows InstanceMaintenanceWindowArrayInput
	// Memory size in GB.
	Memory pulumi.IntPtrInput
	// The number of nodes.
	NodeNumber pulumi.IntPtrInput
	// The specification of primary node and secondary node.
	NodeSpec pulumi.StringPtrInput
	// Instance node information.
	Nodes InstanceNodeArrayInput
	// Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
	Parameters InstanceParameterArrayInput
	// The available zone of primary node.
	PrimaryZoneId pulumi.StringPtrInput
	// The region of the RDS instance.
	RegionId pulumi.StringPtrInput
	// The available zone of secondary node.
	SecondaryZoneId pulumi.StringPtrInput
	// Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
	StorageSpace pulumi.IntPtrInput
	// Instance storage type.
	StorageType pulumi.StringPtrInput
	// The instance has used storage space. Unit: GB.
	StorageUse pulumi.IntPtrInput
	// Subnet ID of the RDS instance.
	SubnetId pulumi.StringPtrInput
	// Time zone.
	TimeZone pulumi.StringPtrInput
	// The update time of the RDS instance.
	UpdateTime pulumi.StringPtrInput
	// CPU size.
	VCpu pulumi.IntPtrInput
	// The vpc ID of the RDS instance.
	VpcId pulumi.StringPtrInput
	// The available zone of the RDS instance.
	ZoneId pulumi.StringPtrInput
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Allow list Ids of the RDS instance.
	AllowListIds []string `pulumi:"allowListIds"`
	// Payment methods.
	ChargeInfo InstanceChargeInfo `pulumi:"chargeInfo"`
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion string `pulumi:"dbEngineVersion"`
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone *string `pulumi:"dbTimeZone"`
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName *string `pulumi:"instanceName"`
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames *string `pulumi:"lowerCaseTableNames"`
	// The specification of primary node and secondary node.
	NodeSpec string `pulumi:"nodeSpec"`
	// Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
	Parameters []InstanceParameter `pulumi:"parameters"`
	// The available zone of primary node.
	PrimaryZoneId string `pulumi:"primaryZoneId"`
	// The available zone of secondary node.
	SecondaryZoneId string `pulumi:"secondaryZoneId"`
	// Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
	StorageSpace *int `pulumi:"storageSpace"`
	// Subnet ID of the RDS instance.
	SubnetId string `pulumi:"subnetId"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Allow list Ids of the RDS instance.
	AllowListIds pulumi.StringArrayInput
	// Payment methods.
	ChargeInfo InstanceChargeInfoInput
	// Instance type. Value:
	// MySQL_5_7
	// MySQL_8_0.
	DbEngineVersion pulumi.StringInput
	// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
	DbTimeZone pulumi.StringPtrInput
	// Instance name. Cannot start with a number or a dash
	// Can only contain Chinese characters, letters, numbers, underscores and dashes
	// The length is limited between 1 ~ 128.
	InstanceName pulumi.StringPtrInput
	// Whether the table name is case sensitive, the default value is 1.
	// Ranges:
	// 0: Table names are stored as fixed and table names are case-sensitive.
	// 1: Table names will be stored in lowercase and table names are not case sensitive.
	LowerCaseTableNames pulumi.StringPtrInput
	// The specification of primary node and secondary node.
	NodeSpec pulumi.StringInput
	// Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
	Parameters InstanceParameterArrayInput
	// The available zone of primary node.
	PrimaryZoneId pulumi.StringInput
	// The available zone of secondary node.
	SecondaryZoneId pulumi.StringInput
	// Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
	StorageSpace pulumi.IntPtrInput
	// Subnet ID of the RDS instance.
	SubnetId pulumi.StringInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

// InstanceArrayInput is an input type that accepts InstanceArray and InstanceArrayOutput values.
// You can construct a concrete instance of `InstanceArrayInput` via:
//
//	InstanceArray{ InstanceArgs{...} }
type InstanceArrayInput interface {
	pulumi.Input

	ToInstanceArrayOutput() InstanceArrayOutput
	ToInstanceArrayOutputWithContext(context.Context) InstanceArrayOutput
}

type InstanceArray []InstanceInput

func (InstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (i InstanceArray) ToInstanceArrayOutput() InstanceArrayOutput {
	return i.ToInstanceArrayOutputWithContext(context.Background())
}

func (i InstanceArray) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceArrayOutput)
}

// InstanceMapInput is an input type that accepts InstanceMap and InstanceMapOutput values.
// You can construct a concrete instance of `InstanceMapInput` via:
//
//	InstanceMap{ "key": InstanceArgs{...} }
type InstanceMapInput interface {
	pulumi.Input

	ToInstanceMapOutput() InstanceMapOutput
	ToInstanceMapOutputWithContext(context.Context) InstanceMapOutput
}

type InstanceMap map[string]InstanceInput

func (InstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (i InstanceMap) ToInstanceMapOutput() InstanceMapOutput {
	return i.ToInstanceMapOutputWithContext(context.Background())
}

func (i InstanceMap) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceMapOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

// Allow list Ids of the RDS instance.
func (o InstanceOutput) AllowListIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.AllowListIds }).(pulumi.StringArrayOutput)
}

// The version of allow list.
func (o InstanceOutput) AllowListVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.AllowListVersion }).(pulumi.StringOutput)
}

// The instance has used backup space. Unit: GB.
func (o InstanceOutput) BackupUse() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.BackupUse }).(pulumi.IntOutput)
}

// Payment methods.
func (o InstanceOutput) ChargeDetails() InstanceChargeDetailArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceChargeDetailArrayOutput { return v.ChargeDetails }).(InstanceChargeDetailArrayOutput)
}

// Payment methods.
func (o InstanceOutput) ChargeInfo() InstanceChargeInfoOutput {
	return o.ApplyT(func(v *Instance) InstanceChargeInfoOutput { return v.ChargeInfo }).(InstanceChargeInfoOutput)
}

// Node creation local time.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Data synchronization mode.
func (o InstanceOutput) DataSyncMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DataSyncMode }).(pulumi.StringOutput)
}

// Instance type. Value:
// MySQL_5_7
// MySQL_8_0.
func (o InstanceOutput) DbEngineVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DbEngineVersion }).(pulumi.StringOutput)
}

// Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
func (o InstanceOutput) DbTimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DbTimeZone }).(pulumi.StringOutput)
}

// The endpoint info of the RDS instance.
func (o InstanceOutput) Endpoints() InstanceEndpointArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceEndpointArrayOutput { return v.Endpoints }).(InstanceEndpointArrayOutput)
}

// Instance ID.
func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// Instance name. Cannot start with a number or a dash
// Can only contain Chinese characters, letters, numbers, underscores and dashes
// The length is limited between 1 ~ 128.
func (o InstanceOutput) InstanceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.InstanceName }).(pulumi.StringPtrOutput)
}

// The status of the RDS instance.
func (o InstanceOutput) InstanceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceStatus }).(pulumi.StringOutput)
}

// Whether the table name is case sensitive, the default value is 1.
// Ranges:
// 0: Table names are stored as fixed and table names are case-sensitive.
// 1: Table names will be stored in lowercase and table names are not case sensitive.
func (o InstanceOutput) LowerCaseTableNames() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringPtrOutput { return v.LowerCaseTableNames }).(pulumi.StringPtrOutput)
}

// Maintenance Window.
func (o InstanceOutput) MaintenanceWindows() InstanceMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceMaintenanceWindowArrayOutput { return v.MaintenanceWindows }).(InstanceMaintenanceWindowArrayOutput)
}

// Memory size in GB.
func (o InstanceOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.Memory }).(pulumi.IntOutput)
}

// The number of nodes.
func (o InstanceOutput) NodeNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.NodeNumber }).(pulumi.IntOutput)
}

// The specification of primary node and secondary node.
func (o InstanceOutput) NodeSpec() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.NodeSpec }).(pulumi.StringOutput)
}

// Instance node information.
func (o InstanceOutput) Nodes() InstanceNodeArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceNodeArrayOutput { return v.Nodes }).(InstanceNodeArrayOutput)
}

// Parameter of the RDS instance. This field can only be added or modified. Deleting this field is invalid.
func (o InstanceOutput) Parameters() InstanceParameterArrayOutput {
	return o.ApplyT(func(v *Instance) InstanceParameterArrayOutput { return v.Parameters }).(InstanceParameterArrayOutput)
}

// The available zone of primary node.
func (o InstanceOutput) PrimaryZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PrimaryZoneId }).(pulumi.StringOutput)
}

// The region of the RDS instance.
func (o InstanceOutput) RegionId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.RegionId }).(pulumi.StringOutput)
}

// The available zone of secondary node.
func (o InstanceOutput) SecondaryZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SecondaryZoneId }).(pulumi.StringOutput)
}

// Instance storage space. Value range: [20, 3000], unit: GB, increments every 100GB. Default value: 100.
func (o InstanceOutput) StorageSpace() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntPtrOutput { return v.StorageSpace }).(pulumi.IntPtrOutput)
}

// Instance storage type.
func (o InstanceOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// The instance has used storage space. Unit: GB.
func (o InstanceOutput) StorageUse() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.StorageUse }).(pulumi.IntOutput)
}

// Subnet ID of the RDS instance.
func (o InstanceOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

// Time zone.
func (o InstanceOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

// The update time of the RDS instance.
func (o InstanceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// CPU size.
func (o InstanceOutput) VCpu() pulumi.IntOutput {
	return o.ApplyT(func(v *Instance) pulumi.IntOutput { return v.VCpu }).(pulumi.IntOutput)
}

// The vpc ID of the RDS instance.
func (o InstanceOutput) VpcId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.VpcId }).(pulumi.StringOutput)
}

// The available zone of the RDS instance.
func (o InstanceOutput) ZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ZoneId }).(pulumi.StringOutput)
}

type InstanceArrayOutput struct{ *pulumi.OutputState }

func (InstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Instance)(nil)).Elem()
}

func (o InstanceArrayOutput) ToInstanceArrayOutput() InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) ToInstanceArrayOutputWithContext(ctx context.Context) InstanceArrayOutput {
	return o
}

func (o InstanceArrayOutput) Index(i pulumi.IntInput) InstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].([]*Instance)[vs[1].(int)]
	}).(InstanceOutput)
}

type InstanceMapOutput struct{ *pulumi.OutputState }

func (InstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Instance)(nil)).Elem()
}

func (o InstanceMapOutput) ToInstanceMapOutput() InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) ToInstanceMapOutputWithContext(ctx context.Context) InstanceMapOutput {
	return o
}

func (o InstanceMapOutput) MapIndex(k pulumi.StringInput) InstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Instance {
		return vs[0].(map[string]*Instance)[vs[1].(string)]
	}).(InstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceArrayInput)(nil)).Elem(), InstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceMapInput)(nil)).Elem(), InstanceMap{})
	pulumi.RegisterOutputType(InstanceOutput{})
	pulumi.RegisterOutputType(InstanceArrayOutput{})
	pulumi.RegisterOutputType(InstanceMapOutput{})
}
