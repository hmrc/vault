#!groovy
pipeline {
    agent any
    environment {
        XDG_CACHE_HOME = "/tmp"  // so golang is able to use the cache
    }

    stages {
        stage('Checkout') {
          steps {
            checkout scm
          }
        }

        stage('Build') {
          agent {
            dockerfile {
              filename 'scripts/cross/Dockerfile'
              args "-v ${env.WORKSPACE}:/gopath/src/github.com/hashicorp/vault -w /gopath/src/github.com/hashicorp/vault"
            }
          }
          steps{
            sh 'docker build -t vault-build -f scripts/cross/Dockerfile .'
            sh "docker run --rm -v ${WORKSPACE}:/gopath/src/github.com/hashicorp/vault -w /gopath/src/github.com/hashicorp/vault -e XC_OSARCH=linux/amd64 vault-build"
            archiveArtifacts 'pkg/linux_amd64/vault'
            stash includes: 'pkg', name: 'pkg'
          }
        }

        stage('Upload to S3') {
            steps {
              unstash 'pkg'
              sh 'aws s3 cp pkg/linux_amd64/vault s3://mdtp-vault-binary-d10af457daa1deed54e2c36b5f295e7e/vault --acl=bucket-owner-full-control'
            }
        }
    }
}

