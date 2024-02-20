resource "aws_key_pair" "ssh_key" {
  key_name = "ssh_key"
  public_key = file(var.ssh_public_key_path)
}

resource "aws_instance" "main" {
  ami           = data.aws_ami.amazonlinux.id
  instance_type = "t3.micro"
  subnet_id = aws_subnet.public1.id
  associate_public_ip_address = true
  vpc_security_group_ids = [aws_security_group.main.id]
  key_name = aws_key_pair.ssh_key.key_name

  tags = {
    Name = "${var.project}-ec2-instance"
    project = var.project
  }
}

resource "aws_eip" "ec2-eip" {
  instance = aws_instance.main.id
}