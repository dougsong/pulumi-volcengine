# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GatewayRouteArgs', 'GatewayRoute']

@pulumi.input_type
class GatewayRouteArgs:
    def __init__(__self__, *,
                 destination_cidr_block: pulumi.Input[str],
                 next_hop_id: pulumi.Input[str],
                 vpn_gateway_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a GatewayRoute resource.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the VPN gateway route.
        :param pulumi.Input[str] next_hop_id: The next hop id of the VPN gateway route.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway of the VPN gateway route.
        """
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "next_hop_id", next_hop_id)
        pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        The destination cidr block of the VPN gateway route.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="nextHopId")
    def next_hop_id(self) -> pulumi.Input[str]:
        """
        The next hop id of the VPN gateway route.
        """
        return pulumi.get(self, "next_hop_id")

    @next_hop_id.setter
    def next_hop_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "next_hop_id", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPN gateway of the VPN gateway route.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpn_gateway_id", value)


@pulumi.input_type
class _GatewayRouteState:
    def __init__(__self__, *,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 next_hop_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_route_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GatewayRoute resources.
        :param pulumi.Input[str] creation_time: The create time of VPN gateway route.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the VPN gateway route.
        :param pulumi.Input[str] next_hop_id: The next hop id of the VPN gateway route.
        :param pulumi.Input[str] status: The status of the VPN gateway route.
        :param pulumi.Input[str] update_time: The update time of VPN gateway route.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway of the VPN gateway route.
        :param pulumi.Input[str] vpn_gateway_route_id: The ID of the VPN gateway route.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if next_hop_id is not None:
            pulumi.set(__self__, "next_hop_id", next_hop_id)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)
        if vpn_gateway_id is not None:
            pulumi.set(__self__, "vpn_gateway_id", vpn_gateway_id)
        if vpn_gateway_route_id is not None:
            pulumi.set(__self__, "vpn_gateway_route_id", vpn_gateway_route_id)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of VPN gateway route.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The destination cidr block of the VPN gateway route.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="nextHopId")
    def next_hop_id(self) -> Optional[pulumi.Input[str]]:
        """
        The next hop id of the VPN gateway route.
        """
        return pulumi.get(self, "next_hop_id")

    @next_hop_id.setter
    def next_hop_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "next_hop_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the VPN gateway route.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of VPN gateway route.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPN gateway of the VPN gateway route.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @vpn_gateway_id.setter
    def vpn_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_id", value)

    @property
    @pulumi.getter(name="vpnGatewayRouteId")
    def vpn_gateway_route_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPN gateway route.
        """
        return pulumi.get(self, "vpn_gateway_route_id")

    @vpn_gateway_route_id.setter
    def vpn_gateway_route_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpn_gateway_route_id", value)


class GatewayRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 next_hop_id: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage vpn gateway route
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpn.GatewayRoute("foo",
            destination_cidr_block="192.168.0.0/20",
            next_hop_id="vgc-2d5ww3ww2lwcg58ozfe61ppc3",
            vpn_gateway_id="vgw-2c012ea9fm5mo2dx0efxg46qi")
        ```

        ## Import

        VpnGatewayRoute can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpn/gatewayRoute:GatewayRoute default vgr-3tex2c6c0v844c****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the VPN gateway route.
        :param pulumi.Input[str] next_hop_id: The next hop id of the VPN gateway route.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway of the VPN gateway route.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage vpn gateway route
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpn.GatewayRoute("foo",
            destination_cidr_block="192.168.0.0/20",
            next_hop_id="vgc-2d5ww3ww2lwcg58ozfe61ppc3",
            vpn_gateway_id="vgw-2c012ea9fm5mo2dx0efxg46qi")
        ```

        ## Import

        VpnGatewayRoute can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpn/gatewayRoute:GatewayRoute default vgr-3tex2c6c0v844c****
        ```

        :param str resource_name: The name of the resource.
        :param GatewayRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 next_hop_id: Optional[pulumi.Input[str]] = None,
                 vpn_gateway_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GatewayRouteArgs.__new__(GatewayRouteArgs)

            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            if next_hop_id is None and not opts.urn:
                raise TypeError("Missing required property 'next_hop_id'")
            __props__.__dict__["next_hop_id"] = next_hop_id
            if vpn_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpn_gateway_id'")
            __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_time"] = None
            __props__.__dict__["vpn_gateway_route_id"] = None
        super(GatewayRoute, __self__).__init__(
            'volcengine:vpn/gatewayRoute:GatewayRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            next_hop_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None,
            vpn_gateway_id: Optional[pulumi.Input[str]] = None,
            vpn_gateway_route_id: Optional[pulumi.Input[str]] = None) -> 'GatewayRoute':
        """
        Get an existing GatewayRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: The create time of VPN gateway route.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the VPN gateway route.
        :param pulumi.Input[str] next_hop_id: The next hop id of the VPN gateway route.
        :param pulumi.Input[str] status: The status of the VPN gateway route.
        :param pulumi.Input[str] update_time: The update time of VPN gateway route.
        :param pulumi.Input[str] vpn_gateway_id: The ID of the VPN gateway of the VPN gateway route.
        :param pulumi.Input[str] vpn_gateway_route_id: The ID of the VPN gateway route.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayRouteState.__new__(_GatewayRouteState)

        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["next_hop_id"] = next_hop_id
        __props__.__dict__["status"] = status
        __props__.__dict__["update_time"] = update_time
        __props__.__dict__["vpn_gateway_id"] = vpn_gateway_id
        __props__.__dict__["vpn_gateway_route_id"] = vpn_gateway_route_id
        return GatewayRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The create time of VPN gateway route.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        The destination cidr block of the VPN gateway route.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="nextHopId")
    def next_hop_id(self) -> pulumi.Output[str]:
        """
        The next hop id of the VPN gateway route.
        """
        return pulumi.get(self, "next_hop_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the VPN gateway route.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of VPN gateway route.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="vpnGatewayId")
    def vpn_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPN gateway of the VPN gateway route.
        """
        return pulumi.get(self, "vpn_gateway_id")

    @property
    @pulumi.getter(name="vpnGatewayRouteId")
    def vpn_gateway_route_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPN gateway route.
        """
        return pulumi.get(self, "vpn_gateway_route_id")

