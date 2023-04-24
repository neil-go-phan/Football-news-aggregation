import React, { FunctionComponent, ReactNode, useState } from 'react';
import { MatchEvent, MatchProgress } from './index';

type Props = {
  matchProcess: MatchProgress | null;
};

const MatchEventsComponent: FunctionComponent<Props> = ({ matchProcess }) => {
  const [expanded, setExpanded] = useState(false);
  const renderEvents = (events: Array<MatchEvent>): ReactNode => {
    const eventsForDisplay = expanded ? events : events.slice(0, 10);
    return (
      <>
        {eventsForDisplay.map((event) => {
          return (
            <div
              className="event-detail d-flex"
              key={`event-detail-${event.time}-${event.content}`}
            >
              <div className="time col-3">
                <span>{event.time}</span>
              </div>
              <div className="content col-9">
                <span>{event.content}</span>
              </div>
            </div>
          );
        })}
      </>
    );
  };
  if ((matchProcess) && (matchProcess.events)) {
    return (
      <div id="process" className="matchDetail__content--process">
        <div className="title">Match process</div>
        <div className="events">
          {matchProcess.events ? renderEvents(matchProcess.events) : <></>}
          {matchProcess.events!.length > 10 ? (
            <p className="showmore" onClick={() => setExpanded(!expanded)}>
              {expanded ? 'Show less' : 'Show more...'}
            </p>
          ) : (
            <></>
          )}
        </div>
      </div>
    );
  }
  return <div className="matchDetail__content--process" id="process"></div>;
};

export default MatchEventsComponent;
