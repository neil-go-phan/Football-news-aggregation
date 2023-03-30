import SearchTagContext from '@/common/contexts/searchTag';
import React, { useContext, useState } from 'react';
import Article, { ArticleType } from './article';
import SearchBar from './searchBar';

export default function News() {
  const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const { searchTags, setSearchTags } = useContext(SearchTagContext);
  const handleSearch = (searchResult: Array<ArticleType>) => {
    setArticles(searchResult);
  };
  return (
    <div className="news px-2">
      <div className="news__searchBar px-3">
        <SearchBar handleSearch={handleSearch} tagsList={searchTags}/>
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
