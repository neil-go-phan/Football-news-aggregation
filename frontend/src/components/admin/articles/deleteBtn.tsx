import React from 'react'
import { Button } from 'react-bootstrap';
type Props = {
  title: string
  index: number
  handleUpdateTable:(title: string, index: number) => void
};

const DeleteArticleBtn: React.FC<Props> = (props: Props) => {
  return (
    <Button
      variant="danger"
      value={props.index}
      onClick={() => props.handleUpdateTable(props.title, props.index)}
    >
      Delete
    </Button>
  )
}

export default DeleteArticleBtn