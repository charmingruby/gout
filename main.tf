module "aws_dynamo_data" {
  source = "./modules/aws/data"

  common_tags = local.common_tags
}
