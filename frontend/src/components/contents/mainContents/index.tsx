import React, { useEffect } from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import News from './news';
import Schedule from './schedule';
import { useRouter } from 'next/router';
import { _ROUTES } from '@/helpers/constants';
export default function MainContents() {
  const router = useRouter();
  // useEffect(() => {
    // router.events.on('routeChangeStart', (url) => {
    // });
  // }, []);

  return (
    <Tabs
      defaultActiveKey="news"
      id="uncontrolled-tab-example"
      className="mb-3"
    >
      <Tab eventKey="news" title="Tin tức">
        <News />
      </Tab>
      <Tab
        eventKey="schedule"
        title="Lịch thi đấu"
        disabled={router.asPath === _ROUTES.NEWS_PAGE ? true : false}
      >
        <Schedule />
      </Tab>
    </Tabs>
  );
}
