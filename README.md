# Todong-Des (like a drum)

## Local Development

### 1. Clone Repository
```bash
git clone https://github.com/username/todong-des.git
cd todong-des
```

### 2. Setup Environment
```bash
cp .env-example .env
```

Edit `.env`:
```env
DB_USER=root
DB_PASSWORD=secret
DB_HOST=db
DB_PORT=3306
DB_NAME=todo_db
```

### 3. Run Locally
```bash
docker compose up --build
```

### 4. Access
http://localhost


---

## Production Deployment (VPS)

### Prerequisites
- VPS dengan Ubuntu
- Docker & Docker Compose terinstall di VPS
- Akun Docker Hub
- Repository GitHub

### Step 1 — Install Docker di VPS
```bash
ssh root@IP_VPS

# Add Docker's official GPG key:
sudo apt update
sudo apt install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
sudo tee /etc/apt/sources.list.d/docker.sources <<EOF
Types: deb
URIs: https://download.docker.com/linux/ubuntu
Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
Components: stable
Architectures: $(dpkg --print-architecture)
Signed-By: /etc/apt/keyrings/docker.asc
EOF

sudo apt update

# install Docker CLI
sudo apt install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# Cek status Docker
sudo systemctl status docker
```

### Step 2 — Setup Folder di VPS
```bash
mkdir -p ~/todong-des/nginx
```

### Step 3 — Buat File ke VPS
Buat`docker-compose.prod.yml` dan `nginx.conf`: 
```bash
scp docker-compose.prod.yml root@IP_VPS:~/todong-des/
scp nginx/nginx.conf root@IP_VPS:~/todong-des/nginx/
```

### Step 4 — Buat `.env` di VPS
```bash
nano ~/todong-des/.env
```
```env
DB_USER=root
DB_PASSWORD=secret
DB_HOST=db
DB_PORT=3306
DB_NAME=todo_db
```

### Step 5 — Setup GitHub Secrets
Masuk ke GitHub repo → Settings → Secrets and variables → Actions → New repository secret

`DOCKERHUB_USERNAME` : Username Docker Hub
`DOCKERHUB_TOKEN` : Access token Docker Hub 
`VPS_HOST` : IP address VPS 
`VPS_USER` : User VPS
`VPS_PASS` : Password VPS 
`VPS_PORT` : Port SSH 

### Step 6 — GitHub Actions Workflow 
Buat file `.github/workflows/deploy.yml`:
```yaml
name: Deploy Todong

on:
  push:
    branches:
      - main

jobs:
  build-and-deploy:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout
        uses: actions/checkout@v3

      - name: Login to Docker Hub
        uses: docker/login-action@v4
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Build and push backend
        uses: docker/build-push-action@v5
        with:
          context: ./backend
          push: true
          tags: ${{ secrets.DOCKERHUB_USERNAME }}/todong-des-backend:latest

      - name: Build and push frontend
        uses: docker/build-push-action@v5
        with:
          context: ./frontend
          push: true
          tags: ${{ secrets.DOCKERHUB_USERNAME }}/todong-des-frontend:latest

      - name: Deploy to VPS
        uses: appleboy/ssh-action@v1
        with:
          host: ${{ secrets.VPS_HOST }}
          username: ${{ secrets.VPS_USER }}
          password: ${{ secrets.VPS_PASS }}
          port: ${{ secrets.VPS_PORT }}
          script: |
            cd ~/todong-des
            docker compose -f docker-compose.prod.yml down
            docker compose -f docker-compose.prod.yml pull
            docker compose -f docker-compose.prod.yml up -d
```

### Step 7 — Push ke GitHub
Setiap push ke branch `main` akan otomatis:
1. Build Docker image backend & frontend
2. Push image ke Docker Hub
3. Deploy ke VPS

```bash
git add .
git commit -m "deploy"
git push origin main
```

### Step 8 — Access
http://IP_VPS
