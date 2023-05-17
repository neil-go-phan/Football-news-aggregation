import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import ArticleAdmin from './articles';
import { _ROUTES } from '@/helpers/constants';
import AdminLeagues from './leagues';
import AdminTags from './tags';
import Crawler from './crawler';
import AddCrawler from './crawler/addCrawler';

function AdminComponent() {
  const router = useRouter();
  const [path, setpath] = useState<string>();
  useEffect(() => {
    const path = router.asPath;
    const beforeQuestionMark = path.split('?')[0];
    setpath(beforeQuestionMark);
  }, [router.asPath]);
  const render = () => {
    switch (path) {
      case _ROUTES.ADMIN_LEAGUES:
        return <AdminLeagues />;
      case _ROUTES.ADMIN_TAGS:
        return <AdminTags />;
      case _ROUTES.ADMIN_CRAWLER:
        return <Crawler />;
      case _ROUTES.ADD_CRAWLER:
        return <AddCrawler />;
      default:
        return <ArticleAdmin />;
    }
  };

  return <>{render()}</>;
}

export default AdminComponent;
