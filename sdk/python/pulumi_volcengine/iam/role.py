# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 role_name: pulumi.Input[str],
                 trust_policy_document: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[str] display_name: The display name of the Role.
        :param pulumi.Input[str] role_name: The name of the Role.
        :param pulumi.Input[str] trust_policy_document: The trust policy document of the Role.
        :param pulumi.Input[str] description: The description of the Role.
        :param pulumi.Input[int] max_session_duration: The max session duration of the Role.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "role_name", role_name)
        pulumi.set(__self__, "trust_policy_document", trust_policy_document)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if max_session_duration is not None:
            pulumi.set(__self__, "max_session_duration", max_session_duration)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the Role.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="trustPolicyDocument")
    def trust_policy_document(self) -> pulumi.Input[str]:
        """
        The trust policy document of the Role.
        """
        return pulumi.get(self, "trust_policy_document")

    @trust_policy_document.setter
    def trust_policy_document(self, value: pulumi.Input[str]):
        pulumi.set(self, "trust_policy_document", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[pulumi.Input[int]]:
        """
        The max session duration of the Role.
        """
        return pulumi.get(self, "max_session_duration")

    @max_session_duration.setter
    def max_session_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_session_duration", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 trn: Optional[pulumi.Input[str]] = None,
                 trust_policy_document: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[str] description: The description of the Role.
        :param pulumi.Input[str] display_name: The display name of the Role.
        :param pulumi.Input[int] max_session_duration: The max session duration of the Role.
        :param pulumi.Input[str] role_name: The name of the Role.
        :param pulumi.Input[str] trn: The resource name of the Role.
        :param pulumi.Input[str] trust_policy_document: The trust policy document of the Role.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if max_session_duration is not None:
            pulumi.set(__self__, "max_session_duration", max_session_duration)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if trn is not None:
            pulumi.set(__self__, "trn", trn)
        if trust_policy_document is not None:
            pulumi.set(__self__, "trust_policy_document", trust_policy_document)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Role.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of the Role.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> Optional[pulumi.Input[int]]:
        """
        The max session duration of the Role.
        """
        return pulumi.get(self, "max_session_duration")

    @max_session_duration.setter
    def max_session_duration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_session_duration", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter
    def trn(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the Role.
        """
        return pulumi.get(self, "trn")

    @trn.setter
    def trn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trn", value)

    @property
    @pulumi.getter(name="trustPolicyDocument")
    def trust_policy_document(self) -> Optional[pulumi.Input[str]]:
        """
        The trust policy document of the Role.
        """
        return pulumi.get(self, "trust_policy_document")

    @trust_policy_document.setter
    def trust_policy_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "trust_policy_document", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 trust_policy_document: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage iam role
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.iam.Role("foo",
            description="created by terraform",
            display_name="terraform role",
            max_session_duration=43200,
            role_name="TerraformTestRole",
            trust_policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"sts:AssumeRole\\"],\\"Principal\\":{\\"Service\\":[\\"auto_scaling\\"]}}]}")
        ```

        ## Import

        Iam role can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:iam/role:Role default TerraformTestRole
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Role.
        :param pulumi.Input[str] display_name: The display name of the Role.
        :param pulumi.Input[int] max_session_duration: The max session duration of the Role.
        :param pulumi.Input[str] role_name: The name of the Role.
        :param pulumi.Input[str] trust_policy_document: The trust policy document of the Role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage iam role
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.iam.Role("foo",
            description="created by terraform",
            display_name="terraform role",
            max_session_duration=43200,
            role_name="TerraformTestRole",
            trust_policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"sts:AssumeRole\\"],\\"Principal\\":{\\"Service\\":[\\"auto_scaling\\"]}}]}")
        ```

        ## Import

        Iam role can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:iam/role:Role default TerraformTestRole
        ```

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 max_session_duration: Optional[pulumi.Input[int]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 trust_policy_document: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            __props__.__dict__["description"] = description
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["max_session_duration"] = max_session_duration
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            if trust_policy_document is None and not opts.urn:
                raise TypeError("Missing required property 'trust_policy_document'")
            __props__.__dict__["trust_policy_document"] = trust_policy_document
            __props__.__dict__["trn"] = None
        super(Role, __self__).__init__(
            'volcengine:iam/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            max_session_duration: Optional[pulumi.Input[int]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            trn: Optional[pulumi.Input[str]] = None,
            trust_policy_document: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Role.
        :param pulumi.Input[str] display_name: The display name of the Role.
        :param pulumi.Input[int] max_session_duration: The max session duration of the Role.
        :param pulumi.Input[str] role_name: The name of the Role.
        :param pulumi.Input[str] trn: The resource name of the Role.
        :param pulumi.Input[str] trust_policy_document: The trust policy document of the Role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["description"] = description
        __props__.__dict__["display_name"] = display_name
        __props__.__dict__["max_session_duration"] = max_session_duration
        __props__.__dict__["role_name"] = role_name
        __props__.__dict__["trn"] = trn
        __props__.__dict__["trust_policy_document"] = trust_policy_document
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the Role.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="maxSessionDuration")
    def max_session_duration(self) -> pulumi.Output[Optional[int]]:
        """
        The max session duration of the Role.
        """
        return pulumi.get(self, "max_session_duration")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def trn(self) -> pulumi.Output[str]:
        """
        The resource name of the Role.
        """
        return pulumi.get(self, "trn")

    @property
    @pulumi.getter(name="trustPolicyDocument")
    def trust_policy_document(self) -> pulumi.Output[str]:
        """
        The trust policy document of the Role.
        """
        return pulumi.get(self, "trust_policy_document")

