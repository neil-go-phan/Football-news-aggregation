import React, { FunctionComponent, ReactNode } from 'react';
import { MatchOverview, OverviewItem } from './index';

type Props = {
  matchOverview: MatchOverview | null;
};

const MatchOverviewComponent: FunctionComponent<Props> = ({ matchOverview }) => {
  const checkImage = (imgType: string): string => {
    switch (imgType) {
      case 'goal':
        return '/images/goal.png';
      case 'yellow-card':
        return '/images/yellow_card.png';
      case 'red-card':
        return '/images/red_card.png';
      case 'substitution':
        return '/images/substitution.png';
      default:
        return '';
    }
  };

  const renderClubOverview = (
    overviewItems: Array<OverviewItem>
  ): ReactNode => {
    return (
      <>
        {overviewItems.map((item) => {
          if (item.info === "") {
            return;
          }
          return (
            <div className="item" key={item.info}>
              {item.info}
              <span className="more-info">
                <img
                  src={checkImage(item.image_type)}
                  alt={item.image_type}
                  className="image"
                />{' '}
                <span className="time">{item.time}</span>
              </span>
            </div>
          );
        })}
      </>
    );
  };
  if (matchOverview) {
    return (
      <div id="overview" className="matchDetail__content--overview d-flex">
        <div className="club col-6">
          {matchOverview.club_1_overview ? (
            renderClubOverview(matchOverview.club_1_overview)
          ) : (
            <></>
          )}
        </div>
        <div className="club col-6">
          {matchOverview.club_2_overview ? (
            renderClubOverview(matchOverview.club_2_overview)
          ) : (
            <></>
          )}
        </div>
      </div>
    );
  }
  return <div className="matchDetail__content--overview" id="overview"></div>;
};

export default MatchOverviewComponent;
