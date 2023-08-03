# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DefaultNodePoolBatchAttachArgs', 'DefaultNodePoolBatchAttach']

@pulumi.input_type
class DefaultNodePoolBatchAttachArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 default_node_pool_id: pulumi.Input[str],
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]]] = None,
                 kubernetes_config: Optional[pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs']] = None):
        """
        The set of arguments for constructing a DefaultNodePoolBatchAttach resource.
        :param pulumi.Input[str] cluster_id: The ClusterId of NodePool.
        :param pulumi.Input[str] default_node_pool_id: The default NodePool ID.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]] instances: The ECS InstanceIds add to NodePool.
        :param pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs'] kubernetes_config: The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "default_node_pool_id", default_node_pool_id)
        if instances is not None:
            pulumi.set(__self__, "instances", instances)
        if kubernetes_config is not None:
            pulumi.set(__self__, "kubernetes_config", kubernetes_config)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The ClusterId of NodePool.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="defaultNodePoolId")
    def default_node_pool_id(self) -> pulumi.Input[str]:
        """
        The default NodePool ID.
        """
        return pulumi.get(self, "default_node_pool_id")

    @default_node_pool_id.setter
    def default_node_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_node_pool_id", value)

    @property
    @pulumi.getter
    def instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]]]:
        """
        The ECS InstanceIds add to NodePool.
        """
        return pulumi.get(self, "instances")

    @instances.setter
    def instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]]]):
        pulumi.set(self, "instances", value)

    @property
    @pulumi.getter(name="kubernetesConfig")
    def kubernetes_config(self) -> Optional[pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs']]:
        """
        The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        """
        return pulumi.get(self, "kubernetes_config")

    @kubernetes_config.setter
    def kubernetes_config(self, value: Optional[pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs']]):
        pulumi.set(self, "kubernetes_config", value)


@pulumi.input_type
class _DefaultNodePoolBatchAttachState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 default_node_pool_id: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]]] = None,
                 is_import: Optional[pulumi.Input[bool]] = None,
                 kubernetes_config: Optional[pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs']] = None,
                 node_configs: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachNodeConfigArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering DefaultNodePoolBatchAttach resources.
        :param pulumi.Input[str] cluster_id: The ClusterId of NodePool.
        :param pulumi.Input[str] default_node_pool_id: The default NodePool ID.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]] instances: The ECS InstanceIds add to NodePool.
        :param pulumi.Input[bool] is_import: Is import of the DefaultNodePool. It only works when imported, set to true.
        :param pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs'] kubernetes_config: The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachNodeConfigArgs']]] node_configs: The Config of NodePool.
        :param pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachTagArgs']]] tags: Tags.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if default_node_pool_id is not None:
            pulumi.set(__self__, "default_node_pool_id", default_node_pool_id)
        if instances is not None:
            pulumi.set(__self__, "instances", instances)
        if is_import is not None:
            pulumi.set(__self__, "is_import", is_import)
        if kubernetes_config is not None:
            pulumi.set(__self__, "kubernetes_config", kubernetes_config)
        if node_configs is not None:
            pulumi.set(__self__, "node_configs", node_configs)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ClusterId of NodePool.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="defaultNodePoolId")
    def default_node_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The default NodePool ID.
        """
        return pulumi.get(self, "default_node_pool_id")

    @default_node_pool_id.setter
    def default_node_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_node_pool_id", value)

    @property
    @pulumi.getter
    def instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]]]:
        """
        The ECS InstanceIds add to NodePool.
        """
        return pulumi.get(self, "instances")

    @instances.setter
    def instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachInstanceArgs']]]]):
        pulumi.set(self, "instances", value)

    @property
    @pulumi.getter(name="isImport")
    def is_import(self) -> Optional[pulumi.Input[bool]]:
        """
        Is import of the DefaultNodePool. It only works when imported, set to true.
        """
        return pulumi.get(self, "is_import")

    @is_import.setter
    def is_import(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_import", value)

    @property
    @pulumi.getter(name="kubernetesConfig")
    def kubernetes_config(self) -> Optional[pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs']]:
        """
        The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        """
        return pulumi.get(self, "kubernetes_config")

    @kubernetes_config.setter
    def kubernetes_config(self, value: Optional[pulumi.Input['DefaultNodePoolBatchAttachKubernetesConfigArgs']]):
        pulumi.set(self, "kubernetes_config", value)

    @property
    @pulumi.getter(name="nodeConfigs")
    def node_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachNodeConfigArgs']]]]:
        """
        The Config of NodePool.
        """
        return pulumi.get(self, "node_configs")

    @node_configs.setter
    def node_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachNodeConfigArgs']]]]):
        pulumi.set(self, "node_configs", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['DefaultNodePoolBatchAttachTagArgs']]]]):
        pulumi.set(self, "tags", value)


