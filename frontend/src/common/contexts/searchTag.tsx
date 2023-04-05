import { createContext } from 'react';

interface SearchTagsContextType {
  searchTags: Array<string>
  setSearchTags: (tags: Array<string>) => void;
}

const SearchTagContext = createContext<SearchTagsContextType>({
  searchTags: [],
  setSearchTags: () => {}
})

export default SearchTagContext;