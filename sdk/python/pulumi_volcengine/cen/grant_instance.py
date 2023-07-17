# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['GrantInstanceArgs', 'GrantInstance']

@pulumi.input_type
class GrantInstanceArgs:
    def __init__(__self__, *,
                 cen_id: pulumi.Input[str],
                 cen_owner_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 instance_region_id: pulumi.Input[str],
                 instance_type: pulumi.Input[str]):
        """
        The set of arguments for constructing a GrantInstance resource.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        :param pulumi.Input[str] cen_owner_id: The owner ID of the cen.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] instance_region_id: The region ID of the instance.
        :param pulumi.Input[str] instance_type: The type of the instance.
        """
        pulumi.set(__self__, "cen_id", cen_id)
        pulumi.set(__self__, "cen_owner_id", cen_owner_id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "instance_region_id", instance_region_id)
        pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Input[str]:
        """
        The ID of the cen.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="cenOwnerId")
    def cen_owner_id(self) -> pulumi.Input[str]:
        """
        The owner ID of the cen.
        """
        return pulumi.get(self, "cen_owner_id")

    @cen_owner_id.setter
    def cen_owner_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_owner_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> pulumi.Input[str]:
        """
        The region ID of the instance.
        """
        return pulumi.get(self, "instance_region_id")

    @instance_region_id.setter
    def instance_region_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_region_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Input[str]:
        """
        The type of the instance.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_type", value)


@pulumi.input_type
class _GrantInstanceState:
    def __init__(__self__, *,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_owner_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GrantInstance resources.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        :param pulumi.Input[str] cen_owner_id: The owner ID of the cen.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] instance_region_id: The region ID of the instance.
        :param pulumi.Input[str] instance_type: The type of the instance.
        """
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)
        if cen_owner_id is not None:
            pulumi.set(__self__, "cen_owner_id", cen_owner_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)
        if instance_region_id is not None:
            pulumi.set(__self__, "instance_region_id", instance_region_id)
        if instance_type is not None:
            pulumi.set(__self__, "instance_type", instance_type)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the cen.
        """
        return pulumi.get(self, "cen_id")

    @cen_id.setter
    def cen_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_id", value)

    @property
    @pulumi.getter(name="cenOwnerId")
    def cen_owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        The owner ID of the cen.
        """
        return pulumi.get(self, "cen_owner_id")

    @cen_owner_id.setter
    def cen_owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_owner_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> Optional[pulumi.Input[str]]:
        """
        The region ID of the instance.
        """
        return pulumi.get(self, "instance_region_id")

    @instance_region_id.setter
    def instance_region_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_region_id", value)

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the instance.
        """
        return pulumi.get(self, "instance_type")

    @instance_type.setter
    def instance_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_type", value)


class GrantInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_owner_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage cen grant instance
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.GrantInstance("foo",
            cen_id="cen-2d6zdn0c1z5s058ozfcyf4lee",
            cen_owner_id="210000****",
            instance_id="vpc-2bysvq1xx543k2dx0eeulpeiv",
            instance_region_id="cn-guilin-boe",
            instance_type="VPC")
        ```

        ## Import

        Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.

        ```sh
         $ pulumi import volcengine:cen/grantInstance:GrantInstance default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        :param pulumi.Input[str] cen_owner_id: The owner ID of the cen.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] instance_region_id: The region ID of the instance.
        :param pulumi.Input[str] instance_type: The type of the instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GrantInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage cen grant instance
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.GrantInstance("foo",
            cen_id="cen-2d6zdn0c1z5s058ozfcyf4lee",
            cen_owner_id="210000****",
            instance_id="vpc-2bysvq1xx543k2dx0eeulpeiv",
            instance_region_id="cn-guilin-boe",
            instance_type="VPC")
        ```

        ## Import

        Cen grant instance can be imported using the CenId:CenOwnerId:InstanceId:InstanceType:RegionId, e.g.

        ```sh
         $ pulumi import volcengine:cen/grantInstance:GrantInstance default cen-7qthudw0ll6jmc***:210000****:vpc-2fexiqjlgjif45oxruvso****:VPC:cn-beijing
        ```

        :param str resource_name: The name of the resource.
        :param GrantInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GrantInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_owner_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 instance_region_id: Optional[pulumi.Input[str]] = None,
                 instance_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = GrantInstanceArgs.__new__(GrantInstanceArgs)

            if cen_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_id'")
            __props__.__dict__["cen_id"] = cen_id
            if cen_owner_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_owner_id'")
            __props__.__dict__["cen_owner_id"] = cen_owner_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
            if instance_region_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_region_id'")
            __props__.__dict__["instance_region_id"] = instance_region_id
            if instance_type is None and not opts.urn:
                raise TypeError("Missing required property 'instance_type'")
            __props__.__dict__["instance_type"] = instance_type
        super(GrantInstance, __self__).__init__(
            'volcengine:cen/grantInstance:GrantInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            cen_owner_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None,
            instance_region_id: Optional[pulumi.Input[str]] = None,
            instance_type: Optional[pulumi.Input[str]] = None) -> 'GrantInstance':
        """
        Get an existing GrantInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        :param pulumi.Input[str] cen_owner_id: The owner ID of the cen.
        :param pulumi.Input[str] instance_id: The ID of the instance.
        :param pulumi.Input[str] instance_region_id: The region ID of the instance.
        :param pulumi.Input[str] instance_type: The type of the instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GrantInstanceState.__new__(_GrantInstanceState)

        __props__.__dict__["cen_id"] = cen_id
        __props__.__dict__["cen_owner_id"] = cen_owner_id
        __props__.__dict__["instance_id"] = instance_id
        __props__.__dict__["instance_region_id"] = instance_region_id
        __props__.__dict__["instance_type"] = instance_type
        return GrantInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The ID of the cen.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="cenOwnerId")
    def cen_owner_id(self) -> pulumi.Output[str]:
        """
        The owner ID of the cen.
        """
        return pulumi.get(self, "cen_owner_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter(name="instanceRegionId")
    def instance_region_id(self) -> pulumi.Output[str]:
        """
        The region ID of the instance.
        """
        return pulumi.get(self, "instance_region_id")

    @property
    @pulumi.getter(name="instanceType")
    def instance_type(self) -> pulumi.Output[str]:
        """
        The type of the instance.
        """
        return pulumi.get(self, "instance_type")

