// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of subnets
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.Subnets({
 *     ids: ["subnet-274zsa5kfmj287fap8soo5e19"],
 * });
 * ```
 */
export function subnets(args?: SubnetsArgs, opts?: pulumi.InvokeOptions): Promise<SubnetsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:vpc/subnets:Subnets", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "routeTableId": args.routeTableId,
        "subnetName": args.subnetName,
        "vpcId": args.vpcId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking Subnets.
 */
export interface SubnetsArgs {
    /**
     * A list of Subnet IDs.
     */
    ids?: string[];
    /**
     * A Name Regex of Subnet.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The ID of route table which subnet associated with.
     */
    routeTableId?: string;
    /**
     * The subnet name to query.
     */
    subnetName?: string;
    /**
     * The ID of VPC which subnet belongs to.
     */
    vpcId?: string;
    /**
     * The ID of zone which subnet belongs to.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by Subnets.
 */
export interface SubnetsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The route table ID.
     */
    readonly routeTableId?: string;
    /**
     * The Name of Subnet.
     */
    readonly subnetName?: string;
    /**
     * The collection of Subnet query.
     */
    readonly subnets: outputs.vpc.SubnetsSubnet[];
    /**
     * The total count of Subnet query.
     */
    readonly totalCount: number;
    /**
     * The Vpc ID of Subnet.
     */
    readonly vpcId?: string;
    /**
     * The ID of Zone.
     */
    readonly zoneId?: string;
}
/**
 * Use this data source to query detailed information of subnets
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.vpc.Subnets({
 *     ids: ["subnet-274zsa5kfmj287fap8soo5e19"],
 * });
 * ```
 */
export function subnetsOutput(args?: SubnetsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<SubnetsResult> {
    return pulumi.output(args).apply((a: any) => subnets(a, opts))
}

/**
 * A collection of arguments for invoking Subnets.
 */
export interface SubnetsOutputArgs {
    /**
     * A list of Subnet IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of Subnet.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ID of route table which subnet associated with.
     */
    routeTableId?: pulumi.Input<string>;
    /**
     * The subnet name to query.
     */
    subnetName?: pulumi.Input<string>;
    /**
     * The ID of VPC which subnet belongs to.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The ID of zone which subnet belongs to.
     */
    zoneId?: pulumi.Input<string>;
}
