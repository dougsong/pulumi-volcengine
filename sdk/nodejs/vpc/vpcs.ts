// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of vpcs
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultVpcs = pulumi.output(volcengine.vpc.Vpcs({
 *     ids: ["vpc-mizl7m1kqccg5smt1bdpijuj"],
 * }));
 * ```
 */
export function vpcs(args?: VpcsArgs, opts?: pulumi.InvokeOptions): Promise<VpcsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:vpc/vpcs:Vpcs", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "tags": args.tags,
        "vpcName": args.vpcName,
    }, opts);
}

/**
 * A collection of arguments for invoking Vpcs.
 */
export interface VpcsArgs {
    /**
     * A list of VPC IDs.
     */
    ids?: string[];
    /**
     * A Name Regex of Vpc.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The ProjectName of the VPC.
     */
    projectName?: string;
    /**
     * Tags.
     */
    tags?: inputs.vpc.VpcsTag[];
    /**
     * The vpc name to query.
     */
    vpcName?: string;
}

/**
 * A collection of values returned by Vpcs.
 */
export interface VpcsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The ProjectName of the VPC.
     */
    readonly projectName?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.vpc.VpcsTag[];
    /**
     * The total count of Vpc query.
     */
    readonly totalCount: number;
    /**
     * The name of VPC.
     */
    readonly vpcName?: string;
    /**
     * The collection of Vpc query.
     */
    readonly vpcs: outputs.vpc.VpcsVpc[];
}

export function vpcsOutput(args?: VpcsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<VpcsResult> {
    return pulumi.output(args).apply(a => vpcs(a, opts))
}

/**
 * A collection of arguments for invoking Vpcs.
 */
export interface VpcsOutputArgs {
    /**
     * A list of VPC IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A Name Regex of Vpc.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ProjectName of the VPC.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.vpc.VpcsTagArgs>[]>;
    /**
     * The vpc name to query.
     */
    vpcName?: pulumi.Input<string>;
}
