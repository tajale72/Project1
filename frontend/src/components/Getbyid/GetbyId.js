import {useRef} from 'react';

const Getbyid = () => {
  const inputRef = useRef(null);

  function handleClick() {
    console.log(inputRef.current.value);
  }

  return (
    <div>
      <input
        ref={inputRef}
        type="text"
        id="message"
        name="message"
      />
      <button onClick={handleClick}>Log message</button>
    </div>
  );
};



export default Getbyid;
