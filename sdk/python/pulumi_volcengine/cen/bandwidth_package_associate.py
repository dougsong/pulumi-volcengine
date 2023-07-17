# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['BandwidthPackageAssociateArgs', 'BandwidthPackageAssociate']

@pulumi.input_type
class BandwidthPackageAssociateArgs:
    def __init__(__self__, *,
                 cen_bandwidth_package_id: pulumi.Input[str],
                 cen_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a BandwidthPackageAssociate resource.
        :param pulumi.Input[str] cen_bandwidth_package_id: The ID of the cen bandwidth package.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        """
        pulumi.set(__self__, "cen_bandwidth_package_id", cen_bandwidth_package_id)
        pulumi.set(__self__, "cen_id", cen_id)

    @property
    @pulumi.getter(name="cenBandwidthPackageId")
    def cen_bandwidth_package_id(self) -> pulumi.Input[str]:
        """
        The ID of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_id")

    @cen_bandwidth_package_id.setter
    def cen_bandwidth_package_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "cen_bandwidth_package_id", value)

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


@pulumi.input_type
class _BandwidthPackageAssociateState:
    def __init__(__self__, *,
                 cen_bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BandwidthPackageAssociate resources.
        :param pulumi.Input[str] cen_bandwidth_package_id: The ID of the cen bandwidth package.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        """
        if cen_bandwidth_package_id is not None:
            pulumi.set(__self__, "cen_bandwidth_package_id", cen_bandwidth_package_id)
        if cen_id is not None:
            pulumi.set(__self__, "cen_id", cen_id)

    @property
    @pulumi.getter(name="cenBandwidthPackageId")
    def cen_bandwidth_package_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_id")

    @cen_bandwidth_package_id.setter
    def cen_bandwidth_package_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_bandwidth_package_id", value)

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


class BandwidthPackageAssociate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage cen bandwidth package associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.BandwidthPackageAssociate("foo",
            cen_bandwidth_package_id="cbp-2bzeew3s8p79c2dx0eeohej4x",
            cen_id="cen-2bzrl3srxsv0g2dx0efyoojn3")
        ```

        ## Import

        Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.

        ```sh
         $ pulumi import volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_bandwidth_package_id: The ID of the cen bandwidth package.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BandwidthPackageAssociateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage cen bandwidth package associate
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.BandwidthPackageAssociate("foo",
            cen_bandwidth_package_id="cbp-2bzeew3s8p79c2dx0eeohej4x",
            cen_id="cen-2bzrl3srxsv0g2dx0efyoojn3")
        ```

        ## Import

        Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.

        ```sh
         $ pulumi import volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
        ```

        :param str resource_name: The name of the resource.
        :param BandwidthPackageAssociateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BandwidthPackageAssociateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cen_bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cen_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = BandwidthPackageAssociateArgs.__new__(BandwidthPackageAssociateArgs)

            if cen_bandwidth_package_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_bandwidth_package_id'")
            __props__.__dict__["cen_bandwidth_package_id"] = cen_bandwidth_package_id
            if cen_id is None and not opts.urn:
                raise TypeError("Missing required property 'cen_id'")
            __props__.__dict__["cen_id"] = cen_id
        super(BandwidthPackageAssociate, __self__).__init__(
            'volcengine:cen/bandwidthPackageAssociate:BandwidthPackageAssociate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cen_bandwidth_package_id: Optional[pulumi.Input[str]] = None,
            cen_id: Optional[pulumi.Input[str]] = None) -> 'BandwidthPackageAssociate':
        """
        Get an existing BandwidthPackageAssociate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cen_bandwidth_package_id: The ID of the cen bandwidth package.
        :param pulumi.Input[str] cen_id: The ID of the cen.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BandwidthPackageAssociateState.__new__(_BandwidthPackageAssociateState)

        __props__.__dict__["cen_bandwidth_package_id"] = cen_bandwidth_package_id
        __props__.__dict__["cen_id"] = cen_id
        return BandwidthPackageAssociate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cenBandwidthPackageId")
    def cen_bandwidth_package_id(self) -> pulumi.Output[str]:
        """
        The ID of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_id")

    @property
    @pulumi.getter(name="cenId")
    def cen_id(self) -> pulumi.Output[str]:
        """
        The ID of the cen.
        """
        return pulumi.get(self, "cen_id")

