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

__all__ = ['GatewayArgs', 'Gateway']

@pulumi.input_type
class GatewayArgs:
    def __init__(__self__, *,
                 subnet_id: pulumi.Input[str],
                 vpc_id: pulumi.Input[str],
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 nat_gateway_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]]] = None):
        """
        The set of arguments for constructing a Gateway resource.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        :param pulumi.Input[str] billing_type: The billing type of the NatGateway, the value is `PostPaid`.
        :param pulumi.Input[str] description: The description of the NatGateway.
        :param pulumi.Input[str] nat_gateway_name: The name of the NatGateway.
        :param pulumi.Input[str] project_name: The ProjectName of the NatGateway.
        :param pulumi.Input[str] spec: The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        :param pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "subnet_id", subnet_id)
        pulumi.set(__self__, "vpc_id", vpc_id)
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if nat_gateway_name is not None:
            pulumi.set(__self__, "nat_gateway_name", nat_gateway_name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Input[str]:
        """
        The ID of the Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Input[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vpc_id", value)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The billing type of the NatGateway, the value is `PostPaid`.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the NatGateway.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="natGatewayName")
    def nat_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NatGateway.
        """
        return pulumi.get(self, "nat_gateway_name")

    @nat_gateway_name.setter
    def nat_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the NatGateway.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _GatewayState:
    def __init__(__self__, *,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 nat_gateway_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Gateway resources.
        :param pulumi.Input[str] billing_type: The billing type of the NatGateway, the value is `PostPaid`.
        :param pulumi.Input[str] description: The description of the NatGateway.
        :param pulumi.Input[str] nat_gateway_name: The name of the NatGateway.
        :param pulumi.Input[str] project_name: The ProjectName of the NatGateway.
        :param pulumi.Input[str] spec: The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet.
        :param pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]] tags: Tags.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if nat_gateway_name is not None:
            pulumi.set(__self__, "nat_gateway_name", nat_gateway_name)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)
        if subnet_id is not None:
            pulumi.set(__self__, "subnet_id", subnet_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vpc_id is not None:
            pulumi.set(__self__, "vpc_id", vpc_id)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The billing type of the NatGateway, the value is `PostPaid`.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the NatGateway.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="natGatewayName")
    def nat_gateway_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the NatGateway.
        """
        return pulumi.get(self, "nat_gateway_name")

    @nat_gateway_name.setter
    def nat_gateway_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "nat_gateway_name", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the NatGateway.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input[str]]:
        """
        The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "spec", value)

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @subnet_id.setter
    def subnet_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subnet_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GatewayTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

    @vpc_id.setter
    def vpc_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vpc_id", value)


class Gateway(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 nat_gateway_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GatewayTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.nat.Gateway("foo",
            description="This nat gateway auto-created by terraform. ",
            nat_gateway_name="tf-auto-demo-1",
            project_name="default",
            spec="Medium",
            subnet_id="subnet-im67x70vxla88gbssz1hy1z2",
            vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        ```

        ## Import

        NatGateway can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:nat/gateway:Gateway default ngw-vv3t043k05sm****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_type: The billing type of the NatGateway, the value is `PostPaid`.
        :param pulumi.Input[str] description: The description of the NatGateway.
        :param pulumi.Input[str] nat_gateway_name: The name of the NatGateway.
        :param pulumi.Input[str] project_name: The ProjectName of the NatGateway.
        :param pulumi.Input[str] spec: The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GatewayTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GatewayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.nat.Gateway("foo",
            description="This nat gateway auto-created by terraform. ",
            nat_gateway_name="tf-auto-demo-1",
            project_name="default",
            spec="Medium",
            subnet_id="subnet-im67x70vxla88gbssz1hy1z2",
            vpc_id="vpc-im67wjcikxkw8gbssx8ufpj8")
        ```

        ## Import

        NatGateway can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:nat/gateway:Gateway default ngw-vv3t043k05sm****
        ```

        :param str resource_name: The name of the resource.
        :param GatewayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GatewayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 nat_gateway_name: Optional[pulumi.Input[str]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[str]] = None,
                 subnet_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GatewayTagArgs']]]]] = None,
                 vpc_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = GatewayArgs.__new__(GatewayArgs)

            __props__.__dict__["billing_type"] = billing_type
            __props__.__dict__["description"] = description
            __props__.__dict__["nat_gateway_name"] = nat_gateway_name
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["spec"] = spec
            if subnet_id is None and not opts.urn:
                raise TypeError("Missing required property 'subnet_id'")
            __props__.__dict__["subnet_id"] = subnet_id
            __props__.__dict__["tags"] = tags
            if vpc_id is None and not opts.urn:
                raise TypeError("Missing required property 'vpc_id'")
            __props__.__dict__["vpc_id"] = vpc_id
        super(Gateway, __self__).__init__(
            'volcengine:nat/gateway:Gateway',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            billing_type: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            nat_gateway_name: Optional[pulumi.Input[str]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            spec: Optional[pulumi.Input[str]] = None,
            subnet_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GatewayTagArgs']]]]] = None,
            vpc_id: Optional[pulumi.Input[str]] = None) -> 'Gateway':
        """
        Get an existing Gateway resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] billing_type: The billing type of the NatGateway, the value is `PostPaid`.
        :param pulumi.Input[str] description: The description of the NatGateway.
        :param pulumi.Input[str] nat_gateway_name: The name of the NatGateway.
        :param pulumi.Input[str] project_name: The ProjectName of the NatGateway.
        :param pulumi.Input[str] spec: The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        :param pulumi.Input[str] subnet_id: The ID of the Subnet.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GatewayTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] vpc_id: The ID of the VPC.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GatewayState.__new__(_GatewayState)

        __props__.__dict__["billing_type"] = billing_type
        __props__.__dict__["description"] = description
        __props__.__dict__["nat_gateway_name"] = nat_gateway_name
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["spec"] = spec
        __props__.__dict__["subnet_id"] = subnet_id
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vpc_id"] = vpc_id
        return Gateway(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> pulumi.Output[Optional[str]]:
        """
        The billing type of the NatGateway, the value is `PostPaid`.
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the NatGateway.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="natGatewayName")
    def nat_gateway_name(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the NatGateway.
        """
        return pulumi.get(self, "nat_gateway_name")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[Optional[str]]:
        """
        The ProjectName of the NatGateway.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output[Optional[str]]:
        """
        The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large`.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter(name="subnetId")
    def subnet_id(self) -> pulumi.Output[str]:
        """
        The ID of the Subnet.
        """
        return pulumi.get(self, "subnet_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.GatewayTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vpcId")
    def vpc_id(self) -> pulumi.Output[str]:
        """
        The ID of the VPC.
        """
        return pulumi.get(self, "vpc_id")

