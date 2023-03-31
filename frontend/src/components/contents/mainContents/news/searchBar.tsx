import React, {
  FunctionComponent,
  useContext,
  useEffect,
  useState,
} from 'react';
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
}

const SearchBar: FunctionComponent<Props> = (props: Props) => {
  const [keyword, setkeyword] = useState<string>('');
  const { searchTags, setSearchTags } = useContext(SearchTagContext);
  const [tagsInSearch, setTagsInSearch] = useState<Array<string>>([]);
  const router = useRouter();
  
  console.log("search render")

  const getIndex = (): string => {
    const isContainNewsPath = router.asPath.search('/news/');
    if (isContainNewsPath === -1) {
      return '';
    }
    return router.asPath.slice(6);
  };

  const getTagParam = ():string => {
    let tagParam:string = ""
    if (searchTags.indexOf(getIndex()) < 0) {
      tagParam += getIndex() + ","
    }
    searchTags.forEach((tag) => tagParam += tag + "," )
    tagParam = tagParam.slice(0, tagParam.length - 1)
    return tagParam
  }

  const requestArticle = async (searchKeyword: string) => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: searchKeyword.trim(), tags : getTagParam()},
      });
      props.handleSearch(data.articles);
    } catch (error) {
      
    }
  };

  // handle when user change route
  useEffect(() => {
    const storedValue = window.sessionStorage.getItem(
      `search_keyword_${getIndex()}`
    );
    if (storedValue !== null) {
      setkeyword(storedValue);
      requestArticle(storedValue);
    } else {
      setkeyword('');
      props.handleSearch([]);
    }
  }, [router.asPath]);

  useEffect(() => {
    setTagsInSearch(searchTags)
  }, [searchTags])
  

  // handle when click add tag

  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if ((keyword.trim() === "") && (searchTags.length === 0)) {
      setkeyword('');
      props.handleSearch([]);
      return
    }
    window.sessionStorage.setItem(`search_keyword_${getIndex()}`, keyword.trim());
    // window.sessionStorage.setItem(`search_tags_${getIndex()}`, tagsInSearch);
    requestArticle(keyword);
  };

  const handleDeleteTag = (tag : string) => {
    const tags = searchTags
    const index = tags.indexOf(tag)
    tags.splice(index, 1)
    setSearchTags(tags)
    setTagsInSearch(tags)
  };

  return (
    <Form onSubmit={(event) => onSearchSubmit(event)}>
      <InputGroup className="mb-3 news__searchBar--search">
        <span className="icon">
          <FontAwesomeIcon icon={faMagnifyingGlass} />
        </span>
        <div className="tags">
          {tagsInSearch.map((tag) => (
            <div key={`search_tag_name_${tag}`} className="tag">
              <span>{tag}</span>
              <span className="tag--icon" onClick={() => handleDeleteTag(tag)}>
                <FontAwesomeIcon icon={faX} />
              </span>
            </div>
          ))}
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
