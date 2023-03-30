import React, { FunctionComponent, useEffect, useState } from 'react';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import { ArticleType } from '../article';
import nextSession from "next-session";


interface Props {
  // eslint-disable-next-line no-unused-vars
  handleSearch: (searchResult: Array<ArticleType>) => void;
}

const SearchBar: FunctionComponent<Props> = (props: Props) => {

  const [keyword, setkeyword] = useState<string>('');
  const router = useRouter();
  const getIndex = (): string => {
    const isContainNewsPath = router.asPath.search('/news/');
    if (isContainNewsPath === -1) {
      return '_all';
    }
    return router.asPath.slice(6);
  };

  const requestArticle = async (searchKeyword: string) => {
    try {
      const { data } = await axiosClient.get('article/search', {
        params: { q: searchKeyword, index: getIndex() },
      }); 
      props.handleSearch(data.articles);
    } catch (error) {}
  }

  // handle when user change route
  useEffect(() => {
    const storedValue = window.sessionStorage.getItem(getIndex());
    if (storedValue !== null) {
      setkeyword(storedValue)
      requestArticle(storedValue)
    } else {
      setkeyword('')
      props.handleSearch([]);
    }
  }, [router.asPath])
  
  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    window.sessionStorage.setItem(getIndex(), keyword);
    requestArticle(keyword)
  };
  
  return (
    <Form onSubmit={(event) => onSearchSubmit(event)}>
      <InputGroup className="mb-3">
        <Form.Control
          placeholder="Tìm kiếm"
          aria-label="search_field"
          value={keyword}
          onChange={(event) => setkeyword(event.target.value)}
        />
      </InputGroup>
    </Form>
  );
};

export default SearchBar;
