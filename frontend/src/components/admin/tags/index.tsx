import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import React, { useEffect, useState } from 'react';
import { Button, Form, InputGroup, Table } from 'react-bootstrap';
import { toast } from 'react-toastify';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass, faPlus } from '@fortawesome/free-solid-svg-icons';
import { Column, useGlobalFilter, usePagination, useTable } from 'react-table';
import { Leagues } from '../leagues';
import { toLowerCaseNonAccentVietnamese } from '@/helpers/format';
import { useRouter } from 'next/router';
import DeleteBtn from './deleteBtn';
import { ERROR_POPUP_ADMIN_TIME, ERROR_POPUP_USER_TIME } from '@/helpers/constants';

const TIN_TUC_BONG_DA_TAG = 'tin tuc bong da';

export type Tags = Array<string>;

type TagsRow = {
  index: number;
  tagName: string;
  isDisabled: boolean;
};

export default function AdminTags() {
  const [tags, setTags] = useState<Tags>();
  const [leagues, setLeagues] = useState<Leagues>();
  const [newTagName, setNewTagName] = useState<string>('');
  const router = useRouter();

  const columns: Column<TagsRow>[] = React.useMemo(
    () => [
      {
        Header: 'STT',
        accessor: 'index',
      },
      {
        Header: 'Tên tag',
        accessor: 'tagName',
      },
      {
        Header: 'Action',
        accessor: 'isDisabled',
        Cell: ({ row }) => (
          <DeleteBtn
            isDisabled={row.values.isDisabled}
            tagName={row.values.tagName}
            handleDeleteTag={handleDeleteTag}
          />
        ),
      },
    ],
    []
  );

  const handleDeleteTag = () => {
    requestListTags();
  };

  const handleAddTag = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const tagFormated = toLowerCaseNonAccentVietnamese(newTagName)
    const exist = tags!.indexOf(tagFormated);
    if (exist !== -1) {
      toast.info('Tag alreay exist', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      setNewTagName('');
      // return
    }
    if (tagFormated !== '') {
      requestAddTag(tagFormated);
    }
    setNewTagName('');
  };

  const removeDefaultTag = (tags: Tags) => {
    if (tags) {
      tags.every((tag, index) => {
        if (tag === TIN_TUC_BONG_DA_TAG) {
          tags.splice(index, 1);
          return false;
        }
      });
    }
    return tags;
  };

  const checkIsTagLeague = (tagName: string): boolean => {
    for (let index = 0; index < leagues!.leagues.length; index++) {
      if (
        toLowerCaseNonAccentVietnamese(leagues!.leagues[index].league_name) ===
        tagName
      ) {
        return true;
      }
    }
    return false;
  };

  const useCreateTableData = (tagRows: Tags | undefined) => {
    const tagRowsAfter = removeDefaultTag(tagRows!);
    return React.useMemo(() => {
      if (!tagRowsAfter) {
        return [];
      }
      return tagRowsAfter.map((tag, index) => ({
        index: index + 1,
        tagName: tag,
        isDisabled: checkIsTagLeague(tag),
      }));
    }, [tagRowsAfter]);
  };

  const data = useCreateTableData(tags);

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

  const requestAddTag = async (tag: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('tags/add', {
        params: { tag: tag },
      });
      if (!data.success) {
        throw 'add fail';
      }
      toast.success('Add tag success', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      if (tags) {
        const newTags = tags;
        newTags.push(tag);
        setTags([...newTags]);
      }
      requestTaggedArticle(tag)
    } catch (error) {
      toast.error('Error occurred while add tags', {
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
  const requestListTags = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('tags/list', {});
      setTags(data.tags.tags);
    } catch (error) {
      toast.error('Error occurred while get tags', {
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
  const requestTaggedArticle = async (tag: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('article/update-tag', {
        params: { tag: tag },
      });
      if (!data.success) {
        throw 'tagged fail';
      }
      toast.success('Tagged article success', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    } catch (error) {
      toast.error('Error occurred while tagged article', {
        position: 'top-right',
        autoClose: ERROR_POPUP_USER_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    }
  };
  const requestListLeagues = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('leagues/list-all', {});

      setLeagues(data.leagues);
    } catch (error) {
      toast.error('Error occurred while get leagues', {
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
    requestListLeagues();
    requestListTags();
  }, [router.asPath]);

  if (tags) {
    return (
      <div className="adminTags">
        <h1 className="adminTags__title">Manage tags</h1>
        <div className="adminTags__overview">
          <div className="adminTags__overview--item">
            <p>
              Tổng số tags: <span>{tags.length}</span>
            </p>
          </div>
        </div>
        <div className="adminTags__list">
          <h2 className="adminTags__list--title">Danh sách tags</h2>
          <div className="adminTags__list--search d-sm-flex">
            <div className="searchBar col-sm-6">
              <InputGroup className="mb-3">
                <InputGroup.Text>
                  <FontAwesomeIcon icon={faMagnifyingGlass} fixedWidth />
                </InputGroup.Text>
                <Form.Control
                  placeholder="Search tags"
                  type="text"
                  value={globalFilter || ''}
                  onChange={(e) => setGlobalFilter(e.target.value)}
                />
              </InputGroup>
            </div>
            <div className="col-sm-1"></div>
            <div className="addBtn col-sm-5">
              <form onSubmit={handleAddTag}>
                <InputGroup className="mb-3">
                  <InputGroup.Text>
                    <FontAwesomeIcon icon={faPlus} fixedWidth />
                  </InputGroup.Text>
                  <Form.Control
                    placeholder="Input tag name"
                    type="text"
                    value={newTagName}
                    onChange={(e) => setNewTagName(e.target.value)}
                  />
                  <Button type="submit" variant="success">
                    Add
                  </Button>
                </InputGroup>
              </form>
            </div>
          </div>

          <div className="adminTags__list--table">
            <Table bordered hover {...getTableProps()}>
              <thead>
                {headerGroups.map((headerGroup, index) => (
                  <tr {...headerGroup.getHeaderGroupProps()} key={`tags-admin-tr-${index}`}>
                    {headerGroup.headers.map((column, index) => (
                      <th {...column.getHeaderProps()} key={`tags-admin-tr-item-${index}`}>
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
                    <tr {...row.getRowProps()} key={`tags-admin-row-tr-${row.index}`}>
                      {row.cells.map((cell, index) => {
                        return (
                          <td {...cell.getCellProps()} key={`tags-admin-row-tr-item-${index}`}>
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
    <div className="adminTags">
      <h1 className="adminTags__title">Manage tags</h1>
      <div className="adminTags__overview">
        <div className="adminTags__overview--item">
          <p>
            Total tags: <span></span>
          </p>
        </div>
      </div>
      <div className="adminTags__list">
        <h2 className="adminTags__list--title">Tags list</h2>
        <div className="adminTags__list--search">
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
        <div className="adminTags__list--table">
          <Table bordered hover>
            <thead>
              <tr>
                <th>STT</th>
                <th>Tag name</th>
                <th>Action</th>
              </tr>
            </thead>
            <tbody></tbody>
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
