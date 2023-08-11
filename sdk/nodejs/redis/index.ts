// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AccountArgs, AccountState } from "./account";
export type Account = import("./account").Account;
export const Account: typeof import("./account").Account = null as any;
utilities.lazyLoad(exports, ["Account"], () => require("./account"));

export { AccountsArgs, AccountsResult, AccountsOutputArgs } from "./accounts";
export const accounts: typeof import("./accounts").accounts = null as any;
export const accountsOutput: typeof import("./accounts").accountsOutput = null as any;
utilities.lazyLoad(exports, ["accounts","accountsOutput"], () => require("./accounts"));

export { AllowListArgs, AllowListState } from "./allowList";
export type AllowList = import("./allowList").AllowList;
export const AllowList: typeof import("./allowList").AllowList = null as any;
utilities.lazyLoad(exports, ["AllowList"], () => require("./allowList"));

export { AllowListAssociateArgs, AllowListAssociateState } from "./allowListAssociate";
export type AllowListAssociate = import("./allowListAssociate").AllowListAssociate;
export const AllowListAssociate: typeof import("./allowListAssociate").AllowListAssociate = null as any;
utilities.lazyLoad(exports, ["AllowListAssociate"], () => require("./allowListAssociate"));

export { AllowListsArgs, AllowListsResult, AllowListsOutputArgs } from "./allowLists";
export const allowLists: typeof import("./allowLists").allowLists = null as any;
export const allowListsOutput: typeof import("./allowLists").allowListsOutput = null as any;
utilities.lazyLoad(exports, ["allowLists","allowListsOutput"], () => require("./allowLists"));

export { BackupArgs, BackupState } from "./backup";
export type Backup = import("./backup").Backup;
export const Backup: typeof import("./backup").Backup = null as any;
utilities.lazyLoad(exports, ["Backup"], () => require("./backup"));

export { BackupRestoreArgs, BackupRestoreState } from "./backupRestore";
export type BackupRestore = import("./backupRestore").BackupRestore;
export const BackupRestore: typeof import("./backupRestore").BackupRestore = null as any;
utilities.lazyLoad(exports, ["BackupRestore"], () => require("./backupRestore"));

export { BackupsArgs, BackupsResult, BackupsOutputArgs } from "./backups";
export const backups: typeof import("./backups").backups = null as any;
export const backupsOutput: typeof import("./backups").backupsOutput = null as any;
utilities.lazyLoad(exports, ["backups","backupsOutput"], () => require("./backups"));

export { ContinuousBackupArgs, ContinuousBackupState } from "./continuousBackup";
export type ContinuousBackup = import("./continuousBackup").ContinuousBackup;
export const ContinuousBackup: typeof import("./continuousBackup").ContinuousBackup = null as any;
utilities.lazyLoad(exports, ["ContinuousBackup"], () => require("./continuousBackup"));

export { EndpointArgs, EndpointState } from "./endpoint";
export type Endpoint = import("./endpoint").Endpoint;
export const Endpoint: typeof import("./endpoint").Endpoint = null as any;
utilities.lazyLoad(exports, ["Endpoint"], () => require("./endpoint"));

export { InstanceArgs, InstanceState } from "./instance";
export type Instance = import("./instance").Instance;
export const Instance: typeof import("./instance").Instance = null as any;
utilities.lazyLoad(exports, ["Instance"], () => require("./instance"));

export { InstancesArgs, InstancesResult, InstancesOutputArgs } from "./instances";
export const instances: typeof import("./instances").instances = null as any;
export const instancesOutput: typeof import("./instances").instancesOutput = null as any;
utilities.lazyLoad(exports, ["instances","instancesOutput"], () => require("./instances"));

export { PitrTimeWindowsArgs, PitrTimeWindowsResult, PitrTimeWindowsOutputArgs } from "./pitrTimeWindows";
export const pitrTimeWindows: typeof import("./pitrTimeWindows").pitrTimeWindows = null as any;
export const pitrTimeWindowsOutput: typeof import("./pitrTimeWindows").pitrTimeWindowsOutput = null as any;
utilities.lazyLoad(exports, ["pitrTimeWindows","pitrTimeWindowsOutput"], () => require("./pitrTimeWindows"));

export { RegionsArgs, RegionsResult, RegionsOutputArgs } from "./regions";
export const regions: typeof import("./regions").regions = null as any;
export const regionsOutput: typeof import("./regions").regionsOutput = null as any;
utilities.lazyLoad(exports, ["regions","regionsOutput"], () => require("./regions"));

export { StateArgs, StateState } from "./state";
export type State = import("./state").State;
export const State: typeof import("./state").State = null as any;
utilities.lazyLoad(exports, ["State"], () => require("./state"));

export { ZonesArgs, ZonesResult, ZonesOutputArgs } from "./zones";
export const zones: typeof import("./zones").zones = null as any;
export const zonesOutput: typeof import("./zones").zonesOutput = null as any;
utilities.lazyLoad(exports, ["zones","zonesOutput"], () => require("./zones"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "volcengine:redis/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "volcengine:redis/allowList:AllowList":
                return new AllowList(name, <any>undefined, { urn })
            case "volcengine:redis/allowListAssociate:AllowListAssociate":
                return new AllowListAssociate(name, <any>undefined, { urn })
            case "volcengine:redis/backup:Backup":
                return new Backup(name, <any>undefined, { urn })
            case "volcengine:redis/backupRestore:BackupRestore":
                return new BackupRestore(name, <any>undefined, { urn })
            case "volcengine:redis/continuousBackup:ContinuousBackup":
                return new ContinuousBackup(name, <any>undefined, { urn })
            case "volcengine:redis/endpoint:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "volcengine:redis/instance:Instance":
                return new Instance(name, <any>undefined, { urn })
            case "volcengine:redis/state:State":
                return new State(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("volcengine", "redis/account", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/allowList", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/allowListAssociate", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/backup", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/backupRestore", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/continuousBackup", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/endpoint", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/instance", _module)
pulumi.runtime.registerResourceModule("volcengine", "redis/state", _module)
