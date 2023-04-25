import React from 'react';
import { Button } from 'react-bootstrap';
type Props = {
  title: string;
  // eslint-disable-next-line no-unused-vars
  handleUpdateTable: (title: string) => void;
};

const DeleteArticleBtn: React.FC<Props> = (props: Props) => {
  return (
    <Button
      variant="danger"
      onClick={() => props.handleUpdateTable(props.title)}
    >
      Delete
    </Button>
  );
};

export default DeleteArticleBtn;
