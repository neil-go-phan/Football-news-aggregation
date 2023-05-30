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
[![](https://mermaid.ink/img/pako:eNrNWFtv2jAU_itRnmnVhEIhb4xmLSulFZdWm5AikxzAmmNHjtOWAf99zoWCk1C0l5FIgHO-Y_vcfcxad5kHuqUDv8VowZE_pZp8OsNxr9u3R9o6fY-fEDhGRMOe9vywpwrsQyiQH2guByTAc5DQyuAo8L6CPSCQwQd4KDimC8kmCBSoHoQux4HAjBYwgunvlLhNf8adu8pogxYORT5oioCZyR0paImciAvsEnASeWva94cCR7xqDt2tfPvYG1RE90juHeteAAIUhu-Me4rkTy_28KVnvzq9sf14fvele_pIuMvY1CVOcEk0K0CpgpjOWZHoowU4YhVAScz74KRjxSbP_c5Pe6gYA1NxRkvIVAMnCo5YpNxamc8JWgF3yuMhwyJ_BlxVlbN3leAyopho1L23bye58vXfjXSAy8kluUAALSJIS0EuZft2525yZvlnjMm64wr8Bv8i-mNn3L137Bd7MK5QkOYCcD83Hp3OvQPMZVQAFWrAjTuyeI_iT7FUVUrxTAk5XTiKJgVQVjLzOGQUfV6ZdMu8J4_M0vpsfFG7zSNYXONkiftq-gHLsVVCdwlelJ7iSjTCh8gFYkLiLKJejpZ61QOBMHH2XU4Oj_0UhTkodBmHUHFctz_5VmW3JXKr50NCImzBSqpOvzewnclz5TVK2oQSteaMS_8p_WzquSXmcb4SxlX3DTuv_XM3A4mEESc5yq5j9fDbESTX1auzynp7haMk9qn8cgLZURXo4iOhH3RanxZ8Gvx4qkwOJMEi0TIMqKciR3LD5eidyNYJe7m2KaIOvAFfOT6mZVePkbbdXFxsNspVxNJmQNjuHExuUaVc8WGCMA1LG_dNOiOtNid4tyzh3R0pee5d-3vIts97Vdp0v0zcYwsq_cqJzYvn_IkJOfFOLf_ZtyYGWx8wKmrtYjbTbFcFLOlmGSrUhae5XtN9kMUEe_Jyn8T2VBdLkMGiW3LowRxFREz1Kd1KVhQJNlpRV7fmiIRQ09Mwzv4S-KQGiP5iTL4LHqWvurXWP3TrwjDMm0ujabTaZqN906y3WjV9pVtG-9Iw69fta_Oq1TTaptna1vQ_yRLmZd2omzdG87reajSuGu369i-jvQb2?type=png)](https://mermaid.live/edit#pako:eNrNWFtv2jAU_itRnmnVhEIhb4xmLSulFZdWm5AikxzAmmNHjtOWAf99zoWCk1C0l5FIgHO-Y_vcfcxad5kHuqUDv8VowZE_pZp8OsNxr9u3R9o6fY-fEDhGRMOe9vywpwrsQyiQH2guByTAc5DQyuAo8L6CPSCQwQd4KDimC8kmCBSoHoQux4HAjBYwgunvlLhNf8adu8pogxYORT5oioCZyR0paImciAvsEnASeWva94cCR7xqDt2tfPvYG1RE90juHeteAAIUhu-Me4rkTy_28KVnvzq9sf14fvele_pIuMvY1CVOcEk0K0CpgpjOWZHoowU4YhVAScz74KRjxSbP_c5Pe6gYA1NxRkvIVAMnCo5YpNxamc8JWgF3yuMhwyJ_BlxVlbN3leAyopho1L23bye58vXfjXSAy8kluUAALSJIS0EuZft2525yZvlnjMm64wr8Bv8i-mNn3L137Bd7MK5QkOYCcD83Hp3OvQPMZVQAFWrAjTuyeI_iT7FUVUrxTAk5XTiKJgVQVjLzOGQUfV6ZdMu8J4_M0vpsfFG7zSNYXONkiftq-gHLsVVCdwlelJ7iSjTCh8gFYkLiLKJejpZ61QOBMHH2XU4Oj_0UhTkodBmHUHFctz_5VmW3JXKr50NCImzBSqpOvzewnclz5TVK2oQSteaMS_8p_WzquSXmcb4SxlX3DTuv_XM3A4mEESc5yq5j9fDbESTX1auzynp7haMk9qn8cgLZURXo4iOhH3RanxZ8Gvx4qkwOJMEi0TIMqKciR3LD5eidyNYJe7m2KaIOvAFfOT6mZVePkbbdXFxsNspVxNJmQNjuHExuUaVc8WGCMA1LG_dNOiOtNid4tyzh3R0pee5d-3vIts97Vdp0v0zcYwsq_cqJzYvn_IkJOfFOLf_ZtyYGWx8wKmrtYjbTbFcFLOlmGSrUhae5XtN9kMUEe_Jyn8T2VBdLkMGiW3LowRxFREz1Kd1KVhQJNlpRV7fmiIRQ09Mwzv4S-KQGiP5iTL4LHqWvurXWP3TrwjDMm0ujabTaZqN906y3WjV9pVtG-9Iw69fta_Oq1TTaptna1vQ_yRLmZd2omzdG87reajSuGu369i-jvQb2)
