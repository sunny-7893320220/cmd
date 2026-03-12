variable "artifacts_repository_param" {
    description = "variables for the artifacts repository"
    type = map(object({
      name = string
      format = string
      description = string
      location = string
      project = string
    }))
}