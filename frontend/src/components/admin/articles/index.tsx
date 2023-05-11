import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import { faMagnifyingGlass } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Form, InputGroup } from 'react-bootstrap';
import ArticleTable from './table';
import axiosClient from '@/helpers/axiosClient';
import { ArticleType } from '@/components/matchDetail/relatedNews/article';
import { ThreeDots } from 'react-loader-spinner';
import AdminArticlePagination from './pagination';

// all article is tagged with this tag. we use it to say to the server that we want to get all article
const DEFAULT_TAGS = 'tin tuc bong da';
export const ROW_PER_PAGE = 10;

function ArticleAdmin() {
  const [today, setToday] = useState<number>();
  const [total, setTotal] = useState<number>();
  const [totalArticleSearch, setTotalArticleSearch] = useState<number>();
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const [keyword, setkeyword] = useState<string>('');

  const requestArticleCount = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('article/count', {});
      setTotal(data.total);
      setToday(data.today);
      setTotalArticleSearch(data.total);
    } catch (error) {
      toast.error('Error occurred while request to count article', {
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
  const requestDeleteArticle = async (id: number) => {
    try {
      const { data } = await axiosProtectedAPI.post('article/delete', {
        id: id,
      });
      if (!data.success) {
        throw 'Throw error occurred while delete article';
      } else {
        toast.success('Delete success', {
          position: 'top-right',
          autoClose: ERROR_POPUP_ADMIN_TIME,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        });
        setArticles([])
        setTimeout(() => {
          //this is time to wait for elastic search delete document. 
          //TODO: mirage database to mongodb
          requestArticleCount();
          requestArticle('', 0);
          setCurrentPage(1)
          pageChangeHandler(1)
        }, 1000);

      }
    } catch (error) {
      toast.error('Error occurred while delete article', {
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

  const requestArticle = async (searchKeyword: string, from: number) => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: searchKeyword.trim(), tags: DEFAULT_TAGS, from: from },
      });
      setArticles(data.articles);
      setTotalArticleSearch(data.total);
    } catch (error) {
      toast.error(
        `Error occurred while searching keyword ${searchKeyword.trim()}`,
        {
          position: 'top-right',
          autoClose: ERROR_POPUP_ADMIN_TIME,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        }
      );
    }
  };

  useEffect(() => {
    requestArticleCount();
    requestArticle('', 0);
  }, []);


  const pageChangeHandler = (currentPage: number) => {
    setCurrentPage(currentPage);
    const skipFactor = (currentPage - 1) * ROW_PER_PAGE;
    requestArticle(keyword, skipFactor);
  };

  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (keyword.trim() === '') {
      setkeyword('');
      requestArticle('', 0);
      setCurrentPage(1)
      pageChangeHandler(1)
      return;
    }
    requestArticle(keyword, 0);
    setCurrentPage(1)
    pageChangeHandler(1)
  };

  const handleUpdateTable = (id: number) => {
    requestDeleteArticle(id);
  };
  return (
    <div className="adminArticles">
      <h1 className="adminArticles__title">Manage Articles</h1>
      <div className="adminArticles__overview">
        <div className="adminArticles__overview--item">
          <p>
            Total article: <span>{total}</span>
          </p>
        </div>
        <div className="adminArticles__overview--item">
          <p>
            Number of articles scratched in the previous 24 hours:{' '}
            <span>{today}</span>
          </p>
        </div>
      </div>
      <div className="adminArticles__list">
        <h2 className="adminArticles__list--title">Articles list</h2>
        <div className="adminArticles__list--search">
          <Form onSubmit={(event) => onSearchSubmit(event)}>
            <InputGroup className="mb-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faMagnifyingGlass} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Search articles"
                type="text"
                value={keyword}
                onChange={(event) => setkeyword(event.target.value)}
              />
            </InputGroup>
          </Form>
        </div>
        {articles ? (
          <ArticleTable
            articles={articles}
            currentPage={currentPage!}
            handleUpdateTable={handleUpdateTable}
          />
        ) : (
          <div className="adminArticles__table--loading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        )}
        <AdminArticlePagination
          totalRows={totalArticleSearch!}
          pageChangeHandler={pageChangeHandler}
        />
      </div>
    </div>
  );
}

export default ArticleAdmin;
