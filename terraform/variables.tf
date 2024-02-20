variable "project" {
  type = string
}

variable "vpc_cidr" {
  type = string
  description = "vpc cidrblock"
}

variable "public_subnet1_cidr" {
  type = string
  description = "public subnet cidrblock"
}

variable "public_subnet2_cidr" {
  type = string
  description = "public subnet2 cidrblock"
}

variable "private_subnet1_cidr" {
  type = string
  description = "private subnet cidrblock"
}

variable "private_subnet2_cidr" {
  type = string
  description = "private subnet2 cidrblock"
}

variable "domain" {
  type = string
  description = "domain name"
}

variable "route53_zone_id" {
  type = string
  description = "route53 zone id"
}

variable "ssh_public_key_path" {
  type = string
  description = "ssh public key path"
}

variable "db_name" {
  description = "Name of the RDS database"
}

variable "db_username" {
  description = "Username for the RDS postgres instance"
}

variable "db_password" {
  description = "Password for the RDS postgres instance"
}