
import React from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLink } from '@fortawesome/free-solid-svg-icons';

type EditTimeProperty = {
  newTime: number;
};

type Props = {
  // eslint-disable-next-line no-unused-vars
  handleIsModalClose: (newTime:number) => void;
};

const EditTimeModal: React.FC<Props> = (props: Props) => {
  const schema = yup.object().shape({
    newTime: yup.number()
      .required('Please enter time'),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<EditTimeProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<EditTimeProperty> = async (data) => {
    let { newTime } = data;
    props.handleIsModalClose(newTime);
  };
  return (
    <div className="adminCronjob__editTimeModal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminCronjob__editTimeModal--title">Input new time</h2>
        <div className="adminCronjob__editTimeModal--line" />
        <label> Time (min) </label>
        <InputGroup className="mb-3">
          <InputGroup.Text>
            <FontAwesomeIcon icon={faLink} fixedWidth />
          </InputGroup.Text>
          <Form.Control
            {...register('newTime')}
            placeholder="Type new time (min)"
            type="text"
            required
          />
        </InputGroup>

        {errors.newTime && (
          <p className="errorMessage">{errors.newTime.message}</p>
        )}
        <Button className="w-100 px-4" variant="success" type="submit">
          Continue
        </Button>
      </form>
    </div>
  );
};

export default EditTimeModal;
