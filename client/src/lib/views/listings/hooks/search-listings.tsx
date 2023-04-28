import { useContext } from 'react'

import { SearchListingContext } from '../context/search-listings'

export const useSearchListings = () => useContext(SearchListingContext)
