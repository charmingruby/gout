resource "aws_dynamodb_table" "outbox" {
  name         = "outbox"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "AggregateID"
  range_key    = "EventTimestamp"

  attribute {
    type = "S"
    name = "AggregateID" # Ex: ORDER#123, PAYMENT#456  
  }

  attribute {
    type = "S"
    name = "EventTimestamp" # Ex: 2024-01-07T10:00:00#unique-id
  }

  attribute {
    name = "EventStatus" # Ex: PENDING, PROCESSED
    type = "S"
  }
  attribute {
    name = "ProcessedAt" # Ex: 2024-01-07T10:00:00
    type = "S"
  }

  global_secondary_index {
    name            = "StatusProcessingIndex"
    hash_key        = "EventStatus"
    range_key       = "ProcessedAt"
    projection_type = "ALL"
  }

  point_in_time_recovery {
    enabled = true
  }

  tags = merge(var.common_tags, {
    Name = format("outbox-%s-dynamodb-table", var.common_tags.service_name)
  })
}
