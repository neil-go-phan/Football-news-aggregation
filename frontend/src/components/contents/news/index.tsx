import React, { useState } from 'react';
import Article from './article';
import { ArticleType } from './article';
import SearchBar from './searchBar';


export default function News() {
  const [articles, setArticles] = useState<Array<ArticleType>>([]);

  const handleSearch = (searchResult: Array<ArticleType>) => {
    setArticles(searchResult);
  };
  console.log(articles)
  return (
    <div className="news">
      <div className="news__searchBar px-3">
        <SearchBar handleSearch={handleSearch}/>
      </div>

      <div className="news__articles">
      <div className="articles">
      {articles.map((article: ArticleType) => (
        <Article key={article.title} article={article}></Article>
      ))}
    </div>
      </div>
    </div>
  );
}
