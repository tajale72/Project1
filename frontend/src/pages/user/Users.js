import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import Progressbar from './progressbar';
import ProgressStatus from './ProgressStatus';

function Users() {
  const { id } = useParams();

  const [user, setUser] = useState([]);

  useEffect(() => {
    axios.get(`http://localhost:8080/user/${id}`).then((res) => {
      setUser(res.data);
    });
  }, []);

  console.log(user);
  return (
    <>

      <div className="h-full w-full flex flex-col mt-32 justify-center items-center">
        {user && (
          <>
          <div className="py-4 inline-block min-w-full sm:px-6 lg:px-8">
          <ProgressStatus/>
         <Progressbar bgcolor="#ff00ff" progress='20'  height={50} />
          </div>
          <div className="w-[700px] h-[200] px-6 py-4 flex shadow-xl rounded-xl justify-center items-center bg-teal-600 mt-16 border-teal-800 border-2">
            <div className="w-5/12 flex flex-col space-y-4">
              <h2 className="text-white font-bold text-3xl border-black border-b-2">
                Name
              </h2>
              <h2 className="text-white font-bold text-3xl border-black border-b-2">
                Email
              </h2>
              <h2 className="text-white font-bold text-3xl border-black border-b-2">
                Phone
              </h2>
            </div>
            <div className="w-7/12 flex flex-col space-y-4  ">
              <h2 className="text-teal-200 font-bold text-3xl border-black border-b-2">
                {user.name}
              </h2>
              <h2 className="text-teal-200 font-bold text-3xl border-black border-b-2">
                {user.email}
              </h2>
              <h2 className="text-teal-200 font-bold text-3xl border-black border-b-2">
                {user.phone}
              </h2>
            </div>
          </div>
          </>
        )}
      </div>
    </>
  );
}

export default Users;
