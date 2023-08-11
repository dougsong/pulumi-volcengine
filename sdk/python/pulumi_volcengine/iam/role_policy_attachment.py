# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RolePolicyAttachmentArgs', 'RolePolicyAttachment']

@pulumi.input_type
class RolePolicyAttachmentArgs:
    def __init__(__self__, *,
                 policy_name: pulumi.Input[str],
                 policy_type: pulumi.Input[str],
                 role_name: pulumi.Input[str]):
        """
        The set of arguments for constructing a RolePolicyAttachment resource.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] policy_type: The type of the Policy.
        :param pulumi.Input[str] role_name: The name of the Role.
        """
        pulumi.set(__self__, "policy_name", policy_name)
        pulumi.set(__self__, "policy_type", policy_type)
        pulumi.set(__self__, "role_name", role_name)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Input[str]:
        """
        The name of the Policy.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Input[str]:
        """
        The type of the Policy.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_type", value)

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


@pulumi.input_type
class _RolePolicyAttachmentState:
    def __init__(__self__, *,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RolePolicyAttachment resources.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] policy_type: The type of the Policy.
        :param pulumi.Input[str] role_name: The name of the Role.
        """
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the Policy.
        """
        return pulumi.get(self, "policy_name")

    @policy_name.setter
    def policy_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_name", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the Policy.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

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


class RolePolicyAttachment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage iam role policy attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        role = volcengine.iam.Role("role",
            role_name="TerraformTestRole",
            display_name="terraform role",
            trust_policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"sts:AssumeRole\\"],\\"Principal\\":{\\"Service\\":[\\"auto_scaling\\"]}}]}",
            description="created by terraform",
            max_session_duration=43200)
        policy = volcengine.iam.Policy("policy",
            policy_name="TerraformResourceTest1",
            description="created by terraform 1",
            policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"auto_scaling:DescribeScalingGroups\\"],\\"Resource\\":[\\"*\\"]}]}")
        foo = volcengine.iam.RolePolicyAttachment("foo",
            role_name=role.id,
            policy_name=policy.id,
            policy_type=policy.policy_type)
        ```

        ## Import

        Iam role policy attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:iam/rolePolicyAttachment:RolePolicyAttachment default TerraformTestRole:TerraformTestPolicy:Custom
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] policy_type: The type of the Policy.
        :param pulumi.Input[str] role_name: The name of the Role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RolePolicyAttachmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage iam role policy attachment
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        role = volcengine.iam.Role("role",
            role_name="TerraformTestRole",
            display_name="terraform role",
            trust_policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"sts:AssumeRole\\"],\\"Principal\\":{\\"Service\\":[\\"auto_scaling\\"]}}]}",
            description="created by terraform",
            max_session_duration=43200)
        policy = volcengine.iam.Policy("policy",
            policy_name="TerraformResourceTest1",
            description="created by terraform 1",
            policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"auto_scaling:DescribeScalingGroups\\"],\\"Resource\\":[\\"*\\"]}]}")
        foo = volcengine.iam.RolePolicyAttachment("foo",
            role_name=role.id,
            policy_name=policy.id,
            policy_type=policy.policy_type)
        ```

        ## Import

        Iam role policy attachment can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:iam/rolePolicyAttachment:RolePolicyAttachment default TerraformTestRole:TerraformTestPolicy:Custom
        ```

        :param str resource_name: The name of the resource.
        :param RolePolicyAttachmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RolePolicyAttachmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RolePolicyAttachmentArgs.__new__(RolePolicyAttachmentArgs)

            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
            if policy_type is None and not opts.urn:
                raise TypeError("Missing required property 'policy_type'")
            __props__.__dict__["policy_type"] = policy_type
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
        super(RolePolicyAttachment, __self__).__init__(
            'volcengine:iam/rolePolicyAttachment:RolePolicyAttachment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            policy_name: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None) -> 'RolePolicyAttachment':
        """
        Get an existing RolePolicyAttachment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] policy_type: The type of the Policy.
        :param pulumi.Input[str] role_name: The name of the Role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RolePolicyAttachmentState.__new__(_RolePolicyAttachmentState)

        __props__.__dict__["policy_name"] = policy_name
        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["role_name"] = role_name
        return RolePolicyAttachment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        The name of the Policy.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        """
        The type of the Policy.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        The name of the Role.
        """
        return pulumi.get(self, "role_name")

