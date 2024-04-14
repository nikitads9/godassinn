import { Col, Row } from 'antd';
import s from './Footer.module.css';

const Footer = () => {
  return (
    <div className={s.footer}>
      <Row>
        <Col
          xxl={{ span: 18, offset: 3 }}
          xl={{ span: 20, offset: 2 }}
          xs={{ span: 22, offset: 1 }}
        >
          <div className={s.footer_content}>
            <div className={s.footer_logo}>В.</div>
            <div>О нас</div>
            <div>Условия пользования</div>
            <div>Контакты</div>
            <div>Мы в соцсетях</div>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default Footer;
