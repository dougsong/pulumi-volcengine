// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of rds mysql instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultInstances = pulumi.output(volcengine.rds_mysql.Instances({
 *     instanceId: "mysql-72da4258c2c7",
 * }));
 * ```
 */
export function instances(args?: InstancesArgs, opts?: pulumi.InvokeOptions): Promise<InstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:rds_mysql/instances:Instances", {
        "chargeType": args.chargeType,
        "createTimeEnd": args.createTimeEnd,
        "createTimeStart": args.createTimeStart,
        "dbEngineVersion": args.dbEngineVersion,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "instanceStatus": args.instanceStatus,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesArgs {
    /**
     * The charge type of the RDS instance.
     */
    chargeType?: string;
    /**
     * The end time of creating RDS instance.
     */
    createTimeEnd?: string;
    /**
     * The start time of creating RDS instance.
     */
    createTimeStart?: string;
    /**
     * The version of the RDS instance.
     */
    dbEngineVersion?: string;
    /**
     * The id of the RDS instance.
     */
    instanceId?: string;
    /**
     * The name of the RDS instance.
     */
    instanceName?: string;
    /**
     * The status of the RDS instance.
     */
    instanceStatus?: string;
    /**
     * A Name Regex of RDS instance.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The available zone of the RDS instance.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by Instances.
 */
export interface InstancesResult {
    /**
     * Payment type. Value:
     * PostPaid - Pay-As-You-Go
     * PrePaid - Yearly and monthly (default).
     */
    readonly chargeType?: string;
    readonly createTimeEnd?: string;
    readonly createTimeStart?: string;
    /**
     * The engine version of the RDS instance.
     */
    readonly dbEngineVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * Instance ID.
     */
    readonly instanceId?: string;
    /**
     * The name of the RDS instance.
     */
    readonly instanceName?: string;
    /**
     * The status of the RDS instance.
     */
    readonly instanceStatus?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of RDS instance query.
     */
    readonly rdsMysqlInstances: outputs.rds_mysql.InstancesRdsMysqlInstance[];
    /**
     * The total count of RDS instance query.
     */
    readonly totalCount: number;
    /**
     * The available zone of the RDS instance.
     */
    readonly zoneId?: string;
}

export function instancesOutput(args?: InstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<InstancesResult> {
    return pulumi.output(args).apply(a => instances(a, opts))
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesOutputArgs {
    /**
     * The charge type of the RDS instance.
     */
    chargeType?: pulumi.Input<string>;
    /**
     * The end time of creating RDS instance.
     */
    createTimeEnd?: pulumi.Input<string>;
    /**
     * The start time of creating RDS instance.
     */
    createTimeStart?: pulumi.Input<string>;
    /**
     * The version of the RDS instance.
     */
    dbEngineVersion?: pulumi.Input<string>;
    /**
     * The id of the RDS instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of the RDS instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The status of the RDS instance.
     */
    instanceStatus?: pulumi.Input<string>;
    /**
     * A Name Regex of RDS instance.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The available zone of the RDS instance.
     */
    zoneId?: pulumi.Input<string>;
}
