import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';
import MatchDetailArticle, { ArticleType } from './article';

function RelatedNews() {
  const [expanded, setExpanded] = useState(false);
  const [articles, setArticles] = useState<Array<ArticleType>>([]);
  const router = useRouter()
  const requestArticle = async () => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: '', tags: router.query.league},
      });
      setArticles(data.articles)
    } catch (error) {
      toast.error(
        `Error occurred while get article`,
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

  useEffect(() => {
    requestArticle()
  }, [])
  
  const tagForDisplay: Array<ArticleType> = expanded ? articles : articles!.slice(0, 3);

  return (
    <div className="matchDetail__relatedNews p-3">
      <div className="matchDetail__relatedNews--title">Tin liên quan</div>
      <div className="matchDetail__relatedNews--line my-3"></div>
      <div className="matchDetail__relatedNews--news">
        {articles.length !== 0 ? (
          tagForDisplay.map((article: ArticleType) => (
            <MatchDetailArticle key={`matchDetail-${article.title}`} article={article}></MatchDetailArticle>
          ))
        ) : (
          <div className="news__articles--noresults">
            Không tìm thấy kết quả
          </div>
        )}
        <p className="showmore" onClick={() => setExpanded(!expanded)}>
          {expanded ? 'Show less' : 'Show more...'}
        </p>
      </div>
    </div>
  );
}

export default RelatedNews;
