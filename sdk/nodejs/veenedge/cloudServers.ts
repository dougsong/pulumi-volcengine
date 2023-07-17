// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of veenedge cloud servers
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultCloudServers = pulumi.output(volcengine.veenedge.CloudServers());
 * ```
 */
export function cloudServers(args?: CloudServersArgs, opts?: pulumi.InvokeOptions): Promise<CloudServersResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:veenedge/cloudServers:CloudServers", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
    }, opts);
}

/**
 * A collection of arguments for invoking CloudServers.
 */
export interface CloudServersArgs {
    /**
     * A list of cloud server IDs.
     */
    ids?: string[];
    /**
     * A Name Regex of Cloud Server.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
}

/**
 * A collection of values returned by CloudServers.
 */
export interface CloudServersResult {
    /**
     * The collection of cloud servers query.
     */
    readonly cloudServers: outputs.veenedge.CloudServersCloudServer[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of cloud servers query.
     */
    readonly totalCount: number;
}

export function cloudServersOutput(args?: CloudServersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CloudServersResult> {
    return pulumi.output(args).apply(a => cloudServers(a, opts))
}

/**
 * A collection of arguments for invoking CloudServers.
 */
export interface CloudServersOutputArgs {
    /**
     * A list of cloud server IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of Cloud Server.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
}
