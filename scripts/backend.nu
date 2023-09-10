#!/usr/bin/env nu
use utils.nu

export def main [] {
    help commands main
}

export def `main dev` [] {
    try {
        telepresence intercept backend --port 8282:http --mechanism tcp --namespace chat-app
    }
    cd $'(utils project-root)/backend';
    go run main.go
}

export def `main build` [] {
    cd $'(utils project-root)/backend';
    docker build . -f Dockerfile -t noxy.ddns.net:5000/chat-app/backend:latest
}

export def `main push` [] {
    cd $'(utils project-root)/backend';
    docker push noxy.ddns.net:5000/chat-app/backend:latest
}

