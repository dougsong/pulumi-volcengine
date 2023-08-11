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

__all__ = ['NetworkInterfaceArgs', 'NetworkInterface']

@pulumi.input_type
class NetworkInterfaceArgs:
    def __init__(__self__, *,
                 security_group_ids: pulumi.Input[Sequence[pulumi.Input[str]]],
                 subnet_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_name: Optional[pulumi.Input[str]] = None,
                 port_security_enabled: Optional[pulumi.Input[bool]] = None,
                 primary_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]]] = None):
        """
        The set of arguments for constructing a NetworkInterface resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The list of the security group id to which the secondary ENI belongs.
        :param pulumi.Input[str] subnet_id: The id of the subnet to which the ENI is connected.
        :param pulumi.Input[str] description: The description of the ENI.
        :param pulumi.Input[str] network_interface_name: The name of the ENI.
        :param pulumi.Input[bool] port_security_enabled: Set port security enable or disable.
        :param pulumi.Input[str] primary_ip_address: The primary IP address of the ENI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_addresses: The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        :param pulumi.Input[str] project_name: The ProjectName of the ENI.
        :param pulumi.Input[int] secondary_private_ip_address_count: The count of secondary private ip address. This field conflicts with `private_ip_address`.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "security_group_ids", security_group_ids)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if network_interface_name is not None:
            pulumi.set(__self__, "network_interface_name", network_interface_name)
        if port_security_enabled is not None:
            pulumi.set(__self__, "port_security_enabled", port_security_enabled)
        if primary_ip_address is not None:
            pulumi.set(__self__, "primary_ip_address", primary_ip_address)
        if private_ip_addresses is not None:
            pulumi.set(__self__, "private_ip_addresses", private_ip_addresses)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if secondary_private_ip_address_count is not None:
            pulumi.set(__self__, "secondary_private_ip_address_count", secondary_private_ip_address_count)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The list of the security group id to which the secondary ENI belongs.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The id of the subnet to which the ENI is connected.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the ENI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="networkInterfaceName")
    def network_interface_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ENI.
        """
        return pulumi.get(self, "network_interface_name")

    @network_interface_name.setter
    def network_interface_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_name", value)

    @property
    @pulumi.getter(name="portSecurityEnabled")
    def port_security_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set port security enable or disable.
        """
        return pulumi.get(self, "port_security_enabled")

    @port_security_enabled.setter
    def port_security_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "port_security_enabled", value)

    @property
    @pulumi.getter(name="primaryIpAddress")
    def primary_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The primary IP address of the ENI.
        """
        return pulumi.get(self, "primary_ip_address")

    @primary_ip_address.setter
    def primary_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_ip_address", value)

    @property
    @pulumi.getter(name="privateIpAddresses")
    def private_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        """
        return pulumi.get(self, "private_ip_addresses")

    @private_ip_addresses.setter
    def private_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "private_ip_addresses", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the ENI.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="secondaryPrivateIpAddressCount")
    def secondary_private_ip_address_count(self) -> Optional[pulumi.Input[int]]:
        """
        The count of secondary private ip address. This field conflicts with `private_ip_address`.
        """
        return pulumi.get(self, "secondary_private_ip_address_count")

    @secondary_private_ip_address_count.setter
    def secondary_private_ip_address_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secondary_private_ip_address_count", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _NetworkInterfaceState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_name: Optional[pulumi.Input[str]] = None,
                 port_security_enabled: Optional[pulumi.Input[bool]] = None,
                 primary_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[int]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering NetworkInterface resources.
        :param pulumi.Input[str] description: The description of the ENI.
        :param pulumi.Input[str] network_interface_name: The name of the ENI.
        :param pulumi.Input[bool] port_security_enabled: Set port security enable or disable.
        :param pulumi.Input[str] primary_ip_address: The primary IP address of the ENI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_addresses: The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        :param pulumi.Input[str] project_name: The ProjectName of the ENI.
        :param pulumi.Input[int] secondary_private_ip_address_count: The count of secondary private ip address. This field conflicts with `private_ip_address`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The list of the security group id to which the secondary ENI belongs.
        :param pulumi.Input[str] status: The status of the ENI.
        :param pulumi.Input[str] subnet_id: The id of the subnet to which the ENI is connected.
        :param pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]] tags: Tags.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if network_interface_name is not None:
            pulumi.set(__self__, "network_interface_name", network_interface_name)
        if port_security_enabled is not None:
            pulumi.set(__self__, "port_security_enabled", port_security_enabled)
        if primary_ip_address is not None:
            pulumi.set(__self__, "primary_ip_address", primary_ip_address)
        if private_ip_addresses is not None:
            pulumi.set(__self__, "private_ip_addresses", private_ip_addresses)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if secondary_private_ip_address_count is not None:
            pulumi.set(__self__, "secondary_private_ip_address_count", secondary_private_ip_address_count)
        if security_group_ids is not None:
            pulumi.set(__self__, "security_group_ids", security_group_ids)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the ENI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="networkInterfaceName")
    def network_interface_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the ENI.
        """
        return pulumi.get(self, "network_interface_name")

    @network_interface_name.setter
    def network_interface_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_name", value)

    @property
    @pulumi.getter(name="portSecurityEnabled")
    def port_security_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Set port security enable or disable.
        """
        return pulumi.get(self, "port_security_enabled")

    @port_security_enabled.setter
    def port_security_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "port_security_enabled", value)

    @property
    @pulumi.getter(name="primaryIpAddress")
    def primary_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The primary IP address of the ENI.
        """
        return pulumi.get(self, "primary_ip_address")

    @primary_ip_address.setter
    def primary_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "primary_ip_address", value)

    @property
    @pulumi.getter(name="privateIpAddresses")
    def private_ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        """
        return pulumi.get(self, "private_ip_addresses")

    @private_ip_addresses.setter
    def private_ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "private_ip_addresses", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the ENI.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="secondaryPrivateIpAddressCount")
    def secondary_private_ip_address_count(self) -> Optional[pulumi.Input[int]]:
        """
        The count of secondary private ip address. This field conflicts with `private_ip_address`.
        """
        return pulumi.get(self, "secondary_private_ip_address_count")

    @secondary_private_ip_address_count.setter
    def secondary_private_ip_address_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secondary_private_ip_address_count", value)

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The list of the security group id to which the secondary ENI belongs.
        """
        return pulumi.get(self, "security_group_ids")

    @security_group_ids.setter
    def security_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "security_group_ids", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the ENI.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the subnet to which the ENI is connected.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkInterfaceTagArgs']]]]):
        pulumi.set(self, "tags", value)


