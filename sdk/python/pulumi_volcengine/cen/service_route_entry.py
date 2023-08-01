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

__all__ = ['ServiceRouteEntryArgs', 'ServiceRouteEntry']

@pulumi.input_type
class ServiceRouteEntryArgs:
    def __init__(__self__, *,
                 cen_id: pulumi.Input[str],
                 destination_cidr_block: pulumi.Input[str],
                 service_region_id: pulumi.Input[str],
                 service_vpc_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 publish_mode: Optional[pulumi.Input[str]] = None,
                 publish_to_instances: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]]] = None):
        """
        The set of arguments for constructing a ServiceRouteEntry resource.
        :param pulumi.Input[str] cen_id: The cen ID of the cen service route entry.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the cen service route entry.
        :param pulumi.Input[str] service_region_id: The service region id of the cen service route entry.
        :param pulumi.Input[str] service_vpc_id: The service VPC id of the cen service route entry.
        :param pulumi.Input[str] description: The description of the cen service route entry.
        :param pulumi.Input[str] publish_mode: Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]] publish_to_instances: The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        """
        pulumi.set(__self__, "cen_id", cen_id)
        pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        pulumi.set(__self__, "service_region_id", service_region_id)
        pulumi.set(__self__, "service_vpc_id", service_vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if publish_mode is not None:
            pulumi.set(__self__, "publish_mode", publish_mode)
        if publish_to_instances is not None:
            pulumi.set(__self__, "publish_to_instances", publish_to_instances)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Input[str]:
        """
        The cen ID of the cen service route entry.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Input[str]:
        """
        The destination cidr block of the cen service route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="serviceRegionId")
    def service_region_id(self) -> pulumi.Input[str]:
        """
        The service region id of the cen service route entry.
        """
        return pulumi.get(self, "service_region_id")

    @service_region_id.setter
    def service_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_region_id", value)

    @property
    @pulumi.getter(name="serviceVpcId")
    def service_vpc_id(self) -> pulumi.Input[str]:
        """
        The service VPC id of the cen service route entry.
        """
        return pulumi.get(self, "service_vpc_id")

    @service_vpc_id.setter
    def service_vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_vpc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cen service route entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="publishMode")
    def publish_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        """
        return pulumi.get(self, "publish_mode")

    @publish_mode.setter
    def publish_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "publish_mode", value)

    @property
    @pulumi.getter(name="publishToInstances")
    def publish_to_instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]]]:
        """
        The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        """
        return pulumi.get(self, "publish_to_instances")

    @publish_to_instances.setter
    def publish_to_instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]]]):
        pulumi.set(self, "publish_to_instances", value)


@pulumi.input_type
class _ServiceRouteEntryState:
    def __init__(__self__, *,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 publish_mode: Optional[pulumi.Input[str]] = None,
                 publish_to_instances: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 service_vpc_id: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ServiceRouteEntry resources.
        :param pulumi.Input[str] cen_id: The cen ID of the cen service route entry.
        :param pulumi.Input[str] creation_time: The create time of the cen service route entry.
        :param pulumi.Input[str] description: The description of the cen service route entry.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the cen service route entry.
        :param pulumi.Input[str] publish_mode: Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        :param pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]] publish_to_instances: The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        :param pulumi.Input[str] service_region_id: The service region id of the cen service route entry.
        :param pulumi.Input[str] service_vpc_id: The service VPC id of the cen service route entry.
        :param pulumi.Input[str] status: The status of the cen service route entry.
        """
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if destination_cidr_block is not None:
            pulumi.set(__self__, "destination_cidr_block", destination_cidr_block)
        if publish_mode is not None:
            pulumi.set(__self__, "publish_mode", publish_mode)
        if publish_to_instances is not None:
            pulumi.set(__self__, "publish_to_instances", publish_to_instances)
        if service_region_id is not None:
            pulumi.set(__self__, "service_region_id", service_region_id)
        if service_vpc_id is not None:
            pulumi.set(__self__, "service_vpc_id", service_vpc_id)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[pulumi.Input[str]]:
        """
        The cen ID of the cen service route entry.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the cen service route entry.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cen service route entry.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> Optional[pulumi.Input[str]]:
        """
        The destination cidr block of the cen service route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @destination_cidr_block.setter
    def destination_cidr_block(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_cidr_block", value)

    @property
    @pulumi.getter(name="publishMode")
    def publish_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        """
        return pulumi.get(self, "publish_mode")

    @publish_mode.setter
    def publish_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "publish_mode", value)

    @property
    @pulumi.getter(name="publishToInstances")
    def publish_to_instances(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]]]:
        """
        The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        """
        return pulumi.get(self, "publish_to_instances")

    @publish_to_instances.setter
    def publish_to_instances(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ServiceRouteEntryPublishToInstanceArgs']]]]):
        pulumi.set(self, "publish_to_instances", value)

    @property
    @pulumi.getter(name="serviceRegionId")
    def service_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The service region id of the cen service route entry.
        """
        return pulumi.get(self, "service_region_id")

    @service_region_id.setter
    def service_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_region_id", value)

    @property
    @pulumi.getter(name="serviceVpcId")
    def service_vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The service VPC id of the cen service route entry.
        """
        return pulumi.get(self, "service_vpc_id")

    @service_vpc_id.setter
    def service_vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_vpc_id", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the cen service route entry.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class ServiceRouteEntry(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 publish_mode: Optional[pulumi.Input[str]] = None,
                 publish_to_instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRouteEntryPublishToInstanceArgs']]]]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 service_vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage cen service route entry
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.ServiceRouteEntry("foo",
            cen_id="cen-12ar8uclj68sg17q7y20v9gil",
            description="test-tf",
            destination_cidr_block="100.64.0.0/11",
            publish_mode="Custom",
            publish_to_instances=[
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-2fepz36a5ra4g59gp67w197xo",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-im67wjcikxkw8gbssx8ufpj8",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
            ],
            service_region_id="cn-beijing",
            service_vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        foo1 = volcengine.cen.ServiceRouteEntry("foo1",
            cen_id="cen-12ar8uclj68sg17q7y20v9gil",
            description="test-tf",
            destination_cidr_block="100.64.0.0/10",
            publish_mode="Custom",
            publish_to_instances=[
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-2fepz36a5ra4g59gp67w197xo",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-im67wjcikxkw8gbssx8ufpj8",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
            ],
            service_region_id="cn-beijing",
            service_vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        foo2 = volcengine.cen.ServiceRouteEntry("foo2",
            cen_id="cen-12ar8uclj68sg17q7y20v9gil",
            description="test-tf",
            destination_cidr_block="100.64.0.0/12",
            publish_mode="Custom",
            publish_to_instances=[
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-2fepz36a5ra4g59gp67w197xo",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-im67wjcikxkw8gbssx8ufpj8",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
            ],
            service_region_id="cn-beijing",
            service_vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        ```

        ## Import

        CenServiceRouteEntry can be imported using the CenId:DestinationCidrBlock:ServiceRegionId:ServiceVpcId, e.g.

        ```sh
         $ pulumi import volcengine:cen/serviceRouteEntry:ServiceRouteEntry default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:cn-beijing:vpc-3rlkeggyn6tc010exd32q****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The cen ID of the cen service route entry.
        :param pulumi.Input[str] description: The description of the cen service route entry.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the cen service route entry.
        :param pulumi.Input[str] publish_mode: Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRouteEntryPublishToInstanceArgs']]]] publish_to_instances: The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        :param pulumi.Input[str] service_region_id: The service region id of the cen service route entry.
        :param pulumi.Input[str] service_vpc_id: The service VPC id of the cen service route entry.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceRouteEntryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage cen service route entry
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.ServiceRouteEntry("foo",
            cen_id="cen-12ar8uclj68sg17q7y20v9gil",
            description="test-tf",
            destination_cidr_block="100.64.0.0/11",
            publish_mode="Custom",
            publish_to_instances=[
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-2fepz36a5ra4g59gp67w197xo",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-im67wjcikxkw8gbssx8ufpj8",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
            ],
            service_region_id="cn-beijing",
            service_vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        foo1 = volcengine.cen.ServiceRouteEntry("foo1",
            cen_id="cen-12ar8uclj68sg17q7y20v9gil",
            description="test-tf",
            destination_cidr_block="100.64.0.0/10",
            publish_mode="Custom",
            publish_to_instances=[
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-2fepz36a5ra4g59gp67w197xo",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-im67wjcikxkw8gbssx8ufpj8",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
            ],
            service_region_id="cn-beijing",
            service_vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        foo2 = volcengine.cen.ServiceRouteEntry("foo2",
            cen_id="cen-12ar8uclj68sg17q7y20v9gil",
            description="test-tf",
            destination_cidr_block="100.64.0.0/12",
            publish_mode="Custom",
            publish_to_instances=[
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-2fepz36a5ra4g59gp67w197xo",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
                volcengine.cen.ServiceRouteEntryPublishToInstanceArgs(
                    instance_id="vpc-im67wjcikxkw8gbssx8ufpj8",
                    instance_region_id="cn-beijing",
                    instance_type="VPC",
                ),
            ],
            service_region_id="cn-beijing",
            service_vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        ```

        ## Import

        CenServiceRouteEntry can be imported using the CenId:DestinationCidrBlock:ServiceRegionId:ServiceVpcId, e.g.

        ```sh
         $ pulumi import volcengine:cen/serviceRouteEntry:ServiceRouteEntry default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:cn-beijing:vpc-3rlkeggyn6tc010exd32q****
        ```

        :param str resource_name: The name of the resource.
        :param ServiceRouteEntryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceRouteEntryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 destination_cidr_block: Optional[pulumi.Input[str]] = None,
                 publish_mode: Optional[pulumi.Input[str]] = None,
                 publish_to_instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRouteEntryPublishToInstanceArgs']]]]] = None,
                 service_region_id: Optional[pulumi.Input[str]] = None,
                 service_vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceRouteEntryArgs.__new__(ServiceRouteEntryArgs)

            if cen_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_id'")
            __props__.__dict__["cen_id"] = cen_id
            __props__.__dict__["description"] = description
            if destination_cidr_block is None and not opts.urn:
                raise TypeError("Missing required property 'destination_cidr_block'")
            __props__.__dict__["destination_cidr_block"] = destination_cidr_block
            __props__.__dict__["publish_mode"] = publish_mode
            __props__.__dict__["publish_to_instances"] = publish_to_instances
            if service_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_region_id'")
            __props__.__dict__["service_region_id"] = service_region_id
            if service_vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_vpc_id'")
            __props__.__dict__["service_vpc_id"] = service_vpc_id
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
        super(ServiceRouteEntry, __self__).__init__(
            'volcengine:cen/serviceRouteEntry:ServiceRouteEntry',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            destination_cidr_block: Optional[pulumi.Input[str]] = None,
            publish_mode: Optional[pulumi.Input[str]] = None,
            publish_to_instances: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRouteEntryPublishToInstanceArgs']]]]] = None,
            service_region_id: Optional[pulumi.Input[str]] = None,
            service_vpc_id: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'ServiceRouteEntry':
        """
        Get an existing ServiceRouteEntry resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The cen ID of the cen service route entry.
        :param pulumi.Input[str] creation_time: The create time of the cen service route entry.
        :param pulumi.Input[str] description: The description of the cen service route entry.
        :param pulumi.Input[str] destination_cidr_block: The destination cidr block of the cen service route entry.
        :param pulumi.Input[str] publish_mode: Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ServiceRouteEntryPublishToInstanceArgs']]]] publish_to_instances: The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        :param pulumi.Input[str] service_region_id: The service region id of the cen service route entry.
        :param pulumi.Input[str] service_vpc_id: The service VPC id of the cen service route entry.
        :param pulumi.Input[str] status: The status of the cen service route entry.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ServiceRouteEntryState.__new__(_ServiceRouteEntryState)

        __props__.__dict__["cen_id"] = cen_id
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["destination_cidr_block"] = destination_cidr_block
        __props__.__dict__["publish_mode"] = publish_mode
        __props__.__dict__["publish_to_instances"] = publish_to_instances
        __props__.__dict__["service_region_id"] = service_region_id
        __props__.__dict__["service_vpc_id"] = service_vpc_id
        __props__.__dict__["status"] = status
        return ServiceRouteEntry(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The cen ID of the cen service route entry.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The create time of the cen service route entry.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the cen service route entry.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="destinationCidrBlock")
    def destination_cidr_block(self) -> pulumi.Output[str]:
        """
        The destination cidr block of the cen service route entry.
        """
        return pulumi.get(self, "destination_cidr_block")

    @property
    @pulumi.getter(name="publishMode")
    def publish_mode(self) -> pulumi.Output[Optional[str]]:
        """
        Publishing scope of cloud service access routes. Valid values are `LocalDCGW`(default), `Custom`.
        """
        return pulumi.get(self, "publish_mode")

    @property
    @pulumi.getter(name="publishToInstances")
    def publish_to_instances(self) -> pulumi.Output[Optional[Sequence['outputs.ServiceRouteEntryPublishToInstance']]]:
        """
        The publish instances. A maximum of 100 can be uploaded in one request. This field needs to be filled in when the `publish_mode` is `Custom`.
        """
        return pulumi.get(self, "publish_to_instances")

    @property
    @pulumi.getter(name="serviceRegionId")
    def service_region_id(self) -> pulumi.Output[str]:
        """
        The service region id of the cen service route entry.
        """
        return pulumi.get(self, "service_region_id")

    @property
    @pulumi.getter(name="serviceVpcId")
    def service_vpc_id(self) -> pulumi.Output[str]:
        """
        The service VPC id of the cen service route entry.
        """
        return pulumi.get(self, "service_vpc_id")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the cen service route entry.
        """
        return pulumi.get(self, "status")

