import React from 'react'

function ProgressStatus() {
      const Status = {
		height: '100px',
		width: '100%',
		backgroundColor: 'light-blue'
	}
	

  return (
    <>
        <div className="w-ful h-16 flex items-center px-20 justify-between" style={Status}>
        <div  className="text-black">Start</div>
        <div className="text-black">On Progress</div>
        <div  className="text-black">Documentation</div>
        <div className=" text-black">Marketting</div>
        </div>
    </>
  )
}

export default ProgressStatus