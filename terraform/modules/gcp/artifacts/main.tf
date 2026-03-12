resource "google_artifact_registry_repository" "gcp_artifacts_repository" {

  for_each = var.artifacts_repository_param
  repository_id = each.value.name
  format = each.value.format
  description = each.value.description
  location = each.value.location
  project = each.value.project
}