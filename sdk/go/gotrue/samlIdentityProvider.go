// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gotrue

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SamlIdentityProvider struct {
	pulumi.CustomResourceState

	AttributeMapping pulumi.StringPtrOutput   `pulumi:"attributeMapping"`
	CreatedAt        pulumi.StringOutput      `pulumi:"createdAt"`
	Domains          pulumi.StringArrayOutput `pulumi:"domains"`
	MetadataUrl      pulumi.StringPtrOutput   `pulumi:"metadataUrl"`
	MetadataXml      pulumi.StringPtrOutput   `pulumi:"metadataXml"`
	UpdatedAt        pulumi.StringOutput      `pulumi:"updatedAt"`
}

// NewSamlIdentityProvider registers a new resource with the given unique name, arguments, and options.
func NewSamlIdentityProvider(ctx *pulumi.Context,
	name string, args *SamlIdentityProviderArgs, opts ...pulumi.ResourceOption) (*SamlIdentityProvider, error) {
	if args == nil {
		args = &SamlIdentityProviderArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource SamlIdentityProvider
	err := ctx.RegisterResource("gotrue:index/samlIdentityProvider:SamlIdentityProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSamlIdentityProvider gets an existing SamlIdentityProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSamlIdentityProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SamlIdentityProviderState, opts ...pulumi.ResourceOption) (*SamlIdentityProvider, error) {
	var resource SamlIdentityProvider
	err := ctx.ReadResource("gotrue:index/samlIdentityProvider:SamlIdentityProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SamlIdentityProvider resources.
type samlIdentityProviderState struct {
	AttributeMapping *string  `pulumi:"attributeMapping"`
	CreatedAt        *string  `pulumi:"createdAt"`
	Domains          []string `pulumi:"domains"`
	MetadataUrl      *string  `pulumi:"metadataUrl"`
	MetadataXml      *string  `pulumi:"metadataXml"`
	UpdatedAt        *string  `pulumi:"updatedAt"`
}

type SamlIdentityProviderState struct {
	AttributeMapping pulumi.StringPtrInput
	CreatedAt        pulumi.StringPtrInput
	Domains          pulumi.StringArrayInput
	MetadataUrl      pulumi.StringPtrInput
	MetadataXml      pulumi.StringPtrInput
	UpdatedAt        pulumi.StringPtrInput
}

func (SamlIdentityProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*samlIdentityProviderState)(nil)).Elem()
}

type samlIdentityProviderArgs struct {
	AttributeMapping *string  `pulumi:"attributeMapping"`
	Domains          []string `pulumi:"domains"`
	MetadataUrl      *string  `pulumi:"metadataUrl"`
	MetadataXml      *string  `pulumi:"metadataXml"`
}

// The set of arguments for constructing a SamlIdentityProvider resource.
type SamlIdentityProviderArgs struct {
	AttributeMapping pulumi.StringPtrInput
	Domains          pulumi.StringArrayInput
	MetadataUrl      pulumi.StringPtrInput
	MetadataXml      pulumi.StringPtrInput
}

func (SamlIdentityProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*samlIdentityProviderArgs)(nil)).Elem()
}

type SamlIdentityProviderInput interface {
	pulumi.Input

	ToSamlIdentityProviderOutput() SamlIdentityProviderOutput
	ToSamlIdentityProviderOutputWithContext(ctx context.Context) SamlIdentityProviderOutput
}

func (*SamlIdentityProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlIdentityProvider)(nil)).Elem()
}

func (i *SamlIdentityProvider) ToSamlIdentityProviderOutput() SamlIdentityProviderOutput {
	return i.ToSamlIdentityProviderOutputWithContext(context.Background())
}

func (i *SamlIdentityProvider) ToSamlIdentityProviderOutputWithContext(ctx context.Context) SamlIdentityProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlIdentityProviderOutput)
}

// SamlIdentityProviderArrayInput is an input type that accepts SamlIdentityProviderArray and SamlIdentityProviderArrayOutput values.
// You can construct a concrete instance of `SamlIdentityProviderArrayInput` via:
//
//	SamlIdentityProviderArray{ SamlIdentityProviderArgs{...} }
type SamlIdentityProviderArrayInput interface {
	pulumi.Input

	ToSamlIdentityProviderArrayOutput() SamlIdentityProviderArrayOutput
	ToSamlIdentityProviderArrayOutputWithContext(context.Context) SamlIdentityProviderArrayOutput
}

type SamlIdentityProviderArray []SamlIdentityProviderInput

func (SamlIdentityProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlIdentityProvider)(nil)).Elem()
}

func (i SamlIdentityProviderArray) ToSamlIdentityProviderArrayOutput() SamlIdentityProviderArrayOutput {
	return i.ToSamlIdentityProviderArrayOutputWithContext(context.Background())
}

