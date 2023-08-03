// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of server groups
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@volcengine/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const fooZones = volcengine.ecs.Zones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooClb = new volcengine.clb.Clb("fooClb", {
 *     type: "public",
 *     subnetId: fooSubnet.id,
 *     loadBalancerSpec: "small_1",
 *     description: "acc0Demo",
 *     loadBalancerName: "acc-test-create",
 *     eipBillingConfig: {
 *         isp: "BGP",
 *         eipBillingType: "PostPaidByBandwidth",
 *         bandwidth: 1,
 *     },
 * });
 * const fooServerGroup = new volcengine.clb.ServerGroup("fooServerGroup", {
 *     loadBalancerId: fooClb.id,
 *     serverGroupName: "acc-test-create",
 *     description: "hello demo11",
 * });
 * const fooServerGroups = volcengine.clb.ServerGroupsOutput({
 *     ids: [fooServerGroup.id],
 * });
 * ```
 */
export function serverGroups(args?: ServerGroupsArgs, opts?: pulumi.InvokeOptions): Promise<ServerGroupsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:clb/serverGroups:ServerGroups", {
        "ids": args.ids,
        "loadBalancerId": args.loadBalancerId,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "serverGroupName": args.serverGroupName,
    }, opts);
}

/**
 * A collection of arguments for invoking ServerGroups.
 */
export interface ServerGroupsArgs {
    /**
     * A list of ServerGroup IDs.
     */
    ids?: string[];
    /**
     * The id of the Clb.
     */
    loadBalancerId?: string;
    /**
     * A Name Regex of ServerGroup.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The name of the ServerGroup.
     */
    serverGroupName?: string;
}

/**
 * A collection of values returned by ServerGroups.
 */
export interface ServerGroupsResult {
    /**
     * The collection of ServerGroup query.
     */
    readonly groups: outputs.clb.ServerGroupsGroup[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly loadBalancerId?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The name of the ServerGroup.
     */
    readonly serverGroupName?: string;
    /**
     * The total count of ServerGroup query.
     */
    readonly totalCount: number;
}

export function serverGroupsOutput(args?: ServerGroupsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ServerGroupsResult> {
    return pulumi.output(args).apply(a => serverGroups(a, opts))
}

/**
 * A collection of arguments for invoking ServerGroups.
 */
export interface ServerGroupsOutputArgs {
    /**
     * A list of ServerGroup IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The id of the Clb.
     */
    loadBalancerId?: pulumi.Input<string>;
    /**
     * A Name Regex of ServerGroup.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The name of the ServerGroup.
     */
    serverGroupName?: pulumi.Input<string>;
}
