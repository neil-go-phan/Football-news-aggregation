import axiosClient from '@/helpers/axiosClient';
import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import Article, { ArticleType } from './article';
import SearchBar from './searchBar';

export default function News() {
  const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const handleSearch = (searchResult: Array<ArticleType>) => {
    setArticles(searchResult);
  };
  useEffect(() => {
    const requestArticleFirstTime = async () => {
      try {
        const { data } = await axiosClient.get('article/search-tag-keyword', {
          params: { q: '', tags: 'tin+tuc+bong+da' },
        });
        setArticles(data.articles)
      } catch (error) {
        toast.error(
          'Error occurred while searching keyword get article',
          {
            position: 'top-right',
            autoClose: 3000,
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

    requestArticleFirstTime();
  }, []);

  return (
    <div className="news px-2">
      <div className="news__searchBar px-3">
        <SearchBar handleSearch={handleSearch} />
      </div>

      <div className="news__articles px-4">
        {articles.length !== 0 ? (
          articles.map((article: ArticleType) => (
            <Article key={article.title} article={article}></Article>
          ))
        ) : (
          <div className="news__articles--noresults">
            Không tìm thấy kết quả
          </div>
        )}
      </div>
    </div>
  );
}
