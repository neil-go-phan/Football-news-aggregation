import React, { useRef, useState } from 'react';
import EmbedWeb from './embedWeb';
import { useRouter } from 'next/router';
import * as yup from 'yup';
import { Button, Form, InputGroup, Table } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import {
  faHeading,
  faInfo,
  faLink,
  faList,
  faTag,
} from '@fortawesome/free-solid-svg-icons';
import { toast } from 'react-toastify';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ArticleType } from '@/components/matchDetail/relatedNews/article';

const AddCrawler: React.FC = () => {
  const [errorMessage, setErrorMessage] = useState({
    trigger: false,
    message: '',
  });
  const [htmlArticleList, setHtmlArticleList] = useState<string>();
  const [htmlArticleDiv, setHtmlArticleDiv] = useState<string>();
  const [htmlArticleTitle, setHtmlArticleTitle] = useState<string>();
  const [htmlArticleDescription, setHtmlArticleDescription] =
    useState<string>();
  const [htmlArticleLink, setHtmlArticleLink] = useState<string>();
  const [htmlNextPage, setHtmlNextPage] = useState<string>();
  const [nextPageType, setNextPageType] = useState<string>('button next');

  const [totalArticleCrawler, setTotalArticleCrawled] = useState<number>();
  const [isRenderResult, setIsRenderResult] = useState<boolean>(false);
  const [articles, setArticles] = useState<Array<ArticleType>>([]);

  const fieldChooseRef = useRef('');

  const handleChoose = (fieldChoosed: string) => {
    fieldChooseRef.current = fieldChoosed;
  };

  // JavaScript closure
  const handleClick = (event : Event): void => {
    const target = event.target as HTMLElement;
    const classname = target.className
    switch (fieldChooseRef.current) {
      case 'list':
        setHtmlArticleList(classname);
        break;
      case 'div':
        setHtmlArticleDiv(classname);
        break;
      case 'title':
        setHtmlArticleTitle(classname);
        break;
      case 'description':
        setHtmlArticleDescription(classname);
        break;
      case 'link':
        setHtmlArticleLink(classname);
      default:
        break;
    }
    toast.success('Get class successsdads', {
      position: 'top-right',
      autoClose: ERROR_POPUP_ADMIN_TIME,
      hideProgressBar: false,
      closeOnClick: true,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
      theme: 'light',
    });
    console.log("Helo")
  };
  const router = useRouter();
  const { url } = router.query;

  const schema = yup.object().shape({
    article_list: yup.string().required('Article list is require'),
    article_div: yup.string().required('Article div is require'),
    article_title: yup.string().required('Article title is require'),
    article_description: yup.string().required('Article descrition is require'),
    article_link: yup.string().required('Article link is require'),
    next_page_type: yup.string().required('Next page type is require'),
  });

  const handleSubmit = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    e.preventDefault();
    // validate
    try {
      await schema.validate(
        {
          article_list: htmlArticleList,
          article_div: htmlArticleDiv,
          article_title: htmlArticleTitle,
          article_description: htmlArticleDescription,
          article_link: htmlArticleLink,
          next_page_type: nextPageType,
        },
        { abortEarly: true }
      );
      requestSubmit();
    } catch (error) {
      if (error instanceof yup.ValidationError) {
        setErrorMessage({
          trigger: true,
          message: error.message,
        });
      }
    }
  };

  const handleTest = async (
    e: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    e.preventDefault();
    // validate
    try {
      await schema.validate(
        {
          article_list: htmlArticleList,
          article_div: htmlArticleDiv,
          article_title: htmlArticleTitle,
          article_description: htmlArticleDescription,
          article_link: htmlArticleLink,
          next_page_type: nextPageType,
        },
        { abortEarly: true }
      );
      requestTest();
    } catch (error) {
      if (error instanceof yup.ValidationError) {
        setErrorMessage({
          trigger: true,
          message: error.message,
        });
      }
    }
  };

  const requestSubmit = async () => {
    try {
      const { data } = await axiosProtectedAPI.post('crawler/upsert', {
        url: String(url),
        article_list: htmlArticleList,
        article_div: htmlArticleDiv,
        article_title: htmlArticleTitle,
        article_description: htmlArticleDescription,
        article_link: htmlArticleLink,
        next_page: htmlNextPage,
        next_page_type: nextPageType,
      });
      if (!data.success) {
        throw 'upsert fail';
      }
      setErrorMessage({
        trigger: false,
        message: data.message,
      });
      toast.success('Upsert success', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    } catch (error: any) {
      setErrorMessage({
        trigger: true,
        message: error.response.data.message,
      });
      toast.error('Error occurred while upsert crawler', {
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

  const requestTest = async () => {
    try {
      const { data } = await axiosProtectedAPI.post('crawler/test', {
        url: String(url),
        article_list: htmlArticleList,
        article_div: htmlArticleDiv,
        article_title: htmlArticleTitle,
        article_description: htmlArticleDescription,
        article_link: htmlArticleLink,
        next_page: htmlNextPage,
        next_page_type: nextPageType,
      });
      if (!data.success) {
        throw 'upsert fail';
      }
      setArticles(data.articles);
      setTotalArticleCrawled(data.amount);
      setIsRenderResult(true);
      setErrorMessage({
        trigger: false,
        message: data.message,
      });
      toast.success('Test success, result below', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    } catch (error: any) {
      setErrorMessage({
        trigger: true,
        message: error.response.data.message,
      });
      setIsRenderResult(false);
      toast.error('Error occurred while test crawler', {
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

  return (
    <div className="adminCrawler__addCrawler ">
      <div className="adminCrawler__addCrawler--add d-flex">
        <div className="adminCrawler__addCrawler--embed col-7">
          <EmbedWeb url={String(url)} handleClick={handleClick} />
        </div>
        <div className="adminCrawler__addCrawler--form col-5">
          <form>
            <h2 className="title">Input article class</h2>
            <div className="line" />
            <label> Article list </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faList} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article list classname"
                type="text"
                required
                className="bg-white"
                defaultValue={htmlArticleList} // Sử dụng defaultValue thay vì value
                onChange={(e) => setHtmlArticleList(e.target.value)}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('list')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article div </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faTag} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article div classname"
                type="text"
                required
                className="bg-white"
                value={htmlArticleDiv}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('div')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article title </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faHeading} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article title classname"
                type="text"
                required
                className="bg-white"
                value={htmlArticleTitle}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('title')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article description </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faInfo} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                placeholder="Type of choose article description classname"
                type="text"
                required
                className="bg-white"
                value={htmlArticleDescription}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('description')}
              >
                Choose
              </Button>
            </InputGroup>

            <label> Article link </label>
            <InputGroup className="my-3">
              <InputGroup.Text>
                <FontAwesomeIcon icon={faLink} fixedWidth />
              </InputGroup.Text>
              <Form.Control
                className="bg-white"
                placeholder="Type of choose article link classname"
                type="text"
                required
                value={htmlArticleLink}
              />
              <Button
                className="px-4"
                variant="primary"
                onClick={() => handleChoose('link')}
              >
                Choose
              </Button>
            </InputGroup>

            {errorMessage.trigger && (
              <p className="errorMessage errorFromServer">
                {errorMessage.message}
              </p>
            )}

            <div className="btnGroup d-flex justify-content-between">
              <Button
                className="w-30 px-4"
                variant="secondary"
                onClick={(e) => handleTest(e)}
                type="submit"
              >
                Test
              </Button>
              <Button
                className="w-30 px-4"
                variant="success"
                onClick={(e) => handleSubmit(e)}
                type="submit"
              >
                Submit
              </Button>
            </div>
          </form>
        </div>
      </div>

      {isRenderResult ? (
        <div className="adminCrawler__addCrawler--testResult">
          <div className="total">
            <p>
              Total article: <span>{totalArticleCrawler}</span>
            </p>
          </div>
          <div className="table">
            <Table striped bordered hover>
              <thead>
                <tr>
                  <th>#</th>
                  <th>Title</th>
                  <th>Description</th>
                  <th>Link</th>
                </tr>
              </thead>
              <tbody>
                {articles.map((article, index) => (
                  <tr>
                    <td>{index}</td>
                    <td>{article.title}</td>
                    <td>{article.description}</td>
                    <td>{article.link}</td>
                  </tr>
                ))}
              </tbody>
            </Table>
          </div>
        </div>
      ) : (
        <></>
      )}
    </div>
  );
};

export default AddCrawler;
