import React, { useState, useEffect } from 'react';
import { Button } from 'react-bootstrap';
import { ROW_PER_PAGE } from '.';

type Props = {
  totalRows: number;
  // eslint-disable-next-line no-unused-vars
  pageChangeHandler: (currentPage: number) => void;
};

const AdminArticlePagination: React.FC<Props> = ({
  pageChangeHandler,
  totalRows,
}) => {
  // Calculating max number of pages
  const noOfPages = Math.ceil(totalRows / ROW_PER_PAGE);
  
  const [currentPage, setCurrentPage] = useState(1);
  
  const [canGoBack, setCanGoBack] = useState(false);
  const [canGoNext, setCanGoNext] = useState(true);

  const onNextPage = () => setCurrentPage(currentPage + 1);
  const onPrevPage = () => setCurrentPage(currentPage - 1);

  useEffect(() => {
    if (noOfPages === currentPage) {
      setCanGoNext(false);
    } else {
      setCanGoNext(true);
    }
    if (currentPage === 1) {
      setCanGoBack(false);
    } else {
      setCanGoBack(true);
    }
  }, [noOfPages, currentPage]);
  useEffect(() => {
    pageChangeHandler(currentPage);
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentPage]);
  useEffect(() => {
    setCurrentPage(1)
  }, [totalRows]);
  return (
    <div className="btnPaging">
      <Button
        onClick={() => onPrevPage()}
        disabled={!canGoBack}
        variant="primary"
      >
        Previous Page
      </Button>
      <Button
        onClick={() => onNextPage()}
        disabled={!canGoNext}
        variant="primary"
      >
        Next Page
      </Button>
      <p>
        Page
        <span>
          {currentPage} of {noOfPages}
        </span>
      </p>
    </div>
  );
};

export default AdminArticlePagination;
