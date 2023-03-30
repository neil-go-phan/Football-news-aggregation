import React, { FunctionComponent, useContext, useEffect, useState } from 'react';
import InputGroup from 'react-bootstrap/InputGroup';
import Form from 'react-bootstrap/Form';
import axiosClient from '@/helpers/axiosClient';
import { useRouter } from 'next/router';
import { ArticleType } from './article';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faMagnifyingGlass, faX } from '@fortawesome/free-solid-svg-icons';
import SearchTagContext from '@/common/contexts/searchTag';

interface Props {
  // eslint-disable-next-line no-unused-vars
  handleSearch: (searchResult: Array<ArticleType>) => void;
  tagsList: Array<string>
}

const SearchBar: FunctionComponent<Props> = (props: Props) => {
  const [keyword, setkeyword] = useState<string>('');
  const [tagsInSearch, setTagsInSearch] = useState<Array<string>>([])
  const router = useRouter();
  


  const getIndex = (): string => {
    const isContainNewsPath = router.asPath.search('/news/');
    if (isContainNewsPath === -1) {
      return '_all';
    }
    return router.asPath.slice(6);
  };
  console.log(props.tagsList)

  const requestArticle = async (searchKeyword: string) => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: searchKeyword, index: getIndex() },
      }); 
      props.handleSearch(data.articles);
    } catch (error) {}
  }

  // handle when user change route
  useEffect(() => {
    const storedValue = window.sessionStorage.getItem(`search_keyword_${getIndex()}`);
    if (storedValue !== null) {
      setkeyword(storedValue)
      requestArticle(storedValue)
    } else {
      setkeyword('')
      props.handleSearch([]);
    }
  }, [router.asPath])

    // handle when click add tag

  
  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    window.sessionStorage.setItem(`search_keyword_${getIndex()}`, keyword);
    // window.sessionStorage.setItem(`search_tags_${getIndex()}`, tagsInSearch);
    requestArticle(keyword)
  };

  const handleDeleteTag = () => {
    console.log("delete tag")
  }

  return (
    <Form onSubmit={(event) => onSearchSubmit(event)}>
      <InputGroup className="mb-3 news__searchBar--search">
        <span className="icon"><FontAwesomeIcon icon={faMagnifyingGlass} /></span>
        <div className="tags">
        
          <div className="tag">
            <span>Heelo</span>
            <span className='tag--icon' onClick={handleDeleteTag}><FontAwesomeIcon icon={faX} /></span>
          </div>
        </div>
        <input
          placeholder="Tìm kiếm"
          value={keyword}
          onChange={(event) => setkeyword(event.target.value)}
        />
      </InputGroup>
    </Form>
  );
};

export default SearchBar;
