import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import MatchDetailArticle, { ArticleType } from './article';
import { ERROR_POPUP_USER_TIME } from '@/helpers/constants';

function RelatedNews() {
  const [expanded, setExpanded] = useState(false);
  const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const router = useRouter();

  const requestFirstPageArticle = async () => {
    try {
      const { data } = await axiosClient.get('article/get-first-page', {
        params: { tags: router.query.league },
      });
      setArticles(data.articles);
    } catch (error) {
      toast.error('Error occurred while get article', {
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

  useEffect(() => {
    requestFirstPageArticle();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const tagForDisplay: Array<ArticleType> = expanded
    ? articles
    : articles!.slice(0, 3);

  return (
    <div className="matchDetail__relatedNews p-3">
      <div className="matchDetail__relatedNews--title">Related news</div>
      <div className="matchDetail__relatedNews--line my-3"></div>
      <div className="matchDetail__relatedNews--news">
        {articles.length !== 0 ? (
          tagForDisplay.map((article: ArticleType) => (
            <MatchDetailArticle
              key={`matchDetail-${article.title}`}
              article={article}
            ></MatchDetailArticle>
          ))
        ) : (
          <div className="news__articles--noresults">Not found</div>
        )}
        {articles.length > 3 ? (
          <p className="showmore" onClick={() => setExpanded(!expanded)}>
            {expanded ? 'Show less' : 'Show more...'}
          </p>
        ) : (
          <></>
        )}
      </div>
    </div>
  );
}

export default RelatedNews;
