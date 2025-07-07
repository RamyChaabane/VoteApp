resource "argocd_application" "vote_app" {
  metadata {
    name      = "vote-app"
    namespace = "argocd"
  }

  spec {
    project = "default"

    source {
      repo_url        = "https://github.com/RamyChaabane/VoteApp"
      path            = "k8s/overlays/dev"
      target_revision = "main"
    }

    destination {
      server    = "https://kubernetes.default.svc"
      namespace = "default"
    }

    sync_policy {
      automated {
        prune       = true
        self_heal   = true
      }

      sync_options = [
        "CreateNamespace=true"
      ]
    }
  }
}
