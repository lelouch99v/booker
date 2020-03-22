import React, { useState } from 'react';
import '../styles/login.scss';
import {Container, Row, Col, Form, Button} from 'react-bootstrap';
import axios from 'axios';

const Login = () => {

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleEmailChange = e => {
    setEmail(e.target.value);
  };

  const handlePasswordChange = e => {
    setPassword(e.target.value);
  };

  const submit = async () => {
    const uri = '/api/auth';
    try {
      const res = await axios.post(uri, {
        email: email,
        password: password
      });
      alert(`ログイン成功\nemail: ${res.data.email}\npassword: ${res.data.password}`);
    } catch (err) {
      alert('ログイン失敗');
    }
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
              <Form.Control onChange={handleEmailChange} type='text' name='id' placeholder='e-mail' />
            </Form.Group>
            <Form.Group>
              <Form.Control onChange={handlePasswordChange} type='password' name='password' placeholder='password' />
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