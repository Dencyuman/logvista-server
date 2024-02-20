terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.7.0"

  backend "s3" {
    bucket = "logvista-tf-state"
    key    = "terraform.tfstate"
    region = "ap-northeast-1"
    encrypt = true
  }
}

provider "aws" {
  region = "ap-northeast-1"
}
