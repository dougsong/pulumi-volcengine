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

__all__ = ['AddressArgs', 'Address']

@pulumi.input_type
class AddressArgs:
    def __init__(__self__, *,
                 billing_type: pulumi.Input[str],
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]]] = None):
        """
        The set of arguments for constructing a Address resource.
        :param pulumi.Input[str] billing_type: The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        :param pulumi.Input[int] bandwidth: The peek bandwidth of the EIP.
        :param pulumi.Input[str] description: The description of the EIP.
        :param pulumi.Input[str] isp: The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        :param pulumi.Input[str] name: The name of the EIP Address.
        :param pulumi.Input[int] period: The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        :param pulumi.Input[str] project_name: The ProjectName of the EIP.
        :param pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]] tags: Tags.
        """
        pulumi.set(__self__, "billing_type", billing_type)
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if isp is not None:
            pulumi.set(__self__, "isp", isp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> pulumi.Input[str]:
        """
        The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The peek bandwidth of the EIP.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the EIP.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def isp(self) -> Optional[pulumi.Input[str]]:
        """
        The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        """
        return pulumi.get(self, "isp")

    @isp.setter
    def isp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the EIP Address.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the EIP.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _AddressState:
    def __init__(__self__, *,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 deleted_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 eip_address: Optional[pulumi.Input[str]] = None,
                 expired_time: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overdue_time: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering Address resources.
        :param pulumi.Input[int] bandwidth: The peek bandwidth of the EIP.
        :param pulumi.Input[str] billing_type: The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        :param pulumi.Input[str] deleted_time: The deleted time of the EIP.
        :param pulumi.Input[str] description: The description of the EIP.
        :param pulumi.Input[str] eip_address: The ip address of the EIP.
        :param pulumi.Input[str] expired_time: The expired time of the EIP.
        :param pulumi.Input[str] isp: The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        :param pulumi.Input[str] name: The name of the EIP Address.
        :param pulumi.Input[str] overdue_time: The overdue time of the EIP.
        :param pulumi.Input[int] period: The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        :param pulumi.Input[str] project_name: The ProjectName of the EIP.
        :param pulumi.Input[str] status: The status of the EIP.
        :param pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]] tags: Tags.
        """
        if bandwidth is not None:
            pulumi.set(__self__, "bandwidth", bandwidth)
        if billing_type is not None:
            pulumi.set(__self__, "billing_type", billing_type)
        if deleted_time is not None:
            pulumi.set(__self__, "deleted_time", deleted_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if eip_address is not None:
            pulumi.set(__self__, "eip_address", eip_address)
        if expired_time is not None:
            pulumi.set(__self__, "expired_time", expired_time)
        if isp is not None:
            pulumi.set(__self__, "isp", isp)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overdue_time is not None:
            pulumi.set(__self__, "overdue_time", overdue_time)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if project_name is not None:
            pulumi.set(__self__, "project_name", project_name)
        if status is not None:
            pulumi.set(__self__, "status", status)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def bandwidth(self) -> Optional[pulumi.Input[int]]:
        """
        The peek bandwidth of the EIP.
        """
        return pulumi.get(self, "bandwidth")

    @bandwidth.setter
    def bandwidth(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bandwidth", value)

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> Optional[pulumi.Input[str]]:
        """
        The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        """
        return pulumi.get(self, "billing_type")

    @billing_type.setter
    def billing_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "billing_type", value)

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> Optional[pulumi.Input[str]]:
        """
        The deleted time of the EIP.
        """
        return pulumi.get(self, "deleted_time")

    @deleted_time.setter
    def deleted_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "deleted_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the EIP.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="eipAddress")
    def eip_address(self) -> Optional[pulumi.Input[str]]:
        """
        The ip address of the EIP.
        """
        return pulumi.get(self, "eip_address")

    @eip_address.setter
    def eip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "eip_address", value)

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> Optional[pulumi.Input[str]]:
        """
        The expired time of the EIP.
        """
        return pulumi.get(self, "expired_time")

    @expired_time.setter
    def expired_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expired_time", value)

    @property
    @pulumi.getter
    def isp(self) -> Optional[pulumi.Input[str]]:
        """
        The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        """
        return pulumi.get(self, "isp")

    @isp.setter
    def isp(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "isp", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the EIP Address.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> Optional[pulumi.Input[str]]:
        """
        The overdue time of the EIP.
        """
        return pulumi.get(self, "overdue_time")

    @overdue_time.setter
    def overdue_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "overdue_time", value)

    @property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[int]]:
        """
        The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "period", value)

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> Optional[pulumi.Input[str]]:
        """
        The ProjectName of the EIP.
        """
        return pulumi.get(self, "project_name")

    @project_name.setter
    def project_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        The status of the EIP.
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AddressTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Address(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressTagArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.eip.Address("foo",
            bandwidth=1,
            billing_type="PostPaidByBandwidth",
            description="acc-test",
            isp="ChinaUnicom",
            project_name="default")
        ```

        ## Import

        Eip address can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:eip/address:Address default eip-274oj9a8rs9a87fap8sf9515b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: The peek bandwidth of the EIP.
        :param pulumi.Input[str] billing_type: The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        :param pulumi.Input[str] description: The description of the EIP.
        :param pulumi.Input[str] isp: The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        :param pulumi.Input[str] name: The name of the EIP Address.
        :param pulumi.Input[int] period: The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        :param pulumi.Input[str] project_name: The ProjectName of the EIP.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressTagArgs']]]] tags: Tags.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AddressArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.eip.Address("foo",
            bandwidth=1,
            billing_type="PostPaidByBandwidth",
            description="acc-test",
            isp="ChinaUnicom",
            project_name="default")
        ```

        ## Import

        Eip address can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:eip/address:Address default eip-274oj9a8rs9a87fap8sf9515b
        ```

        :param str resource_name: The name of the resource.
        :param AddressArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AddressArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 bandwidth: Optional[pulumi.Input[int]] = None,
                 billing_type: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 isp: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[int]] = None,
                 project_name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressTagArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AddressArgs.__new__(AddressArgs)

            __props__.__dict__["bandwidth"] = bandwidth
            if billing_type is None and not opts.urn:
                raise TypeError("Missing required property 'billing_type'")
            __props__.__dict__["billing_type"] = billing_type
            __props__.__dict__["description"] = description
            __props__.__dict__["isp"] = isp
            __props__.__dict__["name"] = name
            __props__.__dict__["period"] = period
            __props__.__dict__["project_name"] = project_name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["deleted_time"] = None
            __props__.__dict__["eip_address"] = None
            __props__.__dict__["expired_time"] = None
            __props__.__dict__["overdue_time"] = None
            __props__.__dict__["status"] = None
        super(Address, __self__).__init__(
            'volcengine:eip/address:Address',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bandwidth: Optional[pulumi.Input[int]] = None,
            billing_type: Optional[pulumi.Input[str]] = None,
            deleted_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            eip_address: Optional[pulumi.Input[str]] = None,
            expired_time: Optional[pulumi.Input[str]] = None,
            isp: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            overdue_time: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[int]] = None,
            project_name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressTagArgs']]]]] = None) -> 'Address':
        """
        Get an existing Address resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] bandwidth: The peek bandwidth of the EIP.
        :param pulumi.Input[str] billing_type: The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        :param pulumi.Input[str] deleted_time: The deleted time of the EIP.
        :param pulumi.Input[str] description: The description of the EIP.
        :param pulumi.Input[str] eip_address: The ip address of the EIP.
        :param pulumi.Input[str] expired_time: The expired time of the EIP.
        :param pulumi.Input[str] isp: The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        :param pulumi.Input[str] name: The name of the EIP Address.
        :param pulumi.Input[str] overdue_time: The overdue time of the EIP.
        :param pulumi.Input[int] period: The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        :param pulumi.Input[str] project_name: The ProjectName of the EIP.
        :param pulumi.Input[str] status: The status of the EIP.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AddressTagArgs']]]] tags: Tags.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AddressState.__new__(_AddressState)

        __props__.__dict__["bandwidth"] = bandwidth
        __props__.__dict__["billing_type"] = billing_type
        __props__.__dict__["deleted_time"] = deleted_time
        __props__.__dict__["description"] = description
        __props__.__dict__["eip_address"] = eip_address
        __props__.__dict__["expired_time"] = expired_time
        __props__.__dict__["isp"] = isp
        __props__.__dict__["name"] = name
        __props__.__dict__["overdue_time"] = overdue_time
        __props__.__dict__["period"] = period
        __props__.__dict__["project_name"] = project_name
        __props__.__dict__["status"] = status
        __props__.__dict__["tags"] = tags
        return Address(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def bandwidth(self) -> pulumi.Output[int]:
        """
        The peek bandwidth of the EIP.
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="billingType")
    def billing_type(self) -> pulumi.Output[str]:
        """
        The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.
        """
        return pulumi.get(self, "billing_type")

    @property
    @pulumi.getter(name="deletedTime")
    def deleted_time(self) -> pulumi.Output[str]:
        """
        The deleted time of the EIP.
        """
        return pulumi.get(self, "deleted_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of the EIP.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="eipAddress")
    def eip_address(self) -> pulumi.Output[str]:
        """
        The ip address of the EIP.
        """
        return pulumi.get(self, "eip_address")

    @property
    @pulumi.getter(name="expiredTime")
    def expired_time(self) -> pulumi.Output[str]:
        """
        The expired time of the EIP.
        """
        return pulumi.get(self, "expired_time")

    @property
    @pulumi.getter
    def isp(self) -> pulumi.Output[str]:
        """
        The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.
        """
        return pulumi.get(self, "isp")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the EIP Address.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="overdueTime")
    def overdue_time(self) -> pulumi.Output[str]:
        """
        The overdue time of the EIP.
        """
        return pulumi.get(self, "overdue_time")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[int]]:
        """
        The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`.This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="projectName")
    def project_name(self) -> pulumi.Output[Optional[str]]:
        """
        The ProjectName of the EIP.
        """
        return pulumi.get(self, "project_name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        The status of the EIP.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.AddressTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

