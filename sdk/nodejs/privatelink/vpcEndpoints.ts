// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of privatelink vpc endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.privatelink.VpcEndpoints({
 *     ids: ["ep-3rel74u229dz45zsk2i6l****"],
 * });
 * ```
 */
export function vpcEndpoints(args?: VpcEndpointsArgs, opts?: pulumi.InvokeOptions): Promise<VpcEndpointsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:privatelink/vpcEndpoints:VpcEndpoints", {
        "endpointName": args.endpointName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serviceName": args.serviceName,
        "status": args.status,
        "vpcId": args.vpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking VpcEndpoints.
 */
export interface VpcEndpointsArgs {
    /**
     * The name of vpc endpoint.
     */
    endpointName?: string;
    /**
     * The IDs of vpc endpoint.
     */
    ids?: string[];
    /**
     * A Name Regex of vpc endpoint.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of vpc endpoint service.
     */
    serviceName?: string;
    /**
     * The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
     */
    status?: string;
    /**
     * The vpc id of vpc endpoint.
     */
    vpcId?: string;
}

/**
 * A collection of values returned by VpcEndpoints.
 */
export interface VpcEndpointsResult {
    /**
     * The name of vpc endpoint.
     */
    readonly endpointName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The name of vpc endpoint service.
     */
    readonly serviceName?: string;
    /**
     * The status of vpc endpoint.
     */
    readonly status?: string;
    /**
     * Returns the total amount of the data list.
     */
    readonly totalCount: number;
    /**
     * The collection of query.
     */
    readonly vpcEndpoints: outputs.privatelink.VpcEndpointsVpcEndpoint[];
    /**
     * The vpc id of vpc endpoint.
     */
    readonly vpcId?: string;
}
/**
 * Use this data source to query detailed information of privatelink vpc endpoints
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.privatelink.VpcEndpoints({
 *     ids: ["ep-3rel74u229dz45zsk2i6l****"],
 * });
 * ```
 */
export function vpcEndpointsOutput(args?: VpcEndpointsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VpcEndpointsResult> {
    return pulumi.output(args).apply((a: any) => vpcEndpoints(a, opts))
}

/**
 * A collection of arguments for invoking VpcEndpoints.
 */
export interface VpcEndpointsOutputArgs {
    /**
     * The name of vpc endpoint.
     */
    endpointName?: pulumi.Input<string>;
    /**
     * The IDs of vpc endpoint.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of vpc endpoint.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of vpc endpoint service.
     */
    serviceName?: pulumi.Input<string>;
    /**
     * The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
     */
    status?: pulumi.Input<string>;
    /**
     * The vpc id of vpc endpoint.
     */
    vpcId?: pulumi.Input<string>;
}
