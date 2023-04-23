import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React, { useEffect, useState } from 'react';
import { toast } from 'react-toastify';

function ArticleAdmin() {
  const [today, setToday] = useState<number>();
  const [total, setTotal] = useState<number>();

  const requestArticleCount = async () => {
    try {
      const { data } = await axiosProtectedAPI.get('article/count', {});
      setTotal(data.total);
      setToday(data.today);
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

  useEffect(() => {
    requestArticleCount();
  }, []);

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
    </div>
  );
}

export default ArticleAdmin;
