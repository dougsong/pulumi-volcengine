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

__all__ = ['NodeArgs', 'Node']

@pulumi.input_type
class NodeArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 additional_container_storage_enabled: Optional[pulumi.Input[bool]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 container_storage_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 initialize_script: Optional[pulumi.Input[str]] = None,
                 keep_instance_name: Optional[pulumi.Input[bool]] = None,
                 kubernetes_config: Optional[pulumi.Input['NodeKubernetesConfigArgs']] = None):
        """
        The set of arguments for constructing a Node resource.
        :param pulumi.Input[str] cluster_id: The cluster id.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[bool] additional_container_storage_enabled: The flag of additional container storage enable, the value is `true` or `false`.
        :param pulumi.Input[str] client_token: The client token.
        :param pulumi.Input[str] container_storage_path: The container storage path.
        :param pulumi.Input[str] image_id: The ImageId of NodeConfig.
        :param pulumi.Input[str] initialize_script: The initializeScript of Node.
        :param pulumi.Input[bool] keep_instance_name: The flag of keep instance name, the value is `true` or `false`.
        :param pulumi.Input['NodeKubernetesConfigArgs'] kubernetes_config: The KubernetesConfig of Node.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        pulumi.set(__self__, "instance_id", instance_id)
        if additional_container_storage_enabled is not None:
            pulumi.set(__self__, "additional_container_storage_enabled", additional_container_storage_enabled)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if container_storage_path is not None:
            pulumi.set(__self__, "container_storage_path", container_storage_path)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if initialize_script is not None:
            pulumi.set(__self__, "initialize_script", initialize_script)
        if keep_instance_name is not None:
            pulumi.set(__self__, "keep_instance_name", keep_instance_name)
        if kubernetes_config is not None:
            pulumi.set(__self__, "kubernetes_config", kubernetes_config)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The cluster id.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="additionalContainerStorageEnabled")
    def additional_container_storage_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The flag of additional container storage enable, the value is `true` or `false`.
        """
        return pulumi.get(self, "additional_container_storage_enabled")

    @additional_container_storage_enabled.setter
    def additional_container_storage_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "additional_container_storage_enabled", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[str]]:
        """
        The client token.
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter(name="containerStoragePath")
    def container_storage_path(self) -> Optional[pulumi.Input[str]]:
        """
        The container storage path.
        """
        return pulumi.get(self, "container_storage_path")

    @container_storage_path.setter
    def container_storage_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_storage_path", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ImageId of NodeConfig.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="initializeScript")
    def initialize_script(self) -> Optional[pulumi.Input[str]]:
        """
        The initializeScript of Node.
        """
        return pulumi.get(self, "initialize_script")

    @initialize_script.setter
    def initialize_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initialize_script", value)

    @property
    @pulumi.getter(name="keepInstanceName")
    def keep_instance_name(self) -> Optional[pulumi.Input[bool]]:
        """
        The flag of keep instance name, the value is `true` or `false`.
        """
        return pulumi.get(self, "keep_instance_name")

    @keep_instance_name.setter
    def keep_instance_name(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_instance_name", value)

    @property
    @pulumi.getter(name="kubernetesConfig")
    def kubernetes_config(self) -> Optional[pulumi.Input['NodeKubernetesConfigArgs']]:
        """
        The KubernetesConfig of Node.
        """
        return pulumi.get(self, "kubernetes_config")

    @kubernetes_config.setter
    def kubernetes_config(self, value: Optional[pulumi.Input['NodeKubernetesConfigArgs']]):
        pulumi.set(self, "kubernetes_config", value)


@pulumi.input_type
class _NodeState:
    def __init__(__self__, *,
                 additional_container_storage_enabled: Optional[pulumi.Input[bool]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 container_storage_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 initialize_script: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 keep_instance_name: Optional[pulumi.Input[bool]] = None,
                 kubernetes_config: Optional[pulumi.Input['NodeKubernetesConfigArgs']] = None,
                 node_pool_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Node resources.
        :param pulumi.Input[bool] additional_container_storage_enabled: The flag of additional container storage enable, the value is `true` or `false`.
        :param pulumi.Input[str] client_token: The client token.
        :param pulumi.Input[str] cluster_id: The cluster id.
        :param pulumi.Input[str] container_storage_path: The container storage path.
        :param pulumi.Input[str] image_id: The ImageId of NodeConfig.
        :param pulumi.Input[str] initialize_script: The initializeScript of Node.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[bool] keep_instance_name: The flag of keep instance name, the value is `true` or `false`.
        :param pulumi.Input['NodeKubernetesConfigArgs'] kubernetes_config: The KubernetesConfig of Node.
        :param pulumi.Input[str] node_pool_id: The node pool id.
        """
        if additional_container_storage_enabled is not None:
            pulumi.set(__self__, "additional_container_storage_enabled", additional_container_storage_enabled)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if container_storage_path is not None:
            pulumi.set(__self__, "container_storage_path", container_storage_path)
        if image_id is not None:
            pulumi.set(__self__, "image_id", image_id)
        if initialize_script is not None:
            pulumi.set(__self__, "initialize_script", initialize_script)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if keep_instance_name is not None:
            pulumi.set(__self__, "keep_instance_name", keep_instance_name)
        if kubernetes_config is not None:
            pulumi.set(__self__, "kubernetes_config", kubernetes_config)
        if node_pool_id is not None:
            pulumi.set(__self__, "node_pool_id", node_pool_id)

    @property
    @pulumi.getter(name="additionalContainerStorageEnabled")
    def additional_container_storage_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        The flag of additional container storage enable, the value is `true` or `false`.
        """
        return pulumi.get(self, "additional_container_storage_enabled")

    @additional_container_storage_enabled.setter
    def additional_container_storage_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "additional_container_storage_enabled", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[str]]:
        """
        The client token.
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster id.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter(name="containerStoragePath")
    def container_storage_path(self) -> Optional[pulumi.Input[str]]:
        """
        The container storage path.
        """
        return pulumi.get(self, "container_storage_path")

    @container_storage_path.setter
    def container_storage_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "container_storage_path", value)

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ImageId of NodeConfig.
        """
        return pulumi.get(self, "image_id")

    @image_id.setter
    def image_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "image_id", value)

    @property
    @pulumi.getter(name="initializeScript")
    def initialize_script(self) -> Optional[pulumi.Input[str]]:
        """
        The initializeScript of Node.
        """
        return pulumi.get(self, "initialize_script")

    @initialize_script.setter
    def initialize_script(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "initialize_script", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The instance id.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="keepInstanceName")
    def keep_instance_name(self) -> Optional[pulumi.Input[bool]]:
        """
        The flag of keep instance name, the value is `true` or `false`.
        """
        return pulumi.get(self, "keep_instance_name")

    @keep_instance_name.setter
    def keep_instance_name(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_instance_name", value)

    @property
    @pulumi.getter(name="kubernetesConfig")
    def kubernetes_config(self) -> Optional[pulumi.Input['NodeKubernetesConfigArgs']]:
        """
        The KubernetesConfig of Node.
        """
        return pulumi.get(self, "kubernetes_config")

    @kubernetes_config.setter
    def kubernetes_config(self, value: Optional[pulumi.Input['NodeKubernetesConfigArgs']]):
        pulumi.set(self, "kubernetes_config", value)

    @property
    @pulumi.getter(name="nodePoolId")
    def node_pool_id(self) -> Optional[pulumi.Input[str]]:
        """
        The node pool id.
        """
        return pulumi.get(self, "node_pool_id")

    @node_pool_id.setter
    def node_pool_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "node_pool_id", value)


class Node(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_container_storage_enabled: Optional[pulumi.Input[bool]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 container_storage_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 initialize_script: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 keep_instance_name: Optional[pulumi.Input[bool]] = None,
                 kubernetes_config: Optional[pulumi.Input[pulumi.InputType['NodeKubernetesConfigArgs']]] = None,
                 __props__=None):
        """
        Provides a resource to manage vke node
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vke.Node("foo",
            additional_container_storage_enabled=False,
            cluster_id="ccahbr0nqtofhiuuuajn0",
            container_storage_path="",
            instance_id="i-ybrfa2vu2t7grbv8qa0j",
            keep_instance_name=True)
        ```

        ## Import

        VKE node can be imported using the node id, e.g.

        ```sh
         $ pulumi import volcengine:vke/node:Node default nc5t5epmrsf****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] additional_container_storage_enabled: The flag of additional container storage enable, the value is `true` or `false`.
        :param pulumi.Input[str] client_token: The client token.
        :param pulumi.Input[str] cluster_id: The cluster id.
        :param pulumi.Input[str] container_storage_path: The container storage path.
        :param pulumi.Input[str] image_id: The ImageId of NodeConfig.
        :param pulumi.Input[str] initialize_script: The initializeScript of Node.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[bool] keep_instance_name: The flag of keep instance name, the value is `true` or `false`.
        :param pulumi.Input[pulumi.InputType['NodeKubernetesConfigArgs']] kubernetes_config: The KubernetesConfig of Node.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NodeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage vke node
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vke.Node("foo",
            additional_container_storage_enabled=False,
            cluster_id="ccahbr0nqtofhiuuuajn0",
            container_storage_path="",
            instance_id="i-ybrfa2vu2t7grbv8qa0j",
            keep_instance_name=True)
        ```

        ## Import

        VKE node can be imported using the node id, e.g.

        ```sh
         $ pulumi import volcengine:vke/node:Node default nc5t5epmrsf****
        ```

        :param str resource_name: The name of the resource.
        :param NodeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NodeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_container_storage_enabled: Optional[pulumi.Input[bool]] = None,
                 client_token: Optional[pulumi.Input[str]] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 container_storage_path: Optional[pulumi.Input[str]] = None,
                 image_id: Optional[pulumi.Input[str]] = None,
                 initialize_script: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 keep_instance_name: Optional[pulumi.Input[bool]] = None,
                 kubernetes_config: Optional[pulumi.Input[pulumi.InputType['NodeKubernetesConfigArgs']]] = None,
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
            __props__ = NodeArgs.__new__(NodeArgs)

            __props__.__dict__["additional_container_storage_enabled"] = additional_container_storage_enabled
            __props__.__dict__["client_token"] = client_token
            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["container_storage_path"] = container_storage_path
            __props__.__dict__["image_id"] = image_id
            __props__.__dict__["initialize_script"] = initialize_script
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["keep_instance_name"] = keep_instance_name
            __props__.__dict__["kubernetes_config"] = kubernetes_config
            __props__.__dict__["node_pool_id"] = None
        super(Node, __self__).__init__(
            'volcengine:vke/node:Node',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            additional_container_storage_enabled: Optional[pulumi.Input[bool]] = None,
            client_token: Optional[pulumi.Input[str]] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            container_storage_path: Optional[pulumi.Input[str]] = None,
            image_id: Optional[pulumi.Input[str]] = None,
            initialize_script: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            keep_instance_name: Optional[pulumi.Input[bool]] = None,
            kubernetes_config: Optional[pulumi.Input[pulumi.InputType['NodeKubernetesConfigArgs']]] = None,
            node_pool_id: Optional[pulumi.Input[str]] = None) -> 'Node':
        """
        Get an existing Node resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] additional_container_storage_enabled: The flag of additional container storage enable, the value is `true` or `false`.
        :param pulumi.Input[str] client_token: The client token.
        :param pulumi.Input[str] cluster_id: The cluster id.
        :param pulumi.Input[str] container_storage_path: The container storage path.
        :param pulumi.Input[str] image_id: The ImageId of NodeConfig.
        :param pulumi.Input[str] initialize_script: The initializeScript of Node.
        :param pulumi.Input[str] instance_id: The instance id.
        :param pulumi.Input[bool] keep_instance_name: The flag of keep instance name, the value is `true` or `false`.
        :param pulumi.Input[pulumi.InputType['NodeKubernetesConfigArgs']] kubernetes_config: The KubernetesConfig of Node.
        :param pulumi.Input[str] node_pool_id: The node pool id.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NodeState.__new__(_NodeState)

        __props__.__dict__["additional_container_storage_enabled"] = additional_container_storage_enabled
        __props__.__dict__["client_token"] = client_token
        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["container_storage_path"] = container_storage_path
        __props__.__dict__["image_id"] = image_id
        __props__.__dict__["initialize_script"] = initialize_script
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["keep_instance_name"] = keep_instance_name
        __props__.__dict__["kubernetes_config"] = kubernetes_config
        __props__.__dict__["node_pool_id"] = node_pool_id
        return Node(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="additionalContainerStorageEnabled")
    def additional_container_storage_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        The flag of additional container storage enable, the value is `true` or `false`.
        """
        return pulumi.get(self, "additional_container_storage_enabled")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[str]:
        """
        The client token.
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The cluster id.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter(name="containerStoragePath")
    def container_storage_path(self) -> pulumi.Output[str]:
        """
        The container storage path.
        """
        return pulumi.get(self, "container_storage_path")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> pulumi.Output[str]:
        """
        The ImageId of NodeConfig.
        """
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="initializeScript")
    def initialize_script(self) -> pulumi.Output[Optional[str]]:
        """
        The initializeScript of Node.
        """
        return pulumi.get(self, "initialize_script")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The instance id.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="keepInstanceName")
    def keep_instance_name(self) -> pulumi.Output[Optional[bool]]:
        """
        The flag of keep instance name, the value is `true` or `false`.
        """
        return pulumi.get(self, "keep_instance_name")

    @property
    @pulumi.getter(name="kubernetesConfig")
    def kubernetes_config(self) -> pulumi.Output[Optional['outputs.NodeKubernetesConfig']]:
        """
        The KubernetesConfig of Node.
        """
        return pulumi.get(self, "kubernetes_config")

    @property
    @pulumi.getter(name="nodePoolId")
    def node_pool_id(self) -> pulumi.Output[str]:
        """
        The node pool id.
        """
        return pulumi.get(self, "node_pool_id")

