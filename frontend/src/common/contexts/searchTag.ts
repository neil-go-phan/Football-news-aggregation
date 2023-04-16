import { createContext } from 'react';

interface SearchTagsContextType {
  searchTags: Array<string>
  // eslint-disable-next-line no-unused-vars
  setSearchTags: (tags: Array<string>) => void;
}

const SearchTagContext = createContext<SearchTagsContextType>({
  searchTags: [],
  setSearchTags: () => {}
})

export default SearchTagContext;