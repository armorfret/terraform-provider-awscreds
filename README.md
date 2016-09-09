terraform-provider-awscreds
=========

[![Build Status](https://img.shields.io/circleci/project/akerl/terraform-provider-awscreds.svg)](https://circleci.com/gh/akerl/terraform-provider-awscreds)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

Terraform Provider to generate IAM access keys without storing the secret in the statefile

## Usage

Add the provider to your ~/.terraformrc file:

```
providers {
    awscreds = "/path/to/terraform-provider-awscreds"
}
```

## Installation

```
git clone git://github.com/akerl/terraform-provider-awscreds
cd terraform-provider-awscreds
go build
# Optionally, move ./terraform-provider-awscreds into your $PATH somewhere, like /usr/local/bin
```

## License

terraform-provider-awscreds is released under the MIT License. See the bundled LICENSE file for details.
