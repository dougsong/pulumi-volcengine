# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AllowListAssociateArgs', 'AllowListAssociate']

@pulumi.input_type
class AllowListAssociateArgs:
    def __init__(__self__, *,
                 allow_list_id: pulumi.Input[str],
                 instance_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a AllowListAssociate resource.
        :param pulumi.Input[str] allow_list_id: Id of allow list to associate.
        :param pulumi.Input[str] instance_id: Id of instance to associate.
        """
        pulumi.set(__self__, "allow_list_id", allow_list_id)
        pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="allowListId")
    def allow_list_id(self) -> pulumi.Input[str]:
        """
        Id of allow list to associate.
        """
        return pulumi.get(self, "allow_list_id")

    @allow_list_id.setter
    def allow_list_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "allow_list_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        Id of instance to associate.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)


@pulumi.input_type
class _AllowListAssociateState:
    def __init__(__self__, *,
                 allow_list_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AllowListAssociate resources.
        :param pulumi.Input[str] allow_list_id: Id of allow list to associate.
        :param pulumi.Input[str] instance_id: Id of instance to associate.
        """
        if allow_list_id is not None:
            pulumi.set(__self__, "allow_list_id", allow_list_id)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="allowListId")
    def allow_list_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of allow list to associate.
        """
        return pulumi.get(self, "allow_list_id")

    @allow_list_id.setter
    def allow_list_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allow_list_id", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of instance to associate.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class AllowListAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_list_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage redis allow list associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        default = volcengine.redis.AllowListAssociate("default",
            allow_list_id="acl-cnlfc5zsxscu1gg2ajh",
            instance_id="redis-cnlfbzifs7bpvundz")
        default1 = volcengine.redis.AllowListAssociate("default1",
            allow_list_id="acl-cnlff2mb31zo85p5am7",
            instance_id="redis-cnlfbzifs7bpvundz")
        ```

        ## Import

        Redis AllowList Association can be imported using the instanceId:allowListId, e.g.

        ```sh
         $ pulumi import volcengine:redis/allowListAssociate:AllowListAssociate default redis-asdljioeixxxx:acl-cn03wk541s55c376xxxx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_list_id: Id of allow list to associate.
        :param pulumi.Input[str] instance_id: Id of instance to associate.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AllowListAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage redis allow list associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        default = volcengine.redis.AllowListAssociate("default",
            allow_list_id="acl-cnlfc5zsxscu1gg2ajh",
            instance_id="redis-cnlfbzifs7bpvundz")
        default1 = volcengine.redis.AllowListAssociate("default1",
            allow_list_id="acl-cnlff2mb31zo85p5am7",
            instance_id="redis-cnlfbzifs7bpvundz")
        ```

        ## Import

        Redis AllowList Association can be imported using the instanceId:allowListId, e.g.

        ```sh
         $ pulumi import volcengine:redis/allowListAssociate:AllowListAssociate default redis-asdljioeixxxx:acl-cn03wk541s55c376xxxx
        ```

        :param str resource_name: The name of the resource.
        :param AllowListAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AllowListAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_list_id: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AllowListAssociateArgs.__new__(AllowListAssociateArgs)

            if allow_list_id is None and not opts.urn:
                raise TypeError("Missing required property 'allow_list_id'")
            __props__.__dict__["allow_list_id"] = allow_list_id
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        super(AllowListAssociate, __self__).__init__(
            'volcengine:redis/allowListAssociate:AllowListAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_list_id: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'AllowListAssociate':
        """
        Get an existing AllowListAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allow_list_id: Id of allow list to associate.
        :param pulumi.Input[str] instance_id: Id of instance to associate.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AllowListAssociateState.__new__(_AllowListAssociateState)

        __props__.__dict__["allow_list_id"] = allow_list_id
        __props__.__dict__["instance_id"] = instance_id
        return AllowListAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowListId")
    def allow_list_id(self) -> pulumi.Output[str]:
        """
        Id of allow list to associate.
        """
        return pulumi.get(self, "allow_list_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        Id of instance to associate.
        """
        return pulumi.get(self, "instance_id")

