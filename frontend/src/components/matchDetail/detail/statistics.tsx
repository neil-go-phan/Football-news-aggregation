import React, { FunctionComponent, ReactNode } from 'react';
import { MatchDetailTitle, MatchStatistics, StatisticsItem } from './index';
import Image from 'next/image';

type Props = {
  matchStatistics: MatchStatistics | null;
  matchTitle: MatchDetailTitle | null;
};
// pixel
const CLUB_LOGO_STATS_SIZE = 40 

const MatchStatsComponent: FunctionComponent<Props> = ({
  matchStatistics,
  matchTitle,
}) => {
  const renderStatsItem = (statsItem: Array<StatisticsItem>): ReactNode => {
    const isStat1Bigger = (stat1: string, stat2: string): boolean => {
      const stat1Number = +stat1;
      const stat2Number = +stat2;
      if (stat1Number >= stat2Number) {
        return true;
      }
      return false;
    };
    return (
      <div className="detail">
        {statsItem.map((item) => (
          <div
            className="stats-item d-flex"
            key={`stats-${item.stat_club_1}-${item.stat_club_2}-${item.stat_content}`}
          >
            <div className="name club1 col-3">
              <span
                className={
                  isStat1Bigger(item.stat_club_1, item.stat_club_2)
                    ? 'active'
                    : ''
                }
              >
                {item.stat_club_1}
              </span>
            </div>
            <div className="statContent col-6">{item.stat_content}</div>
            <div className="name club2 col-3">
              <span
                className={
                  isStat1Bigger(item.stat_club_1, item.stat_club_2)
                    ? ''
                    : 'active'
                }
              >
                {item.stat_club_2}
              </span>
            </div>
          </div>
        ))}
      </div>
    );
  };
  if (matchStatistics && matchTitle && matchStatistics.statistics) {
    return (
      <div id="statistics" className="matchDetail__content--statistics">
        <div className="title">Statistics</div>
        <div className="statHead d-flex">
          <div className="club col-6">
            <div className="club1">
              <Image
                src={matchTitle.club_1.logo}
                alt={`${matchTitle.club_1.name} logo`}
                width={CLUB_LOGO_STATS_SIZE}
                height={CLUB_LOGO_STATS_SIZE}
                className="logo logo1"
              />
              <span className="clubName clubName1">
                {matchTitle.club_1.name}
              </span>
            </div>
          </div>
          <div className="club col-6">
            <div className="club2">
              <span className="clubName clubName2">
                {matchTitle.club_2.name}
              </span>
              <Image
                src={matchTitle.club_2.logo}
                alt={`${matchTitle.club_2.name} logo`}
                width={CLUB_LOGO_STATS_SIZE}
                height={CLUB_LOGO_STATS_SIZE}
                className="logo logo2"
              />
            </div>
          </div>
        </div>
        {matchStatistics.statistics ? (
          renderStatsItem(matchStatistics.statistics)
        ) : (
          <></>
        )}
      </div>
    );
  }
  return (
    <div className="" id="statistics"></div>
  );
};

export default MatchStatsComponent;
