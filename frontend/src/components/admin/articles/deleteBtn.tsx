import React from 'react';
import { Button } from 'react-bootstrap';
type Props = {
  id: number;
  // eslint-disable-next-line no-unused-vars
  handleUpdateTable: (id: number) => void;
};

const DeleteArticleBtn: React.FC<Props> = (props: Props) => {
  return (
    <Button
      variant="danger"
      onClick={() => props.handleUpdateTable(props.id)}
    >
      Delete
    </Button>
  );
};

export default DeleteArticleBtn;
