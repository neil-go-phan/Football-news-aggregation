import React, { FunctionComponent, useState } from 'react';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import { ArticleType } from '../article';


interface Props {
  handleSearch: (searchResult: Array<ArticleType>) => void 
}

const SearchBar: FunctionComponent<Props> = (props:Props) =>{
  const [keyword, setkeyword] = useState<string>('');
  const router = useRouter();
  const getIndex = (): string => {
    const isContainNewsPath = router.asPath.search('/news/');
    if (isContainNewsPath === -1) {
      return '_all';
    }
    return router.asPath.slice(6);
  };
  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    try {
      const { data } = await axiosClient.get('article/search', {
        params: { q: keyword, index: getIndex() }
      });
      setkeyword('')
      props.handleSearch(data.articles)
    } catch (error) {
      
    }
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
}

export default SearchBar;