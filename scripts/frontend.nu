#!/usr/bin/env nu
use utils.nu

export def main [] {
    help commands main
}

export def `main dev` [] {
    # kubectl scale --replicas=0 deployment/frontend -n streampai
    try {
        telepresence intercept frontend --port 3000:http --mechanism tcp --namespace streampai
    }
    cd $'(utils project-root)/frontend';
    bun run dev;
}

export def `main build` [] {
    cd $'(utils project-root)/frontend';
    docker build . -t noxy.ddns.net:5000/chat-app/frontend:latest
}

export def `main push` [] {
    cd $'(utils project-root)/frontend';
    docker push noxy.ddns.net:5000/chat-app/frontend:latest
}

