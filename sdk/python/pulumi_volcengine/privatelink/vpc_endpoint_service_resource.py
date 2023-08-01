# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['VpcEndpointServiceResourceInitArgs', 'VpcEndpointServiceResource']

@pulumi.input_type
class VpcEndpointServiceResourceInitArgs:
    def __init__(__self__, *,
                 resource_id: pulumi.Input[str],
                 service_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a VpcEndpointServiceResource resource.
        :param pulumi.Input[str] resource_id: The id of resource.
        :param pulumi.Input[str] service_id: The id of service.
        """
        pulumi.set(__self__, "resource_id", resource_id)
        pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Input[str]:
        """
        The id of resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_id", value)

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
class _VpcEndpointServiceResourceState:
    def __init__(__self__, *,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VpcEndpointServiceResource resources.
        :param pulumi.Input[str] resource_id: The id of resource.
        :param pulumi.Input[str] service_id: The id of service.
        """
        if resource_id is not None:
            pulumi.set(__self__, "resource_id", resource_id)
        if service_id is not None:
            pulumi.set(__self__, "service_id", service_id)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of resource.
        """
        return pulumi.get(self, "resource_id")

    @resource_id.setter
    def resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_id", value)

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


class VpcEndpointServiceResource(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage privatelink vpc endpoint service resource
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.privatelink.VpcEndpointServiceResource("foo",
            resource_id="clb-3reii8qfbp7gg5zsk2hsrbe3c",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        foo1 = volcengine.privatelink.VpcEndpointServiceResource("foo1",
            resource_id="clb-2d6sfye98rzls58ozfducee1o",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        foo2 = volcengine.privatelink.VpcEndpointServiceResource("foo2",
            resource_id="clb-3refkvae02gow5zsk2ilaev5y",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        ```

        ## Import

        VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.

        ```sh
         $ pulumi import volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_id: The id of resource.
        :param pulumi.Input[str] service_id: The id of service.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcEndpointServiceResourceInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage privatelink vpc endpoint service resource
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.privatelink.VpcEndpointServiceResource("foo",
            resource_id="clb-3reii8qfbp7gg5zsk2hsrbe3c",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        foo1 = volcengine.privatelink.VpcEndpointServiceResource("foo1",
            resource_id="clb-2d6sfye98rzls58ozfducee1o",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        foo2 = volcengine.privatelink.VpcEndpointServiceResource("foo2",
            resource_id="clb-3refkvae02gow5zsk2ilaev5y",
            service_id="epsvc-3rel73uf2ewao5zsk2j2l58ro")
        ```

        ## Import

        VpcEndpointServiceResource can be imported using the serviceId:resourceId, e.g.

        ```sh
         $ pulumi import volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource default epsvc-2fe630gurkl37k5gfuy33****:clb-bp1o94dp5i6ea****
        ```

        :param str resource_name: The name of the resource.
        :param VpcEndpointServiceResourceInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcEndpointServiceResourceInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 resource_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = VpcEndpointServiceResourceInitArgs.__new__(VpcEndpointServiceResourceInitArgs)

            if resource_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_id'")
            __props__.__dict__["resource_id"] = resource_id
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
        super(VpcEndpointServiceResource, __self__).__init__(
            'volcengine:privatelink/vpcEndpointServiceResource:VpcEndpointServiceResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            resource_id: Optional[pulumi.Input[str]] = None,
            service_id: Optional[pulumi.Input[str]] = None) -> 'VpcEndpointServiceResource':
        """
        Get an existing VpcEndpointServiceResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] resource_id: The id of resource.
        :param pulumi.Input[str] service_id: The id of service.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcEndpointServiceResourceState.__new__(_VpcEndpointServiceResourceState)

        __props__.__dict__["resource_id"] = resource_id
        __props__.__dict__["service_id"] = service_id
        return VpcEndpointServiceResource(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="resourceId")
    def resource_id(self) -> pulumi.Output[str]:
        """
        The id of resource.
        """
        return pulumi.get(self, "resource_id")

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Output[str]:
        """
        The id of service.
        """
        return pulumi.get(self, "service_id")

