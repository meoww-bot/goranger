#!/bin/bash

set -eux


function install_ranger() {

    git clone https://github.com/kadensungbincho/apache-ranger-docker-poc.git
    cd apache-ranger-docker-poc/docker-composes/ranger
    docker-compose up -d

}

function run_tests() {
    
}



install_ranger