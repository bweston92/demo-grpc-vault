# Demo

This demo is me learning, it could be completly wrong.

# Setting up Vault

## Ensure the Vault Server has started.

```
$ docker-compose up -d vault_server
```

## Enable the PKI secrets endpoint.

In order to have Vault issue out keys we need to enable the secret.

Read more about the `secrets enable` command here: https://www.vaultproject.io/docs/commands/secrets/enable.html

```
$ docker-compose run --rm vault_server vault secrets enable pki
```

## Set maximum time to live the for any PKI secret.

```
$ docker-compose run --rm vault_server vault secrets tune -max-lease-ttl=8760h pki
```

## Create root certificate.

```
$ docker-compose run --rm vault_server vault write pki/root/generate/internal common_name=apps.internal ttl=8760h
```

## Create endpoints for issuing certificates and revoking.

```
$ docker-compose run --rm vault_server vault write pki/config/urls issuing_certificates="http://vault_server:8200/v1/pki/ca" crl_distribution_points="http://vault_server:8200/v1/pki/crl"
```

## Create new role.

```
$ docker-compose run --rm vault_server vault write -address=http://docker.legalweb:8200 pki/roles/apps allowed_domains=apps.internal allow_subdomains=true max_ttl=8760h
```

## Start the backend server.

