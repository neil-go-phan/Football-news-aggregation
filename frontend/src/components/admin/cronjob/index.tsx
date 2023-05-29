import React, { useEffect, useState } from 'react';
import { Column, usePagination, useTable } from 'react-table';
import CronjobActions from './action';
import { useRouter } from 'next/router';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { toast } from 'react-toastify';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import { Button, Table } from 'react-bootstrap';

type CronjobRow = {
  index: number;
  name: string;
  url: string;
  run_every_min: number;
  action: string;
};

type Cronjob = {
  name: string;
  url: string;
  run_every_min: number;
};

function AdminCronjob() {
  const [cronjobs, setCronjobs] = useState<Array<Cronjob>>();
  const router = useRouter();

  const columns: Column<CronjobRow>[] = React.useMemo(
    () => [
      {
        Header: 'STT',
        accessor: 'index',
      },
      {
        Header: 'Name',
        accessor: 'name',
      },
      {
        Header: 'Run every (min)',
        accessor: 'run_every_min',
      },
      {
        Header: 'Action',
        accessor: 'action',
        Cell: ({ row }) => (
          <CronjobActions
            url={row.values.action}
            name={row.values.name}
            runEveryMinOld={row.values.run_every_min}
            handleChangeTime={handleChangeTime}
            // handleUpdate={handleUpdate}
          />
        ),
      },
    ],
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const useCreateTableData = (cronjobs: Array<Cronjob> | undefined) => {
    return React.useMemo(() => {
      if (!cronjobs) {
        return [];
      }
      return cronjobs.map((cronjob, index) => ({
        index: index + 1,
        name: cronjob.name,
        url: cronjob.url,
        run_every_min: cronjob.run_every_min,
        action: cronjob.url,
      }));
    }, [cronjobs]);
  };

  const data = useCreateTableData(cronjobs);

  const {
    getTableProps,
    getTableBodyProps,
    headerGroups,
    prepareRow,
    page,
    pageOptions,
    state: { pageIndex },
    previousPage,
    nextPage,
    canPreviousPage,
    canNextPage,
  } = useTable(
    {
      columns,
      data,
      initialState: { pageIndex: 0 },
    },
    usePagination
  );

  const handleChangeTime = () => {
    requestListCronjobs()
  };

  const requestListCronjobs = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('cronjob/list-cronjob');
      setCronjobs(data.cronjobs);
    } catch (error) {
      toast.error('Error occurred while get list cronjob', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    }
  };

  useEffect(() => {
    requestListCronjobs();
  }, [router.asPath]);

  return (
    <div className="adminCronjob">
      <h2 className="adminCronjob__list--title">Crawler list</h2>
      <div className="adminCronjob__list--table mt-3">
        <Table bordered hover {...getTableProps()}>
          <thead>
            {headerGroups.map((headerGroup, index) => (
              <tr
                {...headerGroup.getHeaderGroupProps()}
                key={`crawler-admin-tr-${index}`}
              >
                {headerGroup.headers.map((column, index) => (
                  <th
                    {...column.getHeaderProps()}
                    key={`crawler-admin-tr-item-${index}`}
                  >
                    {column.render('Header')}
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
                  key={`crawler-admin-row-tr-${row.index}`}
                >
                  {row.cells.map((cell, index) => {
                    return (
                      <td
                        {...cell.getCellProps()}
                        key={`crawler-admin-row-tr-item-${index}`}
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
  );
}

export default AdminCronjob;
