# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ServerGroupServerArgs', 'ServerGroupServer']

@pulumi.input_type
class ServerGroupServerArgs:
    def __init__(__self__, *,
                 instance_id: pulumi.Input[str],
                 port: pulumi.Input[int],
                 server_group_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a ServerGroupServer resource.
        :param pulumi.Input[str] instance_id: The ID of ecs instance or the network card bound to ecs instance.
        :param pulumi.Input[int] port: The port receiving request.
        :param pulumi.Input[str] server_group_id: The ID of the ServerGroup.
        :param pulumi.Input[str] type: The type of instance. Optional choice contains `ecs`, `eni`.
        :param pulumi.Input[str] description: The description of the instance.
        :param pulumi.Input[str] ip: The private ip of the instance.
        :param pulumi.Input[int] weight: The weight of the instance, range in 0~100.
        """
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "port", port)
        pulumi.set(__self__, "server_group_id", server_group_id)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of ecs instance or the network card bound to ecs instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def port(self) -> pulumi.Input[int]:
        """
        The port receiving request.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: pulumi.Input[int]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> pulumi.Input[str]:
        """
        The ID of the ServerGroup.
        """
        return pulumi.get(self, "server_group_id")

    @server_group_id.setter
    def server_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_group_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        The type of instance. Optional choice contains `ecs`, `eni`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The private ip of the instance.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of the instance, range in 0~100.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


@pulumi.input_type
class _ServerGroupServerState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_group_id: Optional[pulumi.Input[str]] = None,
                 server_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering ServerGroupServer resources.
        :param pulumi.Input[str] description: The description of the instance.
        :param pulumi.Input[str] instance_id: The ID of ecs instance or the network card bound to ecs instance.
        :param pulumi.Input[str] ip: The private ip of the instance.
        :param pulumi.Input[int] port: The port receiving request.
        :param pulumi.Input[str] server_group_id: The ID of the ServerGroup.
        :param pulumi.Input[str] server_id: The server id of instance in ServerGroup.
        :param pulumi.Input[str] type: The type of instance. Optional choice contains `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the instance, range in 0~100.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if port is not None:
            pulumi.set(__self__, "port", port)
        if server_group_id is not None:
            pulumi.set(__self__, "server_group_id", server_group_id)
        if server_id is not None:
            pulumi.set(__self__, "server_id", server_id)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if weight is not None:
            pulumi.set(__self__, "weight", weight)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the instance.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of ecs instance or the network card bound to ecs instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        The private ip of the instance.
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter
    def port(self) -> Optional[pulumi.Input[int]]:
        """
        The port receiving request.
        """
        return pulumi.get(self, "port")

    @port.setter
    def port(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "port", value)

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the ServerGroup.
        """
        return pulumi.get(self, "server_group_id")

    @server_group_id.setter
    def server_group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_group_id", value)

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> Optional[pulumi.Input[str]]:
        """
        The server id of instance in ServerGroup.
        """
        return pulumi.get(self, "server_id")

    @server_id.setter
    def server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "server_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of instance. Optional choice contains `ecs`, `eni`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def weight(self) -> Optional[pulumi.Input[int]]:
        """
        The weight of the instance, range in 0~100.
        """
        return pulumi.get(self, "weight")

    @weight.setter
    def weight(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "weight", value)


class ServerGroupServer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_group_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to manage server group server
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.clb.ServerGroupServer("foo",
            description="This is a server",
            instance_id="i-ybp1scasbe72q1vq35wv",
            port=80,
            server_group_id="rsp-274xltv2sjoxs7fap8tlv3q3s",
            type="ecs",
            weight=100)
        ```

        ## Import

        ServerGroupServer can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:clb/serverGroupServer:ServerGroupServer default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the instance.
        :param pulumi.Input[str] instance_id: The ID of ecs instance or the network card bound to ecs instance.
        :param pulumi.Input[str] ip: The private ip of the instance.
        :param pulumi.Input[int] port: The port receiving request.
        :param pulumi.Input[str] server_group_id: The ID of the ServerGroup.
        :param pulumi.Input[str] type: The type of instance. Optional choice contains `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the instance, range in 0~100.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServerGroupServerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage server group server
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.clb.ServerGroupServer("foo",
            description="This is a server",
            instance_id="i-ybp1scasbe72q1vq35wv",
            port=80,
            server_group_id="rsp-274xltv2sjoxs7fap8tlv3q3s",
            type="ecs",
            weight=100)
        ```

        ## Import

        ServerGroupServer can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:clb/serverGroupServer:ServerGroupServer default rsp-274xltv2*****8tlv3q3s:rs-3ciynux6i1x4w****rszh49sj
        ```

        :param str resource_name: The name of the resource.
        :param ServerGroupServerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServerGroupServerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 port: Optional[pulumi.Input[int]] = None,
                 server_group_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 weight: Optional[pulumi.Input[int]] = None,
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
            __props__ = ServerGroupServerArgs.__new__(ServerGroupServerArgs)

            __props__.__dict__["description"] = description
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            __props__.__dict__["ip"] = ip
            if port is None and not opts.urn:
                raise TypeError("Missing required property 'port'")
            __props__.__dict__["port"] = port
            if server_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'server_group_id'")
            __props__.__dict__["server_group_id"] = server_group_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["weight"] = weight
            __props__.__dict__["server_id"] = None
        super(ServerGroupServer, __self__).__init__(
            'volcengine:clb/serverGroupServer:ServerGroupServer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            port: Optional[pulumi.Input[int]] = None,
            server_group_id: Optional[pulumi.Input[str]] = None,
            server_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            weight: Optional[pulumi.Input[int]] = None) -> 'ServerGroupServer':
        """
        Get an existing ServerGroupServer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the instance.
        :param pulumi.Input[str] instance_id: The ID of ecs instance or the network card bound to ecs instance.
        :param pulumi.Input[str] ip: The private ip of the instance.
        :param pulumi.Input[int] port: The port receiving request.
        :param pulumi.Input[str] server_group_id: The ID of the ServerGroup.
        :param pulumi.Input[str] server_id: The server id of instance in ServerGroup.
        :param pulumi.Input[str] type: The type of instance. Optional choice contains `ecs`, `eni`.
        :param pulumi.Input[int] weight: The weight of the instance, range in 0~100.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServerGroupServerState.__new__(_ServerGroupServerState)

        __props__.__dict__["description"] = description
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["ip"] = ip
        __props__.__dict__["port"] = port
        __props__.__dict__["server_group_id"] = server_group_id
        __props__.__dict__["server_id"] = server_id
        __props__.__dict__["type"] = type
        __props__.__dict__["weight"] = weight
        return ServerGroupServer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of ecs instance or the network card bound to ecs instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        The private ip of the instance.
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter
    def port(self) -> pulumi.Output[int]:
        """
        The port receiving request.
        """
        return pulumi.get(self, "port")

    @property
    @pulumi.getter(name="serverGroupId")
    def server_group_id(self) -> pulumi.Output[str]:
        """
        The ID of the ServerGroup.
        """
        return pulumi.get(self, "server_group_id")

    @property
    @pulumi.getter(name="serverId")
    def server_id(self) -> pulumi.Output[str]:
        """
        The server id of instance in ServerGroup.
        """
        return pulumi.get(self, "server_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of instance. Optional choice contains `ecs`, `eni`.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def weight(self) -> pulumi.Output[Optional[int]]:
        """
        The weight of the instance, range in 0~100.
        """
        return pulumi.get(self, "weight")

