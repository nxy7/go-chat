# Real-Time Chat App 🚀

## Overview

Welcome to the Real-Time Chat App! This project is a showcase of real-time messaging using Golang, MongoDB, Redis, and Kubernetes. Talk with friends, family, or random strangers in a fast, scalable environment.

## How to run
### Prerequisites
```bash
- kubernetes cluster running with cilium kube-proxy replacement
- add `noxy.ddns.net:5000` insecure image registry to the cluster
```

### Run
```bash
kubectl apply -k ./k8s/dev
```

## Features

- 📬 Real-time Messaging: Chat in real-time thanks to WebSockets.
- 🙋 User Authentication: Sign in securely.
- ⚡️ Quick Access: Recently active chats and messages cached in Redis.
- 🛠 Scalability: Deployed in a Kubernetes cluster for easy scaling.

## Tech Stack

- Frontend: Vue
- Backend: Golang
- Authentication: Custom JWT Authentication managed by Backend
- Database: MongoDB
- Cache: Redis
- Orchestration: Kubernetes
