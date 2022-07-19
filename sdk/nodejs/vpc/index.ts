// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./acl";
export * from "./aclEntry";
export * from "./acls";
export * from "./certificates";
export * from "./gateways";
export * from "./networkInterface";
export * from "./networkInterfaceAttach";
export * from "./networkInterfaces";
export * from "./routeEntries";
export * from "./routeEntry";
export * from "./routeTable";
export * from "./routeTableAssociate";
export * from "./routeTables";
export * from "./securityGroup";
export * from "./securityGroupRule";
export * from "./securityGroups";
export * from "./snatEntries";
export * from "./subnet";
export * from "./subnets";
export * from "./vpc";
export * from "./vpcs";
export * from "./zones";

// Import resources to register:
import { Acl } from "./acl";
import { AclEntry } from "./aclEntry";
import { NetworkInterface } from "./networkInterface";
import { NetworkInterfaceAttach } from "./networkInterfaceAttach";
import { RouteEntry } from "./routeEntry";
import { RouteTable } from "./routeTable";
import { RouteTableAssociate } from "./routeTableAssociate";
import { SecurityGroup } from "./securityGroup";
import { SecurityGroupRule } from "./securityGroupRule";
import { Subnet } from "./subnet";
import { Vpc } from "./vpc";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:Vpc/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "volcengine:Vpc/aclEntry:AclEntry":
                return new AclEntry(name, <any>undefined, { urn })
            case "volcengine:Vpc/networkInterface:NetworkInterface":
                return new NetworkInterface(name, <any>undefined, { urn })
            case "volcengine:Vpc/networkInterfaceAttach:NetworkInterfaceAttach":
                return new NetworkInterfaceAttach(name, <any>undefined, { urn })
            case "volcengine:Vpc/routeEntry:RouteEntry":
                return new RouteEntry(name, <any>undefined, { urn })
            case "volcengine:Vpc/routeTable:RouteTable":
                return new RouteTable(name, <any>undefined, { urn })
            case "volcengine:Vpc/routeTableAssociate:RouteTableAssociate":
                return new RouteTableAssociate(name, <any>undefined, { urn })
            case "volcengine:Vpc/securityGroup:SecurityGroup":
                return new SecurityGroup(name, <any>undefined, { urn })
            case "volcengine:Vpc/securityGroupRule:SecurityGroupRule":
                return new SecurityGroupRule(name, <any>undefined, { urn })
            case "volcengine:Vpc/subnet:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "volcengine:Vpc/vpc:Vpc":
                return new Vpc(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "Vpc/acl", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/aclEntry", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/networkInterface", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/networkInterfaceAttach", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/routeEntry", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/routeTable", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/routeTableAssociate", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/securityGroup", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/securityGroupRule", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/subnet", _module)
pulumi.runtime.registerResourceModule("volcengine", "Vpc/vpc", _module)
