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

__all__ = ['BandwidthPackageArgs', 'BandwidthPackage']

@pulumi.input_type
class BandwidthPackageArgs:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 cen_bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 local_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 peer_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]] = None):
        """
        The set of arguments for constructing a BandwidthPackage resource.
        :param pulumi.Input[int] bandwidth: The bandwidth of the cen bandwidth package.
        :param pulumi.Input[str] billing_type: The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
               state file, not actually remove.
        :param pulumi.Input[str] cen_bandwidth_package_name: The name of the cen bandwidth package.
        :param pulumi.Input[str] description: The description of the cen bandwidth package.
        :param pulumi.Input[str] local_geographic_region_set_id: The local geographic region set id of the cen bandwidth package.
        :param pulumi.Input[str] peer_geographic_region_set_id: The peer geographic region set id of the cen bandwidth package.
        :param pulumi.Input[int] period: The period of the cen bandwidth package.
        :param pulumi.Input[str] period_unit: The period unit of the cen bandwidth package.
        :param pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]] tags: Tags.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if cen_bandwidth_package_name is not None:
            pulumi.set(__self__, "cen_bandwidth_package_name", cen_bandwidth_package_name)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if local_geographic_region_set_id is not None:
            pulumi.set(__self__, "local_geographic_region_set_id", local_geographic_region_set_id)
        if peer_geographic_region_set_id is not None:
            pulumi.set(__self__, "peer_geographic_region_set_id", peer_geographic_region_set_id)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The bandwidth of the cen bandwidth package.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
        state file, not actually remove.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter(name="cenBandwidthPackageName")
    def cen_bandwidth_package_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_name")

    @cen_bandwidth_package_name.setter
    def cen_bandwidth_package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_bandwidth_package_name", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cen bandwidth package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="localGeographicRegionSetId")
    def local_geographic_region_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The local geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "local_geographic_region_set_id")

    @local_geographic_region_set_id.setter
    def local_geographic_region_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_geographic_region_set_id", value)

    @property
    @pulumi.getter(name="peerGeographicRegionSetId")
    def peer_geographic_region_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The peer geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "peer_geographic_region_set_id")

    @peer_geographic_region_set_id.setter
    def peer_geographic_region_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_geographic_region_set_id", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The period of the cen bandwidth package.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The period unit of the cen bandwidth package.
        """
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _BandwidthPackageState:
    def __init__(__self__, *,
                 account_id: Optional[pulumi.Input[str]] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 business_status: Optional[pulumi.Input[str]] = None,
                 cen_bandwidth_package_id: Optional[pulumi.Input[str]] = None,
                 cen_bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 cen_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 deleted_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 expired_time: Optional[pulumi.Input[str]] = None,
                 local_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 peer_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None,
                 remaining_bandwidth: Optional[pulumi.Input[int]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BandwidthPackage resources.
        :param pulumi.Input[str] account_id: The account ID of the cen bandwidth package.
        :param pulumi.Input[int] bandwidth: The bandwidth of the cen bandwidth package.
        :param pulumi.Input[str] billing_type: The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
               state file, not actually remove.
        :param pulumi.Input[str] business_status: The business status of the cen bandwidth package.
        :param pulumi.Input[str] cen_bandwidth_package_id: The ID of the cen bandwidth package.
        :param pulumi.Input[str] cen_bandwidth_package_name: The name of the cen bandwidth package.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cen_ids: The cen IDs of the bandwidth package.
        :param pulumi.Input[str] creation_time: The create time of the cen bandwidth package.
        :param pulumi.Input[str] deleted_time: The deleted time of the cen bandwidth package.
        :param pulumi.Input[str] description: The description of the cen bandwidth package.
        :param pulumi.Input[str] expired_time: The expired time of the cen bandwidth package.
        :param pulumi.Input[str] local_geographic_region_set_id: The local geographic region set id of the cen bandwidth package.
        :param pulumi.Input[str] peer_geographic_region_set_id: The peer geographic region set id of the cen bandwidth package.
        :param pulumi.Input[int] period: The period of the cen bandwidth package.
        :param pulumi.Input[str] period_unit: The period unit of the cen bandwidth package.
        :param pulumi.Input[int] remaining_bandwidth: The remain bandwidth of the cen bandwidth package.
        :param pulumi.Input[str] status: The status of the cen bandwidth package.
        :param pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]] tags: Tags.
        :param pulumi.Input[str] update_time: The update time of the cen bandwidth package.
        """
        if account_id is not None:
            pulumi.set(__self__, "account_id", account_id)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if business_status is not None:
            pulumi.set(__self__, "business_status", business_status)
        if cen_bandwidth_package_id is not None:
            pulumi.set(__self__, "cen_bandwidth_package_id", cen_bandwidth_package_id)
        if cen_bandwidth_package_name is not None:
            pulumi.set(__self__, "cen_bandwidth_package_name", cen_bandwidth_package_name)
        if cen_ids is not None:
            pulumi.set(__self__, "cen_ids", cen_ids)
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if deleted_time is not None:
            pulumi.set(__self__, "deleted_time", deleted_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if expired_time is not None:
            pulumi.set(__self__, "expired_time", expired_time)
        if local_geographic_region_set_id is not None:
            pulumi.set(__self__, "local_geographic_region_set_id", local_geographic_region_set_id)
        if peer_geographic_region_set_id is not None:
            pulumi.set(__self__, "peer_geographic_region_set_id", peer_geographic_region_set_id)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if period_unit is not None:
            pulumi.set(__self__, "period_unit", period_unit)
        if remaining_bandwidth is not None:
            pulumi.set(__self__, "remaining_bandwidth", remaining_bandwidth)
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
        The account ID of the cen bandwidth package.
        """
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The bandwidth of the cen bandwidth package.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
        state file, not actually remove.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> Optional[pulumi.Input[str]]:
        """
        The business status of the cen bandwidth package.
        """
        return pulumi.get(self, "business_status")

    @business_status.setter
    def business_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "business_status", value)

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
    @pulumi.getter(name="cenBandwidthPackageName")
    def cen_bandwidth_package_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_name")

    @cen_bandwidth_package_name.setter
    def cen_bandwidth_package_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cen_bandwidth_package_name", value)

    @property
    @pulumi.getter(name="cenIds")
    def cen_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The cen IDs of the bandwidth package.
        """
        return pulumi.get(self, "cen_ids")

    @cen_ids.setter
    def cen_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cen_ids", value)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the cen bandwidth package.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> Optional[pulumi.Input[str]]:
        """
        The deleted time of the cen bandwidth package.
        """
        return pulumi.get(self, "deleted_time")

    @deleted_time.setter
    def deleted_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deleted_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the cen bandwidth package.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> Optional[pulumi.Input[str]]:
        """
        The expired time of the cen bandwidth package.
        """
        return pulumi.get(self, "expired_time")

    @expired_time.setter
    def expired_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expired_time", value)

    @property
    @pulumi.getter(name="localGeographicRegionSetId")
    def local_geographic_region_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The local geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "local_geographic_region_set_id")

    @local_geographic_region_set_id.setter
    def local_geographic_region_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_geographic_region_set_id", value)

    @property
    @pulumi.getter(name="peerGeographicRegionSetId")
    def peer_geographic_region_set_id(self) -> Optional[pulumi.Input[str]]:
        """
        The peer geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "peer_geographic_region_set_id")

    @peer_geographic_region_set_id.setter
    def peer_geographic_region_set_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_geographic_region_set_id", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The period of the cen bandwidth package.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> Optional[pulumi.Input[str]]:
        """
        The period unit of the cen bandwidth package.
        """
        return pulumi.get(self, "period_unit")

    @period_unit.setter
    def period_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "period_unit", value)

    @property
    @pulumi.getter(name="remainingBandwidth")
    def remaining_bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The remain bandwidth of the cen bandwidth package.
        """
        return pulumi.get(self, "remaining_bandwidth")

    @remaining_bandwidth.setter
    def remaining_bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "remaining_bandwidth", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the cen bandwidth package.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BandwidthPackageTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of the cen bandwidth package.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class BandwidthPackage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 cen_bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 local_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 peer_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.BandwidthPackage("foo",
            bandwidth=32,
            billing_type="PrePaid",
            cen_bandwidth_package_name="tf-test",
            description="tf-test1",
            local_geographic_region_set_id="China",
            peer_geographic_region_set_id="China",
            period=1,
            period_unit="Year")
        ```

        ## Import

        CenBandwidthPackage can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:cen/bandwidthPackage:BandwidthPackage default cbp-4c2zaavbvh5f42****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: The bandwidth of the cen bandwidth package.
        :param pulumi.Input[str] billing_type: The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
               state file, not actually remove.
        :param pulumi.Input[str] cen_bandwidth_package_name: The name of the cen bandwidth package.
        :param pulumi.Input[str] description: The description of the cen bandwidth package.
        :param pulumi.Input[str] local_geographic_region_set_id: The local geographic region set id of the cen bandwidth package.
        :param pulumi.Input[str] peer_geographic_region_set_id: The peer geographic region set id of the cen bandwidth package.
        :param pulumi.Input[int] period: The period of the cen bandwidth package.
        :param pulumi.Input[str] period_unit: The period unit of the cen bandwidth package.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]] tags: Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[BandwidthPackageArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cen.BandwidthPackage("foo",
            bandwidth=32,
            billing_type="PrePaid",
            cen_bandwidth_package_name="tf-test",
            description="tf-test1",
            local_geographic_region_set_id="China",
            peer_geographic_region_set_id="China",
            period=1,
            period_unit="Year")
        ```

        ## Import

        CenBandwidthPackage can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:cen/bandwidthPackage:BandwidthPackage default cbp-4c2zaavbvh5f42****
        ```

        :param str resource_name: The name of the resource.
        :param BandwidthPackageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BandwidthPackageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 cen_bandwidth_package_name: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 local_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 peer_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 period_unit: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]]] = None,
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
            __props__ = BandwidthPackageArgs.__new__(BandwidthPackageArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            __props__.__dict__["billing_type"] = billing_type
            __props__.__dict__["cen_bandwidth_package_name"] = cen_bandwidth_package_name
            __props__.__dict__["description"] = description
            __props__.__dict__["local_geographic_region_set_id"] = local_geographic_region_set_id
            __props__.__dict__["peer_geographic_region_set_id"] = peer_geographic_region_set_id
            __props__.__dict__["period"] = period
            __props__.__dict__["period_unit"] = period_unit
            __props__.__dict__["tags"] = tags
            __props__.__dict__["account_id"] = None
            __props__.__dict__["business_status"] = None
            __props__.__dict__["cen_bandwidth_package_id"] = None
            __props__.__dict__["cen_ids"] = None
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["deleted_time"] = None
            __props__.__dict__["expired_time"] = None
            __props__.__dict__["remaining_bandwidth"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["update_time"] = None
        super(BandwidthPackage, __self__).__init__(
            'volcengine:cen/bandwidthPackage:BandwidthPackage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_id: Optional[pulumi.Input[str]] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            billing_type: Optional[pulumi.Input[str]] = None,
            business_status: Optional[pulumi.Input[str]] = None,
            cen_bandwidth_package_id: Optional[pulumi.Input[str]] = None,
            cen_bandwidth_package_name: Optional[pulumi.Input[str]] = None,
            cen_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            deleted_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            expired_time: Optional[pulumi.Input[str]] = None,
            local_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
            peer_geographic_region_set_id: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            period_unit: Optional[pulumi.Input[str]] = None,
            remaining_bandwidth: Optional[pulumi.Input[int]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'BandwidthPackage':
        """
        Get an existing BandwidthPackage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_id: The account ID of the cen bandwidth package.
        :param pulumi.Input[int] bandwidth: The bandwidth of the cen bandwidth package.
        :param pulumi.Input[str] billing_type: The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
               state file, not actually remove.
        :param pulumi.Input[str] business_status: The business status of the cen bandwidth package.
        :param pulumi.Input[str] cen_bandwidth_package_id: The ID of the cen bandwidth package.
        :param pulumi.Input[str] cen_bandwidth_package_name: The name of the cen bandwidth package.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cen_ids: The cen IDs of the bandwidth package.
        :param pulumi.Input[str] creation_time: The create time of the cen bandwidth package.
        :param pulumi.Input[str] deleted_time: The deleted time of the cen bandwidth package.
        :param pulumi.Input[str] description: The description of the cen bandwidth package.
        :param pulumi.Input[str] expired_time: The expired time of the cen bandwidth package.
        :param pulumi.Input[str] local_geographic_region_set_id: The local geographic region set id of the cen bandwidth package.
        :param pulumi.Input[str] peer_geographic_region_set_id: The peer geographic region set id of the cen bandwidth package.
        :param pulumi.Input[int] period: The period of the cen bandwidth package.
        :param pulumi.Input[str] period_unit: The period unit of the cen bandwidth package.
        :param pulumi.Input[int] remaining_bandwidth: The remain bandwidth of the cen bandwidth package.
        :param pulumi.Input[str] status: The status of the cen bandwidth package.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BandwidthPackageTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] update_time: The update time of the cen bandwidth package.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BandwidthPackageState.__new__(_BandwidthPackageState)

        __props__.__dict__["account_id"] = account_id
        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["billing_type"] = billing_type
        __props__.__dict__["business_status"] = business_status
        __props__.__dict__["cen_bandwidth_package_id"] = cen_bandwidth_package_id
        __props__.__dict__["cen_bandwidth_package_name"] = cen_bandwidth_package_name
        __props__.__dict__["cen_ids"] = cen_ids
        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["deleted_time"] = deleted_time
        __props__.__dict__["description"] = description
        __props__.__dict__["expired_time"] = expired_time
        __props__.__dict__["local_geographic_region_set_id"] = local_geographic_region_set_id
        __props__.__dict__["peer_geographic_region_set_id"] = peer_geographic_region_set_id
        __props__.__dict__["period"] = period
        __props__.__dict__["period_unit"] = period_unit
        __props__.__dict__["remaining_bandwidth"] = remaining_bandwidth
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["update_time"] = update_time
        return BandwidthPackage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        """
        The account ID of the cen bandwidth package.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        The bandwidth of the cen bandwidth package.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> pulumi.Output[Optional[str]]:
        """
        The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the
        state file, not actually remove.
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="businessStatus")
    def business_status(self) -> pulumi.Output[str]:
        """
        The business status of the cen bandwidth package.
        """
        return pulumi.get(self, "business_status")

    @property
    @pulumi.getter(name="cenBandwidthPackageId")
    def cen_bandwidth_package_id(self) -> pulumi.Output[str]:
        """
        The ID of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_id")

    @property
    @pulumi.getter(name="cenBandwidthPackageName")
    def cen_bandwidth_package_name(self) -> pulumi.Output[str]:
        """
        The name of the cen bandwidth package.
        """
        return pulumi.get(self, "cen_bandwidth_package_name")

    @property
    @pulumi.getter(name="cenIds")
    def cen_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The cen IDs of the bandwidth package.
        """
        return pulumi.get(self, "cen_ids")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        The create time of the cen bandwidth package.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> pulumi.Output[str]:
        """
        The deleted time of the cen bandwidth package.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the cen bandwidth package.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> pulumi.Output[str]:
        """
        The expired time of the cen bandwidth package.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter(name="localGeographicRegionSetId")
    def local_geographic_region_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        The local geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "local_geographic_region_set_id")

    @property
    @pulumi.getter(name="peerGeographicRegionSetId")
    def peer_geographic_region_set_id(self) -> pulumi.Output[Optional[str]]:
        """
        The peer geographic region set id of the cen bandwidth package.
        """
        return pulumi.get(self, "peer_geographic_region_set_id")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        The period of the cen bandwidth package.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="periodUnit")
    def period_unit(self) -> pulumi.Output[Optional[str]]:
        """
        The period unit of the cen bandwidth package.
        """
        return pulumi.get(self, "period_unit")

    @property
    @pulumi.getter(name="remainingBandwidth")
    def remaining_bandwidth(self) -> pulumi.Output[int]:
        """
        The remain bandwidth of the cen bandwidth package.
        """
        return pulumi.get(self, "remaining_bandwidth")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the cen bandwidth package.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.BandwidthPackageTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The update time of the cen bandwidth package.
        """
        return pulumi.get(self, "update_time")

