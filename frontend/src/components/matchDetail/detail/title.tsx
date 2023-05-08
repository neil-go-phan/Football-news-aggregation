import React, { FunctionComponent } from 'react';
import { MatchDetailTitle } from './index';
import Image from 'next/image';
type Props = {
  matchTitle: MatchDetailTitle | null;
  date: string | string[] | undefined;
};
//pixel
const CLUB_LOGO_TITLE_SIZE = 80;

const MatchTitle: FunctionComponent<Props> = ({ matchTitle, date }) => {
  if (matchTitle) {
    return (
      <div className="matchDetail__content--title">
        <h1>{`Live match results on ${matchTitle.club_1.name} vs ${matchTitle.club_2.name} ngày ${date}`}</h1>
        <div className="result d-flex">
          <div className="col-4">
            <div className="club1 d-none d-xl-block">
              <Image
                src={
                  matchTitle.club_1.logo === '/images/CLB-MacDinh.svg'
                    ? 'https://bongda24h.vn/images/CLB-MacDinh.svg'
                    : matchTitle.club_1.logo
                }
                alt={`${matchTitle.club_1.name} logo`}
                width={CLUB_LOGO_TITLE_SIZE}
                height={CLUB_LOGO_TITLE_SIZE}
                className="logo logo1"
              />
              <span className="clubName clubName1">
                {matchTitle.club_1.name}
              </span>
            </div>
            <div className="club1 clubsm d-block d-xl-none">
              <Image
                src={
                  matchTitle.club_1.logo === '/images/CLB-MacDinh.svg'
                    ? 'https://bongda24h.vn/images/CLB-MacDinh.svg'
                    : matchTitle.club_1.logo
                }
                alt={`${matchTitle.club_1.name} logo`}
                width={CLUB_LOGO_TITLE_SIZE}
                height={CLUB_LOGO_TITLE_SIZE}
                className="logo logo1"
              />
              <span className="clubName clubName1">
                {matchTitle.club_1.name}
              </span>
            </div>
          </div>
          <div className="col-4 match-score">{matchTitle.match_score}</div>
          <div className="col-4">
            <div className="club2 d-none d-xl-block">
              <span className="clubName clubName2">
                {matchTitle.club_2.name}
              </span>
              <Image
                src={
                  matchTitle.club_2.logo === '/images/CLB-MacDinh.svg'
                    ? 'https://bongda24h.vn/images/CLB-MacDinh.svg'
                    : matchTitle.club_2.logo
                }
                alt={`${matchTitle.club_2.name} logo`}
                width={CLUB_LOGO_TITLE_SIZE}
                height={CLUB_LOGO_TITLE_SIZE}
                className="logo logo2"
              />
            </div>
            <div className="club2 clubsm d-block d-xl-none">
              <Image
                src={
                  matchTitle.club_2.logo === '/images/CLB-MacDinh.svg'
                    ? 'https://bongda24h.vn/images/CLB-MacDinh.svg'
                    : matchTitle.club_2.logo
                }
                alt={`${matchTitle.club_2.name} logo`}
                width={CLUB_LOGO_TITLE_SIZE}
                height={CLUB_LOGO_TITLE_SIZE}
                className="logo logo2"
              />
              <span className="clubName clubName2">
                {matchTitle.club_2.name}
              </span>
            </div>
          </div>
        </div>
      </div>
    );
  }
  return (
    <div className="matchDetail__content--title">
      <h1>{`Live match results on ${date}`}</h1>
    </div>
  );
};

export default MatchTitle;
