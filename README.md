# Football news aggregation
## Description
- Website automatically aggregates football news, schedule and match information
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
  - Log: Sentry
- Deploy: 
  - Frontend: Vercel
    - Link: https://football-news-aggregation.vercel.app/
  - Backend: Digital Ocean droplets
## How to run
- Device dev must install: 
  - docker
  - golang version >1.20
  - nodejs version >18.15
- Add file .env following file .env.example (or just delete the .example part in file name)
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
[![](https://mermaid.ink/img/pako:eNqNUsFuwyAM_RWL65ofyKHS1E67L1cuHnE6JDAdMa2qqv8-EE2zbuk0Tsb283vPcFYm9KRaNdJnIja0tbiL6DVDPhtniaVZr586igeKLbySAEaxxtFYe9CIPaAQ1JaarHEBblHwHUdagva0CJ4gTcY3E_Hzv0kLaxX-E_WAz4Wwh5d8PYEQg7echGrp3ssm4tEVMZ3dMTqQAKakbrZgiMGDZaHIJPOIG-11wly5JoroyekbHqHPK5ibvun-NWDR0aMnuF_HHysh7jWrlfIUPdo-_49zKWglH-RJqzaHPQ2YnGil-ZJbMUnoTmxUKzHRSqV9NjF9J9UO6Ea6fAGo-c0W?type=png)](https://mermaid.live/edit#pako:eNqNUsFuwyAM_RWL65ofyKHS1E67L1cuHnE6JDAdMa2qqv8-EE2zbuk0Tsb283vPcFYm9KRaNdJnIja0tbiL6DVDPhtniaVZr586igeKLbySAEaxxtFYe9CIPaAQ1JaarHEBblHwHUdagva0CJ4gTcY3E_Hzv0kLaxX-E_WAz4Wwh5d8PYEQg7echGrp3ssm4tEVMZ3dMTqQAKakbrZgiMGDZaHIJPOIG-11wly5JoroyekbHqHPK5ibvun-NWDR0aMnuF_HHysh7jWrlfIUPdo-_49zKWglH-RJqzaHPQ2YnGil-ZJbMUnoTmxUKzHRSqV9NjF9J9UO6Ea6fAGo-c0W)
