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

__all__ = ['BucketObjectArgs', 'BucketObject']

@pulumi.input_type
class BucketObjectArgs:
    def __init__(__self__, *,
                 bucket_name: pulumi.Input[str],
                 file_path: pulumi.Input[str],
                 object_name: pulumi.Input[str],
                 account_acls: Optional[pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]]] = None,
                 content_md5: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 public_acl: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BucketObject resource.
        :param pulumi.Input[str] bucket_name: The name of the bucket.
        :param pulumi.Input[str] file_path: The file path for upload.
        :param pulumi.Input[str] object_name: The name of the object.
        :param pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]] account_acls: The user set of grant full control.
        :param pulumi.Input[str] content_md5: The file md5 sum (32-bit hexadecimal string) for upload.
        :param pulumi.Input[str] content_type: The content type of the object.
        :param pulumi.Input[str] encryption: The encryption of the object.Valid value is AES256.
        :param pulumi.Input[str] public_acl: The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        :param pulumi.Input[str] storage_class: The storage type of the object.Valid value is STANDARD|IA.
        """
        pulumi.set(__self__, "bucket_name", bucket_name)
        pulumi.set(__self__, "file_path", file_path)
        pulumi.set(__self__, "object_name", object_name)
        if account_acls is not None:
            pulumi.set(__self__, "account_acls", account_acls)
        if content_md5 is not None:
            pulumi.set(__self__, "content_md5", content_md5)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if public_acl is not None:
            pulumi.set(__self__, "public_acl", public_acl)
        if storage_class is not None:
            pulumi.set(__self__, "storage_class", storage_class)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Input[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> pulumi.Input[str]:
        """
        The file path for upload.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> pulumi.Input[str]:
        """
        The name of the object.
        """
        return pulumi.get(self, "object_name")

    @object_name.setter
    def object_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "object_name", value)

    @property
    @pulumi.getter(name="accountAcls")
    def account_acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]]]:
        """
        The user set of grant full control.
        """
        return pulumi.get(self, "account_acls")

    @account_acls.setter
    def account_acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]]]):
        pulumi.set(self, "account_acls", value)

    @property
    @pulumi.getter(name="contentMd5")
    def content_md5(self) -> Optional[pulumi.Input[str]]:
        """
        The file md5 sum (32-bit hexadecimal string) for upload.
        """
        return pulumi.get(self, "content_md5")

    @content_md5.setter
    def content_md5(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_md5", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The content type of the object.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        The encryption of the object.Valid value is AES256.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter(name="publicAcl")
    def public_acl(self) -> Optional[pulumi.Input[str]]:
        """
        The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        """
        return pulumi.get(self, "public_acl")

    @public_acl.setter
    def public_acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_acl", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        The storage type of the object.Valid value is STANDARD|IA.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_class", value)


