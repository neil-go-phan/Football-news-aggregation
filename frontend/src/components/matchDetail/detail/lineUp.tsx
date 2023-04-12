import React, { FunctionComponent, ReactNode } from 'react';
import { MatchDetailTitle, MatchLineup, PitchRows } from './index';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faShirt } from '@fortawesome/free-solid-svg-icons';
import Image from 'next/image';

type Props = {
  matchLineUp: MatchLineup | null;
  matchTitle: MatchDetailTitle | null;
};

const MatchLineUpComponent: FunctionComponent<Props> = ({
  matchLineUp,
  matchTitle,
}) => {
  const renderPitchRow = (pitchRows: Array<PitchRows>): ReactNode => {
    return (
      <>
        {pitchRows.map((row, index) => (
          <div className="pitch-row" key={`pitch-row-${index}`}>{renderPlayer(row)}</div>
        ))}
      </>
    );
  };

  const renderPlayer = (row: PitchRows): ReactNode => {
    return (
      <>
        {row.pitch_rows_detail.map((player) => (
          <div className="pitch-item" key={`pitch-item-${player.player_name}`}>
            <div className="team-player">
              <div className="img">
                <FontAwesomeIcon icon={faShirt} />
                <div className="player-number">{player.player_number}</div>
              </div>
              <div className="name">{player.player_name}</div>
            </div>
          </div>
        ))}
      </>
    );
  };

  if (matchLineUp && matchTitle) {
    return (
      <div id="lineup" className="matchDetail__content--lineup">
        <div className="title">Đội hình</div>
        <div className="team">
          <div className="team-head">
            <div className="team-name">
              <Image
                src={matchTitle.club_1.logo}
                alt={`${matchTitle.club_1.name} logo`}
                className="logo"
              />
              <div className="clubName">{matchTitle.club_1.name}</div>
            </div>
            <div className="formation">
              {matchLineUp.lineup_club_1.formation}
            </div>
          </div>

          <div className="team-body">
            <div className="team-content team1">
              {matchLineUp.lineup_club_1.pitch_row ? (
                renderPitchRow(matchLineUp.lineup_club_1.pitch_row)
              ) : (
                <></>
              )}
            </div>
            <div className="team-content team2">
              {matchLineUp.lineup_club_2.pitch_row ? (
                renderPitchRow(matchLineUp.lineup_club_2.pitch_row)
              ) : (
                <></>
              )}
            </div>
          </div>

          <div className="team-head">
            <div className="team-name">
              <Image
                src={matchTitle.club_2.logo}
                alt={`${matchTitle.club_2.name} logo`}
                className="logo"
              />
              <div className="clubName">{matchTitle.club_2.name}</div>
            </div>
            <div className="formation">
              {matchLineUp.lineup_club_2.formation}
            </div>
          </div>
        </div>
      </div>
    );
  }
  return <div className="matchDetail__content--lineup" id="lineup"></div>;
};

export default MatchLineUpComponent;
