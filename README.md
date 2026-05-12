# Pulumi Provider for GoTrue

**WARNING: This is experimental software and is not indended for production
purposes just yet. Version 2.0.0 and above will be for wider consumption.**

[GoTrue](https://github.com/supabase/gotrue) is the software behind [Supabase
Auth](https://supabase.com/auth).

This package implements a [Pulumi](https://pulumi.com) provider over the
existing
[terraform-provider-gotrue](https://github.com/supabase-community/terraform-provider-gotrue)
that lets you configure some settings in GoTrue as Infrastructure as
Code using Pulumi.

See the `examples` directory for a full example.

## Local Development

This is kind-of complex, so please read the
[pulumi-tf-provider-boilerplate](https://github.com/pulumi/pulumi-tf-provider-boilerplate)
documentation. This provider does very little (if any) customizations on top of
it.

