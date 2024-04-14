import React from 'react';
import { Row, Col } from 'antd';
import Block from '../../shared/ui/Block/Block';
import Input from '../../shared/ui/Input/Input';
import Button from '../../shared/ui/Button/Button';
import { useNavigate } from 'react-router-dom';
import axios from 'axios';

const Auth = () => {
  const [login, setLogin] = React.useState('');
  const [password, setPassword] = React.useState('');
  const navigate = useNavigate();

  const API_AUTH_SERVER_URI = 'http://localhost:5000';
  const $authApi = axios.create({
    baseURL: API_AUTH_SERVER_URI,
  });

  const signIn = () => {
    $authApi
      .get('/auth/sign-in', {
        withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          // 'Access-Control-Allow-Origin': '*',
          // 'Access-Control-Allow-Credentials': 'true',
          Authorization: 'Basic ' + btoa(`${login}:${password}`),
        },
      })
      .then((res) => {
        console.log(res.data);
        localStorage.setItem('_a', res.data.token);
        console.log('Авторизация');
        navigate('/profile');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  return (
    <Row>
      <Col xxl={{ span: 12, offset: 6 }} xl={{ span: 16, offset: 4 }}>
        <Block>
          <div style={{ padding: '0 100px' }}>
            <div style={{ textAlign: 'center' }}>
              <h1>Авторизация</h1>
            </div>
            <br />
            <br />
            <Input
              label="Логин"
              value={login}
              callback={(e) => setLogin(e.target.value)}
            />
            <Input
              label="Пароль"
              value={password}
              callback={(e) => setPassword(e.target.value)}
            />
            <Row>
              <Col xl={{ span: 18 }}>
                <br />
                <Button callback={() => navigate('/')}>
                  Вернуться на главную
                </Button>
              </Col>
              <Col xl={{ span: 6 }}>
                <div style={{ textAlign: 'right' }}>
                  <br />
                  <Button callback={() => signIn()}>Продолжить</Button>
                </div>
              </Col>
            </Row>
            <br />
            <br />
            <br />
          </div>
        </Block>
      </Col>
    </Row>
  );
};

export default Auth;
