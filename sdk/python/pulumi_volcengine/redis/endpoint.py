# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['EndpointArgs', 'Endpoint']

@pulumi.input_type
class EndpointArgs:
    def __init__(__self__, *,
                 eip_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a Endpoint resource.
        :param pulumi.Input[str] eip_id: Id of eip.
        :param pulumi.Input[str] instance_id: Id of instance.
        """
        pulumi.set(__self__, "eip_id", eip_id)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> pulumi.Input[str]:
        """
        Id of eip.
        """
        return pulumi.get(self, "eip_id")

    @eip_id.setter
    def eip_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "eip_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Id of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _EndpointState:
    def __init__(__self__, *,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Endpoint resources.
        :param pulumi.Input[str] eip_id: Id of eip.
        :param pulumi.Input[str] instance_id: Id of instance.
        """
        if eip_id is not None:
            pulumi.set(__self__, "eip_id", eip_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of eip.
        """
        return pulumi.get(self, "eip_id")

    @eip_id.setter
    def eip_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class Endpoint(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage redis endpoint
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.redis.Endpoint("foo",
            eip_id="eip-274ho3mtx543k7fap8tyi****",
            instance_id="redis-cn03bb67g3tr2****")
        ```

        ## Import

        Redis Endpoint can be imported using the instanceId:eipId, e.g.

        ```sh
         $ pulumi import volcengine:redis/endpoint:Endpoint default redis-asdljioeixxxx:eip-2fef2qcfbfw8w5oxruw3w****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_id: Id of eip.
        :param pulumi.Input[str] instance_id: Id of instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EndpointArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage redis endpoint
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.redis.Endpoint("foo",
            eip_id="eip-274ho3mtx543k7fap8tyi****",
            instance_id="redis-cn03bb67g3tr2****")
        ```

        ## Import

        Redis Endpoint can be imported using the instanceId:eipId, e.g.

        ```sh
         $ pulumi import volcengine:redis/endpoint:Endpoint default redis-asdljioeixxxx:eip-2fef2qcfbfw8w5oxruw3w****
        ```

        :param str resource_name: The name of the resource.
        :param EndpointArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EndpointArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 eip_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EndpointArgs.__new__(EndpointArgs)

            if eip_id is None and not opts.urn:
                raise TypeError("Missing required property 'eip_id'")
            __props__.__dict__["eip_id"] = eip_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(Endpoint, __self__).__init__(
            'volcengine:redis/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            eip_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] eip_id: Id of eip.
        :param pulumi.Input[str] instance_id: Id of instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EndpointState.__new__(_EndpointState)

        __props__.__dict__["eip_id"] = eip_id
        __props__.__dict__["instance_id"] = instance_id
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="eipId")
    def eip_id(self) -> pulumi.Output[str]:
        """
        Id of eip.
        """
        return pulumi.get(self, "eip_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Id of instance.
        """
        return pulumi.get(self, "instance_id")

