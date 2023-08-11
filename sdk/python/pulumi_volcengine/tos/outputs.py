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
    'BucketAccountAcl',
    'BucketObjectAccountAcl',
    'BucketObjectsObjectResult',
    'BucketsBucketResult',
]

@pulumi.output_type
class BucketAccountAcl(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"
        elif key == "aclType":
            suggest = "acl_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BucketAccountAcl. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BucketAccountAcl.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BucketAccountAcl.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_id: str,
                 permission: str,
                 acl_type: Optional[str] = None):
        """
        :param str account_id: The accountId to control.
        :param str permission: The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
        :param str acl_type: The acl type to control.Valid value is CanonicalUser.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "permission", permission)
        if acl_type is not None:
            pulumi.set(__self__, "acl_type", acl_type)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The accountId to control.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def permission(self) -> str:
        """
        The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter(name="aclType")
    def acl_type(self) -> Optional[str]:
        """
        The acl type to control.Valid value is CanonicalUser.
        """
        return pulumi.get(self, "acl_type")


@pulumi.output_type
class BucketObjectAccountAcl(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "accountId":
            suggest = "account_id"
        elif key == "aclType":
            suggest = "acl_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in BucketObjectAccountAcl. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        BucketObjectAccountAcl.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        BucketObjectAccountAcl.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 account_id: str,
                 permission: str,
                 acl_type: Optional[str] = None):
        """
        :param str account_id: The accountId to control.
        :param str permission: The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
        :param str acl_type: The acl type to control.Valid value is CanonicalUser.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "permission", permission)
        if acl_type is not None:
            pulumi.set(__self__, "acl_type", acl_type)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> str:
        """
        The accountId to control.
        """
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter
    def permission(self) -> str:
        """
        The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
        """
        return pulumi.get(self, "permission")

    @property
    @pulumi.getter(name="aclType")
    def acl_type(self) -> Optional[str]:
        """
        The acl type to control.Valid value is CanonicalUser.
        """
        return pulumi.get(self, "acl_type")


@pulumi.output_type
class BucketObjectsObjectResult(dict):
    def __init__(__self__, *,
                 name: str,
                 size: int,
                 storage_class: str):
        """
        :param str name: The name the TOS Object.
        :param int size: The name the TOS Object size.
        :param str storage_class: The name the TOS Object storage class.
        """
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "size", size)
        pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name the TOS Object.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def size(self) -> int:
        """
        The name the TOS Object size.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> str:
        """
        The name the TOS Object storage class.
        """
        return pulumi.get(self, "storage_class")


@pulumi.output_type
class BucketsBucketResult(dict):
    def __init__(__self__, *,
                 creation_date: str,
                 extranet_endpoint: str,
                 intranet_endpoint: str,
                 is_truncated: bool,
                 location: str,
                 marker: str,
                 max_keys: int,
                 name: str,
                 prefix: str):
        """
        :param str creation_date: The create date of the TOS bucket.
        :param str extranet_endpoint: The extranet endpoint of the TOS bucket.
        :param str intranet_endpoint: The intranet endpoint the TOS bucket.
        :param bool is_truncated: (**Deprecated**) The Field is Deprecated. The truncated the TOS bucket.
        :param str location: The location of the TOS bucket.
        :param str marker: (**Deprecated**) The Field is Deprecated. The marker the TOS bucket.
        :param int max_keys: (**Deprecated**) The Field is Deprecated. The max keys the TOS bucket.
        :param str name: The name the TOS bucket.
        :param str prefix: (**Deprecated**) The Field is Deprecated. The prefix the TOS bucket.
        """
        pulumi.set(__self__, "creation_date", creation_date)
        pulumi.set(__self__, "extranet_endpoint", extranet_endpoint)
        pulumi.set(__self__, "intranet_endpoint", intranet_endpoint)
        pulumi.set(__self__, "is_truncated", is_truncated)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "marker", marker)
        pulumi.set(__self__, "max_keys", max_keys)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "prefix", prefix)

    @property
    @pulumi.getter(name="creationDate")
    def creation_date(self) -> str:
        """
        The create date of the TOS bucket.
        """
        return pulumi.get(self, "creation_date")

    @property
    @pulumi.getter(name="extranetEndpoint")
    def extranet_endpoint(self) -> str:
        """
        The extranet endpoint of the TOS bucket.
        """
        return pulumi.get(self, "extranet_endpoint")

    @property
    @pulumi.getter(name="intranetEndpoint")
    def intranet_endpoint(self) -> str:
        """
        The intranet endpoint the TOS bucket.
        """
        return pulumi.get(self, "intranet_endpoint")

    @property
    @pulumi.getter(name="isTruncated")
    def is_truncated(self) -> bool:
        """
        (**Deprecated**) The Field is Deprecated. The truncated the TOS bucket.
        """
        warnings.warn("""The Field is Deprecated.""", DeprecationWarning)
        pulumi.log.warn("""is_truncated is deprecated: The Field is Deprecated.""")

        return pulumi.get(self, "is_truncated")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the TOS bucket.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def marker(self) -> str:
        """
        (**Deprecated**) The Field is Deprecated. The marker the TOS bucket.
        """
        warnings.warn("""The Field is Deprecated.""", DeprecationWarning)
        pulumi.log.warn("""marker is deprecated: The Field is Deprecated.""")

        return pulumi.get(self, "marker")

    @property
    @pulumi.getter(name="maxKeys")
    def max_keys(self) -> int:
        """
        (**Deprecated**) The Field is Deprecated. The max keys the TOS bucket.
        """
        warnings.warn("""The Field is Deprecated.""", DeprecationWarning)
        pulumi.log.warn("""max_keys is deprecated: The Field is Deprecated.""")

        return pulumi.get(self, "max_keys")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name the TOS bucket.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def prefix(self) -> str:
        """
        (**Deprecated**) The Field is Deprecated. The prefix the TOS bucket.
        """
        warnings.warn("""The Field is Deprecated.""", DeprecationWarning)
        pulumi.log.warn("""prefix is deprecated: The Field is Deprecated.""")

        return pulumi.get(self, "prefix")


