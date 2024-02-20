resource "aws_eip" "natgw_eip" {
  vpc = true

  tags = {
      Name = "${var.project}-natgw-eip"
  }
}

resource "aws_nat_gateway" "main" {
  allocation_id = aws_eip.natgw_eip.id
  subnet_id     = aws_subnet.public1.id

  tags = {
      Name = "${var.project}-natgw"
  }
}