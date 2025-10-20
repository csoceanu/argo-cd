```yaml
url: https://argocd.example.com
oidc.config: |
  name: Okta
  # this is the authorization server URI
  issuer: https://example.okta.com/oauth2/aus9abcdefgABCDEFGd7
  clientID: 0oa9abcdefgh123AB5d7
  cliClientID: gfedcba0987654321GEFDCBA # Optional if using the CLI for SSO
  clientSecret: ABCDEFG1234567890abcdefg
  requestedScopes: ["openid", "profile", "email", "groups"]
  requestedIDTokenClaims: {"groups": {"essential": true}}
```

You may want to store the `clientSecret` in a Kubernetes secret; see [how to deal with SSO secrets](./index.md/#sensitive-data-and-sso-client-secrets ) for more details.