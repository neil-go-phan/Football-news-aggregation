import React, { FunctionComponent, ReactNode } from 'react';
import { MatchDetailTitle, MatchLineup, PitchRows } from './index';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faShirt } from '@fortawesome/free-solid-svg-icons';
import Image from 'next/image';

type Props = {
  matchLineUp: MatchLineup | null;
  matchTitle: MatchDetailTitle | null;
};
// pixel
const CLUB_LOGO_LINEUP_SIDE = 24;

const MatchLineUpComponent: FunctionComponent<Props> = ({
  matchLineUp,
  matchTitle,
}) => {
  const renderPitchRow = (
    pitchRows: Array<PitchRows>,
    // eslint-disable-next-line camelcase
    shirt_color: string
  ): ReactNode => {
    return (
      <>
        {pitchRows.map((row, index) => (
          <div className="pitch-row" key={`pitch-row-${index}`}>
            {/*eslint-disable-next-line camelcase */}
            {renderPlayer(row, shirt_color)}
          </div>
        ))}
      </>
    );
  };
  // eslint-disable-next-line camelcase
  const renderPlayer = (row: PitchRows, shirt_color: string): ReactNode => {
    return (
      <>
        {row.pitch_rows_detail.map((player) => (
          <div className="pitch-item" key={`pitch-item-${player.player_name}`}>
            <div className="team-player">
              <div className="img">
                <FontAwesomeIcon
                  icon={faShirt}
                  // eslint-disable-next-line camelcase
                  style={{ color: shirt_color }}
                />
                <div className="player-number">{player.player_number}</div>
              </div>
              <div className="name">{player.player_name}</div>
            </div>
          </div>
        ))}
      </>
    );
  };

  if (
    matchLineUp &&
    matchTitle &&
    matchLineUp.lineup_club_1.pitch_row &&
    matchLineUp.lineup_club_2.pitch_row
  ) {
    return (
      <div id="lineup" className="matchDetail__content--lineup">
        <div className="title">Lineup</div>
        <div className="team">
          <div className="team-head">
            <div className="team-name">
              <Image
                src={matchTitle.club_1.logo}
                alt={`${matchTitle.club_1.name} logo`}
                width={CLUB_LOGO_LINEUP_SIDE}
                height={CLUB_LOGO_LINEUP_SIDE}
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
                renderPitchRow(
                  matchLineUp.lineup_club_1.pitch_row,
                  matchLineUp.lineup_club_1.shirt_color
                )
              ) : (
                <></>
              )}
            </div>
            <div className="team-content team2">
              {matchLineUp.lineup_club_2.pitch_row ? (
                renderPitchRow(
                  matchLineUp.lineup_club_2.pitch_row,
                  matchLineUp.lineup_club_2.shirt_color
                )
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
                width={CLUB_LOGO_LINEUP_SIDE}
                height={CLUB_LOGO_LINEUP_SIDE}
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
