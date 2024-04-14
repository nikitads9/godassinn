import { Row, Col } from 'antd';
import Block from '../../shared/ui/Block/Block';
import Button from '../../shared/ui/Button/Button';
import Input from '../../shared/ui/Input/Input';
import { useNavigate } from 'react-router-dom';
import React from 'react';
import axios from 'axios';

const Profile = () => {
  const navigate = useNavigate();

  const [name, setName] = React.useState('');
  const [login, setLogin] = React.useState('');
  const [phoneNumber, setPhoneNumber] = React.useState('');

  const [bookings, setBookings] = React.useState([]);

  //
  const [newName, setNewName] = React.useState('');
  const [newPhone, setNewPhone] = React.useState('');
  const [newLogin, setNewLogin] = React.useState('');
  const [newPassword, setNewPassword] = React.useState('');
  //

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

  const getBookings = () => {
    $bookingApi
      .get(
        '/bookings/get-bookings?start=2024-04-15T17%3A43%3A00&end=2024-04-25T17%3A43%3A00',
        {
          // withCredentials: true,
          headers: {
            'Content-type': 'application/json',
            Authorization: 'Bearer ' + localStorage.getItem('_a'),
          },
        }
      )
      .then((res) => {
        console.log(res.data);
        setBookings(res.data.bookings);
        console.log('Инфа о доступных пользователю бронированиях');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const deleteBooking = (bookingId) => {
    $bookingApi
      .delete(`/bookings/${bookingId}/delete`, {
        // withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('_a'),
        },
      })
      .then((res) => {
        console.log(res.data);
        console.log('Бронирование успешно удалено');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  React.useEffect(() => {
    getInfo();
    getBookings();
  }, []);

  const logout = () => {
    localStorage.removeItem('_a');
    navigate('/auth');
  };

  return (
    <div>
      <Row gutter={[24, 24]}>
        <Col xl={{ span: 7 }}>
          <Block>
            <Row>
              <Col xl={{ span: 18 }}>
                <h2>Личные данные:</h2>
                <p>Имя: {name}</p>
                <p>Логин: {login}</p>
                <p>Телефон: {phoneNumber}</p>
              </Col>
              <Col xl={{ span: 6 }}></Col>
            </Row>
            <Row>
              <Col xl={{ span: 24 }}>
                <Button callback={() => navigate('/bookings')}>
                  Перейти к объявлениям
                </Button>
                <br />
                <br />
                <Button callback={() => logout()}>Выйти</Button>
              </Col>
            </Row>
          </Block>
        </Col>
        <Col xl={{ span: 17 }}>
          <Block>
            <h2>Список ваших бронирований:</h2>
            {bookings ? (
              bookings.map((booking) => {
                return (
                  <div key={booking.BookingID}>
                    <hr />
                    <p>Дата начала бронирования: {booking.startDate}</p>
                    <p>Дата окончания бронирования: {booking.endDate}</p>
                    <p>Уведомить через: {booking.notifyAt}</p>
                    <p>ID Объявления: {booking.offerID}</p>
                    <Button callback={() => deleteBooking(booking.BookingID)}>
                      Отменить
                    </Button>
                  </div>
                );
              })
            ) : (
              <p>Нет бронирований</p>
            )}
          </Block>
        </Col>
        <Col xl={{ span: 7 }}>
          <Block>
            <Input
              label="Редактировать Имя"
              value={newName}
              callback={(e) => setNewName(e.target.value)}
            />
            <Input
              label="Редактировать телефон"
              value={newPhone}
              callback={(e) => setNewPhone(e.target.value)}
            />
            <Input
              label="Редактировать Логин"
              value={newLogin}
              callback={(e) => setNewLogin(e.target.value)}
            />
            <Input
              label="Редактировать Пароль"
              value={newPassword}
              callback={(e) => setNewPassword(e.target.value)}
            />
            <Button>Изменить учетные данные</Button>
          </Block>
          <br />
        </Col>
      </Row>
    </div>
  );
};

export default Profile;
