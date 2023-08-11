// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances v2
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds_v2.RdsInstancesV2({});
 * ```
 */
export function rdsInstancesV2(args?: RdsInstancesV2Args, opts?: pulumi.InvokeOptions): Promise<RdsInstancesV2Result> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:rds_v2/rdsInstancesV2:RdsInstancesV2", {
        "chargeType": args.chargeType,
        "createTimeEnd": args.createTimeEnd,
        "createTimeStart": args.createTimeStart,
        "dbEngineVersion": args.dbEngineVersion,
        "instanceId": args.instanceId,
        "instanceName": args.instanceName,
        "instanceStatus": args.instanceStatus,
        "instanceType": args.instanceType,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking RdsInstancesV2.
 */
export interface RdsInstancesV2Args {
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
     * The version of the RDS instance, Value:
     * MySQL Community:
     * MySQL_5.7 - MySQL 5.7
     * MySQL_8_0 - MySQL 8.0
     * PostgreSQL Community:
     * PostgreSQL_11 - PostgreSQL 11
     * PostgreSQL_12 - PostgreSQL 12
     * Microsoft SQL Server: Not available at this time
     * SQLServer_2019 - SQL Server 2019
     * veDB for MySQL:
     * MySQL_8_0 - MySQL 8.0
     * veDB for PostgreSQL:
     * PostgreSQL_13 - PostgreSQL 13.
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
     * The status of the RDS instance, Value:
     * Running - running
     * Creating - Creating
     * Deleting - Deleting
     * Restarting - Restarting
     * Restoring - Restoring
     * Updating - changing
     * Upgrading - Upgrading
     * Error - the error.
     */
    instanceStatus?: string;
    /**
     * The type of the RDS instance, Value:
     * Value:
     * RDS for MySQL:
     * HA - high availability version;
     * RDS for PostgreSQL:
     * HA - high availability version;
     * Microsoft SQL Server: Not available at this time
     * Enterprise - Enterprise Edition
     * Standard - Standard Edition
     * Web - Web version
     * veDB for MySQL:
     * Cluster - Cluster Edition
     * veDB for PostgreSQL:
     * Cluster - Cluster Edition
     * MySQL Sharding:
     * HA - high availability version;.
     */
    instanceType?: string;
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
 * A collection of values returned by RdsInstancesV2.
 */
export interface RdsInstancesV2Result {
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
    /**
     * The type of the RDS instance.
     */
    readonly instanceType?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of RDS instance query.
     */
    readonly rdsInstances: outputs.rds_v2.RdsInstancesV2RdsInstance[];
    /**
     * The total count of RDS instance query.
     */
    readonly totalCount: number;
    /**
     * The available zone of the RDS instance.
     */
    readonly zoneId?: string;
}
/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds instances v2
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds_v2.RdsInstancesV2({});
 * ```
 */
export function rdsInstancesV2Output(args?: RdsInstancesV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RdsInstancesV2Result> {
    return pulumi.output(args).apply((a: any) => rdsInstancesV2(a, opts))
}

/**
 * A collection of arguments for invoking RdsInstancesV2.
 */
export interface RdsInstancesV2OutputArgs {
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
     * The version of the RDS instance, Value:
     * MySQL Community:
     * MySQL_5.7 - MySQL 5.7
     * MySQL_8_0 - MySQL 8.0
     * PostgreSQL Community:
     * PostgreSQL_11 - PostgreSQL 11
     * PostgreSQL_12 - PostgreSQL 12
     * Microsoft SQL Server: Not available at this time
     * SQLServer_2019 - SQL Server 2019
     * veDB for MySQL:
     * MySQL_8_0 - MySQL 8.0
     * veDB for PostgreSQL:
     * PostgreSQL_13 - PostgreSQL 13.
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
     * The status of the RDS instance, Value:
     * Running - running
     * Creating - Creating
     * Deleting - Deleting
     * Restarting - Restarting
     * Restoring - Restoring
     * Updating - changing
     * Upgrading - Upgrading
     * Error - the error.
     */
    instanceStatus?: pulumi.Input<string>;
    /**
     * The type of the RDS instance, Value:
     * Value:
     * RDS for MySQL:
     * HA - high availability version;
     * RDS for PostgreSQL:
     * HA - high availability version;
     * Microsoft SQL Server: Not available at this time
     * Enterprise - Enterprise Edition
     * Standard - Standard Edition
     * Web - Web version
     * veDB for MySQL:
     * Cluster - Cluster Edition
     * veDB for PostgreSQL:
     * Cluster - Cluster Edition
     * MySQL Sharding:
     * HA - high availability version;.
     */
    instanceType?: pulumi.Input<string>;
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
