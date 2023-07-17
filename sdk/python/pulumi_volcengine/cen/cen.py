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

__all__ = ['CenArgs', 'Cen']

@pulumi.input_type
class CenArgs:
    def __init__(__self__, *,
                 cen_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]]] = None):
        """
        The set of arguments for constructing a Cen resource.
        :param pulumi.Input[str] cen_name: The name of the cen.
        :param pulumi.Input[str] description: The description of the cen.
        :param pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]] tags: Tags.
        """
        if cen_name is not None:
            pulumi.set(__self__, "cen_name", cen_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="cenName")
    def cen_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cen.
        """
        return pulumi.get(self, "cen_name")

    @cen_name.setter
    def cen_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cen.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _CenState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 cen_bandwidth_package_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 cen_name: Optional[pulumi.Input[str]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Cen resources.
        :param pulumi.Input[str] account_id: The account ID of the cen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cen_bandwidth_package_ids: A list of bandwidth package IDs of the cen.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        :param pulumi.Input[str] cen_name: The name of the cen.
        :param pulumi.Input[str] creation_time: The create time of the cen.
        :param pulumi.Input[str] description: The description of the cen.
        :param pulumi.Input[str] status: The status of the cen.
        :param pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]] tags: Tags.
        :param pulumi.Input[str] update_time: The update time of the cen.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if cen_bandwidth_package_ids is not None:
            pulumi.set(__self__, "cen_bandwidth_package_ids", cen_bandwidth_package_ids)
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)
        if cen_name is not None:
            pulumi.set(__self__, "cen_name", cen_name)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The account ID of the cen.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="cenBandwidthPackageIds")
    def cen_bandwidth_package_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of bandwidth package IDs of the cen.
        """
        return pulumi.get(self, "cen_bandwidth_package_ids")

    @cen_bandwidth_package_ids.setter
    def cen_bandwidth_package_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cen_bandwidth_package_ids", value)

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
    @pulumi.getter(name="cenName")
    def cen_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cen.
        """
        return pulumi.get(self, "cen_name")

    @cen_name.setter
    def cen_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_name", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the cen.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cen.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the cen.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CenTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of the cen.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Cen(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CenTagArgs']]]]] = None,
                 __props__=None):
        """
        Provides a resource to manage cen
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.Cen("foo",
            cen_name="tf-test",
            description="tf-test")
        ```

        ## Import

        Cen can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:cen/cen:Cen default cen-7qthudw0ll6jmc****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_name: The name of the cen.
        :param pulumi.Input[str] description: The description of the cen.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CenTagArgs']]]] tags: Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[CenArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage cen
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.Cen("foo",
            cen_name="tf-test",
            description="tf-test")
        ```

        ## Import

        Cen can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:cen/cen:Cen default cen-7qthudw0ll6jmc****
        ```

        :param str resource_name: The name of the resource.
        :param CenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CenTagArgs']]]]] = None,
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
            __props__ = CenArgs.__new__(CenArgs)

            __props__.__dict__["cen_name"] = cen_name
            __props__.__dict__["description"] = description
            __props__.__dict__["tags"] = tags
            __props__.__dict__["account_id"] = None
            __props__.__dict__["cen_bandwidth_package_ids"] = None
            __props__.__dict__["cen_id"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_time"] = None
        super(Cen, __self__).__init__(
            'volcengine:cen/cen:Cen',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            cen_bandwidth_package_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            cen_id: Optional[pulumi.Input[str]] = None,
            cen_name: Optional[pulumi.Input[str]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CenTagArgs']]]]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Cen':
        """
        Get an existing Cen resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account ID of the cen.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cen_bandwidth_package_ids: A list of bandwidth package IDs of the cen.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        :param pulumi.Input[str] cen_name: The name of the cen.
        :param pulumi.Input[str] creation_time: The create time of the cen.
        :param pulumi.Input[str] description: The description of the cen.
        :param pulumi.Input[str] status: The status of the cen.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CenTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] update_time: The update time of the cen.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CenState.__new__(_CenState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["cen_bandwidth_package_ids"] = cen_bandwidth_package_ids
        __props__.__dict__["cen_id"] = cen_id
        __props__.__dict__["cen_name"] = cen_name
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["update_time"] = update_time
        return Cen(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The account ID of the cen.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="cenBandwidthPackageIds")
    def cen_bandwidth_package_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of bandwidth package IDs of the cen.
        """
        return pulumi.get(self, "cen_bandwidth_package_ids")

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The ID of the cen.
        """
        return pulumi.get(self, "cen_id")

    @property
    @pulumi.getter(name="cenName")
    def cen_name(self) -> pulumi.Output[str]:
        """
        The name of the cen.
        """
        return pulumi.get(self, "cen_name")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The create time of the cen.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the cen.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the cen.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.CenTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of the cen.
        """
        return pulumi.get(self, "update_time")

