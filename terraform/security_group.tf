resource "aws_security_group" "main" {
  name = "${var.project}-sg"
  description = "security group for ${var.project}"
  vpc_id = aws_vpc.main.id
  tags = {
      Name = "${var.project}-sg"
      Project = var.project
  }

  ingress {
    from_port = 22
    to_port = 22
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "allow_postgres" {
  name = "${var.project}-allow-pg-sg"
  description = "security group for ${var.project} to allow postgres"
  vpc_id = aws_vpc.main.id
  tags = {
      Name = "${var.project}-allow-postgres"
      Project = var.project
  }

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}