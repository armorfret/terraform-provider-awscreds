terraform-provider-awscreds
=========

[![Build Status](https://img.shields.io/circleci/project/akerl/terraform-provider-awscreds.svg)](https://circleci.com/gh/akerl/terraform-provider-awscreds)
[![GitHub release](https://img.shields.io/github/release/akerl/terraform-provider-awscreds.svg)](https://github.com/akerl/terraform-provider-awscreds/releases)
[![MIT Licensed](https://img.shields.io/badge/license-MIT-green.svg)](https://tldrlegal.com/license/mit-license)

Terraform Provider to generate IAM access keys without storing the secret in the statefile

## Usage

Add the provider to your ~/.terraformrc file:

```
providers {
    awscreds = "/path/to/terraform-provider-awscreds"
}
```

Now you can use the provider in your terraform files to create IAM users:

```
provider "aws" {
    region = "us-east-1"
}

provider "awscreds" {
    region = "us-east-1"
}

resource "aws_iam_user" "my_cool_user" {
    name = "my_cool_user"

resource "awscreds_iam_access_key" {
    user = "${aws_iam_user.my_cool_user.name}"
    file = "./secret_creds"
}
```

The "awscreds" provider takes all the same options as the "aws" provider, and has the same behavior (it checks the same default environment variables and configs for credentials).

The "awscreds_iam_access_key" resource accepts the same options as the "aws_iam_access_key" resource (currently just "user"), and additionally requires the "file" option, which should be the path in which to store the access key and secret key. They are stored newline-delimited in the file.

## Installation

```
git clone git://github.com/akerl/terraform-provider-awscreds
cd terraform-provider-awscreds
go build
# Optionally, move ./terraform-provider-awscreds into your $PATH somewhere, like /usr/local/bin
```

## License

terraform-provider-awscreds is released under the MIT License. See the bundled LICENSE file for details.
