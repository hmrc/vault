#!groovy
pipeline {
  agent any

  stages {
    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Build') {
      steps {
        sh 'docker build -t vault-build -f scripts/cross/Dockerfile .'
        sh "docker run --rm -v ${WORKSPACE}:/gopath/src/github.com/hashicorp/vault -w /gopath/src/github.com/hashicorp/vault -e XC_OSARCH=linux/amd64 vault-build"
        sh "docker run --rm -v ${WORKSPACE}:/gopath/src/github.com/hashicorp/vault -w /gopath/src/github.com/hashicorp/vault alpine chmod -R 0777 ."
        sh 'aws s3 cp pkg/linux_amd64/vault s3://mdtp-vault-binary-d10af457daa1deed54e2c36b5f295e7e/vault --acl=bucket-owner-full-control'
        sh "docker run --rm -v /var/lib/jenkins/workspace/tools/vault-binary:/gopath/src/github.com/hashicorp/vault -w /gopath/src/github.com/hashicorp/vault alpine rm -rf pkg/"
      }
    }
  }
}
