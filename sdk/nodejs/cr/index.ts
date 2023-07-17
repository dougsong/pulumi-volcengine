// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./authorizationTokens";
export * from "./endpoint";
export * from "./endpoints";
export * from "./namespace";
export * from "./namespaces";
export * from "./registries";
export * from "./registry";
export * from "./repositories";
export * from "./repository";
export * from "./state";
export * from "./tag";
export * from "./tags";
export * from "./vpcEndpoint";
export * from "./vpcEndpoints";

// Import resources to register:
import { Endpoint } from "./endpoint";
import { Namespace } from "./namespace";
import { Registry } from "./registry";
import { Repository } from "./repository";
import { State } from "./state";
import { Tag } from "./tag";
import { VpcEndpoint } from "./vpcEndpoint";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:cr/endpoint:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "volcengine:cr/namespace:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "volcengine:cr/registry:Registry":
                return new Registry(name, <any>undefined, { urn })
            case "volcengine:cr/repository:Repository":
                return new Repository(name, <any>undefined, { urn })
            case "volcengine:cr/state:State":
                return new State(name, <any>undefined, { urn })
            case "volcengine:cr/tag:Tag":
                return new Tag(name, <any>undefined, { urn })
            case "volcengine:cr/vpcEndpoint:VpcEndpoint":
                return new VpcEndpoint(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "cr/endpoint", _module)
pulumi.runtime.registerResourceModule("volcengine", "cr/namespace", _module)
pulumi.runtime.registerResourceModule("volcengine", "cr/registry", _module)
pulumi.runtime.registerResourceModule("volcengine", "cr/repository", _module)
pulumi.runtime.registerResourceModule("volcengine", "cr/state", _module)
pulumi.runtime.registerResourceModule("volcengine", "cr/tag", _module)
pulumi.runtime.registerResourceModule("volcengine", "cr/vpcEndpoint", _module)
