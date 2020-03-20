import axios from 'axios';
import '../styles/login.scss';
import {Container, Row, Col, Form, Button} from 'react-bootstrap';

const Login = () => {

  const submit = e => {
    console.log(e);
  };

  return (
    <Container className='login'>
      <Row>
        <Col className='text-center my-3'>
          <h2>ログイン</h2>
        </Col>
      </Row>
      <Row>
        <Col>
          <Form className='col-4 offset-4'>
            <Form.Group>
              <Form.Control type='text' name='id' placeholder='e-mail' />
            </Form.Group>
            <Form.Group>
              <Form.Control type='password' name='password' placeholder='password' />
            </Form.Group>
            <div className='text-center'>
              <Button onClick={submit} className='w-100'>ログイン</Button>
            </div>
          </Form>
        </Col>
      </Row>
    </Container>
  );
};

export default Login;