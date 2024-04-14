import { Row, Col } from 'antd';
import Block from '../../shared/ui/Block/Block';
import Button from '../../shared/ui/Button/Button';
import { useNavigate } from 'react-router-dom';
import React from 'react';
import axios from 'axios';

const Profile = () => {
  const navigate = useNavigate();

  const [name, setName] = React.useState('');
  const [login, setLogin] = React.useState('');
  const [phoneNumber, setPhoneNumber] = React.useState('');

  const API_BOOKING_SERVER_URI = 'http://localhost:3000';
  const $bookingApi = axios.create({
    baseURL: API_BOOKING_SERVER_URI,
  });

  const getInfo = () => {
    $bookingApi
      .get('/bookings/user/me', {
        // withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('_a'),
        },
      })
      .then((res) => {
        setName(res.data.profile.name);
        setLogin(res.data.profile.login);
        setPhoneNumber(res.data.profile.phoneNumber);
        console.log('Инфа о пользователе');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  React.useEffect(() => {
    getInfo();
  }, []);

  const logout = () => {
    localStorage.removeItem('_a');
    navigate('/auth');
  };

  return (
    <div>
      <Row gutter={[40, 40]}>
        <Col xl={{ span: 12 }}>
          <Block>
            <Row>
              <Col xl={{ span: 18 }}>
                <h1>Личные данные:</h1>
                <p>Имя: {name}</p>
                <p>Логин: {login}</p>
                <p>Телефон: {phoneNumber}</p>
              </Col>
              <Col xl={{ span: 6 }}>Фото профиля</Col>
            </Row>
            <Row>
              <Col xl={{ span: 12 }}>
                <Button>Мои отзывы</Button>
              </Col>
              <Col xl={{ span: 12 }}>
                <Button callback={() => logout()}>Выйти</Button>
              </Col>
            </Row>
          </Block>
        </Col>
        <Col xl={{ span: 12 }}>
          <Block>123</Block>
        </Col>
      </Row>
    </div>
  );
};

export default Profile;
