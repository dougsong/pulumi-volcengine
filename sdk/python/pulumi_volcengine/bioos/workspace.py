# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['WorkspaceArgs', 'Workspace']

@pulumi.input_type
class WorkspaceArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 cover_path: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Workspace resource.
        :param pulumi.Input[str] description: The description of the workspace.
        :param pulumi.Input[str] cover_path: Cover path (relative path in tos bucket).
        :param pulumi.Input[str] name: The name of the workspace.
        """
        pulumi.set(__self__, "description", description)
        if cover_path is not None:
            pulumi.set(__self__, "cover_path", cover_path)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of the workspace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="coverPath")
    def cover_path(self) -> Optional[pulumi.Input[str]]:
        """
        Cover path (relative path in tos bucket).
        """
        return pulumi.get(self, "cover_path")

    @cover_path.setter
    def cover_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cover_path", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _WorkspaceState:
    def __init__(__self__, *,
                 cover_path: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 updated: Optional[pulumi.Input[bool]] = None,
                 workspace_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Workspace resources.
        :param pulumi.Input[str] cover_path: Cover path (relative path in tos bucket).
        :param pulumi.Input[str] description: The description of the workspace.
        :param pulumi.Input[str] name: The name of the workspace.
        :param pulumi.Input[bool] updated: Whether the update complete.
        :param pulumi.Input[str] workspace_id: The id of the workspace.
        """
        if cover_path is not None:
            pulumi.set(__self__, "cover_path", cover_path)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if updated is not None:
            pulumi.set(__self__, "updated", updated)
        if workspace_id is not None:
            pulumi.set(__self__, "workspace_id", workspace_id)

    @property
    @pulumi.getter(name="coverPath")
    def cover_path(self) -> Optional[pulumi.Input[str]]:
        """
        Cover path (relative path in tos bucket).
        """
        return pulumi.get(self, "cover_path")

    @cover_path.setter
    def cover_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cover_path", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the workspace.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def updated(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the update complete.
        """
        return pulumi.get(self, "updated")

    @updated.setter
    def updated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "updated", value)

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the workspace.
        """
        return pulumi.get(self, "workspace_id")

    @workspace_id.setter
    def workspace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "workspace_id", value)


class Workspace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cover_path: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage bioos workspace
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.bioos.Workspace("foo",
            cover_path="template-cover/pic5.png",
            description="test-description23")
        #必填
        ```

        ## Import

        Workspace can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:bioos/workspace:Workspace default *****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cover_path: Cover path (relative path in tos bucket).
        :param pulumi.Input[str] description: The description of the workspace.
        :param pulumi.Input[str] name: The name of the workspace.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: WorkspaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage bioos workspace
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.bioos.Workspace("foo",
            cover_path="template-cover/pic5.png",
            description="test-description23")
        #必填
        ```

        ## Import

        Workspace can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:bioos/workspace:Workspace default *****
        ```

        :param str resource_name: The name of the resource.
        :param WorkspaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(WorkspaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cover_path: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
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
            __props__ = WorkspaceArgs.__new__(WorkspaceArgs)

            __props__.__dict__["cover_path"] = cover_path
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["updated"] = None
            __props__.__dict__["workspace_id"] = None
        super(Workspace, __self__).__init__(
            'volcengine:bioos/workspace:Workspace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cover_path: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            updated: Optional[pulumi.Input[bool]] = None,
            workspace_id: Optional[pulumi.Input[str]] = None) -> 'Workspace':
        """
        Get an existing Workspace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cover_path: Cover path (relative path in tos bucket).
        :param pulumi.Input[str] description: The description of the workspace.
        :param pulumi.Input[str] name: The name of the workspace.
        :param pulumi.Input[bool] updated: Whether the update complete.
        :param pulumi.Input[str] workspace_id: The id of the workspace.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _WorkspaceState.__new__(_WorkspaceState)

        __props__.__dict__["cover_path"] = cover_path
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["updated"] = updated
        __props__.__dict__["workspace_id"] = workspace_id
        return Workspace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="coverPath")
    def cover_path(self) -> pulumi.Output[Optional[str]]:
        """
        Cover path (relative path in tos bucket).
        """
        return pulumi.get(self, "cover_path")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the workspace.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the workspace.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def updated(self) -> pulumi.Output[bool]:
        """
        Whether the update complete.
        """
        return pulumi.get(self, "updated")

    @property
    @pulumi.getter(name="workspaceId")
    def workspace_id(self) -> pulumi.Output[str]:
        """
        The id of the workspace.
        """
        return pulumi.get(self, "workspace_id")

