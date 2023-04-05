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
  handleSearch: (searchResult: Array<ArticleType>) => void;
}

const SearchBar: FunctionComponent<Props> = (props: Props) => {
  const [keyword, setkeyword] = useState<string>('');
  const { searchTags, setSearchTags } = useContext(SearchTagContext);
  const router = useRouter();

  const getDefaultTag = (): string => {
    const isContainNewsPath = router.asPath.search('/news/');
    if (isContainNewsPath === -1) {
      return '';
    }
    const defaultTag = router.asPath.slice(6);

    return defaultTag.replace('-', ' ');
  };

  const getTagParam = (): string => {
    let tagParam: string = '';
    if (searchTags.indexOf(getDefaultTag()) < 0) {
      tagParam += getDefaultTag() + ',';
    }
    searchTags.forEach((tag) => (tagParam += tag + ','));
    tagParam = tagParam.slice(0, tagParam.length - 1);
    return tagParam;
  };

  const requestArticle = async (searchKeyword: string) => {
    try {
      const { data } = await axiosClient.get('article/search-tag-keyword', {
        params: { q: searchKeyword.trim(), tags: getTagParam() },
      });
      props.handleSearch(data.articles);
      toast.success(`Search keyword ${searchKeyword.trim()} success`, {
        position: 'top-right',
        autoClose: 100,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    } catch (error) {
      toast.error(
        `Error occurred while searching keyword ${searchKeyword.trim()}`,
        {
          position: 'top-right',
          autoClose: 3000,
          hideProgressBar: false,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
          theme: 'light',
        }
      );
    }
  };

  // handle when user change route
  useEffect(() => {
    setkeyword('');
    requestArticle('');
  }, [router.asPath]);

  const onSearchSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    if (keyword.trim() === '' && searchTags.length === 0) {
      setkeyword('');
      requestArticle('');
      return;
    }
    requestArticle(keyword);
  };

  const handleDeleteTag = (tag: string) => {
    const tags = searchTags;
    const index = tags.indexOf(tag);
    tags.splice(index, 1);
    setSearchTags([...tags]);
  };

  return (
    <Form onSubmit={(event) => onSearchSubmit(event)}>
      <InputGroup className='mb-3 news__searchBar--search'>
        <span className='icon'>
          <FontAwesomeIcon icon={faMagnifyingGlass} />
        </span>
        <div className='tags'>
          {searchTags.map((tag) => (
            <div key={`search_tag_name_${tag}`} className='tag'>
              <span>{tag}</span>
              <span className='tag--icon' onClick={() => handleDeleteTag(tag)}>
                <FontAwesomeIcon icon={faX} />
              </span>
            </div>
          ))}
        </div>
        <input
          placeholder='Tìm kiếm'
          value={keyword}
          onChange={(event) => setkeyword(event.target.value)}
        />
      </InputGroup>
    </Form>
  );
};

export default SearchBar;
