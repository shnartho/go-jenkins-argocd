## A straightforward website built with Go, backed by a Jenkins and ArgoCD-managed CI/CD pipeline.✌️

### 🔰 Overview 

#### 🟢 The website is running in localhost. To update and push the changes in the server, I have Jenkins for CI and argoCD for CD.
![](./images/overview.png)

#### 🟢 Here I prepared the Jenkins server for CI
![](./images/server.png)

#### 🟢 Let's update the code and push it to the git repository
![](./images/code.png)

#### 🟢 The code repository has been updated
![](./images/g.png)

#### 🟢 Update got triggered by the GitHub webhook
![](./images/gw.png)

#### 🟢 Jenkins pulling the updated code repository and executing shell commands
![](./images/j1.png)

#### 🟢 Jenkins workspace after pulling the latest repo
![](./images/j2.png)

#### 🟢 The website is running on Jenkins server port 8091
![](./images/pu.png)

#### 🟢 The website functionalities are working as expected
![](./images/pu2.png)



### 🖥️ Servers Setup 
1. EC2: To connect using aws ec2 instances, first check if "etc/ssh/sshd_config -> PasswordAuthentication yes", incase it is set to no then change it to yes otherwise using putty you cant connet to the ec2 instances. 
2. Jenkins: Install java runtime environment. For Jenkins installing follow https://www.jenkins.io/doc/book/installing/linux/. 
3. ArgoCD: For argocd installation follow 

🎯 In summary, Updating code and pushing it to git repository will trigger the webhook, jenkins will pull the updated code repo in the project workspace and execute necessary shell comands for port and running the website.

### 🖥️ Error Hanlidng
1. Css and static files were not being served. Solution: Make sure to get and set working directory is the executables directory and then use relative path in your code. (os.Chdir(exeDir), templates/index.html, static/styles.css etc)
2. Jenkins issues: sudo usermod -aG docker $username, for example "sudo usermod -aG docker jenkins" and "newgrp docker". And reboot the server. so jenkins will be able to access docker comands. A

