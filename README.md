# 📬 mailing-service

**An implementation of a scalable mailing service in Go**

This project is a containerized, scalable mailing service written in Go. It uses Redis for queueing, NGINX for load balancing, and supports asynchronous email dispatch via workers running in parallel. Here i implemented course rejected mail notification as an example, but the interfaces may be changed so as to adapt to the need accordingly.

---
### 🔧 Prerequisites

- Docker & Docker Compose
- A Gmail account with [App Passwords](https://support.google.com/accounts/answer/185833?hl=en)
- (Optional) Go 1.18+ if running outside Docker

---
### Steps to set-up
- Make sure Docker Daemon is running.
- make a ```.env``` with the same format as ```.env.example``` in the same directory.
- run the following command to run the whole service.
  ```bash
  docker-compose up -d
  ```
- the service is avilable at ```localhost:8080```.
- format of request
  ```bash
  curl --location 'http://localhost:8080/send' \
  --header 'Content-Type: application/json' \
  --data-raw '{
      "student_mail":"example@gmail.com",
      "course_id":"Course-Code-Here"
  }'
  ```


### System Architecture
```mermaid
flowchart TD
    A[Client Request] --> C(Load Balancer)
    C --> D[fa:fa-server server-1]
    C --> E[fa:fa-server server-2]
    C --> F[fa:fa-server server-3]
    G(Push to Queue)
    E --> G
    D --> G
    F --> G
    J(SMTP server)
    H(Server Thread)
    G --> H
    H -->|Send Mail| J

```
