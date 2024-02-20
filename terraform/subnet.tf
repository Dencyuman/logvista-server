resource "aws_subnet" "public1" {
  vpc_id = aws_vpc.main.id
  cidr_block = var.public_subnet1_cidr
  availability_zone = "ap-northeast-1a"

  tags = {
      Name = "${var.project}-public-subnet1"
  }
}

resource "aws_subnet" "public2" {
  vpc_id = aws_vpc.main.id
  cidr_block = var.public_subnet2_cidr
  availability_zone = "ap-northeast-1c"

  tags = {
      Name = "${var.project}-public-subnet2"
  }
}

resource "aws_subnet" "private1" {
  vpc_id = aws_vpc.main.id
  cidr_block = var.private_subnet1_cidr
  availability_zone = "ap-northeast-1a"

  tags = {
      Name = "${var.project}-private-subnet1"
  }
}

resource "aws_subnet" "private2" {
  vpc_id = aws_vpc.main.id
  cidr_block = var.private_subnet2_cidr
  availability_zone = "ap-northeast-1c"

  tags = {
      Name = "${var.project}-private-subnet2"
  }
}