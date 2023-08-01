# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RouteTableAssociateArgs', 'RouteTableAssociate']

@pulumi.input_type
class RouteTableAssociateArgs:
    def __init__(__self__, *,
                 route_table_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a RouteTableAssociate resource.
        :param pulumi.Input[str] route_table_id: The id of the route table.
        :param pulumi.Input[str] subnet_id: The id of the subnet.
        """
        pulumi.set(__self__, "route_table_id", route_table_id)
        pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Input[str]:
        """
        The id of the route table.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The id of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)


@pulumi.input_type
class _RouteTableAssociateState:
    def __init__(__self__, *,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RouteTableAssociate resources.
        :param pulumi.Input[str] route_table_id: The id of the route table.
        :param pulumi.Input[str] subnet_id: The id of the subnet.
        """
        if route_table_id is not None:
            pulumi.set(__self__, "route_table_id", route_table_id)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the route table.
        """
        return pulumi.get(self, "route_table_id")

    @route_table_id.setter
    def route_table_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "route_table_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)


class RouteTableAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage route table associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.RouteTableAssociate("foo",
            route_table_id="vtb-274e19skkuhog7fap8u4i8ird",
            subnet_id="subnet-2744ht7fhjthc7fap8tm10eqg")
        ```

        ## Import

        Route table associate address can be imported using the route_table_id:subnet_id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/routeTableAssociate:RouteTableAssociate default vtb-2fdzao4h726f45******:subnet-2fdzaou4liw3k5oxruv******
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The id of the route table.
        :param pulumi.Input[str] subnet_id: The id of the subnet.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RouteTableAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage route table associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.RouteTableAssociate("foo",
            route_table_id="vtb-274e19skkuhog7fap8u4i8ird",
            subnet_id="subnet-2744ht7fhjthc7fap8tm10eqg")
        ```

        ## Import

        Route table associate address can be imported using the route_table_id:subnet_id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/routeTableAssociate:RouteTableAssociate default vtb-2fdzao4h726f45******:subnet-2fdzaou4liw3k5oxruv******
        ```

        :param str resource_name: The name of the resource.
        :param RouteTableAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RouteTableAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 route_table_id: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = RouteTableAssociateArgs.__new__(RouteTableAssociateArgs)

            if route_table_id is None and not opts.urn:
                raise TypeError("Missing required property 'route_table_id'")
            __props__.__dict__["route_table_id"] = route_table_id
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
        super(RouteTableAssociate, __self__).__init__(
            'volcengine:vpc/routeTableAssociate:RouteTableAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            route_table_id: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None) -> 'RouteTableAssociate':
        """
        Get an existing RouteTableAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] route_table_id: The id of the route table.
        :param pulumi.Input[str] subnet_id: The id of the subnet.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RouteTableAssociateState.__new__(_RouteTableAssociateState)

        __props__.__dict__["route_table_id"] = route_table_id
        __props__.__dict__["subnet_id"] = subnet_id
        return RouteTableAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="routeTableId")
    def route_table_id(self) -> pulumi.Output[str]:
        """
        The id of the route table.
        """
        return pulumi.get(self, "route_table_id")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The id of the subnet.
        """
        return pulumi.get(self, "subnet_id")

