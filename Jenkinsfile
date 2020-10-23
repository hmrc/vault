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
            agent { docker { image 'golang:1.14.7' } }
            steps{
                sh 'go get github.com/mitchellh/gox'
                sh 'make'
                archiveArtifacts 'bin/vault'
                stash includes: 'bin/', name: 'bin'
            }
        }

        stage('Upload to S3') {
            steps{
                unstash 'bin'
                sh 'aws s3 cp bin/vault s3://mdtp-vault-binary-d10af457daa1deed54e2c36b5f295e7e/vault --acl=bucket-owner-full-control'
            }
        }
    }
}

