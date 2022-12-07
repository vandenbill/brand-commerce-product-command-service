pipeline {
    agent any
    stages {
        stage ('get source code') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/main']], extensions: [], userRemoteConfigs: [[url: 'https://github.com/vandenbill/brand-commerce-product-command-service']]])
            }
        }
        stage ('build source code') {
            steps {
                script {
                    sh 'docker build -t vandenbill/brand-commerce-product-command-service .'
                }
            }
        }
        stage ('push image to docker hub') {
            steps {
                script {
                    withCredentials([string(credentialsId: 'dockerhubpw', variable: 'dockerhubpw')]) {
                        sh 'docker login -u vandenbill -p ${dockerhubpw}'
                    }
                    sh 'docker image push vandenbill/brand-commerce-product-command-service'
                }
            }
        }
    }
}