// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of cr tags
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cr.Tags({
 *     namespace: "test",
 *     registry: "enterprise-1",
 *     repository: "repo",
 *     types: ["Image"],
 * });
 * ```
 */
export function tags(args: TagsArgs, opts?: pulumi.InvokeOptions): Promise<TagsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:cr/tags:Tags", {
        "names": args.names,
        "namespace": args.namespace,
        "outputFile": args.outputFile,
        "registry": args.registry,
        "repository": args.repository,
        "types": args.types,
    }, opts);
}

/**
 * A collection of arguments for invoking Tags.
 */
export interface TagsArgs {
    /**
     * The list of instance names.
     */
    names?: string[];
    /**
     * The CR namespace.
     */
    namespace: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The CR instance name.
     */
    registry: string;
    /**
     * The repository name.
     */
    repository: string;
    /**
     * The list of OCI product tag type.
     */
    types?: string[];
}

/**
 * A collection of values returned by Tags.
 */
export interface TagsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly names?: string[];
    readonly namespace: string;
    readonly outputFile?: string;
    readonly registry: string;
    readonly repository: string;
    /**
     * The collection of repository query.
     */
    readonly tags: outputs.cr.TagsTag[];
    /**
     * The total count of tag query.
     */
    readonly totalCount: number;
    readonly types?: string[];
}
/**
 * Use this data source to query detailed information of cr tags
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const foo = volcengine.cr.Tags({
 *     namespace: "test",
 *     registry: "enterprise-1",
 *     repository: "repo",
 *     types: ["Image"],
 * });
 * ```
 */
export function tagsOutput(args: TagsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<TagsResult> {
    return pulumi.output(args).apply((a: any) => tags(a, opts))
}

/**
 * A collection of arguments for invoking Tags.
 */
export interface TagsOutputArgs {
    /**
     * The list of instance names.
     */
    names?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The CR namespace.
     */
    namespace: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The CR instance name.
     */
    registry: pulumi.Input<string>;
    /**
     * The repository name.
     */
    repository: pulumi.Input<string>;
    /**
     * The list of OCI product tag type.
     */
    types?: pulumi.Input<pulumi.Input<string>[]>;
}
