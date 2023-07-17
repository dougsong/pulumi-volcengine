// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of ecs instances
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = pulumi.output(volcengine.ecs.Instances({
 *     ids: ["i-ebgy6xmgjve0384ncgsc"],
 * }));
 * ```
 */
export function instances(args?: InstancesArgs, opts?: pulumi.InvokeOptions): Promise<InstancesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:ecs/instances:Instances", {
        "deploymentSetIds": args.deploymentSetIds,
        "hpcClusterId": args.hpcClusterId,
        "ids": args.ids,
        "instanceChargeType": args.instanceChargeType,
        "keyPairName": args.keyPairName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "primaryIpAddress": args.primaryIpAddress,
        "projectName": args.projectName,
        "status": args.status,
        "tags": args.tags,
        "vpcId": args.vpcId,
        "zoneId": args.zoneId,
    }, opts);
}

/**
 * A collection of arguments for invoking Instances.
 */
export interface InstancesArgs {
    /**
     * A list of DeploymentSet IDs.
     */
    deploymentSetIds?: string[];
    /**
     * The hpc cluster ID of ECS instance.
     */
    hpcClusterId?: string;
    /**
     * A list of ECS instance IDs.
     */
    ids?: string[];
    /**
     * The charge type of ECS instance.
     */
    instanceChargeType?: string;
    /**
     * The key pair name of ECS instance.
     */
    keyPairName?: string;
    /**
     * A Name Regex of ECS instance.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The primary ip address of ECS instance.
     */
    primaryIpAddress?: string;
    /**
     * The ProjectName of ECS instance.
     */
    projectName?: string;
    /**
     * The status of ECS instance.
     */
    status?: string;
    /**
     * Tags.
     */
    tags?: inputs.ecs.InstancesTag[];
    /**
     * The VPC ID of ECS instance.
     */
    vpcId?: string;
    /**
     * The available zone ID of ECS instance.
     */
    zoneId?: string;
}

/**
 * A collection of values returned by Instances.
 */
export interface InstancesResult {
    readonly deploymentSetIds?: string[];
    readonly hpcClusterId?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    /**
     * The charge type of ECS instance.
     */
    readonly instanceChargeType?: string;
    /**
     * The collection of ECS instance query.
     */
    readonly instances: outputs.ecs.InstancesInstance[];
    /**
     * The ssh key name of ECS instance.
     */
    readonly keyPairName?: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The private ip address of networkInterface.
     */
    readonly primaryIpAddress?: string;
    /**
     * The ProjectName of ECS instance.
     */
    readonly projectName?: string;
    /**
     * The status of ECS instance.
     */
    readonly status?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.ecs.InstancesTag[];
    /**
     * The total count of ECS instance query.
     */
    readonly totalCount: number;
    /**
     * The VPC ID of ECS instance.
     */
    readonly vpcId?: string;
    /**
     * The available zone ID of ECS instance.
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
     * A list of DeploymentSet IDs.
     */
    deploymentSetIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The hpc cluster ID of ECS instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * A list of ECS instance IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The charge type of ECS instance.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The key pair name of ECS instance.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * A Name Regex of ECS instance.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The primary ip address of ECS instance.
     */
    primaryIpAddress?: pulumi.Input<string>;
    /**
     * The ProjectName of ECS instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The status of ECS instance.
     */
    status?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ecs.InstancesTagArgs>[]>;
    /**
     * The VPC ID of ECS instance.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The available zone ID of ECS instance.
     */
    zoneId?: pulumi.Input<string>;
}
