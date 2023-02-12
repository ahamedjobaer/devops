pipeline {
  agent any
  stages {
    stage('Unit test') {
      steps {
        sh 'go test ./my_math'
      }
    }

  }
  tools {
    go 'go-1.19'
  }
  environment {
    GO119MODULE = 'on'
  }
}