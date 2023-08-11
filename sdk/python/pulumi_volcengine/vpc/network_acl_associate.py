# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['NetworkAclAssociateArgs', 'NetworkAclAssociate']

@pulumi.input_type
class NetworkAclAssociateArgs:
    def __init__(__self__, *,
                 network_acl_id: pulumi.Input[str],
                 resource_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a NetworkAclAssociate resource.
        :param pulumi.Input[str] network_acl_id: The id of Network Acl.
        :param pulumi.Input[str] resource_id: The resource id of Network Acl.
        """
        pulumi.set(__self__, "network_acl_id", network_acl_id)
        pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> pulumi.Input[str]:
        """
        The id of Network Acl.
        """
        return pulumi.get(self, "network_acl_id")

    @network_acl_id.setter
    def network_acl_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_acl_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The resource id of Network Acl.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)


@pulumi.input_type
class _NetworkAclAssociateState:
    def __init__(__self__, *,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkAclAssociate resources.
        :param pulumi.Input[str] network_acl_id: The id of Network Acl.
        :param pulumi.Input[str] resource_id: The resource id of Network Acl.
        """
        if network_acl_id is not None:
            pulumi.set(__self__, "network_acl_id", network_acl_id)
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of Network Acl.
        """
        return pulumi.get(self, "network_acl_id")

    @network_acl_id.setter
    def network_acl_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_acl_id", value)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The resource id of Network Acl.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)


class NetworkAclAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage network acl associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.NetworkAcl("foo",
            vpc_id="vpc-ru0wv9alfoxsu3nuld85rpp",
            network_acl_name="tf-test-acl")
        foo1 = volcengine.vpc.NetworkAclAssociate("foo1",
            network_acl_id=foo.id,
            resource_id="subnet-637jxq81u5mon3gd6ivc7rj")
        ```

        ## Import

        NetworkAcl associate can be imported using the network_acl_id:resource_id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/networkAclAssociate:NetworkAclAssociate default nacl-172leak37mi9s4d1w33pswqkh:subnet-637jxq81u5mon3gd6ivc7rj
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] network_acl_id: The id of Network Acl.
        :param pulumi.Input[str] resource_id: The resource id of Network Acl.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkAclAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage network acl associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.vpc.NetworkAcl("foo",
            vpc_id="vpc-ru0wv9alfoxsu3nuld85rpp",
            network_acl_name="tf-test-acl")
        foo1 = volcengine.vpc.NetworkAclAssociate("foo1",
            network_acl_id=foo.id,
            resource_id="subnet-637jxq81u5mon3gd6ivc7rj")
        ```

        ## Import

        NetworkAcl associate can be imported using the network_acl_id:resource_id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/networkAclAssociate:NetworkAclAssociate default nacl-172leak37mi9s4d1w33pswqkh:subnet-637jxq81u5mon3gd6ivc7rj
        ```

        :param str resource_name: The name of the resource.
        :param NetworkAclAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkAclAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 network_acl_id: Optional[pulumi.Input[str]] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkAclAssociateArgs.__new__(NetworkAclAssociateArgs)

            if network_acl_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_acl_id'")
            __props__.__dict__["network_acl_id"] = network_acl_id
            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
        super(NetworkAclAssociate, __self__).__init__(
            'volcengine:vpc/networkAclAssociate:NetworkAclAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            network_acl_id: Optional[pulumi.Input[str]] = None,
            resource_id: Optional[pulumi.Input[str]] = None) -> 'NetworkAclAssociate':
        """
        Get an existing NetworkAclAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] network_acl_id: The id of Network Acl.
        :param pulumi.Input[str] resource_id: The resource id of Network Acl.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkAclAssociateState.__new__(_NetworkAclAssociateState)

        __props__.__dict__["network_acl_id"] = network_acl_id
        __props__.__dict__["resource_id"] = resource_id
        return NetworkAclAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="networkAclId")
    def network_acl_id(self) -> pulumi.Output[str]:
        """
        The id of Network Acl.
        """
        return pulumi.get(self, "network_acl_id")

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The resource id of Network Acl.
        """
        return pulumi.get(self, "resource_id")

