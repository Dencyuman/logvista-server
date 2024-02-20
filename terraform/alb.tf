resource "aws_lb" "alb" {
  name = "${var.project}-alb"
  internal = false
  load_balancer_type = "application"
  security_groups = [aws_security_group.main.id]
  subnets = [aws_subnet.public1.id, aws_subnet.public2.id]

  enable_deletion_protection = false
}

resource "aws_lb_target_group" "alb_tg" {
  name = "${var.project}-alb-tg"
  port = 80
  protocol = "HTTP"
  vpc_id = aws_vpc.main.id
  health_check {
    path = "/"
    protocol = "HTTP"
    port = "traffic-port"
    interval = 30
    timeout = 5
    healthy_threshold = 2
    unhealthy_threshold = 2
  }
}

resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.alb.arn
  port = 443
  protocol = "HTTPS"
  ssl_policy = "ELBSecurityPolicy-2016-08"
  certificate_arn = aws_acm_certificate.main.arn

  default_action {
    type = "forward"
    target_group_arn = aws_lb_target_group.alb_tg.arn
  }
}

resource "aws_lb_target_group_attachment" "alb_tg_attachment" {
  target_group_arn = aws_lb_target_group.alb_tg.arn
  target_id        = aws_instance.main.id
  port             = 80
}