// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as pulumi from "@volcengine/pulumi";
 * import * as volcengine from "@pulumi/volcengine";
 *
 * const fooZones = volcengine.ecs.Zones({});
 * const fooVpc = new volcengine.vpc.Vpc("fooVpc", {
 *     vpcName: "acc-test-vpc",
 *     cidrBlock: "172.16.0.0/16",
 * });
 * const fooSubnet = new volcengine.vpc.Subnet("fooSubnet", {
 *     subnetName: "acc-test-subnet",
 *     cidrBlock: "172.16.0.0/24",
 *     zoneId: fooZones.then(fooZones => fooZones.zones?[0]?.id),
 *     vpcId: fooVpc.id,
 * });
 * const fooSecurityGroup = new volcengine.vpc.SecurityGroup("fooSecurityGroup", {
 *     securityGroupName: "acc-test-security-group",
 *     vpcId: fooVpc.id,
 * });
 * const fooImages = volcengine.ecs.Images({
 *     osType: "Linux",
 *     visibility: "public",
 *     instanceTypeId: "ecs.g1.large",
 * });
 * const fooInstance = new volcengine.ecs.Instance("fooInstance", {
 *     instanceName: "acc-test-ecs",
 *     description: "acc-test",
 *     hostName: "tf-acc-test",
 *     imageId: fooImages.then(fooImages => fooImages.images?[0]?.imageId),
 *     instanceType: "ecs.g1.large",
 *     password: "93f0cb0614Aab12",
 *     instanceChargeType: "PostPaid",
 *     systemVolumeType: "ESSD_PL0",
 *     systemVolumeSize: 40,
 *     dataVolumes: [{
 *         volumeType: "ESSD_PL0",
 *         size: 50,
 *         deleteWithInstance: true,
 *     }],
 *     subnetId: fooSubnet.id,
 *     securityGroupIds: [fooSecurityGroup.id],
 *     projectName: "default",
 *     tags: [{
 *         key: "k1",
 *         value: "v1",
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ECS Instance can be imported using the id, e.g. If Import,The data_volumes is sort by volume name
 *
 * ```sh
 *  $ pulumi import volcengine:ecs/instance:Instance default i-mizl7m1kqccg5smt1bdpijuj
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'volcengine:ecs/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * The auto renew flag of ECS instance.Only effective when instanceChargeType is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly autoRenew!: pulumi.Output<boolean | undefined>;
    /**
     * The auto renew period of ECS instance.Only effective when instanceChargeType is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly autoRenewPeriod!: pulumi.Output<number | undefined>;
    /**
     * The option of cpu.
     */
    public readonly cpuOptions!: pulumi.Output<outputs.ecs.InstanceCpuOptions>;
    /**
     * The number of ECS instance CPU cores.
     */
    public /*out*/ readonly cpus!: pulumi.Output<number>;
    /**
     * The create time of ECS instance.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The data volumes collection of  ECS instance.
     */
    public readonly dataVolumes!: pulumi.Output<outputs.ecs.InstanceDataVolume[]>;
    /**
     * The ID of Ecs Deployment Set.
     */
    public readonly deploymentSetId!: pulumi.Output<string>;
    /**
     * The description of ECS instance.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The GPU device info of Instance.
     */
    public /*out*/ readonly gpuDevices!: pulumi.Output<outputs.ecs.InstanceGpuDevice[]>;
    /**
     * The host name of ECS instance.
     */
    public readonly hostName!: pulumi.Output<string>;
    /**
     * The hpc cluster ID of ECS instance.
     */
    public readonly hpcClusterId!: pulumi.Output<string | undefined>;
    /**
     * The Image ID of ECS instance.
     */
    public readonly imageId!: pulumi.Output<string>;
    /**
     * The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.
     */
    public readonly includeDataVolumes!: pulumi.Output<boolean | undefined>;
    /**
     * The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.
     */
    public readonly instanceChargeType!: pulumi.Output<string>;
    /**
     * The ID of ECS instance.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * The name of ECS instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The instance type of ECS instance.
     */
    public readonly instanceType!: pulumi.Output<string>;
    /**
     * The number of IPv6 addresses to be automatically assigned from within the CIDR block of the subnet that hosts the ENI. Valid values: 1 to 10.
     */
    public readonly ipv6AddressCount!: pulumi.Output<number>;
    /**
     * One or more IPv6 addresses selected from within the CIDR block of the subnet that hosts the ENI. Support up to 10.
     * You cannot specify both the ipv6Addresses and ipv6AddressCount parameters.
     */
    public readonly ipv6Addresses!: pulumi.Output<string[]>;
    /**
     * The Flag of GPU instance.If the instance is GPU,The flag is true.
     */
    public /*out*/ readonly isGpu!: pulumi.Output<boolean>;
    /**
     * Whether to keep the mirror settings. Only custom images and shared images support this field.
     * When the value of this field is true, the Password and KeyPairName cannot be specified.
     * When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly keepImageCredential!: pulumi.Output<boolean | undefined>;
    /**
     * The ssh key ID of ECS instance.
     */
    public /*out*/ readonly keyPairId!: pulumi.Output<string>;
    /**
     * The ssh key name of ECS instance.
     */
    public readonly keyPairName!: pulumi.Output<string>;
    /**
     * The memory size of ECS instance.
     */
    public /*out*/ readonly memorySize!: pulumi.Output<number>;
    /**
     * The ID of primary networkInterface.
     */
    public /*out*/ readonly networkInterfaceId!: pulumi.Output<string>;
    /**
     * The os name of ECS instance.
     */
    public /*out*/ readonly osName!: pulumi.Output<string>;
    /**
     * The os type of ECS instance.
     */
    public /*out*/ readonly osType!: pulumi.Output<string>;
    /**
     * The password of ECS instance.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * The period of ECS instance.Only effective when instanceChargeType is PrePaid. Default is 12. Unit is Month.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The private ip address of primary networkInterface.
     */
    public /*out*/ readonly primaryIpAddress!: pulumi.Output<string>;
    /**
     * The ProjectName of the ecs instance.
     */
    public readonly projectName!: pulumi.Output<string | undefined>;
    /**
     * The secondary networkInterface detail collection of ECS instance.
     */
    public readonly secondaryNetworkInterfaces!: pulumi.Output<outputs.ecs.InstanceSecondaryNetworkInterface[]>;
    /**
     * The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    public readonly securityEnhancementStrategy!: pulumi.Output<string | undefined>;
    /**
     * The security group ID set of primary networkInterface.
     */
    public readonly securityGroupIds!: pulumi.Output<string[]>;
    /**
     * The spot strategy will autoremove instance in some conditions.Please make sure you can maintain instance lifecycle before auto remove.The spot strategy of ECS instance, the value can be `NoSpot` or `SpotAsPriceGo`.
     */
    public readonly spotStrategy!: pulumi.Output<string>;
    /**
     * The status of ECS instance.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The stop mode of ECS instance.
     */
    public /*out*/ readonly stoppedMode!: pulumi.Output<string>;
    /**
     * The subnet ID of primary networkInterface.
     */
    public readonly subnetId!: pulumi.Output<string>;
    /**
     * The ID of system volume.
     */
    public /*out*/ readonly systemVolumeId!: pulumi.Output<string>;
    /**
     * The size of system volume. The value range of the system volume size is ESSD_PL0: 20~2048, ESSD_FlexPL: 20~2048, PTSSD: 10~500.
     */
    public readonly systemVolumeSize!: pulumi.Output<number>;
    /**
     * The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
     */
    public readonly systemVolumeType!: pulumi.Output<string>;
    /**
     * Tags.
     */
    public readonly tags!: pulumi.Output<outputs.ecs.InstanceTag[] | undefined>;
    /**
     * The update time of ECS instance.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;
    /**
     * The user data of ECS instance, this field must be encrypted with base64.
     */
    public readonly userData!: pulumi.Output<string>;
    /**
     * The VPC ID of ECS instance.
     */
    public /*out*/ readonly vpcId!: pulumi.Output<string>;
    /**
     * The available zone ID of ECS instance.
     */
    public readonly zoneId!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["autoRenew"] = state ? state.autoRenew : undefined;
            resourceInputs["autoRenewPeriod"] = state ? state.autoRenewPeriod : undefined;
            resourceInputs["cpuOptions"] = state ? state.cpuOptions : undefined;
            resourceInputs["cpus"] = state ? state.cpus : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["dataVolumes"] = state ? state.dataVolumes : undefined;
            resourceInputs["deploymentSetId"] = state ? state.deploymentSetId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["gpuDevices"] = state ? state.gpuDevices : undefined;
            resourceInputs["hostName"] = state ? state.hostName : undefined;
            resourceInputs["hpcClusterId"] = state ? state.hpcClusterId : undefined;
            resourceInputs["imageId"] = state ? state.imageId : undefined;
            resourceInputs["includeDataVolumes"] = state ? state.includeDataVolumes : undefined;
            resourceInputs["instanceChargeType"] = state ? state.instanceChargeType : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["instanceType"] = state ? state.instanceType : undefined;
            resourceInputs["ipv6AddressCount"] = state ? state.ipv6AddressCount : undefined;
            resourceInputs["ipv6Addresses"] = state ? state.ipv6Addresses : undefined;
            resourceInputs["isGpu"] = state ? state.isGpu : undefined;
            resourceInputs["keepImageCredential"] = state ? state.keepImageCredential : undefined;
            resourceInputs["keyPairId"] = state ? state.keyPairId : undefined;
            resourceInputs["keyPairName"] = state ? state.keyPairName : undefined;
            resourceInputs["memorySize"] = state ? state.memorySize : undefined;
            resourceInputs["networkInterfaceId"] = state ? state.networkInterfaceId : undefined;
            resourceInputs["osName"] = state ? state.osName : undefined;
            resourceInputs["osType"] = state ? state.osType : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["primaryIpAddress"] = state ? state.primaryIpAddress : undefined;
            resourceInputs["projectName"] = state ? state.projectName : undefined;
            resourceInputs["secondaryNetworkInterfaces"] = state ? state.secondaryNetworkInterfaces : undefined;
            resourceInputs["securityEnhancementStrategy"] = state ? state.securityEnhancementStrategy : undefined;
            resourceInputs["securityGroupIds"] = state ? state.securityGroupIds : undefined;
            resourceInputs["spotStrategy"] = state ? state.spotStrategy : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["stoppedMode"] = state ? state.stoppedMode : undefined;
            resourceInputs["subnetId"] = state ? state.subnetId : undefined;
            resourceInputs["systemVolumeId"] = state ? state.systemVolumeId : undefined;
            resourceInputs["systemVolumeSize"] = state ? state.systemVolumeSize : undefined;
            resourceInputs["systemVolumeType"] = state ? state.systemVolumeType : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["updatedAt"] = state ? state.updatedAt : undefined;
            resourceInputs["userData"] = state ? state.userData : undefined;
            resourceInputs["vpcId"] = state ? state.vpcId : undefined;
            resourceInputs["zoneId"] = state ? state.zoneId : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.imageId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'imageId'");
            }
            if ((!args || args.instanceType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceType'");
            }
            if ((!args || args.securityGroupIds === undefined) && !opts.urn) {
                throw new Error("Missing required property 'securityGroupIds'");
            }
            if ((!args || args.subnetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetId'");
            }
            if ((!args || args.systemVolumeSize === undefined) && !opts.urn) {
                throw new Error("Missing required property 'systemVolumeSize'");
            }
            if ((!args || args.systemVolumeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'systemVolumeType'");
            }
            resourceInputs["autoRenew"] = args ? args.autoRenew : undefined;
            resourceInputs["autoRenewPeriod"] = args ? args.autoRenewPeriod : undefined;
            resourceInputs["cpuOptions"] = args ? args.cpuOptions : undefined;
            resourceInputs["dataVolumes"] = args ? args.dataVolumes : undefined;
            resourceInputs["deploymentSetId"] = args ? args.deploymentSetId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["hostName"] = args ? args.hostName : undefined;
            resourceInputs["hpcClusterId"] = args ? args.hpcClusterId : undefined;
            resourceInputs["imageId"] = args ? args.imageId : undefined;
            resourceInputs["includeDataVolumes"] = args ? args.includeDataVolumes : undefined;
            resourceInputs["instanceChargeType"] = args ? args.instanceChargeType : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["instanceType"] = args ? args.instanceType : undefined;
            resourceInputs["ipv6AddressCount"] = args ? args.ipv6AddressCount : undefined;
            resourceInputs["ipv6Addresses"] = args ? args.ipv6Addresses : undefined;
            resourceInputs["keepImageCredential"] = args ? args.keepImageCredential : undefined;
            resourceInputs["keyPairName"] = args ? args.keyPairName : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["projectName"] = args ? args.projectName : undefined;
            resourceInputs["secondaryNetworkInterfaces"] = args ? args.secondaryNetworkInterfaces : undefined;
            resourceInputs["securityEnhancementStrategy"] = args ? args.securityEnhancementStrategy : undefined;
            resourceInputs["securityGroupIds"] = args ? args.securityGroupIds : undefined;
            resourceInputs["spotStrategy"] = args ? args.spotStrategy : undefined;
            resourceInputs["subnetId"] = args ? args.subnetId : undefined;
            resourceInputs["systemVolumeSize"] = args ? args.systemVolumeSize : undefined;
            resourceInputs["systemVolumeType"] = args ? args.systemVolumeType : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["userData"] = args ? args.userData : undefined;
            resourceInputs["zoneId"] = args ? args.zoneId : undefined;
            resourceInputs["cpus"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["gpuDevices"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["isGpu"] = undefined /*out*/;
            resourceInputs["keyPairId"] = undefined /*out*/;
            resourceInputs["memorySize"] = undefined /*out*/;
            resourceInputs["networkInterfaceId"] = undefined /*out*/;
            resourceInputs["osName"] = undefined /*out*/;
            resourceInputs["osType"] = undefined /*out*/;
            resourceInputs["primaryIpAddress"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["stoppedMode"] = undefined /*out*/;
            resourceInputs["systemVolumeId"] = undefined /*out*/;
            resourceInputs["updatedAt"] = undefined /*out*/;
            resourceInputs["vpcId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * The auto renew flag of ECS instance.Only effective when instanceChargeType is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The auto renew period of ECS instance.Only effective when instanceChargeType is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The option of cpu.
     */
    cpuOptions?: pulumi.Input<inputs.ecs.InstanceCpuOptions>;
    /**
     * The number of ECS instance CPU cores.
     */
    cpus?: pulumi.Input<number>;
    /**
     * The create time of ECS instance.
     */
    createdAt?: pulumi.Input<string>;
    /**
     * The data volumes collection of  ECS instance.
     */
    dataVolumes?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceDataVolume>[]>;
    /**
     * The ID of Ecs Deployment Set.
     */
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The description of ECS instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The GPU device info of Instance.
     */
    gpuDevices?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceGpuDevice>[]>;
    /**
     * The host name of ECS instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The hpc cluster ID of ECS instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * The Image ID of ECS instance.
     */
    imageId?: pulumi.Input<string>;
    /**
     * The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.
     */
    includeDataVolumes?: pulumi.Input<boolean>;
    /**
     * The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The ID of ECS instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The name of ECS instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The instance type of ECS instance.
     */
    instanceType?: pulumi.Input<string>;
    /**
     * The number of IPv6 addresses to be automatically assigned from within the CIDR block of the subnet that hosts the ENI. Valid values: 1 to 10.
     */
    ipv6AddressCount?: pulumi.Input<number>;
    /**
     * One or more IPv6 addresses selected from within the CIDR block of the subnet that hosts the ENI. Support up to 10.
     * You cannot specify both the ipv6Addresses and ipv6AddressCount parameters.
     */
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Flag of GPU instance.If the instance is GPU,The flag is true.
     */
    isGpu?: pulumi.Input<boolean>;
    /**
     * Whether to keep the mirror settings. Only custom images and shared images support this field.
     * When the value of this field is true, the Password and KeyPairName cannot be specified.
     * When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    keepImageCredential?: pulumi.Input<boolean>;
    /**
     * The ssh key ID of ECS instance.
     */
    keyPairId?: pulumi.Input<string>;
    /**
     * The ssh key name of ECS instance.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The memory size of ECS instance.
     */
    memorySize?: pulumi.Input<number>;
    /**
     * The ID of primary networkInterface.
     */
    networkInterfaceId?: pulumi.Input<string>;
    /**
     * The os name of ECS instance.
     */
    osName?: pulumi.Input<string>;
    /**
     * The os type of ECS instance.
     */
    osType?: pulumi.Input<string>;
    /**
     * The password of ECS instance.
     */
    password?: pulumi.Input<string>;
    /**
     * The period of ECS instance.Only effective when instanceChargeType is PrePaid. Default is 12. Unit is Month.
     */
    period?: pulumi.Input<number>;
    /**
     * The private ip address of primary networkInterface.
     */
    primaryIpAddress?: pulumi.Input<string>;
    /**
     * The ProjectName of the ecs instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The secondary networkInterface detail collection of ECS instance.
     */
    secondaryNetworkInterfaces?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceSecondaryNetworkInterface>[]>;
    /**
     * The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    securityEnhancementStrategy?: pulumi.Input<string>;
    /**
     * The security group ID set of primary networkInterface.
     */
    securityGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The spot strategy will autoremove instance in some conditions.Please make sure you can maintain instance lifecycle before auto remove.The spot strategy of ECS instance, the value can be `NoSpot` or `SpotAsPriceGo`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The status of ECS instance.
     */
    status?: pulumi.Input<string>;
    /**
     * The stop mode of ECS instance.
     */
    stoppedMode?: pulumi.Input<string>;
    /**
     * The subnet ID of primary networkInterface.
     */
    subnetId?: pulumi.Input<string>;
    /**
     * The ID of system volume.
     */
    systemVolumeId?: pulumi.Input<string>;
    /**
     * The size of system volume. The value range of the system volume size is ESSD_PL0: 20~2048, ESSD_FlexPL: 20~2048, PTSSD: 10~500.
     */
    systemVolumeSize?: pulumi.Input<number>;
    /**
     * The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
     */
    systemVolumeType?: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceTag>[]>;
    /**
     * The update time of ECS instance.
     */
    updatedAt?: pulumi.Input<string>;
    /**
     * The user data of ECS instance, this field must be encrypted with base64.
     */
    userData?: pulumi.Input<string>;
    /**
     * The VPC ID of ECS instance.
     */
    vpcId?: pulumi.Input<string>;
    /**
     * The available zone ID of ECS instance.
     */
    zoneId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The auto renew flag of ECS instance.Only effective when instanceChargeType is PrePaid. Default is true.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    autoRenew?: pulumi.Input<boolean>;
    /**
     * The auto renew period of ECS instance.Only effective when instanceChargeType is PrePaid. Default is 1.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    autoRenewPeriod?: pulumi.Input<number>;
    /**
     * The option of cpu.
     */
    cpuOptions?: pulumi.Input<inputs.ecs.InstanceCpuOptions>;
    /**
     * The data volumes collection of  ECS instance.
     */
    dataVolumes?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceDataVolume>[]>;
    /**
     * The ID of Ecs Deployment Set.
     */
    deploymentSetId?: pulumi.Input<string>;
    /**
     * The description of ECS instance.
     */
    description?: pulumi.Input<string>;
    /**
     * The host name of ECS instance.
     */
    hostName?: pulumi.Input<string>;
    /**
     * The hpc cluster ID of ECS instance.
     */
    hpcClusterId?: pulumi.Input<string>;
    /**
     * The Image ID of ECS instance.
     */
    imageId: pulumi.Input<string>;
    /**
     * The include data volumes flag of ECS instance.Only effective when change instance charge type.include_data_volumes.
     */
    includeDataVolumes?: pulumi.Input<boolean>;
    /**
     * The charge type of ECS instance, the value can be `PrePaid` or `PostPaid`.
     */
    instanceChargeType?: pulumi.Input<string>;
    /**
     * The name of ECS instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The instance type of ECS instance.
     */
    instanceType: pulumi.Input<string>;
    /**
     * The number of IPv6 addresses to be automatically assigned from within the CIDR block of the subnet that hosts the ENI. Valid values: 1 to 10.
     */
    ipv6AddressCount?: pulumi.Input<number>;
    /**
     * One or more IPv6 addresses selected from within the CIDR block of the subnet that hosts the ENI. Support up to 10.
     * You cannot specify both the ipv6Addresses and ipv6AddressCount parameters.
     */
    ipv6Addresses?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Whether to keep the mirror settings. Only custom images and shared images support this field.
     * When the value of this field is true, the Password and KeyPairName cannot be specified.
     * When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    keepImageCredential?: pulumi.Input<boolean>;
    /**
     * The ssh key name of ECS instance.
     */
    keyPairName?: pulumi.Input<string>;
    /**
     * The password of ECS instance.
     */
    password?: pulumi.Input<string>;
    /**
     * The period of ECS instance.Only effective when instanceChargeType is PrePaid. Default is 12. Unit is Month.
     */
    period?: pulumi.Input<number>;
    /**
     * The ProjectName of the ecs instance.
     */
    projectName?: pulumi.Input<string>;
    /**
     * The secondary networkInterface detail collection of ECS instance.
     */
    secondaryNetworkInterfaces?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceSecondaryNetworkInterface>[]>;
    /**
     * The security enhancement strategy of ECS instance. The value can be Active or InActive. Default is Active.When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignoreChanges ignore changes in fields.
     */
    securityEnhancementStrategy?: pulumi.Input<string>;
    /**
     * The security group ID set of primary networkInterface.
     */
    securityGroupIds: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The spot strategy will autoremove instance in some conditions.Please make sure you can maintain instance lifecycle before auto remove.The spot strategy of ECS instance, the value can be `NoSpot` or `SpotAsPriceGo`.
     */
    spotStrategy?: pulumi.Input<string>;
    /**
     * The subnet ID of primary networkInterface.
     */
    subnetId: pulumi.Input<string>;
    /**
     * The size of system volume. The value range of the system volume size is ESSD_PL0: 20~2048, ESSD_FlexPL: 20~2048, PTSSD: 10~500.
     */
    systemVolumeSize: pulumi.Input<number>;
    /**
     * The type of system volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
     */
    systemVolumeType: pulumi.Input<string>;
    /**
     * Tags.
     */
    tags?: pulumi.Input<pulumi.Input<inputs.ecs.InstanceTag>[]>;
    /**
     * The user data of ECS instance, this field must be encrypted with base64.
     */
    userData?: pulumi.Input<string>;
    /**
     * The available zone ID of ECS instance.
     */
    zoneId?: pulumi.Input<string>;
}
