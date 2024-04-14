import { useState } from 'react';
import './App.css';
import axios from 'axios';

function App() {
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [phone, setphone] = useState('');
  const [login, setLogin] = useState('');

  const API_AUTH_SERVER_URI = 'http://localhost:5000';
  const $authApi = axios.create({
    baseURL: API_AUTH_SERVER_URI,
  });

  const API_BOOKING_SERVER_URI = 'http://localhost:3000';
  const $bookingApi = axios.create({
    baseURL: API_BOOKING_SERVER_URI,
  });

  const auth = () => {
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
        localStorage.setItem('_a', res.data.token); /*console.log('Ауф сука!'); **/
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const getInfo = () => {
    $bookingApi
      .get(
        '/bookings/user/me',
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
        console.log('Инфа о пользователе');
      })
      .catch((e) => {
        console.log(e);
        console.log('Ты кто вообще такой то?');
      });
  };

  return (
    <>
      <p>Сервис бронирования</p>
      <input
        type="text"
        value={name}
        placeholder="Имя"
        onChange={(e) => setName(e.target.value)}
      />
      <br />
      <input
        type="text"
        value={login}
        placeholder="Логин"
        onChange={(e) => setLogin(e.target.value)}
      /><br />
      <input
        type="text"
        value={password}
        placeholder="Пароль"
        onChange={(e) => setPassword(e.target.value)}
      />
      
      <br />
      <input
        type="text"
        value={phone}
        placeholder="Телефон"
        onChange={(e) => setphone(e.target.value)}
      />
      <br />
      <button onClick={() => auth()}>Зарегистрироваться</button>
      <br />
      <br />
      <button onClick={() => getInfo()}>
        Посмотреть информацию о пользователе
      </button>
    </>
  );
}

export default App;
