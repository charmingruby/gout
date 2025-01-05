terraform {
  required_version = "1.10.3"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "s3" {
    region         = "us-east-1"
    bucket         = "gout-tf"
    key            = "backend/terraform.tfstate"
    dynamodb_table = "tfstate-lock"
  }
}

provider "aws" {
  region = var.aws_region
}
