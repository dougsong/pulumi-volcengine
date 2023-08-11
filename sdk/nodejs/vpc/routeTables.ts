// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of route tables
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.RouteTables({
 *     ids: [
 *         "vtb-274e19skkuhog7fap8u4i8ird",
 *         "vtb-2744hslq5b7r47fap8tjomgnj",
 *     ],
 *     routeTableName: "vpc-fast",
 * });
 * ```
 */
export function routeTables(args?: RouteTablesArgs, opts?: pulumi.InvokeOptions): Promise<RouteTablesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/routeTables:RouteTables", {
        "ids": args.ids,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "routeTableName": args.routeTableName,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking RouteTables.
 */
export interface RouteTablesArgs {
    /**
     * A list of route table ids.
     */
    ids?: string[];
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The ProjectName of the route table.
     */
    projectName?: string;
    /**
     * A name of route table.
     */
    routeTableName?: string;
    /**
     * An id of VPC.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by RouteTables.
 */
export interface RouteTablesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly outputFile?: string;
    /**
     * The ProjectName of the route table.
     */
    readonly projectName?: string;
    /**
     * The name of the route table.
     */
    readonly routeTableName?: string;
    /**
     * The collection of route tables.
     */
    readonly routeTables: outputs.vpc.RouteTablesRouteTable[];
    /**
     * The total count of route table query.
     */
    readonly totalCount: number;
    /**
     * The id of the virtual private cloud (VPC) to which the route entry belongs.
     */
    readonly vpcId?: string;
}
/**
 * Use this data source to query detailed information of route tables
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.RouteTables({
 *     ids: [
 *         "vtb-274e19skkuhog7fap8u4i8ird",
 *         "vtb-2744hslq5b7r47fap8tjomgnj",
 *     ],
 *     routeTableName: "vpc-fast",
 * });
 * ```
 */
export function routeTablesOutput(args?: RouteTablesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RouteTablesResult> {
    return pulumi.output(args).apply((a: any) => routeTables(a, opts))
}

/**
 * A collection of arguments for invoking RouteTables.
 */
export interface RouteTablesOutputArgs {
    /**
     * A list of route table ids.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ProjectName of the route table.
     */
    projectName?: pulumi.Input<string>;
    /**
     * A name of route table.
     */
    routeTableName?: pulumi.Input<string>;
    /**
     * An id of VPC.
     */
    vpcId?: pulumi.Input<string>;
}
