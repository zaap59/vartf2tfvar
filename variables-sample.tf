variable "alb_name" {
  type        = "string"
  description = "ALB Name"
}

variable "alb_sg" {
  type        = "string"
  description = "ID of ALB Security group"
}

variable "alb_target_health_check" {
  type        = "string"
  description = "HTTP page for health check"
  default     = "/"
}

variable "alb_target_port" {
  type        = "string"
  description = "Instance listner"
  default     = "7777"
}

variable "app_name" {
  type        = "string"
  description = "Name of application"
}

variable "app_private_dns" {
  type        = "string"
  description = "DNS A entry without .sudsidia.org"
}
variable "owner" {
  type        = "string"
  description = "User ID of server reponsible, ex : JNAHEL22"
}

variable "user_sg_ids" {
  type        = "list"
  description = "List of extra security groups applied on instances"
  default     = []
}
