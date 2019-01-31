node {
    stage ("Fetch source code from Git"){
        git 'https://github.com/saburovga/simple-go-webapp'
    }
    stage ("Build the source code") {
        def go="/usr/local/go/bin/go"
        sh "${go} build -o app"
    }
    stage ("Build a docker image") {
        sh "docker build --tag=saburovga/mygoapp:0.0.${env.BUILD_NUMBER} ."
        echo "build number is ${env.BUILD_NUMBER}"
    }
    stage ("Push the docker image") {
        withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'pwd', usernameVariable: 'username')]) {
            sh "docker login -u ${username} -p ${pwd}"
        }
        sh "docker push saburovga/mygoapp:0.0.${env.BUILD_NUMBER}"
    }
    stage ("Deploy the image on AWS instance") {
        def DockerPull = "docker pull saburovga/mygoapp:0.0.${env.BUILD_NUMBER}"
        sshagent(['e68d8dc5-e4ee-41cb-a226-d6428db6a610']) {
            sh "ssh -o StrictHostKeyChecking=no ec2-user@54.175.169.242 ${DockerPull}"
        }
    }
    stage ("Stop any running container") {
        def DockerStop="\"if [ \"\$(ssh -o StrictHostKeyChecking=no ec2-user@54.175.169.242 \"docker ps -q|wc -l\")\" -gt 0 ]; then docker stop \$(ssh -o StrictHostKeyChecking=no ec2-user@54.175.169.242 \"docker ps -q\"); fi\""
        sshagent(['e68d8dc5-e4ee-41cb-a226-d6428db6a610']) {
            def DockerPS=""
            sh "ssh -o StrictHostKeyChecking=no ec2-user@54.175.169.242 ${DockerStop}"
        }        
    }
    stage ("Run the container on AWS instance") {
        def DockerRun = "docker run -d -p 8088:8088 saburovga/mygoapp:0.0.${env.BUILD_NUMBER}"
        sshagent(['e68d8dc5-e4ee-41cb-a226-d6428db6a610']) {
            sh "ssh -o StrictHostKeyChecking=no ec2-user@54.175.169.242 ${DockerRun}"
        }
    }
}