// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds.IpLists({
 *     instanceId: "mysql-0fdd3bab2e7c",
 * });
 * ```
 */
export function ipLists(args: IpListsArgs, opts?: pulumi.InvokeOptions): Promise<IpListsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:rds/ipLists:IpLists", {
        "instanceId": args.instanceId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking IpLists.
 */
export interface IpListsArgs {
    /**
     * The id of the RDS instance.
     */
    instanceId: string;
    /**
     * A Name Regex of RDS ip list.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by IpLists.
 */
export interface IpListsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instanceId: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The collection of RDS ip list account query.
     */
    readonly rdsIpLists: outputs.rds.IpListsRdsIpList[];
    /**
     * The total count of RDS ip list query.
     */
    readonly totalCount: number;
}
/**
 * (Deprecated! Recommend use volcengine_rds_mysql_*** replace) Use this data source to query detailed information of rds ip lists
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.rds.IpLists({
 *     instanceId: "mysql-0fdd3bab2e7c",
 * });
 * ```
 */
export function ipListsOutput(args: IpListsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<IpListsResult> {
    return pulumi.output(args).apply((a: any) => ipLists(a, opts))
}

/**
 * A collection of arguments for invoking IpLists.
 */
export interface IpListsOutputArgs {
    /**
     * The id of the RDS instance.
     */
    instanceId: pulumi.Input<string>;
    /**
     * A Name Regex of RDS ip list.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
