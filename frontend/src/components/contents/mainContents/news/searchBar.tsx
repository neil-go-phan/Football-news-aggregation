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
import { toast } from 'react-toastify';

interface Props {
  // eslint-disable-next-line no-unused-vars
  handleSearchArticle: (keywordSearch: string, from: number) => void
}

const DEFAULT_PAGE = 0

const SearchBar: FunctionComponent<Props> = (props: Props) => {
  const [keyword, setkeyword] = useState<string>('');
  const { searchTags, setSearchTags } = useContext(SearchTagContext);
  const router = useRouter();

  // handle when user change route
  useEffect(() => {
    setkeyword('');
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [router.asPath]);

  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (keyword.trim() === '' && searchTags.length === 0) {
      setkeyword('');
      props.handleSearchArticle('', DEFAULT_PAGE);
      return;
    }
    props.handleSearchArticle(keyword, DEFAULT_PAGE);
  };

  const handleDeleteTag = (tag: string) => {
    const tags = searchTags;
    const index = tags.indexOf(tag);
    tags.splice(index, 1);
    setSearchTags([...tags]);
  };

  return (
    <Form onSubmit={(event) => onSearchSubmit(event)}>
      <InputGroup className="mb-3 news__searchBar--search">
        <span className="icon">
          <FontAwesomeIcon icon={faMagnifyingGlass} />
        </span>
        <div className="tags">
          {searchTags.map((tag) => (
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
