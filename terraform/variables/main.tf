module "gcp_vm" {
  source = "../modules/gcp/vm_module"

  vm_param = {
    vm1 = {
      name         = "vm1"
      machine_type = "n1-standard-1"
      zone         = "us-central1-a"
      tags         = ["web", "dev"]
      boot_disk = {
        initialize_params = {
          image = "debian-cloud/debian-11"
          size  = 50
        }
      }
    }
  }
}