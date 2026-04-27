# 011 正解例（課題別）

## 課題1: 最小Terraform構成

```hcl
terraform {
  required_version = ">= 1.7.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.region
}

variable "region" {
  type    = string
  default = "ap-northeast-1"
}
```

## 課題2: Go API配置想定（例: ECSの前段としてログバケット）

```hcl
resource "aws_s3_bucket" "app_logs" {
  bucket = "${var.project}-app-logs"
}

variable "project" {
  type    = string
  default = "go-study"
}
```

## 課題3: module化例（呼び出し側）

```hcl
module "dev_app_logs" {
  source  = "./modules/app_logs"
  project = "go-study-dev"
}
```

```hcl
// modules/app_logs/main.tf
variable "project" { type = string }

resource "aws_s3_bucket" "this" {
  bucket = "${var.project}-app-logs"
}
```
