resource "aws_db_instance" "main" {
  allocated_storage = 20
  engine = "postgres"
  engine_version = "12.17"
  instance_class = "db.t3.micro"
  db_name = var.db_name
  username = var.db_username
  password = var.db_password
  parameter_group_name = "default.postgres12"
  skip_final_snapshot = true

  vpc_security_group_ids = [aws_security_group.allow_postgres.id]
  db_subnet_group_name = aws_db_subnet_group.main.name
}

resource "aws_db_subnet_group" "main" {
  name = "db-subnet-group"
  subnet_ids = [aws_subnet.private1.id, aws_subnet.private2.id]

  tags = {
      Name = "db-subnet-group"
  }
}