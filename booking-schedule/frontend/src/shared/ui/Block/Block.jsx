import s from './Block.module.css';

const Block = ({ children }) => {
  return <div className={s.block}>{children}</div>;
};

export default Block;
