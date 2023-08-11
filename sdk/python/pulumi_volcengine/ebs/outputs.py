# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'VolumesVolumeResult',
]

@pulumi.output_type
class VolumesVolumeResult(dict):
    def __init__(__self__, *,
                 billing_type: int,
                 created_at: str,
                 delete_with_instance: bool,
                 description: str,
                 device_name: str,
                 expired_time: str,
                 id: str,
                 image_id: str,
                 instance_id: str,
                 kind: str,
                 pay_type: str,
                 renew_type: int,
                 size: int,
                 status: str,
                 trade_status: int,
                 updated_at: str,
                 volume_id: str,
                 volume_name: str,
                 volume_type: str,
                 zone_id: str):
        """
        :param str instance_id: The Id of instance.
        :param str kind: The Kind of Volume.
        :param str volume_name: The name of Volume.
        :param str volume_type: The type of Volume.
        :param str zone_id: The Id of Zone.
        """
        pulumi.set(__self__, "billing_type", billing_type)
        pulumi.set(__self__, "created_at", created_at)
        pulumi.set(__self__, "delete_with_instance", delete_with_instance)
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "device_name", device_name)
        pulumi.set(__self__, "expired_time", expired_time)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "image_id", image_id)
        pulumi.set(__self__, "instance_id", instance_id)
        pulumi.set(__self__, "kind", kind)
        pulumi.set(__self__, "pay_type", pay_type)
        pulumi.set(__self__, "renew_type", renew_type)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "trade_status", trade_status)
        pulumi.set(__self__, "updated_at", updated_at)
        pulumi.set(__self__, "volume_id", volume_id)
        pulumi.set(__self__, "volume_name", volume_name)
        pulumi.set(__self__, "volume_type", volume_type)
        pulumi.set(__self__, "zone_id", zone_id)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> int:
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> str:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter(name="deleteWithInstance")
    def delete_with_instance(self) -> bool:
        return pulumi.get(self, "delete_with_instance")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="deviceName")
    def device_name(self) -> str:
        return pulumi.get(self, "device_name")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> str:
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="imageId")
    def image_id(self) -> str:
        return pulumi.get(self, "image_id")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> str:
        """
        The Id of instance.
        """
        return pulumi.get(self, "instance_id")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The Kind of Volume.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="payType")
    def pay_type(self) -> str:
        return pulumi.get(self, "pay_type")

    @property
    @pulumi.getter(name="renewType")
    def renew_type(self) -> int:
        return pulumi.get(self, "renew_type")

    @property
    @pulumi.getter
    def size(self) -> int:
        return pulumi.get(self, "size")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="tradeStatus")
    def trade_status(self) -> int:
        return pulumi.get(self, "trade_status")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> str:
        return pulumi.get(self, "updated_at")

    @property
    @pulumi.getter(name="volumeId")
    def volume_id(self) -> str:
        return pulumi.get(self, "volume_id")

    @property
    @pulumi.getter(name="volumeName")
    def volume_name(self) -> str:
        """
        The name of Volume.
        """
        return pulumi.get(self, "volume_name")

    @property
    @pulumi.getter(name="volumeType")
    def volume_type(self) -> str:
        """
        The type of Volume.
        """
        return pulumi.get(self, "volume_type")

    @property
    @pulumi.getter(name="zoneId")
    def zone_id(self) -> str:
        """
        The Id of Zone.
        """
        return pulumi.get(self, "zone_id")


