import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React, { useCallback, useEffect, useRef, useState } from 'react';
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
      element.removeAttribute('onmouseover');
      element.removeAttribute('onmouseout');
    });
    const sanitizedHtml = tempContainer.innerHTML;
    setHtmlContent(sanitizedHtml);
  };

  const iframeRef = useRef<any>(null);

  useEffect(() => {
    removeOnClickEvents()
    const iframe = iframeRef.current;
    if (iframe) {
      const iframeDocument =
        iframe.contentDocument || iframe.contentWindow.document;

      iframeDocument.open();
      iframeDocument.write(htmlContent);
      iframeDocument.close();

      const elements = iframeDocument.querySelectorAll('*');
      elements.forEach((element: any) => {
        element.addEventListener('mouseover', handleMouseOver);
        element.addEventListener('mouseout', handleMouseOut);
        element.addEventListener('click', handleClick);
      });

      return () => {
        elements.forEach((element: any) => {
          element.removeEventListener('mouseover', handleMouseOver);
          element.removeEventListener('mouseout', handleMouseOut);
          element.removeEventListener('click', handleClick);
        });
      };
    }
      // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [htmlContent]);

  const handleMouseOver = (event:Event) => {
    const target = event.target  as HTMLElement; 
    target.style.backgroundColor = 'rgba(255, 0, 0, 0.2)';
  };

  const handleMouseOut = (event:Event) => {
    const target = event.target as HTMLElement;
    target.style.backgroundColor = '';
  }

  const handleClick = useCallback((event:Event) => {
    // Xử lý logic khi phần tử HTML được click
    props.handleClick(event)
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

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
      toast.error('Error occurred', {
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
          <iframe
            className="adminCrawler__iFrame--embed"
            ref={iframeRef}
          ></iframe>
          // <div
          //   className="adminCrawler__iFrame--embed"
          //   id="embedContainer"
          //   dangerouslySetInnerHTML={{ __html: htmlContent }}
          // ></div>
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
