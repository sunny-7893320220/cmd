variable "vm_param" {
    description = "variables for the vm module"
    type = map(object({
      name = string
      machine_type = string
      zone = string
      tags = list(string)
      boot_disk = object({
        initialize_params = object({
          image = string
          size = number
        })
      })
    })) 
}