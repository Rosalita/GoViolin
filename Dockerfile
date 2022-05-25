
FROM golang:1.18-alpine3.15
WORKDIR /app
COPY . .
RUN go build -o main main.go
EXPOSE 8080 
CMD [ "/app/main" ]

FROM jenkins
COPY https.pem /var/lib/jenkins/cert
COPY https.key /var/lib/jenkins/pk
ENV JENKINS_OPTS --httpPort=-1 --httpsPort=8083 --httpsCertificate=/var/lib/jenkins/cert --httpsPrivateKey=/var/lib/jenkins/pk
EXPOSE 8083