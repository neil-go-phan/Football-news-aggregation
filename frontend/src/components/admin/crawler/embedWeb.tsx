import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React, { useEffect, useState } from 'react';
import { ThreeDots } from 'react-loader-spinner';
import { toast } from 'react-toastify';

type Props = {
  url: string;
  // eslint-disable-next-line no-unused-vars
  handleClick: (event: Event) => void;
};

const EmbedWeb: React.FC<Props> = (props: Props) => {
  const [htmlContent, setHtmlContent] = useState<string>();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [isError, setIsError] = useState<boolean>(false);

  const handleMouseOver = (event: Event) => {
    const target = event.target as HTMLElement;
    target.style.backgroundColor = 'rgba(255, 0, 0, 0.2)';
  };

  const handleMouseOut = (event: Event) => {
    const target = event.target as HTMLElement;
    target.style.backgroundColor = '';
  };

  const removeOnClickEvents = () => {
    const tempContainer = document.createElement('div');
    tempContainer.innerHTML = htmlContent!;
    const anchorElements = tempContainer.querySelectorAll('a');
    // gỡ sự kiện onClick chuyển trang của thẻ a
    anchorElements.forEach((element) => {
      element.removeAttribute('href');
      element.addEventListener('click', (event) => {
        event.preventDefault();
      });
    });
    const elements = tempContainer.querySelectorAll('*');
    // Gỡ bỏ sự kiện onclick mạc định của tất cả thẻ khác
    elements.forEach((element) => {
      element.removeAttribute('onclick');
    });
    const sanitizedHtml = tempContainer.innerHTML;
    setHtmlContent(sanitizedHtml);
  };
  const addClickEventToContainer = () => {
    const container = document.getElementById('container');
    if (container) {
      container.addEventListener('click', props.handleClick);
      const elements = container.querySelectorAll('*');

      elements.forEach((element) => {
        element.addEventListener('mouseover', handleMouseOver);
        element.addEventListener('mouseout', handleMouseOut);
      });
    }
  };
  useEffect(() => {
    removeOnClickEvents();
    addClickEventToContainer();
  }, [htmlContent]);

  const requestHtmlPage = async (url: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('crawler/get-html-page', {
        params: { url: url },
      });
      if (data.success === false) {
        throw 'Throw error occurred while get html page';
      } else {
        toast.success('Get page success', {
          position: 'top-right',
          autoClose: ERROR_POPUP_ADMIN_TIME,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        });
        setHtmlContent(data);
        setIsLoading(false);
      }
    } catch (error) {
      setIsError(true);
      toast.error(`Error occurred`, {
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
    setIsLoading(true);
    requestHtmlPage(props.url);
  }, [props.url]);

  if (isError) {
    return (
      <div className="adminCrawler__iFrame">
        <div className="adminCrawler__iFrame--loading">Error</div>
      </div>
    );
  }
  if (htmlContent) {
    return (
      <div className="adminCrawler__iFrame">
        {isLoading ? (
          <div className="adminCrawler__iFrame--loading">
            <ThreeDots
              height="50"
              width="50"
              radius="9"
              color="#4fa94d"
              ariaLabel="three-dots-loading"
              visible={true}
            />
          </div>
        ) : (
          <div
            className="adminCrawler__iFrame--embed"
            id="container"
            dangerouslySetInnerHTML={{ __html: htmlContent }}
          ></div>
        )}
      </div>
    );
  }

  return (
    <div className="adminCrawler__iFrame">
      <div className="adminCrawler__iFrame--loading">
        <ThreeDots
          height="50"
          width="50"
          radius="9"
          color="#4fa94d"
          ariaLabel="three-dots-loading"
          visible={true}
        />
      </div>
    </div>
  );
};

export default EmbedWeb;
