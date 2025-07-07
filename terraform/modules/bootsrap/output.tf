output "dns_zone" {
  value =  format("argocd.%s", scaleway_domain_record.argo_hostname.dns_zone)
}

output "argocd_admin_password" {
  value = data.kubernetes_secret.argocd_admin_password.data["password"]
  sensitive = true
}
