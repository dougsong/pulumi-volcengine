// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to manage tos bucket policy
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const _default = new volcengine.tos.BucketPolicy("default", {
 *     bucketName: "bucket-20230418",
 *     policy: JSON.stringify({
 *         Statement: [{
 *             Sid: "test",
 *             Effect: "Allow",
 *             Principal: ["AccountId/subUserName"],
 *             Action: ["tos:List*"],
 *             Resource: ["trn:tos:::bucket-20230418"],
 *         }],
 *     }),
 * });
 * ```
 *
 * ## Import
 *
 * Tos Bucket can be imported using the id, e.g.
 *
 * ```sh
 *  $ pulumi import volcengine:tos/bucketPolicy:BucketPolicy default bucketName:policy
 * ```
 */
export class BucketPolicy extends pulumi.CustomResource {
    /**
     * Get an existing BucketPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BucketPolicyState, opts?: pulumi.CustomResourceOptions): BucketPolicy {
        return new BucketPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:tos/bucketPolicy:BucketPolicy';

    /**
     * Returns true if the given object is an instance of BucketPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BucketPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BucketPolicy.__pulumiType;
    }

    /**
     * The name of the bucket.
     */
    public readonly bucketName!: pulumi.Output<string>;
    /**
     * The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
     * documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
     */
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a BucketPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BucketPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BucketPolicyArgs | BucketPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BucketPolicyState | undefined;
            resourceInputs["bucketName"] = state ? state.bucketName : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as BucketPolicyArgs | undefined;
            if ((!args || args.bucketName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bucketName'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(BucketPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BucketPolicy resources.
 */
export interface BucketPolicyState {
    /**
     * The name of the bucket.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
     * documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
     */
    policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BucketPolicy resource.
 */
export interface BucketPolicyArgs {
    /**
     * The name of the bucket.
     */
    bucketName: pulumi.Input<string>;
    /**
     * The policy document. This is a JSON formatted string. For more information about building Volcengine IAM policy
     * documents with Terraform, see the [Volcengine IAM Policy Document Guide](https://www.volcengine.com/docs/6349/102127).
     */
    policy: pulumi.Input<string>;
}
