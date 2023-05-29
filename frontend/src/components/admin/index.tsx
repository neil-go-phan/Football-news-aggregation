import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import ArticleAdmin from './articles';
import { _ROUTES } from '@/helpers/constants';
import AdminLeagues from './leagues';
import AdminTags from './tags';
import AddCrawler from './crawler/addCrawler';
import CrawlerComponent from './crawler';
import AdminCronjob from './cronjob';
import Dashboard from './dashboard';

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
        return <CrawlerComponent />;
      case _ROUTES.ADD_CRAWLER:
        return <AddCrawler />;
      case _ROUTES.ADMIN_CRONJOB:
        return <AdminCronjob />
      case _ROUTES.ADMIN_ARTICLES:
        return <ArticleAdmin />;
      default:
        return <Dashboard />
    }
  };

  return <>{render()}</>;
}

export default AdminComponent;
