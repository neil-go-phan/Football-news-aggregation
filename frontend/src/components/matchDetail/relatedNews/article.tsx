import React, { FunctionComponent } from 'react';

interface ArticleProps {
  article: ArticleType;
}

export type ArticleType = {
  title: string;
  description: string;
  link: string;
  tags: Array<string>;
};

const MatchDetailArticle: FunctionComponent<ArticleProps> = (props: ArticleProps) => {
  const { article } = props;
  return (
    <div className="article mb-3 py-2">
      <a className="" href={article.link} target="_blank">
        <div className="title">{article.title}</div>
        <div className="description">{article.description}</div>
      </a>
    </div>
  );
};

export default MatchDetailArticle;
