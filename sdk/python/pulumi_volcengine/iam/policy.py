# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['PolicyArgs', 'Policy']

@pulumi.input_type
class PolicyArgs:
    def __init__(__self__, *,
                 policy_document: pulumi.Input[str],
                 policy_name: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Policy resource.
        :param pulumi.Input[str] policy_document: The document of the Policy.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] description: The description of the Policy.
        """
        pulumi.set(__self__, "policy_document", policy_document)
        pulumi.set(__self__, "policy_name", policy_name)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Input[str]:
        """
        The document of the Policy.
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_document", value)

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
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class _PolicyState:
    def __init__(__self__, *,
                 create_date: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 policy_trn: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 update_date: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Policy resources.
        :param pulumi.Input[str] create_date: The create time of the Policy.
        :param pulumi.Input[str] description: The description of the Policy.
        :param pulumi.Input[str] policy_document: The document of the Policy.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] policy_trn: The resource name of the Policy.
        :param pulumi.Input[str] policy_type: The type of the Policy.
        :param pulumi.Input[str] update_date: The update time of the Policy.
        """
        if create_date is not None:
            pulumi.set(__self__, "create_date", create_date)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policy_document is not None:
            pulumi.set(__self__, "policy_document", policy_document)
        if policy_name is not None:
            pulumi.set(__self__, "policy_name", policy_name)
        if policy_trn is not None:
            pulumi.set(__self__, "policy_trn", policy_trn)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if update_date is not None:
            pulumi.set(__self__, "update_date", update_date)

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the Policy.
        """
        return pulumi.get(self, "create_date")

    @create_date.setter
    def create_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_date", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the Policy.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> Optional[pulumi.Input[str]]:
        """
        The document of the Policy.
        """
        return pulumi.get(self, "policy_document")

    @policy_document.setter
    def policy_document(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_document", value)

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
    @pulumi.getter(name="policyTrn")
    def policy_trn(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the Policy.
        """
        return pulumi.get(self, "policy_trn")

    @policy_trn.setter
    def policy_trn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_trn", value)

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
    @pulumi.getter(name="updateDate")
    def update_date(self) -> Optional[pulumi.Input[str]]:
        """
        The update time of the Policy.
        """
        return pulumi.get(self, "update_date")

    @update_date.setter
    def update_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_date", value)


class Policy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage iam policy
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.iam.Policy("foo",
            description="created by terraform 1",
            policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"auto_scaling:DescribeScalingGroups\\"],\\"Resource\\":[\\"*\\"]}]}",
            policy_name="TerraformResourceTest1")
        ```

        ## Import

        Iam policy can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:iam/policy:Policy default TerraformTestPolicy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: The description of the Policy.
        :param pulumi.Input[str] policy_document: The document of the Policy.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage iam policy
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.iam.Policy("foo",
            description="created by terraform 1",
            policy_document="{\\"Statement\\":[{\\"Effect\\":\\"Allow\\",\\"Action\\":[\\"auto_scaling:DescribeScalingGroups\\"],\\"Resource\\":[\\"*\\"]}]}",
            policy_name="TerraformResourceTest1")
        ```

        ## Import

        Iam policy can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:iam/policy:Policy default TerraformTestPolicy
        ```

        :param str resource_name: The name of the resource.
        :param PolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policy_document: Optional[pulumi.Input[str]] = None,
                 policy_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PolicyArgs.__new__(PolicyArgs)

            __props__.__dict__["description"] = description
            if policy_document is None and not opts.urn:
                raise TypeError("Missing required property 'policy_document'")
            __props__.__dict__["policy_document"] = policy_document
            if policy_name is None and not opts.urn:
                raise TypeError("Missing required property 'policy_name'")
            __props__.__dict__["policy_name"] = policy_name
            __props__.__dict__["create_date"] = None
            __props__.__dict__["policy_trn"] = None
            __props__.__dict__["policy_type"] = None
            __props__.__dict__["update_date"] = None
        super(Policy, __self__).__init__(
            'volcengine:iam/policy:Policy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            create_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            policy_document: Optional[pulumi.Input[str]] = None,
            policy_name: Optional[pulumi.Input[str]] = None,
            policy_trn: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            update_date: Optional[pulumi.Input[str]] = None) -> 'Policy':
        """
        Get an existing Policy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] create_date: The create time of the Policy.
        :param pulumi.Input[str] description: The description of the Policy.
        :param pulumi.Input[str] policy_document: The document of the Policy.
        :param pulumi.Input[str] policy_name: The name of the Policy.
        :param pulumi.Input[str] policy_trn: The resource name of the Policy.
        :param pulumi.Input[str] policy_type: The type of the Policy.
        :param pulumi.Input[str] update_date: The update time of the Policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PolicyState.__new__(_PolicyState)

        __props__.__dict__["create_date"] = create_date
        __props__.__dict__["description"] = description
        __props__.__dict__["policy_document"] = policy_document
        __props__.__dict__["policy_name"] = policy_name
        __props__.__dict__["policy_trn"] = policy_trn
        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["update_date"] = update_date
        return Policy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createDate")
    def create_date(self) -> pulumi.Output[str]:
        """
        The create time of the Policy.
        """
        return pulumi.get(self, "create_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the Policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="policyDocument")
    def policy_document(self) -> pulumi.Output[str]:
        """
        The document of the Policy.
        """
        return pulumi.get(self, "policy_document")

    @property
    @pulumi.getter(name="policyName")
    def policy_name(self) -> pulumi.Output[str]:
        """
        The name of the Policy.
        """
        return pulumi.get(self, "policy_name")

    @property
    @pulumi.getter(name="policyTrn")
    def policy_trn(self) -> pulumi.Output[str]:
        """
        The resource name of the Policy.
        """
        return pulumi.get(self, "policy_trn")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[str]:
        """
        The type of the Policy.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="updateDate")
    def update_date(self) -> pulumi.Output[str]:
        """
        The update time of the Policy.
        """
        return pulumi.get(self, "update_date")

