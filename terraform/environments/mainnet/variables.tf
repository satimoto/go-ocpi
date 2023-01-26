variable "region" {
  description = "The AWS region"
  default     = "eu-central-1"
}

variable "availability_zones" {
  description = "A list of Availability Zones where subnets and DB instances can be created"
}

variable "deployment_stage" {
  description = "The deployment stage"
  default     = "mainnet"
}

variable "forbidden_account_ids" {
  description = "The forbidden account IDs"
  type        = list(string)
  default     = []
}

# -----------------------------------------------------------------------------
# Module service-ocpi
# -----------------------------------------------------------------------------


variable "rds_satimoto_db_password_ssm_key" {
  description = "Systems Manager key where the password for the satimoto DB is stored"
}

variable "service_name" {
  description = "The name of the service"
}

variable "service_desired_count" {
  description = "The number of instances of the task definition to place and keep running"
}

variable "service_container_port" {
  description = "The port on the container to associate with the load balancer"
}

variable "service_metric_port" {
  description = "The port to associate with metric collection"
}

variable "task_network_mode" {
  description = "The Docker networking mode to use for the containers in the task"
}

variable "task_cpu" {
  description = "The number of cpu units used by the task"
}

variable "task_memory" {
  description = "The amount (in MiB) of memory used by the task"
}

variable "target_health_path" {
  description = "The path to check the target's health"
}

variable "target_health_interval" {
  description = "The approximate amount of time, in seconds, between health checks of an individual target"
}

variable "target_health_timeout" {
  description = "The amount of time, in seconds, during which no response means a failed health check"
}

variable "target_health_matcher" {
  description = "The HTTP codes to use when checking for a successful response from a target"
}

variable "subdomain_name" {
  description = "The subdomain name of the service"
}

variable "env_country_code" {
  description = "The environment variable to set the country code"
}

variable "env_party_id" {
  description = "The environment variable to set the party id"
}

variable "env_issuer" {
  description = "The environment variable to set the issuer"
}

variable "env_fcm_api_key" {
  description = "The environment variable to set the FCM API key"
}

variable "env_rpc_port" {
  description = "The environment variable to set the RPC port"
}

variable "env_record_evse_status_periods" {
  description = "The environment variable to set if EVSE status periods are recorded"
}

variable "env_shutdown_timeout" {
  description = "The environment variable to set the shutdown timeout"
}

variable "env_token_authorization_timeout" {
  description = "The environment variable to set the token authorization timeout"
}

variable "env_wait_for_evse_status_timeout" {
  description = "The environment variable to set the wait for evse status timeout"
}
