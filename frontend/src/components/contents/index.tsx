import React from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import News from './news';
import Schedule from './schedule';
import { useRouter } from 'next/router';
export default function Contents() {
  const router = useRouter();
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
        disabled={router.asPath === '/' ? true : false}
      >
        <Schedule />
      </Tab>
    </Tabs>
  );
}
