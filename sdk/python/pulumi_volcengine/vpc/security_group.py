# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecurityGroupArgs', 'SecurityGroup']

@pulumi.input_type
class SecurityGroupArgs:
    def __init__(__self__, *,
                 vpc_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 security_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]]] = None):
        """
        The set of arguments for constructing a SecurityGroup resource.
        :param pulumi.Input[str] vpc_id: Id of the VPC.
        :param pulumi.Input[str] description: Description of SecurityGroup.
        :param pulumi.Input[str] project_name: The ProjectName of SecurityGroup.
        :param pulumi.Input[str] security_group_name: Name of SecurityGroup.
        :param pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "vpc_id", vpc_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if security_group_name is not None:
            pulumi.set(__self__, "security_group_name", security_group_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        Id of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of SecurityGroup.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of SecurityGroup.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="securityGroupName")
    def security_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of SecurityGroup.
        """
        return pulumi.get(self, "security_group_name")

    @security_group_name.setter
    def security_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SecurityGroupState:
    def __init__(__self__, *,
                 creation_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 security_group_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecurityGroup resources.
        :param pulumi.Input[str] creation_time: Creation time of SecurityGroup.
        :param pulumi.Input[str] description: Description of SecurityGroup.
        :param pulumi.Input[str] project_name: The ProjectName of SecurityGroup.
        :param pulumi.Input[str] security_group_name: Name of SecurityGroup.
        :param pulumi.Input[str] status: Status of SecurityGroup.
        :param pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]] tags: Tags.
        :param pulumi.Input[str] vpc_id: Id of the VPC.
        """
        if creation_time is not None:
            pulumi.set(__self__, "creation_time", creation_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if security_group_name is not None:
            pulumi.set(__self__, "security_group_name", security_group_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> Optional[pulumi.Input[str]]:
        """
        Creation time of SecurityGroup.
        """
        return pulumi.get(self, "creation_time")

    @creation_time.setter
    def creation_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "creation_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of SecurityGroup.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of SecurityGroup.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter(name="securityGroupName")
    def security_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of SecurityGroup.
        """
        return pulumi.get(self, "security_group_name")

    @security_group_name.setter
    def security_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "security_group_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of SecurityGroup.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecurityGroupTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        Id of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class SecurityGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 security_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage security group
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        g1test1 = volcengine.vpc.SecurityGroup("g1test1",
            project_name="yuwenhao",
            vpc_id="vpc-2feppmy1ugt1c59gp688n1fld")
        ```

        ## Import

        SecurityGroup can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/securityGroup:SecurityGroup default sg-273ycgql3ig3k7fap8t3dyvqx
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description of SecurityGroup.
        :param pulumi.Input[str] project_name: The ProjectName of SecurityGroup.
        :param pulumi.Input[str] security_group_name: Name of SecurityGroup.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] vpc_id: Id of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecurityGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage security group
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        g1test1 = volcengine.vpc.SecurityGroup("g1test1",
            project_name="yuwenhao",
            vpc_id="vpc-2feppmy1ugt1c59gp688n1fld")
        ```

        ## Import

        SecurityGroup can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:vpc/securityGroup:SecurityGroup default sg-273ycgql3ig3k7fap8t3dyvqx
        ```

        :param str resource_name: The name of the resource.
        :param SecurityGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecurityGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 security_group_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecurityGroupArgs.__new__(SecurityGroupArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["security_group_name"] = security_group_name
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["status"] = None
        super(SecurityGroup, __self__).__init__(
            'volcengine:vpc/securityGroup:SecurityGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            creation_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            security_group_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupTagArgs']]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'SecurityGroup':
        """
        Get an existing SecurityGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] creation_time: Creation time of SecurityGroup.
        :param pulumi.Input[str] description: Description of SecurityGroup.
        :param pulumi.Input[str] project_name: The ProjectName of SecurityGroup.
        :param pulumi.Input[str] security_group_name: Name of SecurityGroup.
        :param pulumi.Input[str] status: Status of SecurityGroup.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecurityGroupTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] vpc_id: Id of the VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecurityGroupState.__new__(_SecurityGroupState)

        __props__.__dict__["creation_time"] = creation_time
        __props__.__dict__["description"] = description
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["security_group_name"] = security_group_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        return SecurityGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        Creation time of SecurityGroup.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of SecurityGroup.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[Optional[str]]:
        """
        The ProjectName of SecurityGroup.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter(name="securityGroupName")
    def security_group_name(self) -> pulumi.Output[str]:
        """
        Name of SecurityGroup.
        """
        return pulumi.get(self, "security_group_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of SecurityGroup.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.SecurityGroupTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        Id of the VPC.
        """
        return pulumi.get(self, "vpc_id")

