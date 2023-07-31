// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cen service route entries
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultServiceRouteEntries = pulumi.output(volcengine.cen.ServiceRouteEntries({
 *     cenId: "cen-12ar8uclj68sg17q7y20v9gil",
 * }));
 * ```
 */
export function serviceRouteEntries(args?: ServiceRouteEntriesArgs, opts?: pulumi.InvokeOptions): Promise<ServiceRouteEntriesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:cen/serviceRouteEntries:ServiceRouteEntries", {
        "cenId": args.cenId,
        "destinationCidrBlock": args.destinationCidrBlock,
        "outputFile": args.outputFile,
        "serviceRegionId": args.serviceRegionId,
        "serviceVpcId": args.serviceVpcId,
    }, opts);
}

/**
 * A collection of arguments for invoking ServiceRouteEntries.
 */
export interface ServiceRouteEntriesArgs {
    /**
     * A cen ID.
     */
    cenId?: string;
    /**
     * A destination cidr block.
     */
    destinationCidrBlock?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A service region id.
     */
    serviceRegionId?: string;
    /**
     * A service VPC id.
     */
    serviceVpcId?: string;
}

/**
 * A collection of values returned by ServiceRouteEntries.
 */
export interface ServiceRouteEntriesResult {
    /**
     * The cen ID of the cen service route entry.
     */
    readonly cenId?: string;
    /**
     * The destination cidr block of the cen service route entry.
     */
    readonly destinationCidrBlock?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly outputFile?: string;
    /**
     * The service region id of the cen service route entry.
     */
    readonly serviceRegionId?: string;
    /**
     * The collection of cen service route entry query.
     */
    readonly serviceRouteEntries: outputs.cen.ServiceRouteEntriesServiceRouteEntry[];
    /**
     * The service VPC id of the cen service route entry.
     */
    readonly serviceVpcId?: string;
    /**
     * The total count of cen service route entry.
     */
    readonly totalCount: number;
}

export function serviceRouteEntriesOutput(args?: ServiceRouteEntriesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ServiceRouteEntriesResult> {
    return pulumi.output(args).apply(a => serviceRouteEntries(a, opts))
}

/**
 * A collection of arguments for invoking ServiceRouteEntries.
 */
export interface ServiceRouteEntriesOutputArgs {
    /**
     * A cen ID.
     */
    cenId?: pulumi.Input<string>;
    /**
     * A destination cidr block.
     */
    destinationCidrBlock?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A service region id.
     */
    serviceRegionId?: pulumi.Input<string>;
    /**
     * A service VPC id.
     */
    serviceVpcId?: pulumi.Input<string>;
}
