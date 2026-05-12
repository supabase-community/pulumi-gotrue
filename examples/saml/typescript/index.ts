import * as pulumi from "@pulumi/pulumi";

import { SamlIdentityProvider } from "@supabase-community/pulumi-gotrue";

const samltest = new SamlIdentityProvider("samltest", {
    domains: ["samltest.id", "example.com"],
    metadataUrl: "https://samltest.id/saml/idp",
    attributeMapping: JSON.stringify({
        keys: {
            user_name: {
                name: "mail",
            },
        },
    }),
});

