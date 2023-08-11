# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MongoAllowListArgs', 'MongoAllowList']

@pulumi.input_type
class MongoAllowListArgs:
    def __init__(__self__, *,
                 allow_list: pulumi.Input[str],
                 allow_list_name: pulumi.Input[str],
                 allow_list_desc: Optional[pulumi.Input[str]] = None,
                 allow_list_type: Optional[pulumi.Input[str]] = None,
                 modify_mode: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MongoAllowList resource.
        :param pulumi.Input[str] allow_list: IP address or IP address segment in CIDR format.
        :param pulumi.Input[str] allow_list_name: The name of allow list.
        :param pulumi.Input[str] allow_list_desc: The description of allow list.
        :param pulumi.Input[str] allow_list_type: The IP address type of allow list, valid value contains `IPv4`.
        :param pulumi.Input[str] modify_mode: The modify mode. Only support Cover mode.
        """
        pulumi.set(__self__, "allow_list", allow_list)
        pulumi.set(__self__, "allow_list_name", allow_list_name)
        if allow_list_desc is not None:
            pulumi.set(__self__, "allow_list_desc", allow_list_desc)
        if allow_list_type is not None:
            pulumi.set(__self__, "allow_list_type", allow_list_type)
        if modify_mode is not None:
            pulumi.set(__self__, "modify_mode", modify_mode)

    @property
    @pulumi.getter(name="allowList")
    def allow_list(self) -> pulumi.Input[str]:
        """
        IP address or IP address segment in CIDR format.
        """
        return pulumi.get(self, "allow_list")

    @allow_list.setter
    def allow_list(self, value: pulumi.Input[str]):
        pulumi.set(self, "allow_list", value)

    @property
    @pulumi.getter(name="allowListName")
    def allow_list_name(self) -> pulumi.Input[str]:
        """
        The name of allow list.
        """
        return pulumi.get(self, "allow_list_name")

    @allow_list_name.setter
    def allow_list_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "allow_list_name", value)

    @property
    @pulumi.getter(name="allowListDesc")
    def allow_list_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of allow list.
        """
        return pulumi.get(self, "allow_list_desc")

    @allow_list_desc.setter
    def allow_list_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list_desc", value)

    @property
    @pulumi.getter(name="allowListType")
    def allow_list_type(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address type of allow list, valid value contains `IPv4`.
        """
        return pulumi.get(self, "allow_list_type")

    @allow_list_type.setter
    def allow_list_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list_type", value)

    @property
    @pulumi.getter(name="modifyMode")
    def modify_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The modify mode. Only support Cover mode.
        """
        return pulumi.get(self, "modify_mode")

    @modify_mode.setter
    def modify_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_mode", value)


@pulumi.input_type
class _MongoAllowListState:
    def __init__(__self__, *,
                 allow_list: Optional[pulumi.Input[str]] = None,
                 allow_list_desc: Optional[pulumi.Input[str]] = None,
                 allow_list_name: Optional[pulumi.Input[str]] = None,
                 allow_list_type: Optional[pulumi.Input[str]] = None,
                 modify_mode: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MongoAllowList resources.
        :param pulumi.Input[str] allow_list: IP address or IP address segment in CIDR format.
        :param pulumi.Input[str] allow_list_desc: The description of allow list.
        :param pulumi.Input[str] allow_list_name: The name of allow list.
        :param pulumi.Input[str] allow_list_type: The IP address type of allow list, valid value contains `IPv4`.
        :param pulumi.Input[str] modify_mode: The modify mode. Only support Cover mode.
        """
        if allow_list is not None:
            pulumi.set(__self__, "allow_list", allow_list)
        if allow_list_desc is not None:
            pulumi.set(__self__, "allow_list_desc", allow_list_desc)
        if allow_list_name is not None:
            pulumi.set(__self__, "allow_list_name", allow_list_name)
        if allow_list_type is not None:
            pulumi.set(__self__, "allow_list_type", allow_list_type)
        if modify_mode is not None:
            pulumi.set(__self__, "modify_mode", modify_mode)

    @property
    @pulumi.getter(name="allowList")
    def allow_list(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or IP address segment in CIDR format.
        """
        return pulumi.get(self, "allow_list")

    @allow_list.setter
    def allow_list(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list", value)

    @property
    @pulumi.getter(name="allowListDesc")
    def allow_list_desc(self) -> Optional[pulumi.Input[str]]:
        """
        The description of allow list.
        """
        return pulumi.get(self, "allow_list_desc")

    @allow_list_desc.setter
    def allow_list_desc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list_desc", value)

    @property
    @pulumi.getter(name="allowListName")
    def allow_list_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of allow list.
        """
        return pulumi.get(self, "allow_list_name")

    @allow_list_name.setter
    def allow_list_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list_name", value)

    @property
    @pulumi.getter(name="allowListType")
    def allow_list_type(self) -> Optional[pulumi.Input[str]]:
        """
        The IP address type of allow list, valid value contains `IPv4`.
        """
        return pulumi.get(self, "allow_list_type")

    @allow_list_type.setter
    def allow_list_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list_type", value)

    @property
    @pulumi.getter(name="modifyMode")
    def modify_mode(self) -> Optional[pulumi.Input[str]]:
        """
        The modify mode. Only support Cover mode.
        """
        return pulumi.get(self, "modify_mode")

    @modify_mode.setter
    def modify_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_mode", value)


class MongoAllowList(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_list: Optional[pulumi.Input[str]] = None,
                 allow_list_desc: Optional[pulumi.Input[str]] = None,
                 allow_list_name: Optional[pulumi.Input[str]] = None,
                 allow_list_type: Optional[pulumi.Input[str]] = None,
                 modify_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage mongodb allow list
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.mongodb.MongoAllowList("foo",
            allow_list="10.1.1.3,10.2.3.0/24,10.1.1.1",
            allow_list_desc="test1",
            allow_list_name="tf-test-hh",
            allow_list_type="IPv4")
        ```

        ## Import

        mongodb allow list can be imported using the allowListId, e.g.

        ```sh
         $ pulumi import volcengine:mongodb/mongoAllowList:MongoAllowList default acl-d1fd76693bd54e658912e7337d5b****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_list: IP address or IP address segment in CIDR format.
        :param pulumi.Input[str] allow_list_desc: The description of allow list.
        :param pulumi.Input[str] allow_list_name: The name of allow list.
        :param pulumi.Input[str] allow_list_type: The IP address type of allow list, valid value contains `IPv4`.
        :param pulumi.Input[str] modify_mode: The modify mode. Only support Cover mode.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MongoAllowListArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage mongodb allow list
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.mongodb.MongoAllowList("foo",
            allow_list="10.1.1.3,10.2.3.0/24,10.1.1.1",
            allow_list_desc="test1",
            allow_list_name="tf-test-hh",
            allow_list_type="IPv4")
        ```

        ## Import

        mongodb allow list can be imported using the allowListId, e.g.

        ```sh
         $ pulumi import volcengine:mongodb/mongoAllowList:MongoAllowList default acl-d1fd76693bd54e658912e7337d5b****
        ```

        :param str resource_name: The name of the resource.
        :param MongoAllowListArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MongoAllowListArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_list: Optional[pulumi.Input[str]] = None,
                 allow_list_desc: Optional[pulumi.Input[str]] = None,
                 allow_list_name: Optional[pulumi.Input[str]] = None,
                 allow_list_type: Optional[pulumi.Input[str]] = None,
                 modify_mode: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MongoAllowListArgs.__new__(MongoAllowListArgs)

            if allow_list is None and not opts.urn:
                raise TypeError("Missing required property 'allow_list'")
            __props__.__dict__["allow_list"] = allow_list
            __props__.__dict__["allow_list_desc"] = allow_list_desc
            if allow_list_name is None and not opts.urn:
                raise TypeError("Missing required property 'allow_list_name'")
            __props__.__dict__["allow_list_name"] = allow_list_name
            __props__.__dict__["allow_list_type"] = allow_list_type
            __props__.__dict__["modify_mode"] = modify_mode
        super(MongoAllowList, __self__).__init__(
            'volcengine:mongodb/mongoAllowList:MongoAllowList',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_list: Optional[pulumi.Input[str]] = None,
            allow_list_desc: Optional[pulumi.Input[str]] = None,
            allow_list_name: Optional[pulumi.Input[str]] = None,
            allow_list_type: Optional[pulumi.Input[str]] = None,
            modify_mode: Optional[pulumi.Input[str]] = None) -> 'MongoAllowList':
        """
        Get an existing MongoAllowList resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_list: IP address or IP address segment in CIDR format.
        :param pulumi.Input[str] allow_list_desc: The description of allow list.
        :param pulumi.Input[str] allow_list_name: The name of allow list.
        :param pulumi.Input[str] allow_list_type: The IP address type of allow list, valid value contains `IPv4`.
        :param pulumi.Input[str] modify_mode: The modify mode. Only support Cover mode.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MongoAllowListState.__new__(_MongoAllowListState)

        __props__.__dict__["allow_list"] = allow_list
        __props__.__dict__["allow_list_desc"] = allow_list_desc
        __props__.__dict__["allow_list_name"] = allow_list_name
        __props__.__dict__["allow_list_type"] = allow_list_type
        __props__.__dict__["modify_mode"] = modify_mode
        return MongoAllowList(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowList")
    def allow_list(self) -> pulumi.Output[str]:
        """
        IP address or IP address segment in CIDR format.
        """
        return pulumi.get(self, "allow_list")

    @property
    @pulumi.getter(name="allowListDesc")
    def allow_list_desc(self) -> pulumi.Output[str]:
        """
        The description of allow list.
        """
        return pulumi.get(self, "allow_list_desc")

    @property
    @pulumi.getter(name="allowListName")
    def allow_list_name(self) -> pulumi.Output[str]:
        """
        The name of allow list.
        """
        return pulumi.get(self, "allow_list_name")

    @property
    @pulumi.getter(name="allowListType")
    def allow_list_type(self) -> pulumi.Output[Optional[str]]:
        """
        The IP address type of allow list, valid value contains `IPv4`.
        """
        return pulumi.get(self, "allow_list_type")

    @property
    @pulumi.getter(name="modifyMode")
    def modify_mode(self) -> pulumi.Output[Optional[str]]:
        """
        The modify mode. Only support Cover mode.
        """
        return pulumi.get(self, "modify_mode")

