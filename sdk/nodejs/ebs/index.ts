// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./volume";
export * from "./volumeAttach";
export * from "./volumes";

// Import resources to register:
import { Volume } from "./volume";
import { VolumeAttach } from "./volumeAttach";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:ebs/volume:Volume":
                return new Volume(name, <any>undefined, { urn })
            case "volcengine:ebs/volumeAttach:VolumeAttach":
                return new VolumeAttach(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "ebs/volume", _module)
pulumi.runtime.registerResourceModule("volcengine", "ebs/volumeAttach", _module)
