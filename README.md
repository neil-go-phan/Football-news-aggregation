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
[![](https://mermaid.ink/img/pako:eNqNUsFuwyAM_RWL65ofyKHS1E67L1cuHjgdEpiOmFZV1X8fKE2zbuk0Tsb283vPcFYmWlKtGugzExvaOtwlDJqhnI13xNKs108dpQOlFl5JAJM442kYe9CIO6AQjC1jcowrcIuC7zjQEtTSIniCNAXfTMTP_yatrKPwn6gHfD7GPbyU6wmEGILjLBPk3swm4dFXNZ3bMXqQCKambr6gTzGAY6HEJPOIG-91wly5JqrqyeobHsGWHcxN34T_GrBo6dEb3O_jj50QW81qpQKlgM6WD3KuBa3kgwJp1ZbQUo_Zi1aaL6UVs8TuxEa1kjKtVN4XE9N_Um2PfqDLF6IgzYk?type=png)](https://mermaid.live/edit#pako:eNqNUsFuwyAM_RWL65ofyKHS1E67L1cuHjgdEpiOmFZV1X8fKE2zbuk0Tsb283vPcFYmWlKtGugzExvaOtwlDJqhnI13xNKs108dpQOlFl5JAJM442kYe9CIO6AQjC1jcowrcIuC7zjQEtTSIniCNAXfTMTP_yatrKPwn6gHfD7GPbyU6wmEGILjLBPk3swm4dFXNZ3bMXqQCKambr6gTzGAY6HEJPOIG-91wly5JqrqyeobHsGWHcxN34T_GrBo6dEb3O_jj50QW81qpQKlgM6WD3KuBa3kgwJp1ZbQUo_Zi1aaL6UVs8TuxEa1kjKtVN4XE9N_Um2PfqDLF6IgzYk)
