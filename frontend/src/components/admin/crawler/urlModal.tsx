
import React from 'react';
import * as yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button, Form, InputGroup } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLink } from '@fortawesome/free-solid-svg-icons';

type UrlFormProperty = {
  url: string;
};

type Props = {
  handleIsUrlModalClose: (url:string) => void;
};

const UrlModal: React.FC<Props> = (props: Props) => {
  const schema = yup.object().shape({
    url: yup.string()
      .url('Enter correct url!')
      .required('Please enter website'),
  });
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<UrlFormProperty>({
    resolver: yupResolver(schema),
  });
  const onSubmit: SubmitHandler<UrlFormProperty> = async (data) => {
    let { url } = data;
    props.handleIsUrlModalClose(url);
  };
  return (
    <div className="adminCrawler__addModal">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h2 className="adminCrawler__addModal--title">Input url</h2>
        <div className="adminCrawler__addModal--line" />
        <label> URL </label>
        <InputGroup className="mb-3">
          <InputGroup.Text>
            <FontAwesomeIcon icon={faLink} fixedWidth />
          </InputGroup.Text>
          <Form.Control
            {...register('url')}
            placeholder="Type url"
            type="text"
            required
          />
        </InputGroup>

        {errors.url && (
          <p className="errorMessage">{errors.url.message}</p>
        )}
        <Button className="w-100 px-4" variant="success" type="submit">
          Continue
        </Button>
      </form>
    </div>
  );
};

export default UrlModal;
