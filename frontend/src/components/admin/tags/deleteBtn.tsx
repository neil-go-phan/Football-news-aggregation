import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React from 'react';
import { Button } from 'react-bootstrap';
import { toast } from 'react-toastify';

type Props = {
  isDisabled: boolean;
  tagName: string;
  handleDeleteTag: () => void
};

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const requestDeleteTag = async (tag: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('tags/delete', {
        params: { tag: tag },
      });
      if (!data.success) {
        throw 'delete fail';
      }
      toast.success('Delete tag success', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      props.handleDeleteTag()
    } catch (error) {
      toast.error('Error occurred while delete tags', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
    }
  };

  const handlerDeteleTag = (
    event: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    const target = event.target as HTMLButtonElement;
    requestDeleteTag(target.value);
  };

  return (
    <Button
      variant="danger"
      value={props.tagName}
      onClick={(event) => handlerDeteleTag(event)}
      disabled={props.isDisabled}
    >
      Delete
    </Button>
  );
};

export default DeleteBtn;
