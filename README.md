# Football news aggregation
## Description
- Website aggregating news, schedule and automatic match information
## Admin
- To access: 
  - go to `/admin`
- Admin account: 
  - username: `admin2023`
  - password: `12345678`
## Tech stack
- Frontend: 
  - Nextjs
  - Bootstrapt
  - Code convention: eslint
- Backend: 
  - Crawler: goquery
  - Message between services: gRPC
  - Database: elasticsearch
  - Code convention: golangci-lint
## How to run
- Device dev must install: 
  - docker
  - golang version >1.20
  - nodejs version >18.15
- 
- Frontend: 
  - Open terminal
  - CD to frontend folder
  - Run commands:
    - `npm insall`
    - `npm run dev`
- Backend: 
  - Open terminal
  - CD to backend folder
  - Run commands:
    - `docker compose up`

## Sequence diagram

