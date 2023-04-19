import axiosClient from '@/helpers/axiosClient';
import React, { useContext, useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import Article, { ArticleType } from './article';
import SearchBar from './searchBar';
import InfiniteScroll from 'react-infinite-scroll-component';
import { useRouter } from 'next/router';
import SearchTagContext from '@/common/contexts/searchTag';
import useWindowDimensions from '@/helpers/useWindowResize';
import { ThreeDots } from 'react-loader-spinner';

const DEFAULT_PAGE = 0;

export default function News() {
  const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const [keyword, setkeyword] = useState<string>('');
  const [from, setFrom] = useState<number>(0);
  const [hasMore, setHasMore] = useState<boolean>(true);

  const { searchTags } = useContext(SearchTagContext);
  const router = useRouter();
  const { height } = useWindowDimensions();

  const handleSearchArticle = (keywordSearch: string) => {
    setkeyword(keywordSearch);
    setFrom(DEFAULT_PAGE + 10);
    requestFirstPageArticle(keywordSearch, DEFAULT_PAGE);
  };

  const handleRequestMoreArticle = () => {
    requestArticle(keyword, from);
    setFrom(from + 10);
  };

  const getDefaultTag = (): string => {
    const isContainNewsPath = router.asPath.search('/news/');
    if (isContainNewsPath === -1) {
      return '';
    }
    let defaultTag = router.asPath.slice(6);
    defaultTag = defaultTag.substring(0, defaultTag.indexOf('?'));
    return defaultTag.replace('-', ' ');
  };

  const getTagParam = (): string => {
    let tagParam: string = '';
    if (searchTags.indexOf(getDefaultTag()) < 0) {
      tagParam += `${getDefaultTag()},`;
    }
    searchTags.forEach((tag) => (tagParam += `${tag},`));
    tagParam = tagParam.slice(0, tagParam.length - 1);
    return tagParam;
  };

  const requestArticle = async (searchKeyword: string, from: number) => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: searchKeyword.trim(), tags: getTagParam(), from: from },
      });
      if (data.articles.length === 0) {
        setHasMore(false);
      }
      const newArticle = articles.concat(data.articles);
      setArticles([...newArticle]);
    } catch (error) {
      toast.error(
        `Error occurred while searching keyword ${searchKeyword.trim()}`,
        {
          position: 'top-right',
          autoClose: 1000,
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

  const requestFirstPageArticle = async (
    searchKeyword: string,
    from: number
  ) => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: searchKeyword.trim(), tags: getTagParam(), from: from },
      });
      setArticles(data.articles);
    } catch (error) {
      toast.error(
        `Error occurred while searching keyword ${searchKeyword.trim()}`,
        {
          position: 'top-right',
          autoClose: 1000,
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

  const resetState = () => {
    setkeyword('');
    setHasMore(true)
    setFrom(DEFAULT_PAGE);
  }

  // handle when user change route
  useEffect(() => {
    requestFirstPageArticle('', DEFAULT_PAGE);
    resetState()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router.asPath]);

  return (
    <div className="news px-2">
      <div className="news__searchBar px-3">
        <SearchBar handleSearchArticle={handleSearchArticle} />
      </div>

      <div className="news__articles px-4" id="scrollableDiv">
        {articles.length !== 0 ? (
          <InfiniteScroll
            dataLength={articles.length}
            next={handleRequestMoreArticle}
            height={height ? height - 200 : 500}
            hasMore={hasMore}
            loader={
              <div className="news__articles--loading">
                <ThreeDots
                  height="50"
                  width="50"
                  radius="9"
                  color="#4fa94d"
                  ariaLabel="three-dots-loading"
                  visible={true}
                />
              </div>
            }
            endMessage={
              <div className="news__articles--loading">
                <p>
                  <b>There is no more articles</b>
                </p>
              </div>
            }
          >
            {articles.map((article: ArticleType, index) => (
              <Article
                key={`news-${getTagParam()}-${article.title}-${index}`}
                article={article}
              ></Article>
            ))}
          </InfiniteScroll>
        ) : (
          <div className="news__articles--noresults">
            Không tìm thấy kết quả
          </div>
        )}
      </div>
    </div>
  );
}
