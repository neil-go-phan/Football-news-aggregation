import React, { FunctionComponent } from 'react';

interface Props {
  article: ArticleType;
}

export type ArticleType = {
  title: string;
  description: string;
  link: string;
};

const Article: FunctionComponent<Props> = (props: Props) => {
  const { article } = props;
  return (
    <div className="news__articles--article mb-3 py-2">
      <a className="" href={article.link} target='_blank'>
          <div className="title">{article.title}</div>
          <div className="description">{article.description}</div>
      </a>
    </div>
  );
};

export default Article;
