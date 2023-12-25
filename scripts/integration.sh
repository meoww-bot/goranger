#!/bin/bash

set -ux

function install_ranger() {

    docker network create ranger-env
    git clone https://github.com/kadensungbincho/apache-ranger-docker-poc.git
    cd apache-ranger-docker-poc/docker-composes/ranger
    docker-compose up -d

}

function wait_ranger() {
    # Wait for port to open
    counter=0
    while [ 1 ]; do
        curl http://127.0.0.1:6080
        r=$?
        echo $r
        if [ $r -eq 0 ]; then
            break
        fi
        counter=$((counter + 1))
        if [[ "$counter" -gt 18 ]]; then
            # Just fail because the port didn't open
            echo "Waited for three minutes and ranger didn't appear to start"
            docker logs ranger-admin
            exit 1
        fi
        echo "Waiting for ranger port to open"
        sleep 10
    done
}

function run_tests() {
    go test -v ./client/.
}

install_ranger
wait_ranger
run_tests