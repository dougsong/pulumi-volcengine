// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vpc

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage acl
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-volcengine/sdk/go/volcengine/Vpc"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := Vpc.NewAcl(ctx, "foo", &Vpc.AclArgs{
// 			AclEntries: vpc.AclAclEntryArray{
// 				&vpc.AclAclEntryArgs{
// 					Description: pulumi.String("e1"),
// 					Entry:       pulumi.String("172.20.1.0/24"),
// 				},
// 				&vpc.AclAclEntryArgs{
// 					Description: pulumi.String("e3"),
// 					Entry:       pulumi.String("172.20.3.0/24"),
// 				},
// 			},
// 			AclName: pulumi.String("tf-test-2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Acl can be imported using the id, e.g.
//
// ```sh
//  $ pulumi import volcengine:Vpc/acl:Acl default acl-mizl7m1kqccg5smt1bdpijuj
// ```
type Acl struct {
	pulumi.CustomResourceState

	// The acl entry set of the Acl.
	AclEntries AclAclEntryArrayOutput `pulumi:"aclEntries"`
	// The name of Acl.
	AclName pulumi.StringOutput `pulumi:"aclName"`
	// Create time of Acl.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the Acl.
	Description pulumi.StringPtrOutput `pulumi:"description"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		args = &AclArgs{}
	}

	var resource Acl
	err := ctx.RegisterResource("volcengine:Vpc/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("volcengine:Vpc/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// The acl entry set of the Acl.
	AclEntries []AclAclEntry `pulumi:"aclEntries"`
	// The name of Acl.
	AclName *string `pulumi:"aclName"`
	// Create time of Acl.
	CreateTime *string `pulumi:"createTime"`
	// The description of the Acl.
	Description *string `pulumi:"description"`
}

type AclState struct {
	// The acl entry set of the Acl.
	AclEntries AclAclEntryArrayInput
	// The name of Acl.
	AclName pulumi.StringPtrInput
	// Create time of Acl.
	CreateTime pulumi.StringPtrInput
	// The description of the Acl.
	Description pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// The acl entry set of the Acl.
	AclEntries []AclAclEntry `pulumi:"aclEntries"`
	// The name of Acl.
	AclName *string `pulumi:"aclName"`
	// The description of the Acl.
	Description *string `pulumi:"description"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// The acl entry set of the Acl.
	AclEntries AclAclEntryArrayInput
	// The name of Acl.
	AclName pulumi.StringPtrInput
	// The description of the Acl.
	Description pulumi.StringPtrInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

// AclArrayInput is an input type that accepts AclArray and AclArrayOutput values.
// You can construct a concrete instance of `AclArrayInput` via:
//
//          AclArray{ AclArgs{...} }
type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

// AclMapInput is an input type that accepts AclMap and AclMapOutput values.
// You can construct a concrete instance of `AclMapInput` via:
//
//          AclMap{ "key": AclArgs{...} }
type AclMapInput interface {
	pulumi.Input

	ToAclMapOutput() AclMapOutput
	ToAclMapOutputWithContext(context.Context) AclMapOutput
}

type AclMap map[string]AclInput

func (AclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (i AclMap) ToAclMapOutput() AclMapOutput {
	return i.ToAclMapOutputWithContext(context.Background())
}

func (i AclMap) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclMapOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil)).Elem()
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

// The acl entry set of the Acl.
func (o AclOutput) AclEntries() AclAclEntryArrayOutput {
	return o.ApplyT(func(v *Acl) AclAclEntryArrayOutput { return v.AclEntries }).(AclAclEntryArrayOutput)
}

// The name of Acl.
func (o AclOutput) AclName() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.AclName }).(pulumi.StringOutput)
}

// Create time of Acl.
func (o AclOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The description of the Acl.
func (o AclOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Acl) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].([]*Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclMapOutput struct{ *pulumi.OutputState }

func (AclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (o AclMapOutput) ToAclMapOutput() AclMapOutput {
	return o
}

func (o AclMapOutput) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return o
}

func (o AclMapOutput) MapIndex(k pulumi.StringInput) AclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Acl {
		return vs[0].(map[string]*Acl)[vs[1].(string)]
	}).(AclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AclInput)(nil)).Elem(), &Acl{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclArrayInput)(nil)).Elem(), AclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AclMapInput)(nil)).Elem(), AclMap{})
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclMapOutput{})
}
