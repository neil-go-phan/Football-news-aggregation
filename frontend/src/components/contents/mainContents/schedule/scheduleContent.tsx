import React, { FunctionComponent } from 'react';
import { Schedules } from '.';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faChevronRight } from '@fortawesome/free-solid-svg-icons';
import Link from 'next/link';
import { _ROUTES } from '@/helpers/constants';
import { formatRoute, formatVietnameseDate } from '@/helpers/format';
import Image from 'next/image';

type Props = {
  schedule: Schedules | undefined;
};

const ScheduleContent: FunctionComponent<Props> = ({ schedule }) => {
  // TODO: refactor this duck-typing logic
  if (schedule) {
    if (schedule.schedule_on_leagues !== null) {
      const title = formatVietnameseDate(new Date(schedule.date));

      return (
        <div className="schedule__content p-3">
          <h2 className="schedule__content--title">{title}</h2>
          {schedule?.schedule_on_leagues.map((scheduleOnLeague) => (
            <div
              key={`schedule__content--scheduleOnleague--${scheduleOnLeague.league_name}`}
              className="schedule__content--scheduleOnleague"
            >
              <div className="leagueName p-2">
                <h3>{scheduleOnLeague.league_name}</h3>
              </div>
              {scheduleOnLeague.matchs.map((match) => (
                <div
                  key={`scheduleOnleague--match--${match.match_detail_link}`}
                  className="match"
                >
                  <div className="timeAndRound">
                    <div className="time">{match.time}</div>
                    <div className="round">{match.round}</div>{' '}
                  </div>
                  <div className="club1">
                    <p>
                      {match.club_1.name}
                      <Image
                        alt="club-logo"
                        src={match.club_1.logo}
                        width={20}
                        height={20}
                      ></Image>
                    </p>
                  </div>
                  <div className="score">
                    <span>{match.scores}</span>
                  </div>
                  <div className="club2">
                    <p>
                      <Image
                        alt="club-logo"
                        src={match.club_2.logo}
                        width={20}
                        height={20}
                      ></Image>
                      {match.club_2.name}
                    </p>
                  </div>
                  <div className="detail">
                    <Link
                    href={{ pathname: `${_ROUTES.MATCH_DETAIL_PAGE}/${formatRoute(
                      match.match_detail_link
                    )}`, query: { date: schedule.date, club_1: match.club_1.name, club_2: match.club_2.name } }}
                    >
                      <FontAwesomeIcon icon={faChevronRight} />
                    </Link>
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
