import React, { FunctionComponent } from 'react';
import { Schedules } from '.';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronRight } from '@fortawesome/free-solid-svg-icons';

type Props = {
  schedule: Schedules | undefined;
};

const ScheduleContent: FunctionComponent<Props> = ({ schedule }) => {
  // TODI: refactor this duck-typing login
  if (schedule) {
    if (schedule.schedule_on_leagues !== null) {
      const title = new Date(schedule.date).toISOString().split('T')[0];
      return (
        <div className="schedule__content">
          <h2 className="schedule__content--title">{title}</h2>
          {schedule?.schedule_on_leagues.map((scheduleOnLeague) => (
            <div className="schedule__content--scheduleOnleague">
              <div className="leagueName">
                <h3>{scheduleOnLeague.league_name}</h3>
              </div>
              {scheduleOnLeague.matchs.map((match) => (
                <div className="match">
                  <div className="time">{match.time}</div>
                  <div className="club1">
                    <p>
                      {match.club_1.name}
                      <img src={match.club_1.logo}></img>
                    </p>
                  </div>
                  <div className="score">{match.scores}</div>
                  <div className="club2">
                    <p>
                      {match.club_2.name}
                      <img src={match.club_2.logo}></img>
                    </p>
                  </div>
                  <div className="detail">
                    <FontAwesomeIcon icon={faChevronRight} />
                  </div>
                </div>
              ))}
            </div>
          ))}
        </div>
      );
    }
    return <div className="schedule__content">Chưa có thông tin</div>;
  }

  return <div className="schedule__content">Chưa có thông tin</div>;
};

export default ScheduleContent;
