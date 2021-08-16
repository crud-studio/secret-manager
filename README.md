# secret manager

## Why?

This is an internal tool that is used to create, edit and list specific secrets in the [AWS Secrets Manager](https://aws.amazon.com/secrets-manager/). It will only recognize secrets annotated with the tag `studio.crud.secrets/type=applicationProperties`.

## Installation
1. Ensure AWS CLI is configured with `aws configure`
2. Copy the binary for your platform to a location on PATH (See `/bin` directory for binaries)
3. Run `sm --help` for specific instructions
