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
  // const addClickEventToContainer = () => {
  //   const container = document.getElementById('embedContainer');
  //   if (container) {
  //     container.addEventListener('click', props.handleClick);
  //     const elements = container.querySelectorAll('*');

  //     elements.forEach((element) => {
  //       element.addEventListener('mouseover', handleMouseOver);
  //       element.addEventListener('mouseout', handleMouseOut);
  //     });
  //   }
  // };
  // useEffect(() => {
  //   removeOnClickEvents();
  // }, [htmlContent]);

  const iframeRef = useRef<any>(null);

  useEffect(() => {
    removeOnClickEvents()
    const iframe = iframeRef.current;
    console.log(iframe)
    if (iframe) {
      const iframeDocument =
        iframe.contentDocument || iframe.contentWindow.document;

      // Gắn HTML content vào iframe
      iframeDocument.open();
      iframeDocument.write(htmlContent);
      iframeDocument.close();

      // Gắn sự kiện onMouseOver vào các phần tử HTML
      const elements = iframeDocument.querySelectorAll('*');
      elements.forEach((element: any) => {
        element.addEventListener('mouseover', handleMouseOver);
        element.addEventListener('mouseout', handleMouseOut);
        element.addEventListener('click', handleClick);
      });

      return () => {
        // Hủy bỏ sự kiện khi component unmount
        elements.forEach((element: any) => {
          element.removeEventListener('mouseover', handleMouseOver);
          element.removeEventListener('mouseout', handleMouseOut);
          element.removeEventListener('click', handleClick);
        });
      };
    }
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
  console.log('render');
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
