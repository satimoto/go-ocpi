region = "eu-central-1"

availability_zones = ["eu-central-1a", "eu-central-1b", "eu-central-1c"]

deployment_stage = "mainnet"

forbidden_account_ids = ["909899099608"]

# -----------------------------------------------------------------------------
# Module service-ocpi
# -----------------------------------------------------------------------------

rds_satimoto_db_password_ssm_key = "/rds/satimoto_db_password"

service_name = "ocpi"

service_desired_count = 1

service_container_port = 9001

service_metric_port = 9101

task_network_mode = "awsvpc"

task_cpu = 256

task_memory = 512

target_health_path = "/health"

target_health_interval = 120

target_health_timeout = 5

target_health_matcher = "200"

subdomain_name = "ocpi"

env_country_code = "DE"

env_party_id = "BTC"

env_issuer = "Satimoto"

env_fcm_api_key = "AAAA3ZJBhzw:APA91bHCxjgYChy0QnDfUihevkVyni_klXxH5GkVLdHAdcjgnWbSAxnpeP9b0GmMiUTPbStB8uAzNw147CPUWbbBlCMUDiFOCMp9Mqp9YGNZhYHTiv0AMSV3BAAmWn6_vQraENT4CTQ8"

env_rpc_port = 50000

env_shutdown_timeout = 20
