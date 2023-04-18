import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button, Form, InputGroup, Table } from 'react-bootstrap';
import { toast } from 'react-toastify';
import Status from './status';
import { Column, useGlobalFilter, usePagination, useTable } from 'react-table';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
export type Leagues = {
  leagues: Array<League>;
};

export type League = {
  league_name: string;
  active: boolean;
};

type LeagueRender = {
  index: number;
  leagueName: string;
  active: boolean;
};

// const TIN_TUC_BONG_DA: League = {
//   league_name: 'Tin tức bóng đá',
//   active: true,
// };

const TIN_TUC_BONG_DA = 'Tin tức bóng đá';

export default function AdminLeagues() {
  const [leagues, setLeagues] = useState<Leagues>();

  const getActiveLeagues = (): number => {
    let count = 0;
    leagues?.leagues.forEach((league) => {
      if (league.active) {
        count += 1;
      }
    });
    return count;
  };

  const activeLeague = getActiveLeagues();
  const columns: Column<LeagueRender>[] = React.useMemo(
    () => [
      {
        header: 'STT',
        accessor: 'index',
      },
      {
        header: 'Tên giải đấu',
        accessor: 'leagueName',
      },
      {
        header: 'Trạng thái',
        accessor: 'active',
        Cell: ({ row }) => (
          <Status
            active={row.values.active}
            leagueName={row.values.leagueName}
          />
        ),
      },
    ],
    []
  );

  const removeDefaultLeague = (leagues: Leagues | undefined) => {
    if (leagues) {
      leagues.leagues.every((league, index) => {
        if (league.league_name === TIN_TUC_BONG_DA) {
          leagues.leagues.splice(index, 1);
          return false;
        }
      });
    }
    return leagues;
  };

  const useCreateTableData = (leagues: Leagues | undefined) => {
    const leagueAfter = removeDefaultLeague(leagues);
    return React.useMemo(() => {
      if (!leagueAfter) {
        return [];
      }
      return leagueAfter.leagues.map((league, index) => ({
        index: index + 1,
        leagueName: league.league_name,
        active: league.active,
      }));
    }, [leagueAfter]);
  };

  const data = useCreateTableData(leagues);

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    prepareRow,
    page,
    pageOptions,
    state: { pageIndex, globalFilter },
    previousPage,
    nextPage,
    canPreviousPage,
    canNextPage,
    setGlobalFilter,
  } = useTable(
    {
      columns,
      data,
      initialState: { pageIndex: 0 },
    },
    useGlobalFilter,
    usePagination
  );

  useEffect(() => {
    const requestListLeagues = async () => {
      try {
        const { data } = await axiosProtectedAPI.get('leagues/list-all', {});

        setLeagues(data.leagues);
      } catch (error) {
        toast.error('Error occurred while get leagues', {
          position: 'top-right',
          autoClose: 1000,
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
              Tổng số giải đấu: <span>{leagues.leagues.length}</span>
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
              <span>{leagues.leagues.length - activeLeague}</span>
            </p>
          </div>
        </div>
        <div className="adminLeagues__list">
          <h2 className="adminLeagues__list--title">Danh sách giải đấu</h2>
          <div className="adminLeagues__list--search">
            <InputGroup className="mb-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faMagnifyingGlass} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Search league"
                type="text"
                value={globalFilter || ''}
                onChange={(e) => setGlobalFilter(e.target.value)}
              />
            </InputGroup>
          </div>
          <div className="adminLeagues__list--table">
            <Table bordered hover {...getTableProps()}>
              <thead>
                {headerGroups.map((headerGroup, index) => (
                  <tr
                    {...headerGroup.getHeaderGroupProps()}
                    key={`league-admin-collum-${index}`}
                  >
                    {headerGroup.headers.map((column) => (
                      <th
                        {...column.getHeaderProps()}
                        key={`league-admin-collum-${column.render('header')}}`}
                      >
                        {column.render('header')}
                      </th>
                    ))}
                  </tr>
                ))}
              </thead>
              <tbody {...getTableBodyProps()}>
                {page.map((row) => {
                  prepareRow(row);
                  return (
                    <tr
                      {...row.getRowProps()}
                      key={`league-admin-row-${row.index}`}
                    >
                      {row.cells.map((cell) => {
                        return (
                          <td
                            {...cell.getCellProps()}
                            key={`league-admin-row-item-${cell}`}
                          >
                            {cell.render('Cell')}
                          </td>
                        );
                      })}
                    </tr>
                  );
                })}
              </tbody>
            </Table>
            <div className="btnPaging">
              <Button
                onClick={() => previousPage()}
                disabled={!canPreviousPage}
                variant="primary"
              >
                Previous Page
              </Button>
              <Button
                onClick={() => nextPage()}
                disabled={!canNextPage}
                variant="primary"
              >
                Next Page
              </Button>
              <p>
                Page
                <span>
                  {pageIndex + 1} of {pageOptions.length}
                </span>
              </p>
            </div>
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
                <th>STT</th>
                <th>Tên giải đấu</th>
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
