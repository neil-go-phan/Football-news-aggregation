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
[![](https://mermaid.ink/img/pako:eNrNWFtv2jAU_itRnmlFgALhjdGsZaO04tJqE1JkkgNYc-zIcdoy4L_PuVBwEsr2MoIETc537Jy7v3SjO8wFvaMDv8VoyZE3o5r8dEeTfm9gjbVNch99AuAYEQ272tP3g1RgDwKBPF9zOCABro2EVgSHvvsZ7AKBFD7CA8ExXUo1QSAndSFwOPYFZjSHEUx_JcJd8mfSvSuNN2hpU-SBphiYhtyWhhbYibjADgE7treiff2e04h2zaD7nW8f-sOS-B7KZ0e-5wAfBcEb465i-eOzNXruWy92f2I9XD59yTM9JJxVFOqCJDgknOegxEFMFywv9NASbLH2oaDmPbCTayUmT4PuD2ukBANTccFIyFYDO_RPRKQ4WmnOCVoDt4vrIcVCbw5cdZWzN1XgMKKEaNy7t26nmfH134N0hMvFBb1AAC1DSEZBpmUHVvduemH754zJueMI_Ar_YvpDd9K7t61nazgpUZFmCvCwNro633tHmMOoACrUgpt05fAeR9_8qCqV46kTcrmwFU9yoJxktdOQkc95adotzZ48Mgvns_HJ7K6dwKIZJ0fcZ8uPVE7tEjgrcMPkFFeqEd5FphBjEWchdTOyJKsuCISJfWA5GTzKUxhkoMBhHAIlcb3B9EuZ0xbbrZ4PsYiwJSuYOoP-0LKnT6X3KKYJBW4tGJf5U_hskrkV5lG_EsbV9I26L4NLk4G_8DfkJCPZ81kXv55AMpxfXVXE_BWNgs6g8sf2Jd_KycV7LD_iYR_xfRx-eyxNh8ShlWgRBtRVkROd43D0RiSxwm6GVIXUhlfga9vDtOjFZKzttldX263yotLR5kDY_pSM37EKtaKjBmEaFNL6bbIimUVndHcs1t0fOFntPTk-VjtMBdXa5Hmpuac2VNjMmYfnWcCZBRnzzm3_wWrjgG2OFBO39IrugRwg2JUv9HHFznSxAlkCekdeurBAIREzfUZ3UhWFgo3X1NE7C0QCqOhJcab_BviQ-oj-ZEzeCx4mt3pno7_rnat6u31dNcymWa82amaj0azoayluVa-b9ZppGi2jbVRb5s2uov-OdzCub1pmo9mqmc1mtVpvGO3dH4beA8g?type=png)](https://mermaid.live/edit#pako:eNrNWFtv2jAU_itRnmlFgALhjdGsZaO04tJqE1JkkgNYc-zIcdoy4L_PuVBwEsr2MoIETc537Jy7v3SjO8wFvaMDv8VoyZE3o5r8dEeTfm9gjbVNch99AuAYEQ272tP3g1RgDwKBPF9zOCABro2EVgSHvvsZ7AKBFD7CA8ExXUo1QSAndSFwOPYFZjSHEUx_JcJd8mfSvSuNN2hpU-SBphiYhtyWhhbYibjADgE7treiff2e04h2zaD7nW8f-sOS-B7KZ0e-5wAfBcEb465i-eOzNXruWy92f2I9XD59yTM9JJxVFOqCJDgknOegxEFMFywv9NASbLH2oaDmPbCTayUmT4PuD2ukBANTccFIyFYDO_RPRKQ4WmnOCVoDt4vrIcVCbw5cdZWzN1XgMKKEaNy7t26nmfH134N0hMvFBb1AAC1DSEZBpmUHVvduemH754zJueMI_Ar_YvpDd9K7t61nazgpUZFmCvCwNro633tHmMOoACrUgpt05fAeR9_8qCqV46kTcrmwFU9yoJxktdOQkc95adotzZ48Mgvns_HJ7K6dwKIZJ0fcZ8uPVE7tEjgrcMPkFFeqEd5FphBjEWchdTOyJKsuCISJfWA5GTzKUxhkoMBhHAIlcb3B9EuZ0xbbrZ4PsYiwJSuYOoP-0LKnT6X3KKYJBW4tGJf5U_hskrkV5lG_EsbV9I26L4NLk4G_8DfkJCPZ81kXv55AMpxfXVXE_BWNgs6g8sf2Jd_KycV7LD_iYR_xfRx-eyxNh8ShlWgRBtRVkROd43D0RiSxwm6GVIXUhlfga9vDtOjFZKzttldX263yotLR5kDY_pSM37EKtaKjBmEaFNL6bbIimUVndHcs1t0fOFntPTk-VjtMBdXa5Hmpuac2VNjMmYfnWcCZBRnzzm3_wWrjgG2OFBO39IrugRwg2JUv9HHFznSxAlkCekdeurBAIREzfUZ3UhWFgo3X1NE7C0QCqOhJcab_BviQ-oj-ZEzeCx4mt3pno7_rnat6u31dNcymWa82amaj0azoayluVa-b9ZppGi2jbVRb5s2uov-OdzCub1pmo9mqmc1mtVpvGO3dH4beA8g)
