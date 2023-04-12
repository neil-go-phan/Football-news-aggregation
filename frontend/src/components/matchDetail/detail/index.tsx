import { Club } from '@/components/contents/mainContents/schedule';
import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import MatchTitle from './title';
import MatchDetail from '..';
import MatchLineUpComponent from './lineUp';
import MatchOverviewComponent from './overview';
import MatchStatsComponent from './statistics';
import MatchEventsComponent from './events';

export type MatchDetail = {
  match_detail_title: MatchDetailTitle;
  match_overview: MatchOverview;
  match_statistics: MatchStatistics;
  match_lineup: MatchLineup;
  match_progress: MatchProgress;
};

export type MatchDetailTitle = {
  club_1: Club;
  club_2: Club;
  match_score: string;
};

export type MatchOverview = {
  club_1_overview: Array<OverviewItem>;
  club_2_overview: Array<OverviewItem>;
};

export type OverviewItem = {
  info: string;
  image_type: string;
  time: string;
};

export type MatchStatistics = {
  statistics: Array<StatisticsItem>;
};

export type StatisticsItem = {
  stat_club_1: string;
  stat_content: string;
  stat_club_2: string;
};

export type MatchProgress = {
  events: Array<MatchEvent>;
};
export type MatchEvent = {
  time: string;
  content: string;
};

export type MatchLineup = {
  lineup_club_1: MatchLineUpDetail;
  lineup_club_2: MatchLineUpDetail;
};

export type MatchLineUpDetail = {
  club_name: string;
  formation: string;
  pitch_row: Array<PitchRows>;
};

export type PitchRows = {
  pitch_rows_detail: Array<PitchRowsDetail>;
};

export type PitchRowsDetail = {
  player_name: string;
  player_number: string;
};

function Detail() {
  const router = useRouter();
  const { date, club_1, club_2 } = router.query;
  const [matchDetail, setMatchDetail] = useState<MatchDetail>();
  console.log(matchDetail);

  useEffect(() => {
    const requestGetMatchDetail = async () => {
      try {
        const { data } = await axiosClient.get('match-detail/get', {
          // eslint-disable-next-line camelcase
          params: { date: date, club_1: club_1, club_2: club_2 },
        });
        setMatchDetail(data.match_detail);
      } catch (error) {
        toast.error(`Error occurred while get match detail`, {
          position: 'top-right',
          autoClose: 3000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        });
      }
    };
    requestGetMatchDetail();
  }, [router.asPath]);
  if (matchDetail) {
    return (
      <div className="matchDetail__content px-5">
        <MatchTitle matchTitle={matchDetail.match_detail_title} date={date} />

        <div className="matchDetail__content--navbar d-flex">
          <div className="nav--item col-3">
            {matchDetail.match_overview.club_1_overview ||
            matchDetail.match_overview.club_2_overview ? (
              <a href="#overview" className="item">
                Tổng quan
              </a>
            ) : (
              <span className="item hidden">Tổng quan</span>
            )}
          </div>
          <div className="nav--item col-3">
            {matchDetail.match_statistics.statistics ? (
              <a href="#statistics" className="item">
                Thống kê
              </a>
            ) : (
              <span className="item hidden">Thống kê</span>
            )}
          </div>
          <div className="nav--item col-3">
            {matchDetail.match_lineup.lineup_club_1.pitch_row &&
            matchDetail.match_lineup.lineup_club_2.pitch_row ? (
              <a href="#lineup" className="item">
                Đội hình
              </a>
            ) : (
              <span className="item hidden">Đội hình</span>
            )}
          </div>
          <div className="nav--item col-3">
            {matchDetail.match_progress.events ? (
              <a href="#process" className="item">
                Diễn biến
              </a>
            ) : (
              <span className="item hidden">Diễn biến</span>
            )}
          </div>
        </div>

        <MatchOverviewComponent matchOverview={matchDetail.match_overview}/>

        <MatchStatsComponent matchStatistics={matchDetail.match_statistics} matchTitle={matchDetail.match_detail_title}/>

        <MatchLineUpComponent matchLineUp={matchDetail.match_lineup} matchTitle={matchDetail.match_detail_title}/>

        <MatchEventsComponent matchProcess={matchDetail.match_progress}/>
      </div>
    );
  } else {
    return <div className="matchDetail__content">Loading</div>;
  }
}

export default Detail;
