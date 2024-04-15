import React from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import Block from '../../shared/ui/Block/Block';
import Button from '../../shared/ui/Button/Button';
import axios from 'axios';
import { Row, Col } from 'antd';
//
import dayjs from 'dayjs';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
import StarIcon from '../../shared/ui/StarIcon/StarIcon';

const Booking = () => {
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();

  const API_SEARCH_SERVER_URI = 'http://127.0.0.1:14881';
  const $searchApi = axios.create({
    baseURL: API_SEARCH_SERVER_URI,
  });

  const API_BOOKING_SERVER_URI = 'http://localhost:3000';
  const $bookingApi = axios.create({
    baseURL: API_BOOKING_SERVER_URI,
  });

  const [offer, setOffer] = React.useState({
    street: {
      city: {},
    },
  });
  const [start, setStart] = React.useState(dayjs('2024-04-15T00:00'));
  const [end, setEnd] = React.useState(dayjs('2024-04-20T00:00'));
  const [intervals, setIntervals] = React.useState([]);

  const getOffer = () => {
    $searchApi
      .get('/MyBookings', {
        // .get(
        // `/MyBookings/time?rangeStart=${start.toJSON()}&rangeEnd=${end.toJSON()}`,
        // {
        withCredentials: false,
        headers: {
          'Content-type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('_a'),
        },
      })
      .then((res) => {
        console.log(
          res.data.filter((item) => item.id === +searchParams.get('id'))
        );
        setOffer(
          res.data.filter((item) => item.id === +searchParams.get('id'))[0]
        );
        console.log('Инфа о доступных офферах');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const bookOffer = () => {
    const formData = {
      endDate: '2024-04-23T00:00:00Z',
      notifyAt: '24h',
      offerID: +searchParams.get('id'),
      startDate: '2024-04-22T00:00:00Z',
    };

    $bookingApi
      .post('/bookings/add', formData, {
        withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          Authorization: 'Bearer ' + localStorage.getItem('_a'),
        },
      })
      .then((res) => {
        console.log(res.data);
        console.log(
          'Бронирование на даты:',
          start.toJSON(),
          ' и ',
          end.toJSON()
        );
        // navigate('/profile');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  const getVacantDates = () => {
    $bookingApi
      .get(`/bookings/${searchParams.get('id')}/get-vacant-dates`, {
        withCredentials: true,
        headers: {
          'Content-type': 'application/json',
          // Authorization: 'Bearer ' + localStorage.getItem('_a'),
        },
      })
      .then((res) => {
        console.log(res.data);
        setIntervals(res.data.intervals);
        console.log('Списко вакантных дат');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  React.useEffect(() => {
    getOffer();
    getVacantDates();
    console.log(searchParams.get('id'));
  }, []);

  return (
    <div>
      <Row gutter={[12, 12]}>
        <Col xl={{ span: 7 }}>
          <Block>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DemoContainer components={['DateTimePicker']}>
                <DateTimePicker
                  label="Дата начала"
                  value={start}
                  onChange={(newValue) => setStart(newValue)}
                />
              </DemoContainer>
            </LocalizationProvider>
            <br />
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DemoContainer components={['DateTimePicker']}>
                <DateTimePicker
                  label="Дата окончания"
                  value={end}
                  onChange={(newValue) => setEnd(newValue)}
                />
              </DemoContainer>
            </LocalizationProvider>
            <br />
            <Button callback={() => bookOffer()}>Забронировать</Button>
          </Block>
        </Col>
        <Col xl={{ span: 17 }}>
          <Block>
            <Row>
              <Col xl={{ span: 16 }}>
                <h1>Объявление</h1>
                <h2>{offer.name}</h2>
                <p>ID Объявления {offer.id}</p>
                <p>
                  Местонахождение: {offer.street.city.name}, {offer.street.name}
                </p>
                <p>Описание: {offer.shortDescription}</p>
                <p>Количество комнат: {offer.bedsCount}</p>
                <p>Цена: {offer.cost}</p>
                <Button callback={() => navigate('/bookings')}>
                  Вернуться ко всем объявлениям
                </Button>
              </Col>
              <Col xl={{ span: 8 }}>
                <h2>Рейтинг</h2>
                {+offer.rating === 1 ? (
                  <p>
                    <StarIcon />
                  </p>
                ) : +offer.rating === 2 ? (
                  <p>
                    <StarIcon />
                    <StarIcon />
                  </p>
                ) : +offer.rating === 3 ? (
                  <p>
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                  </p>
                ) : +offer.rating === 4 ? (
                  <p>
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                  </p>
                ) : +offer.rating === 5 ? (
                  <p>
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                  </p>
                ) : (
                  ''
                )}
              </Col>
            </Row>
          </Block>
        </Col>
        <Col xl={{ span: 7 }}></Col>
        <Col xl={{ span: 17 }}>
          <Block>
            <h2>Список вакантных дат</h2>
            {intervals.map((interval) => {
              return (
                <>
                  <hr />
                  <p>Начало: {interval.start}</p>
                  <p>Конец: {interval.end}</p>
                </>
              );
            })}
          </Block>
        </Col>
      </Row>
    </div>
  );
};

export default Booking;