class NetworkInterface(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_name: Optional[pulumi.Input[str]] = None,
                 port_security_enabled: Optional[pulumi.Input[bool]] = None,
                 primary_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[int]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkInterfaceTagArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage network interface
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.NetworkInterface("foo",
            description="tf-test-up",
            network_interface_name="tf-test-up",
            port_security_enabled=False,
            primary_ip_address="192.168.5.253",
            private_ip_addresses=["192.168.5.2"],
            project_name="default",
            security_group_ids=["sg-2fepz3c793g1s59gp67y21r34"],
            subnet_id="subnet-2fe79j7c8o5c059gp68ksxr93")
        ```

        ## Import

        Network interface can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/networkInterface:NetworkInterface default eni-bp1fgnh68xyz9****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the ENI.
        :param pulumi.Input[str] network_interface_name: The name of the ENI.
        :param pulumi.Input[bool] port_security_enabled: Set port security enable or disable.
        :param pulumi.Input[str] primary_ip_address: The primary IP address of the ENI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_addresses: The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        :param pulumi.Input[str] project_name: The ProjectName of the ENI.
        :param pulumi.Input[int] secondary_private_ip_address_count: The count of secondary private ip address. This field conflicts with `private_ip_address`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The list of the security group id to which the secondary ENI belongs.
        :param pulumi.Input[str] subnet_id: The id of the subnet to which the ENI is connected.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkInterfaceTagArgs']]]] tags: Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkInterfaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage network interface
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.NetworkInterface("foo",
            description="tf-test-up",
            network_interface_name="tf-test-up",
            port_security_enabled=False,
            primary_ip_address="192.168.5.253",
            private_ip_addresses=["192.168.5.2"],
            project_name="default",
            security_group_ids=["sg-2fepz3c793g1s59gp67y21r34"],
            subnet_id="subnet-2fe79j7c8o5c059gp68ksxr93")
        ```

        ## Import

        Network interface can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/networkInterface:NetworkInterface default eni-bp1fgnh68xyz9****
        ```

        :param str resource_name: The name of the resource.
        :param NetworkInterfaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkInterfaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 network_interface_name: Optional[pulumi.Input[str]] = None,
                 port_security_enabled: Optional[pulumi.Input[bool]] = None,
                 primary_ip_address: Optional[pulumi.Input[str]] = None,
                 private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 secondary_private_ip_address_count: Optional[pulumi.Input[int]] = None,
                 security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkInterfaceTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkInterfaceArgs.__new__(NetworkInterfaceArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["network_interface_name"] = network_interface_name
            __props__.__dict__["port_security_enabled"] = port_security_enabled
            __props__.__dict__["primary_ip_address"] = primary_ip_address
            __props__.__dict__["private_ip_addresses"] = private_ip_addresses
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["secondary_private_ip_address_count"] = secondary_private_ip_address_count
            if security_group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'security_group_ids'")
            __props__.__dict__["security_group_ids"] = security_group_ids
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            __props__.__dict__["status"] = None
        super(NetworkInterface, __self__).__init__(
            'volcengine:vpc/networkInterface:NetworkInterface',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            network_interface_name: Optional[pulumi.Input[str]] = None,
            port_security_enabled: Optional[pulumi.Input[bool]] = None,
            primary_ip_address: Optional[pulumi.Input[str]] = None,
            private_ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            secondary_private_ip_address_count: Optional[pulumi.Input[int]] = None,
            security_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            status: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkInterfaceTagArgs']]]]] = None) -> 'NetworkInterface':
        """
        Get an existing NetworkInterface resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the ENI.
        :param pulumi.Input[str] network_interface_name: The name of the ENI.
        :param pulumi.Input[bool] port_security_enabled: Set port security enable or disable.
        :param pulumi.Input[str] primary_ip_address: The primary IP address of the ENI.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] private_ip_addresses: The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        :param pulumi.Input[str] project_name: The ProjectName of the ENI.
        :param pulumi.Input[int] secondary_private_ip_address_count: The count of secondary private ip address. This field conflicts with `private_ip_address`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] security_group_ids: The list of the security group id to which the secondary ENI belongs.
        :param pulumi.Input[str] status: The status of the ENI.
        :param pulumi.Input[str] subnet_id: The id of the subnet to which the ENI is connected.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkInterfaceTagArgs']]]] tags: Tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkInterfaceState.__new__(_NetworkInterfaceState)

        __props__.__dict__["description"] = description
        __props__.__dict__["network_interface_name"] = network_interface_name
        __props__.__dict__["port_security_enabled"] = port_security_enabled
        __props__.__dict__["primary_ip_address"] = primary_ip_address
        __props__.__dict__["private_ip_addresses"] = private_ip_addresses
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["secondary_private_ip_address_count"] = secondary_private_ip_address_count
        __props__.__dict__["security_group_ids"] = security_group_ids
        __props__.__dict__["status"] = status
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        return NetworkInterface(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the ENI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="networkInterfaceName")
    def network_interface_name(self) -> pulumi.Output[str]:
        """
        The name of the ENI.
        """
        return pulumi.get(self, "network_interface_name")

    @property
    @pulumi.getter(name="portSecurityEnabled")
    def port_security_enabled(self) -> pulumi.Output[bool]:
        """
        Set port security enable or disable.
        """
        return pulumi.get(self, "port_security_enabled")

    @property
    @pulumi.getter(name="primaryIpAddress")
    def primary_ip_address(self) -> pulumi.Output[str]:
        """
        The primary IP address of the ENI.
        """
        return pulumi.get(self, "primary_ip_address")

    @property
    @pulumi.getter(name="privateIpAddresses")
    def private_ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of private ip address. This field conflicts with `secondary_private_ip_address_count`.
        """
        return pulumi.get(self, "private_ip_addresses")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[Optional[str]]:
        """
        The ProjectName of the ENI.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="secondaryPrivateIpAddressCount")
    def secondary_private_ip_address_count(self) -> pulumi.Output[int]:
        """
        The count of secondary private ip address. This field conflicts with `private_ip_address`.
        """
        return pulumi.get(self, "secondary_private_ip_address_count")

    @property
    @pulumi.getter(name="securityGroupIds")
    def security_group_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of the security group id to which the secondary ENI belongs.
        """
        return pulumi.get(self, "security_group_ids")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the ENI.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The id of the subnet to which the ENI is connected.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkInterfaceTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

