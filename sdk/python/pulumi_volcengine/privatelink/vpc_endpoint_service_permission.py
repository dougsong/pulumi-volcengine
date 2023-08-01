# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcEndpointServicePermissionArgs', 'VpcEndpointServicePermission']

@pulumi.input_type
class VpcEndpointServicePermissionArgs:
    def __init__(__self__, *,
                 permit_account_id: pulumi.Input[str],
                 service_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a VpcEndpointServicePermission resource.
        :param pulumi.Input[str] permit_account_id: The id of account.
        :param pulumi.Input[str] service_id: The id of service.
        """
        pulumi.set(__self__, "permit_account_id", permit_account_id)
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="permitAccountId")
    def permit_account_id(self) -> pulumi.Input[str]:
        """
        The id of account.
        """
        return pulumi.get(self, "permit_account_id")

    @permit_account_id.setter
    def permit_account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "permit_account_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        """
        The id of service.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)


@pulumi.input_type
class _VpcEndpointServicePermissionState:
    def __init__(__self__, *,
                 permit_account_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointServicePermission resources.
        :param pulumi.Input[str] permit_account_id: The id of account.
        :param pulumi.Input[str] service_id: The id of service.
        """
        if permit_account_id is not None:
            pulumi.set(__self__, "permit_account_id", permit_account_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="permitAccountId")
    def permit_account_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of account.
        """
        return pulumi.get(self, "permit_account_id")

    @permit_account_id.setter
    def permit_account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "permit_account_id", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of service.
        """
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_id", value)


class VpcEndpointServicePermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permit_account_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage privatelink vpc endpoint service permission
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.privatelink.VpcEndpointServicePermission("foo",
            permit_account_id="210000000",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        foo1 = volcengine.privatelink.VpcEndpointServicePermission("foo1",
            permit_account_id="210000001",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        ```

        ## Import

        VpcEndpointServicePermission can be imported using the serviceId:permitAccountId, e.g.

        ```sh
         $ pulumi import volcengine:privatelink/vpcEndpointServicePermission:VpcEndpointServicePermission default epsvc-2fe630gurkl37k5gfuy33****:2100000000
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] permit_account_id: The id of account.
        :param pulumi.Input[str] service_id: The id of service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointServicePermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage privatelink vpc endpoint service permission
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.privatelink.VpcEndpointServicePermission("foo",
            permit_account_id="210000000",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        foo1 = volcengine.privatelink.VpcEndpointServicePermission("foo1",
            permit_account_id="210000001",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        ```

        ## Import

        VpcEndpointServicePermission can be imported using the serviceId:permitAccountId, e.g.

        ```sh
         $ pulumi import volcengine:privatelink/vpcEndpointServicePermission:VpcEndpointServicePermission default epsvc-2fe630gurkl37k5gfuy33****:2100000000
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointServicePermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointServicePermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 permit_account_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcEndpointServicePermissionArgs.__new__(VpcEndpointServicePermissionArgs)

            if permit_account_id is None and not opts.urn:
                raise TypeError("Missing required property 'permit_account_id'")
            __props__.__dict__["permit_account_id"] = permit_account_id
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
        super(VpcEndpointServicePermission, __self__).__init__(
            'volcengine:privatelink/vpcEndpointServicePermission:VpcEndpointServicePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            permit_account_id: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointServicePermission':
        """
        Get an existing VpcEndpointServicePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] permit_account_id: The id of account.
        :param pulumi.Input[str] service_id: The id of service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointServicePermissionState.__new__(_VpcEndpointServicePermissionState)

        __props__.__dict__["permit_account_id"] = permit_account_id
        __props__.__dict__["service_id"] = service_id
        return VpcEndpointServicePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="permitAccountId")
    def permit_account_id(self) -> pulumi.Output[str]:
        """
        The id of account.
        """
        return pulumi.get(self, "permit_account_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        The id of service.
        """
        return pulumi.get(self, "service_id")

