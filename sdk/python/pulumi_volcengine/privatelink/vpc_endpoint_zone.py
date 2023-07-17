# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcEndpointZoneArgs', 'VpcEndpointZone']

@pulumi.input_type
class VpcEndpointZoneArgs:
    def __init__(__self__, *,
                 endpoint_id: pulumi.Input[str],
                 subnet_id: pulumi.Input[str],
                 private_ip_address: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VpcEndpointZone resource.
        :param pulumi.Input[str] endpoint_id: The endpoint id of vpc endpoint zone.
        :param pulumi.Input[str] subnet_id: The subnet id of vpc endpoint zone.
        :param pulumi.Input[str] private_ip_address: The private ip address of vpc endpoint zone.
        """
        pulumi.set(__self__, "endpoint_id", endpoint_id)
        pulumi.set(__self__, "subnet_id", subnet_id)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Input[str]:
        """
        The endpoint id of vpc endpoint zone.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The subnet id of vpc endpoint zone.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The private ip address of vpc endpoint zone.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)


@pulumi.input_type
class _VpcEndpointZoneState:
    def __init__(__self__, *,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 network_interface_id: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 zone_domain: Optional[pulumi.Input[str]] = None,
                 zone_id: Optional[pulumi.Input[str]] = None,
                 zone_status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointZone resources.
        :param pulumi.Input[str] endpoint_id: The endpoint id of vpc endpoint zone.
        :param pulumi.Input[str] network_interface_id: The network interface id of vpc endpoint.
        :param pulumi.Input[str] private_ip_address: The private ip address of vpc endpoint zone.
        :param pulumi.Input[str] subnet_id: The subnet id of vpc endpoint zone.
        :param pulumi.Input[str] zone_domain: The domain of vpc endpoint zone.
        :param pulumi.Input[str] zone_id: The Id of vpc endpoint zone.
        :param pulumi.Input[str] zone_status: The status of vpc endpoint zone.
        """
        if endpoint_id is not None:
            pulumi.set(__self__, "endpoint_id", endpoint_id)
        if network_interface_id is not None:
            pulumi.set(__self__, "network_interface_id", network_interface_id)
        if private_ip_address is not None:
            pulumi.set(__self__, "private_ip_address", private_ip_address)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if zone_domain is not None:
            pulumi.set(__self__, "zone_domain", zone_domain)
        if zone_id is not None:
            pulumi.set(__self__, "zone_id", zone_id)
        if zone_status is not None:
            pulumi.set(__self__, "zone_status", zone_status)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        The endpoint id of vpc endpoint zone.
        """
        return pulumi.get(self, "endpoint_id")

    @endpoint_id.setter
    def endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "endpoint_id", value)

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> Optional[pulumi.Input[str]]:
        """
        The network interface id of vpc endpoint.
        """
        return pulumi.get(self, "network_interface_id")

    @network_interface_id.setter
    def network_interface_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_interface_id", value)

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The private ip address of vpc endpoint zone.
        """
        return pulumi.get(self, "private_ip_address")

    @private_ip_address.setter
    def private_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_ip_address", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The subnet id of vpc endpoint zone.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="zoneDomain")
    def zone_domain(self) -> Optional[pulumi.Input[str]]:
        """
        The domain of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_domain")

    @zone_domain.setter
    def zone_domain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_domain", value)

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> Optional[pulumi.Input[str]]:
        """
        The Id of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_id")

    @zone_id.setter
    def zone_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_id", value)

    @property
    @pulumi.getter(name="zoneStatus")
    def zone_status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_status")

    @zone_status.setter
    def zone_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "zone_status", value)


class VpcEndpointZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage privatelink vpc endpoint zone
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.privatelink.VpcEndpointZone("foo",
            endpoint_id="ep-2byz5nlkimc5c2dx0ef9g****",
            private_ip_address="172.16.0.251",
            subnet_id="subnet-2bz47q19zhx4w2dx0eevn****")
        ```

        ## Import

        VpcEndpointZone can be imported using the endpointId:subnetId, e.g.

        ```sh
         $ pulumi import volcengine:privatelink/vpcEndpointZone:VpcEndpointZone default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_id: The endpoint id of vpc endpoint zone.
        :param pulumi.Input[str] private_ip_address: The private ip address of vpc endpoint zone.
        :param pulumi.Input[str] subnet_id: The subnet id of vpc endpoint zone.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage privatelink vpc endpoint zone
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.privatelink.VpcEndpointZone("foo",
            endpoint_id="ep-2byz5nlkimc5c2dx0ef9g****",
            private_ip_address="172.16.0.251",
            subnet_id="subnet-2bz47q19zhx4w2dx0eevn****")
        ```

        ## Import

        VpcEndpointZone can be imported using the endpointId:subnetId, e.g.

        ```sh
         $ pulumi import volcengine:privatelink/vpcEndpointZone:VpcEndpointZone default ep-3rel75r081l345zsk2i59****:subnet-2bz47q19zhx4w2dx0eevn****
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 endpoint_id: Optional[pulumi.Input[str]] = None,
                 private_ip_address: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcEndpointZoneArgs.__new__(VpcEndpointZoneArgs)

            if endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint_id'")
            __props__.__dict__["endpoint_id"] = endpoint_id
            __props__.__dict__["private_ip_address"] = private_ip_address
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["network_interface_id"] = None
            __props__.__dict__["zone_domain"] = None
            __props__.__dict__["zone_id"] = None
            __props__.__dict__["zone_status"] = None
        super(VpcEndpointZone, __self__).__init__(
            'volcengine:privatelink/vpcEndpointZone:VpcEndpointZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            endpoint_id: Optional[pulumi.Input[str]] = None,
            network_interface_id: Optional[pulumi.Input[str]] = None,
            private_ip_address: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            zone_domain: Optional[pulumi.Input[str]] = None,
            zone_id: Optional[pulumi.Input[str]] = None,
            zone_status: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointZone':
        """
        Get an existing VpcEndpointZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] endpoint_id: The endpoint id of vpc endpoint zone.
        :param pulumi.Input[str] network_interface_id: The network interface id of vpc endpoint.
        :param pulumi.Input[str] private_ip_address: The private ip address of vpc endpoint zone.
        :param pulumi.Input[str] subnet_id: The subnet id of vpc endpoint zone.
        :param pulumi.Input[str] zone_domain: The domain of vpc endpoint zone.
        :param pulumi.Input[str] zone_id: The Id of vpc endpoint zone.
        :param pulumi.Input[str] zone_status: The status of vpc endpoint zone.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointZoneState.__new__(_VpcEndpointZoneState)

        __props__.__dict__["endpoint_id"] = endpoint_id
        __props__.__dict__["network_interface_id"] = network_interface_id
        __props__.__dict__["private_ip_address"] = private_ip_address
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["zone_domain"] = zone_domain
        __props__.__dict__["zone_id"] = zone_id
        __props__.__dict__["zone_status"] = zone_status
        return VpcEndpointZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="endpointId")
    def endpoint_id(self) -> pulumi.Output[str]:
        """
        The endpoint id of vpc endpoint zone.
        """
        return pulumi.get(self, "endpoint_id")

    @property
    @pulumi.getter(name="networkInterfaceId")
    def network_interface_id(self) -> pulumi.Output[str]:
        """
        The network interface id of vpc endpoint.
        """
        return pulumi.get(self, "network_interface_id")

    @property
    @pulumi.getter(name="privateIpAddress")
    def private_ip_address(self) -> pulumi.Output[str]:
        """
        The private ip address of vpc endpoint zone.
        """
        return pulumi.get(self, "private_ip_address")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The subnet id of vpc endpoint zone.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter(name="zoneDomain")
    def zone_domain(self) -> pulumi.Output[str]:
        """
        The domain of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_domain")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> pulumi.Output[str]:
        """
        The Id of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_id")

    @property
    @pulumi.getter(name="zoneStatus")
    def zone_status(self) -> pulumi.Output[str]:
        """
        The status of vpc endpoint zone.
        """
        return pulumi.get(self, "zone_status")

