import { CustomOauthProvider } from "@supabase-community/pulumi-gotrue";

// OAuth2 example: GitHub Enterprise
const githubEnterprise = new CustomOauthProvider("github-enterprise", {
    providerType: "oauth2",
    identifier: "custom:github-enterprise",
    name: "GitHub Enterprise",
    clientId: "your-client-id",
    clientSecret: "your-client-secret",
    authorizationUrl: "https://ghe.example.com/login/oauth/authorize",
    tokenUrl: "https://ghe.example.com/login/oauth/access_token",
    userinfoUrl: "https://ghe.example.com/api/v3/user",
    scopes: ["read:user", "user:email"],
});

// OIDC example: corporate identity provider
const corporateOidc = new CustomOauthProvider("corporate-oidc", {
    providerType: "oidc",
    identifier: "custom:corporate-oidc",
    name: "Corporate SSO",
    clientId: "your-oidc-client-id",
    clientSecret: "your-oidc-client-secret",
    issuer: "https://idp.example.com",
    discoveryUrl: "https://idp.example.com/.well-known/openid-configuration",
    scopes: ["profile", "email"],
});
