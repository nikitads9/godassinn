import s from './Input.module.css';

const Input = ({ label, value, callback }) => {
  return (
    <>
      <span>{label}</span>
      <input
        className={s.input}
        type="text"
        placeholder={label}
        value={value}
        onChange={callback}
      />
    </>
  );
};

export default Input;
