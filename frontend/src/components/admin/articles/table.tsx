import { ArticleType } from '@/components/matchDetail/relatedNews/article';
import React, { useEffect, useState } from 'react';
import { Table } from 'react-bootstrap';
import { Column, useTable } from 'react-table';
import DeleteArticleBtn from './deleteBtn';
import axiosClient from '@/helpers/axiosClient';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import { toast } from 'react-toastify';
import { ThreeDots } from 'react-loader-spinner';

type ArticlesRender = {
  index: number;
  title: string;
  description: string;
  source: string;
  tags: Array<string>;
  action: boolean;
};
type Props = {
  articles: Array<ArticleType>;
  currentPage: number;
  handleUpdateTable: (title: string, index: number) => void;
};

const ArticleTable: React.FC<Props> = (props: Props) => {
  // const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const columns: Column<ArticlesRender>[] = React.useMemo(
    () => [
      {
        header: 'Index',
        accessor: 'index',
      },
      {
        header: 'Title',
        accessor: 'title',
      },
      {
        header: 'Description',
        accessor: 'description',
      },
      {
        header: 'Source',
        accessor: 'source',
        Cell: ({ row }) => (
          <a target="_blank" className="" href={row.values.source}>
            {getDomainName(row.values.source)}
          </a>
        ),
      },
      {
        header: 'action',
        accessor: 'action',
        Cell: ({ row }) => (
          <DeleteArticleBtn
            title={row.values.title}
            index={row.index}
            handleUpdateTable={props.handleUpdateTable}
            key={`delete-article-btn-${row.values.title}`}
          />
        ),
      },
    ],
    []
  );

  const getDomainName = (url: string): string => {
    let domain = new URL(url);
    return domain.hostname.replace('www.', '');
  };

  const useCreateTableData = (articles: Array<ArticleType> | undefined) => {
    return React.useMemo(() => {
      if (!articles) {
        return [];
      }
      return articles.map((article, index) => ({
        index: index + 1 + 10 * (props.currentPage - 1),
        title: article.title,
        description: article.description,
        source: article.link,
        tags: article.tags,
        action: true,
      }));
    }, [articles]);
  };

  const data = useCreateTableData(props.articles);
  const { getTableProps, getTableBodyProps, headerGroups, prepareRow, rows } =
    useTable({
      columns,
      data,
    });
  if (props.articles.length === 0) {
    return (
      <div className="sidebar--loading">
        <ThreeDots
          height="50"
          width="50"
          radius="9"
          color="#4fa94d"
          ariaLabel="three-dots-loading"
          visible={true}
        />
      </div>
    );
  }
  return (
    <>
      <div className="adminArticles__list--table">
        <Table bordered hover {...getTableProps()}>
          <thead>
            {headerGroups.map((headerGroup, index) => (
              <tr
                {...headerGroup.getHeaderGroupProps()}
                key={`articles-admin-collum-${index}`}
              >
                {headerGroup.headers.map((column) => (
                  <th
                    {...column.getHeaderProps()}
                    key={`articles-admin-collum-${column.render('header')}}`}
                  >
                    {column.render('header')}
                  </th>
                ))}
              </tr>
            ))}
          </thead>
          <tbody {...getTableBodyProps()}>
            {rows.map((row, i) => {
              prepareRow(row);
              return (
                <tr
                  {...row.getRowProps()}
                  key={`articles-admin-tr-${i}-${row.values.title}`}
                >
                  {row.cells.map((cell, i) => {
                    return (
                      <td
                        {...cell.getCellProps()}
                        key={`articles-admin-td-${i}-${row.values.title}`}
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
      </div>
    </>
  );
  // return (
  //   <>
  //     <div className="adminArticles__list--table">
  //       <Table bordered hover {...getTableProps()}>
  //         <thead>
  //           <tr>
  //             <td>Index</td>
  //             <td>Title</td>
  //             <td>Description</td>
  //             <td>Source</td>
  //             <td>Action</td>
  //           </tr>
  //         </thead>
  //         <tbody>
  //           {props.articles.map((article, index) => (
  //             <tr>
  //               <td>{index + 1}</td>
  //               <td>{article.title}</td>
  //               <td>{article.description}</td>
  //               <td>{getDomainName(article.link)}</td>
  //               <td>
  //                 {' '}
  //                 <DeleteArticleBtn
  //                   title={article.title}
  //                   index={index + 1}
  //                   handleUpdateTable={props.handleUpdateTable}
  //                   key={`delete-article-btn-${article.title}`}
  //                 />
  //               </td>
  //             </tr>
  //           ))}
  //         </tbody>
  //       </Table>
  //     </div>
  //   </>
  // );
};

export default ArticleTable;
