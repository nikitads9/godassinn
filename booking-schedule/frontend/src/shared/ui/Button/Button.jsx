import s from './Button.module.css';

const Button = ({ callback, children }) => {
  return (
    <button onClick={callback} className={s.button}>
      {children}
    </button>
  );
};

export default Button;
