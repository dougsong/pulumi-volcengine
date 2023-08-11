// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Provides a resource to manage rds instance v2
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 * import * as volcengine from "@volcengine/pulumi";
 *
 * const fooZones = volcengine.ecs.Zones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooRdsInstanceV2 = new volcengine.rds_v2.RdsInstanceV2("fooRdsInstanceV2", {
 *     dbEngineVersion: "MySQL_5_7",
 *     nodeInfos: [
 *         {
 *             nodeType: "Primary",
 *             nodeSpec: "rds.mysql.2c4g",
 *             zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *         },
 *         {
 *             nodeType: "Secondary",
 *             nodeSpec: "rds.mysql.2c4g",
 *             zoneId: fooZones.then(fooZones => fooZones.zones?.[0]?.id),
 *         },
 *     ],
 *     storageType: "LocalSSD",
 *     storageSpace: 100,
 *     vpcId: fooVpc.id,
 *     subnetId: fooSubnet.id,
 *     instanceName: "tf-test-v2",
 *     lowerCaseTableNames: "1",
 *     chargeInfo: {
 *         chargeType: "PostPaid",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * RDS Instance can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:rds_v2/rdsInstanceV2:RdsInstanceV2 default mysql-42b38c769c4b
 * ```
 */
export class RdsInstanceV2 extends pulumi.CustomResource {
    /**
     * Get an existing RdsInstanceV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdsInstanceV2State, opts?: pulumi.CustomResourceOptions): RdsInstanceV2 {
        return new RdsInstanceV2(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:rds_v2/rdsInstanceV2:RdsInstanceV2';

    /**
     * Returns true if the given object is an instance of RdsInstanceV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdsInstanceV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdsInstanceV2.__pulumiType;
    }

    /**
     * Payment methods.
     */
    public readonly chargeInfo!: pulumi.Output<outputs.rds_v2.RdsInstanceV2ChargeInfo>;
    /**
     * The connection info ot the RDS instance.
     */
    public /*out*/ readonly connectionInfos!: pulumi.Output<outputs.rds_v2.RdsInstanceV2ConnectionInfo[]>;
    /**
     * Instance type. Value:
     * MySQL_5_7
     * MySQL_8_0.
     */
    public readonly dbEngineVersion!: pulumi.Output<string>;
    /**
     * Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly dbParamGroupId!: pulumi.Output<string | undefined>;
    /**
     * Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly dbTimeZone!: pulumi.Output<string | undefined>;
    /**
     * Instance name. Cannot start with a number or a dash
     * Can only contain Chinese characters, letters, numbers, underscores and dashes
     * The length is limited between 1 ~ 128.
     */
    public readonly instanceName!: pulumi.Output<string | undefined>;
    /**
     * The field instanceType is no longer support. The type of Instance.
     *
     * @deprecated The field instance_type is no longer support.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * Whether the table name is case sensitive, the default value is 1.
     * Ranges:
     * 0: Table names are stored as fixed and table names are case-sensitive.
     * 1: Table names will be stored in lowercase and table names are not case sensitive.
     */
    public readonly lowerCaseTableNames!: pulumi.Output<string | undefined>;
    /**
     * Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
     */
    public readonly nodeInfos!: pulumi.Output<outputs.rds_v2.RdsInstanceV2NodeInfo[]>;
    /**
     * Subordinate to the project.
     */
    public readonly projectName!: pulumi.Output<string>;
    /**
     * Instance storage space.
     * When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
     * When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
     */
    public readonly storageSpace!: pulumi.Output<number | undefined>;
    /**
     * Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
     * LocalSSD - local SSD disk
     * When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
     * DistributedStorage - Distributed Storage.
     */
    public readonly storageType!: pulumi.Output<string>;
    /**
     * Subnet ID.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
     */
    public readonly vpcId!: pulumi.Output<string>;

    /**
     * Create a RdsInstanceV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdsInstanceV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdsInstanceV2Args | RdsInstanceV2State, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdsInstanceV2State | undefined;
            resourceInputs["chargeInfo"] = state ? state.chargeInfo : undefined;
            resourceInputs["connectionInfos"] = state ? state.connectionInfos : undefined;
            resourceInputs["dbEngineVersion"] = state ? state.dbEngineVersion : undefined;
            resourceInputs["dbParamGroupId"] = state ? state.dbParamGroupId : undefined;
            resourceInputs["dbTimeZone"] = state ? state.dbTimeZone : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["lowerCaseTableNames"] = state ? state.lowerCaseTableNames : undefined;
            resourceInputs["nodeInfos"] = state ? state.nodeInfos : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["storageSpace"] = state ? state.storageSpace : undefined;
            resourceInputs["storageType"] = state ? state.storageType : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
        } else {
            const args = argsOrState as RdsInstanceV2Args | undefined;
            if ((!args || args.chargeInfo === undefined) && !opts.urn) {
                throw new Error("Missing required property 'chargeInfo'");
            }
            if ((!args || args.dbEngineVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dbEngineVersion'");
            }
            if ((!args || args.nodeInfos === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeInfos'");
            }
            if ((!args || args.storageType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageType'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.vpcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vpcId'");
            }
            resourceInputs["chargeInfo"] = args ? args.chargeInfo : undefined;
            resourceInputs["dbEngineVersion"] = args ? args.dbEngineVersion : undefined;
            resourceInputs["dbParamGroupId"] = args ? args.dbParamGroupId : undefined;
            resourceInputs["dbTimeZone"] = args ? args.dbTimeZone : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["lowerCaseTableNames"] = args ? args.lowerCaseTableNames : undefined;
            resourceInputs["nodeInfos"] = args ? args.nodeInfos : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["storageSpace"] = args ? args.storageSpace : undefined;
            resourceInputs["storageType"] = args ? args.storageType : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["vpcId"] = args ? args.vpcId : undefined;
            resourceInputs["connectionInfos"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdsInstanceV2.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdsInstanceV2 resources.
 */
export interface RdsInstanceV2State {
    /**
     * Payment methods.
     */
    chargeInfo?: pulumi.Input<inputs.rds_v2.RdsInstanceV2ChargeInfo>;
    /**
     * The connection info ot the RDS instance.
     */
    connectionInfos?: pulumi.Input<pulumi.Input<inputs.rds_v2.RdsInstanceV2ConnectionInfo>[]>;
    /**
     * Instance type. Value:
     * MySQL_5_7
     * MySQL_8_0.
     */
    dbEngineVersion?: pulumi.Input<string>;
    /**
     * Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    dbParamGroupId?: pulumi.Input<string>;
    /**
     * Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    dbTimeZone?: pulumi.Input<string>;
    /**
     * Instance name. Cannot start with a number or a dash
     * Can only contain Chinese characters, letters, numbers, underscores and dashes
     * The length is limited between 1 ~ 128.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The field instanceType is no longer support. The type of Instance.
     *
     * @deprecated The field instance_type is no longer support.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Whether the table name is case sensitive, the default value is 1.
     * Ranges:
     * 0: Table names are stored as fixed and table names are case-sensitive.
     * 1: Table names will be stored in lowercase and table names are not case sensitive.
     */
    lowerCaseTableNames?: pulumi.Input<string>;
    /**
     * Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
     */
    nodeInfos?: pulumi.Input<pulumi.Input<inputs.rds_v2.RdsInstanceV2NodeInfo>[]>;
    /**
     * Subordinate to the project.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Instance storage space.
     * When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
     * When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
     */
    storageSpace?: pulumi.Input<number>;
    /**
     * Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
     * LocalSSD - local SSD disk
     * When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
     * DistributedStorage - Distributed Storage.
     */
    storageType?: pulumi.Input<string>;
    /**
     * Subnet ID.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
     */
    vpcId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdsInstanceV2 resource.
 */
export interface RdsInstanceV2Args {
    /**
     * Payment methods.
     */
    chargeInfo: pulumi.Input<inputs.rds_v2.RdsInstanceV2ChargeInfo>;
    /**
     * Instance type. Value:
     * MySQL_5_7
     * MySQL_8_0.
     */
    dbEngineVersion: pulumi.Input<string>;
    /**
     * Parameter template ID. It only takes effect when the database type is MySQL/PostgreSQL/SQL_Server. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    dbParamGroupId?: pulumi.Input<string>;
    /**
     * Time zone. Support UTC -12:00 ~ +13:00. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    dbTimeZone?: pulumi.Input<string>;
    /**
     * Instance name. Cannot start with a number or a dash
     * Can only contain Chinese characters, letters, numbers, underscores and dashes
     * The length is limited between 1 ~ 128.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The field instanceType is no longer support. The type of Instance.
     *
     * @deprecated The field instance_type is no longer support.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * Whether the table name is case sensitive, the default value is 1.
     * Ranges:
     * 0: Table names are stored as fixed and table names are case-sensitive.
     * 1: Table names will be stored in lowercase and table names are not case sensitive.
     */
    lowerCaseTableNames?: pulumi.Input<string>;
    /**
     * Instance specification configuration. This parameter is required for RDS for MySQL, RDS for PostgreSQL and MySQL Sharding. There is one and only one Primary node, one and only one Secondary node, and 0-10 Read-Only nodes.
     */
    nodeInfos: pulumi.Input<pulumi.Input<inputs.rds_v2.RdsInstanceV2NodeInfo>[]>;
    /**
     * Subordinate to the project.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Instance storage space.
     * When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, value range: [20, 3000], unit: GB, increments every 100GB.
     * When the database type is veDB_MySQL/veDB_PostgreSQL, this parameter does not need to be passed.
     */
    storageSpace?: pulumi.Input<number>;
    /**
     * Instance storage type. When the database type is MySQL/PostgreSQL/SQL_Server/MySQL Sharding, the value is:
     * LocalSSD - local SSD disk
     * When the database type is veDB_MySQL/veDB_PostgreSQL, the value is:
     * DistributedStorage - Distributed Storage.
     */
    storageType: pulumi.Input<string>;
    /**
     * Subnet ID.
     */
    subnetId: pulumi.Input<string>;
    /**
     * Private network (VPC) ID. You can call the DescribeVpcs query and use this parameter to specify the VPC where the instance is to be created.
     */
    vpcId: pulumi.Input<string>;
}
