import { Col, Row } from 'antd';
import s from './Header.module.css';
import Button from '../../shared/ui/Button/Button';
import { useNavigate } from 'react-router-dom';

const Header = () => {
  const navigate = useNavigate();

  return (
    <div className={s.header}>
      <Row>
        <Col
          xxl={{ span: 18, offset: 3 }}
          xl={{ span: 20, offset: 2 }}
          xs={{ span: 22, offset: 1 }}
        >
          <div className={s.header_content}>
            <div className={s.header_logo}>Booking-service</div>
            {/* <div>456</div> */}
            <div className={s.header_auth}>
              <Button callback={() => navigate('/auth')}>Войти</Button>
              <Button callback={() => navigate('/profile')}>ЛК</Button>
            </div>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default Header;
