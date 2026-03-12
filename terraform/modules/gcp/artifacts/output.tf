output "artifacts_repository_name" {
    value = google_artifact_registry_repository.gcp_artifacts_repository[*].name
  
}

output "artifacts_repository_id" {
    value = google_artifact_registry_repository.gcp_artifacts_repository[*].id
}

output "artifacts_repository_format" {
    value = google_artifact_registry_repository.gcp_artifacts_repository[*].format
}

output "artifacts_repository_description" {
    value = google_artifact_registry_repository.gcp_artifacts_repository[*].description
}