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

__all__ = ['AccountArgs', 'Account']

@pulumi.input_type
class AccountArgs:
    def __init__(__self__, *,
                 account_name: pulumi.Input[str],
                 account_password: pulumi.Input[str],
                 account_type: pulumi.Input[str],
                 instance_id: pulumi.Input[str],
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]]] = None):
        """
        The set of arguments for constructing a Account resource.
        :param pulumi.Input[str] account_name: Database account name. The rules are as follows:
               Unique name.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, or underscores (_).
               The length is 2~32 characters.
               The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        :param pulumi.Input[str] account_password: The password of the database account.
               Illustrate:
               Cannot start with `!` or `@`.
               The length is 8~32 characters.
               It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
               The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[str] account_type: Database account type, value:
               Super: A high-privilege account. Only one database account can be created for an instance.
               Normal: An account with ordinary privileges.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        :param pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]] account_privileges: The privilege information of account.
        """
        pulumi.set(__self__, "account_name", account_name)
        pulumi.set(__self__, "account_password", account_password)
        pulumi.set(__self__, "account_type", account_type)
        pulumi.set(__self__, "instance_id", instance_id)
        if account_privileges is not None:
            pulumi.set(__self__, "account_privileges", account_privileges)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Input[str]:
        """
        Database account name. The rules are as follows:
        Unique name.
        Start with a letter and end with a letter or number.
        Consists of lowercase letters, numbers, or underscores (_).
        The length is 2~32 characters.
        The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Input[str]:
        """
        The password of the database account.
        Illustrate:
        Cannot start with `!` or `@`.
        The length is 8~32 characters.
        It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
        The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> pulumi.Input[str]:
        """
        Database account type, value:
        Super: A high-privilege account. Only one database account can be created for an instance.
        Normal: An account with ordinary privileges.
        """
        return pulumi.get(self, "account_type")

    @account_type.setter
    def account_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Input[str]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "instance_id", value)

    @property
    @pulumi.getter(name="accountPrivileges")
    def account_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]]]:
        """
        The privilege information of account.
        """
        return pulumi.get(self, "account_privileges")

    @account_privileges.setter
    def account_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]]]):
        pulumi.set(self, "account_privileges", value)


@pulumi.input_type
class _AccountState:
    def __init__(__self__, *,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Account resources.
        :param pulumi.Input[str] account_name: Database account name. The rules are as follows:
               Unique name.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, or underscores (_).
               The length is 2~32 characters.
               The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        :param pulumi.Input[str] account_password: The password of the database account.
               Illustrate:
               Cannot start with `!` or `@`.
               The length is 8~32 characters.
               It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
               The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]] account_privileges: The privilege information of account.
        :param pulumi.Input[str] account_type: Database account type, value:
               Super: A high-privilege account. Only one database account can be created for an instance.
               Normal: An account with ordinary privileges.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        """
        if account_name is not None:
            pulumi.set(__self__, "account_name", account_name)
        if account_password is not None:
            pulumi.set(__self__, "account_password", account_password)
        if account_privileges is not None:
            pulumi.set(__self__, "account_privileges", account_privileges)
        if account_type is not None:
            pulumi.set(__self__, "account_type", account_type)
        if instance_id is not None:
            pulumi.set(__self__, "instance_id", instance_id)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> Optional[pulumi.Input[str]]:
        """
        Database account name. The rules are as follows:
        Unique name.
        Start with a letter and end with a letter or number.
        Consists of lowercase letters, numbers, or underscores (_).
        The length is 2~32 characters.
        The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        """
        return pulumi.get(self, "account_name")

    @account_name.setter
    def account_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_name", value)

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> Optional[pulumi.Input[str]]:
        """
        The password of the database account.
        Illustrate:
        Cannot start with `!` or `@`.
        The length is 8~32 characters.
        It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
        The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "account_password")

    @account_password.setter
    def account_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_password", value)

    @property
    @pulumi.getter(name="accountPrivileges")
    def account_privileges(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]]]:
        """
        The privilege information of account.
        """
        return pulumi.get(self, "account_privileges")

    @account_privileges.setter
    def account_privileges(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AccountAccountPrivilegeArgs']]]]):
        pulumi.set(self, "account_privileges", value)

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> Optional[pulumi.Input[str]]:
        """
        Database account type, value:
        Super: A high-privilege account. Only one database account can be created for an instance.
        Normal: An account with ordinary privileges.
        """
        return pulumi.get(self, "account_type")

    @account_type.setter
    def account_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "account_type", value)

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

    @instance_id.setter
    def instance_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "instance_id", value)


