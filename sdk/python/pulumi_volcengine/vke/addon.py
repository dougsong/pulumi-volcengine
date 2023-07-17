# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AddonArgs', 'Addon']

@pulumi.input_type
class AddonArgs:
    def __init__(__self__, *,
                 cluster_id: pulumi.Input[str],
                 config: Optional[pulumi.Input[str]] = None,
                 deploy_mode: Optional[pulumi.Input[str]] = None,
                 deploy_node_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Addon resource.
        :param pulumi.Input[str] cluster_id: The cluster id of the addon.
        :param pulumi.Input[str] config: The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        :param pulumi.Input[str] deploy_mode: The deploy mode.
        :param pulumi.Input[str] deploy_node_type: The deploy node type.
        :param pulumi.Input[str] name: The name of the addon.
        :param pulumi.Input[str] version: The version info of the cluster.
        """
        pulumi.set(__self__, "cluster_id", cluster_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if deploy_mode is not None:
            pulumi.set(__self__, "deploy_mode", deploy_mode)
        if deploy_node_type is not None:
            pulumi.set(__self__, "deploy_node_type", deploy_node_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Input[str]:
        """
        The cluster id of the addon.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="deployMode")
    def deploy_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The deploy mode.
        """
        return pulumi.get(self, "deploy_mode")

    @deploy_mode.setter
    def deploy_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_mode", value)

    @property
    @pulumi.getter(name="deployNodeType")
    def deploy_node_type(self) -> Optional[pulumi.Input[str]]:
        """
        The deploy node type.
        """
        return pulumi.get(self, "deploy_node_type")

    @deploy_node_type.setter
    def deploy_node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_node_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the addon.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version info of the cluster.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


@pulumi.input_type
class _AddonState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 deploy_mode: Optional[pulumi.Input[str]] = None,
                 deploy_node_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Addon resources.
        :param pulumi.Input[str] cluster_id: The cluster id of the addon.
        :param pulumi.Input[str] config: The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        :param pulumi.Input[str] deploy_mode: The deploy mode.
        :param pulumi.Input[str] deploy_node_type: The deploy node type.
        :param pulumi.Input[str] name: The name of the addon.
        :param pulumi.Input[str] version: The version info of the cluster.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if deploy_mode is not None:
            pulumi.set(__self__, "deploy_mode", deploy_mode)
        if deploy_node_type is not None:
            pulumi.set(__self__, "deploy_node_type", deploy_node_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The cluster id of the addon.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[str]]:
        """
        The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="deployMode")
    def deploy_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The deploy mode.
        """
        return pulumi.get(self, "deploy_mode")

    @deploy_mode.setter
    def deploy_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_mode", value)

    @property
    @pulumi.getter(name="deployNodeType")
    def deploy_node_type(self) -> Optional[pulumi.Input[str]]:
        """
        The deploy node type.
        """
        return pulumi.get(self, "deploy_node_type")

    @deploy_node_type.setter
    def deploy_node_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deploy_node_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the addon.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        The version info of the cluster.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Addon(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 deploy_mode: Optional[pulumi.Input[str]] = None,
                 deploy_node_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage vke addon
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vke.Addon("foo",
            cluster_id="cccctv1vqtofp49d96ujg",
            config="{\"xxx\":\"true\"}",
            deploy_mode="Unmanaged",
            deploy_node_type="Node",
            version="v0.1.3")
        ```

        ## Import

        VkeAddon can be imported using the clusterId:Name, e.g.

        ```sh
         $ pulumi import volcengine:vke/addon:Addon default cc9l74mvqtofjnoj5****:nginx-ingress
        ```

         Notice Some kind of VKEAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy. If you want to remove it from terraform state, please use command $ terraform state rm volcengine_vke_addon.${name}

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The cluster id of the addon.
        :param pulumi.Input[str] config: The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        :param pulumi.Input[str] deploy_mode: The deploy mode.
        :param pulumi.Input[str] deploy_node_type: The deploy node type.
        :param pulumi.Input[str] name: The name of the addon.
        :param pulumi.Input[str] version: The version info of the cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AddonArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage vke addon
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vke.Addon("foo",
            cluster_id="cccctv1vqtofp49d96ujg",
            config="{\"xxx\":\"true\"}",
            deploy_mode="Unmanaged",
            deploy_node_type="Node",
            version="v0.1.3")
        ```

        ## Import

        VkeAddon can be imported using the clusterId:Name, e.g.

        ```sh
         $ pulumi import volcengine:vke/addon:Addon default cc9l74mvqtofjnoj5****:nginx-ingress
        ```

         Notice Some kind of VKEAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy. If you want to remove it from terraform state, please use command $ terraform state rm volcengine_vke_addon.${name}

        :param str resource_name: The name of the resource.
        :param AddonArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AddonArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[str]] = None,
                 deploy_mode: Optional[pulumi.Input[str]] = None,
                 deploy_node_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AddonArgs.__new__(AddonArgs)

            if cluster_id is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_id'")
            __props__.__dict__["cluster_id"] = cluster_id
            __props__.__dict__["config"] = config
            __props__.__dict__["deploy_mode"] = deploy_mode
            __props__.__dict__["deploy_node_type"] = deploy_node_type
            __props__.__dict__["name"] = name
            __props__.__dict__["version"] = version
        super(Addon, __self__).__init__(
            'volcengine:vke/addon:Addon',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            config: Optional[pulumi.Input[str]] = None,
            deploy_mode: Optional[pulumi.Input[str]] = None,
            deploy_node_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'Addon':
        """
        Get an existing Addon resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The cluster id of the addon.
        :param pulumi.Input[str] config: The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        :param pulumi.Input[str] deploy_mode: The deploy mode.
        :param pulumi.Input[str] deploy_node_type: The deploy node type.
        :param pulumi.Input[str] name: The name of the addon.
        :param pulumi.Input[str] version: The version info of the cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AddonState.__new__(_AddonState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["config"] = config
        __props__.__dict__["deploy_mode"] = deploy_mode
        __props__.__dict__["deploy_node_type"] = deploy_node_type
        __props__.__dict__["name"] = name
        __props__.__dict__["version"] = version
        return Addon(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The cluster id of the addon.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[str]:
        """
        The config info of addon. Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="deployMode")
    def deploy_mode(self) -> pulumi.Output[str]:
        """
        The deploy mode.
        """
        return pulumi.get(self, "deploy_mode")

    @property
    @pulumi.getter(name="deployNodeType")
    def deploy_node_type(self) -> pulumi.Output[str]:
        """
        The deploy node type.
        """
        return pulumi.get(self, "deploy_node_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the addon.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        The version info of the cluster.
        """
        return pulumi.get(self, "version")

