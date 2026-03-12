resource "google_compute_instance" "vm_instance" {
    for_each = var.vm_param
    name = each.value.name
    machine_type = each.value.machine_type
    zone = each.value.zone
    tags = each.value.tags
    boot_disk {
        initialize_params {
            image = each.value.boot_disk.initialize_params.image
            size = each.value.boot_disk.initialize_params.size
        }
    }
    network_interface {
        network = "default"
    }
}