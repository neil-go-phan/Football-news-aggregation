import { useRouter } from 'next/router'
import React, { useEffect, useState } from 'react'
import ArticleAdmin from './articles'
import { _ROUTES } from '@/helpers/constants'
import AdminLeagues from './leagues'
import AdminTags from './tags'

function AdminComponent() {
  const router = useRouter()
  const [path, setpath] = useState<string>()
  useEffect(() => {
    setpath(router.asPath)
  }, [router.asPath])
  
  const render = () => {
    switch (path) {
      case _ROUTES.ADMIN_LEAGUES:
        return <AdminLeagues />
      case (_ROUTES.ADMIN_TAGS):
        return <AdminTags />
      default:
        return <ArticleAdmin />
    }
  }

  return (
    <>
    {render()}
    </>
  )
}

export default AdminComponent