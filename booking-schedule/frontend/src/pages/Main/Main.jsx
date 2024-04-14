import React from 'react';
import { Row, Col } from 'antd';
import s from './Main.module.css';
import image from '../../shared/assets/main.png';
import Button from '../../shared/ui/Button/Button';
import Input from '../../shared/ui/Input/Input';
import { Link } from 'react-router-dom';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

const Main = () => {
  const [name, setName] = React.useState('');
  const [phone, setPhone] = React.useState('');
  const [login, setLogin] = React.useState('');
  const [password, setPassword] = React.useState('');

  const navigate = useNavigate();

  const API_AUTH_SERVER_URI = 'http://localhost:5000';
  const $authApi = axios.create({
    baseURL: API_AUTH_SERVER_URI,
  });

  const signUp = () => {
    const formData = {
      name: name,
      password: password,
      login: login,
      phoneNumber: phone,
    };

    $authApi
      .post('/auth/sign-up', formData, {
        withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          // 'Access-Control-Allow-Origin': '*',
          // 'Access-Control-Allow-Credentials': 'true',
          // Authorization: 'Basic ' + btoa('client:secret'),
        },
      })
      .then((res) => {
        console.log(res.data);
        localStorage.setItem('_a', res.data.token);
        console.log('Регистрация');
        navigate('/profile');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  return (
    <Row gutter={[80, 24]}>
      <Col xl={{ span: 12 }}>
        <div className={s.block}>
          <h1>Попробуйте пока это бесплатно</h1>
          <p>
            Уже зарегистрированы? <Link to="/auth">Войти</Link>
          </p>
          <br />
          <hr />
          <br />
          <Input
            label="Имя"
            value={name}
            callback={(e) => setName(e.target.value)}
          />
          <Input
            label="Телефон"
            value={phone}
            callback={(e) => setPhone(e.target.value)}
          />
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
              <input type="checkbox" name="conditions" id="checkbox" />
              <label htmlFor="checkbox">
                Я принимаю <a href="#">Правила пользования сервисом</a> и{' '}
                <a href="#">Политику конфиденциальности</a>
              </label>
            </Col>
            <Col xl={{ span: 6 }}>
              <div style={{ textAlign: 'right' }}>
                <Button callback={() => signUp()}>Продолжить</Button>
              </div>
            </Col>
          </Row>
        </div>
      </Col>
      <Col xl={{ span: 12 }}>
        <img
          style={{ width: '100%', borderRadius: '12px' }}
          src={image}
          alt="Main"
        />
      </Col>
    </Row>
  );
};

export default Main;
