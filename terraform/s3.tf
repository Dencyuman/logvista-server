resource "aws_s3_bucket" "logvista_tf_state" {
  bucket = "${var.project}-tf-state"
}

resource "aws_s3_bucket_versioning" "logvista_tf_state" {
  bucket = aws_s3_bucket.logvista_tf_state.bucket
  versioning_configuration {
    status = "Enabled"
  }
}
