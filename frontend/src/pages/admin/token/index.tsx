import { _ROUTES } from '@/helpers/constants'
import { setCookie } from 'cookies-next'
import { useRouter } from 'next/router'
import React, { useEffect } from 'react'

function Token() {
  const route = useRouter()
  useEffect(()=>{
    if (route.query.token) {
      const token = route.query.token
      setCookie('token', token)
      route.push(_ROUTES.ADMIN_PAGE)
    }
  })
  return (
    <div></div>
  )
}

export default Token