import React from 'react';
import axios from 'axios';
import { Row, Col } from 'antd';
import Block from '../../shared/ui/Block/Block';
import Button from '../../shared/ui/Button/Button';
import StarIcon from '../../shared/ui/StarIcon/StarIcon';
import { useNavigate } from 'react-router-dom';
//
import dayjs from 'dayjs';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DateTimePicker } from '@mui/x-date-pickers/DateTimePicker';
//
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';

const Bookings = () => {
  const API_SEARCH_SERVER_URI = 'http://127.0.0.1:14881';
  const $searchApi = axios.create({
    baseURL: API_SEARCH_SERVER_URI,
  });

  const [offers, setOffers] = React.useState([]);

  const [start, setStart] = React.useState(dayjs('2024-04-15T00:00'));
  const [end, setEnd] = React.useState(dayjs('2024-04-20T00:00'));
  const [rate, setRate] = React.useState('');

  const navigate = useNavigate();

  const getOffers = () => {
    $searchApi
      .get('/MyBookings?streetId=1&rating=5', {
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
        console.log(res.data);
        setOffers(res.data);
        console.log('Инфа о доступных офферах');
      })
      .catch((e) => {
        console.log(e);
      });
  };

  React.useEffect(() => {
    getOffers();
  }, []);

  return (
    <div>
      <h1>Список доступного жилья</h1>
      <br />
      <Block>
        <Row gutter={[12, 12]}>
          <Col xl={{ span: 10 }}>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DemoContainer components={['DateTimePicker']}>
                <DateTimePicker
                  label="Дата начала"
                  value={start}
                  onChange={(newValue) => setStart(newValue)}
                />
              </DemoContainer>
            </LocalizationProvider>
          </Col>
          <Col xl={{ span: 10 }}>
            <LocalizationProvider dateAdapter={AdapterDayjs}>
              <DemoContainer components={['DateTimePicker']}>
                <DateTimePicker
                  label="Дата окончания"
                  value={end}
                  onChange={(newValue) => setEnd(newValue)}
                />
              </DemoContainer>
            </LocalizationProvider>
          </Col>
          <Col xl={{ span: 2 }}>
            <div style={{ paddingTop: '8px' }}>
              <FormControl fullWidth>
                <InputLabel id="demo-simple-select-label">Рейтинг</InputLabel>
                <Select
                  labelId="demo-simple-select-label"
                  id="demo-simple-select"
                  value={rate}
                  label="Rating"
                  onChange={(e) => setRate(e.target.value)}
                >
                  <MenuItem value={1}>
                    <StarIcon />
                  </MenuItem>
                  <MenuItem value={2}>
                    <StarIcon /> <StarIcon />
                  </MenuItem>
                  <MenuItem value={3}>
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                  </MenuItem>
                  <MenuItem value={4}>
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                  </MenuItem>
                  <MenuItem value={5}>
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                    <StarIcon />
                  </MenuItem>
                </Select>
              </FormControl>
            </div>
          </Col>
          <Col xl={{ span: 2 }}>
            <br />
            <Button>Поиск</Button>
          </Col>
        </Row>
      </Block>
      <br />

      <Row gutter={[12, 12]}>
        {offers.map((offer) => {
          return (
            <Col xl={{ span: 8 }} key={offer.id}>
              <Block>
                <h2>{offer.name}</h2>
                <p>ID Объявления {offer.id}</p>
                <p>{offer.shortDescription}</p>
                <p>{offer.street.name}</p>
                <p>Цена: {offer.cost} рублей в день</p>
                <Button callback={() => navigate(`/booking?id=${offer.id}`)}>
                  Подробнее
                </Button>
              </Block>
            </Col>
          );
        })}
      </Row>
      <br />
    </div>
  );
};

export default Bookings;
