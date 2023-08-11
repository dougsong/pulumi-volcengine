# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ClusterArgs', 'Cluster']

@pulumi.input_type
class ClusterArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shared_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]]] = None,
                 vke_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]]] = None):
        """
        The set of arguments for constructing a Cluster resource.
        :param pulumi.Input[str] description: The description of the cluster.
        :param pulumi.Input[str] name: The name of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]] shared_configs: The configuration of the shared cluster.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]] vke_configs: The configuration of the vke cluster.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if shared_configs is not None:
            pulumi.set(__self__, "shared_configs", shared_configs)
        if vke_configs is not None:
            pulumi.set(__self__, "vke_configs", vke_configs)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cluster.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sharedConfigs")
    def shared_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]]]:
        """
        The configuration of the shared cluster.
        """
        return pulumi.get(self, "shared_configs")

    @shared_configs.setter
    def shared_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]]]):
        pulumi.set(self, "shared_configs", value)

    @property
    @pulumi.getter(name="vkeConfigs")
    def vke_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]]]:
        """
        The configuration of the vke cluster.
        """
        return pulumi.get(self, "vke_configs")

    @vke_configs.setter
    def vke_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]]]):
        pulumi.set(self, "vke_configs", value)


@pulumi.input_type
class _ClusterState:
    def __init__(__self__, *,
                 cluster_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shared_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]]] = None,
                 vke_configs: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]]] = None):
        """
        Input properties used for looking up and filtering Cluster resources.
        :param pulumi.Input[str] cluster_id: The id of the vke cluster.
        :param pulumi.Input[str] description: The description of the cluster.
        :param pulumi.Input[str] name: The name of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]] shared_configs: The configuration of the shared cluster.
        :param pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]] vke_configs: The configuration of the vke cluster.
        """
        if cluster_id is not None:
            pulumi.set(__self__, "cluster_id", cluster_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if shared_configs is not None:
            pulumi.set(__self__, "shared_configs", shared_configs)
        if vke_configs is not None:
            pulumi.set(__self__, "vke_configs", vke_configs)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the vke cluster.
        """
        return pulumi.get(self, "cluster_id")

    @cluster_id.setter
    def cluster_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cluster_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cluster.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="sharedConfigs")
    def shared_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]]]:
        """
        The configuration of the shared cluster.
        """
        return pulumi.get(self, "shared_configs")

    @shared_configs.setter
    def shared_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterSharedConfigArgs']]]]):
        pulumi.set(self, "shared_configs", value)

    @property
    @pulumi.getter(name="vkeConfigs")
    def vke_configs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]]]:
        """
        The configuration of the vke cluster.
        """
        return pulumi.get(self, "vke_configs")

    @vke_configs.setter
    def vke_configs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ClusterVkeConfigArgs']]]]):
        pulumi.set(self, "vke_configs", value)


class Cluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shared_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSharedConfigArgs']]]]] = None,
                 vke_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterVkeConfigArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage bioos cluster
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.bioos.Cluster("foo",
            description="test-description",
            shared_configs=[volcengine.bioos.ClusterSharedConfigArgs(
                enable=True,
            )])
        ```

        ## Import

        Cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:bioos/cluster:Cluster default *****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the cluster.
        :param pulumi.Input[str] name: The name of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSharedConfigArgs']]]] shared_configs: The configuration of the shared cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterVkeConfigArgs']]]] vke_configs: The configuration of the vke cluster.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ClusterArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage bioos cluster
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.bioos.Cluster("foo",
            description="test-description",
            shared_configs=[volcengine.bioos.ClusterSharedConfigArgs(
                enable=True,
            )])
        ```

        ## Import

        Cluster can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:bioos/cluster:Cluster default *****
        ```

        :param str resource_name: The name of the resource.
        :param ClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 shared_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSharedConfigArgs']]]]] = None,
                 vke_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterVkeConfigArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClusterArgs.__new__(ClusterArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["shared_configs"] = shared_configs
            __props__.__dict__["vke_configs"] = vke_configs
            __props__.__dict__["cluster_id"] = None
        super(Cluster, __self__).__init__(
            'volcengine:bioos/cluster:Cluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cluster_id: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            shared_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSharedConfigArgs']]]]] = None,
            vke_configs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterVkeConfigArgs']]]]] = None) -> 'Cluster':
        """
        Get an existing Cluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_id: The id of the vke cluster.
        :param pulumi.Input[str] description: The description of the cluster.
        :param pulumi.Input[str] name: The name of the cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterSharedConfigArgs']]]] shared_configs: The configuration of the shared cluster.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ClusterVkeConfigArgs']]]] vke_configs: The configuration of the vke cluster.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClusterState.__new__(_ClusterState)

        __props__.__dict__["cluster_id"] = cluster_id
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["shared_configs"] = shared_configs
        __props__.__dict__["vke_configs"] = vke_configs
        return Cluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterId")
    def cluster_id(self) -> pulumi.Output[str]:
        """
        The id of the vke cluster.
        """
        return pulumi.get(self, "cluster_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the cluster.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="sharedConfigs")
    def shared_configs(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterSharedConfig']]]:
        """
        The configuration of the shared cluster.
        """
        return pulumi.get(self, "shared_configs")

    @property
    @pulumi.getter(name="vkeConfigs")
    def vke_configs(self) -> pulumi.Output[Optional[Sequence['outputs.ClusterVkeConfig']]]:
        """
        The configuration of the vke cluster.
        """
        return pulumi.get(self, "vke_configs")

