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
  - Framework: Gin
  - Docker
  - Crawler: goquery
  - Message between services: gRPC
  - Database: postgres
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

## ERD 
[![](https://mermaid.ink/img/pako:eNrNV9tu4jAQ_ZUoz7TiUkLhjaVpi0ppxaXVrpAskwxgrRNHjtOWAv--ThwKTkLZfVlAApI5M_bcPV6ZDnPBbJnAbwiec-xNfEN-2oNRt9Ozh8ZKvcefEDjB1CCu8fywowriQSiwFxgOByzARVgYRXAUuN_BLlBI4T08FJz4c8kmKOSoLoQOJ4EgzM9hlPi_FXGj_kbtu7OxBs-Rjz0wNAVTlyOpaIGemAviUECJviXj9iHHEa-aQbcr3zx2-2dieyT3jm3PAQEOw3fGXU3zpxd78NK1X1F3ZD-ePnxqTw8LZxG7uiAIDo2mOUgZSPwZyxM9PAcklgEU5LwHSD1rPnnutX_aA80ZxBcn9IQsNUBRcMAjxd5KY07xEjgqzocUi7wpcN1Uzt51gsOo5qJh596-GWfa13930h4uhQtqgQKeR6BaQaZke3b7bnxi_aeMyb7jCPIG_6L6Y3vUuUf2i90fnVGSZhJwJxs_Ha-9PcxhvgBf6Ak3asvmPYy_-VZ1VoanRkhxgTRLcqDsZNXDUCUf87MptzR68sgs7M-Vb3p39QAW9zjZ4r4T32M5tEroLMCN1CmuZSN8iEwiJiTOIt_N0FRUXRCYULSbcjJ4HKcozEChwziEWuA6vfGPcw5bord-PiQkyuasoOv0un0bjZ_P3qJkTCgwa8a4jJ82z6rILQiP65UyrofvqX_bvUOdQfu1d-qZ4C_MjjjNULZjrUveDiCZ0V-XKroAaBwFBeLLHxTIsStHFx8JfW8c22RuQ5v1xcV6rY3qLWMKlG3PieSWUcgVN1tM_LBwsF0rCVWNR3g3LOHdttws93Y83Gfb1YWurdovVffQgtp5fmTz_Dl4RCCj3rHlv-a6xGGrPUZlllkyPZAlRFx5pU2KYWKKBcg6M1vy0YUZjqiYmBN_I1lxJNhw6Ttma4ZpCCVT5X16Ef6iBtj_xZh8FzxSr2ZrZX6YrYtqo3JZrjTrjbpVq1lX9YpVMpeSXru-lIRyuWw1y1ateb0pmZ_JCpXLeqN5ZTWqTcsql2tXFWvzBzYqtew?type=png)](https://mermaid.live/edit#pako:eNrNV9tu4jAQ_ZUoz7TiUkLhjaVpi0ppxaXVrpAskwxgrRNHjtOWAv--ThwKTkLZfVlAApI5M_bcPV6ZDnPBbJnAbwiec-xNfEN-2oNRt9Ozh8ZKvcefEDjB1CCu8fywowriQSiwFxgOByzARVgYRXAUuN_BLlBI4T08FJz4c8kmKOSoLoQOJ4EgzM9hlPi_FXGj_kbtu7OxBs-Rjz0wNAVTlyOpaIGemAviUECJviXj9iHHEa-aQbcr3zx2-2dieyT3jm3PAQEOw3fGXU3zpxd78NK1X1F3ZD-ePnxqTw8LZxG7uiAIDo2mOUgZSPwZyxM9PAcklgEU5LwHSD1rPnnutX_aA80ZxBcn9IQsNUBRcMAjxd5KY07xEjgqzocUi7wpcN1Uzt51gsOo5qJh596-GWfa13930h4uhQtqgQKeR6BaQaZke3b7bnxi_aeMyb7jCPIG_6L6Y3vUuUf2i90fnVGSZhJwJxs_Ha-9PcxhvgBf6Ak3asvmPYy_-VZ1VoanRkhxgTRLcqDsZNXDUCUf87MptzR68sgs7M-Vb3p39QAW9zjZ4r4T32M5tEroLMCN1CmuZSN8iEwiJiTOIt_N0FRUXRCYULSbcjJ4HKcozEChwziEWuA6vfGPcw5bord-PiQkyuasoOv0un0bjZ_P3qJkTCgwa8a4jJ82z6rILQiP65UyrofvqX_bvUOdQfu1d-qZ4C_MjjjNULZjrUveDiCZ0V-XKroAaBwFBeLLHxTIsStHFx8JfW8c22RuQ5v1xcV6rY3qLWMKlG3PieSWUcgVN1tM_LBwsF0rCVWNR3g3LOHdttws93Y83Gfb1YWurdovVffQgtp5fmTz_Dl4RCCj3rHlv-a6xGGrPUZlllkyPZAlRFx5pU2KYWKKBcg6M1vy0YUZjqiYmBN_I1lxJNhw6Ttma4ZpCCVT5X16Ef6iBtj_xZh8FzxSr2ZrZX6YrYtqo3JZrjTrjbpVq1lX9YpVMpeSXru-lIRyuWw1y1ateb0pmZ_JCpXLeqN5ZTWqTcsql2tXFWvzBzYqtew)