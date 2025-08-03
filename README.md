# 🗳 VoteApp

A production-grade, cloud-native voting application built with a clean Go backend, modern Vue.js frontend, and a fully automated GitOps-powered CI/CD pipeline.

---
## 🌟 Overview

**VoteApp** is a microservice-style voting application designed to demonstrate:

- Robust **backend architecture** using Go and Clean Architecture principles
- Modern **frontend UI** built with Vue 3 + Vite
- End-to-end **DevOps** automation including:
    - CI/CD with GitHub Actions
    - Docker image builds and versioning
    - GitOps-driven deployment via **Argo CD**
    - Dependency updates managed with **Renovate**
    - Automated **E2E tests** after deployment
- Tag-based release promotion: every push to `main` triggers a deterministic release based on timestamp and run number

This project simulates real-world infrastructure, workflows, and testing strategies adopted by top engineering teams.

---

## 📁 Repository Structure

```
vote-app/
├── backend/ # Go backend API using Clean Architecture
│ ├── cmd/server # Entry point
│ └── internal/ # Domain logic, use cases, infrastructure, handlers
├── frontend/ # Vue 3 frontend served via Nginx
├── docker-compose.yaml # Local dev stack
├── scripts/ # Workflow coordination scripts (Renovate, Argo CD sync)
└── .github/workflows/ci.yml # Full CI pipeline (build, release, test)
```

---

## ⚙️ Architecture

### Backend (Go)

- Follows **Clean Architecture**:
    - `domain/`: vote entity and validation
    - `usecase/`: business logic
    - `infrastructure/redis`: persistent vote storage
    - `handler/`: HTTP handlers and routes
- Unit tested using **Testify** and **table-driven tests**
- Redis-backed data store with mocked interfaces for testing

### Frontend (Vue.js)

- Built with **Vue 3** and **Vite**
- Simple but extensible UI
- Served via Nginx (configured with `nginx.conf`)

---

## 🚀 CI/CD & GitOps Workflow

Every time code is merged to `main`:

1. ✅ **Unit tests** run for backend services
2. ✅ **Docker images** are built and pushed to Docker Hub:
    - `vote-frontend`
    - `vote-backend`
    - Tagged as `vYYYY.MM.DD.RUN_ID`
3. ✅ **Renovate** is triggered in the [`argocd-voteapp`](https://github.com/RamyChaabane/argocd-voteapp) repo to update the deployment manifests
4. ✅ A Renovate PR is created and **auto-merged** (thanks to semantic PR titles + labels)
5. ✅ The workflow **waits for Argo CD** to finish syncing the app
6. ✅ **E2E tests** are executed against the live dev environment

This approach guarantees only tested, deployed code reaches the environment.

---

## ✅ E2E Tests

Once the application is deployed and synced via Argo CD, the workflow runs integration tests using `curl`:

```bash
# Valid votes
curl -X POST https://voteapp.dev.rch.domain-buy-5.com/vote?vote=Cats    # expect 200
curl -X POST https://voteapp.dev.rch.domain-buy-5.com/vote?vote=Dogs    # expect 200

# Invalid votes
curl -X POST https://voteapp.dev.rch.domain-buy-5.com/vote?vote=Unicorn # expect 400
curl -X POST https://voteapp.dev.rch.domain-buy-5.com/vote?vote=Cat     # expect 400 (typo)
curl -X POST https://voteapp.dev.rch.domain-buy-5.com/vote              # expect 400 (missing param)
```

These tests confirm the system is functionally correct in a production-like environment.

## 🏷️ Image Tagging Strategy

Docker images are tagged using the format: vYYYY.MM.DD.RUN_ID

For example:
```
v2025.07.26.16543231189
```

### ✅ Benefits

- **Uniqueness**: Each tag corresponds to a single GitHub Actions run.
- **Traceability**: Tags map directly to a specific CI run and commit.
- **Deterministic promotion**: GitOps tools like Argo CD can safely deploy new images based on tags, not `latest`.

---

## 🔄 Renovate Pull Requests

After pushing new images, the [Renovate](https://github.com/renovatebot/renovate) bot is triggered in the [argocd-voteapp](https://github.com/RamyChaabane/argocd-voteapp) repository.

### 🧠 What It Does

- Scans for new tags of `vote-backend` and `vote-frontend`
- Creates a PR updating the Kubernetes manifests with the new image tag
- PR title example: chore(deps): update vote app dev to v2025.07.26.16543231189
- Adds labels: `env::dev`
- The PR is automatically merged when the branch passes CI checks
---

## 🧪 Test Coverage

| Type           | Description                                        |
|----------------|----------------------------------------------------|
| ✅ Unit Tests   | Clean architecture logic (entities, use cases)     |
| ✅ E2E Tests    | Live curl-based HTTP tests after deployment        |

### 🧪 E2E Test Cases

| Test Description                     | Expected Status |
|-------------------------------------|-----------------|
| Valid vote: `Cats`                  | `200 OK`        |
| Valid vote: `Dogs`                  | `200 OK`        |
| Invalid vote: `Cat` (typo)          | `400 BadRequest`|
| Invalid vote: `Unicorn`             | `400 BadRequest`|
| Missing vote param                  | `400 BadRequest`|

These tests ensure the application works end-to-end in the live environment post-deployment.

---

## 📦 Technologies Used

| Category        | Technology                          |
|------------------|-------------------------------------|
| **Backend**       | Go, Clean Architecture              |
| **Frontend**      | Vue 3, Vite, Nginx                  |
| **CI/CD**         | GitHub Actions, Renovate            |
| **GitOps**        | Argo CD                             |
| **Testing**       | Testify (unit tests), `curl` (E2E)  |
| **Registry**      | Docker Hub                          |
| **Infrastructure**| Deployed via Argo CD (see [argocd-voteapp](https://github.com/RamyChaabane/argocd-voteapp)) |

---

## 👔 Why This Project Stands Out

This repository goes beyond just building an app — it showcases how real-world teams operate:

- **Well-architected Go codebase** following clean separation of concerns
- **CI/CD pipelines** that are production-grade and observable
- **Complete GitOps automation** with controlled rollouts
- **Automated validation gates** via unit and E2E tests
- **Release traceability and control** via semantic Docker tags and Renovate-managed PRs

---

## 🔗 Related Repositories
- [`argocd-voteapp`](https://github.com/RamyChaabane/argocd-voteapp) – GitOps deployment definitions managed by Argo CD
- [`provisioning-voteapp`](https://github.com/RamyChaabane/infrastructure) – Terraform for cluster & cloud provisioning
---
