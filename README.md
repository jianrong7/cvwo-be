# Gossip with Go (Backend) [22/23 CVWO Winter Assignment]

A web forum built with React and Golang.

You can find the live version of the project [here](https://3.1.102.180.nip.io).
Postman API [here](https://www.postman.com/spacecraft-candidate-84168725/workspace/cvwo/collection/16590827-e7b9e933-5a8b-4297-939b-6372028f8dfc).
You can find the frontend [here](https://d3mj3t330xelda.cloudfront.net) and its repository [here](https://github.com/jianrong7/cvwo-fe).

## Submission Details

**Name:** Loh Jian Rong

**Matric No.:** A0252735A

[Mid-Assignment Writeup](https://docs.google.com/document/d/1-RYiu5qhJFxY_yzrtO3-t6H8u4rrveW-IbFkb_v6Nwo/edit?usp=sharing)

## Getting started on your local machine

1. Clone the reponsitory.

```
$ git clone git@github.com:jianrong7/cvwo-fe.git
```

2. Copy template env file (`.env.example`).

```
PORT=3000 # Port number that your server will listen to

DB_URI=YOUR_DB_URI # PostgreSQL DB_URI
# SAMPLE_DB_URI=postgres://USER:PASSWORD@localhost:5432/DB_NAME

JWT_SECRET=YOUR_JWT_SECRET # Random string to sign JWT Tokens

OPENAI_API=YOUR_OPENAI_API # OpenAI API Key to generate fetch AI generated posts

AWS_ACCESS_KEY_ID=YOUR_AWS_ACCESS_KEY # AWS S3 Key for uploading of profile pictures
AWS_SECRET_ACCESS_KEY=YOUR_AWS_SECRET_ACCESS_KEY # AWS S3 Key for uploading of profile pictures
AWS_DEFAULT_REGION=YOUR_AWS_DEFAULT_REGION # AWS S3 Key for uploading of profile pictures
```

3. Run the server!

```
<!-- If you have gin set up, you can use hot reload. -->
$ gin

<!-- If not, the default command would be. -->
$ go run main.go
```

## Tools/Technologies used

- Gin with Golang for backend framework
- GORM for object relational mapping
- AWS SDK to upload images to S3
- OpenAI SDK to connect with OpenAI
- bcrypt to hash password

## Database Schema

![Database schema](<https://cvwo-user-profiles.s3.ap-southeast-1.amazonaws.com/cvwo+(1).png>)

## Project Structure

```sh
backup.sh          # Cron job to backup PostgreSQL database onto the EC2 instance.
controllers        # Handles GET, POST, PUT, DELETE logic.
initializers       # Connect database, load environment variables, start AWS and OpenAI clients.
middleware         # Contains the CORS middleware as well as the authentication middlware.
models             # Contains the models as defined by the database schema.
routes             # Contains routes to endpoints, separated by respective entities.
utils              # Contains JWT related functions.
docker-compose.yml # Dockerize application and start database image simultaneously.
Dockerfile          # To dockerize application.
main.go            # Main app file
```

## Deployment

This app is deployed in an AWS EC2 instance. As I do not have a custom domain and I do not want to pay for one, I used a workaround to obtain HTTPS support.
Since the AWS EC2 instance did not allow us to generate SSL certificates, I made use of [Caddy](https://caddyserver.com/) to create a reverse proxy. This web server
is also handy in giving automatic HTTPS support. That way, I do not need a custom domain for HTTPS support, albeit sacrificing the "niceness" of the URL.

## Reflections

Reflections [here](https://github.com/jianrong7/cvwo-fe#reflections).