class DefaultNodePoolBatchAttach(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 default_node_pool_id: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachInstanceArgs']]]]] = None,
                 kubernetes_config: Optional[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachKubernetesConfigArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to manage vke default node pool batch attach
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-project1",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-subnet-test-2",
            cidr_block="172.16.0.0/24",
            zone_id="cn-beijing-a",
            vpc_id=foo_vpc.id)
        foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
            vpc_id=foo_vpc.id,
            security_group_name="acc-test-security-group2")
        foo_instance = volcengine.ecs.Instance("fooInstance",
            image_id="image-ybqi99s7yq8rx7mnk44b",
            instance_type="ecs.g1ie.large",
            instance_name="acc-test-ecs-name2",
            password="93f0cb0614Aab12",
            instance_charge_type="PostPaid",
            system_volume_type="ESSD_PL0",
            system_volume_size=40,
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id])
        foo_cluster = volcengine.vke.Cluster("fooCluster",
            description="created by terraform",
            delete_protection_enabled=False,
            cluster_config=volcengine.vke.ClusterClusterConfigArgs(
                subnet_ids=[foo_subnet.id],
                api_server_public_access_enabled=True,
                api_server_public_access_config=volcengine.vke.ClusterClusterConfigApiServerPublicAccessConfigArgs(
                    public_access_network_config=volcengine.vke.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs(
                        billing_type="PostPaidByBandwidth",
                        bandwidth=1,
                    ),
                ),
                resource_public_access_default_enabled=True,
            ),
            pods_config=volcengine.vke.ClusterPodsConfigArgs(
                pod_network_mode="VpcCniShared",
                vpc_cni_config=volcengine.vke.ClusterPodsConfigVpcCniConfigArgs(
                    subnet_ids=[foo_subnet.id],
                ),
            ),
            services_config=volcengine.vke.ClusterServicesConfigArgs(
                service_cidrsv4s=["172.30.0.0/18"],
            ),
            tags=[volcengine.vke.ClusterTagArgs(
                key="tf-k1",
                value="tf-v1",
            )])
        default_default_node_pool = volcengine.vke.DefaultNodePool("defaultDefaultNodePool",
            cluster_id=foo_cluster.id,
            node_config=volcengine.vke.DefaultNodePoolNodeConfigArgs(
                security=volcengine.vke.DefaultNodePoolNodeConfigSecurityArgs(
                    login=volcengine.vke.DefaultNodePoolNodeConfigSecurityLoginArgs(
                        password="amw4WTdVcTRJVVFsUXpVTw==",
                    ),
                    security_group_ids=[foo_security_group.id],
                    security_strategies=["Hids"],
                ),
                initialize_script="ISMvYmluL2Jhc2gKZWNobyAx",
            ),
            kubernetes_config=volcengine.vke.DefaultNodePoolKubernetesConfigArgs(
                labels=[
                    volcengine.vke.DefaultNodePoolKubernetesConfigLabelArgs(
                        key="tf-key1",
                        value="tf-value1",
                    ),
                    volcengine.vke.DefaultNodePoolKubernetesConfigLabelArgs(
                        key="tf-key2",
                        value="tf-value2",
                    ),
                ],
                taints=[
                    volcengine.vke.DefaultNodePoolKubernetesConfigTaintArgs(
                        key="tf-key3",
                        value="tf-value3",
                        effect="NoSchedule",
                    ),
                    volcengine.vke.DefaultNodePoolKubernetesConfigTaintArgs(
                        key="tf-key4",
                        value="tf-value4",
                        effect="NoSchedule",
                    ),
                ],
                cordon=True,
            ),
            tags=[volcengine.vke.DefaultNodePoolTagArgs(
                key="k1",
                value="v1",
            )])
        default_default_node_pool_batch_attach = volcengine.vke.DefaultNodePoolBatchAttach("defaultDefaultNodePoolBatchAttach",
            cluster_id=foo_cluster.id,
            default_node_pool_id=default_default_node_pool.id,
            instances=[volcengine.vke.DefaultNodePoolBatchAttachInstanceArgs(
                instance_id=foo_instance.id,
                keep_instance_name=True,
                additional_container_storage_enabled=False,
            )],
            kubernetes_config=volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigArgs(
                labels=[
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigLabelArgs(
                        key="tf-key1",
                        value="tf-value1",
                    ),
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigLabelArgs(
                        key="tf-key2",
                        value="tf-value2",
                    ),
                ],
                taints=[
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigTaintArgs(
                        key="tf-key3",
                        value="tf-value3",
                        effect="NoSchedule",
                    ),
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigTaintArgs(
                        key="tf-key4",
                        value="tf-value4",
                        effect="NoSchedule",
                    ),
                ],
                cordon=True,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ClusterId of NodePool.
        :param pulumi.Input[str] default_node_pool_id: The default NodePool ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachInstanceArgs']]]] instances: The ECS InstanceIds add to NodePool.
        :param pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachKubernetesConfigArgs']] kubernetes_config: The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefaultNodePoolBatchAttachArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage vke default node pool batch attach
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo_vpc = volcengine.vpc.Vpc("fooVpc",
            vpc_name="acc-test-project1",
            cidr_block="172.16.0.0/16")
        foo_subnet = volcengine.vpc.Subnet("fooSubnet",
            subnet_name="acc-subnet-test-2",
            cidr_block="172.16.0.0/24",
            zone_id="cn-beijing-a",
            vpc_id=foo_vpc.id)
        foo_security_group = volcengine.vpc.SecurityGroup("fooSecurityGroup",
            vpc_id=foo_vpc.id,
            security_group_name="acc-test-security-group2")
        foo_instance = volcengine.ecs.Instance("fooInstance",
            image_id="image-ybqi99s7yq8rx7mnk44b",
            instance_type="ecs.g1ie.large",
            instance_name="acc-test-ecs-name2",
            password="93f0cb0614Aab12",
            instance_charge_type="PostPaid",
            system_volume_type="ESSD_PL0",
            system_volume_size=40,
            subnet_id=foo_subnet.id,
            security_group_ids=[foo_security_group.id])
        foo_cluster = volcengine.vke.Cluster("fooCluster",
            description="created by terraform",
            delete_protection_enabled=False,
            cluster_config=volcengine.vke.ClusterClusterConfigArgs(
                subnet_ids=[foo_subnet.id],
                api_server_public_access_enabled=True,
                api_server_public_access_config=volcengine.vke.ClusterClusterConfigApiServerPublicAccessConfigArgs(
                    public_access_network_config=volcengine.vke.ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs(
                        billing_type="PostPaidByBandwidth",
                        bandwidth=1,
                    ),
                ),
                resource_public_access_default_enabled=True,
            ),
            pods_config=volcengine.vke.ClusterPodsConfigArgs(
                pod_network_mode="VpcCniShared",
                vpc_cni_config=volcengine.vke.ClusterPodsConfigVpcCniConfigArgs(
                    subnet_ids=[foo_subnet.id],
                ),
            ),
            services_config=volcengine.vke.ClusterServicesConfigArgs(
                service_cidrsv4s=["172.30.0.0/18"],
            ),
            tags=[volcengine.vke.ClusterTagArgs(
                key="tf-k1",
                value="tf-v1",
            )])
        default_default_node_pool = volcengine.vke.DefaultNodePool("defaultDefaultNodePool",
            cluster_id=foo_cluster.id,
            node_config=volcengine.vke.DefaultNodePoolNodeConfigArgs(
                security=volcengine.vke.DefaultNodePoolNodeConfigSecurityArgs(
                    login=volcengine.vke.DefaultNodePoolNodeConfigSecurityLoginArgs(
                        password="amw4WTdVcTRJVVFsUXpVTw==",
                    ),
                    security_group_ids=[foo_security_group.id],
                    security_strategies=["Hids"],
                ),
                initialize_script="ISMvYmluL2Jhc2gKZWNobyAx",
            ),
            kubernetes_config=volcengine.vke.DefaultNodePoolKubernetesConfigArgs(
                labels=[
                    volcengine.vke.DefaultNodePoolKubernetesConfigLabelArgs(
                        key="tf-key1",
                        value="tf-value1",
                    ),
                    volcengine.vke.DefaultNodePoolKubernetesConfigLabelArgs(
                        key="tf-key2",
                        value="tf-value2",
                    ),
                ],
                taints=[
                    volcengine.vke.DefaultNodePoolKubernetesConfigTaintArgs(
                        key="tf-key3",
                        value="tf-value3",
                        effect="NoSchedule",
                    ),
                    volcengine.vke.DefaultNodePoolKubernetesConfigTaintArgs(
                        key="tf-key4",
                        value="tf-value4",
                        effect="NoSchedule",
                    ),
                ],
                cordon=True,
            ),
            tags=[volcengine.vke.DefaultNodePoolTagArgs(
                key="k1",
                value="v1",
            )])
        default_default_node_pool_batch_attach = volcengine.vke.DefaultNodePoolBatchAttach("defaultDefaultNodePoolBatchAttach",
            cluster_id=foo_cluster.id,
            default_node_pool_id=default_default_node_pool.id,
            instances=[volcengine.vke.DefaultNodePoolBatchAttachInstanceArgs(
                instance_id=foo_instance.id,
                keep_instance_name=True,
                additional_container_storage_enabled=False,
            )],
            kubernetes_config=volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigArgs(
                labels=[
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigLabelArgs(
                        key="tf-key1",
                        value="tf-value1",
                    ),
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigLabelArgs(
                        key="tf-key2",
                        value="tf-value2",
                    ),
                ],
                taints=[
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigTaintArgs(
                        key="tf-key3",
                        value="tf-value3",
                        effect="NoSchedule",
                    ),
                    volcengine.vke.DefaultNodePoolBatchAttachKubernetesConfigTaintArgs(
                        key="tf-key4",
                        value="tf-value4",
                        effect="NoSchedule",
                    ),
                ],
                cordon=True,
            ))
        ```

        :param str resource_name: The name of the resource.
        :param DefaultNodePoolBatchAttachArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultNodePoolBatchAttachArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 default_node_pool_id: Optional[pulumi.Input[str]] = None,
                 instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachInstanceArgs']]]]] = None,
                 kubernetes_config: Optional[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachKubernetesConfigArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultNodePoolBatchAttachArgs.__new__(DefaultNodePoolBatchAttachArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            if default_node_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'default_node_pool_id'")
            __props__.__dict__["default_node_pool_id"] = default_node_pool_id
            __props__.__dict__["instances"] = instances
            __props__.__dict__["kubernetes_config"] = kubernetes_config
            __props__.__dict__["is_import"] = None
            __props__.__dict__["node_configs"] = None
            __props__.__dict__["tags"] = None
        super(DefaultNodePoolBatchAttach, __self__).__init__(
            'volcengine:vke/defaultNodePoolBatchAttach:DefaultNodePoolBatchAttach',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            default_node_pool_id: Optional[pulumi.Input[str]] = None,
            instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachInstanceArgs']]]]] = None,
            is_import: Optional[pulumi.Input[bool]] = None,
            kubernetes_config: Optional[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachKubernetesConfigArgs']]] = None,
            node_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachNodeConfigArgs']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachTagArgs']]]]] = None) -> 'DefaultNodePoolBatchAttach':
        """
        Get an existing DefaultNodePoolBatchAttach resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The ClusterId of NodePool.
        :param pulumi.Input[str] default_node_pool_id: The default NodePool ID.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachInstanceArgs']]]] instances: The ECS InstanceIds add to NodePool.
        :param pulumi.Input[bool] is_import: Is import of the DefaultNodePool. It only works when imported, set to true.
        :param pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachKubernetesConfigArgs']] kubernetes_config: The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachNodeConfigArgs']]]] node_configs: The Config of NodePool.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['DefaultNodePoolBatchAttachTagArgs']]]] tags: Tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultNodePoolBatchAttachState.__new__(_DefaultNodePoolBatchAttachState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["default_node_pool_id"] = default_node_pool_id
        __props__.__dict__["instances"] = instances
        __props__.__dict__["is_import"] = is_import
        __props__.__dict__["kubernetes_config"] = kubernetes_config
        __props__.__dict__["node_configs"] = node_configs
        __props__.__dict__["tags"] = tags
        return DefaultNodePoolBatchAttach(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The ClusterId of NodePool.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="defaultNodePoolId")
    def default_node_pool_id(self) -> pulumi.Output[str]:
        """
        The default NodePool ID.
        """
        return pulumi.get(self, "default_node_pool_id")

    @property
    @pulumi.getter
    def instances(self) -> pulumi.Output[Optional[Sequence['outputs.DefaultNodePoolBatchAttachInstance']]]:
        """
        The ECS InstanceIds add to NodePool.
        """
        return pulumi.get(self, "instances")

    @property
    @pulumi.getter(name="isImport")
    def is_import(self) -> pulumi.Output[bool]:
        """
        Is import of the DefaultNodePool. It only works when imported, set to true.
        """
        return pulumi.get(self, "is_import")

    @property
    @pulumi.getter(name="kubernetesConfig")
    def kubernetes_config(self) -> pulumi.Output[Optional['outputs.DefaultNodePoolBatchAttachKubernetesConfig']]:
        """
        The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.
        """
        return pulumi.get(self, "kubernetes_config")

    @property
    @pulumi.getter(name="nodeConfigs")
    def node_configs(self) -> pulumi.Output[Sequence['outputs.DefaultNodePoolBatchAttachNodeConfig']]:
        """
        The Config of NodePool.
        """
        return pulumi.get(self, "node_configs")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence['outputs.DefaultNodePoolBatchAttachTag']]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

