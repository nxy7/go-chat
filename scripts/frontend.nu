#!/usr/bin/env nu
use utils.nu

# display help message
export def main [] {
    help commands frontend.nu 
}

# start development (needs k8s cluster with service running)
export def `main dev` [] {
    try {
        telepresence helm install
        telepresence connect
    }
    try {
        telepresence intercept frontend --port 3000:http --mechanism tcp --namespace chat-app
    }
    cd $'(utils project-root)/frontend';
    bun run dev --host;
}

# build docker image and tag it
export def `main build` [] {
    cd $'(utils project-root)/frontend';
    docker build . --target prod -t noxy.ddns.net:5000/chat-app/frontend:latest
}

# push image to registry
export def `main push` [] {
    cd $'(utils project-root)/frontend';
    docker push noxy.ddns.net:5000/chat-app/frontend:latest
}

