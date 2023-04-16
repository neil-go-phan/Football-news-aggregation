import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Table } from 'react-bootstrap';
import { toast } from 'react-toastify';
import Status from './status';

type Leagues = {
  leagues: Array<League>;
};

type League = {
  league_name: string;
  active: boolean;
};

const TIN_TUC_BONG_DA_LEAGUE = 1

export default function AdminLeagues() {
  const [leagues, setLeagues] = useState<Leagues>();

  const getActiveLeagues = (): number => {
    let count = 0;
    leagues?.leagues.forEach((league) => {
      if (league.active) {
        count += 1;
      }
    });
    return count - TIN_TUC_BONG_DA_LEAGUE;
  };

  const activeLeague = getActiveLeagues();

  useEffect(() => {
    const requestListLeagues = async () => {
      try {
        const { data } = await axiosProtectedAPI.get('leagues/list-all', {});
        setLeagues(data.leagues);
      } catch (error) {
        toast.error('Error occurred while get leagues', {
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
    requestListLeagues();
  }, []);

  if (leagues) {
    return (
      <div className="adminLeagues">
        <h1 className="adminLeagues__title">Manage Leagues</h1>
        <div className="adminLeagues__overview">
          <div className="adminLeagues__overview--item">
            <p>
              Tổng số giải đấu: <span>{leagues.leagues.length - TIN_TUC_BONG_DA_LEAGUE}</span>
            </p>
          </div>
          <div className="adminLeagues__overview--item">
            <p>
              Số giải đấu hiển thị: <span>{activeLeague}</span>
            </p>
          </div>
          <div className="adminLeagues__overview--item">
            <p>
              Số giải đấu không hiển thị:{' '}
              <span>{leagues.leagues.length - TIN_TUC_BONG_DA_LEAGUE - activeLeague}</span>
            </p>
          </div>
        </div>
        <div className="adminLeagues__list">
          <h2 className="adminLeagues__list--title">Danh sách giải đấu</h2>
          <div className="adminLeagues__list--table">
            <Table bordered hover>
              <thead>
                <tr>
                  <th>#</th>
                  <th>Tên giải đấu</th>
                  <th>Trạng thái</th>
                </tr>
              </thead>
              <tbody>
                {leagues.leagues.map((league, index) => {
                  if (league.league_name !== 'Tin tức bóng đá') {
                    return (
                      <tr key={`league_name ${league.league_name}`}>
                        <td>{index}</td>
                        <td>{league.league_name}</td>
                        <td>
                          {
                            <Status
                              active={league.active}
                              leagueName={league.league_name}
                            />
                          }
                        </td>
                      </tr>
                    );
                  }
                })}
              </tbody>
            </Table>
          </div>
        </div>
      </div>
    );
  }
  return (
    <div className="adminLeagues">
      <h1 className="adminLeagues__title">Manage Leagues</h1>
      <div className="adminLeagues__overview">
        <div className="adminLeagues__overview--item">
          <p>
            Tổng số giải đấu: <span></span>
          </p>
        </div>
        <div className="adminLeagues__overview--item">
          <p>
            Số giải đấu hiển thị: <span></span>
          </p>
        </div>
        <div className="adminLeagues__overview--item">
          <p>
            Số giải đấu không hiển thị: <span></span>
          </p>
        </div>
      </div>
      <div className="adminLeagues__list">
        <h2 className="adminLeagues__list--title">Danh sách giải đấu</h2>
        <div className="adminLeagues__list--table">
          <Table bordered hover>
            <thead>
              <tr>
                <th>Tên giải đấu</th>
                <th>Số trận đấu</th>
                <th>Trạng thái</th>
              </tr>
            </thead>
            <tbody></tbody>
          </Table>
        </div>
      </div>
    </div>
  );
}
