import axiosProtectedAPI from '@/helpers/axiosProtectedAPI';
import { ERROR_POPUP_ADMIN_TIME } from '@/helpers/constants';
import React from 'react';
import { Button } from 'react-bootstrap';
import { toast } from 'react-toastify';

type Props = {
  url: string;
  handleDelete: () => void
};

const DeleteBtn: React.FC<Props> = (props: Props) => {
  const requestDelete= async (url: string) => {
    try {
      const { data } = await axiosProtectedAPI.get('crawler/delete', {
        params: { url: url },
      });
      if (!data.success) {
        throw 'delete fail';
      }
      toast.success('Delete success', {
        position: 'top-right',
        autoClose: ERROR_POPUP_ADMIN_TIME,
        hideProgressBar: false,
        closeOnClick: true,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: 'light',
      });
      props.handleDelete();
    } catch (error) {
      toast.error('Error occurred while delete crawler', {
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

  const handlerDetele = (
    event: React.MouseEvent<HTMLButtonElement, MouseEvent>
  ) => {
    const target = event.target as HTMLButtonElement;
    requestDelete(target.value);
  };

  return (
    <>
      <Button
        variant="danger"
        value={props.url}
        onClick={(event) => handlerDetele(event)}
      >
        Delete
      </Button>
    </>
  );
};

export default DeleteBtn;
