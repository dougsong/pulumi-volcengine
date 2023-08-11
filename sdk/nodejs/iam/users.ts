// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of iam users
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.iam.Users({});
 * ```
 */
export function users(args?: UsersArgs, opts?: pulumi.InvokeOptions): Promise<UsersResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("volcengine:iam/users:Users", {
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "userNames": args.userNames,
    }, opts);
}

/**
 * A collection of arguments for invoking Users.
 */
export interface UsersArgs {
    /**
     * A Name Regex of IAM.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * A list of user names.
     */
    userNames?: string[];
}

/**
 * A collection of values returned by Users.
 */
export interface UsersResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The total count of user query.
     */
    readonly totalCount: number;
    readonly userNames?: string[];
    /**
     * The collection of user.
     */
    readonly users: outputs.iam.UsersUser[];
}
/**
 * Use this data source to query detailed information of iam users
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const default = volcengine.iam.Users({});
 * ```
 */
export function usersOutput(args?: UsersOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<UsersResult> {
    return pulumi.output(args).apply((a: any) => users(a, opts))
}

/**
 * A collection of arguments for invoking Users.
 */
export interface UsersOutputArgs {
    /**
     * A Name Regex of IAM.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * A list of user names.
     */
    userNames?: pulumi.Input<pulumi.Input<string>[]>;
}
