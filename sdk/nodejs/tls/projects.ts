// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of tls projects
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.tls.Projects({
 *     projectId: "e020c978-4f05-40e1-9167-0113d3ef****",
 * });
 * ```
 */
export function projects(args?: ProjectsArgs, opts?: pulumi.InvokeOptions): Promise<ProjectsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:tls/projects:Projects", {
        "iamProjectName": args.iamProjectName,
        "isFullName": args.isFullName,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectId": args.projectId,
        "projectName": args.projectName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking Projects.
 */
export interface ProjectsArgs {
    /**
     * The IAM project name of the tls project.
     */
    iamProjectName?: string;
    /**
     * Whether to match accurately when filtering based on ProjectName.
     */
    isFullName?: boolean;
    /**
     * A Name Regex of tls project.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
     */
    projectId?: string;
    /**
     * The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
     */
    projectName?: string;
    /**
     * Tags.
     */
    tags?: inputs.tls.ProjectsTag[];
}

/**
 * A collection of values returned by Projects.
 */
export interface ProjectsResult {
    /**
     * The IAM project name of the tls project.
     */
    readonly iamProjectName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isFullName?: boolean;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The ID of the tls project.
     */
    readonly projectId?: string;
    /**
     * The name of the tls project.
     */
    readonly projectName?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.tls.ProjectsTag[];
    /**
     * The collection of tls project query.
     */
    readonly tlsProjects: outputs.tls.ProjectsTlsProject[];
    /**
     * The total count of tls project query.
     */
    readonly totalCount: number;
}
/**
 * Use this data source to query detailed information of tls projects
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.tls.Projects({
 *     projectId: "e020c978-4f05-40e1-9167-0113d3ef****",
 * });
 * ```
 */
export function projectsOutput(args?: ProjectsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<ProjectsResult> {
    return pulumi.output(args).apply((a: any) => projects(a, opts))
}

/**
 * A collection of arguments for invoking Projects.
 */
export interface ProjectsOutputArgs {
    /**
     * The IAM project name of the tls project.
     */
    iamProjectName?: pulumi.Input<string>;
    /**
     * Whether to match accurately when filtering based on ProjectName.
     */
    isFullName?: pulumi.Input<boolean>;
    /**
     * A Name Regex of tls project.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The id of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
     */
    projectId?: pulumi.Input<string>;
    /**
     * The name of tls project. This field supports fuzzy queries. It is not supported to specify both ProjectName and ProjectId at the same time.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.tls.ProjectsTagArgs>[]>;
}
