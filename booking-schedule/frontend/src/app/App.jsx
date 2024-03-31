import { useState } from 'react';
import './App.css';
import axios from 'axios';

function App() {
  const [name, setName] = useState('');
  const [password, setPassword] = useState('');
  const [tgName, setTgName] = useState('');

  const API_AUTH_SERVER_URI = 'http://127.0.0.1:5000';
  const $authApi = axios.create({
    baseURL: API_AUTH_SERVER_URI,
  });

  const auth = () => {
    const formData = {
      name: name,
      password: password,
      telegramID: Date.now(),
      telegramNickname: tgName,
    };

    $authApi
      .post('/auth/sign-up', formData, {
        withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          'Access-Control-Allow-Credentials': 'true',
          // Authorization: 'Basic ' + btoa('client:secret'),
        },
      })
      .then((res) => {
        console.log(res.data);
        console.log('Ауф сука!');
      })
      .catch((e) => {
        console.log(e);
        console.log('Не нихуя!');
      });
  };

  return (
    <>
      <p>Сервис бронирования</p>
      {/* <input
        type="text"
        value={token}
        onChange={(e) => setToken(e.target.value)}
      /> */}
      <button onClick={() => auth()}>Зарегистрироваться</button>
    </>
  );
}

export default App;
