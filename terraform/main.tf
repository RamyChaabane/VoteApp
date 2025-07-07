module "cluster" {
  source = "./modules/cluster"

  access_key = var.access_key
  secret_key = var.secret_key
  project_id = var.project_id
}

module "bootstrap" {
  source = "./modules/bootsrap"

  kube_host = module.cluster.kube_host
  kube_token = module.cluster.kube_token
  base64_cluster_ca_certificate = base64decode(module.cluster.cluster_ca_certificate)

  default_access_key = var.default_access_key
  default_secret_key = var.default_secret_key
  default_project_id = var.default_project_id
}

module "argocd" {
  source = "./modules/argocd"

  argocd_server_addr    = module.bootstrap.argocd_serer_addr
  argocd_admin_password = module.bootstrap.argocd_admin_password
}
