// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to query detailed information of certificates
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const defaultCertificates = pulumi.output(volcengine.clb.Certificates({
 *     ids: ["cert-274scdwqufwg07fap8u5fu8pi"],
 * }));
 * ```
 */
export function certificates(args?: CertificatesArgs, opts?: pulumi.InvokeOptions): Promise<CertificatesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("volcengine:clb/certificates:Certificates", {
        "certificateName": args.certificateName,
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "outputFile": args.outputFile,
        "projectName": args.projectName,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking Certificates.
 */
export interface CertificatesArgs {
    /**
     * The name of the Certificate.
     */
    certificateName?: string;
    /**
     * The list of Certificate IDs.
     */
    ids?: string[];
    /**
     * The Name Regex of Certificate.
     */
    nameRegex?: string;
    /**
     * File name where to save data source results.
     */
    outputFile?: string;
    /**
     * The ProjectName of Certificate.
     */
    projectName?: string;
    /**
     * Tags.
     */
    tags?: inputs.clb.CertificatesTag[];
}

/**
 * A collection of values returned by Certificates.
 */
export interface CertificatesResult {
    /**
     * The name of the Certificate.
     */
    readonly certificateName?: string;
    /**
     * The collection of Certificate query.
     */
    readonly certificates: outputs.clb.CertificatesCertificate[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ids?: string[];
    readonly nameRegex?: string;
    readonly outputFile?: string;
    /**
     * The ProjectName of the Certificate.
     */
    readonly projectName?: string;
    /**
     * Tags.
     */
    readonly tags?: outputs.clb.CertificatesTag[];
    /**
     * The total count of Certificate query.
     */
    readonly totalCount: number;
}

export function certificatesOutput(args?: CertificatesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<CertificatesResult> {
    return pulumi.output(args).apply(a => certificates(a, opts))
}

/**
 * A collection of arguments for invoking Certificates.
 */
export interface CertificatesOutputArgs {
    /**
     * The name of the Certificate.
     */
    certificateName?: pulumi.Input<string>;
    /**
     * The list of Certificate IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Name Regex of Certificate.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * File name where to save data source results.
     */
    outputFile?: pulumi.Input<string>;
    /**
     * The ProjectName of Certificate.
     */
    projectName?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.clb.CertificatesTagArgs>[]>;
}