func (i SamlIdentityProviderArray) ToSamlIdentityProviderArrayOutputWithContext(ctx context.Context) SamlIdentityProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlIdentityProviderArrayOutput)
}

// SamlIdentityProviderMapInput is an input type that accepts SamlIdentityProviderMap and SamlIdentityProviderMapOutput values.
// You can construct a concrete instance of `SamlIdentityProviderMapInput` via:
//
//	SamlIdentityProviderMap{ "key": SamlIdentityProviderArgs{...} }
type SamlIdentityProviderMapInput interface {
	pulumi.Input

	ToSamlIdentityProviderMapOutput() SamlIdentityProviderMapOutput
	ToSamlIdentityProviderMapOutputWithContext(context.Context) SamlIdentityProviderMapOutput
}

type SamlIdentityProviderMap map[string]SamlIdentityProviderInput

func (SamlIdentityProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlIdentityProvider)(nil)).Elem()
}

func (i SamlIdentityProviderMap) ToSamlIdentityProviderMapOutput() SamlIdentityProviderMapOutput {
	return i.ToSamlIdentityProviderMapOutputWithContext(context.Background())
}

func (i SamlIdentityProviderMap) ToSamlIdentityProviderMapOutputWithContext(ctx context.Context) SamlIdentityProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SamlIdentityProviderMapOutput)
}

type SamlIdentityProviderOutput struct{ *pulumi.OutputState }

func (SamlIdentityProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SamlIdentityProvider)(nil)).Elem()
}

func (o SamlIdentityProviderOutput) ToSamlIdentityProviderOutput() SamlIdentityProviderOutput {
	return o
}

func (o SamlIdentityProviderOutput) ToSamlIdentityProviderOutputWithContext(ctx context.Context) SamlIdentityProviderOutput {
	return o
}

func (o SamlIdentityProviderOutput) AttributeMapping() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamlIdentityProvider) pulumi.StringPtrOutput { return v.AttributeMapping }).(pulumi.StringPtrOutput)
}

func (o SamlIdentityProviderOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlIdentityProvider) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SamlIdentityProviderOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SamlIdentityProvider) pulumi.StringArrayOutput { return v.Domains }).(pulumi.StringArrayOutput)
}

func (o SamlIdentityProviderOutput) MetadataUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamlIdentityProvider) pulumi.StringPtrOutput { return v.MetadataUrl }).(pulumi.StringPtrOutput)
}

func (o SamlIdentityProviderOutput) MetadataXml() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SamlIdentityProvider) pulumi.StringPtrOutput { return v.MetadataXml }).(pulumi.StringPtrOutput)
}

func (o SamlIdentityProviderOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *SamlIdentityProvider) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type SamlIdentityProviderArrayOutput struct{ *pulumi.OutputState }

func (SamlIdentityProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SamlIdentityProvider)(nil)).Elem()
}

func (o SamlIdentityProviderArrayOutput) ToSamlIdentityProviderArrayOutput() SamlIdentityProviderArrayOutput {
	return o
}

func (o SamlIdentityProviderArrayOutput) ToSamlIdentityProviderArrayOutputWithContext(ctx context.Context) SamlIdentityProviderArrayOutput {
	return o
}

func (o SamlIdentityProviderArrayOutput) Index(i pulumi.IntInput) SamlIdentityProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SamlIdentityProvider {
		return vs[0].([]*SamlIdentityProvider)[vs[1].(int)]
	}).(SamlIdentityProviderOutput)
}

type SamlIdentityProviderMapOutput struct{ *pulumi.OutputState }

func (SamlIdentityProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SamlIdentityProvider)(nil)).Elem()
}

func (o SamlIdentityProviderMapOutput) ToSamlIdentityProviderMapOutput() SamlIdentityProviderMapOutput {
	return o
}

func (o SamlIdentityProviderMapOutput) ToSamlIdentityProviderMapOutputWithContext(ctx context.Context) SamlIdentityProviderMapOutput {
	return o
}

func (o SamlIdentityProviderMapOutput) MapIndex(k pulumi.StringInput) SamlIdentityProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SamlIdentityProvider {
		return vs[0].(map[string]*SamlIdentityProvider)[vs[1].(string)]
	}).(SamlIdentityProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SamlIdentityProviderInput)(nil)).Elem(), &SamlIdentityProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlIdentityProviderArrayInput)(nil)).Elem(), SamlIdentityProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SamlIdentityProviderMapInput)(nil)).Elem(), SamlIdentityProviderMap{})
	pulumi.RegisterOutputType(SamlIdentityProviderOutput{})
	pulumi.RegisterOutputType(SamlIdentityProviderArrayOutput{})
	pulumi.RegisterOutputType(SamlIdentityProviderMapOutput{})
}
