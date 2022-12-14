// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gotrue
{
    [GotrueResourceType("gotrue:index/samlIdentityProvider:SamlIdentityProvider")]
    public partial class SamlIdentityProvider : global::Pulumi.CustomResource
    {
        [Output("attributeMapping")]
        public Output<string?> AttributeMapping { get; private set; } = null!;

        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("domains")]
        public Output<ImmutableArray<string>> Domains { get; private set; } = null!;

        [Output("metadataUrl")]
        public Output<string?> MetadataUrl { get; private set; } = null!;

        [Output("metadataXml")]
        public Output<string?> MetadataXml { get; private set; } = null!;

        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;


        /// <summary>
        /// Create a SamlIdentityProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SamlIdentityProvider(string name, SamlIdentityProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("gotrue:index/samlIdentityProvider:SamlIdentityProvider", name, args ?? new SamlIdentityProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SamlIdentityProvider(string name, Input<string> id, SamlIdentityProviderState? state = null, CustomResourceOptions? options = null)
            : base("gotrue:index/samlIdentityProvider:SamlIdentityProvider", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/supabase-community/pulumi-gotrue",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SamlIdentityProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SamlIdentityProvider Get(string name, Input<string> id, SamlIdentityProviderState? state = null, CustomResourceOptions? options = null)
        {
            return new SamlIdentityProvider(name, id, state, options);
        }
    }

    public sealed class SamlIdentityProviderArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributeMapping")]
        public Input<string>? AttributeMapping { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        [Input("metadataUrl")]
        public Input<string>? MetadataUrl { get; set; }

        [Input("metadataXml")]
        public Input<string>? MetadataXml { get; set; }

        public SamlIdentityProviderArgs()
        {
        }
        public static new SamlIdentityProviderArgs Empty => new SamlIdentityProviderArgs();
    }

    public sealed class SamlIdentityProviderState : global::Pulumi.ResourceArgs
    {
        [Input("attributeMapping")]
        public Input<string>? AttributeMapping { get; set; }

        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("domains")]
        private InputList<string>? _domains;
        public InputList<string> Domains
        {
            get => _domains ?? (_domains = new InputList<string>());
            set => _domains = value;
        }

        [Input("metadataUrl")]
        public Input<string>? MetadataUrl { get; set; }

        [Input("metadataXml")]
        public Input<string>? MetadataXml { get; set; }

        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        public SamlIdentityProviderState()
        {
        }
        public static new SamlIdentityProviderState Empty => new SamlIdentityProviderState();
    }
}
