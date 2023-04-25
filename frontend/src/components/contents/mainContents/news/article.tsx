import React, { FunctionComponent } from 'react';
import Tag from '../../../../common/tag';

interface ArticleProps {
  article: ArticleType;
}

export type ArticleType = {
  title: string;
  description: string;
  link: string;
  tags: Array<string>;
};

const Article: FunctionComponent<ArticleProps> = (props: ArticleProps) => {
  const getDomainName = (url: string): string => {
    let domain = (new URL(url));
    return domain.hostname.replace('www.','');
  };
  const { article } = props;
  return (
    <div className="news__articles--article mb-3 py-2">
      <a className="" href={article.link} target="_blank">
        <div className="title">{article.title}</div>
        <div className="source mt-1">{getDomainName(article.link)}</div>
        <div className="description">{article.description}</div>
      </a>
      <div className="tags d-flex mt-2">
        {article.tags.map((tag) => {
          if (tag !== 'tin tuc bong da') {
            return <Tag key={`news_articles_${tag}`} tagName={tag} />;
          }
        })}
      </div>
    </div>
  );
};

export default Article;