@pulumi.input_type
class _BucketObjectState:
    def __init__(__self__, *,
                 account_acls: Optional[pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]]] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 content_md5: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 enable_version: Optional[pulumi.Input[bool]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 object_name: Optional[pulumi.Input[str]] = None,
                 public_acl: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 version_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering BucketObject resources.
        :param pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]] account_acls: The user set of grant full control.
        :param pulumi.Input[str] bucket_name: The name of the bucket.
        :param pulumi.Input[str] content_md5: The file md5 sum (32-bit hexadecimal string) for upload.
        :param pulumi.Input[str] content_type: The content type of the object.
        :param pulumi.Input[bool] enable_version: The flag of enable tos version.
        :param pulumi.Input[str] encryption: The encryption of the object.Valid value is AES256.
        :param pulumi.Input[str] file_path: The file path for upload.
        :param pulumi.Input[str] object_name: The name of the object.
        :param pulumi.Input[str] public_acl: The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        :param pulumi.Input[str] storage_class: The storage type of the object.Valid value is STANDARD|IA.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] version_ids: The version ids of the object if exist.
        """
        if account_acls is not None:
            pulumi.set(__self__, "account_acls", account_acls)
        if bucket_name is not None:
            pulumi.set(__self__, "bucket_name", bucket_name)
        if content_md5 is not None:
            pulumi.set(__self__, "content_md5", content_md5)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if enable_version is not None:
            pulumi.set(__self__, "enable_version", enable_version)
        if encryption is not None:
            pulumi.set(__self__, "encryption", encryption)
        if file_path is not None:
            pulumi.set(__self__, "file_path", file_path)
        if object_name is not None:
            pulumi.set(__self__, "object_name", object_name)
        if public_acl is not None:
            pulumi.set(__self__, "public_acl", public_acl)
        if storage_class is not None:
            pulumi.set(__self__, "storage_class", storage_class)
        if version_ids is not None:
            pulumi.set(__self__, "version_ids", version_ids)

    @property
    @pulumi.getter(name="accountAcls")
    def account_acls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]]]:
        """
        The user set of grant full control.
        """
        return pulumi.get(self, "account_acls")

    @account_acls.setter
    def account_acls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BucketObjectAccountAclArgs']]]]):
        pulumi.set(self, "account_acls", value)

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket_name")

    @bucket_name.setter
    def bucket_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bucket_name", value)

    @property
    @pulumi.getter(name="contentMd5")
    def content_md5(self) -> Optional[pulumi.Input[str]]:
        """
        The file md5 sum (32-bit hexadecimal string) for upload.
        """
        return pulumi.get(self, "content_md5")

    @content_md5.setter
    def content_md5(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_md5", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The content type of the object.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter(name="enableVersion")
    def enable_version(self) -> Optional[pulumi.Input[bool]]:
        """
        The flag of enable tos version.
        """
        return pulumi.get(self, "enable_version")

    @enable_version.setter
    def enable_version(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_version", value)

    @property
    @pulumi.getter
    def encryption(self) -> Optional[pulumi.Input[str]]:
        """
        The encryption of the object.Valid value is AES256.
        """
        return pulumi.get(self, "encryption")

    @encryption.setter
    def encryption(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption", value)

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> Optional[pulumi.Input[str]]:
        """
        The file path for upload.
        """
        return pulumi.get(self, "file_path")

    @file_path.setter
    def file_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_path", value)

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the object.
        """
        return pulumi.get(self, "object_name")

    @object_name.setter
    def object_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_name", value)

    @property
    @pulumi.getter(name="publicAcl")
    def public_acl(self) -> Optional[pulumi.Input[str]]:
        """
        The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        """
        return pulumi.get(self, "public_acl")

    @public_acl.setter
    def public_acl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_acl", value)

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> Optional[pulumi.Input[str]]:
        """
        The storage type of the object.Valid value is STANDARD|IA.
        """
        return pulumi.get(self, "storage_class")

    @storage_class.setter
    def storage_class(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_class", value)

    @property
    @pulumi.getter(name="versionIds")
    def version_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The version ids of the object if exist.
        """
        return pulumi.get(self, "version_ids")

    @version_ids.setter
    def version_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "version_ids", value)


class BucketObject(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketObjectAccountAclArgs']]]]] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 content_md5: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 object_name: Optional[pulumi.Input[str]] = None,
                 public_acl: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage tos object
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        default = volcengine.tos.BucketObject("default",
            account_acls=[
                volcengine.tos.BucketObjectAccountAclArgs(
                    account_id="1",
                    permission="READ",
                ),
                volcengine.tos.BucketObjectAccountAclArgs(
                    account_id="2001",
                    permission="WRITE_ACP",
                ),
            ],
            bucket_name="test-xym-1",
            encryption="AES256",
            file_path="/Users/bytedance/Work/Go/build/test.txt",
            object_name="demo_xym",
            public_acl="private")
        ```

        ## Import

        TOS Object can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:tos/bucketObject:BucketObject default bucketName:objectName
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketObjectAccountAclArgs']]]] account_acls: The user set of grant full control.
        :param pulumi.Input[str] bucket_name: The name of the bucket.
        :param pulumi.Input[str] content_md5: The file md5 sum (32-bit hexadecimal string) for upload.
        :param pulumi.Input[str] content_type: The content type of the object.
        :param pulumi.Input[str] encryption: The encryption of the object.Valid value is AES256.
        :param pulumi.Input[str] file_path: The file path for upload.
        :param pulumi.Input[str] object_name: The name of the object.
        :param pulumi.Input[str] public_acl: The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        :param pulumi.Input[str] storage_class: The storage type of the object.Valid value is STANDARD|IA.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketObjectArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage tos object
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        default = volcengine.tos.BucketObject("default",
            account_acls=[
                volcengine.tos.BucketObjectAccountAclArgs(
                    account_id="1",
                    permission="READ",
                ),
                volcengine.tos.BucketObjectAccountAclArgs(
                    account_id="2001",
                    permission="WRITE_ACP",
                ),
            ],
            bucket_name="test-xym-1",
            encryption="AES256",
            file_path="/Users/bytedance/Work/Go/build/test.txt",
            object_name="demo_xym",
            public_acl="private")
        ```

        ## Import

        TOS Object can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:tos/bucketObject:BucketObject default bucketName:objectName
        ```

        :param str resource_name: The name of the resource.
        :param BucketObjectArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketObjectArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketObjectAccountAclArgs']]]]] = None,
                 bucket_name: Optional[pulumi.Input[str]] = None,
                 content_md5: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 encryption: Optional[pulumi.Input[str]] = None,
                 file_path: Optional[pulumi.Input[str]] = None,
                 object_name: Optional[pulumi.Input[str]] = None,
                 public_acl: Optional[pulumi.Input[str]] = None,
                 storage_class: Optional[pulumi.Input[str]] = None,
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
            __props__ = BucketObjectArgs.__new__(BucketObjectArgs)

            __props__.__dict__["account_acls"] = account_acls
            if bucket_name is None and not opts.urn:
                raise TypeError("Missing required property 'bucket_name'")
            __props__.__dict__["bucket_name"] = bucket_name
            __props__.__dict__["content_md5"] = content_md5
            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["encryption"] = encryption
            if file_path is None and not opts.urn:
                raise TypeError("Missing required property 'file_path'")
            __props__.__dict__["file_path"] = file_path
            if object_name is None and not opts.urn:
                raise TypeError("Missing required property 'object_name'")
            __props__.__dict__["object_name"] = object_name
            __props__.__dict__["public_acl"] = public_acl
            __props__.__dict__["storage_class"] = storage_class
            __props__.__dict__["enable_version"] = None
            __props__.__dict__["version_ids"] = None
        super(BucketObject, __self__).__init__(
            'volcengine:tos/bucketObject:BucketObject',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_acls: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketObjectAccountAclArgs']]]]] = None,
            bucket_name: Optional[pulumi.Input[str]] = None,
            content_md5: Optional[pulumi.Input[str]] = None,
            content_type: Optional[pulumi.Input[str]] = None,
            enable_version: Optional[pulumi.Input[bool]] = None,
            encryption: Optional[pulumi.Input[str]] = None,
            file_path: Optional[pulumi.Input[str]] = None,
            object_name: Optional[pulumi.Input[str]] = None,
            public_acl: Optional[pulumi.Input[str]] = None,
            storage_class: Optional[pulumi.Input[str]] = None,
            version_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'BucketObject':
        """
        Get an existing BucketObject resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BucketObjectAccountAclArgs']]]] account_acls: The user set of grant full control.
        :param pulumi.Input[str] bucket_name: The name of the bucket.
        :param pulumi.Input[str] content_md5: The file md5 sum (32-bit hexadecimal string) for upload.
        :param pulumi.Input[str] content_type: The content type of the object.
        :param pulumi.Input[bool] enable_version: The flag of enable tos version.
        :param pulumi.Input[str] encryption: The encryption of the object.Valid value is AES256.
        :param pulumi.Input[str] file_path: The file path for upload.
        :param pulumi.Input[str] object_name: The name of the object.
        :param pulumi.Input[str] public_acl: The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        :param pulumi.Input[str] storage_class: The storage type of the object.Valid value is STANDARD|IA.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] version_ids: The version ids of the object if exist.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BucketObjectState.__new__(_BucketObjectState)

        __props__.__dict__["account_acls"] = account_acls
        __props__.__dict__["bucket_name"] = bucket_name
        __props__.__dict__["content_md5"] = content_md5
        __props__.__dict__["content_type"] = content_type
        __props__.__dict__["enable_version"] = enable_version
        __props__.__dict__["encryption"] = encryption
        __props__.__dict__["file_path"] = file_path
        __props__.__dict__["object_name"] = object_name
        __props__.__dict__["public_acl"] = public_acl
        __props__.__dict__["storage_class"] = storage_class
        __props__.__dict__["version_ids"] = version_ids
        return BucketObject(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountAcls")
    def account_acls(self) -> pulumi.Output[Optional[Sequence['outputs.BucketObjectAccountAcl']]]:
        """
        The user set of grant full control.
        """
        return pulumi.get(self, "account_acls")

    @property
    @pulumi.getter(name="bucketName")
    def bucket_name(self) -> pulumi.Output[str]:
        """
        The name of the bucket.
        """
        return pulumi.get(self, "bucket_name")

    @property
    @pulumi.getter(name="contentMd5")
    def content_md5(self) -> pulumi.Output[Optional[str]]:
        """
        The file md5 sum (32-bit hexadecimal string) for upload.
        """
        return pulumi.get(self, "content_md5")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[str]:
        """
        The content type of the object.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter(name="enableVersion")
    def enable_version(self) -> pulumi.Output[bool]:
        """
        The flag of enable tos version.
        """
        return pulumi.get(self, "enable_version")

    @property
    @pulumi.getter
    def encryption(self) -> pulumi.Output[Optional[str]]:
        """
        The encryption of the object.Valid value is AES256.
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter(name="filePath")
    def file_path(self) -> pulumi.Output[str]:
        """
        The file path for upload.
        """
        return pulumi.get(self, "file_path")

    @property
    @pulumi.getter(name="objectName")
    def object_name(self) -> pulumi.Output[str]:
        """
        The name of the object.
        """
        return pulumi.get(self, "object_name")

    @property
    @pulumi.getter(name="publicAcl")
    def public_acl(self) -> pulumi.Output[Optional[str]]:
        """
        The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.
        """
        return pulumi.get(self, "public_acl")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> pulumi.Output[Optional[str]]:
        """
        The storage type of the object.Valid value is STANDARD|IA.
        """
        return pulumi.get(self, "storage_class")

    @property
    @pulumi.getter(name="versionIds")
    def version_ids(self) -> pulumi.Output[Sequence[str]]:
        """
        The version ids of the object if exist.
        """
        return pulumi.get(self, "version_ids")