class Account(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountAccountPrivilegeArgs']]]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage rds mysql account
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        default = volcengine.rds_mysql.Account("default",
            account_name="test",
            account_password="xdjsuiahHUH@",
            account_type="Normal",
            instance_id="mysql-e9293705eed6")
        ```

        ## Import

        RDS mysql account can be imported using the instance_id:account_name, e.g.

        ```sh
         $ pulumi import volcengine:rds_mysql/account:Account default mysql-42b38c769c4b:test
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Database account name. The rules are as follows:
               Unique name.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, or underscores (_).
               The length is 2~32 characters.
               The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        :param pulumi.Input[str] account_password: The password of the database account.
               Illustrate:
               Cannot start with `!` or `@`.
               The length is 8~32 characters.
               It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
               The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountAccountPrivilegeArgs']]]] account_privileges: The privilege information of account.
        :param pulumi.Input[str] account_type: Database account type, value:
               Super: A high-privilege account. Only one database account can be created for an instance.
               Normal: An account with ordinary privileges.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage rds mysql account
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        default = volcengine.rds_mysql.Account("default",
            account_name="test",
            account_password="xdjsuiahHUH@",
            account_type="Normal",
            instance_id="mysql-e9293705eed6")
        ```

        ## Import

        RDS mysql account can be imported using the instance_id:account_name, e.g.

        ```sh
         $ pulumi import volcengine:rds_mysql/account:Account default mysql-42b38c769c4b:test
        ```

        :param str resource_name: The name of the resource.
        :param AccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_name: Optional[pulumi.Input[str]] = None,
                 account_password: Optional[pulumi.Input[str]] = None,
                 account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountAccountPrivilegeArgs']]]]] = None,
                 account_type: Optional[pulumi.Input[str]] = None,
                 instance_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AccountArgs.__new__(AccountArgs)

            if account_name is None and not opts.urn:
                raise TypeError("Missing required property 'account_name'")
            __props__.__dict__["account_name"] = account_name
            if account_password is None and not opts.urn:
                raise TypeError("Missing required property 'account_password'")
            __props__.__dict__["account_password"] = None if account_password is None else pulumi.Output.secret(account_password)
            __props__.__dict__["account_privileges"] = account_privileges
            if account_type is None and not opts.urn:
                raise TypeError("Missing required property 'account_type'")
            __props__.__dict__["account_type"] = account_type
            if instance_id is None and not opts.urn:
                raise TypeError("Missing required property 'instance_id'")
            __props__.__dict__["instance_id"] = instance_id
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accountPassword"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Account, __self__).__init__(
            'volcengine:rds_mysql/account:Account',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            account_name: Optional[pulumi.Input[str]] = None,
            account_password: Optional[pulumi.Input[str]] = None,
            account_privileges: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountAccountPrivilegeArgs']]]]] = None,
            account_type: Optional[pulumi.Input[str]] = None,
            instance_id: Optional[pulumi.Input[str]] = None) -> 'Account':
        """
        Get an existing Account resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] account_name: Database account name. The rules are as follows:
               Unique name.
               Start with a letter and end with a letter or number.
               Consists of lowercase letters, numbers, or underscores (_).
               The length is 2~32 characters.
               The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        :param pulumi.Input[str] account_password: The password of the database account.
               Illustrate:
               Cannot start with `!` or `@`.
               The length is 8~32 characters.
               It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
               The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['AccountAccountPrivilegeArgs']]]] account_privileges: The privilege information of account.
        :param pulumi.Input[str] account_type: Database account type, value:
               Super: A high-privilege account. Only one database account can be created for an instance.
               Normal: An account with ordinary privileges.
        :param pulumi.Input[str] instance_id: The ID of the RDS instance.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AccountState.__new__(_AccountState)

        __props__.__dict__["account_name"] = account_name
        __props__.__dict__["account_password"] = account_password
        __props__.__dict__["account_privileges"] = account_privileges
        __props__.__dict__["account_type"] = account_type
        __props__.__dict__["instance_id"] = instance_id
        return Account(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountName")
    def account_name(self) -> pulumi.Output[str]:
        """
        Database account name. The rules are as follows:
        Unique name.
        Start with a letter and end with a letter or number.
        Consists of lowercase letters, numbers, or underscores (_).
        The length is 2~32 characters.
        The [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.
        """
        return pulumi.get(self, "account_name")

    @property
    @pulumi.getter(name="accountPassword")
    def account_password(self) -> pulumi.Output[str]:
        """
        The password of the database account.
        Illustrate:
        Cannot start with `!` or `@`.
        The length is 8~32 characters.
        It consists of any three of uppercase letters, lowercase letters, numbers, and special characters.
        The special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        """
        return pulumi.get(self, "account_password")

    @property
    @pulumi.getter(name="accountPrivileges")
    def account_privileges(self) -> pulumi.Output[Optional[Sequence['outputs.AccountAccountPrivilege']]]:
        """
        The privilege information of account.
        """
        return pulumi.get(self, "account_privileges")

    @property
    @pulumi.getter(name="accountType")
    def account_type(self) -> pulumi.Output[str]:
        """
        Database account type, value:
        Super: A high-privilege account. Only one database account can be created for an instance.
        Normal: An account with ordinary privileges.
        """
        return pulumi.get(self, "account_type")

    @property
    @pulumi.getter(name="instanceId")
    def instance_id(self) -> pulumi.Output[str]:
        """
        The ID of the RDS instance.
        """
        return pulumi.get(self, "instance_id")

