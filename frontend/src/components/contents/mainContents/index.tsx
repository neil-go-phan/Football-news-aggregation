import React from 'react';
import Tab from 'react-bootstrap/Tab';
import Tabs from 'react-bootstrap/Tabs';
import News from './news';
import Schedule from './schedule';
export default function MainContents() {
  return (
    <Tabs
      defaultActiveKey="news"
      id="uncontrolled-tab-example"
      className="mb-3"
    >
      <Tab eventKey="news" title="News">
        <News />
      </Tab>
      <Tab
        eventKey="schedule"
        title="Schedule"
      >
        <Schedule />
      </Tab>
    </Tabs>
  );
}
